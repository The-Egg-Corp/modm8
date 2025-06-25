package thunderstore

import (
	"encoding/json"
	"fmt"
	"io"
	"modm8/backend/common/fileutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ECOSYSTEM_URL        = "https://raw.githubusercontent.com/ebkr/r2modmanPlus/refs/heads/develop/src/assets/data/ecosystem.json"
	ECOSYSTEM_SCHEMA_URL = "https://raw.githubusercontent.com/ebkr/r2modmanPlus/refs/heads/develop/src/assets/data/ecosystemJsonSchema.json"
)

type R2ModmanEntry struct {
	Name string `json:"name"`
}

type ThunderstoreEcosystem struct {
	Games map[string]struct {
		R2Modman []*R2ModmanEntry `json:"r2modman"`
	} `json:"games"`
}

type ThunderstoreSchema struct {
	validated bool
	ecosystem *ThunderstoreEcosystem
}

func NewThunderstoreSchema() *ThunderstoreSchema {
	schema := &ThunderstoreSchema{}

	_, err := schema.GetEcosystem()
	if err != nil {
		return nil
	}

	return schema
}

// func (schema *ThunderstoreSchema) Validated() bool {
// 	return schema.validated
// }

func (schema *ThunderstoreSchema) GetEcosystem() (*ThunderstoreEcosystem, error) {
	if schema.validated {
		return schema.ecosystem, nil
	}

	json, err := FetchExternalEcosystem()
	if err != nil {
		return nil, err
	}

	ecosys, err := ParseEcosystemJSON(json)
	if err != nil {
		// Fall back to file that the user will usually have cached.
		json, err := FetchFallbackEcosystem()
		if err != nil {
			return nil, err
		}

		ecosys, err = ParseEcosystemJSON(json)
		if err != nil {
			return nil, fmt.Errorf("failed to parse fallback ecosystem.json file:\n%v", err)
		}
	} else {
		ecosysPath, err := GetFallbackEcosystemPath()
		if err != nil {
			return nil, err
		}

		// No error, we can cache the valid JSON into fallback file.
		dir := filepath.Dir(*ecosysPath)
		if err := fileutil.MkDirAll(dir); err == nil {
			_ = fileutil.WriteFile(*ecosysPath, json)
		}
	}

	schema.ecosystem = ecosys
	schema.validated = ecosys != nil // change when schema validation is implemented

	return ecosys, nil
}

func GetFallbackEcosystemPath() (*string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get path to fallback ecosystem.json file:\n%v", err)
	}

	path := filepath.Join(cfgDir, "modm8", "thunderstore", "ecosystem.json")
	return &path, nil
}

// Reads the fallback ecosystem.json file from the user's config directory and returns its contents.
//
// This file is cached from a previous app launch where GetEcosystem() was successful and should usually exist.
func FetchFallbackEcosystem() ([]byte, error) {
	ecosysPath, err := GetFallbackEcosystemPath()
	if err != nil {
		return nil, err
	}

	json, err := os.ReadFile(*ecosysPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse fetched ecosystem and failed to read fallback file %s: %v", *ecosysPath, err)
	}

	return json, nil
}

// Downloads an external ecosystem.json file located at ECOSYSTEM_URL and returns its contents.
func FetchExternalEcosystem() ([]byte, error) {
	res, err := http.Get(ECOSYSTEM_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %s\n%v", ECOSYSTEM_URL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch: %s\nstatus: %d", ECOSYSTEM_URL, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseEcosystemJSON(data []byte) (*ThunderstoreEcosystem, error) {
	// TODO: Before we unmarshal, validate the schema is valid using file at ECOSYSTEM_SCHEMA_URL.
	var parsed ThunderstoreEcosystem
	err := json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ecosystem JSON: %v", err)
	}

	return &parsed, nil
}

package thunderstore

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	ecosys, err := schema.GetEcosystem()
	if err != nil {
		return nil
	}

	schema.ecosystem = ecosys
	return schema
}

func (schema *ThunderstoreSchema) Validated() bool {
	return schema.validated
}

func (schema *ThunderstoreSchema) GetEcosystem() (*ThunderstoreEcosystem, error) {
	if schema.validated {
		return schema.ecosystem, nil
	}

	json, err := FetchEcosystem()
	if err != nil {
		return nil, err
	}

	// TODO: Implement schema validation here using schema from ECOSYSTEM_SCHEMA_URL.

	ecosys, err := ParseEcosystemJSON(*json)
	if err != nil {
		return nil, err
	}

	schema.validated = true
	return ecosys, nil
}

func ParseEcosystemJSON(body io.ReadCloser) (*ThunderstoreEcosystem, error) {
	var parsed ThunderstoreEcosystem
	if err := json.NewDecoder(body).Decode(&parsed); err != nil {
		log.Printf("failed to decode JSON response of ecosystem.json\n%v", err)
		return nil, err
	}

	return &parsed, nil
}

func FetchEcosystem() (*io.ReadCloser, error) {
	resp, err := http.Get(ECOSYSTEM_URL)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("failed to download ecosystem.json from: %s", ECOSYSTEM_URL)
		return nil, err
	}
	defer resp.Body.Close()

	return &resp.Body, nil
}

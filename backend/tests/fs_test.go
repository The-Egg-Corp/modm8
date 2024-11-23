package backend

import (
	"fmt"
	"modm8/backend"
	"modm8/backend/common/profile"
	"testing"
)

func TestGetDirsAtPath(t *testing.T) {
	dirs, err := backend.GetDirsAtPath(profile.GameProfilesPath("Lethal Company"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dirs)
}

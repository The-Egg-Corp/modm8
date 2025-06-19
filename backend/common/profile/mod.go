package profile

import (
	"strings"
)

type ProfileMod struct {
	Author  string `json:"author"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewProfileMod(verFullName string) ProfileMod {
	info := strings.Split(verFullName, "-")

	ver := "latest"
	if len(info) == 3 {
		ver = info[2]
	}

	return ProfileMod{
		Author:  info[0],
		Name:    info[1],
		Version: ver,
	}
}

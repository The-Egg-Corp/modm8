package loaders

import (
	"fmt"
	"modm8/backend/common/fileutil"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
)

const DOORSTOP_VER_FILE_NAME = ".doorstop_version"

// Reads the ".doorstop_version" file in dirPath (usually the game or profile directory) and returns the major version specified within it.
// This file only exists from Unity Doorstop v4 onwards, and should contain a single line representing a semantic version such as "v4.0.0".
//
// Even if the file content is malformed with spaces or other characters before or after, this func will try to find the first line with a valid semantic version.
// If we dont find it or the version is somehow <= 3, defaults and returns 3.
func GetUnityDoorstopVersion(dirPath string) (uint64, error) {
	semverLine, err := fileutil.FindFirstSemverLine(filepath.Join(dirPath, DOORSTOP_VER_FILE_NAME))
	if err != nil {
		return 3, err
	}

	ver, err := semver.StrictNewVersion(*semverLine)
	if err != nil {
		return 3, err
	}

	majorVer := ver.Major()
	if majorVer <= 3 {
		return 3, err
	}

	return majorVer, nil
}

// I'll probably be dead by the time there are >255 loaders.
type ModLoaderType uint8

func GetModLoaderType(loaderName string) ModLoaderType {
	switch loaderName {
	case "bepinex":
		return ModLoaderType(BEPINEX)
	case "lovely":
		return ModLoaderType(LOVELY)
	case "melonloader":
	case "recursive-melonloader":
		return ModLoaderType(MELON)
	}

	return ModLoaderType(0)
}

func (mlt ModLoaderType) Index() int {
	return int(mlt)
}

func (mlt ModLoaderType) Name() string {
	switch mlt {
	case BEPINEX:
		return "BEPINEX"
	case LOVELY:
		return "LOVELY"
	case MELON:
		return "MELON"
	default:
		return "UNKNOWN"
	}
}

const (
	BEPINEX ModLoaderType = iota + 1
	LOVELY
	MELON
	// NORTHSTAR
	// SHIMLOADER
	// GODOT_ML
	// RETURN_OF_MODDING
	// ANCIENT_DUNGEON_VR
)

type LoaderInstructions struct {
	ModdedParams  []string
	VanillaParams []string
}

type IModLoader interface {
	//IsLoaderPackage(fullName string) bool
	GetModLinkPath(profileDir string) string
	GenerateInstructions(profileDir string) (*LoaderInstructions, error)
}

var MOD_LOADERS = map[ModLoaderType]IModLoader{
	BEPINEX: &BepinexModLoader{},
	MELON:   &MelonModLoader{},
	LOVELY:  &LovelyModLoader{},
}

func GetModLoader(loader ModLoaderType) (IModLoader, error) {
	ldr, ok := MOD_LOADERS[loader]
	if !ok {
		return nil, fmt.Errorf("failed to get mod loader. no loader type exists with index %d", loader.Index())
	}

	return ldr, nil
}

func GetLoaderInstructions(loader ModLoaderType, profileDir string) (*LoaderInstructions, error) {
	ldr, err := GetModLoader(loader)
	if err != nil {
		return nil, err
	}

	instructions, err := ldr.GenerateInstructions(profileDir)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}

func GetModLinkPath(loader ModLoaderType, profileDir string) (string, error) {
	ldr, err := GetModLoader(loader)
	if err != nil {
		return "", err
	}

	return ldr.GetModLinkPath(profileDir), nil
}

func IsLoaderPackage(loader ModLoaderType, fullName string) bool {
	switch loader {
	case BEPINEX:
		return strings.HasPrefix(fullName, "BepInEx-BepInExPack")
	case MELON:
		return fullName == "LavaGang-MelonLoader"
	case LOVELY:
		return fullName == "Thunderstore-lovely"
	default:
		return false
	}
}

// The actual loader's mod/pack. (BepInExPack etc.)
// type LoaderPackageInfo struct {
// 	LoaderType         ModLoader
// 	RecommendedVersion *semver.Version
// 	RootDirName        *string
// }

// Note: Since Go can't do overloading, recommendedVer is kwargs - but only the first element will matter.
// func NewLoaderPackageInfo(loaderType ModLoader, rootDirName string, recommendedVer ...*semver.Version) LoaderPackageInfo {
// 	var version *semver.Version
// 	if len(recommendedVer) > 0 {
// 		version = recommendedVer[0]
// 	}

// 	var dir *string
// 	if rootDirName != "" {
// 		dir = &rootDirName
// 	}

// 	return LoaderPackageInfo{
// 		LoaderType:         loaderType,
// 		RecommendedVersion: version,
// 		RootDirName:        dir,
// 	}
// }

// TODO: Finishing implementing and using this where necessary.
// var ModLoaderPackages = map[string]LoaderPackageInfo{
// 	"bbepis-BepInExPack":                                     NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"bbepisTaleSpire-BepInExPack":                            NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"xiaoxiao921-BepInExPack":                                NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"xiaoye97-BepInEx":                                       NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInExPack_Skul":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInExPack_IL2CPP":                             NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInExPack_x86":                                NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInExPack_Thronefall":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInExPack_WizardWithAGun":                     NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx_Wormtown-BepInExPack":                           NewLoaderPackageInfo(BEPINEX, "BepInExPack"),
// 	"BepInEx-BepInEx_Rogue_Tower":                            NewLoaderPackageInfo(BEPINEX, ""),
// 	"BepInEx-BepInExPack_GTFO":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack_GTFO"),
// 	"BepInEx-BepInExPack_Outward":                            NewLoaderPackageInfo(BEPINEX, "BepInExPack_Outward"),
// 	"BepInEx-BepInExPack_H3VR":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack_H3VR"),
// 	"BepInEx-BepInExPack_ROUNDS":                             NewLoaderPackageInfo(BEPINEX, "BepInExPack_ROUNDS"),
// 	"BepInEx-BepInExPack_Muck":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack_Muck"),
// 	"BepInEx-BepInExPack_LLBlaze":                            NewLoaderPackageInfo(BEPINEX, "BepInExPack_LLBlaze"),
// 	"BepInEx-BepInExPack_Timberborn":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_Timberborn"),
// 	"BepInEx-BepInExPack_TABS":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack_TABS"),
// 	"BepInEx-BepInExPack_NASB":                               NewLoaderPackageInfo(BEPINEX, "BepInExPack_NASB"),
// 	"BepInEx-BepInExPack_Inscryption":                        NewLoaderPackageInfo(BEPINEX, "BepInExPack_Inscryption"),
// 	"BepInEx-BepInExPack_Starsand":                           NewLoaderPackageInfo(BEPINEX, "BepInExPack_Starsand"),
// 	"BepInEx-BepInExPack_CaLABP":                             NewLoaderPackageInfo(BEPINEX, "BepInExPack_CaLABP"),
// 	"BepInEx-BepInExPack_PotionCraft":                        NewLoaderPackageInfo(BEPINEX, "BepInExPack_PotionCraft"),
// 	"BepInEx-BepInExPack_NearlyDead":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_NearlyDead"),
// 	"BepInEx-BepInExPack_AGAINST":                            NewLoaderPackageInfo(BEPINEX, "BepInExPack_AGAINST"),
// 	"BepInEx-BepInExPack_HOTDS":                              NewLoaderPackageInfo(BEPINEX, "BepInExPack_HOTDS"),
// 	"BepInEx-BepInExPack_ForTheKing":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_ForTheKing"),
// 	"BepInEx-BepInExPack_Core_Keeper":                        NewLoaderPackageInfo(BEPINEX, "BepInExPack_Core_Keeper"),
// 	"BepInEx-BepInExPack_Peglin":                             NewLoaderPackageInfo(BEPINEX, "BepInExPack_Peglin"),
// 	"BepInEx-BepInExPack_V_Rising":                           NewLoaderPackageInfo(BEPINEX, "BepInExPack_V_Rising"),
// 	"BepInEx-BepInExPack_VTOL_VR":                            NewLoaderPackageInfo(BEPINEX, "BepInExPack_VTOL_VR"),
// 	"BepInEx-BepInExPack_Stacklands":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_Stacklands"),
// 	"BepInEx-BepInExPack_EtG":                                NewLoaderPackageInfo(BEPINEX, "BepInExPack_EtG"),
// 	"BepInEx-BepInExPack_Ravenfield":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_Ravenfield"),
// 	"BepInEx-BepInExPack_Aloft":                              NewLoaderPackageInfo(BEPINEX, "BepInExPack_Aloft"),
// 	"BepInEx-BepInExPack_CultOfTheLamb":                      NewLoaderPackageInfo(BEPINEX, "BepInExPack_CultOfTheLamb"),
// 	"BepInEx-BepInExPack_Chrono_Ark":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_Chrono_Ark"),
// 	"BepInEx-BepInExPack_TromboneChamp":                      NewLoaderPackageInfo(BEPINEX, "BepInExPack_TromboneChamp"),
// 	"BepInEx-BepInExPack_RogueGenesia":                       NewLoaderPackageInfo(BEPINEX, "BepInExPack_RogueGenesia"),
// 	"BepInEx-BepInExPack_AcrossTheObelisk":                   NewLoaderPackageInfo(BEPINEX, "BepInExPack_AcrossTheObelisk"),
// 	"BepInExPackMTD-BepInExPack_20MTD":                       NewLoaderPackageInfo(BEPINEX, "BepInExPack_20MTD"),
// 	"Subnautica_Modding-BepInExPack_Subnautica":              NewLoaderPackageInfo(BEPINEX, "BepInExPack_Subnautica"),
// 	"Subnautica_Modding-BepInExPack_Subnautica_Experimental": NewLoaderPackageInfo(BEPINEX, "BepInExPack_Subnautica_Experimental"),
// 	"Subnautica_Modding-BepInExPack_BelowZero":               NewLoaderPackageInfo(BEPINEX, "BepInExPack_BelowZero"),
// 	"PCVR_Modders-BepInExPack_GHVR":                          NewLoaderPackageInfo(BEPINEX, "BepInExPack_GHVR"),
// 	"Zinal001-BepInExPack_MECHANICA":                         NewLoaderPackageInfo(BEPINEX, "BepInExPack_MECHANICA"),
// 	"Modding_Council-BepInExPack_of_Legend":                  NewLoaderPackageInfo(BEPINEX, "BepInExPack_of_Legend"),
// 	"SunkenlandModding-BepInExPack_Sunkenland":               NewLoaderPackageInfo(BEPINEX, "BepInExPack_Sunkenland"),
// 	"1F31A-BepInEx_Valheim_Full":                             NewLoaderPackageInfo(BEPINEX, "BepInEx_Valheim_Full"),
// 	"denikson-BepInExPack_Valheim":                           NewLoaderPackageInfo(BEPINEX, "BepInExPack_Valheim"),
// 	"0xFFF7-votv_shimloader":                                 NewLoaderPackageInfo(SHIMLOADER, ""),
// 	"Thunderstore-unreal_shimloader":                         NewLoaderPackageInfo(SHIMLOADER, ""),
// 	"Thunderstore-lovely":                                    NewLoaderPackageInfo(LOVELY, ""),
// 	"ReturnOfModding-ReturnOfModding":                        NewLoaderPackageInfo(RETURN_OF_MODDING, "ReturnOfModdingPack"),
// 	"Hell2Modding-Hell2Modding":                              NewLoaderPackageInfo(RETURN_OF_MODDING, "ReturnOfModdingPack"),
// 	"Northstar-Northstar":                                    NewLoaderPackageInfo(NORTHSTAR, "Northstar"),
// }

package versioning

import (
	"fmt"
	"strings"
)

import "strconv"

var defaultVersion = "0.0.1"

func setVersion(v *Version, versionRaw string) (err error) {

	// @TODO: Verify if we have a leading "v"
	// config.VERSION = strings.Split(config.TAG, "v")[1]

	versionData := strings.Split(versionRaw, ".")

	// fetch and parse major
	if i, err := strconv.Atoi(versionData[0]); err == nil {
		v.major = i
	}
	// fetch and parse minor
	if len(versionData) >= 2 {
		if i, err := strconv.Atoi(versionData[1]); err == nil {
			v.minor = i
		}
	}
	if len(versionData) == 3 {
		// fetch and parse patch
		if i, err := strconv.Atoi(versionData[2]); err == nil {
			v.patch = i
		}
	}

	return
}

// GetVersionString will Generate a version string based on your struct
func getVersionString(v Version) string {

	postfix := ""

	if v.hash != "" {
		postfix = fmt.Sprintf("-%s", v.hash)
	}
	if v.dirty == true {
		postfix = fmt.Sprintf("%s-dirty", postfix)
	}

	return fmt.Sprintf("%d.%d.%d%s", v.major, v.minor, v.patch, postfix)

}

// GetVersionTag will Generate version tag string
func getVersionTag(v Version) string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

// bump value
func bump(value int) int {
	return value + 1
}

// CreateVersion will create a struct and populate this with config data
func CreateVersion(versionStr string) (vrsion Version) {

	if versionStr == "" {
		versionStr = defaultVersion
	}

	vrsion.SetVersion(versionStr)
	return vrsion
}

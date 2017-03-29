package versioning

import "fmt"
import "strings"
import "strconv"

var currentVersion Version
var newVersion Version

// SetVersion will set the correct values of the major, minor and patch attributes of Version struct
func (v *Version) SetVersion(versionRaw string) {

	// glog.Info("Setting Version: ", versionRaw)

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
}

// GetVersionString will Generate a version string based on your struct
func (v *Version) GetVersionString() (version string) {

	postfix := ""

	// If differs from release, append rev
	// if git.DiffSinceRelease(git.GetTagLatest()) == true {
	// 	postfix = fmt.Sprintf("-%s", git.GetRevParse())
	// }
	//
	// // If dirty, append --dirty
	// if git.HasChanges() == true {
	// 	postfix = fmt.Sprintf("%s-dirty", postfix)
	// }
	if v.hash != "" {
		postfix = fmt.Sprintf("-%s", v.hash)
	}
	if v.dirty == true {
		postfix = fmt.Sprintf("%s-dirty", postfix)
	}

	version = fmt.Sprintf("%d.%d.%d%s", v.major, v.minor, v.patch, postfix)

	return
}

// GetVersionTag will Generate version tag string
func (v *Version) GetVersionTag() (version string) {

	version = fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)

	return
}

// Write new JSON to disk
// func (v *Version) Write(configObject config.JSONObject) {
//
// 	configObject.VERSION = v.GetVersionString()
// 	config.Write()
//
// }

// Patch version of given version struct
func (v *Version) Patch() (version string) {
	v.patch = v.patch + 1
	return
}

// Minor version of given version struct
func (v *Version) Minor() (version string) {
	v.minor = v.minor + 1
	return
}

// Major version of given version struct
func (v *Version) Major() (version string) {
	v.major = v.major + 1
	return
}

// CreateVersion will create a struct and populate this with config data
func CreateVersion(versionStr string) (vrsion Version) {
	vrsion.SetVersion(versionStr)
	return vrsion
}

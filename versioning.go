// Copyright (c) 2017 Xebia Nederland B.V. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package versioning is a versioning library that helps generate proper version numbers.
package versioning

import (
	"fmt"
	"strings"
)

import "strconv"

var defaultVersion = "0.0.0"

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
	return fmt.Sprintf("%s%d.%d.%d", v.GetPrefix(), v.major, v.minor, v.patch)
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

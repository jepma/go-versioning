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

// Version struct
type Version struct {
	major, minor, patch int
	hash                string
	dirty               bool
}

// SetVersion will set the correct values of the major, minor and patch attributes of Version struct
func (v *Version) SetVersion(versionRaw string) {
	setVersion(v, versionRaw)
}

// GetVersionString will Generate a version string based on your struct
func (v Version) GetVersionString() (version string) {
	return getVersionString(v)
}

// GetVersionTag will Generate version tag string
func (v Version) GetVersionTag() string {
	return getVersionTag(v)
}

// Patch version of given version struct
func (v *Version) Patch() {
	v.patch = bump(v.patch)
}

// Minor version of given version struct
func (v *Version) Minor() {
	v.minor = bump(v.minor)
	v.patch = 0
}

// Major version of given version struct
func (v *Version) Major() {
	v.major = bump(v.major)
	v.patch = 0
	v.minor = 0
}

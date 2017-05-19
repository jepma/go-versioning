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
	prefix              string
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
	v.SetDirty(false)
	v.SetHash("")
}

// Minor version of given version struct
func (v *Version) Minor() {
	v.minor = bump(v.minor)
	v.patch = 0
	v.SetDirty(false)
	v.SetHash("")
}

// Major version of given version struct
func (v *Version) Major() {
	v.major = bump(v.major)
	v.patch = 0
	v.minor = 0
	v.SetDirty(false)
	v.SetHash("")
}

// SetHash sets the hash value in struct
func (v *Version) SetHash(hash string) {
	v.hash = hash
}

// SetDirty sets the dirty bool in struct
func (v *Version) SetDirty(dirty bool) {
	v.dirty = dirty
}

// SetPrefix sets the dirty bool in struct
func (v *Version) SetPrefix(prefix string) {
	v.prefix = prefix
}

// GetPrefix gets the prefix string in struct
func (v *Version) GetPrefix() (prefix string) {
	return v.prefix
}

// GetDirty gets the dirty bool in struct
func (v *Version) GetDirty() (dirty bool) {
	return v.dirty
}

// GetHash gets the hash string in struct
func (v *Version) GetHash() (hash string) {
	return v.hash
}

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
	"os"
	"testing"
)

func setDir() {
	// Set current work-dir
	os.Chdir("/Workspace/playground/demo-repo")
}

func TestVersionNoDirtyNoHash(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")

	var v string
	v = version.GetVersionString()
	if v != "1.0.0" {
		t.Error("Expected 1.0.0, got ", v)
	}

}

func TestVersionDirtyNoHash(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")
	version.dirty = true

	var v string
	v = version.GetVersionString()
	if v != "1.0.0-dirty" {
		t.Error("Expected 1.0.0-dirty, got ", v)
	}

}

func TestVersionDirtyHash(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")
	version.dirty = true
	version.hash = "2ash5a"

	var v string
	v = version.GetVersionString()
	if v != "1.0.0-2ash5a-dirty" {
		t.Error("Expected 1.0.0-2ash5a-dirty, got ", v)
	}

}

func TestVersionGetTag(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")

	var v string
	v = version.GetVersionTag()
	if v != "v1.0.0" {
		t.Error("Expected v1.0.0, got ", v)
	}

}

func TestVersionMultipleObjects(t *testing.T) {

	// Create version struct
	var version1 Version
	version1.SetVersion("1.0.0")
	version1.Patch()

	var version2 Version
	version2.SetVersion("2.0.0")
	version2.Patch()

	var v1, v2 string
	v2 = version2.GetVersionString()
	v1 = version1.GetVersionString()

	if v1 != "1.0.1" {
		t.Error("Expected v1.0.1, got ", v1)
	}

	if v2 != "2.0.1" {
		t.Error("Expected v2.0.1, got ", v2)
	}

}

func TestVersionPatch(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")
	version.Patch()

	var v string
	v = version.GetVersionString()
	if v != "1.0.1" {
		t.Error("Expected v1.0.1, got ", v)
	}
}

func TestVersionMinor(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")
	version.Minor()

	var v string
	v = version.GetVersionString()
	if v != "1.1.0" {
		t.Error("Expected v1.1.0, got ", v)
	}

}

func TestVersionMajor(t *testing.T) {

	// Create version struct
	var version Version
	version.SetVersion("1.0.0")
	version.Major()

	var v string
	v = version.GetVersionString()
	if v != "2.0.0" {
		t.Error("Expected v2.0.0, got ", v)
	}

}

func TestVersionCreate(t *testing.T) {

	// Create version struct
	version := CreateVersion("1.0.0")

	var v string
	v = version.GetVersionString()
	if v != "1.0.0" {
		t.Error("Expected v1.0.0, got ", v)
	}

}

func TestVersionCreateNoValue(t *testing.T) {

	// Create version struct
	version := CreateVersion("")

	var v string
	v = version.GetVersionString()
	if v != "0.0.1" {
		t.Error("Expected v0.0.1, got ", v)
	}

}

func TestVersionSetHash(t *testing.T) {

	// Create version struct
	version := CreateVersion("1.0.1")
	version.SetHash("hash")

	var v string
	v = version.GetVersionString()
	if v != "1.0.1-hash" {
		t.Error("Expected v1.0.1, got ", v)
	}

}

func TestVersionSetDirty(t *testing.T) {

	// Create version struct
	version := CreateVersion("1.0.1")
	version.SetDirty(true)

	var v string
	v = version.GetVersionString()
	if v != "1.0.1-dirty" {
		t.Error("Expected v1.0.1-dirty, got ", v)
	}

}

func TestVersionSetHashDirty(t *testing.T) {

	// Create version struct
	version := CreateVersion("1.0.1")
	version.SetDirty(true)
	version.SetHash("hash")

	var v string
	v = version.GetVersionString()
	if v != "1.0.1-hash-dirty" {
		t.Error("Expected v1.0.1-hash-dirty, got ", v)
	}

}

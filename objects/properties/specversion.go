// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/nextpart/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
SpecVersionProperty - A property used by all STIX objects that captures the
STIX specification version.
*/
type SpecVersionProperty struct {
	SpecVersion string `json:"spec_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the spec version property on an
object is present if required. It will return a boolean, an integer that tracks
the number of problems found, and a slice of strings that contain the detailed
results, whether good or bad.
*/
func (o *SpecVersionProperty) VerifyExists() error {
	if o.SpecVersion == "" {
		return objects.PropertyMissing("spec_version")
	}

	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *SpecVersionProperty) Compare(obj2 *SpecVersionProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Spec Version Value
	if o.SpecVersion != obj2.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- The spec version values do not match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The spec version values match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

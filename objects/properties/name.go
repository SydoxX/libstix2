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
NameProperty - A property used by one or more STIX objects that captures a
vanity name for the STIX object in string format.
*/
type NameProperty struct {
	Name string `json:"name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Functions - NameProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the name property on an object
is present. It will return a boolean, an integer that tracks the number of
problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *NameProperty) VerifyExists() error {
	if o.Name == "" {
		return objects.PropertyMissing("name")
	}

	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *NameProperty) Compare(obj2 *NameProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Name Value
	if o.Name != obj2.Name {
		problemsFound++
		str := fmt.Sprintf("-- The names do not match: %s | %s", o.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The names match: %s | %s", o.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

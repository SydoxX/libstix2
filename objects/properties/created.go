// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/avast/libstix2/objects"

	timestamp2 "github.com/avast/libstix2/datatypes/timestamp"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
CreatedProperty - A property used by all STIX objects that captures the date
and time that the object was created.
*/
type CreatedProperty struct {
	Created timestamp2.Timestamp `json:"created,omitempty"`
}

/*
VerifyExists - This method will verify that the created property on an object
is present if required. It will return a boolean, an integer that tracks the
number of problems found, and a slice of strings that contain the detailed
results, whether good or bad.
*/
func (o *CreatedProperty) VerifyExists() error {
	if o.Created.IsZero() {
		return objects.PropertyMissing("created")
	}
	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *CreatedProperty) Compare(obj2 *CreatedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Created Value
	if o.Created.Equal(obj2.Created.Time) {
		problemsFound++
		str := fmt.Sprintf("-- The created dates do not match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The created dates match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

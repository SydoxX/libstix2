// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/avast/libstix2/datatypes/timestamp"
	"github.com/avast/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ModifiedProperty - A property used by all STIX objects that captures the date
and time that the object was modified or changed. This property effectively
tracks the version of the object.
*/
type ModifiedProperty struct {
	Modified *timestamp.Timestamp `json:"modified,omitempty"`
}

func (o *ModifiedProperty) VerifyExists() error {
	if o.Modified == nil || o.Modified.IsZero() {
		return objects.PropertyMissing("modified")
	}
	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *ModifiedProperty) Compare(obj2 *ModifiedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Modified Value
	if o.Modified.Equal(obj2.Modified.Time) {
		problemsFound++
		str := fmt.Sprintf("-- The modified dates do not match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The modified dates match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

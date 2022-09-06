// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/avast/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
TypeProperty - A property used by one or more STIX objects that captures the
STIX object type in string format.
*/
type TypeProperty struct {
	ObjectType objects.ObjectType `json:"type"`
}

// ----------------------------------------------------------------------
// Public Methods - TypeProperty - Checks
// ----------------------------------------------------------------------

func (o *TypeProperty) VerifyExists() error {
	if o.ObjectType == "" {
		return objects.PropertyMissing("type")
	}

	if !objects.IsValidObjectType(string(o.ObjectType)) {
		return objects.PropertyInvalid("type", o.ObjectType, "unknown type")
	}
	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *TypeProperty) Compare(obj2 *TypeProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Type Value
	if o.ObjectType != obj2.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- The type values do not match: %s | %s", o.ObjectType, obj2.ObjectType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The type values match: %s | %s", o.ObjectType, obj2.ObjectType)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

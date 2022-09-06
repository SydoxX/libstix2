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
ObjectRefsProperty - A property used by one or more STIX objects.
*/
type ObjectRefsProperty struct {
	ObjectRefs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectRefsProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the object refs property on an
object is present. It will return a boolean, an integer that tracks the number
of problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *ObjectRefsProperty) VerifyExists() error {
	if len(o.ObjectRefs) == 0 {
		return objects.PropertyMissing("object_refs")
	}
	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *ObjectRefsProperty) Compare(obj2 *ObjectRefsProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	if len(o.ObjectRefs) != len(obj2.ObjectRefs) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in object refs do not match: %d | %d", len(o.ObjectRefs), len(obj2.ObjectRefs))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in object refs match: %d | %d", len(o.ObjectRefs), len(obj2.ObjectRefs))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.ObjectRefs {
			if o.ObjectRefs[index] != obj2.ObjectRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- The object ref values do not match: %s | %s", o.ObjectRefs[index], obj2.ObjectRefs[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The object ref values match: %s | %s", o.ObjectRefs[index], obj2.ObjectRefs[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

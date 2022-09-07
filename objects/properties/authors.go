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
AuthorsProperty - A property used by one or more STIX objects.
*/
type AuthorsProperty struct {
	Authors []string `json:"authors,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - AuthorsProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the authors property on an
object is present. It will return a boolean, an integer that tracks the number
of problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *AuthorsProperty) VerifyExists() error {
	if len(o.Authors) == 0 {
		return objects.PropertyMissing("authors")
	}

	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *AuthorsProperty) Compare(obj2 *AuthorsProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	if len(o.Authors) != len(obj2.Authors) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in authors do not match: %d | %d", len(o.Authors), len(obj2.Authors))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in authors match: %d | %d", len(o.Authors), len(obj2.Authors))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.Authors {
			if o.Authors[index] != obj2.Authors[index] {
				problemsFound++
				str := fmt.Sprintf("-- The author values do not match: %s | %s", o.Authors[index], obj2.Authors[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The author values match: %s | %s", o.Authors[index], obj2.Authors[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

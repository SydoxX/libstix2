// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package common

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
ValidSDO - This method will verify and test all of the properties on a STIX
Domain Object to make sure they are valid per the specification. It will return
a boolean, an integer that tracks the number of problems found, and a slice of
strings that contain the detailed results, whether good or bad.
*/
func (o *CommonObjectProperties) ValidSDO() []error {
	var errors []error

	// Verify object Type property is present
	if err := o.TypeProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if err := o.SpecVersionProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if err := o.IDProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if typPrefix := string(o.ObjectType) + "--"; !strings.HasPrefix(o.ID, typPrefix) {
		errors = append(errors, fmt.Errorf("object ID '%s' does not start with type prefix '%s'", o.ID, typPrefix))
	}

	if err := o.CreatedProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if err := o.ModifiedProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if err := o.ExtensionsProperty.Valid(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

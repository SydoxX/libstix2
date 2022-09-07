// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/nextpart/libstix2/datatypes/stixid"

	"github.com/google/uuid"
	"github.com/nextpart/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
IDProperty - A property used by one or more STIX objects that captures the
STIX That identifier in string format.
*/
type IDProperty struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IDProperty - Setters
// ----------------------------------------------------------------------

/*
CreateSTIXUUID - This method takes in a string value representing a STIX
object type and creates and returns a new ID based on the approved STIX UUIDv4
format.
*/
func (o *IDProperty) CreateSTIXUUID(objType objects.ObjectType) (string, error) {
	uuid4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := string(objType) + "--" + uuid4.String()
	return id, nil
}

/*
SetNewSTIXID - This method takes in a string value representing a STIX object
type and creates a new ID based on the approved STIX UUIDv4 format and update
the id property for the object.
*/
func (o *IDProperty) SetNewSTIXID(objType objects.ObjectType) error {
	// TODO Add check to validate input value
	o.ID, _ = o.CreateSTIXUUID(objType)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - IDProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the id property on an object is
present if required. It will return a boolean, an integer that tracks the number
of problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *IDProperty) VerifyExists() error {
	if !stixid.ValidSTIXID(o.ID) {
		return objects.PropertyInvalid("id", o.ID, "invalid id format")
	}
	return nil
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *IDProperty) Compare(obj2 *IDProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check ID Value
	if o.ID != obj2.ID {
		problemsFound++
		str := fmt.Sprintf("-- The id values do not match: %s | %s", o.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The id values match: %s | %s", o.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package stixid

import (
	"fmt"
	"strings"

	"github.com/avast/libstix2/objects"
	"github.com/google/uuid"
)

/*
ValidSTIXID - This function will take in a STIX ID and return true if the
string represents an actual STIX ID in the correct format.
*/
func ValidSTIXID(id string) bool {
	idparts := strings.Split(id, "--")

	if len(idparts) != 2 {
		return false
	}

	// First check to see if the object type is valid, if not return false.
	if valid := objects.IsValidObjectType(idparts[0]); valid == false {
		// Short circuit if the STIX type part is wrong
		return false
	}

	// If the type is valid, then check to see if the ID is a UUID, if not return
	// false.
	if _, err := uuid.Parse(idparts[1]); err != nil {
		return false
	}

	return true
}

func BuildId(objectType string, uuid uuid.UUID) string {
	return fmt.Sprintf("%s--%s", objectType, uuid.String())
}

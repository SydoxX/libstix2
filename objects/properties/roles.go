// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/avast/libstix2/vocabs"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
RolesProperty - A property used by one or more STIX objects that captures a
list of roles that are part of the STIX object.
*/
type RolesProperty struct {
	Roles []vocabs.ThreatActorRole `json:"roles,omitempty"`
}

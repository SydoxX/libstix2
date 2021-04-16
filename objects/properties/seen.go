// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/avast/libstix2/datatypes/timestamp"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
SeenProperties - Properties used by one or more STIX objects that
captures the time that this object was first and last seen in STIX timestamp
format.
*/
type SeenProperties struct {
	FirstSeen *timestamp.Timestamp `json:"first_seen,omitempty"`
	LastSeen  *timestamp.Timestamp `json:"last_seen,omitempty"`
}

func (o *SeenProperties) Valid() error {
	return timestamp.CheckRange(o.FirstSeen, o.LastSeen, "first_seen", "last_seen")
}

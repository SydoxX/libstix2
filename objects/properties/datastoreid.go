// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
DatastoreIDProperty - A property used by all STIX objects that captures the
unique database ID for this object. This is not included in the JSON
serialization, but is used with some datastores.
*/
type DatastoreIDProperty struct {
	DatastoreID int `json:"-"`
}

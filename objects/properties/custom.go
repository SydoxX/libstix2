// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "encoding/json"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
CustomProperties - A property used by all STIX objects that captures any
custom properties. These are all stored in a map.
*/
type CustomProperties struct {
	Custom map[string]*json.RawMessage `json:"custom,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty - Setters
// ----------------------------------------------------------------------

func (c *CustomProperties) SetCustomProperties(custom map[string]*json.RawMessage) {
	c.Custom = custom
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/nextpart/libstix2/vocabs"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ResourceLevelProperty - A property used by one or more STIX objects that
captures the resource level.
*/
type ResourceLevelProperty struct {
	ResourceLevel vocabs.AttackResourceLevel `json:"resource_level,omitempty"`
}

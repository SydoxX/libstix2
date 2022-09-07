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
MotivationProperties - Properties used by one or more STIX objects that
capture the primary and secondary motivations.
*/
type MotivationProperties struct {
	PrimaryMotivation    vocabs.AttackMotivation   `json:"primary_motivation,omitempty"`
	SecondaryMotivations []vocabs.AttackMotivation `json:"secondary_motivations,omitempty"`
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package location

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
	"github.com/nextpart/libstix2/vocabs"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Location - This type implements the STIX 2 Location SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Location struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	Latitude           float64       `json:"latitude,omitempty"`
	Longitude          float64       `json:"longitude,omitempty"`
	Precision          float64       `json:"precision,omitempty"`
	Region             vocabs.Region `json:"region,omitempty"`
	Country            string        `json:"country,omitempty"`
	AdministrativeArea string        `json:"administrative_area,omitempty"`
	City               string        `json:"city,omitempty"`
	StreetAddress      string        `json:"street_address,omitempty"`
	PostalCode         string        `json:"postal_code,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeLocation, func() common.STIXObject {
		return New()
	})
}

func New() *Location {
	var obj Location
	obj.InitSDO(objects.TypeLocation)
	return &obj
}

func (o *Location) Valid() []error {
	return o.CommonObjectProperties.ValidSDO()
}

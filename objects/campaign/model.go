// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package campaign

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Campaign - This type implements the STIX 2 Campaign SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Campaign struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.SeenProperties
	Objective string `json:"objective,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeCampaign, func() common.STIXObject {
		return New()
	})
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Campaign object and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Campaign {
	var obj Campaign
	obj.InitSDO(objects.TypeCampaign)
	return &obj
}

func (o *Campaign) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

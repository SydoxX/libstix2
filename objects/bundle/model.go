// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"
	"fmt"

	"github.com/avast/libstix2/objects/factory"

	"github.com/avast/libstix2/defs"
	"github.com/avast/libstix2/objects"

	"github.com/avast/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Bundle - This type implements the STIX 2 Bundle SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Bundle struct {
	common.CommonObjectProperties
	Objects []common.STIXObject `json:"objects,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeBundle, func() common.STIXObject {
		return New()
	})
}

/*
bundleRawDecode - This type is used for decoding a STIX bundle since the
Objects property needs special handling.
*/
type bundleRawDecode struct {
	common.CommonObjectProperties
	Objects []json.RawMessage `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Bundle object and return it as a
pointer. This function can not use the InitNewObject() function as a Bundle does
not have all of the fields that are common to a standard object.
*/
func New() *Bundle {
	var obj Bundle
	obj.ObjectType = objects.TypeBundle
	obj.SetNewSTIXID(objects.TypeBundle)
	obj.SpecVersion = defs.STIX_VERSION
	return &obj
}

func (o *Bundle) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()
	for _, obj := range o.Objects {
		for _, e := range obj.Valid() {
			errors = append(errors, fmt.Errorf("%s: %w", obj.GetCommonProperties().ID, e))
		}
	}
	return errors
}

func (o *Bundle) AddObject(i common.STIXObject) {
	o.Objects = append(o.Objects, i)
}

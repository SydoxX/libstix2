// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package common

import (
	"encoding/json"
	"fmt"

	"github.com/avast/libstix2/datatypes/timestamp"
	"github.com/avast/libstix2/objects"

	"github.com/avast/libstix2/defs"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
STIXObject - This interface defines what methods an object must have to be
considered a STIX Object. So any new object that is created that inherits the
CommonObjectProperties is considered a STIX Object by this code. This interface
is currently used by the Bundle object to add objects to the Bundle.
*/
type STIXObject interface {
	GetCommonProperties() *CommonObjectProperties
	Valid() []error

	SetRawData(data []byte)
	SetCustomProperties(custom map[string]*json.RawMessage)
}

/*
CommonObjectProperties - This type defines the properties that are common to
most STIX objects. If an object does not use all of these properties, then the
Encode() function for that object will clean up and remove the properties that
might get populated by mistake. Also, there will be Init() functions for each
type of STIX object to help with populating the right properties for that type
of object. This was done so that we would only need one type that could be used
by all objects, to simplify the code.
*/
type CommonObjectProperties struct {
	properties.DatastoreIDProperty
	properties.TypeProperty
	properties.SpecVersionProperty
	properties.IDProperty
	properties.CreatedByRefProperty
	properties.CreatedProperty
	properties.ModifiedProperty
	properties.RevokedProperty
	properties.LabelsProperty
	properties.ConfidenceProperty
	properties.LangProperty
	properties.ExternalReferencesProperty
	properties.MarkingProperties
	properties.CustomProperties
	properties.RawProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitSDO - This method will initialize a STIX Domain Object by setting all
of the basic properties and is called by the New() function from each object.
*/
func (o *CommonObjectProperties) InitSDO(objectType objects.ObjectType) {
	if defs.STRICT_TYPES {
		if valid := objects.IsValidObjectType(string(objectType)); valid != true {
			panic(fmt.Sprintf("invalid object type for InitSDO %s with strict checks enabled", objectType))
		}
	}

	o.SpecVersion = defs.STIX_VERSION
	o.ObjectType = objectType
	o.SetNewSTIXID(objectType)
	o.Created = timestamp.Now()
	o.Modified = o.Created
}

/*
InitSRO - This method will initialize a STIX Relationship Object by setting
all of the basic properties and is called by the New() function from each
object.
*/
func (o *CommonObjectProperties) InitSRO(objectType objects.ObjectType) {
	if defs.STRICT_TYPES {
		if valid := objects.IsValidObjectType(string(objectType)); valid != true {
			panic(fmt.Sprintf("invalid object type for InitSRO %s with strict checks enabled", objectType))
		}
	}

	o.SpecVersion = defs.STIX_VERSION
	o.ObjectType = objectType
	o.SetNewSTIXID(objectType)
	o.Created = timestamp.Now()
	o.Modified = o.Created
}

/*
InitSCO - This method will initialize a STIX Cyber Observable Object by
setting all of the basic properties and is called by the New() function from
each object.
*/
func (o *CommonObjectProperties) InitSCO(objectType objects.ObjectType) {
	if defs.STRICT_TYPES {
		if valid := objects.IsValidObjectType(string(objectType)); valid != true {
			panic(fmt.Sprintf("invalid object type for InitSCO %s with strict checks enabled", objectType))
		}
	}

	o.SpecVersion = defs.STIX_VERSION
	o.ObjectType = objectType
	o.SetNewSTIXID(objectType)
}

/*
GetCommonProperties - This method will return a pointer to the common
properties of this object.
*/
func (o *CommonObjectProperties) GetCommonProperties() *CommonObjectProperties {
	return o
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package common

import (
	"encoding/json"

	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoders
// ----------------------------------------------------------------------

/*
DecodeType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
This is called from the Bundle Decode() to determine which type of STIX object
the data represents, so that the data can be dispatched to the right object
decoder.
*/
func DecodeType(data []byte) (objects.ObjectType, error) {
	var o properties.TypeProperty
	err := json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}

	// This will call the Valid function on the TypeProperty type
	if err := o.VerifyExists(); err != nil {
		return "", err
	}

	return o.ObjectType, nil
}

/*
Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. This is called from the Bundle Decode() if the
object type can not be determined. So for custom objects, it will at least
decode any of the common object properties that might be found.
*/
func Decode(data []byte) (*CommonObjectProperties, error) {
	var o CommonObjectProperties

	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}

// ----------------------------------------------------------------------
// Public Methods JSON Encoders
// The encoding is done here at the individual object level instead of at
// the STIX Object level so that individual pre/post processing rules can
// be applied. Since some of the STIX Objects do not follow a universal
// model, we need to cleanup some things that were inherited but not valid
// for the object.
// ----------------------------------------------------------------------

/*
Encode - This method is a simple wrapper for encoding an object into JSON
*/
func (o *CommonObjectProperties) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object into
JSON
*/
func (o *CommonObjectProperties) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

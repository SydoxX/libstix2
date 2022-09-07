// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/nextpart/libstix2/objects/factory"

	_ "github.com/nextpart/libstix2/objects/factory/importall_except_bundle"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/*
Decode - This function will decode a bundle and return the object as a pointer
along with any errors found.
*/
func Decode(r io.Reader) (*Bundle, []error) {
	allErrors := make([]error, 0)

	var b Bundle
	var rawBundle bundleRawDecode

	// This will decode the outer layer of the bundle and leave all of the
	// objects as a slice of json.rawMessage bytes.
	err := json.NewDecoder(r).Decode(&rawBundle)
	if err != nil {
		// If we can not decode the outer Bundle, we can not do anything so return
		allErrors = append(allErrors, err)
		return nil, allErrors
	}

	// Populate the Common properties
	b.CommonObjectProperties = rawBundle.CommonObjectProperties

	// Loop through all of the raw objects and decode them
	for _, v := range rawBundle.Objects {
		obj, err := factory.Decode(v)
		if err != nil {
			allErrors = append(allErrors, err)
			continue
		}
		b.AddObject(obj)
	}

	return &b, allErrors
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
func (o *Bundle) Encode() ([]byte, error) {
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
func (o *Bundle) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (o *Bundle) UnmarshalJSON(data []byte) error {
	b, err := Decode(bytes.NewReader(data))
	if len(err) != 0 {
		return fmt.Errorf("%v", err)
	}
	*o = *b
	return nil
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
RawProperty - A property used to store the raw bytes of the JSON object.
*/
type RawProperty struct {
	Raw []byte `json:"-"`
}

func (o *RawProperty) SetRawData(data []byte) {
	o.Raw = data
}

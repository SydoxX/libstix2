// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/avast/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ValueProperty -
*/
type ValueProperty struct {
	Value string `json:"value,omitempty"`
}

func (o *ValueProperty) VerifyExists() error {
	if o.Value == "" {
		return objects.PropertyMissing("value")
	}

	return nil
}

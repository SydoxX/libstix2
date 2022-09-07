// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package x509certificate

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
X509Certificate - This type implements the STIX 2 X509Certificate SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type X509Certificate struct {
	common.CommonObjectProperties
	properties.ValueProperty
	IsSelfSigned              bool              `json:"is_self_signed,omitempty" bson:"is_self_signed,omitempty"`
	Hashes                    map[string]string `json:"hashes,omitempty" bson:"hashes,omitempty"`
	Version                   string            `json:"version,omitempty" bson:"version,omitempty"`
	SerialNumber              string            `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	SignatureAlgorithm        string            `json:"signature_algorithm,omitempty" bson:"signature_algorithm,omitempty"`
	Issuer                    string            `json:"issuer,omitempty" bson:"issuer,omitempty"`
	ValidityNotBefore         string            `json:"validity_not_before,omitempty" bson:"validity_not_before,omitempty"`
	ValidityNotAfter          string            `json:"validity_not_after,omitempty" bson:"validity_not_after,omitempty"`
	Subject                   string            `json:"subject,omitempty" bson:"subject,omitempty"`
	SubjectPublicKeyAlgorithm string            `json:"subject_public_key_algorithm,omitempty" bson:"subject_public_key_algorithm,omitempty"`
	SubjectPublicKeyModulus   string            `json:"subject_public_key_modulus,omitempty" bson:"subject_public_key_modulus,omitempty"`
	SubjectPublicKeyExponent  int               `json:"subject_public_key_exponent,omitempty" bson:"subject_public_key_exponent,omitempty"`
	//TODO add X.509 V3 extension
}

func init() {
	factory.RegisterObjectCreator(objects.TypeX509Certificate, func() common.STIXObject {
		return New()
	})
}

func New() *X509Certificate {
	var obj X509Certificate
	obj.InitSCO(objects.TypeX509Certificate)
	return &obj
}

func (o *X509Certificate) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

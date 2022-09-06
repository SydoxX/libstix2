// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package useraccount

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
UserAccount - This type implements the STIX 2 UserAccount SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type UserAccount struct {
	common.CommonObjectProperties
	properties.ExtensionsProperty
	UserId                string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Credential            string `json:"credential,omitempty" bson:"credential,omitempty"`
	AccountLogin          string `json:"account_login,omitempty" bson:"account_login,omitempty"`
	AccountType           string `json:"account_type,omitempty" bson:"account_type,omitempty"`
	DisplayName           string `json:"display_name,omitempty" bson:"display_name,omitempty"`
	IsServiceAccount      bool   `json:"is_service_account,omitempty" bson:"is_service_account,omitempty"`
	IsPrivileged          bool   `json:"is_privileged,omitempty" bson:"is_privileged,omitempty"`
	CanEscalatePrivs      bool   `json:"can_escalate_privs,omitempty" bson:"can_escalate_privs,omitempty"`
	IsDisabled            bool   `json:"is_disabled,omitempty" bson:"is_disabled,omitempty"`
	AccountCreated        string `json:"account_created,omitempty" bson:"account_created,omitempty"`
	AccountExpires        string `json:"account_expires,omitempty" bson:"account_expires,omitempty"`
	CredentialLastChanged string `json:"credential_last_changed,omitempty" bson:"credential_last_changed,omitempty"`
	AccountFirstLogin     string `json:"account_first_login,omitempty" bson:"account_first_login,omitempty"`
	AccountLastLogin      string `json:"account_last_login,omitempty" bson:"account_last_login,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *UserAccount) GetPropertyList() []string {
	return []string{
		"extensions",
		"user_id",
		"credential",
		"account_login",
		"account_type",
		"display_name",
		"is_service_account",
		"is_privileged",
		"can_escalate_privs",
		"is_disabled",
		"account_created",
		"account_expires",
		"credential_last_changed",
		"account_first_login",
		"account_last_login",
	}
}

func init() {
	factory.RegisterObjectCreator(objects.TypeUserAccount, func() common.STIXObject {
		return New()
	})
}

func New() *UserAccount {
	var obj UserAccount
	obj.InitSCO(objects.TypeUserAccount)
	return &obj
}

func (o *UserAccount) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}

// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package emailmessage

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
EmailMessage - This type implements the STIX 2 Email Message SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type EmailMessage struct {
	common.CommonObjectProperties
	IsMultipart            bool           `json:"is_multipart,omitempty" bson:"is_multipart,omitempty"`
	Date                   string         `json:"date,omitempty" bson:"date,omitempty"`
	ContentType            string         `json:"content_type,omitempty" bson:"content_type,omitempty"`
	FromRef                string         `json:"from_ref,omitempty" bson:"from_ref,omitempty"`
	SenderRef              string         `json:"sender_ref,omitempty" bson:"sender_ref,omitempty"`
	ToRefs                 []string       `json:"to_refs,omitempty" bson:"to_refs,omitempty"`
	CcRefs                 []string       `json:"cc_refs,omitempty" bson:"cc_refs,omitempty"`
	BccRefs                []string       `json:"bcc_refs,omitempty" bson:"bcc_refs,omitempty"`
	MessageId              string         `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Subject                string         `json:"subject,omitempty" bson:"subject,omitempty"`
	ReceivedLines          []string       `json:"received_lines,omitempty" bson:"received_lines,omitempty"`
	AdditionalHeaderFields map[string]any `json:"additional_header_fields,omitempty" bson:"additional_header_fields,omitempty"`
	Body                   string         `json:"body,omitempty" bson:"body,omitempty"`
	BodyMultipart          []MimePartType `json:"body_multipart,omitempty" bson:"body_multipart,omitempty"`
	RawEmailRef            string         `json:"raw_email_ref,omitempty" bson:"raw_email_ref,omitempty"`
}

type MimePartType struct {
	Body               string `json:"body,omitempty" bson:"body,omitempty"`
	BodyRawRef         string `json:"body_raw_ref,omitempty" bson:"body_raw_ref,omitempty"`
	ContentType        string `json:"content_type,omitempty" bson:"content_type,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty" bson:"content_disposition,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeEmailMessage, func() common.STIXObject {
		return New()
	})
}

func New() *EmailMessage {
	var obj EmailMessage
	obj.InitSCO(objects.TypeEmailMessage)
	return &obj
}

func (o *EmailMessage) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}

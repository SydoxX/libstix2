// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
)

// ----------------------------------------------------------------------
//
// Private Function
//
// ----------------------------------------------------------------------

/*
sqlAddBaseObject - This function will return an SQL statement that will add the
base object properties to the database.
*/
func sqlAddBaseObject() (string, error) {
	tblBaseObj := datastore.DB_TABLE_STIX_BASE_OBJECT

	/*
		INSERT INTO
			s_base_object (
				"object_id",
				"spec_version",
				"date_added",
				"type",
				"id",
				"created_by_ref",
				"created",
				"modified",
				"revoked",
				"confidence",
				"lang"
			)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblBaseObj)
	s.WriteString(" (")
	s.WriteString("\"object_id\", \"spec_version\", \"date_added\", ")
	s.WriteString("\"type\", \"id\", \"created_by_ref\", \"created\", ")
	s.WriteString("\"modified\", \"revoked\", \"confidence\", \"lang\") ")
	s.WriteString("values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
sqlAddIndicatorObject - This function will return an SQL statement that will add
an indicator to the database.
*/
func sqlAddIndicatorObject() (string, error) {
	tblInd := datastore.DB_TABLE_STIX_INDICATOR

	/*
		INSERT INTO
			s_indicator (
				"object_id",
				"name",
				"description",
				"pattern",
				"valid_from",
				"valid_until"
			)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblInd)
	s.WriteString(" (")
	s.WriteString("\"object_id\", \"name\", \"description\", ")
	s.WriteString("\"pattern\", \"valid_from\", \"valid_until\") ")
	s.WriteString("values (?, ?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
sqlAddObjectLabel - This function will return an SQL statement that will add a
label to the database for a given object.
*/
func sqlAddObjectLabel() (string, error) {
	tblLabels := datastore.DB_TABLE_STIX_LABELS

	/*
		INSERT INTO
			s_labels (
				"object_id",
				"label"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblLabels)
	s.WriteString(" (\"object_id\", \"label\") values (?, ?)")

	return s.String(), nil
}

/*
sqlAddExternalReference - This function will return an SQL statement that will add
an external reference to the database for a given object.
*/
func sqlAddExternalReference() (string, error) {
	tblExtRef := datastore.DB_TABLE_STIX_EXTERNAL_REFERENCES

	/*
		INSERT INTO
			s_external_references (
				"object_id",
				"source_name",
				"description",
				"url",
				"external_id"
			)
			values (?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblExtRef)
	s.WriteString(" (\"object_id\", \"source_name\", \"description\", \"url\", \"external_id\") values (?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
sqlAddObjectMarkingRef - This function will return an SQL statement that will add
an object marking ref to the database for a given object.
*/
func sqlAddObjectMarkingRef() (string, error) {
	tblObjMarking := datastore.DB_TABLE_STIX_OBJECT_MARKING_REFS

	/*
		INSERT INTO
			s_object_marking_refs (
				"object_id",
				"object_marking_refs"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblObjMarking)
	s.WriteString(" (\"object_id\", \"object_marking_refs\") values (?, ?)")

	return s.String(), nil
}

/*
sqlAddKillChainPhase - This function will return an SQL statement that will add a
kill chain phase to the database for a given object.
*/
func sqlAddKillChainPhase() (string, error) {
	tblKillChain := datastore.DB_TABLE_STIX_KILL_CHAIN_PHASES

	/*
		INSERT INTO
			s_kill_chain_phases (
				"object_id",
				"kill_chain_name",
				"phase_name"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblKillChain)
	s.WriteString(" (\"object_id\", \"kill_chain_name\", \"phase_name\") values (?, ?, ?)")

	return s.String(), nil
}

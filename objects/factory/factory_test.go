package factory_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/yudai/gojsondiff/formatter"

	"github.com/avast/libstix2/objects/bundle"
	_ "github.com/avast/libstix2/objects/factory/importall"
	"github.com/gowebpki/jcs"
	"github.com/yudai/gojsondiff"
)

func TestDecode(t *testing.T) {
	expectedNormalized, err := jcs.Transform([]byte(testData))
	if err != nil {
		t.Fatal(err)
	}

	var bund bundle.Bundle
	if err := json.Unmarshal([]byte(testData), &bund); err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(&bund)
	if err != nil {
		t.Fatal(err)
	}

	encodedNormalized, err := jcs.Transform(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expectedNormalized, encodedNormalized) {
		differ := gojsondiff.New()
		diff, err := differ.Compare(expectedNormalized, encodedNormalized)
		if err != nil {
			t.Fatal(err)
		}

		var baseJson map[string]interface{}
		if err := json.Unmarshal(expectedNormalized, &baseJson); err != nil {
			t.Fatal(err)
		}

		diffFmt := formatter.NewAsciiFormatter(baseJson, formatter.AsciiFormatterConfig{
			Coloring: true,
		})

		diffString, err := diffFmt.Format(diff)
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Re-encoded form does not match expected one!\n%s\n", diffString)
	}
}

const testData = `{
    "type": "bundle",
    "spec_version": "2.1",
    "id": "bundle--78cd02ff-61cf-4875-87c5-d93f836d4ddc",
    "objects": [
        {
            "type": "identity",
            "spec_version": "2.1",
            "id": "identity--0040fc3a-d738-54a0-a81f-3df31dad89cc",
            "created_by_ref": "identity--0040fc3a-d738-54a0-a81f-3df31dad89cc",
            "created": "2020-10-20T00:00:00Z",
            "modified": "2020-10-20T00:00:00Z",
            "name": "Avast Software s.r.o.",
            "identity_class": "organization"
        },
        {
            "type": "vulnerability",
            "spec_version": "2.1",
            "id": "vulnerability--e424600d-ffff-ffff-ffff-25ec58547683",
            "created": "2021-04-14T10:31:43.569Z",
            "modified": "2021-04-14T10:31:43.569Z",
            "external_references": [
                {
                    "source_name": "cve",
                    "external_id": "CVE-0000-0000"
                }
            ],
            "name": "CVE-0000-0000"
        },
        {
            "type": "file",
            "spec_version": "2.1",
            "id": "file--d4762fad-ffff-ffff-ffff-99aecdf74b76",
            "hashes": {
                "SHA-256": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
            }
        },
        {
            "type": "file",
            "spec_version": "2.1",
            "id": "file--9b8e2243-ffff-ffff-ffff-e0fdae7e8681",
            "hashes": {
                "SHA-256": "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
            }
        },
        {
            "type": "extension-definition",
            "spec_version": "2.1",
            "id": "extension-definition--645a395b-7349-5196-ad75-2f634c63f80e",
            "created_by_ref": "identity--0040fc3a-d738-54a0-a81f-3df31dad89cc",
            "created": "2021-04-14T09:18:00Z",
            "modified": "2021-04-14T09:18:00Z",
            "name": "Vulnerability Prevalence Extension",
            "description": "A vulnerability prevalence object defines how prevalent a vulnerability is occurring across both geographic and malware instances.",
            "schema": "",
            "version": "1.0",
            "extension_types": [
                "property-extension"
            ],
            "extension_properties": [
                "x_avast_com_vulnerability_ref",
                "x_avast_com_regional_prevalence_map"
            ]
        },
        {
            "type": "observed-data",
            "spec_version": "2.1",
            "id": "observed-data--9a572a8f-eb83-404b-84cf-b8ae3d60fb68",
            "created": "2021-04-14T10:31:43.569Z",
            "modified": "2021-04-14T10:31:43.569Z",
            "first_observed": "2021-04-14T10:31:43.569678Z",
            "last_observed": "2021-04-14T10:31:43.569678Z",
            "number_observed": 1,
            "object_refs": [
                "file--d4762fad-ffff-ffff-ffff-99aecdf74b76",
				"file--9b8e2243-ffff-ffff-ffff-e0fdae7e8681"
            ],
            "extensions": {
                "extension-definition--645a395b-7349-5196-ad75-2f634c63f80e": {
                    "x_avast_com_vulnerability_ref": "vulnerability--e424600d-ffff-ffff-ffff-25ec58547683",
                    "x_avast_com_regional_prevalence_map": {
                        "AA": 2,
						"BB": 3
                    }
                }
            }
        }
    ]
}`

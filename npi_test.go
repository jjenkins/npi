package npi

import (
	"fmt"
	"testing"

	"github.com/parnurzeal/gorequest"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestLookup(t *testing.T) {
	json := `{
		"result_count":1,
		"results":[
			{
				"taxonomies":[
					{
						"state": "CA",
						"code": "208000000X",
						"primary": true,
						"license": "G36822",
						"desc": "Pediatrics"
					}
				],
				"addresses": [
					{
						"city": "WALNUT CREEK",
						"address_2": "SUITE 15",
						"telephone_number": "925-930-8770",
						"fax_number": "925-930-9338",
						"state": "CA",
						"postal_code": "945965279",
						"address_1": "1855 SAN MIGUEL DR",
						"country_code": "US",
						"country_name": "United States",
						"address_type": "DOM",
						"address_purpose": "LOCATION"
					},
					{
						"city": "WALNUT CREEK",
						"address_2": "SUITE 15",
						"telephone_number": "925-930-8770",
						"fax_number": "925-930-9338",
						"state": "CA",
						"postal_code": "945965279",
						"address_1": "1855 SAN MIGUEL DR",
						"country_code": "US",
						"country_name": "United States",
						"address_type": "DOM",
						"address_purpose": "MAILING"
					}
				],
				"created_epoch": 1155600000,
				"identifiers": [],
				"other_names": [],
				"number": 1245243567,
				"last_updated_epoch": 1183852800,
				"basic": {
					"status": "A",
					"credential": "M.D.",
					"first_name": "CHARLES",
					"last_name": "HANSON",
					"middle_name": "HARTMAN",
					"name": "HANSON CHARLES",
					"sole_proprietor": "YES",
					"name_suffix": "X",
					"gender": "M",
					"last_updated": "2007-07-08",
					"name_prefix": "DR.",
					"enumeration_date": "2006-08-15"
				},
				"enumeration_type": "NPI-1"
			}
		]
	}`

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	number := 1245243567
	url := fmt.Sprintf("https://npiregistry.cms.hhs.gov/api?number=%d", number)
	responder := httpmock.NewStringResponder(200, json)
	httpmock.RegisterResponder("GET", url, responder)

	registry := New()
	resp, _ := registry.Lookup(number)

	if resp.Number != number {
		t.Errorf("Expected string %v, got %v", number, resp.Number)
	}
}

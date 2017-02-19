package npi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// RegistryEndpoint is theNational Plan and Provider Enumeration System
// API endpoint.
const RegistryEndpoint = "https://npiregistry.cms.hhs.gov/api"

// RegistryResponse is the response from the NPI Registry
// Public Search
type RegistryResponse struct {
	ResultCount int64  `json:"result_count"`
	Result      []*NPI `json:"results"`
}

// NPI represents a National Provider Identifier issued to
// health care providers in the United States.
type NPI struct {
	Number           int         `json:"number"`
	EnumerationType  string      `json:"enumeration_type"`
	Taxonomies       []*Taxonomy `json:"taxonomies"`
	Addresses        []*Address  `json:"addresses"`
	Provider         *Provider   `json:"basic"`
	CreatedEpoch     int         `json:"created_epoch"`
	LastUpdatedEpoch int         `json:"last_updated_epoch"`
}

// Provider represents the health care provider who was
// issued the NPI number.
type Provider struct {
	Status          string `json:"status"`
	Credential      string `json:"credential"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	MiddleName      string `json:"middle_name"`
	Name            string `json:"name"`
	NameSuffix      string `json:"name_suffix"`
	Gender          string `json:"gender"`
	NamePrefix      string `json:"name_prefix"`
	SoleProprietor  string `json:"sole_proprietor"`
	EnumerationDate string `json:"enumeration_date"`
	LastUpdated     string `json:"last_updated"`
}

// Address represents the location or mailing address of
// the provider.
type Address struct {
	AddressType     string `json:"address_type"`
	AddressPurpose  string `json:"address_purpose"`
	Address1        string `json:"address_1"`
	Address2        string `json:"address_2"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	TelephoneNumber string `json:"telephone_number"`
	FaxNumber       string `json:"fax_number"`
}

// Taxonomy represents the category of services the health
// care provider offers.
type Taxonomy struct {
	State       string `json:"state"`
	Code        string `json:"code"`
	Primary     bool   `json:"primary"`
	License     string `json:"license"`
	Description string `json:"desc"`
}

// Registry is a object storing all the request
// data for the client.
type Registry struct{}

// New used to create a new registry.
func New() *Registry {
	return &Registry{}
}

// Lookup searches the NPI Registry Public Registry for
// a given NPI number. Returns the first match only.
func (r *Registry) Lookup(npi int) (*NPI, error) {
	var rr RegistryResponse

	request := gorequest.New()
	resp, _, errs := request.Get(RegistryEndpoint).
		Query(fmt.Sprintf("number=%d", npi)).EndStruct(&rr)

	if errs != nil {
		log.Fatalf("Unexpected errors: %q", errs)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Couldn't lookup NPI: %s", resp.Status)
	}

	if rr.ResultCount == 0 {
		return nil, fmt.Errorf("Couldn't find NPI recording with %d", npi)
	}

	return rr.Result[0], nil
}

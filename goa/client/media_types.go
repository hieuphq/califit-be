// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "califit-be Backend API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hieuphq/califit-be/goa/design
// --out=$(GOPATH)/src/github.com/hieuphq/califit-be/goa
// --version=v1.4.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// a address detail (default view)
//
// Identifier: application/vnd.address+json; view=default
type Address struct {
	City      *string    `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	Country   *string    `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Lat       *float64   `form:"lat,omitempty" json:"lat,omitempty" yaml:"lat,omitempty" xml:"lat,omitempty"`
	Lng       *float64   `form:"lng,omitempty" json:"lng,omitempty" yaml:"lng,omitempty" xml:"lng,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	State     *string    `form:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty" xml:"state,omitempty"`
	Street1   *string    `form:"street1,omitempty" json:"street1,omitempty" yaml:"street1,omitempty" xml:"street1,omitempty"`
	Street2   *string    `form:"street2,omitempty" json:"street2,omitempty" yaml:"street2,omitempty" xml:"street2,omitempty"`
	ZipCode   *string    `form:"zip_code,omitempty" json:"zip_code,omitempty" yaml:"zip_code,omitempty" xml:"zip_code,omitempty"`
}

// a address detail (general view)
//
// Identifier: application/vnd.address+json; view=general
type AddressGeneral struct {
	City      *string    `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	Country   *string    `form:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty" xml:"country,omitempty"`
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Lat       *float64   `form:"lat,omitempty" json:"lat,omitempty" yaml:"lat,omitempty" xml:"lat,omitempty"`
	Lng       *float64   `form:"lng,omitempty" json:"lng,omitempty" yaml:"lng,omitempty" xml:"lng,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	State     *string    `form:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty" xml:"state,omitempty"`
	Street1   *string    `form:"street1,omitempty" json:"street1,omitempty" yaml:"street1,omitempty" xml:"street1,omitempty"`
	Street2   *string    `form:"street2,omitempty" json:"street2,omitempty" yaml:"street2,omitempty" xml:"street2,omitempty"`
	ZipCode   *string    `form:"zip_code,omitempty" json:"zip_code,omitempty" yaml:"zip_code,omitempty" xml:"zip_code,omitempty"`
}

// DecodeAddress decodes the Address instance encoded in resp body.
func (c *Client) DecodeAddress(resp *http.Response) (*Address, error) {
	var decoded Address
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAddressGeneral decodes the AddressGeneral instance encoded in resp body.
func (c *Client) DecodeAddressGeneral(resp *http.Response) (*AddressGeneral, error) {
	var decoded AddressGeneral
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// a center detail (default view)
//
// Identifier: application/vnd.center+json; view=default
type Center struct {
	Address   *Address   `form:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty" xml:"address,omitempty"`
	AddressID *string    `form:"address_id,omitempty" json:"address_id,omitempty" yaml:"address_id,omitempty" xml:"address_id,omitempty"`
	City      *City      `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	CityID    *string    `form:"city_id,omitempty" json:"city_id,omitempty" yaml:"city_id,omitempty" xml:"city_id,omitempty"`
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// a center detail (general view)
//
// Identifier: application/vnd.center+json; view=general
type CenterGeneral struct {
	Address   *Address   `form:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty" xml:"address,omitempty"`
	AddressID *string    `form:"address_id,omitempty" json:"address_id,omitempty" yaml:"address_id,omitempty" xml:"address_id,omitempty"`
	City      *City      `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	CityID    *string    `form:"city_id,omitempty" json:"city_id,omitempty" yaml:"city_id,omitempty" xml:"city_id,omitempty"`
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// DecodeCenter decodes the Center instance encoded in resp body.
func (c *Client) DecodeCenter(resp *http.Response) (*Center, error) {
	var decoded Center
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCenterGeneral decodes the CenterGeneral instance encoded in resp body.
func (c *Client) DecodeCenterGeneral(resp *http.Response) (*CenterGeneral, error) {
	var decoded CenterGeneral
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A list of Center (default view)
//
// Identifier: application/vnd.centers+json; view=default
type Centers struct {
	Data     []*Center `form:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty" xml:"data,omitempty"`
	Paginate *Paginate `form:"paginate,omitempty" json:"paginate,omitempty" yaml:"paginate,omitempty" xml:"paginate,omitempty"`
}

// DecodeCenters decodes the Centers instance encoded in resp body.
func (c *Client) DecodeCenters(resp *http.Response) (*Centers, error) {
	var decoded Centers
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A list of City (default view)
//
// Identifier: application/vnd.cities+json; view=default
type Cities struct {
	Data     []*City   `form:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty" xml:"data,omitempty"`
	Paginate *Paginate `form:"paginate,omitempty" json:"paginate,omitempty" yaml:"paginate,omitempty" xml:"paginate,omitempty"`
}

// DecodeCities decodes the Cities instance encoded in resp body.
func (c *Client) DecodeCities(resp *http.Response) (*Cities, error) {
	var decoded Cities
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// a city detail (default view)
//
// Identifier: application/vnd.city+json; view=default
type City struct {
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// a city detail (general view)
//
// Identifier: application/vnd.city+json; view=general
type CityGeneral struct {
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty" xml:"created_at,omitempty"`
	ID        *string    `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Name      *string    `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// DecodeCity decodes the City instance encoded in resp body.
func (c *Client) DecodeCity(resp *http.Response) (*City, error) {
	var decoded City
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCityGeneral decodes the CityGeneral instance encoded in resp body.
func (c *Client) DecodeCityGeneral(resp *http.Response) (*CityGeneral, error) {
	var decoded CityGeneral
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Login response (default view)
//
// Identifier: application/vnd.login_response+json; view=default
type LoginResponse struct {
	Token    *string   `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
	UserInfo *UserInfo `form:"user_info,omitempty" json:"user_info,omitempty" yaml:"user_info,omitempty" xml:"user_info,omitempty"`
}

// DecodeLoginResponse decodes the LoginResponse instance encoded in resp body.
func (c *Client) DecodeLoginResponse(resp *http.Response) (*LoginResponse, error) {
	var decoded LoginResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// pagination of a resources (default view)
//
// Identifier: application/vnd.paginate+json; view=default
type Paginate struct {
	CurrentPage *int `form:"current_page,omitempty" json:"current_page,omitempty" yaml:"current_page,omitempty" xml:"current_page,omitempty"`
	TotalItem   *int `form:"total_item,omitempty" json:"total_item,omitempty" yaml:"total_item,omitempty" xml:"total_item,omitempty"`
	TotalPage   *int `form:"total_page,omitempty" json:"total_page,omitempty" yaml:"total_page,omitempty" xml:"total_page,omitempty"`
}

// DecodePaginate decodes the Paginate instance encoded in resp body.
func (c *Client) DecodePaginate(resp *http.Response) (*Paginate, error) {
	var decoded Paginate
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A token (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// A JWT token
	Token *string `form:"token,omitempty" json:"token,omitempty" yaml:"token,omitempty" xml:"token,omitempty"`
}

// DecodeToken decodes the Token instance encoded in resp body.
func (c *Client) DecodeToken(resp *http.Response) (*Token, error) {
	var decoded Token
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User info (default view)
//
// Identifier: application/vnd.user_info+json; view=default
type UserInfo struct {
	CompanyAddress *string `form:"company_address,omitempty" json:"company_address,omitempty" yaml:"company_address,omitempty" xml:"company_address,omitempty"`
	Email          *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Name           *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Phone          *string `form:"phone,omitempty" json:"phone,omitempty" yaml:"phone,omitempty" xml:"phone,omitempty"`
}

// DecodeUserInfo decodes the UserInfo instance encoded in resp body.
func (c *Client) DecodeUserInfo(resp *http.Response) (*UserInfo, error) {
	var decoded UserInfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

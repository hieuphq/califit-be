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
)

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
	CompanyName    *string `form:"company_name,omitempty" json:"company_name,omitempty" yaml:"company_name,omitempty" xml:"company_name,omitempty"`
	ContactName    *string `form:"contact_name,omitempty" json:"contact_name,omitempty" yaml:"contact_name,omitempty" xml:"contact_name,omitempty"`
	Email          *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Phone          *string `form:"phone,omitempty" json:"phone,omitempty" yaml:"phone,omitempty" xml:"phone,omitempty"`
}

// DecodeUserInfo decodes the UserInfo instance encoded in resp body.
func (c *Client) DecodeUserInfo(resp *http.Response) (*UserInfo, error) {
	var decoded UserInfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
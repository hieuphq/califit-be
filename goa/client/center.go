// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "califit-be Backend API": center Resource Client
//
// Command:
// $ goagen
// --design=github.com/hieuphq/califit-be/goa/design
// --out=$(GOPATH)/src/github.com/hieuphq/califit-be/goa
// --version=v1.4.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListCenterPath computes a request path to the list action of center.
func ListCenterPath() string {

	return fmt.Sprintf("/api/center")
}

// List centers
func (c *Client) ListCenter(ctx context.Context, path string, cityID *string, limit *int, name *string, offset *int) (*http.Response, error) {
	req, err := c.NewListCenterRequest(ctx, path, cityID, limit, name, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCenterRequest create the request corresponding to the list action endpoint of the center resource.
func (c *Client) NewListCenterRequest(ctx context.Context, path string, cityID *string, limit *int, name *string, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if cityID != nil {
		values.Set("cityID", *cityID)
	}
	if limit != nil {
		tmp8 := strconv.Itoa(*limit)
		values.Set("limit", tmp8)
	}
	if name != nil {
		values.Set("name", *name)
	}
	if offset != nil {
		tmp9 := strconv.Itoa(*offset)
		values.Set("offset", tmp9)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowCenterPath computes a request path to the show action of center.
func ShowCenterPath(centerID string) string {
	param0 := centerID

	return fmt.Sprintf("/api/center/%s", param0)
}

// Get a center by ID
func (c *Client) ShowCenter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCenterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCenterRequest create the request corresponding to the show action endpoint of the center resource.
func (c *Client) NewShowCenterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
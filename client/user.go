// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Data Table": user Resource Client
//
// Command:
// $ goagen
// --design=github.com/arisgk/goa-api-data-table/design
// --out=$(GOPATH)/src/github.com/arisgk/goa-api-data-table
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	// Age
	Age *int `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// First Name
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last Name
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
}

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {

	return fmt.Sprintf("/api/users")
}

// Create new user
func (c *Client) CreateUser(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *CreateUserPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/users/%s", param0)
}

// DeleteUser makes a request to the delete action endpoint of the user resource
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {

	return fmt.Sprintf("/api/users")
}

// Retrieve all users.
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/users/%s", param0)
}

// Retrieve user with given id.
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	// Age
	Age *int `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// First Name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Last Name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/users/%s", param0)
}

// Update a user
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUserPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user_images_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetMultipleUserImagesV1URL generates an URL for the get multiple user images v1 operation
type GetMultipleUserImagesV1URL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMultipleUserImagesV1URL) WithBasePath(bp string) *GetMultipleUserImagesV1URL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetMultipleUserImagesV1URL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetMultipleUserImagesV1URL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/images/multiple"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetMultipleUserImagesV1URL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetMultipleUserImagesV1URL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetMultipleUserImagesV1URL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetMultipleUserImagesV1URL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetMultipleUserImagesV1URL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetMultipleUserImagesV1URL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
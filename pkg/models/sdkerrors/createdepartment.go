// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type CreateDepartmentErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *CreateDepartmentErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *CreateDepartmentErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// CreateDepartmentDepartmentsResponseResponseBody - Invalid data posted
type CreateDepartmentDepartmentsResponseResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response           `json:"-"`
	Errors      []CreateDepartmentErrors `json:"errors,omitempty"`
	Message     *string                  `json:"message,omitempty"`
}

var _ error = &CreateDepartmentDepartmentsResponseResponseBody{}

func (e *CreateDepartmentDepartmentsResponseResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateDepartmentDepartmentsResponseBody - Forbidden
type CreateDepartmentDepartmentsResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	Message     *string        `json:"message,omitempty"`
}

var _ error = &CreateDepartmentDepartmentsResponseBody{}

func (e *CreateDepartmentDepartmentsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateDepartmentResponseBody - Unauthenticated
type CreateDepartmentResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	Message     *string        `json:"message,omitempty"`
}

var _ error = &CreateDepartmentResponseBody{}

func (e *CreateDepartmentResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

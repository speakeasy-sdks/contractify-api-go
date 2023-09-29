// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type ListUsersRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// The page to retrieve
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *ListUsersRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *ListUsersRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

// ListUsers403ApplicationJSON - Forbidden
type ListUsers403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListUsers403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// ListUsers401ApplicationJSON - Unauthenticated
type ListUsers401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListUsers401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type ListUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UserCollection *shared.UserCollection
	// Unauthenticated
	ListUsers401ApplicationJSONObject *ListUsers401ApplicationJSON
	// Forbidden
	ListUsers403ApplicationJSONObject *ListUsers403ApplicationJSON
}

func (o *ListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUsersResponse) GetUserCollection() *shared.UserCollection {
	if o == nil {
		return nil
	}
	return o.UserCollection
}

func (o *ListUsersResponse) GetListUsers401ApplicationJSONObject() *ListUsers401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListUsers401ApplicationJSONObject
}

func (o *ListUsersResponse) GetListUsers403ApplicationJSONObject() *ListUsers403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListUsers403ApplicationJSONObject
}

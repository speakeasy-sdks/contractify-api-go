// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type ListRelationsRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// The page to retrieve
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Return relations with a certain reference
	Reference *string `queryParam:"style=form,explode=true,name=reference"`
}

func (o *ListRelationsRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *ListRelationsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRelationsRequest) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

// ListRelations403ApplicationJSON - Forbidden
type ListRelations403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListRelations403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// ListRelations401ApplicationJSON - Unauthenticated
type ListRelations401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListRelations401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type ListRelationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	RelationCollection *shared.RelationCollection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthenticated
	ListRelations401ApplicationJSONObject *ListRelations401ApplicationJSON
	// Forbidden
	ListRelations403ApplicationJSONObject *ListRelations403ApplicationJSON
}

func (o *ListRelationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRelationsResponse) GetRelationCollection() *shared.RelationCollection {
	if o == nil {
		return nil
	}
	return o.RelationCollection
}

func (o *ListRelationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRelationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListRelationsResponse) GetListRelations401ApplicationJSONObject() *ListRelations401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListRelations401ApplicationJSONObject
}

func (o *ListRelationsResponse) GetListRelations403ApplicationJSONObject() *ListRelations403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListRelations403ApplicationJSONObject
}

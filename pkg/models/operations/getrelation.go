// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type GetRelationRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the relation
	Relation int64 `pathParam:"style=simple,explode=false,name=relation"`
}

func (o *GetRelationRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *GetRelationRequest) GetRelation() int64 {
	if o == nil {
		return 0
	}
	return o.Relation
}

// GetRelation404ApplicationJSON - Not Found
type GetRelation404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetRelation404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetRelation403ApplicationJSON - Forbidden
type GetRelation403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetRelation403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetRelation401ApplicationJSON - Unauthenticated
type GetRelation401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetRelation401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetRelation200ApplicationJSON - OK
type GetRelation200ApplicationJSON struct {
	Data *shared.RelationRead `json:"data,omitempty"`
}

func (o *GetRelation200ApplicationJSON) GetData() *shared.RelationRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRelationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetRelation200ApplicationJSONObject *GetRelation200ApplicationJSON
	// Unauthenticated
	GetRelation401ApplicationJSONObject *GetRelation401ApplicationJSON
	// Forbidden
	GetRelation403ApplicationJSONObject *GetRelation403ApplicationJSON
	// Not Found
	GetRelation404ApplicationJSONObject *GetRelation404ApplicationJSON
}

func (o *GetRelationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRelationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRelationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRelationResponse) GetGetRelation200ApplicationJSONObject() *GetRelation200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetRelation200ApplicationJSONObject
}

func (o *GetRelationResponse) GetGetRelation401ApplicationJSONObject() *GetRelation401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetRelation401ApplicationJSONObject
}

func (o *GetRelationResponse) GetGetRelation403ApplicationJSONObject() *GetRelation403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetRelation403ApplicationJSONObject
}

func (o *GetRelationResponse) GetGetRelation404ApplicationJSONObject() *GetRelation404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetRelation404ApplicationJSONObject
}

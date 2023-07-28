// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type UpdateRelationSecurity struct {
	OAuth2              string `security:"scheme,type=oauth2,name=Authorization"`
	PersonalAccessToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *UpdateRelationSecurity) GetOAuth2() string {
	if o == nil {
		return ""
	}
	return o.OAuth2
}

func (o *UpdateRelationSecurity) GetPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.PersonalAccessToken
}

type UpdateRelationRequest struct {
	RelationWrite *shared.RelationWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the relation
	Relation int64 `pathParam:"style=simple,explode=false,name=relation"`
}

func (o *UpdateRelationRequest) GetRelationWrite() *shared.RelationWrite {
	if o == nil {
		return nil
	}
	return o.RelationWrite
}

func (o *UpdateRelationRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *UpdateRelationRequest) GetRelation() int64 {
	if o == nil {
		return 0
	}
	return o.Relation
}

type UpdateRelation422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *UpdateRelation422ApplicationJSONErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateRelation422ApplicationJSONErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// UpdateRelation422ApplicationJSON - Invalid data posted
type UpdateRelation422ApplicationJSON struct {
	Errors  []UpdateRelation422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                                  `json:"message,omitempty"`
}

func (o *UpdateRelation422ApplicationJSON) GetErrors() []UpdateRelation422ApplicationJSONErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateRelation422ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateRelation404ApplicationJSON - Not Found
type UpdateRelation404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateRelation404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateRelation403ApplicationJSON - Forbidden
type UpdateRelation403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateRelation403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateRelation401ApplicationJSON - Unauthenticated
type UpdateRelation401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateRelation401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateRelation200ApplicationJSON - OK
type UpdateRelation200ApplicationJSON struct {
	Data *shared.RelationRead `json:"data,omitempty"`
}

func (o *UpdateRelation200ApplicationJSON) GetData() *shared.RelationRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateRelationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UpdateRelation200ApplicationJSONObject *UpdateRelation200ApplicationJSON
	// Unauthenticated
	UpdateRelation401ApplicationJSONObject *UpdateRelation401ApplicationJSON
	// Forbidden
	UpdateRelation403ApplicationJSONObject *UpdateRelation403ApplicationJSON
	// Not Found
	UpdateRelation404ApplicationJSONObject *UpdateRelation404ApplicationJSON
	// Invalid data posted
	UpdateRelation422ApplicationJSONObject *UpdateRelation422ApplicationJSON
}

func (o *UpdateRelationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRelationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRelationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRelationResponse) GetUpdateRelation200ApplicationJSONObject() *UpdateRelation200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRelation200ApplicationJSONObject
}

func (o *UpdateRelationResponse) GetUpdateRelation401ApplicationJSONObject() *UpdateRelation401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRelation401ApplicationJSONObject
}

func (o *UpdateRelationResponse) GetUpdateRelation403ApplicationJSONObject() *UpdateRelation403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRelation403ApplicationJSONObject
}

func (o *UpdateRelationResponse) GetUpdateRelation404ApplicationJSONObject() *UpdateRelation404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRelation404ApplicationJSONObject
}

func (o *UpdateRelationResponse) GetUpdateRelation422ApplicationJSONObject() *UpdateRelation422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRelation422ApplicationJSONObject
}

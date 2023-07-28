// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type GetDocumentSecurity struct {
	OAuth2              string `security:"scheme,type=oauth2,name=Authorization"`
	PersonalAccessToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *GetDocumentSecurity) GetOAuth2() string {
	if o == nil {
		return ""
	}
	return o.OAuth2
}

func (o *GetDocumentSecurity) GetPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.PersonalAccessToken
}

type GetDocumentRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the document
	Document int64 `pathParam:"style=simple,explode=false,name=document"`
}

func (o *GetDocumentRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *GetDocumentRequest) GetDocument() int64 {
	if o == nil {
		return 0
	}
	return o.Document
}

// GetDocument404ApplicationJSON - Not Found
type GetDocument404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetDocument404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetDocument403ApplicationJSON - Forbidden
type GetDocument403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetDocument403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetDocument401ApplicationJSON - Unauthenticated
type GetDocument401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetDocument401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetDocument200ApplicationJSON - OK
type GetDocument200ApplicationJSON struct {
	Data *shared.DocumentRead `json:"data,omitempty"`
}

func (o *GetDocument200ApplicationJSON) GetData() *shared.DocumentRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetDocumentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetDocument200ApplicationJSONObject *GetDocument200ApplicationJSON
	// Unauthenticated
	GetDocument401ApplicationJSONObject *GetDocument401ApplicationJSON
	// Forbidden
	GetDocument403ApplicationJSONObject *GetDocument403ApplicationJSON
	// Not Found
	GetDocument404ApplicationJSONObject *GetDocument404ApplicationJSON
}

func (o *GetDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDocumentResponse) GetGetDocument200ApplicationJSONObject() *GetDocument200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetDocument200ApplicationJSONObject
}

func (o *GetDocumentResponse) GetGetDocument401ApplicationJSONObject() *GetDocument401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetDocument401ApplicationJSONObject
}

func (o *GetDocumentResponse) GetGetDocument403ApplicationJSONObject() *GetDocument403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetDocument403ApplicationJSONObject
}

func (o *GetDocumentResponse) GetGetDocument404ApplicationJSONObject() *GetDocument404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetDocument404ApplicationJSONObject
}

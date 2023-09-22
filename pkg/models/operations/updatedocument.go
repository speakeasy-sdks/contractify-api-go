// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type UpdateDocumentRequest struct {
	DocumentWrite *shared.DocumentWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the document
	Document int64 `pathParam:"style=simple,explode=false,name=document"`
}

func (o *UpdateDocumentRequest) GetDocumentWrite() *shared.DocumentWrite {
	if o == nil {
		return nil
	}
	return o.DocumentWrite
}

func (o *UpdateDocumentRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *UpdateDocumentRequest) GetDocument() int64 {
	if o == nil {
		return 0
	}
	return o.Document
}

// UpdateDocument404ApplicationJSON - Not Found
type UpdateDocument404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDocument404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDocument403ApplicationJSON - Forbidden
type UpdateDocument403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDocument403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDocument401ApplicationJSON - Unauthenticated
type UpdateDocument401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDocument401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDocument200ApplicationJSON - OK
type UpdateDocument200ApplicationJSON struct {
	Data *shared.DocumentRead `json:"data,omitempty"`
}

func (o *UpdateDocument200ApplicationJSON) GetData() *shared.DocumentRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateDocumentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UpdateDocument200ApplicationJSONObject *UpdateDocument200ApplicationJSON
	// Unauthenticated
	UpdateDocument401ApplicationJSONObject *UpdateDocument401ApplicationJSON
	// Forbidden
	UpdateDocument403ApplicationJSONObject *UpdateDocument403ApplicationJSON
	// Not Found
	UpdateDocument404ApplicationJSONObject *UpdateDocument404ApplicationJSON
}

func (o *UpdateDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDocumentResponse) GetUpdateDocument200ApplicationJSONObject() *UpdateDocument200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDocument200ApplicationJSONObject
}

func (o *UpdateDocumentResponse) GetUpdateDocument401ApplicationJSONObject() *UpdateDocument401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDocument401ApplicationJSONObject
}

func (o *UpdateDocumentResponse) GetUpdateDocument403ApplicationJSONObject() *UpdateDocument403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDocument403ApplicationJSONObject
}

func (o *UpdateDocumentResponse) GetUpdateDocument404ApplicationJSONObject() *UpdateDocument404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDocument404ApplicationJSONObject
}

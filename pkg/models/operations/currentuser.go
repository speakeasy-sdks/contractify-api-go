// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

// CurrentUser403ApplicationJSON - Forbidden
type CurrentUser403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CurrentUser403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CurrentUser401ApplicationJSON - Unauthenticated
type CurrentUser401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CurrentUser401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CurrentUser200ApplicationJSON - OK
type CurrentUser200ApplicationJSON struct {
	Data *shared.UserCurrent `json:"data,omitempty"`
}

func (o *CurrentUser200ApplicationJSON) GetData() *shared.UserCurrent {
	if o == nil {
		return nil
	}
	return o.Data
}

type CurrentUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CurrentUser200ApplicationJSONObject *CurrentUser200ApplicationJSON
	// Unauthenticated
	CurrentUser401ApplicationJSONObject *CurrentUser401ApplicationJSON
	// Forbidden
	CurrentUser403ApplicationJSONObject *CurrentUser403ApplicationJSON
}

func (o *CurrentUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CurrentUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CurrentUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CurrentUserResponse) GetCurrentUser200ApplicationJSONObject() *CurrentUser200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CurrentUser200ApplicationJSONObject
}

func (o *CurrentUserResponse) GetCurrentUser401ApplicationJSONObject() *CurrentUser401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CurrentUser401ApplicationJSONObject
}

func (o *CurrentUserResponse) GetCurrentUser403ApplicationJSONObject() *CurrentUser403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CurrentUser403ApplicationJSONObject
}

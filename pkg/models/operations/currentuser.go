// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

// CurrentUserResponseBody - OK
type CurrentUserResponseBody struct {
	Data *shared.UserCurrent `json:"data,omitempty"`
}

func (o *CurrentUserResponseBody) GetData() *shared.UserCurrent {
	if o == nil {
		return nil
	}
	return o.Data
}

type CurrentUserResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CurrentUserResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CurrentUserResponse) GetTwoHundredApplicationJSONObject() *CurrentUserResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
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

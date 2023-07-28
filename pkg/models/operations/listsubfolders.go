// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type ListSubfoldersSecurity struct {
	OAuth2              string `security:"scheme,type=oauth2,name=Authorization"`
	PersonalAccessToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *ListSubfoldersSecurity) GetOAuth2() string {
	if o == nil {
		return ""
	}
	return o.OAuth2
}

func (o *ListSubfoldersSecurity) GetPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.PersonalAccessToken
}

type ListSubfoldersRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

func (o *ListSubfoldersRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

// ListSubfolders403ApplicationJSON - Forbidden
type ListSubfolders403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListSubfolders403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// ListSubfolders401ApplicationJSON - Unauthenticated
type ListSubfolders401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *ListSubfolders401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type ListSubfoldersResponse struct {
	ContentType string
	// OK
	DossierCollection *shared.DossierCollection
	StatusCode        int
	RawResponse       *http.Response
	// Unauthenticated
	ListSubfolders401ApplicationJSONObject *ListSubfolders401ApplicationJSON
	// Forbidden
	ListSubfolders403ApplicationJSONObject *ListSubfolders403ApplicationJSON
}

func (o *ListSubfoldersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSubfoldersResponse) GetDossierCollection() *shared.DossierCollection {
	if o == nil {
		return nil
	}
	return o.DossierCollection
}

func (o *ListSubfoldersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSubfoldersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSubfoldersResponse) GetListSubfolders401ApplicationJSONObject() *ListSubfolders401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListSubfolders401ApplicationJSONObject
}

func (o *ListSubfoldersResponse) GetListSubfolders403ApplicationJSONObject() *ListSubfolders403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListSubfolders403ApplicationJSONObject
}

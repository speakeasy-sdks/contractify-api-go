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

type ListSubfoldersRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

// ListSubfolders403ApplicationJSON - Forbidden
type ListSubfolders403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// ListSubfolders401ApplicationJSON - Unauthenticated
type ListSubfolders401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
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

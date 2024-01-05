// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type ListDepartmentsRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

func (o *ListDepartmentsRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

type ListDepartmentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DepartmentCollection *shared.DepartmentCollection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListDepartmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDepartmentsResponse) GetDepartmentCollection() *shared.DepartmentCollection {
	if o == nil {
		return nil
	}
	return o.DepartmentCollection
}

func (o *ListDepartmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDepartmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

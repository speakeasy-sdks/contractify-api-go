// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package contractifyproduction

import (
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/sdkerrors"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type users struct {
	sdkConfiguration sdkConfiguration
}

func newUsers(sdkConfig sdkConfiguration) *users {
	return &users{
		sdkConfiguration: sdkConfig,
	}
}

// CurrentUser - Current User
// Get the current user details
func (s *users) CurrentUser(ctx context.Context) (*operations.CurrentUserResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/user"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CurrentUserResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.CurrentUser200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CurrentUser200ApplicationJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.CurrentUser401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CurrentUser401ApplicationJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.CurrentUser403ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CurrentUser403ApplicationJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListUsers - List users
// List all the users within a company
func (s *users) ListUsers(ctx context.Context, request operations.ListUsersRequest) (*operations.ListUsersResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/companies/{company}/users", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListUsersResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.UserCollection
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.UserCollection = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.ListUsers401ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ListUsers401ApplicationJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.ListUsers403ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ListUsers403ApplicationJSONObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

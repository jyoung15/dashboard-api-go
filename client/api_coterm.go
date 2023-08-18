/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CotermAPIService CotermAPI service
type CotermAPIService service

type CotermAPIGetOrganizationLicensingCotermLicensesRequest struct {
	ctx context.Context
	ApiService *CotermAPIService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	invalidated *bool
	expired *bool
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) PerPage(perPage int32) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) StartingAfter(startingAfter string) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) EndingBefore(endingBefore string) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Filter for licenses that are invalidated
func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) Invalidated(invalidated bool) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	r.invalidated = &invalidated
	return r
}

// Filter for licenses that are expired
func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) Expired(expired bool) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	r.expired = &expired
	return r
}

func (r CotermAPIGetOrganizationLicensingCotermLicensesRequest) Execute() ([]GetOrganizationLicensingCotermLicenses200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationLicensingCotermLicensesExecute(r)
}

/*
GetOrganizationLicensingCotermLicenses List the licenses in a coterm organization

List the licenses in a coterm organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CotermAPIGetOrganizationLicensingCotermLicensesRequest
*/
func (a *CotermAPIService) GetOrganizationLicensingCotermLicenses(ctx context.Context, organizationId string) CotermAPIGetOrganizationLicensingCotermLicensesRequest {
	return CotermAPIGetOrganizationLicensingCotermLicensesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationLicensingCotermLicenses200ResponseInner
func (a *CotermAPIService) GetOrganizationLicensingCotermLicensesExecute(r CotermAPIGetOrganizationLicensingCotermLicensesRequest) ([]GetOrganizationLicensingCotermLicenses200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationLicensingCotermLicenses200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CotermAPIService.GetOrganizationLicensingCotermLicenses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/licensing/coterm/licenses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.invalidated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invalidated", r.invalidated, "")
	}
	if r.expired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expired", r.expired, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CotermAPIMoveOrganizationLicensingCotermLicensesRequest struct {
	ctx context.Context
	ApiService *CotermAPIService
	organizationId string
	moveOrganizationLicensingCotermLicensesRequest *MoveOrganizationLicensingCotermLicensesRequest
}

func (r CotermAPIMoveOrganizationLicensingCotermLicensesRequest) MoveOrganizationLicensingCotermLicensesRequest(moveOrganizationLicensingCotermLicensesRequest MoveOrganizationLicensingCotermLicensesRequest) CotermAPIMoveOrganizationLicensingCotermLicensesRequest {
	r.moveOrganizationLicensingCotermLicensesRequest = &moveOrganizationLicensingCotermLicensesRequest
	return r
}

func (r CotermAPIMoveOrganizationLicensingCotermLicensesRequest) Execute() (*MoveOrganizationLicensingCotermLicenses200Response, *http.Response, error) {
	return r.ApiService.MoveOrganizationLicensingCotermLicensesExecute(r)
}

/*
MoveOrganizationLicensingCotermLicenses Moves a license to a different organization (coterm only)

Moves a license to a different organization (coterm only)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CotermAPIMoveOrganizationLicensingCotermLicensesRequest
*/
func (a *CotermAPIService) MoveOrganizationLicensingCotermLicenses(ctx context.Context, organizationId string) CotermAPIMoveOrganizationLicensingCotermLicensesRequest {
	return CotermAPIMoveOrganizationLicensingCotermLicensesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return MoveOrganizationLicensingCotermLicenses200Response
func (a *CotermAPIService) MoveOrganizationLicensingCotermLicensesExecute(r CotermAPIMoveOrganizationLicensingCotermLicensesRequest) (*MoveOrganizationLicensingCotermLicenses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoveOrganizationLicensingCotermLicenses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CotermAPIService.MoveOrganizationLicensingCotermLicenses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/licensing/coterm/licenses/move"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.moveOrganizationLicensingCotermLicensesRequest == nil {
		return localVarReturnValue, nil, reportError("moveOrganizationLicensingCotermLicensesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.moveOrganizationLicensingCotermLicensesRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

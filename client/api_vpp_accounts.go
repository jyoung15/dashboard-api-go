/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
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


// VppAccountsAPIService VppAccountsAPI service
type VppAccountsAPIService service

type VppAccountsAPIGetOrganizationSmVppAccountRequest struct {
	ctx context.Context
	ApiService *VppAccountsAPIService
	organizationId string
	vppAccountId string
}

func (r VppAccountsAPIGetOrganizationSmVppAccountRequest) Execute() (*GetOrganizationSmVppAccounts200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSmVppAccountExecute(r)
}

/*
GetOrganizationSmVppAccount Get a hash containing the unparsed token of the VPP account with the given ID

Get a hash containing the unparsed token of the VPP account with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param vppAccountId Vpp account ID
 @return VppAccountsAPIGetOrganizationSmVppAccountRequest
*/
func (a *VppAccountsAPIService) GetOrganizationSmVppAccount(ctx context.Context, organizationId string, vppAccountId string) VppAccountsAPIGetOrganizationSmVppAccountRequest {
	return VppAccountsAPIGetOrganizationSmVppAccountRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		vppAccountId: vppAccountId,
	}
}

// Execute executes the request
//  @return GetOrganizationSmVppAccounts200ResponseInner
func (a *VppAccountsAPIService) GetOrganizationSmVppAccountExecute(r VppAccountsAPIGetOrganizationSmVppAccountRequest) (*GetOrganizationSmVppAccounts200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationSmVppAccounts200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VppAccountsAPIService.GetOrganizationSmVppAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/vppAccounts/{vppAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vppAccountId"+"}", url.PathEscape(parameterValueToString(r.vppAccountId, "vppAccountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type VppAccountsAPIGetOrganizationSmVppAccountsRequest struct {
	ctx context.Context
	ApiService *VppAccountsAPIService
	organizationId string
}

func (r VppAccountsAPIGetOrganizationSmVppAccountsRequest) Execute() ([]GetOrganizationSmVppAccounts200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSmVppAccountsExecute(r)
}

/*
GetOrganizationSmVppAccounts List the VPP accounts in the organization

List the VPP accounts in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return VppAccountsAPIGetOrganizationSmVppAccountsRequest
*/
func (a *VppAccountsAPIService) GetOrganizationSmVppAccounts(ctx context.Context, organizationId string) VppAccountsAPIGetOrganizationSmVppAccountsRequest {
	return VppAccountsAPIGetOrganizationSmVppAccountsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSmVppAccounts200ResponseInner
func (a *VppAccountsAPIService) GetOrganizationSmVppAccountsExecute(r VppAccountsAPIGetOrganizationSmVppAccountsRequest) ([]GetOrganizationSmVppAccounts200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSmVppAccounts200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VppAccountsAPIService.GetOrganizationSmVppAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/vppAccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

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


// DelegatedAPIService DelegatedAPI service
type DelegatedAPIService service

type DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	networkId string
	createNetworkAppliancePrefixesDelegatedStaticRequest *CreateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest) CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest CreateNetworkAppliancePrefixesDelegatedStaticRequest) DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.createNetworkAppliancePrefixesDelegatedStaticRequest = &createNetworkAppliancePrefixesDelegatedStaticRequest
	return r
}

func (r DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
CreateNetworkAppliancePrefixesDelegatedStatic Add a static delegated prefix from a network

Add a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedAPIService) CreateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string) DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DelegatedAPIService) CreateNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedAPICreateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.CreateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkAppliancePrefixesDelegatedStaticRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkAppliancePrefixesDelegatedStaticRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkAppliancePrefixesDelegatedStaticRequest
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

type DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	networkId string
	staticDelegatedPrefixId string
}

func (r DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
DeleteNetworkAppliancePrefixesDelegatedStatic Delete a static delegated prefix from a network

Delete a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param staticDelegatedPrefixId Static delegated prefix ID
 @return DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedAPIService) DeleteNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
func (a *DelegatedAPIService) DeleteNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedAPIDeleteNetworkAppliancePrefixesDelegatedStaticRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.DeleteNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	serial string
}

func (r DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegated Return current delegated IPv6 prefixes on an appliance.

Return current delegated IPv6 prefixes on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest
*/
func (a *DelegatedAPIService) GetDeviceAppliancePrefixesDelegated(ctx context.Context, serial string) DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest {
	return DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DelegatedAPIService) GetDeviceAppliancePrefixesDelegatedExecute(r DelegatedAPIGetDeviceAppliancePrefixesDelegatedRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.GetDeviceAppliancePrefixesDelegated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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

type DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	serial string
}

func (r DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r)
}

/*
GetDeviceAppliancePrefixesDelegatedVlanAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest
*/
func (a *DelegatedAPIService) GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx context.Context, serial string) DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest {
	return DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *DelegatedAPIService) GetDeviceAppliancePrefixesDelegatedVlanAssignmentsExecute(r DelegatedAPIGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.GetDeviceAppliancePrefixesDelegatedVlanAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/appliance/prefixes/delegated/vlanAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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

type DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	networkId string
	staticDelegatedPrefixId string
}

func (r DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatic Return a static delegated prefix from a network

Return a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param staticDelegatedPrefixId Static delegated prefix ID
 @return DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedAPIService) GetNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *DelegatedAPIService) GetNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticRequest) (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.GetNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

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

type DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	networkId string
}

func (r DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest) Execute() ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkAppliancePrefixesDelegatedStaticsExecute(r)
}

/*
GetNetworkAppliancePrefixesDelegatedStatics List static delegated prefixes for a network

List static delegated prefixes for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest
*/
func (a *DelegatedAPIService) GetNetworkAppliancePrefixesDelegatedStatics(ctx context.Context, networkId string) DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest {
	return DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
func (a *DelegatedAPIService) GetNetworkAppliancePrefixesDelegatedStaticsExecute(r DelegatedAPIGetNetworkAppliancePrefixesDelegatedStaticsRequest) ([]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.GetNetworkAppliancePrefixesDelegatedStatics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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

type DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	ctx context.Context
	ApiService *DelegatedAPIService
	networkId string
	staticDelegatedPrefixId string
	updateNetworkAppliancePrefixesDelegatedStaticRequest *UpdateNetworkAppliancePrefixesDelegatedStaticRequest
}

func (r DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest) UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest UpdateNetworkAppliancePrefixesDelegatedStaticRequest) DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	r.updateNetworkAppliancePrefixesDelegatedStaticRequest = &updateNetworkAppliancePrefixesDelegatedStaticRequest
	return r
}

func (r DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r)
}

/*
UpdateNetworkAppliancePrefixesDelegatedStatic Update a static delegated prefix from a network

Update a static delegated prefix from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param staticDelegatedPrefixId Static delegated prefix ID
 @return DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest
*/
func (a *DelegatedAPIService) UpdateNetworkAppliancePrefixesDelegatedStatic(ctx context.Context, networkId string, staticDelegatedPrefixId string) DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	return DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		staticDelegatedPrefixId: staticDelegatedPrefixId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DelegatedAPIService) UpdateNetworkAppliancePrefixesDelegatedStaticExecute(r DelegatedAPIUpdateNetworkAppliancePrefixesDelegatedStaticRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatedAPIService.UpdateNetworkAppliancePrefixesDelegatedStatic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"staticDelegatedPrefixId"+"}", url.PathEscape(parameterValueToString(r.staticDelegatedPrefixId, "staticDelegatedPrefixId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateNetworkAppliancePrefixesDelegatedStaticRequest
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

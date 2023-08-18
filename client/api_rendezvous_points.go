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


// RendezvousPointsAPIService RendezvousPointsAPI service
type RendezvousPointsAPIService service

type RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsAPIService
	networkId string
	createNetworkSwitchRoutingMulticastRendezvousPointRequest *CreateNetworkSwitchRoutingMulticastRendezvousPointRequest
}

func (r RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest) CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest CreateNetworkSwitchRoutingMulticastRendezvousPointRequest) RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.createNetworkSwitchRoutingMulticastRendezvousPointRequest = &createNetworkSwitchRoutingMulticastRendezvousPointRequest
	return r
}

func (r RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
CreateNetworkSwitchRoutingMulticastRendezvousPoint Create a multicast rendezvous point

Create a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsAPIService) CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string) RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsAPIService) CreateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsAPICreateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsAPIService.CreateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkSwitchRoutingMulticastRendezvousPointRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchRoutingMulticastRendezvousPointRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchRoutingMulticastRendezvousPointRequest
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

type RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsAPIService
	networkId string
	rendezvousPointId string
}

func (r RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
DeleteNetworkSwitchRoutingMulticastRendezvousPoint Delete a multicast rendezvous point

Delete a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param rendezvousPointId Rendezvous point ID
 @return RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsAPIService) DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
func (a *RendezvousPointsAPIService) DeleteNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsAPIDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsAPIService.DeleteNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

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

type RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsAPIService
	networkId string
	rendezvousPointId string
}

func (r RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoint Return a multicast rendezvous point

Return a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param rendezvousPointId Rendezvous point ID
 @return RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsAPIService) GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsAPIService) GetNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsAPIService.GetNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

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

type RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsAPIService
	networkId string
}

func (r RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) Execute() ([][]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r)
}

/*
GetNetworkSwitchRoutingMulticastRendezvousPoints List multicast rendezvous points

List multicast rendezvous points

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest
*/
func (a *RendezvousPointsAPIService) GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx context.Context, networkId string) RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest {
	return RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return [][]map[string]interface{}
func (a *RendezvousPointsAPIService) GetNetworkSwitchRoutingMulticastRendezvousPointsExecute(r RendezvousPointsAPIGetNetworkSwitchRoutingMulticastRendezvousPointsRequest) ([][]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  [][]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsAPIService.GetNetworkSwitchRoutingMulticastRendezvousPoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints"
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

type RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	ctx context.Context
	ApiService *RendezvousPointsAPIService
	networkId string
	rendezvousPointId string
	updateNetworkSwitchRoutingMulticastRendezvousPointRequest *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
}

func (r RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest(updateNetworkSwitchRoutingMulticastRendezvousPointRequest UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	r.updateNetworkSwitchRoutingMulticastRendezvousPointRequest = &updateNetworkSwitchRoutingMulticastRendezvousPointRequest
	return r
}

func (r RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r)
}

/*
UpdateNetworkSwitchRoutingMulticastRendezvousPoint Update a multicast rendezvous point

Update a multicast rendezvous point

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param rendezvousPointId Rendezvous point ID
 @return RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
*/
func (a *RendezvousPointsAPIService) UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx context.Context, networkId string, rendezvousPointId string) RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		rendezvousPointId: rendezvousPointId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RendezvousPointsAPIService) UpdateNetworkSwitchRoutingMulticastRendezvousPointExecute(r RendezvousPointsAPIUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RendezvousPointsAPIService.UpdateNetworkSwitchRoutingMulticastRendezvousPoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rendezvousPointId"+"}", url.PathEscape(parameterValueToString(r.rendezvousPointId, "rendezvousPointId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkSwitchRoutingMulticastRendezvousPointRequest == nil {
		return localVarReturnValue, nil, reportError("updateNetworkSwitchRoutingMulticastRendezvousPointRequest is required and must be specified")
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
	localVarPostBody = r.updateNetworkSwitchRoutingMulticastRendezvousPointRequest
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

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


// TargetGroupsAPIService TargetGroupsAPI service
type TargetGroupsAPIService service

type TargetGroupsAPICreateNetworkSmTargetGroupRequest struct {
	ctx context.Context
	ApiService *TargetGroupsAPIService
	networkId string
	createNetworkSmTargetGroupRequest *CreateNetworkSmTargetGroupRequest
}

func (r TargetGroupsAPICreateNetworkSmTargetGroupRequest) CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest CreateNetworkSmTargetGroupRequest) TargetGroupsAPICreateNetworkSmTargetGroupRequest {
	r.createNetworkSmTargetGroupRequest = &createNetworkSmTargetGroupRequest
	return r
}

func (r TargetGroupsAPICreateNetworkSmTargetGroupRequest) Execute() (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkSmTargetGroupExecute(r)
}

/*
CreateNetworkSmTargetGroup Add a target group

Add a target group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return TargetGroupsAPICreateNetworkSmTargetGroupRequest
*/
func (a *TargetGroupsAPIService) CreateNetworkSmTargetGroup(ctx context.Context, networkId string) TargetGroupsAPICreateNetworkSmTargetGroupRequest {
	return TargetGroupsAPICreateNetworkSmTargetGroupRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSmTargetGroups200ResponseInner
func (a *TargetGroupsAPIService) CreateNetworkSmTargetGroupExecute(r TargetGroupsAPICreateNetworkSmTargetGroupRequest) (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSmTargetGroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsAPIService.CreateNetworkSmTargetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/targetGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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
	localVarPostBody = r.createNetworkSmTargetGroupRequest
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

type TargetGroupsAPIDeleteNetworkSmTargetGroupRequest struct {
	ctx context.Context
	ApiService *TargetGroupsAPIService
	networkId string
	targetGroupId string
}

func (r TargetGroupsAPIDeleteNetworkSmTargetGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSmTargetGroupExecute(r)
}

/*
DeleteNetworkSmTargetGroup Delete a target group from a network

Delete a target group from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param targetGroupId Target group ID
 @return TargetGroupsAPIDeleteNetworkSmTargetGroupRequest
*/
func (a *TargetGroupsAPIService) DeleteNetworkSmTargetGroup(ctx context.Context, networkId string, targetGroupId string) TargetGroupsAPIDeleteNetworkSmTargetGroupRequest {
	return TargetGroupsAPIDeleteNetworkSmTargetGroupRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		targetGroupId: targetGroupId,
	}
}

// Execute executes the request
func (a *TargetGroupsAPIService) DeleteNetworkSmTargetGroupExecute(r TargetGroupsAPIDeleteNetworkSmTargetGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsAPIService.DeleteNetworkSmTargetGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", url.PathEscape(parameterValueToString(r.targetGroupId, "targetGroupId")), -1)

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

type TargetGroupsAPIGetNetworkSmTargetGroupRequest struct {
	ctx context.Context
	ApiService *TargetGroupsAPIService
	networkId string
	targetGroupId string
	withDetails *bool
}

// Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
func (r TargetGroupsAPIGetNetworkSmTargetGroupRequest) WithDetails(withDetails bool) TargetGroupsAPIGetNetworkSmTargetGroupRequest {
	r.withDetails = &withDetails
	return r
}

func (r TargetGroupsAPIGetNetworkSmTargetGroupRequest) Execute() (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSmTargetGroupExecute(r)
}

/*
GetNetworkSmTargetGroup Return a target group

Return a target group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param targetGroupId Target group ID
 @return TargetGroupsAPIGetNetworkSmTargetGroupRequest
*/
func (a *TargetGroupsAPIService) GetNetworkSmTargetGroup(ctx context.Context, networkId string, targetGroupId string) TargetGroupsAPIGetNetworkSmTargetGroupRequest {
	return TargetGroupsAPIGetNetworkSmTargetGroupRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		targetGroupId: targetGroupId,
	}
}

// Execute executes the request
//  @return GetNetworkSmTargetGroups200ResponseInner
func (a *TargetGroupsAPIService) GetNetworkSmTargetGroupExecute(r TargetGroupsAPIGetNetworkSmTargetGroupRequest) (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSmTargetGroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsAPIService.GetNetworkSmTargetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", url.PathEscape(parameterValueToString(r.targetGroupId, "targetGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withDetails", r.withDetails, "")
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

type TargetGroupsAPIGetNetworkSmTargetGroupsRequest struct {
	ctx context.Context
	ApiService *TargetGroupsAPIService
	networkId string
	withDetails *bool
}

// Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response
func (r TargetGroupsAPIGetNetworkSmTargetGroupsRequest) WithDetails(withDetails bool) TargetGroupsAPIGetNetworkSmTargetGroupsRequest {
	r.withDetails = &withDetails
	return r
}

func (r TargetGroupsAPIGetNetworkSmTargetGroupsRequest) Execute() ([]GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSmTargetGroupsExecute(r)
}

/*
GetNetworkSmTargetGroups List the target groups in this network

List the target groups in this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return TargetGroupsAPIGetNetworkSmTargetGroupsRequest
*/
func (a *TargetGroupsAPIService) GetNetworkSmTargetGroups(ctx context.Context, networkId string) TargetGroupsAPIGetNetworkSmTargetGroupsRequest {
	return TargetGroupsAPIGetNetworkSmTargetGroupsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSmTargetGroups200ResponseInner
func (a *TargetGroupsAPIService) GetNetworkSmTargetGroupsExecute(r TargetGroupsAPIGetNetworkSmTargetGroupsRequest) ([]GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSmTargetGroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsAPIService.GetNetworkSmTargetGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/targetGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withDetails", r.withDetails, "")
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

type TargetGroupsAPIUpdateNetworkSmTargetGroupRequest struct {
	ctx context.Context
	ApiService *TargetGroupsAPIService
	networkId string
	targetGroupId string
	createNetworkSmTargetGroupRequest *CreateNetworkSmTargetGroupRequest
}

func (r TargetGroupsAPIUpdateNetworkSmTargetGroupRequest) CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest CreateNetworkSmTargetGroupRequest) TargetGroupsAPIUpdateNetworkSmTargetGroupRequest {
	r.createNetworkSmTargetGroupRequest = &createNetworkSmTargetGroupRequest
	return r
}

func (r TargetGroupsAPIUpdateNetworkSmTargetGroupRequest) Execute() (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkSmTargetGroupExecute(r)
}

/*
UpdateNetworkSmTargetGroup Update a target group

Update a target group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param targetGroupId Target group ID
 @return TargetGroupsAPIUpdateNetworkSmTargetGroupRequest
*/
func (a *TargetGroupsAPIService) UpdateNetworkSmTargetGroup(ctx context.Context, networkId string, targetGroupId string) TargetGroupsAPIUpdateNetworkSmTargetGroupRequest {
	return TargetGroupsAPIUpdateNetworkSmTargetGroupRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		targetGroupId: targetGroupId,
	}
}

// Execute executes the request
//  @return GetNetworkSmTargetGroups200ResponseInner
func (a *TargetGroupsAPIService) UpdateNetworkSmTargetGroupExecute(r TargetGroupsAPIUpdateNetworkSmTargetGroupRequest) (*GetNetworkSmTargetGroups200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSmTargetGroups200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsAPIService.UpdateNetworkSmTargetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/targetGroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", url.PathEscape(parameterValueToString(r.targetGroupId, "targetGroupId")), -1)

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
	localVarPostBody = r.createNetworkSmTargetGroupRequest
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

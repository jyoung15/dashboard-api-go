/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
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


// ActionBatchesApiService ActionBatchesApi service
type ActionBatchesApiService service

type ActionBatchesApiCreateOrganizationActionBatchRequest struct {
	ctx context.Context
	ApiService *ActionBatchesApiService
	organizationId string
	createOrganizationActionBatchRequest *CreateOrganizationActionBatchRequest
}

func (r ActionBatchesApiCreateOrganizationActionBatchRequest) CreateOrganizationActionBatchRequest(createOrganizationActionBatchRequest CreateOrganizationActionBatchRequest) ActionBatchesApiCreateOrganizationActionBatchRequest {
	r.createOrganizationActionBatchRequest = &createOrganizationActionBatchRequest
	return r
}

func (r ActionBatchesApiCreateOrganizationActionBatchRequest) Execute() (*CreateOrganizationActionBatch201Response, *http.Response, error) {
	return r.ApiService.CreateOrganizationActionBatchExecute(r)
}

/*
CreateOrganizationActionBatch Create an action batch

Create an action batch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ActionBatchesApiCreateOrganizationActionBatchRequest
*/
func (a *ActionBatchesApiService) CreateOrganizationActionBatch(ctx context.Context, organizationId string) ActionBatchesApiCreateOrganizationActionBatchRequest {
	return ActionBatchesApiCreateOrganizationActionBatchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return CreateOrganizationActionBatch201Response
func (a *ActionBatchesApiService) CreateOrganizationActionBatchExecute(r ActionBatchesApiCreateOrganizationActionBatchRequest) (*CreateOrganizationActionBatch201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateOrganizationActionBatch201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionBatchesApiService.CreateOrganizationActionBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/actionBatches"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationActionBatchRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationActionBatchRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationActionBatchRequest
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

type ActionBatchesApiDeleteOrganizationActionBatchRequest struct {
	ctx context.Context
	ApiService *ActionBatchesApiService
	organizationId string
	actionBatchId string
}

func (r ActionBatchesApiDeleteOrganizationActionBatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationActionBatchExecute(r)
}

/*
DeleteOrganizationActionBatch Delete an action batch

Delete an action batch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param actionBatchId Action batch ID
 @return ActionBatchesApiDeleteOrganizationActionBatchRequest
*/
func (a *ActionBatchesApiService) DeleteOrganizationActionBatch(ctx context.Context, organizationId string, actionBatchId string) ActionBatchesApiDeleteOrganizationActionBatchRequest {
	return ActionBatchesApiDeleteOrganizationActionBatchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		actionBatchId: actionBatchId,
	}
}

// Execute executes the request
func (a *ActionBatchesApiService) DeleteOrganizationActionBatchExecute(r ActionBatchesApiDeleteOrganizationActionBatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionBatchesApiService.DeleteOrganizationActionBatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/actionBatches/{actionBatchId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"actionBatchId"+"}", url.PathEscape(parameterValueToString(r.actionBatchId, "actionBatchId")), -1)

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

type ActionBatchesApiGetOrganizationActionBatchRequest struct {
	ctx context.Context
	ApiService *ActionBatchesApiService
	organizationId string
	actionBatchId string
}

func (r ActionBatchesApiGetOrganizationActionBatchRequest) Execute() (*CreateOrganizationActionBatch201Response, *http.Response, error) {
	return r.ApiService.GetOrganizationActionBatchExecute(r)
}

/*
GetOrganizationActionBatch Return an action batch

Return an action batch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param actionBatchId Action batch ID
 @return ActionBatchesApiGetOrganizationActionBatchRequest
*/
func (a *ActionBatchesApiService) GetOrganizationActionBatch(ctx context.Context, organizationId string, actionBatchId string) ActionBatchesApiGetOrganizationActionBatchRequest {
	return ActionBatchesApiGetOrganizationActionBatchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		actionBatchId: actionBatchId,
	}
}

// Execute executes the request
//  @return CreateOrganizationActionBatch201Response
func (a *ActionBatchesApiService) GetOrganizationActionBatchExecute(r ActionBatchesApiGetOrganizationActionBatchRequest) (*CreateOrganizationActionBatch201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateOrganizationActionBatch201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionBatchesApiService.GetOrganizationActionBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/actionBatches/{actionBatchId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"actionBatchId"+"}", url.PathEscape(parameterValueToString(r.actionBatchId, "actionBatchId")), -1)

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

type ActionBatchesApiGetOrganizationActionBatchesRequest struct {
	ctx context.Context
	ApiService *ActionBatchesApiService
	organizationId string
	status *string
}

// Filter batches by status. Valid types are pending, completed, and failed.
func (r ActionBatchesApiGetOrganizationActionBatchesRequest) Status(status string) ActionBatchesApiGetOrganizationActionBatchesRequest {
	r.status = &status
	return r
}

func (r ActionBatchesApiGetOrganizationActionBatchesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOrganizationActionBatchesExecute(r)
}

/*
GetOrganizationActionBatches Return the list of action batches in the organization

Return the list of action batches in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ActionBatchesApiGetOrganizationActionBatchesRequest
*/
func (a *ActionBatchesApiService) GetOrganizationActionBatches(ctx context.Context, organizationId string) ActionBatchesApiGetOrganizationActionBatchesRequest {
	return ActionBatchesApiGetOrganizationActionBatchesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ActionBatchesApiService) GetOrganizationActionBatchesExecute(r ActionBatchesApiGetOrganizationActionBatchesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionBatchesApiService.GetOrganizationActionBatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/actionBatches"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
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

type ActionBatchesApiUpdateOrganizationActionBatchRequest struct {
	ctx context.Context
	ApiService *ActionBatchesApiService
	organizationId string
	actionBatchId string
	updateOrganizationActionBatchRequest *UpdateOrganizationActionBatchRequest
}

func (r ActionBatchesApiUpdateOrganizationActionBatchRequest) UpdateOrganizationActionBatchRequest(updateOrganizationActionBatchRequest UpdateOrganizationActionBatchRequest) ActionBatchesApiUpdateOrganizationActionBatchRequest {
	r.updateOrganizationActionBatchRequest = &updateOrganizationActionBatchRequest
	return r
}

func (r ActionBatchesApiUpdateOrganizationActionBatchRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateOrganizationActionBatchExecute(r)
}

/*
UpdateOrganizationActionBatch Update an action batch

Update an action batch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param actionBatchId Action batch ID
 @return ActionBatchesApiUpdateOrganizationActionBatchRequest
*/
func (a *ActionBatchesApiService) UpdateOrganizationActionBatch(ctx context.Context, organizationId string, actionBatchId string) ActionBatchesApiUpdateOrganizationActionBatchRequest {
	return ActionBatchesApiUpdateOrganizationActionBatchRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		actionBatchId: actionBatchId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ActionBatchesApiService) UpdateOrganizationActionBatchExecute(r ActionBatchesApiUpdateOrganizationActionBatchRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionBatchesApiService.UpdateOrganizationActionBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/actionBatches/{actionBatchId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"actionBatchId"+"}", url.PathEscape(parameterValueToString(r.actionBatchId, "actionBatchId")), -1)

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
	localVarPostBody = r.updateOrganizationActionBatchRequest
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

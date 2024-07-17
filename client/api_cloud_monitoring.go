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
	"reflect"
)


// CloudMonitoringAPIService CloudMonitoringAPI service
type CloudMonitoringAPIService service

type CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringAPIService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest = &createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
	return r
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringExportEventExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent Imports event logs related to the onboarding app into elastisearch

Imports event logs related to the onboarding app into elastisearch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
*/
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(ctx context.Context, organizationId string) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringExportEventExecute(r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringAPIService.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
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

type CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringAPIService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringImportRequest *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest = &createOrganizationInventoryOnboardingCloudMonitoringImportRequest
	return r
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Execute() ([]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
*/
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx context.Context, organizationId string) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ([]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringAPIService.CreateOrganizationInventoryOnboardingCloudMonitoringImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringImportRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest
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

type CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringAPIService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest = &createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
	return r
}

func (r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Execute() ([]CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringPrepare Initiates or updates an import session

Initiates or updates an import session. An import ID will be generated and used when you are ready to commit the import.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
*/
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx context.Context, organizationId string) CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
func (a *CloudMonitoringAPIService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r CloudMonitoringAPICreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ([]CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringAPIService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
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

type CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringAPIService
	organizationId string
	importIds *[]string
}

// import ids from an imports
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ImportIds(importIds []string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	r.importIds = &importIds
	return r
}

func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) Execute() ([]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation

Check the status of a committed Import operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest
*/
func (a *CloudMonitoringAPIService) GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx context.Context, organizationId string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	return CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
func (a *CloudMonitoringAPIService) GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ([]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringAPIService.GetOrganizationInventoryOnboardingCloudMonitoringImports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importIds == nil {
		return localVarReturnValue, nil, reportError("importIds is required and must be specified")
	}

	{
		t := *r.importIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "importIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "importIds", t, "multi")
		}
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

type CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct {
	ctx context.Context
	ApiService *CloudMonitoringAPIService
	organizationId string
	deviceType *string
	search *string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// Device Type switch or wireless controller
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) DeviceType(deviceType string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	r.deviceType = &deviceType
	return r
}

// Optional parameter to search on network name
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) Search(search string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	r.search = &search
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000.
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) PerPage(perPage int32) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) StartingAfter(startingAfter string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) EndingBefore(endingBefore string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) Execute() ([]GetNetwork200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringNetworksExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringNetworks Returns list of networks eligible for adding cloud monitored device

Returns list of networks eligible for adding cloud monitored device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest
*/
func (a *CloudMonitoringAPIService) GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx context.Context, organizationId string) CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest {
	return CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetNetwork200Response
func (a *CloudMonitoringAPIService) GetOrganizationInventoryOnboardingCloudMonitoringNetworksExecute(r CloudMonitoringAPIGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest) ([]GetNetwork200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetwork200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMonitoringAPIService.GetOrganizationInventoryOnboardingCloudMonitoringNetworks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceType == nil {
		return localVarReturnValue, nil, reportError("deviceType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "deviceType", r.deviceType, "")
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
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

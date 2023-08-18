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


// AclsAPIService AclsAPI service
type AclsAPIService service

type AclsAPICreateOrganizationAdaptivePolicyAclRequest struct {
	ctx context.Context
	ApiService *AclsAPIService
	organizationId string
	createOrganizationAdaptivePolicyAclRequest *CreateOrganizationAdaptivePolicyAclRequest
}

func (r AclsAPICreateOrganizationAdaptivePolicyAclRequest) CreateOrganizationAdaptivePolicyAclRequest(createOrganizationAdaptivePolicyAclRequest CreateOrganizationAdaptivePolicyAclRequest) AclsAPICreateOrganizationAdaptivePolicyAclRequest {
	r.createOrganizationAdaptivePolicyAclRequest = &createOrganizationAdaptivePolicyAclRequest
	return r
}

func (r AclsAPICreateOrganizationAdaptivePolicyAclRequest) Execute() (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateOrganizationAdaptivePolicyAclExecute(r)
}

/*
CreateOrganizationAdaptivePolicyAcl Creates new adaptive policy ACL

Creates new adaptive policy ACL

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AclsAPICreateOrganizationAdaptivePolicyAclRequest
*/
func (a *AclsAPIService) CreateOrganizationAdaptivePolicyAcl(ctx context.Context, organizationId string) AclsAPICreateOrganizationAdaptivePolicyAclRequest {
	return AclsAPICreateOrganizationAdaptivePolicyAclRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationAdaptivePolicyAcls200ResponseInner
func (a *AclsAPIService) CreateOrganizationAdaptivePolicyAclExecute(r AclsAPICreateOrganizationAdaptivePolicyAclRequest) (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationAdaptivePolicyAcls200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsAPIService.CreateOrganizationAdaptivePolicyAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/adaptivePolicy/acls"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationAdaptivePolicyAclRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationAdaptivePolicyAclRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationAdaptivePolicyAclRequest
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

type AclsAPIDeleteOrganizationAdaptivePolicyAclRequest struct {
	ctx context.Context
	ApiService *AclsAPIService
	organizationId string
	aclId string
}

func (r AclsAPIDeleteOrganizationAdaptivePolicyAclRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationAdaptivePolicyAclExecute(r)
}

/*
DeleteOrganizationAdaptivePolicyAcl Deletes the specified adaptive policy ACL

Deletes the specified adaptive policy ACL. Note this adaptive policy ACL will also be removed from policies using it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param aclId Acl ID
 @return AclsAPIDeleteOrganizationAdaptivePolicyAclRequest
*/
func (a *AclsAPIService) DeleteOrganizationAdaptivePolicyAcl(ctx context.Context, organizationId string, aclId string) AclsAPIDeleteOrganizationAdaptivePolicyAclRequest {
	return AclsAPIDeleteOrganizationAdaptivePolicyAclRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		aclId: aclId,
	}
}

// Execute executes the request
func (a *AclsAPIService) DeleteOrganizationAdaptivePolicyAclExecute(r AclsAPIDeleteOrganizationAdaptivePolicyAclRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsAPIService.DeleteOrganizationAdaptivePolicyAcl")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclId"+"}", url.PathEscape(parameterValueToString(r.aclId, "aclId")), -1)

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

type AclsAPIGetOrganizationAdaptivePolicyAclRequest struct {
	ctx context.Context
	ApiService *AclsAPIService
	organizationId string
	aclId string
}

func (r AclsAPIGetOrganizationAdaptivePolicyAclRequest) Execute() (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationAdaptivePolicyAclExecute(r)
}

/*
GetOrganizationAdaptivePolicyAcl Returns the adaptive policy ACL information

Returns the adaptive policy ACL information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param aclId Acl ID
 @return AclsAPIGetOrganizationAdaptivePolicyAclRequest
*/
func (a *AclsAPIService) GetOrganizationAdaptivePolicyAcl(ctx context.Context, organizationId string, aclId string) AclsAPIGetOrganizationAdaptivePolicyAclRequest {
	return AclsAPIGetOrganizationAdaptivePolicyAclRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		aclId: aclId,
	}
}

// Execute executes the request
//  @return GetOrganizationAdaptivePolicyAcls200ResponseInner
func (a *AclsAPIService) GetOrganizationAdaptivePolicyAclExecute(r AclsAPIGetOrganizationAdaptivePolicyAclRequest) (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationAdaptivePolicyAcls200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsAPIService.GetOrganizationAdaptivePolicyAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclId"+"}", url.PathEscape(parameterValueToString(r.aclId, "aclId")), -1)

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

type AclsAPIGetOrganizationAdaptivePolicyAclsRequest struct {
	ctx context.Context
	ApiService *AclsAPIService
	organizationId string
}

func (r AclsAPIGetOrganizationAdaptivePolicyAclsRequest) Execute() ([]GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationAdaptivePolicyAclsExecute(r)
}

/*
GetOrganizationAdaptivePolicyAcls List adaptive policy ACLs in a organization

List adaptive policy ACLs in a organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AclsAPIGetOrganizationAdaptivePolicyAclsRequest
*/
func (a *AclsAPIService) GetOrganizationAdaptivePolicyAcls(ctx context.Context, organizationId string) AclsAPIGetOrganizationAdaptivePolicyAclsRequest {
	return AclsAPIGetOrganizationAdaptivePolicyAclsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationAdaptivePolicyAcls200ResponseInner
func (a *AclsAPIService) GetOrganizationAdaptivePolicyAclsExecute(r AclsAPIGetOrganizationAdaptivePolicyAclsRequest) ([]GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationAdaptivePolicyAcls200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsAPIService.GetOrganizationAdaptivePolicyAcls")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/adaptivePolicy/acls"
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

type AclsAPIUpdateOrganizationAdaptivePolicyAclRequest struct {
	ctx context.Context
	ApiService *AclsAPIService
	organizationId string
	aclId string
	updateOrganizationAdaptivePolicyAclRequest *UpdateOrganizationAdaptivePolicyAclRequest
}

func (r AclsAPIUpdateOrganizationAdaptivePolicyAclRequest) UpdateOrganizationAdaptivePolicyAclRequest(updateOrganizationAdaptivePolicyAclRequest UpdateOrganizationAdaptivePolicyAclRequest) AclsAPIUpdateOrganizationAdaptivePolicyAclRequest {
	r.updateOrganizationAdaptivePolicyAclRequest = &updateOrganizationAdaptivePolicyAclRequest
	return r
}

func (r AclsAPIUpdateOrganizationAdaptivePolicyAclRequest) Execute() (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateOrganizationAdaptivePolicyAclExecute(r)
}

/*
UpdateOrganizationAdaptivePolicyAcl Updates an adaptive policy ACL

Updates an adaptive policy ACL

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param aclId Acl ID
 @return AclsAPIUpdateOrganizationAdaptivePolicyAclRequest
*/
func (a *AclsAPIService) UpdateOrganizationAdaptivePolicyAcl(ctx context.Context, organizationId string, aclId string) AclsAPIUpdateOrganizationAdaptivePolicyAclRequest {
	return AclsAPIUpdateOrganizationAdaptivePolicyAclRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		aclId: aclId,
	}
}

// Execute executes the request
//  @return GetOrganizationAdaptivePolicyAcls200ResponseInner
func (a *AclsAPIService) UpdateOrganizationAdaptivePolicyAclExecute(r AclsAPIUpdateOrganizationAdaptivePolicyAclRequest) (*GetOrganizationAdaptivePolicyAcls200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationAdaptivePolicyAcls200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AclsAPIService.UpdateOrganizationAdaptivePolicyAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/adaptivePolicy/acls/{aclId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclId"+"}", url.PathEscape(parameterValueToString(r.aclId, "aclId")), -1)

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
	localVarPostBody = r.updateOrganizationAdaptivePolicyAclRequest
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

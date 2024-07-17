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


// BrandingPoliciesAPIService BrandingPoliciesAPI service
type BrandingPoliciesAPIService service

type BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
	createOrganizationBrandingPolicyRequest *CreateOrganizationBrandingPolicyRequest
}

func (r BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest) CreateOrganizationBrandingPolicyRequest(createOrganizationBrandingPolicyRequest CreateOrganizationBrandingPolicyRequest) BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest {
	r.createOrganizationBrandingPolicyRequest = &createOrganizationBrandingPolicyRequest
	return r
}

func (r BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest) Execute() (*CreateOrganizationBrandingPolicy201Response, *http.Response, error) {
	return r.ApiService.CreateOrganizationBrandingPolicyExecute(r)
}

/*
CreateOrganizationBrandingPolicy Add a new branding policy to an organization

Add a new branding policy to an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest
*/
func (a *BrandingPoliciesAPIService) CreateOrganizationBrandingPolicy(ctx context.Context, organizationId string) BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest {
	return BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return CreateOrganizationBrandingPolicy201Response
func (a *BrandingPoliciesAPIService) CreateOrganizationBrandingPolicyExecute(r BrandingPoliciesAPICreateOrganizationBrandingPolicyRequest) (*CreateOrganizationBrandingPolicy201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateOrganizationBrandingPolicy201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.CreateOrganizationBrandingPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	localVarPostBody = r.createOrganizationBrandingPolicyRequest
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

type BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
	brandingPolicyId string
}

func (r BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationBrandingPolicyExecute(r)
}

/*
DeleteOrganizationBrandingPolicy Delete a branding policy

Delete a branding policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param brandingPolicyId Branding policy ID
 @return BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest
*/
func (a *BrandingPoliciesAPIService) DeleteOrganizationBrandingPolicy(ctx context.Context, organizationId string, brandingPolicyId string) BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest {
	return BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		brandingPolicyId: brandingPolicyId,
	}
}

// Execute executes the request
func (a *BrandingPoliciesAPIService) DeleteOrganizationBrandingPolicyExecute(r BrandingPoliciesAPIDeleteOrganizationBrandingPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.DeleteOrganizationBrandingPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandingPolicyId"+"}", url.PathEscape(parameterValueToString(r.brandingPolicyId, "brandingPolicyId")), -1)

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

type BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
}

func (r BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest) Execute() ([]GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationBrandingPoliciesExecute(r)
}

/*
GetOrganizationBrandingPolicies List the branding policies of an organization

List the branding policies of an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest
*/
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPolicies(ctx context.Context, organizationId string) BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest {
	return BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationBrandingPolicies200ResponseInner
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPoliciesExecute(r BrandingPoliciesAPIGetOrganizationBrandingPoliciesRequest) ([]GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationBrandingPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.GetOrganizationBrandingPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies"
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

type BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
}

func (r BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest) Execute() (*GetOrganizationBrandingPoliciesPriorities200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationBrandingPoliciesPrioritiesExecute(r)
}

/*
GetOrganizationBrandingPoliciesPriorities Return the branding policy IDs of an organization in priority order

Return the branding policy IDs of an organization in priority order. IDs are ordered in ascending order of priority (IDs later in the array have higher priority).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest
*/
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPoliciesPriorities(ctx context.Context, organizationId string) BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest {
	return BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationBrandingPoliciesPriorities200Response
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPoliciesPrioritiesExecute(r BrandingPoliciesAPIGetOrganizationBrandingPoliciesPrioritiesRequest) (*GetOrganizationBrandingPoliciesPriorities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationBrandingPoliciesPriorities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.GetOrganizationBrandingPoliciesPriorities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies/priorities"
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

type BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
	brandingPolicyId string
}

func (r BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest) Execute() (*GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationBrandingPolicyExecute(r)
}

/*
GetOrganizationBrandingPolicy Return a branding policy

Return a branding policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param brandingPolicyId Branding policy ID
 @return BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest
*/
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPolicy(ctx context.Context, organizationId string, brandingPolicyId string) BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest {
	return BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		brandingPolicyId: brandingPolicyId,
	}
}

// Execute executes the request
//  @return GetOrganizationBrandingPolicies200ResponseInner
func (a *BrandingPoliciesAPIService) GetOrganizationBrandingPolicyExecute(r BrandingPoliciesAPIGetOrganizationBrandingPolicyRequest) (*GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationBrandingPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.GetOrganizationBrandingPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandingPolicyId"+"}", url.PathEscape(parameterValueToString(r.brandingPolicyId, "brandingPolicyId")), -1)

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

type BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
	updateOrganizationBrandingPoliciesPrioritiesRequest *UpdateOrganizationBrandingPoliciesPrioritiesRequest
}

func (r BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest) UpdateOrganizationBrandingPoliciesPrioritiesRequest(updateOrganizationBrandingPoliciesPrioritiesRequest UpdateOrganizationBrandingPoliciesPrioritiesRequest) BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest {
	r.updateOrganizationBrandingPoliciesPrioritiesRequest = &updateOrganizationBrandingPoliciesPrioritiesRequest
	return r
}

func (r BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest) Execute() (*GetOrganizationBrandingPoliciesPriorities200Response, *http.Response, error) {
	return r.ApiService.UpdateOrganizationBrandingPoliciesPrioritiesExecute(r)
}

/*
UpdateOrganizationBrandingPoliciesPriorities Update the priority ordering of an organization's branding policies.

Update the priority ordering of an organization's branding policies.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest
*/
func (a *BrandingPoliciesAPIService) UpdateOrganizationBrandingPoliciesPriorities(ctx context.Context, organizationId string) BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest {
	return BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationBrandingPoliciesPriorities200Response
func (a *BrandingPoliciesAPIService) UpdateOrganizationBrandingPoliciesPrioritiesExecute(r BrandingPoliciesAPIUpdateOrganizationBrandingPoliciesPrioritiesRequest) (*GetOrganizationBrandingPoliciesPriorities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationBrandingPoliciesPriorities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.UpdateOrganizationBrandingPoliciesPriorities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies/priorities"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	localVarPostBody = r.updateOrganizationBrandingPoliciesPrioritiesRequest
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

type BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest struct {
	ctx context.Context
	ApiService *BrandingPoliciesAPIService
	organizationId string
	brandingPolicyId string
	updateOrganizationBrandingPolicyRequest *UpdateOrganizationBrandingPolicyRequest
}

func (r BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest) UpdateOrganizationBrandingPolicyRequest(updateOrganizationBrandingPolicyRequest UpdateOrganizationBrandingPolicyRequest) BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest {
	r.updateOrganizationBrandingPolicyRequest = &updateOrganizationBrandingPolicyRequest
	return r
}

func (r BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest) Execute() (*GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateOrganizationBrandingPolicyExecute(r)
}

/*
UpdateOrganizationBrandingPolicy Update a branding policy

Update a branding policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param brandingPolicyId Branding policy ID
 @return BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest
*/
func (a *BrandingPoliciesAPIService) UpdateOrganizationBrandingPolicy(ctx context.Context, organizationId string, brandingPolicyId string) BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest {
	return BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		brandingPolicyId: brandingPolicyId,
	}
}

// Execute executes the request
//  @return GetOrganizationBrandingPolicies200ResponseInner
func (a *BrandingPoliciesAPIService) UpdateOrganizationBrandingPolicyExecute(r BrandingPoliciesAPIUpdateOrganizationBrandingPolicyRequest) (*GetOrganizationBrandingPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationBrandingPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingPoliciesAPIService.UpdateOrganizationBrandingPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/brandingPolicies/{brandingPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brandingPolicyId"+"}", url.PathEscape(parameterValueToString(r.brandingPolicyId, "brandingPolicyId")), -1)

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
	localVarPostBody = r.updateOrganizationBrandingPolicyRequest
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

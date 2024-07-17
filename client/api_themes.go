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


// ThemesAPIService ThemesAPI service
type ThemesAPIService service

type ThemesAPICreateOrganizationSplashThemeRequest struct {
	ctx context.Context
	ApiService *ThemesAPIService
	organizationId string
	createOrganizationSplashThemeRequest *CreateOrganizationSplashThemeRequest
}

func (r ThemesAPICreateOrganizationSplashThemeRequest) CreateOrganizationSplashThemeRequest(createOrganizationSplashThemeRequest CreateOrganizationSplashThemeRequest) ThemesAPICreateOrganizationSplashThemeRequest {
	r.createOrganizationSplashThemeRequest = &createOrganizationSplashThemeRequest
	return r
}

func (r ThemesAPICreateOrganizationSplashThemeRequest) Execute() (*GetOrganizationSplashThemes200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateOrganizationSplashThemeExecute(r)
}

/*
CreateOrganizationSplashTheme Create a Splash Theme

Create a Splash Theme

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ThemesAPICreateOrganizationSplashThemeRequest
*/
func (a *ThemesAPIService) CreateOrganizationSplashTheme(ctx context.Context, organizationId string) ThemesAPICreateOrganizationSplashThemeRequest {
	return ThemesAPICreateOrganizationSplashThemeRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationSplashThemes200ResponseInner
func (a *ThemesAPIService) CreateOrganizationSplashThemeExecute(r ThemesAPICreateOrganizationSplashThemeRequest) (*GetOrganizationSplashThemes200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationSplashThemes200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemesAPIService.CreateOrganizationSplashTheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/splash/themes"
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
	localVarPostBody = r.createOrganizationSplashThemeRequest
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

type ThemesAPICreateOrganizationSplashThemeAssetRequest struct {
	ctx context.Context
	ApiService *ThemesAPIService
	organizationId string
	themeIdentifier string
	createOrganizationSplashThemeAssetRequest *CreateOrganizationSplashThemeAssetRequest
}

func (r ThemesAPICreateOrganizationSplashThemeAssetRequest) CreateOrganizationSplashThemeAssetRequest(createOrganizationSplashThemeAssetRequest CreateOrganizationSplashThemeAssetRequest) ThemesAPICreateOrganizationSplashThemeAssetRequest {
	r.createOrganizationSplashThemeAssetRequest = &createOrganizationSplashThemeAssetRequest
	return r
}

func (r ThemesAPICreateOrganizationSplashThemeAssetRequest) Execute() (*GetOrganizationSplashAsset200Response, *http.Response, error) {
	return r.ApiService.CreateOrganizationSplashThemeAssetExecute(r)
}

/*
CreateOrganizationSplashThemeAsset Create a Splash Theme Asset

Create a Splash Theme Asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param themeIdentifier Theme identifier
 @return ThemesAPICreateOrganizationSplashThemeAssetRequest
*/
func (a *ThemesAPIService) CreateOrganizationSplashThemeAsset(ctx context.Context, organizationId string, themeIdentifier string) ThemesAPICreateOrganizationSplashThemeAssetRequest {
	return ThemesAPICreateOrganizationSplashThemeAssetRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		themeIdentifier: themeIdentifier,
	}
}

// Execute executes the request
//  @return GetOrganizationSplashAsset200Response
func (a *ThemesAPIService) CreateOrganizationSplashThemeAssetExecute(r ThemesAPICreateOrganizationSplashThemeAssetRequest) (*GetOrganizationSplashAsset200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationSplashAsset200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemesAPIService.CreateOrganizationSplashThemeAsset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/splash/themes/{themeIdentifier}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"themeIdentifier"+"}", url.PathEscape(parameterValueToString(r.themeIdentifier, "themeIdentifier")), -1)

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
	localVarPostBody = r.createOrganizationSplashThemeAssetRequest
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

type ThemesAPIDeleteOrganizationSplashThemeRequest struct {
	ctx context.Context
	ApiService *ThemesAPIService
	organizationId string
	id string
}

func (r ThemesAPIDeleteOrganizationSplashThemeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOrganizationSplashThemeExecute(r)
}

/*
DeleteOrganizationSplashTheme Delete a Splash Theme

Delete a Splash Theme

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @param id ID
 @return ThemesAPIDeleteOrganizationSplashThemeRequest
*/
func (a *ThemesAPIService) DeleteOrganizationSplashTheme(ctx context.Context, organizationId string, id string) ThemesAPIDeleteOrganizationSplashThemeRequest {
	return ThemesAPIDeleteOrganizationSplashThemeRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		id: id,
	}
}

// Execute executes the request
func (a *ThemesAPIService) DeleteOrganizationSplashThemeExecute(r ThemesAPIDeleteOrganizationSplashThemeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemesAPIService.DeleteOrganizationSplashTheme")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/splash/themes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ThemesAPIGetOrganizationSplashThemesRequest struct {
	ctx context.Context
	ApiService *ThemesAPIService
	organizationId string
}

func (r ThemesAPIGetOrganizationSplashThemesRequest) Execute() ([]GetOrganizationSplashThemes200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSplashThemesExecute(r)
}

/*
GetOrganizationSplashThemes List Splash Themes

List Splash Themes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ThemesAPIGetOrganizationSplashThemesRequest
*/
func (a *ThemesAPIService) GetOrganizationSplashThemes(ctx context.Context, organizationId string) ThemesAPIGetOrganizationSplashThemesRequest {
	return ThemesAPIGetOrganizationSplashThemesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSplashThemes200ResponseInner
func (a *ThemesAPIService) GetOrganizationSplashThemesExecute(r ThemesAPIGetOrganizationSplashThemesRequest) ([]GetOrganizationSplashThemes200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSplashThemes200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemesAPIService.GetOrganizationSplashThemes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/splash/themes"
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
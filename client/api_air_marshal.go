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


// AirMarshalAPIService AirMarshalAPI service
type AirMarshalAPIService service

type AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	networkId string
	createNetworkWirelessAirMarshalRuleRequest *CreateNetworkWirelessAirMarshalRuleRequest
}

func (r AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest) CreateNetworkWirelessAirMarshalRuleRequest(createNetworkWirelessAirMarshalRuleRequest CreateNetworkWirelessAirMarshalRuleRequest) AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest {
	r.createNetworkWirelessAirMarshalRuleRequest = &createNetworkWirelessAirMarshalRuleRequest
	return r
}

func (r AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest) Execute() (*CreateNetworkWirelessAirMarshalRule201Response, *http.Response, error) {
	return r.ApiService.CreateNetworkWirelessAirMarshalRuleExecute(r)
}

/*
CreateNetworkWirelessAirMarshalRule Creates a new rule

Creates a new rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest
*/
func (a *AirMarshalAPIService) CreateNetworkWirelessAirMarshalRule(ctx context.Context, networkId string) AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest {
	return AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return CreateNetworkWirelessAirMarshalRule201Response
func (a *AirMarshalAPIService) CreateNetworkWirelessAirMarshalRuleExecute(r AirMarshalAPICreateNetworkWirelessAirMarshalRuleRequest) (*CreateNetworkWirelessAirMarshalRule201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateNetworkWirelessAirMarshalRule201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.CreateNetworkWirelessAirMarshalRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/airMarshal/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkWirelessAirMarshalRuleRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkWirelessAirMarshalRuleRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkWirelessAirMarshalRuleRequest
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

type AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	networkId string
	ruleId string
}

func (r AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkWirelessAirMarshalRuleExecute(r)
}

/*
DeleteNetworkWirelessAirMarshalRule Delete an Air Marshal rule.

Delete an Air Marshal rule.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param ruleId Rule ID
 @return AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest
*/
func (a *AirMarshalAPIService) DeleteNetworkWirelessAirMarshalRule(ctx context.Context, networkId string, ruleId string) AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest {
	return AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		ruleId: ruleId,
	}
}

// Execute executes the request
func (a *AirMarshalAPIService) DeleteNetworkWirelessAirMarshalRuleExecute(r AirMarshalAPIDeleteNetworkWirelessAirMarshalRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.DeleteNetworkWirelessAirMarshalRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/airMarshal/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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

type AirMarshalAPIGetNetworkWirelessAirMarshalRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	networkId string
	t0 *string
	timespan *float32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r AirMarshalAPIGetNetworkWirelessAirMarshalRequest) T0(t0 string) AirMarshalAPIGetNetworkWirelessAirMarshalRequest {
	r.t0 = &t0
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r AirMarshalAPIGetNetworkWirelessAirMarshalRequest) Timespan(timespan float32) AirMarshalAPIGetNetworkWirelessAirMarshalRequest {
	r.timespan = &timespan
	return r
}

func (r AirMarshalAPIGetNetworkWirelessAirMarshalRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessAirMarshalExecute(r)
}

/*
GetNetworkWirelessAirMarshal List Air Marshal scan results from a network

List Air Marshal scan results from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AirMarshalAPIGetNetworkWirelessAirMarshalRequest
*/
func (a *AirMarshalAPIService) GetNetworkWirelessAirMarshal(ctx context.Context, networkId string) AirMarshalAPIGetNetworkWirelessAirMarshalRequest {
	return AirMarshalAPIGetNetworkWirelessAirMarshalRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AirMarshalAPIService) GetNetworkWirelessAirMarshalExecute(r AirMarshalAPIGetNetworkWirelessAirMarshalRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.GetNetworkWirelessAirMarshal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/airMarshal"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
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

type AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	organizationId string
	networkIds *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// (optional) The set of network IDs to include.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) NetworkIds(networkIds []string) AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest {
	r.networkIds = &networkIds
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) PerPage(perPage int32) AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) StartingAfter(startingAfter string) AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) EndingBefore(endingBefore string) AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) Execute() (*GetOrganizationWirelessAirMarshalRules200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessAirMarshalRulesExecute(r)
}

/*
GetOrganizationWirelessAirMarshalRules Returns the current Air Marshal rules for this organization

Returns the current Air Marshal rules for this organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest
*/
func (a *AirMarshalAPIService) GetOrganizationWirelessAirMarshalRules(ctx context.Context, organizationId string) AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest {
	return AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationWirelessAirMarshalRules200Response
func (a *AirMarshalAPIService) GetOrganizationWirelessAirMarshalRulesExecute(r AirMarshalAPIGetOrganizationWirelessAirMarshalRulesRequest) (*GetOrganizationWirelessAirMarshalRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationWirelessAirMarshalRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.GetOrganizationWirelessAirMarshalRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/airMarshal/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
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

type AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	organizationId string
	networkIds *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The network IDs to include in the result set.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) NetworkIds(networkIds []string) AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) PerPage(perPage int32) AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) StartingAfter(startingAfter string) AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) EndingBefore(endingBefore string) AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) Execute() (*GetOrganizationWirelessAirMarshalSettingsByNetwork200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessAirMarshalSettingsByNetworkExecute(r)
}

/*
GetOrganizationWirelessAirMarshalSettingsByNetwork Returns the current Air Marshal settings for this network

Returns the current Air Marshal settings for this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest
*/
func (a *AirMarshalAPIService) GetOrganizationWirelessAirMarshalSettingsByNetwork(ctx context.Context, organizationId string) AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest {
	return AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationWirelessAirMarshalSettingsByNetwork200Response
func (a *AirMarshalAPIService) GetOrganizationWirelessAirMarshalSettingsByNetworkExecute(r AirMarshalAPIGetOrganizationWirelessAirMarshalSettingsByNetworkRequest) (*GetOrganizationWirelessAirMarshalSettingsByNetwork200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.GetOrganizationWirelessAirMarshalSettingsByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/airMarshal/settings/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
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

type AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	networkId string
	ruleId string
	updateNetworkWirelessAirMarshalRuleRequest *UpdateNetworkWirelessAirMarshalRuleRequest
}

func (r AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest) UpdateNetworkWirelessAirMarshalRuleRequest(updateNetworkWirelessAirMarshalRuleRequest UpdateNetworkWirelessAirMarshalRuleRequest) AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest {
	r.updateNetworkWirelessAirMarshalRuleRequest = &updateNetworkWirelessAirMarshalRuleRequest
	return r
}

func (r AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest) Execute() (*CreateNetworkWirelessAirMarshalRule201Response, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessAirMarshalRuleExecute(r)
}

/*
UpdateNetworkWirelessAirMarshalRule Update a rule

Update a rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param ruleId Rule ID
 @return AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest
*/
func (a *AirMarshalAPIService) UpdateNetworkWirelessAirMarshalRule(ctx context.Context, networkId string, ruleId string) AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest {
	return AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return CreateNetworkWirelessAirMarshalRule201Response
func (a *AirMarshalAPIService) UpdateNetworkWirelessAirMarshalRuleExecute(r AirMarshalAPIUpdateNetworkWirelessAirMarshalRuleRequest) (*CreateNetworkWirelessAirMarshalRule201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateNetworkWirelessAirMarshalRule201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.UpdateNetworkWirelessAirMarshalRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/airMarshal/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
	localVarPostBody = r.updateNetworkWirelessAirMarshalRuleRequest
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

type AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest struct {
	ctx context.Context
	ApiService *AirMarshalAPIService
	networkId string
	updateNetworkWirelessAirMarshalSettingsRequest *UpdateNetworkWirelessAirMarshalSettingsRequest
}

func (r AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest) UpdateNetworkWirelessAirMarshalSettingsRequest(updateNetworkWirelessAirMarshalSettingsRequest UpdateNetworkWirelessAirMarshalSettingsRequest) AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest {
	r.updateNetworkWirelessAirMarshalSettingsRequest = &updateNetworkWirelessAirMarshalSettingsRequest
	return r
}

func (r AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest) Execute() (*UpdateNetworkWirelessAirMarshalSettings200Response, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessAirMarshalSettingsExecute(r)
}

/*
UpdateNetworkWirelessAirMarshalSettings Updates Air Marshal settings.

Updates Air Marshal settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest
*/
func (a *AirMarshalAPIService) UpdateNetworkWirelessAirMarshalSettings(ctx context.Context, networkId string) AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest {
	return AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return UpdateNetworkWirelessAirMarshalSettings200Response
func (a *AirMarshalAPIService) UpdateNetworkWirelessAirMarshalSettingsExecute(r AirMarshalAPIUpdateNetworkWirelessAirMarshalSettingsRequest) (*UpdateNetworkWirelessAirMarshalSettings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateNetworkWirelessAirMarshalSettings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AirMarshalAPIService.UpdateNetworkWirelessAirMarshalSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/airMarshal/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkWirelessAirMarshalSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("updateNetworkWirelessAirMarshalSettingsRequest is required and must be specified")
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
	localVarPostBody = r.updateNetworkWirelessAirMarshalSettingsRequest
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

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


// ApplicationsAPIService ApplicationsAPI service
type ApplicationsAPIService service

type ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	networkId string
	applicationId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
func (r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) T0(t0 string) ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) T1(t1 string) ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
func (r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) Timespan(timespan float32) ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300.
func (r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) Resolution(resolution int32) ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest {
	r.resolution = &resolution
	return r
}

func (r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) Execute() ([]GetNetworkInsightApplicationHealthByTime200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkInsightApplicationHealthByTimeExecute(r)
}

/*
GetNetworkInsightApplicationHealthByTime Get application health by time

Get application health by time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param applicationId Application ID
 @return ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest
*/
func (a *ApplicationsAPIService) GetNetworkInsightApplicationHealthByTime(ctx context.Context, networkId string, applicationId string) ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest {
	return ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []GetNetworkInsightApplicationHealthByTime200ResponseInner
func (a *ApplicationsAPIService) GetNetworkInsightApplicationHealthByTimeExecute(r ApplicationsAPIGetNetworkInsightApplicationHealthByTimeRequest) ([]GetNetworkInsightApplicationHealthByTime200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkInsightApplicationHealthByTime200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetNetworkInsightApplicationHealthByTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/insight/applications/{applicationId}/healthByTime"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "")
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

type ApplicationsAPIGetOrganizationInsightApplicationsRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	organizationId string
}

func (r ApplicationsAPIGetOrganizationInsightApplicationsRequest) Execute() ([]GetOrganizationInsightApplications200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightApplicationsExecute(r)
}

/*
GetOrganizationInsightApplications List all Insight tracked applications

List all Insight tracked applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsAPIGetOrganizationInsightApplicationsRequest
*/
func (a *ApplicationsAPIService) GetOrganizationInsightApplications(ctx context.Context, organizationId string) ApplicationsAPIGetOrganizationInsightApplicationsRequest {
	return ApplicationsAPIGetOrganizationInsightApplicationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationInsightApplications200ResponseInner
func (a *ApplicationsAPIService) GetOrganizationInsightApplicationsExecute(r ApplicationsAPIGetOrganizationInsightApplicationsRequest) ([]GetOrganizationInsightApplications200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationInsightApplications200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetOrganizationInsightApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/applications"
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

type ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	organizationId string
	networkTag *string
	device *string
	networkId *string
	quantity *int32
	ssidName *string
	usageUplink *string
	t0 *string
	t1 *string
	timespan *float32
}

// Match result to an exact network tag
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) NetworkTag(networkTag string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) Device(device string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.device = &device
	return r
}

// Match result to an exact network id
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) NetworkId(networkId string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) Quantity(quantity int32) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) SsidName(ssidName string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) UsageUplink(usageUplink string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) T0(t0 string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) T1(t1 string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) Timespan(timespan float32) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) Execute() ([]GetOrganizationSummaryTopApplicationsByUsage200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopApplicationsByUsageExecute(r)
}

/*
GetOrganizationSummaryTopApplicationsByUsage Return the top applications sorted by data usage over given time range

Return the top applications sorted by data usage over given time range. Default unit is megabytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest
*/
func (a *ApplicationsAPIService) GetOrganizationSummaryTopApplicationsByUsage(ctx context.Context, organizationId string) ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest {
	return ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSummaryTopApplicationsByUsage200ResponseInner
func (a *ApplicationsAPIService) GetOrganizationSummaryTopApplicationsByUsageExecute(r ApplicationsAPIGetOrganizationSummaryTopApplicationsByUsageRequest) ([]GetOrganizationSummaryTopApplicationsByUsage200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSummaryTopApplicationsByUsage200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetOrganizationSummaryTopApplicationsByUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/applications/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkTag", r.networkTag, "")
	}
	if r.device != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device", r.device, "")
	}
	if r.networkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkId", r.networkId, "")
	}
	if r.quantity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "")
	}
	if r.ssidName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssidName", r.ssidName, "")
	}
	if r.usageUplink != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "usageUplink", r.usageUplink, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
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

type ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	organizationId string
	networkTag *string
	deviceTag *string
	networkId *string
	quantity *int32
	ssidName *string
	usageUplink *string
	t0 *string
	t1 *string
	timespan *float32
}

// Match result to an exact network tag
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkTag(networkTag string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) DeviceTag(deviceTag string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.deviceTag = &deviceTag
	return r
}

// Match result to an exact network id
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkId(networkId string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Quantity(quantity int32) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) SsidName(ssidName string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) UsageUplink(usageUplink string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T0(t0 string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T1(t1 string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.
func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Timespan(timespan float32) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Execute() ([]GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r)
}

/*
GetOrganizationSummaryTopApplicationsCategoriesByUsage Return the top application categories sorted by data usage over given time range

Return the top application categories sorted by data usage over given time range. Default unit is megabytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest
*/
func (a *ApplicationsAPIService) GetOrganizationSummaryTopApplicationsCategoriesByUsage(ctx context.Context, organizationId string) ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	return ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner
func (a *ApplicationsAPIService) GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r ApplicationsAPIGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) ([]GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetOrganizationSummaryTopApplicationsCategoriesByUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/applications/categories/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkTag", r.networkTag, "")
	}
	if r.deviceTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceTag", r.deviceTag, "")
	}
	if r.networkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkId", r.networkId, "")
	}
	if r.quantity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "")
	}
	if r.ssidName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssidName", r.ssidName, "")
	}
	if r.usageUplink != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "usageUplink", r.usageUplink, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
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

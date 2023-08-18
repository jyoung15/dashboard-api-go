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
	"reflect"
)


// ApiRequestsAPIService ApiRequestsAPI service
type ApiRequestsAPIService service

type ApiRequestsAPIGetOrganizationApiRequestsRequest struct {
	ctx context.Context
	ApiService *ApiRequestsAPIService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	adminId *string
	path *string
	method *string
	responseCode *int32
	sourceIp *string
	userAgent *string
	version *int32
	operationIds *[]string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) T0(t0 string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) T1(t1 string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) Timespan(timespan float32) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) PerPage(perPage int32) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) StartingAfter(startingAfter string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) EndingBefore(endingBefore string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.endingBefore = &endingBefore
	return r
}

// Filter the results by the ID of the admin who made the API requests
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) AdminId(adminId string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.adminId = &adminId
	return r
}

// Filter the results by the path of the API requests
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) Path(path string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.path = &path
	return r
}

// Filter the results by the method of the API requests (must be &#39;GET&#39;, &#39;PUT&#39;, &#39;POST&#39; or &#39;DELETE&#39;)
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) Method(method string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.method = &method
	return r
}

// Filter the results by the response code of the API requests
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) ResponseCode(responseCode int32) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.responseCode = &responseCode
	return r
}

// Filter the results by the IP address of the originating API request
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) SourceIp(sourceIp string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.sourceIp = &sourceIp
	return r
}

// Filter the results by the user agent string of the API request
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) UserAgent(userAgent string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.userAgent = &userAgent
	return r
}

// Filter the results by the API version of the API request
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) Version(version int32) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.version = &version
	return r
}

// Filter the results by one or more operation IDs for the API request
func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) OperationIds(operationIds []string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	r.operationIds = &operationIds
	return r
}

func (r ApiRequestsAPIGetOrganizationApiRequestsRequest) Execute() ([]GetOrganizationApiRequests200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsExecute(r)
}

/*
GetOrganizationApiRequests List the API requests made by an organization

List the API requests made by an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsAPIGetOrganizationApiRequestsRequest
*/
func (a *ApiRequestsAPIService) GetOrganizationApiRequests(ctx context.Context, organizationId string) ApiRequestsAPIGetOrganizationApiRequestsRequest {
	return ApiRequestsAPIGetOrganizationApiRequestsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationApiRequests200ResponseInner
func (a *ApiRequestsAPIService) GetOrganizationApiRequestsExecute(r ApiRequestsAPIGetOrganizationApiRequestsRequest) ([]GetOrganizationApiRequests200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationApiRequests200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsAPIService.GetOrganizationApiRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.adminId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adminId", r.adminId, "")
	}
	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
	}
	if r.method != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "method", r.method, "")
	}
	if r.responseCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "responseCode", r.responseCode, "")
	}
	if r.sourceIp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sourceIp", r.sourceIp, "")
	}
	if r.userAgent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userAgent", r.userAgent, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.operationIds != nil {
		t := *r.operationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", t, "multi")
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

type ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest struct {
	ctx context.Context
	ApiService *ApiRequestsAPIService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest) T0(t0 string) ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest) T1(t1 string) ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest) Timespan(timespan float32) ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest {
	r.timespan = &timespan
	return r
}

func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest) Execute() (*GetOrganizationApiRequestsOverview200Response, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewExecute(r)
}

/*
GetOrganizationApiRequestsOverview Return an aggregated overview of API requests data

Return an aggregated overview of API requests data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest
*/
func (a *ApiRequestsAPIService) GetOrganizationApiRequestsOverview(ctx context.Context, organizationId string) ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest {
	return ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return GetOrganizationApiRequestsOverview200Response
func (a *ApiRequestsAPIService) GetOrganizationApiRequestsOverviewExecute(r ApiRequestsAPIGetOrganizationApiRequestsOverviewRequest) (*GetOrganizationApiRequestsOverview200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationApiRequestsOverview200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsAPIService.GetOrganizationApiRequestsOverview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

type ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct {
	ctx context.Context
	ApiService *ApiRequestsAPIService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
	version *int32
	operationIds *[]string
	sourceIps *[]string
	adminIds *[]string
	userAgent *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T0(t0 string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T1(t1 string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Timespan(timespan float32) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Interval(interval int32) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.interval = &interval
	return r
}

// Filter by API version of the endpoint. Allowable values are: [0, 1]
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Version(version int32) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.version = &version
	return r
}

// Filter by operation ID of the endpoint
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) OperationIds(operationIds []string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.operationIds = &operationIds
	return r
}

// Filter by source IP that made the API request
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) SourceIps(sourceIps []string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.sourceIps = &sourceIps
	return r
}

// Filter by admin ID of user that made the API request
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) AdminIds(adminIds []string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.adminIds = &adminIds
	return r
}

// Filter by user agent string for API request. This will filter by a complete or partial match.
func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) UserAgent(userAgent string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.userAgent = &userAgent
	return r
}

func (r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Execute() ([]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r)
}

/*
GetOrganizationApiRequestsOverviewResponseCodesByInterval Tracks organizations' API requests by response code across a given time period

Tracks organizations' API requests by response code across a given time period

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest
*/
func (a *ApiRequestsAPIService) GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx context.Context, organizationId string) ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	return ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
func (a *ApiRequestsAPIService) GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r ApiRequestsAPIGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) ([]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsAPIService.GetOrganizationApiRequestsOverviewResponseCodesByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.operationIds != nil {
		t := *r.operationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", t, "multi")
		}
	}
	if r.sourceIps != nil {
		t := *r.sourceIps
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sourceIps", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sourceIps", t, "multi")
		}
	}
	if r.adminIds != nil {
		t := *r.adminIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "adminIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "adminIds", t, "multi")
		}
	}
	if r.userAgent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userAgent", r.userAgent, "")
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

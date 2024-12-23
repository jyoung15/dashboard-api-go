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


// ByMetricAPIService ByMetricAPI service
type ByMetricAPIService service

type ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest struct {
	ctx context.Context
	ApiService *ByMetricAPIService
	networkId string
}

func (r ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest) Execute() (*GetNetworkSensorAlertsCurrentOverviewByMetric200Response, *http.Response, error) {
	return r.ApiService.GetNetworkSensorAlertsCurrentOverviewByMetricExecute(r)
}

/*
GetNetworkSensorAlertsCurrentOverviewByMetric Return an overview of currently alerting sensors by metric

Return an overview of currently alerting sensors by metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest
*/
func (a *ByMetricAPIService) GetNetworkSensorAlertsCurrentOverviewByMetric(ctx context.Context, networkId string) ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest {
	return ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSensorAlertsCurrentOverviewByMetric200Response
func (a *ByMetricAPIService) GetNetworkSensorAlertsCurrentOverviewByMetricExecute(r ByMetricAPIGetNetworkSensorAlertsCurrentOverviewByMetricRequest) (*GetNetworkSensorAlertsCurrentOverviewByMetric200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSensorAlertsCurrentOverviewByMetric200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByMetricAPIService.GetNetworkSensorAlertsCurrentOverviewByMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sensor/alerts/current/overview/byMetric"
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

type ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest struct {
	ctx context.Context
	ApiService *ByMetricAPIService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) T0(t0 string) ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) T1(t1 string) ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) Timespan(timespan float32) ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 86400, 604800. The default is 604800.
func (r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) Interval(interval int32) ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest {
	r.interval = &interval
	return r
}

func (r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) Execute() ([]GetNetworkSensorAlertsOverviewByMetric200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSensorAlertsOverviewByMetricExecute(r)
}

/*
GetNetworkSensorAlertsOverviewByMetric Return an overview of alert occurrences over a timespan, by metric

Return an overview of alert occurrences over a timespan, by metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest
*/
func (a *ByMetricAPIService) GetNetworkSensorAlertsOverviewByMetric(ctx context.Context, networkId string) ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest {
	return ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSensorAlertsOverviewByMetric200ResponseInner
func (a *ByMetricAPIService) GetNetworkSensorAlertsOverviewByMetricExecute(r ByMetricAPIGetNetworkSensorAlertsOverviewByMetricRequest) ([]GetNetworkSensorAlertsOverviewByMetric200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSensorAlertsOverviewByMetric200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByMetricAPIService.GetNetworkSensorAlertsOverviewByMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sensor/alerts/overview/byMetric"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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

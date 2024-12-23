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


// DesktopLogsAPIService DesktopLogsAPI service
type DesktopLogsAPIService service

type DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest struct {
	ctx context.Context
	ApiService *DesktopLogsAPIService
	networkId string
	deviceId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest) PerPage(perPage int32) DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest) StartingAfter(startingAfter string) DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest) EndingBefore(endingBefore string) DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest) Execute() ([]GetNetworkSmDeviceDesktopLogs200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSmDeviceDesktopLogsExecute(r)
}

/*
GetNetworkSmDeviceDesktopLogs Return historical records of various Systems Manager network connection details for desktop devices.

Return historical records of various Systems Manager network connection details for desktop devices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param deviceId Device ID
 @return DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest
*/
func (a *DesktopLogsAPIService) GetNetworkSmDeviceDesktopLogs(ctx context.Context, networkId string, deviceId string) DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest {
	return DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return []GetNetworkSmDeviceDesktopLogs200ResponseInner
func (a *DesktopLogsAPIService) GetNetworkSmDeviceDesktopLogsExecute(r DesktopLogsAPIGetNetworkSmDeviceDesktopLogsRequest) ([]GetNetworkSmDeviceDesktopLogs200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSmDeviceDesktopLogs200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DesktopLogsAPIService.GetNetworkSmDeviceDesktopLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/sm/devices/{deviceId}/desktopLogs"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

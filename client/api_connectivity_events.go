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


// ConnectivityEventsAPIService ConnectivityEventsAPI service
type ConnectivityEventsAPIService service

type ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest struct {
	ctx context.Context
	ApiService *ConnectivityEventsAPIService
	networkId string
	clientId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	types *[]string
	band *string
	ssidNumber *int32
	includedSeverities *[]string
	deviceSerial *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) PerPage(perPage int32) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) StartingAfter(startingAfter string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) EndingBefore(endingBefore string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) T0(t0 string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) T1(t1 string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) Timespan(timespan float32) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.timespan = &timespan
	return r
}

// A list of event types to include. If not specified, events of all types will be returned. Valid types are &#39;assoc&#39;, &#39;disassoc&#39;, &#39;auth&#39;, &#39;deauth&#39;, &#39;dns&#39;, &#39;dhcp&#39;, &#39;roam&#39;, &#39;connection&#39; and/or &#39;sticky&#39;.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) Types(types []string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.types = &types
	return r
}

// Filter results by band. Valid bands are &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) Band(band string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.band = &band
	return r
}

// Filter results by SSID. If not specified, events for all SSIDs will be returned.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) SsidNumber(ssidNumber int32) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.ssidNumber = &ssidNumber
	return r
}

// A list of severities to include. If not specified, events of all severities will be returned. Valid severities are &#39;good&#39;, &#39;info&#39;, &#39;warn&#39; and/or &#39;bad&#39;.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) IncludedSeverities(includedSeverities []string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.includedSeverities = &includedSeverities
	return r
}

// Filter results by an AP&#39;s serial number.
func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) DeviceSerial(deviceSerial string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	r.deviceSerial = &deviceSerial
	return r
}

func (r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) Execute() ([]GetNetworkWirelessClientConnectivityEvents200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientConnectivityEventsExecute(r)
}

/*
GetNetworkWirelessClientConnectivityEvents List the wireless connectivity events for a client within a network in the timespan.

List the wireless connectivity events for a client within a network in the timespan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest
*/
func (a *ConnectivityEventsAPIService) GetNetworkWirelessClientConnectivityEvents(ctx context.Context, networkId string, clientId string) ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest {
	return ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessClientConnectivityEvents200ResponseInner
func (a *ConnectivityEventsAPIService) GetNetworkWirelessClientConnectivityEventsExecute(r ConnectivityEventsAPIGetNetworkWirelessClientConnectivityEventsRequest) ([]GetNetworkWirelessClientConnectivityEvents200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessClientConnectivityEvents200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityEventsAPIService.GetNetworkWirelessClientConnectivityEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/connectivityEvents"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

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
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssidNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssidNumber", r.ssidNumber, "")
	}
	if r.includedSeverities != nil {
		t := *r.includedSeverities
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includedSeverities", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includedSeverities", t, "multi")
		}
	}
	if r.deviceSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceSerial", r.deviceSerial, "")
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

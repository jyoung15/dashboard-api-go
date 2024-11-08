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


// Ipv6APIService Ipv6API service
type Ipv6APIService service

type Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct {
	ctx context.Context
	ApiService *Ipv6APIService
	serial string
	updateDeviceWirelessAlternateManagementInterfaceIpv6Request *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request
}

func (r Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(updateDeviceWirelessAlternateManagementInterfaceIpv6Request UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	r.updateDeviceWirelessAlternateManagementInterfaceIpv6Request = &updateDeviceWirelessAlternateManagementInterfaceIpv6Request
	return r
}

func (r Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Execute() (*UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response, *http.Response, error) {
	return r.ApiService.UpdateDeviceWirelessAlternateManagementInterfaceIpv6Execute(r)
}

/*
UpdateDeviceWirelessAlternateManagementInterfaceIpv6 Update alternate management interface IPv6 address

Update alternate management interface IPv6 address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request
*/
func (a *Ipv6APIService) UpdateDeviceWirelessAlternateManagementInterfaceIpv6(ctx context.Context, serial string) Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	return Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
func (a *Ipv6APIService) UpdateDeviceWirelessAlternateManagementInterfaceIpv6Execute(r Ipv6APIUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) (*UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Ipv6APIService.UpdateDeviceWirelessAlternateManagementInterfaceIpv6")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/wireless/alternateManagementInterface/ipv6"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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
	localVarPostBody = r.updateDeviceWirelessAlternateManagementInterfaceIpv6Request
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

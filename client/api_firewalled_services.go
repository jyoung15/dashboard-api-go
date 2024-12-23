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


// FirewalledServicesAPIService FirewalledServicesAPI service
type FirewalledServicesAPIService service

type FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest struct {
	ctx context.Context
	ApiService *FirewalledServicesAPIService
	networkId string
	service string
}

func (r FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest) Execute() (*GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallFirewalledServiceExecute(r)
}

/*
GetNetworkApplianceFirewallFirewalledService Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')

Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param service Service
 @return FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest
*/
func (a *FirewalledServicesAPIService) GetNetworkApplianceFirewallFirewalledService(ctx context.Context, networkId string, service string) FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest {
	return FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		service: service,
	}
}

// Execute executes the request
//  @return GetNetworkApplianceFirewallFirewalledServices200ResponseInner
func (a *FirewalledServicesAPIService) GetNetworkApplianceFirewallFirewalledServiceExecute(r FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServiceRequest) (*GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewalledServicesAPIService.GetNetworkApplianceFirewallFirewalledService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", url.PathEscape(parameterValueToString(r.service, "service")), -1)

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

type FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest struct {
	ctx context.Context
	ApiService *FirewalledServicesAPIService
	networkId string
}

func (r FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest) Execute() ([]GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceFirewallFirewalledServicesExecute(r)
}

/*
GetNetworkApplianceFirewallFirewalledServices List the appliance services and their accessibility rules

List the appliance services and their accessibility rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest
*/
func (a *FirewalledServicesAPIService) GetNetworkApplianceFirewallFirewalledServices(ctx context.Context, networkId string) FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest {
	return FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkApplianceFirewallFirewalledServices200ResponseInner
func (a *FirewalledServicesAPIService) GetNetworkApplianceFirewallFirewalledServicesExecute(r FirewalledServicesAPIGetNetworkApplianceFirewallFirewalledServicesRequest) ([]GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewalledServicesAPIService.GetNetworkApplianceFirewallFirewalledServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/firewalledServices"
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

type FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest struct {
	ctx context.Context
	ApiService *FirewalledServicesAPIService
	networkId string
	service string
	updateNetworkApplianceFirewallFirewalledServiceRequest *UpdateNetworkApplianceFirewallFirewalledServiceRequest
}

func (r FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest) UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest UpdateNetworkApplianceFirewallFirewalledServiceRequest) FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest {
	r.updateNetworkApplianceFirewallFirewalledServiceRequest = &updateNetworkApplianceFirewallFirewalledServiceRequest
	return r
}

func (r FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest) Execute() (*GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceFirewallFirewalledServiceExecute(r)
}

/*
UpdateNetworkApplianceFirewallFirewalledService Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')

Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param service Service
 @return FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest
*/
func (a *FirewalledServicesAPIService) UpdateNetworkApplianceFirewallFirewalledService(ctx context.Context, networkId string, service string) FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest {
	return FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		service: service,
	}
}

// Execute executes the request
//  @return GetNetworkApplianceFirewallFirewalledServices200ResponseInner
func (a *FirewalledServicesAPIService) UpdateNetworkApplianceFirewallFirewalledServiceExecute(r FirewalledServicesAPIUpdateNetworkApplianceFirewallFirewalledServiceRequest) (*GetNetworkApplianceFirewallFirewalledServices200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirewalledServicesAPIService.UpdateNetworkApplianceFirewallFirewalledService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/firewall/firewalledServices/{service}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", url.PathEscape(parameterValueToString(r.service, "service")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkApplianceFirewallFirewalledServiceRequest == nil {
		return localVarReturnValue, nil, reportError("updateNetworkApplianceFirewallFirewalledServiceRequest is required and must be specified")
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
	localVarPostBody = r.updateNetworkApplianceFirewallFirewalledServiceRequest
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

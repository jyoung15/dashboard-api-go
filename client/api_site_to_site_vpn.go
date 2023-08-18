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
)


// SiteToSiteVpnAPIService SiteToSiteVpnAPI service
type SiteToSiteVpnAPIService service

type SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest struct {
	ctx context.Context
	ApiService *SiteToSiteVpnAPIService
	networkId string
}

func (r SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest) Execute() (*GetNetworkApplianceVpnSiteToSiteVpn200Response, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceVpnSiteToSiteVpnExecute(r)
}

/*
GetNetworkApplianceVpnSiteToSiteVpn Return the site-to-site VPN settings of a network

Return the site-to-site VPN settings of a network. Only valid for MX networks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest
*/
func (a *SiteToSiteVpnAPIService) GetNetworkApplianceVpnSiteToSiteVpn(ctx context.Context, networkId string) SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest {
	return SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkApplianceVpnSiteToSiteVpn200Response
func (a *SiteToSiteVpnAPIService) GetNetworkApplianceVpnSiteToSiteVpnExecute(r SiteToSiteVpnAPIGetNetworkApplianceVpnSiteToSiteVpnRequest) (*GetNetworkApplianceVpnSiteToSiteVpn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkApplianceVpnSiteToSiteVpn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteToSiteVpnAPIService.GetNetworkApplianceVpnSiteToSiteVpn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/vpn/siteToSiteVpn"
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

type SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct {
	ctx context.Context
	ApiService *SiteToSiteVpnAPIService
	networkId string
	updateNetworkApplianceVpnSiteToSiteVpnRequest *UpdateNetworkApplianceVpnSiteToSiteVpnRequest
}

func (r SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest) UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	r.updateNetworkApplianceVpnSiteToSiteVpnRequest = &updateNetworkApplianceVpnSiteToSiteVpnRequest
	return r
}

func (r SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest) Execute() (*GetNetworkApplianceVpnSiteToSiteVpn200Response, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceVpnSiteToSiteVpnExecute(r)
}

/*
UpdateNetworkApplianceVpnSiteToSiteVpn Update the site-to-site VPN settings of a network

Update the site-to-site VPN settings of a network. Only valid for MX networks in NAT mode.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest
*/
func (a *SiteToSiteVpnAPIService) UpdateNetworkApplianceVpnSiteToSiteVpn(ctx context.Context, networkId string) SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	return SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkApplianceVpnSiteToSiteVpn200Response
func (a *SiteToSiteVpnAPIService) UpdateNetworkApplianceVpnSiteToSiteVpnExecute(r SiteToSiteVpnAPIUpdateNetworkApplianceVpnSiteToSiteVpnRequest) (*GetNetworkApplianceVpnSiteToSiteVpn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkApplianceVpnSiteToSiteVpn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteToSiteVpnAPIService.UpdateNetworkApplianceVpnSiteToSiteVpn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/vpn/siteToSiteVpn"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateNetworkApplianceVpnSiteToSiteVpnRequest == nil {
		return localVarReturnValue, nil, reportError("updateNetworkApplianceVpnSiteToSiteVpnRequest is required and must be specified")
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
	localVarPostBody = r.updateNetworkApplianceVpnSiteToSiteVpnRequest
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

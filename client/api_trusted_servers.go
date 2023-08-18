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


// TrustedServersAPIService TrustedServersAPI service
type TrustedServersAPIService service

type TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *TrustedServersAPIService
	networkId string
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
}

func (r TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest = &createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
	return r
}

func (r TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network

Add a server to be trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *TrustedServersAPIService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string) TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *TrustedServersAPIService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r TrustedServersAPICreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedServersAPIService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
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

type TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *TrustedServersAPIService
	networkId string
	trustedServerId string
}

func (r TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network

Remove a server from being trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *TrustedServersAPIService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
func (a *TrustedServersAPIService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r TrustedServersAPIDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedServersAPIService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterValueToString(r.trustedServerId, "trustedServerId")), -1)

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

type TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct {
	ctx context.Context
	ApiService *TrustedServersAPIService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) PerPage(perPage int32) TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) StartingAfter(startingAfter string) TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) EndingBefore(endingBefore string) TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) Execute() ([]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network

Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as allow listed snoop entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest
*/
func (a *TrustedServersAPIService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx context.Context, networkId string) TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	return TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *TrustedServersAPIService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r TrustedServersAPIGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) ([]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedServersAPIService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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

type TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *TrustedServersAPIService
	networkId string
	trustedServerId string
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
}

func (r TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest = &updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
	return r
}

func (r TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network

Update a server that is trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *TrustedServersAPIService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *TrustedServersAPIService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r TrustedServersAPIUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrustedServersAPIService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterValueToString(r.trustedServerId, "trustedServerId")), -1)

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
	localVarPostBody = r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
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

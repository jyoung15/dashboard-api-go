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


// StagesAPIService StagesAPI service
type StagesAPIService service

type StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest struct {
	ctx context.Context
	ApiService *StagesAPIService
	networkId string
}

func (r StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest) Execute() ([]GetNetworkFirmwareUpgradesStagedStages200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkFirmwareUpgradesStagedStagesExecute(r)
}

/*
GetNetworkFirmwareUpgradesStagedStages Order of Staged Upgrade Groups in a network

Order of Staged Upgrade Groups in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest
*/
func (a *StagesAPIService) GetNetworkFirmwareUpgradesStagedStages(ctx context.Context, networkId string) StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest {
	return StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
func (a *StagesAPIService) GetNetworkFirmwareUpgradesStagedStagesExecute(r StagesAPIGetNetworkFirmwareUpgradesStagedStagesRequest) ([]GetNetworkFirmwareUpgradesStagedStages200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StagesAPIService.GetNetworkFirmwareUpgradesStagedStages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/firmwareUpgrades/staged/stages"
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

type StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest struct {
	ctx context.Context
	ApiService *StagesAPIService
	networkId string
	updateNetworkFirmwareUpgradesStagedStagesRequest *UpdateNetworkFirmwareUpgradesStagedStagesRequest
}

func (r StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest) UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest UpdateNetworkFirmwareUpgradesStagedStagesRequest) StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest {
	r.updateNetworkFirmwareUpgradesStagedStagesRequest = &updateNetworkFirmwareUpgradesStagedStagesRequest
	return r
}

func (r StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest) Execute() ([]GetNetworkFirmwareUpgradesStagedStages200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkFirmwareUpgradesStagedStagesExecute(r)
}

/*
UpdateNetworkFirmwareUpgradesStagedStages Assign Staged Upgrade Group order in the sequence.

Assign Staged Upgrade Group order in the sequence.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest
*/
func (a *StagesAPIService) UpdateNetworkFirmwareUpgradesStagedStages(ctx context.Context, networkId string) StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest {
	return StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
func (a *StagesAPIService) UpdateNetworkFirmwareUpgradesStagedStagesExecute(r StagesAPIUpdateNetworkFirmwareUpgradesStagedStagesRequest) ([]GetNetworkFirmwareUpgradesStagedStages200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StagesAPIService.UpdateNetworkFirmwareUpgradesStagedStages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/firmwareUpgrades/staged/stages"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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
	localVarPostBody = r.updateNetworkFirmwareUpgradesStagedStagesRequest
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

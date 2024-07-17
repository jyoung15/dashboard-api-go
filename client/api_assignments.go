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


// AssignmentsAPIService AssignmentsAPI service
type AssignmentsAPIService service

type AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest struct {
	ctx context.Context
	ApiService *AssignmentsAPIService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	serials *[]string
	productTypes *[]string
	stackIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) PerPage(perPage int32) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) StartingAfter(startingAfter string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) EndingBefore(endingBefore string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) Serials(serials []string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter devices by product types.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) ProductTypes(productTypes []string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter devices by Switch Stack ids.
func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) StackIds(stackIds []string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.stackIds = &stackIds
	return r
}

func (r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) Execute() ([]GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkVlanProfilesAssignmentsByDeviceExecute(r)
}

/*
GetNetworkVlanProfilesAssignmentsByDevice Get the assigned VLAN Profiles for devices in a network

Get the assigned VLAN Profiles for devices in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest
*/
func (a *AssignmentsAPIService) GetNetworkVlanProfilesAssignmentsByDevice(ctx context.Context, networkId string) AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	return AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
func (a *AssignmentsAPIService) GetNetworkVlanProfilesAssignmentsByDeviceExecute(r AssignmentsAPIGetNetworkVlanProfilesAssignmentsByDeviceRequest) ([]GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsAPIService.GetNetworkVlanProfilesAssignmentsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/vlanProfiles/assignments/byDevice"
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
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.productTypes != nil {
		t := *r.productTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", t, "multi")
		}
	}
	if r.stackIds != nil {
		t := *r.stackIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "stackIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "stackIds", t, "multi")
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

type AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest struct {
	ctx context.Context
	ApiService *AssignmentsAPIService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) PerPage(perPage int32) AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) StartingAfter(startingAfter string) AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) EndingBefore(endingBefore string) AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter Sentry Policies by Network Id
func (r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) NetworkIds(networkIds []string) AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

func (r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) Execute() ([]GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r)
}

/*
GetOrganizationSmSentryPoliciesAssignmentsByNetwork List the Sentry Policies for an organization ordered in ascending order of priority

List the Sentry Policies for an organization ordered in ascending order of priority

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest
*/
func (a *AssignmentsAPIService) GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx context.Context, organizationId string) AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	return AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner
func (a *AssignmentsAPIService) GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r AssignmentsAPIGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) ([]GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsAPIService.GetOrganizationSmSentryPoliciesAssignmentsByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
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

type AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest struct {
	ctx context.Context
	ApiService *AssignmentsAPIService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	name *string
	mac *string
	serial *string
	model *string
	macs *[]string
	serials *[]string
	models *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) PerPage(perPage int32) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) StartingAfter(startingAfter string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) EndingBefore(endingBefore string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter devices by network.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) NetworkIds(networkIds []string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) ProductTypes(productTypes []string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Name(name string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.name = &name
	return r
}

// Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Mac(mac string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.mac = &mac
	return r
}

// Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Serial(serial string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.serial = &serial
	return r
}

// Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Model(model string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.model = &model
	return r
}

// Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Macs(macs []string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Serials(serials []string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match.
func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Models(models []string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.models = &models
	return r
}

func (r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Execute() ([]GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessRfProfilesAssignmentsByDeviceExecute(r)
}

/*
GetOrganizationWirelessRfProfilesAssignmentsByDevice List the RF profiles of an organization by device

List the RF profiles of an organization by device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest
*/
func (a *AssignmentsAPIService) GetOrganizationWirelessRfProfilesAssignmentsByDevice(ctx context.Context, organizationId string) AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	return AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner
func (a *AssignmentsAPIService) GetOrganizationWirelessRfProfilesAssignmentsByDeviceExecute(r AssignmentsAPIGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) ([]GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsAPIService.GetOrganizationWirelessRfProfilesAssignmentsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/rfProfiles/assignments/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.productTypes != nil {
		t := *r.productTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "productTypes", t, "multi")
		}
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.serial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "serial", r.serial, "")
	}
	if r.model != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "model", r.model, "")
	}
	if r.macs != nil {
		t := *r.macs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "macs", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "macs", t, "multi")
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.models != nil {
		t := *r.models
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "models", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "models", t, "multi")
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

type AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest struct {
	ctx context.Context
	ApiService *AssignmentsAPIService
	networkId string
	reassignNetworkVlanProfilesAssignmentsRequest *ReassignNetworkVlanProfilesAssignmentsRequest
}

func (r AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest) ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest ReassignNetworkVlanProfilesAssignmentsRequest) AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest {
	r.reassignNetworkVlanProfilesAssignmentsRequest = &reassignNetworkVlanProfilesAssignmentsRequest
	return r
}

func (r AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest) Execute() (*ReassignNetworkVlanProfilesAssignments200Response, *http.Response, error) {
	return r.ApiService.ReassignNetworkVlanProfilesAssignmentsExecute(r)
}

/*
ReassignNetworkVlanProfilesAssignments Update the assigned VLAN Profile for devices in a network

Update the assigned VLAN Profile for devices in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest
*/
func (a *AssignmentsAPIService) ReassignNetworkVlanProfilesAssignments(ctx context.Context, networkId string) AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest {
	return AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return ReassignNetworkVlanProfilesAssignments200Response
func (a *AssignmentsAPIService) ReassignNetworkVlanProfilesAssignmentsExecute(r AssignmentsAPIReassignNetworkVlanProfilesAssignmentsRequest) (*ReassignNetworkVlanProfilesAssignments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReassignNetworkVlanProfilesAssignments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsAPIService.ReassignNetworkVlanProfilesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/vlanProfiles/assignments/reassign"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reassignNetworkVlanProfilesAssignmentsRequest == nil {
		return localVarReturnValue, nil, reportError("reassignNetworkVlanProfilesAssignmentsRequest is required and must be specified")
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
	localVarPostBody = r.reassignNetworkVlanProfilesAssignmentsRequest
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

type AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest struct {
	ctx context.Context
	ApiService *AssignmentsAPIService
	organizationId string
	updateOrganizationSmSentryPoliciesAssignmentsRequest *UpdateOrganizationSmSentryPoliciesAssignmentsRequest
}

func (r AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest) UpdateOrganizationSmSentryPoliciesAssignmentsRequest(updateOrganizationSmSentryPoliciesAssignmentsRequest UpdateOrganizationSmSentryPoliciesAssignmentsRequest) AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	r.updateOrganizationSmSentryPoliciesAssignmentsRequest = &updateOrganizationSmSentryPoliciesAssignmentsRequest
	return r
}

func (r AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest) Execute() (*UpdateOrganizationSmSentryPoliciesAssignments200Response, *http.Response, error) {
	return r.ApiService.UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r)
}

/*
UpdateOrganizationSmSentryPoliciesAssignments Update an Organizations Sentry Policies using the provided list

Update an Organizations Sentry Policies using the provided list. Sentry Policies are ordered in descending order of priority (i.e. highest priority at the bottom, this is opposite the Dashboard UI). Policies not present in the request will be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest
*/
func (a *AssignmentsAPIService) UpdateOrganizationSmSentryPoliciesAssignments(ctx context.Context, organizationId string) AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	return AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return UpdateOrganizationSmSentryPoliciesAssignments200Response
func (a *AssignmentsAPIService) UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r AssignmentsAPIUpdateOrganizationSmSentryPoliciesAssignmentsRequest) (*UpdateOrganizationSmSentryPoliciesAssignments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateOrganizationSmSentryPoliciesAssignments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsAPIService.UpdateOrganizationSmSentryPoliciesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrganizationSmSentryPoliciesAssignmentsRequest == nil {
		return localVarReturnValue, nil, reportError("updateOrganizationSmSentryPoliciesAssignmentsRequest is required and must be specified")
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
	localVarPostBody = r.updateOrganizationSmSentryPoliciesAssignmentsRequest
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

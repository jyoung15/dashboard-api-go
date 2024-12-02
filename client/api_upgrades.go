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


// UpgradesAPIService UpgradesAPI service
type UpgradesAPIService service

type UpgradesAPIGetOrganizationFirmwareUpgradesRequest struct {
	ctx context.Context
	ApiService *UpgradesAPIService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	status *[]string
	productTypes *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) PerPage(perPage int32) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) StartingAfter(startingAfter string) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) EndingBefore(endingBefore string) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter the upgrade by status.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) Status(status []string) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	r.status = &status
	return r
}

// Optional parameter to filter the upgrade by product type.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) ProductTypes(productTypes []string) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	r.productTypes = &productTypes
	return r
}

func (r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) Execute() ([]GetOrganizationFirmwareUpgrades200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesExecute(r)
}

/*
GetOrganizationFirmwareUpgrades Get firmware upgrade information for an organization

Get firmware upgrade information for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return UpgradesAPIGetOrganizationFirmwareUpgradesRequest
*/
func (a *UpgradesAPIService) GetOrganizationFirmwareUpgrades(ctx context.Context, organizationId string) UpgradesAPIGetOrganizationFirmwareUpgradesRequest {
	return UpgradesAPIGetOrganizationFirmwareUpgradesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationFirmwareUpgrades200ResponseInner
func (a *UpgradesAPIService) GetOrganizationFirmwareUpgradesExecute(r UpgradesAPIGetOrganizationFirmwareUpgradesRequest) ([]GetOrganizationFirmwareUpgrades200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationFirmwareUpgrades200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradesAPIService.GetOrganizationFirmwareUpgrades")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades"
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
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
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

type UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest struct {
	ctx context.Context
	ApiService *UpgradesAPIService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	macs *[]string
	firmwareUpgradeBatchIds *[]string
	upgradeStatuses *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) PerPage(perPage int32) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) StartingAfter(startingAfter string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) EndingBefore(endingBefore string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter by network
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) NetworkIds(networkIds []string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) Serials(serials []string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) Macs(macs []string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter by firmware upgrade batch ids.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeBatchIds = &firmwareUpgradeBatchIds
	return r
}

// Optional parameter to filter by firmware upgrade statuses.
func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) UpgradeStatuses(upgradeStatuses []string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.upgradeStatuses = &upgradeStatuses
	return r
}

func (r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) Execute() ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesByDeviceExecute(r)
}

/*
GetOrganizationFirmwareUpgradesByDevice Get firmware upgrade status for the filtered devices

Get firmware upgrade status for the filtered devices. This endpoint currently only supports Meraki switches.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest
*/
func (a *UpgradesAPIService) GetOrganizationFirmwareUpgradesByDevice(ctx context.Context, organizationId string) UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest {
	return UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
func (a *UpgradesAPIService) GetOrganizationFirmwareUpgradesByDeviceExecute(r UpgradesAPIGetOrganizationFirmwareUpgradesByDeviceRequest) ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradesAPIService.GetOrganizationFirmwareUpgradesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades/byDevice"
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
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds[]", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds[]", t, "multi")
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
	if r.firmwareUpgradeBatchIds != nil {
		t := *r.firmwareUpgradeBatchIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "firmwareUpgradeBatchIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "firmwareUpgradeBatchIds", t, "multi")
		}
	}
	if r.upgradeStatuses != nil {
		t := *r.upgradeStatuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "upgradeStatuses", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "upgradeStatuses", t, "multi")
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

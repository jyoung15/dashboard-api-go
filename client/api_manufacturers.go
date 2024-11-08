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


// ManufacturersAPIService ManufacturersAPI service
type ManufacturersAPIService service

type ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest struct {
	ctx context.Context
	ApiService *ManufacturersAPIService
	organizationId string
	networkTag *string
	deviceTag *string
	networkId *string
	quantity *int32
	ssidName *string
	usageUplink *string
	t0 *string
	t1 *string
	timespan *float32
}

// Match result to an exact network tag
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) NetworkTag(networkTag string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) DeviceTag(deviceTag string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.deviceTag = &deviceTag
	return r
}

// Match result to an exact network id
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) NetworkId(networkId string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) Quantity(quantity int32) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) SsidName(ssidName string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) UsageUplink(usageUplink string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) T0(t0 string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) T1(t1 string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day.
func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) Timespan(timespan float32) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) Execute() ([]GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopClientsManufacturersByUsageExecute(r)
}

/*
GetOrganizationSummaryTopClientsManufacturersByUsage Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.

Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest
*/
func (a *ManufacturersAPIService) GetOrganizationSummaryTopClientsManufacturersByUsage(ctx context.Context, organizationId string) ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	return ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
func (a *ManufacturersAPIService) GetOrganizationSummaryTopClientsManufacturersByUsageExecute(r ManufacturersAPIGetOrganizationSummaryTopClientsManufacturersByUsageRequest) ([]GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManufacturersAPIService.GetOrganizationSummaryTopClientsManufacturersByUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/clients/manufacturers/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkTag", r.networkTag, "")
	}
	if r.deviceTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceTag", r.deviceTag, "")
	}
	if r.networkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "networkId", r.networkId, "")
	}
	if r.quantity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "")
	}
	if r.ssidName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssidName", r.ssidName, "")
	}
	if r.usageUplink != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "usageUplink", r.usageUplink, "")
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

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


// SubscriptionsAPIService SubscriptionsAPI service
type SubscriptionsAPIService service

type SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest struct {
	ctx context.Context
	ApiService *SubscriptionsAPIService
	subscriptionId string
	bindAdministeredLicensingSubscriptionSubscriptionRequest *BindAdministeredLicensingSubscriptionSubscriptionRequest
	validate *bool
}

func (r SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest) BindAdministeredLicensingSubscriptionSubscriptionRequest(bindAdministeredLicensingSubscriptionSubscriptionRequest BindAdministeredLicensingSubscriptionSubscriptionRequest) SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest {
	r.bindAdministeredLicensingSubscriptionSubscriptionRequest = &bindAdministeredLicensingSubscriptionSubscriptionRequest
	return r
}

// Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results.
func (r SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest) Validate(validate bool) SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest {
	r.validate = &validate
	return r
}

func (r SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest) Execute() (*BindAdministeredLicensingSubscriptionSubscription200Response, *http.Response, error) {
	return r.ApiService.BindAdministeredLicensingSubscriptionSubscriptionExecute(r)
}

/*
BindAdministeredLicensingSubscriptionSubscription Bind networks to a subscription

Bind networks to a subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription ID
 @return SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest
*/
func (a *SubscriptionsAPIService) BindAdministeredLicensingSubscriptionSubscription(ctx context.Context, subscriptionId string) SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest {
	return SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return BindAdministeredLicensingSubscriptionSubscription200Response
func (a *SubscriptionsAPIService) BindAdministeredLicensingSubscriptionSubscriptionExecute(r SubscriptionsAPIBindAdministeredLicensingSubscriptionSubscriptionRequest) (*BindAdministeredLicensingSubscriptionSubscription200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BindAdministeredLicensingSubscriptionSubscription200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.BindAdministeredLicensingSubscriptionSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions/{subscriptionId}/bind"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterValueToString(r.subscriptionId, "subscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bindAdministeredLicensingSubscriptionSubscriptionRequest == nil {
		return localVarReturnValue, nil, reportError("bindAdministeredLicensingSubscriptionSubscriptionRequest is required and must be specified")
	}

	if r.validate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validate", r.validate, "")
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
	localVarPostBody = r.bindAdministeredLicensingSubscriptionSubscriptionRequest
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

type SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest struct {
	ctx context.Context
	ApiService *SubscriptionsAPIService
	claimAdministeredLicensingSubscriptionSubscriptionsRequest *ClaimAdministeredLicensingSubscriptionSubscriptionsRequest
	validate *bool
}

func (r SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest) ClaimAdministeredLicensingSubscriptionSubscriptionsRequest(claimAdministeredLicensingSubscriptionSubscriptionsRequest ClaimAdministeredLicensingSubscriptionSubscriptionsRequest) SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.claimAdministeredLicensingSubscriptionSubscriptionsRequest = &claimAdministeredLicensingSubscriptionSubscriptionsRequest
	return r
}

// Check if the provided claim key is valid and can be claimed into the organization.
func (r SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest) Validate(validate bool) SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.validate = &validate
	return r
}

func (r SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest) Execute() (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	return r.ApiService.ClaimAdministeredLicensingSubscriptionSubscriptionsExecute(r)
}

/*
ClaimAdministeredLicensingSubscriptionSubscriptions Claim a subscription into an organization.

Claim a subscription into an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest
*/
func (a *SubscriptionsAPIService) ClaimAdministeredLicensingSubscriptionSubscriptions(ctx context.Context) SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest {
	return SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
func (a *SubscriptionsAPIService) ClaimAdministeredLicensingSubscriptionSubscriptionsExecute(r SubscriptionsAPIClaimAdministeredLicensingSubscriptionSubscriptionsRequest) (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.ClaimAdministeredLicensingSubscriptionSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions/claim"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.claimAdministeredLicensingSubscriptionSubscriptionsRequest == nil {
		return localVarReturnValue, nil, reportError("claimAdministeredLicensingSubscriptionSubscriptionsRequest is required and must be specified")
	}

	if r.validate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validate", r.validate, "")
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
	localVarPostBody = r.claimAdministeredLicensingSubscriptionSubscriptionsRequest
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

type SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest struct {
	ctx context.Context
	ApiService *SubscriptionsAPIService
	perPage *int32
	startingAfter *string
	endingBefore *string
	subscriptionIds *[]string
	organizationIds *[]string
	statuses *[]string
	productTypes *[]string
	startDate *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter
	endDate *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) PerPage(perPage int32) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) StartingAfter(startingAfter string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) EndingBefore(endingBefore string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.endingBefore = &endingBefore
	return r
}

// List of subscription ids to fetch
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) SubscriptionIds(subscriptionIds []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.subscriptionIds = &subscriptionIds
	return r
}

// Organizations to get associated subscriptions for
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) OrganizationIds(organizationIds []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.organizationIds = &organizationIds
	return r
}

// List of statuses that returned subscriptions can have
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) Statuses(statuses []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.statuses = &statuses
	return r
}

// List of product types that returned subscriptions need to have entitlements for.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) ProductTypes(productTypes []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.productTypes = &productTypes
	return r
}

// Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use &#39;startDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) StartDate(startDate GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.startDate = &startDate
	return r
}

// Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use &#39;endDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte.
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) EndDate(endDate GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	r.endDate = &endDate
	return r
}

func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) Execute() ([]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	return r.ApiService.GetAdministeredLicensingSubscriptionSubscriptionsExecute(r)
}

/*
GetAdministeredLicensingSubscriptionSubscriptions List available subscriptions

List available subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest
*/
func (a *SubscriptionsAPIService) GetAdministeredLicensingSubscriptionSubscriptions(ctx context.Context) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest {
	return SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
func (a *SubscriptionsAPIService) GetAdministeredLicensingSubscriptionSubscriptionsExecute(r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsRequest) ([]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.GetAdministeredLicensingSubscriptionSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions"

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
	if r.subscriptionIds != nil {
		t := *r.subscriptionIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionIds", t, "multi")
		}
	}
	if r.organizationIds != nil {
		t := *r.organizationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "organizationIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "organizationIds", t, "multi")
		}
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "multi")
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
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

type SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest struct {
	ctx context.Context
	ApiService *SubscriptionsAPIService
	organizationIds *[]string
	subscriptionIds *[]string
}

// Organizations to get subscription compliance information for
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest) OrganizationIds(organizationIds []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest {
	r.organizationIds = &organizationIds
	return r
}

// Subscription ids
func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest) SubscriptionIds(subscriptionIds []string) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest {
	r.subscriptionIds = &subscriptionIds
	return r
}

func (r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest) Execute() ([]GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner, *http.Response, error) {
	return r.ApiService.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesExecute(r)
}

/*
GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses Get compliance status for requested subscriptions

Get compliance status for requested subscriptions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest
*/
func (a *SubscriptionsAPIService) GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(ctx context.Context) SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest {
	return SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner
func (a *SubscriptionsAPIService) GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesExecute(r SubscriptionsAPIGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest) ([]GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions/compliance/statuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationIds == nil {
		return localVarReturnValue, nil, reportError("organizationIds is required and must be specified")
	}

	{
		t := *r.organizationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "organizationIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "organizationIds", t, "multi")
		}
	}
	if r.subscriptionIds != nil {
		t := *r.subscriptionIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionIds", t, "multi")
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

type SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct {
	ctx context.Context
	ApiService *SubscriptionsAPIService
	validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
}

func (r SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest = &validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
	return r
}

func (r SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) Execute() (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	return r.ApiService.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyExecute(r)
}

/*
ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey Find a subscription by claim key

Find a subscription by claim key. Returns 400 if the key has already been claimed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
*/
func (a *SubscriptionsAPIService) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx context.Context) SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	return SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
func (a *SubscriptionsAPIService) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyExecute(r SubscriptionsAPIValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions/claimKey/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest == nil {
		return localVarReturnValue, nil, reportError("validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest is required and must be specified")
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
	localVarPostBody = r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
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
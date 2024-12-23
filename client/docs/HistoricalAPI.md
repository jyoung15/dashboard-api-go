# \HistoricalAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationAssuranceAlertsOverviewHistorical**](HistoricalAPI.md#GetOrganizationAssuranceAlertsOverviewHistorical) | **Get** /organizations/{organizationId}/assurance/alerts/overview/historical | Returns historical health alert overviews



## GetOrganizationAssuranceAlertsOverviewHistorical

> GetOrganizationAssuranceAlertsOverviewHistorical200Response GetOrganizationAssuranceAlertsOverviewHistorical(ctx, organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).Execute()

Returns historical health alert overviews



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	segmentDuration := int32(56) // int32 | Amount of time in seconds for each segment in the returned dataset
	tsStart := time.Now() // time.Time | Parameter to define starting timestamp of historical totals
	networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
	severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
	types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
	tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp defaults to the current time (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
	deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HistoricalAPI.GetOrganizationAssuranceAlertsOverviewHistorical(context.Background(), organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HistoricalAPI.GetOrganizationAssuranceAlertsOverviewHistorical``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlertsOverviewHistorical`: GetOrganizationAssuranceAlertsOverviewHistorical200Response
	fmt.Fprintf(os.Stdout, "Response from `HistoricalAPI.GetOrganizationAssuranceAlertsOverviewHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentDuration** | **int32** | Amount of time in seconds for each segment in the returned dataset | 
 **tsStart** | **time.Time** | Parameter to define starting timestamp of historical totals | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp defaults to the current time | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 

### Return type

[**GetOrganizationAssuranceAlertsOverviewHistorical200Response**](GetOrganizationAssuranceAlertsOverviewHistorical200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


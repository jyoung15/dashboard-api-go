# \AlertsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSensorAlertsProfile**](AlertsAPI.md#CreateNetworkSensorAlertsProfile) | **Post** /networks/{networkId}/sensor/alerts/profiles | Creates a sensor alert profile for a network.
[**CreateOrganizationAlertsProfile**](AlertsAPI.md#CreateOrganizationAlertsProfile) | **Post** /organizations/{organizationId}/alerts/profiles | Create an organization-wide alert configuration
[**DeleteNetworkSensorAlertsProfile**](AlertsAPI.md#DeleteNetworkSensorAlertsProfile) | **Delete** /networks/{networkId}/sensor/alerts/profiles/{id} | Deletes a sensor alert profile from a network.
[**DeleteOrganizationAlertsProfile**](AlertsAPI.md#DeleteOrganizationAlertsProfile) | **Delete** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Removes an organization-wide alert config
[**DismissOrganizationAssuranceAlerts**](AlertsAPI.md#DismissOrganizationAssuranceAlerts) | **Post** /organizations/{organizationId}/assurance/alerts/dismiss | Dismiss health alerts
[**GetNetworkAlertsHistory**](AlertsAPI.md#GetNetworkAlertsHistory) | **Get** /networks/{networkId}/alerts/history | Return the alert history for this network
[**GetNetworkAlertsSettings**](AlertsAPI.md#GetNetworkAlertsSettings) | **Get** /networks/{networkId}/alerts/settings | Return the alert configuration for this network
[**GetNetworkHealthAlerts**](AlertsAPI.md#GetNetworkHealthAlerts) | **Get** /networks/{networkId}/health/alerts | Return all global alerts on this network
[**GetNetworkSensorAlertsCurrentOverviewByMetric**](AlertsAPI.md#GetNetworkSensorAlertsCurrentOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/current/overview/byMetric | Return an overview of currently alerting sensors by metric
[**GetNetworkSensorAlertsOverviewByMetric**](AlertsAPI.md#GetNetworkSensorAlertsOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/overview/byMetric | Return an overview of alert occurrences over a timespan, by metric
[**GetNetworkSensorAlertsProfile**](AlertsAPI.md#GetNetworkSensorAlertsProfile) | **Get** /networks/{networkId}/sensor/alerts/profiles/{id} | Show details of a sensor alert profile for a network.
[**GetNetworkSensorAlertsProfiles**](AlertsAPI.md#GetNetworkSensorAlertsProfiles) | **Get** /networks/{networkId}/sensor/alerts/profiles | Lists all sensor alert profiles for a network.
[**GetOrganizationAlertsProfiles**](AlertsAPI.md#GetOrganizationAlertsProfiles) | **Get** /organizations/{organizationId}/alerts/profiles | List all organization-wide alert configurations
[**GetOrganizationAssuranceAlert**](AlertsAPI.md#GetOrganizationAssuranceAlert) | **Get** /organizations/{organizationId}/assurance/alerts/{id} | Return a singular Health Alert by its id
[**GetOrganizationAssuranceAlerts**](AlertsAPI.md#GetOrganizationAssuranceAlerts) | **Get** /organizations/{organizationId}/assurance/alerts | Return all health alerts for an organization
[**GetOrganizationAssuranceAlertsOverview**](AlertsAPI.md#GetOrganizationAssuranceAlertsOverview) | **Get** /organizations/{organizationId}/assurance/alerts/overview | Return overview of active health alerts for an organization
[**GetOrganizationAssuranceAlertsOverviewByNetwork**](AlertsAPI.md#GetOrganizationAssuranceAlertsOverviewByNetwork) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byNetwork | Return a Summary of Alerts grouped by network and severity
[**GetOrganizationAssuranceAlertsOverviewByType**](AlertsAPI.md#GetOrganizationAssuranceAlertsOverviewByType) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byType | Return a Summary of Alerts grouped by type and severity
[**GetOrganizationAssuranceAlertsOverviewHistorical**](AlertsAPI.md#GetOrganizationAssuranceAlertsOverviewHistorical) | **Get** /organizations/{organizationId}/assurance/alerts/overview/historical | Returns historical health alert overviews
[**RestoreOrganizationAssuranceAlerts**](AlertsAPI.md#RestoreOrganizationAssuranceAlerts) | **Post** /organizations/{organizationId}/assurance/alerts/restore | Restore health alerts from dismissed
[**UpdateNetworkAlertsSettings**](AlertsAPI.md#UpdateNetworkAlertsSettings) | **Put** /networks/{networkId}/alerts/settings | Update the alert configuration for this network
[**UpdateNetworkSensorAlertsProfile**](AlertsAPI.md#UpdateNetworkSensorAlertsProfile) | **Put** /networks/{networkId}/sensor/alerts/profiles/{id} | Updates a sensor alert profile for a network.
[**UpdateOrganizationAlertsProfile**](AlertsAPI.md#UpdateOrganizationAlertsProfile) | **Put** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Update an organization-wide alert config



## CreateNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner CreateNetworkSensorAlertsProfile(ctx, networkId).CreateNetworkSensorAlertsProfileRequest(createNetworkSensorAlertsProfileRequest).Execute()

Creates a sensor alert profile for a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSensorAlertsProfileRequest := *openapiclient.NewCreateNetworkSensorAlertsProfileRequest("Name_example", []openapiclient.GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{*openapiclient.NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner("Metric_example", *openapiclient.NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold())}) // CreateNetworkSensorAlertsProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateNetworkSensorAlertsProfile(context.Background(), networkId).CreateNetworkSensorAlertsProfileRequest(createNetworkSensorAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSensorAlertsProfileRequest** | [**CreateNetworkSensorAlertsProfileRequest**](CreateNetworkSensorAlertsProfileRequest.md) |  | 

### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAlertsProfile

> GetOrganizationAlertsProfiles200ResponseInner CreateOrganizationAlertsProfile(ctx, organizationId).CreateOrganizationAlertsProfileRequest(createOrganizationAlertsProfileRequest).Execute()

Create an organization-wide alert configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	createOrganizationAlertsProfileRequest := *openapiclient.NewCreateOrganizationAlertsProfileRequest("Type_example", *openapiclient.NewCreateOrganizationAlertsProfileRequestAlertCondition(), *openapiclient.NewGetOrganizationAlertsProfiles200ResponseInnerRecipients(), []string{"NetworkTags_example"}) // CreateOrganizationAlertsProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateOrganizationAlertsProfile(context.Background(), organizationId).CreateOrganizationAlertsProfileRequest(createOrganizationAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateOrganizationAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAlertsProfile`: GetOrganizationAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAlertsProfileRequest** | [**CreateOrganizationAlertsProfileRequest**](CreateOrganizationAlertsProfileRequest.md) |  | 

### Return type

[**GetOrganizationAlertsProfiles200ResponseInner**](GetOrganizationAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSensorAlertsProfile

> DeleteNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Deletes a sensor alert profile from a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DeleteNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAlertsProfile

> DeleteOrganizationAlertsProfile(ctx, organizationId, alertConfigId).Execute()

Removes an organization-wide alert config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	alertConfigId := "alertConfigId_example" // string | Alert config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DeleteOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteOrganizationAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DismissOrganizationAssuranceAlerts

> DismissOrganizationAssuranceAlerts(ctx, organizationId).DismissOrganizationAssuranceAlertsRequest(dismissOrganizationAssuranceAlertsRequest).Execute()

Dismiss health alerts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	dismissOrganizationAssuranceAlertsRequest := *openapiclient.NewDismissOrganizationAssuranceAlertsRequest([]string{"AlertIds_example"}) // DismissOrganizationAssuranceAlertsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.DismissOrganizationAssuranceAlerts(context.Background(), organizationId).DismissOrganizationAssuranceAlertsRequest(dismissOrganizationAssuranceAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DismissOrganizationAssuranceAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDismissOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dismissOrganizationAssuranceAlertsRequest** | [**DismissOrganizationAssuranceAlertsRequest**](DismissOrganizationAssuranceAlertsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsHistory

> []GetNetworkAlertsHistory200ResponseInner GetNetworkAlertsHistory(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the alert history for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkAlertsHistory(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkAlertsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlertsHistory`: []GetNetworkAlertsHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkAlertsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkAlertsHistory200ResponseInner**](GetNetworkAlertsHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsSettings

> GetNetworkAlertsSettings200Response GetNetworkAlertsSettings(ctx, networkId).Execute()

Return the alert configuration for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkAlertsSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkAlertsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlertsSettings`: GetNetworkAlertsSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkAlertsSettings200Response**](GetNetworkAlertsSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkHealthAlerts

> []GetNetworkHealthAlerts200ResponseInner GetNetworkHealthAlerts(ctx, networkId).Execute()

Return all global alerts on this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkHealthAlerts(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkHealthAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkHealthAlerts`: []GetNetworkHealthAlerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkHealthAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkHealthAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkHealthAlerts200ResponseInner**](GetNetworkHealthAlerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsCurrentOverviewByMetric

> GetNetworkSensorAlertsCurrentOverviewByMetric200Response GetNetworkSensorAlertsCurrentOverviewByMetric(ctx, networkId).Execute()

Return an overview of currently alerting sensors by metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkSensorAlertsCurrentOverviewByMetric(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkSensorAlertsCurrentOverviewByMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsCurrentOverviewByMetric`: GetNetworkSensorAlertsCurrentOverviewByMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkSensorAlertsCurrentOverviewByMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsCurrentOverviewByMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSensorAlertsCurrentOverviewByMetric200Response**](GetNetworkSensorAlertsCurrentOverviewByMetric200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsOverviewByMetric

> []GetNetworkSensorAlertsOverviewByMetric200ResponseInner GetNetworkSensorAlertsOverviewByMetric(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Return an overview of alert occurrences over a timespan, by metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 86400, 604800. The default is 604800. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkSensorAlertsOverviewByMetric(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkSensorAlertsOverviewByMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsOverviewByMetric`: []GetNetworkSensorAlertsOverviewByMetric200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkSensorAlertsOverviewByMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsOverviewByMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 86400, 604800. The default is 604800. | 

### Return type

[**[]GetNetworkSensorAlertsOverviewByMetric200ResponseInner**](GetNetworkSensorAlertsOverviewByMetric200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner GetNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Show details of a sensor alert profile for a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsProfiles

> []GetNetworkSensorAlertsProfiles200ResponseInner GetNetworkSensorAlertsProfiles(ctx, networkId).Execute()

Lists all sensor alert profiles for a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetNetworkSensorAlertsProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetNetworkSensorAlertsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsProfiles`: []GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetNetworkSensorAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAlertsProfiles

> []GetOrganizationAlertsProfiles200ResponseInner GetOrganizationAlertsProfiles(ctx, organizationId).Execute()

List all organization-wide alert configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAlertsProfiles(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAlertsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAlertsProfiles`: []GetOrganizationAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationAlertsProfiles200ResponseInner**](GetOrganizationAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlert

> GetOrganizationAssuranceAlerts200ResponseInner GetOrganizationAssuranceAlert(ctx, organizationId, id).Execute()

Return a singular Health Alert by its id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlert(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlert`: GetOrganizationAssuranceAlerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationAssuranceAlerts200ResponseInner**](GetOrganizationAssuranceAlerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlerts

> []GetOrganizationAssuranceAlerts200ResponseInner GetOrganizationAssuranceAlerts(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return all health alerts for an organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 4 - 300. Default is 30. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
	networkId := "networkId_example" // string | Optional parameter to filter alerts by network ids. (optional)
	severity := "severity_example" // string | Optional parameter to filter by severity type. (optional)
	types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
	tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
	tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
	category := "category_example" // string | Optional parameter to filter by category. (optional)
	sortBy := "sortBy_example" // string | Optional parameter to set column to sort by. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
	deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
	deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
	active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
	dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
	resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
	suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlerts(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlerts`: []GetOrganizationAssuranceAlerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 4 - 300. Default is 30. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts by network ids. | 
 **severity** | **string** | Optional parameter to filter by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **category** | **string** | Optional parameter to filter by category. | 
 **sortBy** | **string** | Optional parameter to set column to sort by. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**[]GetOrganizationAssuranceAlerts200ResponseInner**](GetOrganizationAssuranceAlerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverview

> GetOrganizationAssuranceAlertsOverview200Response GetOrganizationAssuranceAlertsOverview(ctx, organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return overview of active health alerts for an organization



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
	networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
	severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
	types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
	tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
	tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
	deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
	deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
	active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
	dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
	resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
	suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlertsOverview(context.Background(), organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlertsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlertsOverview`: GetOrganizationAssuranceAlertsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlertsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**GetOrganizationAssuranceAlertsOverview200Response**](GetOrganizationAssuranceAlertsOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByNetwork

> GetOrganizationAssuranceAlertsOverviewByNetwork200Response GetOrganizationAssuranceAlertsOverviewByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return a Summary of Alerts grouped by network and severity



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
	networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network id. (optional)
	severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
	types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
	tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
	tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
	deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
	deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
	active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
	dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
	resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
	suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlertsOverviewByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlertsOverviewByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlertsOverviewByNetwork`: GetOrganizationAssuranceAlertsOverviewByNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlertsOverviewByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network id. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**GetOrganizationAssuranceAlertsOverviewByNetwork200Response**](GetOrganizationAssuranceAlertsOverviewByNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByType

> GetOrganizationAssuranceAlertsOverviewByType200Response GetOrganizationAssuranceAlertsOverviewByType(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return a Summary of Alerts grouped by type and severity



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
	networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
	severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
	types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
	tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
	tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
	sortBy := "sortBy_example" // string | Optional parameter to set column to sort by. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
	deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
	deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
	active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
	dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
	resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
	suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlertsOverviewByType(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlertsOverviewByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlertsOverviewByType`: GetOrganizationAssuranceAlertsOverviewByType200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlertsOverviewByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **sortBy** | **string** | Optional parameter to set column to sort by. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**GetOrganizationAssuranceAlertsOverviewByType200Response**](GetOrganizationAssuranceAlertsOverviewByType200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.AlertsAPI.GetOrganizationAssuranceAlertsOverviewHistorical(context.Background(), organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Serials(serials).DeviceTypes(deviceTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetOrganizationAssuranceAlertsOverviewHistorical``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAssuranceAlertsOverviewHistorical`: GetOrganizationAssuranceAlertsOverviewHistorical200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetOrganizationAssuranceAlertsOverviewHistorical`: %v\n", resp)
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


## RestoreOrganizationAssuranceAlerts

> RestoreOrganizationAssuranceAlerts(ctx, organizationId).RestoreOrganizationAssuranceAlertsRequest(restoreOrganizationAssuranceAlertsRequest).Execute()

Restore health alerts from dismissed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	restoreOrganizationAssuranceAlertsRequest := *openapiclient.NewRestoreOrganizationAssuranceAlertsRequest([]string{"AlertIds_example"}) // RestoreOrganizationAssuranceAlertsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertsAPI.RestoreOrganizationAssuranceAlerts(context.Background(), organizationId).RestoreOrganizationAssuranceAlertsRequest(restoreOrganizationAssuranceAlertsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.RestoreOrganizationAssuranceAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreOrganizationAssuranceAlertsRequest** | [**RestoreOrganizationAssuranceAlertsRequest**](RestoreOrganizationAssuranceAlertsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAlertsSettings

> GetNetworkAlertsSettings200Response UpdateNetworkAlertsSettings(ctx, networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()

Update the alert configuration for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkAlertsSettingsRequest := *openapiclient.NewUpdateNetworkAlertsSettingsRequest() // UpdateNetworkAlertsSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateNetworkAlertsSettings(context.Background(), networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateNetworkAlertsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAlertsSettings`: GetNetworkAlertsSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkAlertsSettingsRequest** | [**UpdateNetworkAlertsSettingsRequest**](UpdateNetworkAlertsSettingsRequest.md) |  | 

### Return type

[**GetNetworkAlertsSettings200Response**](GetNetworkAlertsSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner UpdateNetworkSensorAlertsProfile(ctx, networkId, id).UpdateNetworkSensorAlertsProfileRequest(updateNetworkSensorAlertsProfileRequest).Execute()

Updates a sensor alert profile for a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	id := "id_example" // string | ID
	updateNetworkSensorAlertsProfileRequest := *openapiclient.NewUpdateNetworkSensorAlertsProfileRequest() // UpdateNetworkSensorAlertsProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateNetworkSensorAlertsProfile(context.Background(), networkId, id).UpdateNetworkSensorAlertsProfileRequest(updateNetworkSensorAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSensorAlertsProfileRequest** | [**UpdateNetworkSensorAlertsProfileRequest**](UpdateNetworkSensorAlertsProfileRequest.md) |  | 

### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAlertsProfile

> GetOrganizationAlertsProfiles200ResponseInner UpdateOrganizationAlertsProfile(ctx, organizationId, alertConfigId).UpdateOrganizationAlertsProfileRequest(updateOrganizationAlertsProfileRequest).Execute()

Update an organization-wide alert config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	alertConfigId := "alertConfigId_example" // string | Alert config ID
	updateOrganizationAlertsProfileRequest := *openapiclient.NewUpdateOrganizationAlertsProfileRequest() // UpdateOrganizationAlertsProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).UpdateOrganizationAlertsProfileRequest(updateOrganizationAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateOrganizationAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAlertsProfile`: GetOrganizationAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAlertsProfileRequest** | [**UpdateOrganizationAlertsProfileRequest**](UpdateOrganizationAlertsProfileRequest.md) |  | 

### Return type

[**GetOrganizationAlertsProfiles200ResponseInner**](GetOrganizationAlertsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

# \SensorAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSensorCommand**](SensorAPI.md#CreateDeviceSensorCommand) | **Post** /devices/{serial}/sensor/commands | Sends a command to a sensor
[**CreateNetworkSensorAlertsProfile**](SensorAPI.md#CreateNetworkSensorAlertsProfile) | **Post** /networks/{networkId}/sensor/alerts/profiles | Creates a sensor alert profile for a network.
[**DeleteNetworkSensorAlertsProfile**](SensorAPI.md#DeleteNetworkSensorAlertsProfile) | **Delete** /networks/{networkId}/sensor/alerts/profiles/{id} | Deletes a sensor alert profile from a network.
[**GetDeviceSensorCommand**](SensorAPI.md#GetDeviceSensorCommand) | **Get** /devices/{serial}/sensor/commands/{commandId} | Returns information about the command&#39;s execution, including the status
[**GetDeviceSensorCommands**](SensorAPI.md#GetDeviceSensorCommands) | **Get** /devices/{serial}/sensor/commands | Returns a historical log of all commands
[**GetDeviceSensorRelationships**](SensorAPI.md#GetDeviceSensorRelationships) | **Get** /devices/{serial}/sensor/relationships | List the sensor roles for a given sensor or camera device.
[**GetNetworkSensorAlertsCurrentOverviewByMetric**](SensorAPI.md#GetNetworkSensorAlertsCurrentOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/current/overview/byMetric | Return an overview of currently alerting sensors by metric
[**GetNetworkSensorAlertsOverviewByMetric**](SensorAPI.md#GetNetworkSensorAlertsOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/overview/byMetric | Return an overview of alert occurrences over a timespan, by metric
[**GetNetworkSensorAlertsProfile**](SensorAPI.md#GetNetworkSensorAlertsProfile) | **Get** /networks/{networkId}/sensor/alerts/profiles/{id} | Show details of a sensor alert profile for a network.
[**GetNetworkSensorAlertsProfiles**](SensorAPI.md#GetNetworkSensorAlertsProfiles) | **Get** /networks/{networkId}/sensor/alerts/profiles | Lists all sensor alert profiles for a network.
[**GetNetworkSensorMqttBroker**](SensorAPI.md#GetNetworkSensorMqttBroker) | **Get** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Return the sensor settings of an MQTT broker
[**GetNetworkSensorMqttBrokers**](SensorAPI.md#GetNetworkSensorMqttBrokers) | **Get** /networks/{networkId}/sensor/mqttBrokers | List the sensor settings of all MQTT brokers for this network
[**GetNetworkSensorRelationships**](SensorAPI.md#GetNetworkSensorRelationships) | **Get** /networks/{networkId}/sensor/relationships | List the sensor roles for devices in a given network
[**GetOrganizationSensorReadingsHistory**](SensorAPI.md#GetOrganizationSensorReadingsHistory) | **Get** /organizations/{organizationId}/sensor/readings/history | Return all reported readings from sensors in a given timespan, sorted by timestamp
[**GetOrganizationSensorReadingsLatest**](SensorAPI.md#GetOrganizationSensorReadingsLatest) | **Get** /organizations/{organizationId}/sensor/readings/latest | Return the latest available reading for each metric from each sensor, sorted by sensor serial
[**UpdateDeviceSensorRelationships**](SensorAPI.md#UpdateDeviceSensorRelationships) | **Put** /devices/{serial}/sensor/relationships | Assign one or more sensor roles to a given sensor or camera device.
[**UpdateNetworkSensorAlertsProfile**](SensorAPI.md#UpdateNetworkSensorAlertsProfile) | **Put** /networks/{networkId}/sensor/alerts/profiles/{id} | Updates a sensor alert profile for a network.
[**UpdateNetworkSensorMqttBroker**](SensorAPI.md#UpdateNetworkSensorMqttBroker) | **Put** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Update the sensor settings of an MQTT broker



## CreateDeviceSensorCommand

> GetDeviceSensorCommands200ResponseInner CreateDeviceSensorCommand(ctx, serial).CreateDeviceSensorCommandRequest(createDeviceSensorCommandRequest).Execute()

Sends a command to a sensor



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
	serial := "serial_example" // string | Serial
	createDeviceSensorCommandRequest := *openapiclient.NewCreateDeviceSensorCommandRequest("Operation_example") // CreateDeviceSensorCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.CreateDeviceSensorCommand(context.Background(), serial).CreateDeviceSensorCommandRequest(createDeviceSensorCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.CreateDeviceSensorCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSensorCommand`: GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.CreateDeviceSensorCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSensorCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSensorCommandRequest** | [**CreateDeviceSensorCommandRequest**](CreateDeviceSensorCommandRequest.md) |  | 

### Return type

[**GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.SensorAPI.CreateNetworkSensorAlertsProfile(context.Background(), networkId).CreateNetworkSensorAlertsProfileRequest(createNetworkSensorAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.CreateNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.CreateNetworkSensorAlertsProfile`: %v\n", resp)
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
	r, err := apiClient.SensorAPI.DeleteNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.DeleteNetworkSensorAlertsProfile``: %v\n", err)
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


## GetDeviceSensorCommand

> GetDeviceSensorCommands200ResponseInner GetDeviceSensorCommand(ctx, serial, commandId).Execute()

Returns information about the command's execution, including the status



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
	serial := "serial_example" // string | Serial
	commandId := "commandId_example" // string | Command ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetDeviceSensorCommand(context.Background(), serial, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetDeviceSensorCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorCommand`: GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetDeviceSensorCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**commandId** | **string** | Command ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSensorCommands

> []GetDeviceSensorCommands200ResponseInner GetDeviceSensorCommands(ctx, serial).Operations(operations).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns a historical log of all commands



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
	serial := "serial_example" // string | Serial
	operations := []string{"Inner_example"} // []string | Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetDeviceSensorCommands(context.Background(), serial).Operations(operations).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetDeviceSensorCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorCommands`: []GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetDeviceSensorCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operations** | **[]string** | Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;descending&#39;. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days. | 

### Return type

[**[]GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSensorRelationships

> GetDeviceSensorRelationships200Response GetDeviceSensorRelationships(ctx, serial).Execute()

List the sensor roles for a given sensor or camera device.



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetDeviceSensorRelationships(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetDeviceSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorRelationships`: GetDeviceSensorRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceSensorRelationships200Response**](GetDeviceSensorRelationships200Response.md)

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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorAlertsCurrentOverviewByMetric(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorAlertsCurrentOverviewByMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsCurrentOverviewByMetric`: GetNetworkSensorAlertsCurrentOverviewByMetric200Response
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorAlertsCurrentOverviewByMetric`: %v\n", resp)
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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorAlertsOverviewByMetric(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorAlertsOverviewByMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsOverviewByMetric`: []GetNetworkSensorAlertsOverviewByMetric200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorAlertsOverviewByMetric`: %v\n", resp)
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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorAlertsProfile`: %v\n", resp)
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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorAlertsProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorAlertsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorAlertsProfiles`: []GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorAlertsProfiles`: %v\n", resp)
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


## GetNetworkSensorMqttBroker

> GetNetworkSensorMqttBrokers200ResponseInner GetNetworkSensorMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return the sensor settings of an MQTT broker



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
	mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorMqttBrokers

> []GetNetworkSensorMqttBrokers200ResponseInner GetNetworkSensorMqttBrokers(ctx, networkId).Execute()

List the sensor settings of all MQTT brokers for this network



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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorMqttBrokers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorMqttBrokers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorMqttBrokers`: []GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorMqttBrokersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorRelationships

> []GetNetworkSensorRelationships200ResponseInner GetNetworkSensorRelationships(ctx, networkId).Execute()

List the sensor roles for devices in a given network



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
	resp, r, err := apiClient.SensorAPI.GetNetworkSensorRelationships(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetNetworkSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorRelationships`: []GetNetworkSensorRelationships200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetNetworkSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorRelationships200ResponseInner**](GetNetworkSensorRelationships200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSensorReadingsHistory

> []GetOrganizationSensorReadingsHistory200ResponseInner GetOrganizationSensorReadingsHistory(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()

Return all reported readings from sensors in a given timespan, sorted by timestamp



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter readings by network. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter readings by sensor. (optional)
	metrics := []string{"Inner_example"} // []string | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetOrganizationSensorReadingsHistory(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetOrganizationSensorReadingsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSensorReadingsHistory`: []GetOrganizationSensorReadingsHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetOrganizationSensorReadingsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSensorReadingsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. | 
 **networkIds** | **[]string** | Optional parameter to filter readings by network. | 
 **serials** | **[]string** | Optional parameter to filter readings by sensor. | 
 **metrics** | **[]string** | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water. | 

### Return type

[**[]GetOrganizationSensorReadingsHistory200ResponseInner**](GetOrganizationSensorReadingsHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSensorReadingsLatest

> []GetOrganizationSensorReadingsLatest200ResponseInner GetOrganizationSensorReadingsLatest(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()

Return the latest available reading for each metric from each sensor, sorted by sensor serial



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter readings by network. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter readings by sensor. (optional)
	metrics := []string{"Inner_example"} // []string | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.GetOrganizationSensorReadingsLatest(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.GetOrganizationSensorReadingsLatest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSensorReadingsLatest`: []GetOrganizationSensorReadingsLatest200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.GetOrganizationSensorReadingsLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSensorReadingsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter readings by network. | 
 **serials** | **[]string** | Optional parameter to filter readings by sensor. | 
 **metrics** | **[]string** | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water. | 

### Return type

[**[]GetOrganizationSensorReadingsLatest200ResponseInner**](GetOrganizationSensorReadingsLatest200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSensorRelationships

> GetDeviceSensorRelationships200Response UpdateDeviceSensorRelationships(ctx, serial).UpdateDeviceSensorRelationshipsRequest(updateDeviceSensorRelationshipsRequest).Execute()

Assign one or more sensor roles to a given sensor or camera device.



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
	serial := "serial_example" // string | Serial
	updateDeviceSensorRelationshipsRequest := *openapiclient.NewUpdateDeviceSensorRelationshipsRequest() // UpdateDeviceSensorRelationshipsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.UpdateDeviceSensorRelationships(context.Background(), serial).UpdateDeviceSensorRelationshipsRequest(updateDeviceSensorRelationshipsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.UpdateDeviceSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSensorRelationships`: GetDeviceSensorRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.UpdateDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSensorRelationshipsRequest** | [**UpdateDeviceSensorRelationshipsRequest**](UpdateDeviceSensorRelationshipsRequest.md) |  | 

### Return type

[**GetDeviceSensorRelationships200Response**](GetDeviceSensorRelationships200Response.md)

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
	resp, r, err := apiClient.SensorAPI.UpdateNetworkSensorAlertsProfile(context.Background(), networkId, id).UpdateNetworkSensorAlertsProfileRequest(updateNetworkSensorAlertsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.UpdateNetworkSensorAlertsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.UpdateNetworkSensorAlertsProfile`: %v\n", resp)
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


## UpdateNetworkSensorMqttBroker

> GetNetworkSensorMqttBrokers200ResponseInner UpdateNetworkSensorMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkSensorMqttBrokerRequest(updateNetworkSensorMqttBrokerRequest).Execute()

Update the sensor settings of an MQTT broker



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
	mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID
	updateNetworkSensorMqttBrokerRequest := *openapiclient.NewUpdateNetworkSensorMqttBrokerRequest(false) // UpdateNetworkSensorMqttBrokerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SensorAPI.UpdateNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkSensorMqttBrokerRequest(updateNetworkSensorMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensorAPI.UpdateNetworkSensorMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SensorAPI.UpdateNetworkSensorMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSensorMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSensorMqttBrokerRequest** | [**UpdateNetworkSensorMqttBrokerRequest**](UpdateNetworkSensorMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


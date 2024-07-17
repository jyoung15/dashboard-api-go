# \MqttBrokersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMqttBroker**](MqttBrokersAPI.md#CreateNetworkMqttBroker) | **Post** /networks/{networkId}/mqttBrokers | Add an MQTT broker
[**DeleteNetworkMqttBroker**](MqttBrokersAPI.md#DeleteNetworkMqttBroker) | **Delete** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Delete an MQTT broker
[**GetNetworkMqttBroker**](MqttBrokersAPI.md#GetNetworkMqttBroker) | **Get** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Return an MQTT broker
[**GetNetworkMqttBrokers**](MqttBrokersAPI.md#GetNetworkMqttBrokers) | **Get** /networks/{networkId}/mqttBrokers | List the MQTT brokers for this network
[**GetNetworkSensorMqttBroker**](MqttBrokersAPI.md#GetNetworkSensorMqttBroker) | **Get** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Return the sensor settings of an MQTT broker
[**GetNetworkSensorMqttBrokers**](MqttBrokersAPI.md#GetNetworkSensorMqttBrokers) | **Get** /networks/{networkId}/sensor/mqttBrokers | List the sensor settings of all MQTT brokers for this network
[**UpdateNetworkMqttBroker**](MqttBrokersAPI.md#UpdateNetworkMqttBroker) | **Put** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Update an MQTT broker
[**UpdateNetworkSensorMqttBroker**](MqttBrokersAPI.md#UpdateNetworkSensorMqttBroker) | **Put** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Update the sensor settings of an MQTT broker



## CreateNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner CreateNetworkMqttBroker(ctx, networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()

Add an MQTT broker



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
	createNetworkMqttBrokerRequest := *openapiclient.NewCreateNetworkMqttBrokerRequest("Name_example", "Host_example", int32(123)) // CreateNetworkMqttBrokerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqttBrokersAPI.CreateNetworkMqttBroker(context.Background(), networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.CreateNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.CreateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMqttBrokerRequest** | [**CreateNetworkMqttBrokerRequest**](CreateNetworkMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkMqttBroker

> DeleteNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Delete an MQTT broker



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
	r, err := apiClient.MqttBrokersAPI.DeleteNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.DeleteNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMqttBrokerRequest struct via the builder pattern


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


## GetNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner GetNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return an MQTT broker



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
	resp, r, err := apiClient.MqttBrokersAPI.GetNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.GetNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.GetNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMqttBrokers

> []GetNetworkMqttBrokers200ResponseInner GetNetworkMqttBrokers(ctx, networkId).Execute()

List the MQTT brokers for this network



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
	resp, r, err := apiClient.MqttBrokersAPI.GetNetworkMqttBrokers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.GetNetworkMqttBrokers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMqttBrokers`: []GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.GetNetworkMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

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
	resp, r, err := apiClient.MqttBrokersAPI.GetNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.GetNetworkSensorMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.GetNetworkSensorMqttBroker`: %v\n", resp)
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
	resp, r, err := apiClient.MqttBrokersAPI.GetNetworkSensorMqttBrokers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.GetNetworkSensorMqttBrokers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorMqttBrokers`: []GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.GetNetworkSensorMqttBrokers`: %v\n", resp)
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


## UpdateNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner UpdateNetworkMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()

Update an MQTT broker



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
	updateNetworkMqttBrokerRequest := *openapiclient.NewUpdateNetworkMqttBrokerRequest() // UpdateNetworkMqttBrokerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqttBrokersAPI.UpdateNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.UpdateNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.UpdateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMqttBrokerRequest** | [**UpdateNetworkMqttBrokerRequest**](UpdateNetworkMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

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
	resp, r, err := apiClient.MqttBrokersAPI.UpdateNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkSensorMqttBrokerRequest(updateNetworkSensorMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersAPI.UpdateNetworkSensorMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MqttBrokersAPI.UpdateNetworkSensorMqttBroker`: %v\n", resp)
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


# \MulticastAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchRoutingMulticastRendezvousPoint**](MulticastAPI.md#CreateNetworkSwitchRoutingMulticastRendezvousPoint) | **Post** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | Create a multicast rendezvous point
[**DeleteNetworkSwitchRoutingMulticastRendezvousPoint**](MulticastAPI.md#DeleteNetworkSwitchRoutingMulticastRendezvousPoint) | **Delete** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Delete a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticast**](MulticastAPI.md#GetNetworkSwitchRoutingMulticast) | **Get** /networks/{networkId}/switch/routing/multicast | Return multicast settings for a network
[**GetNetworkSwitchRoutingMulticastRendezvousPoint**](MulticastAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoint) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Return a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoints**](MulticastAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoints) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | List multicast rendezvous points
[**UpdateNetworkSwitchRoutingMulticast**](MulticastAPI.md#UpdateNetworkSwitchRoutingMulticast) | **Put** /networks/{networkId}/switch/routing/multicast | Update multicast settings for a network
[**UpdateNetworkSwitchRoutingMulticastRendezvousPoint**](MulticastAPI.md#UpdateNetworkSwitchRoutingMulticastRendezvousPoint) | **Put** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Update a multicast rendezvous point



## CreateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Create a multicast rendezvous point



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
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MulticastAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchRoutingMulticastRendezvousPoint

> DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Delete a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MulticastAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


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


## GetNetworkSwitchRoutingMulticast

> GetNetworkSwitchRoutingMulticast200Response GetNetworkSwitchRoutingMulticast(ctx, networkId).Execute()

Return multicast settings for a network



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
	resp, r, err := apiClient.MulticastAPI.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.GetNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.GetNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchRoutingMulticast200Response**](GetNetworkSwitchRoutingMulticast200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Return a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoints

> []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx, networkId).Execute()

List multicast rendezvous points



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
	resp, r, err := apiClient.MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoints`: []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticast

> GetNetworkSwitchRoutingMulticast200Response UpdateNetworkSwitchRoutingMulticast(ctx, networkId).UpdateNetworkSwitchRoutingMulticastRequest(updateNetworkSwitchRoutingMulticastRequest).Execute()

Update multicast settings for a network



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
	updateNetworkSwitchRoutingMulticastRequest := *openapiclient.NewUpdateNetworkSwitchRoutingMulticastRequest() // UpdateNetworkSwitchRoutingMulticastRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MulticastAPI.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).UpdateNetworkSwitchRoutingMulticastRequest(updateNetworkSwitchRoutingMulticastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.UpdateNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.UpdateNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingMulticastRequest** | [**UpdateNetworkSwitchRoutingMulticastRequest**](UpdateNetworkSwitchRoutingMulticastRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticast200Response**](GetNetworkSwitchRoutingMulticast200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Update a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MulticastAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MulticastAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MulticastAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


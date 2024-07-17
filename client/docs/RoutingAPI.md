# \RoutingAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSwitchRoutingInterface**](RoutingAPI.md#CreateDeviceSwitchRoutingInterface) | **Post** /devices/{serial}/switch/routing/interfaces | Create a layer 3 interface for a switch
[**CreateDeviceSwitchRoutingStaticRoute**](RoutingAPI.md#CreateDeviceSwitchRoutingStaticRoute) | **Post** /devices/{serial}/switch/routing/staticRoutes | Create a layer 3 static route for a switch
[**CreateNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingAPI.md#CreateNetworkSwitchRoutingMulticastRendezvousPoint) | **Post** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | Create a multicast rendezvous point
[**CreateNetworkSwitchStackRoutingInterface**](RoutingAPI.md#CreateNetworkSwitchStackRoutingInterface) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | Create a layer 3 interface for a switch stack
[**CreateNetworkSwitchStackRoutingStaticRoute**](RoutingAPI.md#CreateNetworkSwitchStackRoutingStaticRoute) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | Create a layer 3 static route for a switch stack
[**DeleteDeviceSwitchRoutingInterface**](RoutingAPI.md#DeleteDeviceSwitchRoutingInterface) | **Delete** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Delete a layer 3 interface from the switch
[**DeleteDeviceSwitchRoutingStaticRoute**](RoutingAPI.md#DeleteDeviceSwitchRoutingStaticRoute) | **Delete** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch
[**DeleteNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingAPI.md#DeleteNetworkSwitchRoutingMulticastRendezvousPoint) | **Delete** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Delete a multicast rendezvous point
[**DeleteNetworkSwitchStackRoutingInterface**](RoutingAPI.md#DeleteNetworkSwitchStackRoutingInterface) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Delete a layer 3 interface from a switch stack
[**DeleteNetworkSwitchStackRoutingStaticRoute**](RoutingAPI.md#DeleteNetworkSwitchStackRoutingStaticRoute) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch stack
[**GetDeviceSwitchRoutingInterface**](RoutingAPI.md#GetDeviceSwitchRoutingInterface) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Return a layer 3 interface for a switch
[**GetDeviceSwitchRoutingInterfaceDhcp**](RoutingAPI.md#GetDeviceSwitchRoutingInterfaceDhcp) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch
[**GetDeviceSwitchRoutingInterfaces**](RoutingAPI.md#GetDeviceSwitchRoutingInterfaces) | **Get** /devices/{serial}/switch/routing/interfaces | List layer 3 interfaces for a switch
[**GetDeviceSwitchRoutingStaticRoute**](RoutingAPI.md#GetDeviceSwitchRoutingStaticRoute) | **Get** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch
[**GetDeviceSwitchRoutingStaticRoutes**](RoutingAPI.md#GetDeviceSwitchRoutingStaticRoutes) | **Get** /devices/{serial}/switch/routing/staticRoutes | List layer 3 static routes for a switch
[**GetNetworkSwitchRoutingMulticast**](RoutingAPI.md#GetNetworkSwitchRoutingMulticast) | **Get** /networks/{networkId}/switch/routing/multicast | Return multicast settings for a network
[**GetNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoint) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Return a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoints**](RoutingAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoints) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | List multicast rendezvous points
[**GetNetworkSwitchRoutingOspf**](RoutingAPI.md#GetNetworkSwitchRoutingOspf) | **Get** /networks/{networkId}/switch/routing/ospf | Return layer 3 OSPF routing configuration
[**GetNetworkSwitchStackRoutingInterface**](RoutingAPI.md#GetNetworkSwitchStackRoutingInterface) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Return a layer 3 interface from a switch stack
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](RoutingAPI.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**GetNetworkSwitchStackRoutingInterfaces**](RoutingAPI.md#GetNetworkSwitchStackRoutingInterfaces) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | List layer 3 interfaces for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoute**](RoutingAPI.md#GetNetworkSwitchStackRoutingStaticRoute) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoutes**](RoutingAPI.md#GetNetworkSwitchStackRoutingStaticRoutes) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | List layer 3 static routes for a switch stack
[**UpdateDeviceSwitchRoutingInterface**](RoutingAPI.md#UpdateDeviceSwitchRoutingInterface) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch
[**UpdateDeviceSwitchRoutingInterfaceDhcp**](RoutingAPI.md#UpdateDeviceSwitchRoutingInterfaceDhcp) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch
[**UpdateDeviceSwitchRoutingStaticRoute**](RoutingAPI.md#UpdateDeviceSwitchRoutingStaticRoute) | **Put** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch
[**UpdateNetworkSwitchRoutingMulticast**](RoutingAPI.md#UpdateNetworkSwitchRoutingMulticast) | **Put** /networks/{networkId}/switch/routing/multicast | Update multicast settings for a network
[**UpdateNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingAPI.md#UpdateNetworkSwitchRoutingMulticastRendezvousPoint) | **Put** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Update a multicast rendezvous point
[**UpdateNetworkSwitchRoutingOspf**](RoutingAPI.md#UpdateNetworkSwitchRoutingOspf) | **Put** /networks/{networkId}/switch/routing/ospf | Update layer 3 OSPF routing configuration
[**UpdateNetworkSwitchStackRoutingInterface**](RoutingAPI.md#UpdateNetworkSwitchStackRoutingInterface) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch stack
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](RoutingAPI.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack
[**UpdateNetworkSwitchStackRoutingStaticRoute**](RoutingAPI.md#UpdateNetworkSwitchStackRoutingStaticRoute) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch stack



## CreateDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner CreateDeviceSwitchRoutingInterface(ctx, serial).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()

Create a layer 3 interface for a switch



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
	createDeviceSwitchRoutingInterfaceRequest := *openapiclient.NewCreateDeviceSwitchRoutingInterfaceRequest() // CreateDeviceSwitchRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreateDeviceSwitchRoutingInterface(context.Background(), serial).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingInterfaceRequest** | [**CreateDeviceSwitchRoutingInterfaceRequest**](CreateDeviceSwitchRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner CreateDeviceSwitchRoutingStaticRoute(ctx, serial).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()

Create a layer 3 static route for a switch



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
	createDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewCreateDeviceSwitchRoutingStaticRouteRequest("Subnet_example", "NextHopIp_example") // CreateDeviceSwitchRoutingStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingStaticRouteRequest** | [**CreateDeviceSwitchRoutingStaticRouteRequest**](CreateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.RoutingAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
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


## CreateNetworkSwitchStackRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner CreateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingInterfaceRequest(createNetworkSwitchStackRoutingInterfaceRequest).Execute()

Create a layer 3 interface for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	createNetworkSwitchStackRoutingInterfaceRequest := *openapiclient.NewCreateNetworkSwitchStackRoutingInterfaceRequest("Name_example", int32(123)) // CreateNetworkSwitchStackRoutingInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingInterfaceRequest(createNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingInterfaceRequest** | [**CreateNetworkSwitchStackRoutingInterfaceRequest**](CreateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner CreateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()

Create a layer 3 static route for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	createDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewCreateDeviceSwitchRoutingStaticRouteRequest("Subnet_example", "NextHopIp_example") // CreateDeviceSwitchRoutingStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.CreateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.CreateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDeviceSwitchRoutingStaticRouteRequest** | [**CreateDeviceSwitchRoutingStaticRouteRequest**](CreateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceSwitchRoutingInterface

> DeleteDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Delete a layer 3 interface from the switch



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
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


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


## DeleteDeviceSwitchRoutingStaticRoute

> DeleteDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Delete a layer 3 static route for a switch



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
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


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
	r, err := apiClient.RoutingAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
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


## DeleteNetworkSwitchStackRoutingInterface

> DeleteNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Delete a layer 3 interface from a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


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


## DeleteNetworkSwitchStackRoutingStaticRoute

> DeleteNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Delete a layer 3 static route for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


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


## GetDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner GetDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Return a layer 3 interface for a switch



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
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response GetDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch



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
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaces

> []GetDeviceSwitchRoutingInterfaces200ResponseInner GetDeviceSwitchRoutingInterfaces(ctx, serial).Execute()

List layer 3 interfaces for a switch



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
	resp, r, err := apiClient.RoutingAPI.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetDeviceSwitchRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetDeviceSwitchRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Return a layer 3 static route for a switch



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
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoutes

> []GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetDeviceSwitchRoutingStaticRoutes(ctx, serial).Execute()

List layer 3 static routes for a switch



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
	resp, r, err := apiClient.RoutingAPI.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetDeviceSwitchRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetDeviceSwitchRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchRoutingMulticast`: %v\n", resp)
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
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
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
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoints`: []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints`: %v\n", resp)
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


## GetNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response GetNetworkSwitchRoutingOspf(ctx, networkId).Execute()

Return layer 3 OSPF routing configuration



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
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner GetNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface from a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response GetNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaces

> []GetDeviceSwitchRoutingInterfaces200ResponseInner GetNetworkSwitchStackRoutingInterfaces(ctx, networkId, switchStackId).Execute()

List layer 3 interfaces for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchStackRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchStackRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Return a layer 3 static route for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoutes

> []GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetNetworkSwitchStackRoutingStaticRoutes(ctx, networkId, switchStackId).Execute()

List layer 3 static routes for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetNetworkSwitchStackRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetNetworkSwitchStackRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner UpdateDeviceSwitchRoutingInterface(ctx, serial, interfaceId).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()

Update a layer 3 interface for a switch



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
	interfaceId := "interfaceId_example" // string | Interface ID
	createDeviceSwitchRoutingInterfaceRequest := *openapiclient.NewCreateDeviceSwitchRoutingInterfaceRequest() // CreateDeviceSwitchRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDeviceSwitchRoutingInterfaceRequest** | [**CreateDeviceSwitchRoutingInterfaceRequest**](CreateDeviceSwitchRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response UpdateDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()

Update a layer 3 interface DHCP configuration for a switch



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
	interfaceId := "interfaceId_example" // string | Interface ID
	updateDeviceSwitchRoutingInterfaceDhcpRequest := *openapiclient.NewUpdateDeviceSwitchRoutingInterfaceDhcpRequest() // UpdateDeviceSwitchRoutingInterfaceDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingInterfaceDhcpRequest** | [**UpdateDeviceSwitchRoutingInterfaceDhcpRequest**](UpdateDeviceSwitchRoutingInterfaceDhcpRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner UpdateDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()

Update a layer 3 static route for a switch



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
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewUpdateDeviceSwitchRoutingStaticRouteRequest() // UpdateDeviceSwitchRoutingStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingStaticRouteRequest** | [**UpdateDeviceSwitchRoutingStaticRouteRequest**](UpdateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).UpdateNetworkSwitchRoutingMulticastRequest(updateNetworkSwitchRoutingMulticastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchRoutingMulticast`: %v\n", resp)
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
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
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


## UpdateNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response UpdateNetworkSwitchRoutingOspf(ctx, networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()

Update layer 3 OSPF routing configuration



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
	updateNetworkSwitchRoutingOspfRequest := *openapiclient.NewUpdateNetworkSwitchRoutingOspfRequest() // UpdateNetworkSwitchRoutingOspfRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingOspfRequest** | [**UpdateNetworkSwitchRoutingOspfRequest**](UpdateNetworkSwitchRoutingOspfRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterface

> UpdateNetworkSwitchStackRoutingInterface200Response UpdateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceRequest(updateNetworkSwitchStackRoutingInterfaceRequest).Execute()

Update a layer 3 interface for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID
	updateNetworkSwitchStackRoutingInterfaceRequest := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceRequest() // UpdateNetworkSwitchStackRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceRequest(updateNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterface`: UpdateNetworkSwitchStackRoutingInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceRequest** | [**UpdateNetworkSwitchStackRoutingInterfaceRequest**](UpdateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

### Return type

[**UpdateNetworkSwitchStackRoutingInterface200Response**](UpdateNetworkSwitchStackRoutingInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()

Update a layer 3 interface DHCP configuration for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID
	updateNetworkSwitchStackRoutingInterfaceDhcpRequest := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest() // UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceDhcpRequest** | [**UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest**](UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner UpdateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()

Update a layer 3 static route for a switch stack



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
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewUpdateDeviceSwitchRoutingStaticRouteRequest() // UpdateDeviceSwitchRoutingStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.UpdateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.UpdateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateDeviceSwitchRoutingStaticRouteRequest** | [**UpdateDeviceSwitchRoutingStaticRouteRequest**](UpdateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \StaticRoutesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSwitchRoutingStaticRoute**](StaticRoutesAPI.md#CreateDeviceSwitchRoutingStaticRoute) | **Post** /devices/{serial}/switch/routing/staticRoutes | Create a layer 3 static route for a switch
[**CreateNetworkApplianceStaticRoute**](StaticRoutesAPI.md#CreateNetworkApplianceStaticRoute) | **Post** /networks/{networkId}/appliance/staticRoutes | Add a static route for an MX or teleworker network
[**CreateNetworkSwitchStackRoutingStaticRoute**](StaticRoutesAPI.md#CreateNetworkSwitchStackRoutingStaticRoute) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | Create a layer 3 static route for a switch stack
[**DeleteDeviceSwitchRoutingStaticRoute**](StaticRoutesAPI.md#DeleteDeviceSwitchRoutingStaticRoute) | **Delete** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch
[**DeleteNetworkApplianceStaticRoute**](StaticRoutesAPI.md#DeleteNetworkApplianceStaticRoute) | **Delete** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Delete a static route from an MX or teleworker network
[**DeleteNetworkSwitchStackRoutingStaticRoute**](StaticRoutesAPI.md#DeleteNetworkSwitchStackRoutingStaticRoute) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch stack
[**GetDeviceSwitchRoutingStaticRoute**](StaticRoutesAPI.md#GetDeviceSwitchRoutingStaticRoute) | **Get** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch
[**GetDeviceSwitchRoutingStaticRoutes**](StaticRoutesAPI.md#GetDeviceSwitchRoutingStaticRoutes) | **Get** /devices/{serial}/switch/routing/staticRoutes | List layer 3 static routes for a switch
[**GetNetworkApplianceStaticRoute**](StaticRoutesAPI.md#GetNetworkApplianceStaticRoute) | **Get** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Return a static route for an MX or teleworker network
[**GetNetworkApplianceStaticRoutes**](StaticRoutesAPI.md#GetNetworkApplianceStaticRoutes) | **Get** /networks/{networkId}/appliance/staticRoutes | List the static routes for an MX or teleworker network
[**GetNetworkSwitchStackRoutingStaticRoute**](StaticRoutesAPI.md#GetNetworkSwitchStackRoutingStaticRoute) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoutes**](StaticRoutesAPI.md#GetNetworkSwitchStackRoutingStaticRoutes) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | List layer 3 static routes for a switch stack
[**UpdateDeviceSwitchRoutingStaticRoute**](StaticRoutesAPI.md#UpdateDeviceSwitchRoutingStaticRoute) | **Put** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch
[**UpdateNetworkApplianceStaticRoute**](StaticRoutesAPI.md#UpdateNetworkApplianceStaticRoute) | **Put** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Update a static route for an MX or teleworker network
[**UpdateNetworkSwitchStackRoutingStaticRoute**](StaticRoutesAPI.md#UpdateNetworkSwitchStackRoutingStaticRoute) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch stack



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
	resp, r, err := apiClient.StaticRoutesAPI.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.CreateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.CreateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
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


## CreateNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner CreateNetworkApplianceStaticRoute(ctx, networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()

Add a static route for an MX or teleworker network



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
	createNetworkApplianceStaticRouteRequest := *openapiclient.NewCreateNetworkApplianceStaticRouteRequest("Name_example", "Subnet_example", "GatewayIp_example") // CreateNetworkApplianceStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticRoutesAPI.CreateNetworkApplianceStaticRoute(context.Background(), networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.CreateNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.CreateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceStaticRouteRequest** | [**CreateNetworkApplianceStaticRouteRequest**](CreateNetworkApplianceStaticRouteRequest.md) |  | 

### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

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
	resp, r, err := apiClient.StaticRoutesAPI.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.CreateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.CreateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
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
	r, err := apiClient.StaticRoutesAPI.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.DeleteDeviceSwitchRoutingStaticRoute``: %v\n", err)
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


## DeleteNetworkApplianceStaticRoute

> DeleteNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Delete a static route from an MX or teleworker network



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
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StaticRoutesAPI.DeleteNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.DeleteNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceStaticRouteRequest struct via the builder pattern


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
	r, err := apiClient.StaticRoutesAPI.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.DeleteNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
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
	resp, r, err := apiClient.StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoute`: %v\n", resp)
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
	resp, r, err := apiClient.StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetDeviceSwitchRoutingStaticRoutes`: %v\n", resp)
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


## GetNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner GetNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Return a static route for an MX or teleworker network



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
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticRoutesAPI.GetNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoutes

> []GetNetworkApplianceStaticRoutes200ResponseInner GetNetworkApplianceStaticRoutes(ctx, networkId).Execute()

List the static routes for an MX or teleworker network



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
	resp, r, err := apiClient.StaticRoutesAPI.GetNetworkApplianceStaticRoutes(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetNetworkApplianceStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceStaticRoutes`: []GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetNetworkApplianceStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

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
	resp, r, err := apiClient.StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
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
	resp, r, err := apiClient.StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.GetNetworkSwitchStackRoutingStaticRoutes`: %v\n", resp)
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
	resp, r, err := apiClient.StaticRoutesAPI.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.UpdateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.UpdateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
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


## UpdateNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner UpdateNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()

Update a static route for an MX or teleworker network



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
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateNetworkApplianceStaticRouteRequest := *openapiclient.NewUpdateNetworkApplianceStaticRouteRequest() // UpdateNetworkApplianceStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StaticRoutesAPI.UpdateNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.UpdateNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.UpdateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceStaticRouteRequest** | [**UpdateNetworkApplianceStaticRouteRequest**](UpdateNetworkApplianceStaticRouteRequest.md) |  | 

### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

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
	resp, r, err := apiClient.StaticRoutesAPI.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticRoutesAPI.UpdateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StaticRoutesAPI.UpdateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
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


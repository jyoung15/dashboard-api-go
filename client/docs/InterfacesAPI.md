# \InterfacesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSwitchRoutingInterface**](InterfacesAPI.md#CreateDeviceSwitchRoutingInterface) | **Post** /devices/{serial}/switch/routing/interfaces | Create a layer 3 interface for a switch
[**CreateNetworkSwitchStackRoutingInterface**](InterfacesAPI.md#CreateNetworkSwitchStackRoutingInterface) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | Create a layer 3 interface for a switch stack
[**DeleteDeviceSwitchRoutingInterface**](InterfacesAPI.md#DeleteDeviceSwitchRoutingInterface) | **Delete** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Delete a layer 3 interface from the switch
[**DeleteNetworkSwitchStackRoutingInterface**](InterfacesAPI.md#DeleteNetworkSwitchStackRoutingInterface) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Delete a layer 3 interface from a switch stack
[**GetDeviceSwitchRoutingInterface**](InterfacesAPI.md#GetDeviceSwitchRoutingInterface) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Return a layer 3 interface for a switch
[**GetDeviceSwitchRoutingInterfaceDhcp**](InterfacesAPI.md#GetDeviceSwitchRoutingInterfaceDhcp) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch
[**GetDeviceSwitchRoutingInterfaces**](InterfacesAPI.md#GetDeviceSwitchRoutingInterfaces) | **Get** /devices/{serial}/switch/routing/interfaces | List layer 3 interfaces for a switch
[**GetNetworkSwitchStackRoutingInterface**](InterfacesAPI.md#GetNetworkSwitchStackRoutingInterface) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Return a layer 3 interface from a switch stack
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](InterfacesAPI.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**GetNetworkSwitchStackRoutingInterfaces**](InterfacesAPI.md#GetNetworkSwitchStackRoutingInterfaces) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | List layer 3 interfaces for a switch stack
[**UpdateDeviceSwitchRoutingInterface**](InterfacesAPI.md#UpdateDeviceSwitchRoutingInterface) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch
[**UpdateDeviceSwitchRoutingInterfaceDhcp**](InterfacesAPI.md#UpdateDeviceSwitchRoutingInterfaceDhcp) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch
[**UpdateNetworkSwitchStackRoutingInterface**](InterfacesAPI.md#UpdateNetworkSwitchStackRoutingInterface) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch stack
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](InterfacesAPI.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack



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
	resp, r, err := apiClient.InterfacesAPI.CreateDeviceSwitchRoutingInterface(context.Background(), serial).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.CreateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.CreateDeviceSwitchRoutingInterface`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingInterfaceRequest(createNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.CreateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.CreateNetworkSwitchStackRoutingInterface`: %v\n", resp)
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
	r, err := apiClient.InterfacesAPI.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.DeleteDeviceSwitchRoutingInterface``: %v\n", err)
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
	r, err := apiClient.InterfacesAPI.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.DeleteNetworkSwitchStackRoutingInterface``: %v\n", err)
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
	resp, r, err := apiClient.InterfacesAPI.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetDeviceSwitchRoutingInterface`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetDeviceSwitchRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetDeviceSwitchRoutingInterfaces`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetNetworkSwitchStackRoutingInterface`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.GetNetworkSwitchStackRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.GetNetworkSwitchStackRoutingInterfaces`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.UpdateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.UpdateDeviceSwitchRoutingInterface`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.UpdateDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.UpdateDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceRequest(updateNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.UpdateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterface`: UpdateNetworkSwitchStackRoutingInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.UpdateNetworkSwitchStackRoutingInterface`: %v\n", resp)
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
	resp, r, err := apiClient.InterfacesAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InterfacesAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `InterfacesAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
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


# \DhcpAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceApplianceDhcpSubnets**](DhcpAPI.md#GetDeviceApplianceDhcpSubnets) | **Get** /devices/{serial}/appliance/dhcp/subnets | Return the DHCP subnet information for an appliance
[**GetDeviceSwitchRoutingInterfaceDhcp**](DhcpAPI.md#GetDeviceSwitchRoutingInterfaceDhcp) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch
[**GetNetworkCellularGatewayDhcp**](DhcpAPI.md#GetNetworkCellularGatewayDhcp) | **Get** /networks/{networkId}/cellularGateway/dhcp | List common DHCP settings of MGs
[**GetNetworkSwitchDhcpV4ServersSeen**](DhcpAPI.md#GetNetworkSwitchDhcpV4ServersSeen) | **Get** /networks/{networkId}/switch/dhcp/v4/servers/seen | Return the network&#39;s DHCPv4 servers seen within the selected timeframe (default 1 day)
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](DhcpAPI.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**UpdateDeviceSwitchRoutingInterfaceDhcp**](DhcpAPI.md#UpdateDeviceSwitchRoutingInterfaceDhcp) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch
[**UpdateNetworkCellularGatewayDhcp**](DhcpAPI.md#UpdateNetworkCellularGatewayDhcp) | **Put** /networks/{networkId}/cellularGateway/dhcp | Update common DHCP settings of MGs
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](DhcpAPI.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack



## GetDeviceApplianceDhcpSubnets

> []map[string]interface{} GetDeviceApplianceDhcpSubnets(ctx, serial).Execute()

Return the DHCP subnet information for an appliance



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
	resp, r, err := apiClient.DhcpAPI.GetDeviceApplianceDhcpSubnets(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.GetDeviceApplianceDhcpSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApplianceDhcpSubnets`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.GetDeviceApplianceDhcpSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceDhcpSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

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
	resp, r, err := apiClient.DhcpAPI.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.GetDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.GetDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
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


## GetNetworkCellularGatewayDhcp

> GetNetworkCellularGatewayDhcp200Response GetNetworkCellularGatewayDhcp(ctx, networkId).Execute()

List common DHCP settings of MGs



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
	resp, r, err := apiClient.DhcpAPI.GetNetworkCellularGatewayDhcp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.GetNetworkCellularGatewayDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewayDhcp`: GetNetworkCellularGatewayDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.GetNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkCellularGatewayDhcp200Response**](GetNetworkCellularGatewayDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpV4ServersSeen

> []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner GetNetworkSwitchDhcpV4ServersSeen(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpAPI.GetNetworkSwitchDhcpV4ServersSeen(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.GetNetworkSwitchDhcpV4ServersSeen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpV4ServersSeen`: []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.GetNetworkSwitchDhcpV4ServersSeen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpV4ServersSeenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInner**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInner.md)

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
	resp, r, err := apiClient.DhcpAPI.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
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
	resp, r, err := apiClient.DhcpAPI.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.UpdateDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.UpdateDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
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


## UpdateNetworkCellularGatewayDhcp

> GetNetworkCellularGatewayDhcp200Response UpdateNetworkCellularGatewayDhcp(ctx, networkId).UpdateNetworkCellularGatewayDhcpRequest(updateNetworkCellularGatewayDhcpRequest).Execute()

Update common DHCP settings of MGs



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
	updateNetworkCellularGatewayDhcpRequest := *openapiclient.NewUpdateNetworkCellularGatewayDhcpRequest() // UpdateNetworkCellularGatewayDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpAPI.UpdateNetworkCellularGatewayDhcp(context.Background(), networkId).UpdateNetworkCellularGatewayDhcpRequest(updateNetworkCellularGatewayDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.UpdateNetworkCellularGatewayDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewayDhcp`: GetNetworkCellularGatewayDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.UpdateNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayDhcpRequest** | [**UpdateNetworkCellularGatewayDhcpRequest**](UpdateNetworkCellularGatewayDhcpRequest.md) |  | 

### Return type

[**GetNetworkCellularGatewayDhcp200Response**](GetNetworkCellularGatewayDhcp200Response.md)

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
	resp, r, err := apiClient.DhcpAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
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


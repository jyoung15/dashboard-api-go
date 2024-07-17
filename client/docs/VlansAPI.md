# \VlansAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceVlan**](VlansAPI.md#CreateNetworkApplianceVlan) | **Post** /networks/{networkId}/appliance/vlans | Add a VLAN
[**DeleteNetworkApplianceVlan**](VlansAPI.md#DeleteNetworkApplianceVlan) | **Delete** /networks/{networkId}/appliance/vlans/{vlanId} | Delete a VLAN from a network
[**GetNetworkApplianceVlan**](VlansAPI.md#GetNetworkApplianceVlan) | **Get** /networks/{networkId}/appliance/vlans/{vlanId} | Return a VLAN
[**GetNetworkApplianceVlans**](VlansAPI.md#GetNetworkApplianceVlans) | **Get** /networks/{networkId}/appliance/vlans | List the VLANs for an MX network
[**GetNetworkApplianceVlansSettings**](VlansAPI.md#GetNetworkApplianceVlansSettings) | **Get** /networks/{networkId}/appliance/vlans/settings | Returns the enabled status of VLANs for the network
[**UpdateNetworkApplianceVlan**](VlansAPI.md#UpdateNetworkApplianceVlan) | **Put** /networks/{networkId}/appliance/vlans/{vlanId} | Update a VLAN
[**UpdateNetworkApplianceVlansSettings**](VlansAPI.md#UpdateNetworkApplianceVlansSettings) | **Put** /networks/{networkId}/appliance/vlans/settings | Enable/Disable VLANs for the given network



## CreateNetworkApplianceVlan

> CreateNetworkApplianceVlan201Response CreateNetworkApplianceVlan(ctx, networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()

Add a VLAN



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
	createNetworkApplianceVlanRequest := *openapiclient.NewCreateNetworkApplianceVlanRequest("Id_example", "Name_example") // CreateNetworkApplianceVlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlansAPI.CreateNetworkApplianceVlan(context.Background(), networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.CreateNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceVlan`: CreateNetworkApplianceVlan201Response
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.CreateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceVlanRequest** | [**CreateNetworkApplianceVlanRequest**](CreateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**CreateNetworkApplianceVlan201Response**](CreateNetworkApplianceVlan201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceVlan

> DeleteNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Delete a VLAN from a network



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
	vlanId := "vlanId_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VlansAPI.DeleteNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.DeleteNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceVlanRequest struct via the builder pattern


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


## GetNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Return a VLAN



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
	vlanId := "vlanId_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlansAPI.GetNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.GetNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.GetNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlans

> []GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlans(ctx, networkId).Execute()

List the VLANs for an MX network



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
	resp, r, err := apiClient.VlansAPI.GetNetworkApplianceVlans(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.GetNetworkApplianceVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlans`: []GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.GetNetworkApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlansSettings

> GetNetworkApplianceVlansSettings200Response GetNetworkApplianceVlansSettings(ctx, networkId).Execute()

Returns the enabled status of VLANs for the network



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
	resp, r, err := apiClient.VlansAPI.GetNetworkApplianceVlansSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.GetNetworkApplianceVlansSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlansSettings`: GetNetworkApplianceVlansSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.GetNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVlansSettings200Response**](GetNetworkApplianceVlansSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner UpdateNetworkApplianceVlan(ctx, networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()

Update a VLAN



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
	vlanId := "vlanId_example" // string | Vlan ID
	updateNetworkApplianceVlanRequest := *openapiclient.NewUpdateNetworkApplianceVlanRequest() // UpdateNetworkApplianceVlanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlansAPI.UpdateNetworkApplianceVlan(context.Background(), networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.UpdateNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.UpdateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceVlanRequest** | [**UpdateNetworkApplianceVlanRequest**](UpdateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlansSettings

> GetNetworkApplianceVlansSettings200Response UpdateNetworkApplianceVlansSettings(ctx, networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()

Enable/Disable VLANs for the given network



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
	updateNetworkApplianceVlansSettingsRequest := *openapiclient.NewUpdateNetworkApplianceVlansSettingsRequest() // UpdateNetworkApplianceVlansSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlansAPI.UpdateNetworkApplianceVlansSettings(context.Background(), networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlansAPI.UpdateNetworkApplianceVlansSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVlansSettings`: GetNetworkApplianceVlansSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `VlansAPI.UpdateNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVlansSettingsRequest** | [**UpdateNetworkApplianceVlansSettingsRequest**](UpdateNetworkApplianceVlansSettingsRequest.md) |  | 

### Return type

[**GetNetworkApplianceVlansSettings200Response**](GetNetworkApplianceVlansSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SnmpAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSnmp**](SnmpAPI.md#GetNetworkSnmp) | **Get** /networks/{networkId}/snmp | Return the SNMP settings for a network
[**GetOrganizationSnmp**](SnmpAPI.md#GetOrganizationSnmp) | **Get** /organizations/{organizationId}/snmp | Return the SNMP settings for an organization
[**UpdateNetworkSnmp**](SnmpAPI.md#UpdateNetworkSnmp) | **Put** /networks/{networkId}/snmp | Update the SNMP settings for a network
[**UpdateOrganizationSnmp**](SnmpAPI.md#UpdateOrganizationSnmp) | **Put** /organizations/{organizationId}/snmp | Update the SNMP settings for an organization



## GetNetworkSnmp

> GetNetworkSnmp200Response GetNetworkSnmp(ctx, networkId).Execute()

Return the SNMP settings for a network



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
	resp, r, err := apiClient.SnmpAPI.GetNetworkSnmp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpAPI.GetNetworkSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSnmp`: GetNetworkSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `SnmpAPI.GetNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSnmp200Response**](GetNetworkSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSnmp

> GetOrganizationSnmp200Response GetOrganizationSnmp(ctx, organizationId).Execute()

Return the SNMP settings for an organization



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
	resp, r, err := apiClient.SnmpAPI.GetOrganizationSnmp(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpAPI.GetOrganizationSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSnmp`: GetOrganizationSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `SnmpAPI.GetOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationSnmp200Response**](GetOrganizationSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSnmp

> GetNetworkSnmp200Response UpdateNetworkSnmp(ctx, networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()

Update the SNMP settings for a network



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
	updateNetworkSnmpRequest := *openapiclient.NewUpdateNetworkSnmpRequest() // UpdateNetworkSnmpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnmpAPI.UpdateNetworkSnmp(context.Background(), networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpAPI.UpdateNetworkSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSnmp`: GetNetworkSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `SnmpAPI.UpdateNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSnmpRequest** | [**UpdateNetworkSnmpRequest**](UpdateNetworkSnmpRequest.md) |  | 

### Return type

[**GetNetworkSnmp200Response**](GetNetworkSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSnmp

> GetOrganizationSnmp200Response UpdateOrganizationSnmp(ctx, organizationId).UpdateOrganizationSnmpRequest(updateOrganizationSnmpRequest).Execute()

Update the SNMP settings for an organization



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
	updateOrganizationSnmpRequest := *openapiclient.NewUpdateOrganizationSnmpRequest() // UpdateOrganizationSnmpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnmpAPI.UpdateOrganizationSnmp(context.Background(), organizationId).UpdateOrganizationSnmpRequest(updateOrganizationSnmpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnmpAPI.UpdateOrganizationSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSnmp`: GetOrganizationSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `SnmpAPI.UpdateOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSnmpRequest** | [**UpdateOrganizationSnmpRequest**](UpdateOrganizationSnmpRequest.md) |  | 

### Return type

[**GetOrganizationSnmp200Response**](GetOrganizationSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


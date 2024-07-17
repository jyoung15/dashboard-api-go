# \MonitoredMediaServersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersAPI.md#CreateOrganizationInsightMonitoredMediaServer) | **Post** /organizations/{organizationId}/insight/monitoredMediaServers | Add a media server to be monitored for this organization
[**DeleteOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersAPI.md#DeleteOrganizationInsightMonitoredMediaServer) | **Delete** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Delete a monitored media server from this organization
[**GetOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersAPI.md#GetOrganizationInsightMonitoredMediaServer) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Return a monitored media server for this organization
[**GetOrganizationInsightMonitoredMediaServers**](MonitoredMediaServersAPI.md#GetOrganizationInsightMonitoredMediaServers) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers | List the monitored media servers for this organization
[**UpdateOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersAPI.md#UpdateOrganizationInsightMonitoredMediaServer) | **Put** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Update a monitored media server for this organization



## CreateOrganizationInsightMonitoredMediaServer

> GetOrganizationInsightMonitoredMediaServers200ResponseInner CreateOrganizationInsightMonitoredMediaServer(ctx, organizationId).CreateOrganizationInsightMonitoredMediaServerRequest(createOrganizationInsightMonitoredMediaServerRequest).Execute()

Add a media server to be monitored for this organization



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
	createOrganizationInsightMonitoredMediaServerRequest := *openapiclient.NewCreateOrganizationInsightMonitoredMediaServerRequest("Name_example", "Address_example") // CreateOrganizationInsightMonitoredMediaServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoredMediaServersAPI.CreateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId).CreateOrganizationInsightMonitoredMediaServerRequest(createOrganizationInsightMonitoredMediaServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersAPI.CreateOrganizationInsightMonitoredMediaServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInsightMonitoredMediaServer`: GetOrganizationInsightMonitoredMediaServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersAPI.CreateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInsightMonitoredMediaServerRequest** | [**CreateOrganizationInsightMonitoredMediaServerRequest**](CreateOrganizationInsightMonitoredMediaServerRequest.md) |  | 

### Return type

[**GetOrganizationInsightMonitoredMediaServers200ResponseInner**](GetOrganizationInsightMonitoredMediaServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationInsightMonitoredMediaServer

> DeleteOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Delete a monitored media server from this organization



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
	monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MonitoredMediaServersAPI.DeleteOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersAPI.DeleteOrganizationInsightMonitoredMediaServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


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


## GetOrganizationInsightMonitoredMediaServer

> GetOrganizationInsightMonitoredMediaServers200ResponseInner GetOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Return a monitored media server for this organization



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
	monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInsightMonitoredMediaServer`: GetOrganizationInsightMonitoredMediaServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationInsightMonitoredMediaServers200ResponseInner**](GetOrganizationInsightMonitoredMediaServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInsightMonitoredMediaServers

> []GetOrganizationInsightMonitoredMediaServers200ResponseInner GetOrganizationInsightMonitoredMediaServers(ctx, organizationId).Execute()

List the monitored media servers for this organization



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
	resp, r, err := apiClient.MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInsightMonitoredMediaServers`: []GetOrganizationInsightMonitoredMediaServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersAPI.GetOrganizationInsightMonitoredMediaServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationInsightMonitoredMediaServers200ResponseInner**](GetOrganizationInsightMonitoredMediaServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationInsightMonitoredMediaServer

> GetOrganizationInsightMonitoredMediaServers200ResponseInner UpdateOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServerRequest(updateOrganizationInsightMonitoredMediaServerRequest).Execute()

Update a monitored media server for this organization



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
	monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID
	updateOrganizationInsightMonitoredMediaServerRequest := *openapiclient.NewUpdateOrganizationInsightMonitoredMediaServerRequest() // UpdateOrganizationInsightMonitoredMediaServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoredMediaServersAPI.UpdateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServerRequest(updateOrganizationInsightMonitoredMediaServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersAPI.UpdateOrganizationInsightMonitoredMediaServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationInsightMonitoredMediaServer`: GetOrganizationInsightMonitoredMediaServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersAPI.UpdateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationInsightMonitoredMediaServerRequest** | [**UpdateOrganizationInsightMonitoredMediaServerRequest**](UpdateOrganizationInsightMonitoredMediaServerRequest.md) |  | 

### Return type

[**GetOrganizationInsightMonitoredMediaServers200ResponseInner**](GetOrganizationInsightMonitoredMediaServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


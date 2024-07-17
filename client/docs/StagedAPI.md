# \StagedAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesStagedEvent**](StagedAPI.md#CreateNetworkFirmwareUpgradesStagedEvent) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events | Create a Staged Upgrade Event for a network
[**CreateNetworkFirmwareUpgradesStagedGroup**](StagedAPI.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**DeferNetworkFirmwareUpgradesStagedEvents**](StagedAPI.md#DeferNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/defer | Postpone by 1 week all pending staged upgrade stages for a network
[**DeleteNetworkFirmwareUpgradesStagedGroup**](StagedAPI.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**GetNetworkFirmwareUpgradesStagedEvents**](StagedAPI.md#GetNetworkFirmwareUpgradesStagedEvents) | **Get** /networks/{networkId}/firmwareUpgrades/staged/events | Get the Staged Upgrade Event from a network
[**GetNetworkFirmwareUpgradesStagedGroup**](StagedAPI.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](StagedAPI.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetNetworkFirmwareUpgradesStagedStages**](StagedAPI.md#GetNetworkFirmwareUpgradesStagedStages) | **Get** /networks/{networkId}/firmwareUpgrades/staged/stages | Order of Staged Upgrade Groups in a network
[**RollbacksNetworkFirmwareUpgradesStagedEvents**](StagedAPI.md#RollbacksNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/rollbacks | Rollback a Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedEvents**](StagedAPI.md#UpdateNetworkFirmwareUpgradesStagedEvents) | **Put** /networks/{networkId}/firmwareUpgrades/staged/events | Update the Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedGroup**](StagedAPI.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateNetworkFirmwareUpgradesStagedStages**](StagedAPI.md#UpdateNetworkFirmwareUpgradesStagedStages) | **Put** /networks/{networkId}/firmwareUpgrades/staged/stages | Assign Staged Upgrade Group order in the sequence.



## CreateNetworkFirmwareUpgradesStagedEvent

> GetNetworkFirmwareUpgradesStagedEvents200Response CreateNetworkFirmwareUpgradesStagedEvent(ctx, networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()

Create a Staged Upgrade Event for a network



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
	createNetworkFirmwareUpgradesStagedEventRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedEventRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // CreateNetworkFirmwareUpgradesStagedEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.CreateNetworkFirmwareUpgradesStagedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirmwareUpgradesStagedEvent`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.CreateNetworkFirmwareUpgradesStagedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedEventRequest** | [**CreateNetworkFirmwareUpgradesStagedEventRequest**](CreateNetworkFirmwareUpgradesStagedEventRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner CreateNetworkFirmwareUpgradesStagedGroup(ctx, networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Create a Staged Upgrade Group for a network



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
	createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeferNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response DeferNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Postpone by 1 week all pending staged upgrade stages for a network



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
	resp, r, err := apiClient.StagedAPI.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.DeferNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeferNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.DeferNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeferNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFirmwareUpgradesStagedGroup

> DeleteNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Delete a Staged Upgrade Group



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
	groupId := "groupId_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StagedAPI.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


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


## GetNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response GetNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Get the Staged Upgrade Event from a network



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
	resp, r, err := apiClient.StagedAPI.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.GetNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.GetNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Get a Staged Upgrade Group from a network



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
	groupId := "groupId_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroups

> []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroups(ctx, networkId).Execute()

List of Staged Upgrade Groups in a network



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
	resp, r, err := apiClient.StagedAPI.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedGroups`: []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner GetNetworkFirmwareUpgradesStagedStages(ctx, networkId).Execute()

Order of Staged Upgrade Groups in a network



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
	resp, r, err := apiClient.StagedAPI.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.GetNetworkFirmwareUpgradesStagedStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.GetNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbacksNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response RollbacksNetworkFirmwareUpgradesStagedEvents(ctx, networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Rollback a Staged Upgrade Event for a network



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
	rollbacksNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // RollbacksNetworkFirmwareUpgradesStagedEventsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.RollbacksNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbacksNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.RollbacksNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbacksNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbacksNetworkFirmwareUpgradesStagedEventsRequest** | [**RollbacksNetworkFirmwareUpgradesStagedEventsRequest**](RollbacksNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response UpdateNetworkFirmwareUpgradesStagedEvents(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Update the Staged Upgrade Event for a network



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
	updateNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // UpdateNetworkFirmwareUpgradesStagedEventsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.UpdateNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.UpdateNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedEventsRequest** | [**UpdateNetworkFirmwareUpgradesStagedEventsRequest**](UpdateNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner UpdateNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Update a Staged Upgrade Group for a network



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
	groupId := "groupId_example" // string | Group ID
	createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner UpdateNetworkFirmwareUpgradesStagedStages(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()

Assign Staged Upgrade Group order in the sequence.



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
	updateNetworkFirmwareUpgradesStagedStagesRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedStagesRequest() // UpdateNetworkFirmwareUpgradesStagedStagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StagedAPI.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagedAPI.UpdateNetworkFirmwareUpgradesStagedStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StagedAPI.UpdateNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedStagesRequest** | [**UpdateNetworkFirmwareUpgradesStagedStagesRequest**](UpdateNetworkFirmwareUpgradesStagedStagesRequest.md) |  | 

### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


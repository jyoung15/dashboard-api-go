# \UsersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmUserDeviceProfiles**](UsersAPI.md#GetNetworkSmUserDeviceProfiles) | **Get** /networks/{networkId}/sm/users/{userId}/deviceProfiles | Get the profiles associated with a user
[**GetNetworkSmUserSoftwares**](UsersAPI.md#GetNetworkSmUserSoftwares) | **Get** /networks/{networkId}/sm/users/{userId}/softwares | Get a list of softwares associated with a user
[**GetNetworkSmUsers**](UsersAPI.md#GetNetworkSmUsers) | **Get** /networks/{networkId}/sm/users | List the owners in an SM network with various specified fields and filters



## GetNetworkSmUserDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmUserDeviceProfiles(ctx, networkId, userId).Execute()

Get the profiles associated with a user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetNetworkSmUserDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetNetworkSmUserDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmUserSoftwares(ctx, networkId, userId).Execute()

Get a list of softwares associated with a user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetNetworkSmUserSoftwares(context.Background(), networkId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetNetworkSmUserSoftwares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetNetworkSmUserSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUsers

> []GetNetworkSmUsers200ResponseInner GetNetworkSmUsers(ctx, networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()

List the owners in an SM network with various specified fields and filters



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
	ids := []string{"Inner_example"} // []string | Filter users by id(s). (optional)
	usernames := []string{"Inner_example"} // []string | Filter users by username(s). (optional)
	emails := []string{"Inner_example"} // []string | Filter users by email(s). (optional)
	scope := []string{"Inner_example"} // []string | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetNetworkSmUsers(context.Background(), networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetNetworkSmUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUsers`: []GetNetworkSmUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetNetworkSmUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Filter users by id(s). | 
 **usernames** | **[]string** | Filter users by username(s). | 
 **emails** | **[]string** | Filter users by email(s). | 
 **scope** | **[]string** | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. | 

### Return type

[**[]GetNetworkSmUsers200ResponseInner**](GetNetworkSmUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AssignmentsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkVlanProfilesAssignmentsByDevice**](AssignmentsAPI.md#GetNetworkVlanProfilesAssignmentsByDevice) | **Get** /networks/{networkId}/vlanProfiles/assignments/byDevice | Get the assigned VLAN Profiles for devices in a network
[**GetOrganizationSmSentryPoliciesAssignmentsByNetwork**](AssignmentsAPI.md#GetOrganizationSmSentryPoliciesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork | List the Sentry Policies for an organization ordered in ascending order of priority
[**GetOrganizationWirelessRfProfilesAssignmentsByDevice**](AssignmentsAPI.md#GetOrganizationWirelessRfProfilesAssignmentsByDevice) | **Get** /organizations/{organizationId}/wireless/rfProfiles/assignments/byDevice | List the RF profiles of an organization by device
[**ReassignNetworkVlanProfilesAssignments**](AssignmentsAPI.md#ReassignNetworkVlanProfilesAssignments) | **Post** /networks/{networkId}/vlanProfiles/assignments/reassign | Update the assigned VLAN Profile for devices in a network
[**UpdateOrganizationSmSentryPoliciesAssignments**](AssignmentsAPI.md#UpdateOrganizationSmSentryPoliciesAssignments) | **Put** /organizations/{organizationId}/sm/sentry/policies/assignments | Update an Organizations Sentry Policies using the provided list



## GetNetworkVlanProfilesAssignmentsByDevice

> []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner GetNetworkVlanProfilesAssignmentsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()

Get the assigned VLAN Profiles for devices in a network



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product types. (optional)
	stackIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by Switch Stack ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssignmentsAPI.GetNetworkVlanProfilesAssignmentsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsAPI.GetNetworkVlanProfilesAssignmentsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkVlanProfilesAssignmentsByDevice`: []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AssignmentsAPI.GetNetworkVlanProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product types. | 
 **stackIds** | **[]string** | Optional parameter to filter devices by Switch Stack ids. | 

### Return type

[**[]GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner**](GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmSentryPoliciesAssignmentsByNetwork

> []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List the Sentry Policies for an organization ordered in ascending order of priority



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter Sentry Policies by Network Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssignmentsAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AssignmentsAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter Sentry Policies by Network Id | 

### Return type

[**[]GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner**](GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRfProfilesAssignmentsByDevice

> []GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner GetOrganizationWirelessRfProfilesAssignmentsByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).Models(models).Execute()

List the RF profiles of an organization by device



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. (optional)
	name := "name_example" // string | Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match. (optional)
	mac := "mac_example" // string | Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. (optional)
	serial := "serial_example" // string | Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match. (optional)
	model := "model_example" // string | Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match. (optional)
	macs := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match. (optional)
	models := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssignmentsAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).Models(models).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessRfProfilesAssignmentsByDevice`: []GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AssignmentsAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. | 
 **name** | **string** | Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. | 
 **serial** | **string** | Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match. | 
 **model** | **string** | Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match. | 
 **models** | **[]string** | Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match. | 

### Return type

[**[]GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner**](GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignNetworkVlanProfilesAssignments

> ReassignNetworkVlanProfilesAssignments200Response ReassignNetworkVlanProfilesAssignments(ctx, networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()

Update the assigned VLAN Profile for devices in a network



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
	reassignNetworkVlanProfilesAssignmentsRequest := *openapiclient.NewReassignNetworkVlanProfilesAssignmentsRequest([]string{"Serials_example"}, []string{"StackIds_example"}) // ReassignNetworkVlanProfilesAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssignmentsAPI.ReassignNetworkVlanProfilesAssignments(context.Background(), networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsAPI.ReassignNetworkVlanProfilesAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReassignNetworkVlanProfilesAssignments`: ReassignNetworkVlanProfilesAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `AssignmentsAPI.ReassignNetworkVlanProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignNetworkVlanProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reassignNetworkVlanProfilesAssignmentsRequest** | [**ReassignNetworkVlanProfilesAssignmentsRequest**](ReassignNetworkVlanProfilesAssignmentsRequest.md) |  | 

### Return type

[**ReassignNetworkVlanProfilesAssignments200Response**](ReassignNetworkVlanProfilesAssignments200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmSentryPoliciesAssignments

> UpdateOrganizationSmSentryPoliciesAssignments200Response UpdateOrganizationSmSentryPoliciesAssignments(ctx, organizationId).UpdateOrganizationSmSentryPoliciesAssignmentsRequest(updateOrganizationSmSentryPoliciesAssignmentsRequest).Execute()

Update an Organizations Sentry Policies using the provided list



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
	updateOrganizationSmSentryPoliciesAssignmentsRequest := *openapiclient.NewUpdateOrganizationSmSentryPoliciesAssignmentsRequest([]openapiclient.UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{*openapiclient.NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner("NetworkId_example")}) // UpdateOrganizationSmSentryPoliciesAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssignmentsAPI.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).UpdateOrganizationSmSentryPoliciesAssignmentsRequest(updateOrganizationSmSentryPoliciesAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsAPI.UpdateOrganizationSmSentryPoliciesAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSmSentryPoliciesAssignments`: UpdateOrganizationSmSentryPoliciesAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `AssignmentsAPI.UpdateOrganizationSmSentryPoliciesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSmSentryPoliciesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSmSentryPoliciesAssignmentsRequest** | [**UpdateOrganizationSmSentryPoliciesAssignmentsRequest**](UpdateOrganizationSmSentryPoliciesAssignmentsRequest.md) |  | 

### Return type

[**UpdateOrganizationSmSentryPoliciesAssignments200Response**](UpdateOrganizationSmSentryPoliciesAssignments200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


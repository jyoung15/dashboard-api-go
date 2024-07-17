# \VlanAssignmentsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceAppliancePrefixesDelegatedVlanAssignments**](VlanAssignmentsAPI.md#GetDeviceAppliancePrefixesDelegatedVlanAssignments) | **Get** /devices/{serial}/appliance/prefixes/delegated/vlanAssignments | Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



## GetDeviceAppliancePrefixesDelegatedVlanAssignments

> []map[string]interface{} GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx, serial).Execute()

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



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
	resp, r, err := apiClient.VlanAssignmentsAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAssignmentsAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePrefixesDelegatedVlanAssignments`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VlanAssignmentsAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct via the builder pattern


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


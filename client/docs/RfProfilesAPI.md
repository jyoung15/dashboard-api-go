# \RfProfilesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceRfProfile**](RfProfilesAPI.md#CreateNetworkApplianceRfProfile) | **Post** /networks/{networkId}/appliance/rfProfiles | Creates new RF profile for this network
[**CreateNetworkWirelessRfProfile**](RfProfilesAPI.md#CreateNetworkWirelessRfProfile) | **Post** /networks/{networkId}/wireless/rfProfiles | Creates new RF profile for this network
[**DeleteNetworkApplianceRfProfile**](RfProfilesAPI.md#DeleteNetworkApplianceRfProfile) | **Delete** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkWirelessRfProfile**](RfProfilesAPI.md#DeleteNetworkWirelessRfProfile) | **Delete** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Delete a RF Profile
[**GetNetworkApplianceRfProfile**](RfProfilesAPI.md#GetNetworkApplianceRfProfile) | **Get** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkApplianceRfProfiles**](RfProfilesAPI.md#GetNetworkApplianceRfProfiles) | **Get** /networks/{networkId}/appliance/rfProfiles | List the RF profiles for this network
[**GetNetworkWirelessRfProfile**](RfProfilesAPI.md#GetNetworkWirelessRfProfile) | **Get** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkWirelessRfProfiles**](RfProfilesAPI.md#GetNetworkWirelessRfProfiles) | **Get** /networks/{networkId}/wireless/rfProfiles | List RF profiles for this network
[**GetOrganizationWirelessRfProfilesAssignmentsByDevice**](RfProfilesAPI.md#GetOrganizationWirelessRfProfilesAssignmentsByDevice) | **Get** /organizations/{organizationId}/wireless/rfProfiles/assignments/byDevice | List the RF profiles of an organization by device
[**UpdateNetworkApplianceRfProfile**](RfProfilesAPI.md#UpdateNetworkApplianceRfProfile) | **Put** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkWirelessRfProfile**](RfProfilesAPI.md#UpdateNetworkWirelessRfProfile) | **Put** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Updates specified RF profile for this network



## CreateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner CreateNetworkApplianceRfProfile(ctx, networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()

Creates new RF profile for this network



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
	createNetworkApplianceRfProfileRequest := *openapiclient.NewCreateNetworkApplianceRfProfileRequest("Name_example") // CreateNetworkApplianceRfProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.CreateNetworkApplianceRfProfile(context.Background(), networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.CreateNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.CreateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceRfProfileRequest** | [**CreateNetworkApplianceRfProfileRequest**](CreateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response CreateNetworkWirelessRfProfile(ctx, networkId).CreateNetworkWirelessRfProfileRequest(createNetworkWirelessRfProfileRequest).Execute()

Creates new RF profile for this network



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
	createNetworkWirelessRfProfileRequest := *openapiclient.NewCreateNetworkWirelessRfProfileRequest("Name_example", "BandSelectionType_example") // CreateNetworkWirelessRfProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.CreateNetworkWirelessRfProfile(context.Background(), networkId).CreateNetworkWirelessRfProfileRequest(createNetworkWirelessRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.CreateNetworkWirelessRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.CreateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessRfProfileRequest** | [**CreateNetworkWirelessRfProfileRequest**](CreateNetworkWirelessRfProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceRfProfile

> DeleteNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RfProfilesAPI.DeleteNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.DeleteNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceRfProfileRequest struct via the builder pattern


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


## DeleteNetworkWirelessRfProfile

> DeleteNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RfProfilesAPI.DeleteNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.DeleteNetworkWirelessRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessRfProfileRequest struct via the builder pattern


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


## GetNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner GetNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.GetNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.GetNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.GetNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfiles

> GetNetworkApplianceRfProfiles200Response GetNetworkApplianceRfProfiles(ctx, networkId).Execute()

List the RF profiles for this network



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
	resp, r, err := apiClient.RfProfilesAPI.GetNetworkApplianceRfProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.GetNetworkApplianceRfProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceRfProfiles`: GetNetworkApplianceRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.GetNetworkApplianceRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceRfProfiles200Response**](GetNetworkApplianceRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response GetNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.GetNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.GetNetworkWirelessRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.GetNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfiles

> GetNetworkWirelessRfProfiles200Response GetNetworkWirelessRfProfiles(ctx, networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()

List RF profiles for this network



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
	includeTemplateProfiles := true // bool | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.GetNetworkWirelessRfProfiles(context.Background(), networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.GetNetworkWirelessRfProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessRfProfiles`: GetNetworkWirelessRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.GetNetworkWirelessRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTemplateProfiles** | **bool** | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

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
	resp, r, err := apiClient.RfProfilesAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).Models(models).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessRfProfilesAssignmentsByDevice`: []GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.GetOrganizationWirelessRfProfilesAssignmentsByDevice`: %v\n", resp)
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


## UpdateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner UpdateNetworkApplianceRfProfile(ctx, networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()

Updates specified RF profile for this network



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID
	updateNetworkApplianceRfProfileRequest := *openapiclient.NewUpdateNetworkApplianceRfProfileRequest() // UpdateNetworkApplianceRfProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.UpdateNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.UpdateNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.UpdateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceRfProfileRequest** | [**UpdateNetworkApplianceRfProfileRequest**](UpdateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response UpdateNetworkWirelessRfProfile(ctx, networkId, rfProfileId).UpdateNetworkWirelessRfProfileRequest(updateNetworkWirelessRfProfileRequest).Execute()

Updates specified RF profile for this network



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
	rfProfileId := "rfProfileId_example" // string | Rf profile ID
	updateNetworkWirelessRfProfileRequest := *openapiclient.NewUpdateNetworkWirelessRfProfileRequest() // UpdateNetworkWirelessRfProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RfProfilesAPI.UpdateNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkWirelessRfProfileRequest(updateNetworkWirelessRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesAPI.UpdateNetworkWirelessRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RfProfilesAPI.UpdateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessRfProfileRequest** | [**UpdateNetworkWirelessRfProfileRequest**](UpdateNetworkWirelessRfProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


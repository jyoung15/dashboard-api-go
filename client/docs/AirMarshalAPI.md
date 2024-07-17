# \AirMarshalAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessAirMarshalRule**](AirMarshalAPI.md#CreateNetworkWirelessAirMarshalRule) | **Post** /networks/{networkId}/wireless/airMarshal/rules | Creates a new rule
[**DeleteNetworkWirelessAirMarshalRule**](AirMarshalAPI.md#DeleteNetworkWirelessAirMarshalRule) | **Delete** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Delete an Air Marshal rule.
[**GetNetworkWirelessAirMarshal**](AirMarshalAPI.md#GetNetworkWirelessAirMarshal) | **Get** /networks/{networkId}/wireless/airMarshal | List Air Marshal scan results from a network
[**GetOrganizationWirelessAirMarshalRules**](AirMarshalAPI.md#GetOrganizationWirelessAirMarshalRules) | **Get** /organizations/{organizationId}/wireless/airMarshal/rules | Returns the current Air Marshal rules for this organization
[**GetOrganizationWirelessAirMarshalSettingsByNetwork**](AirMarshalAPI.md#GetOrganizationWirelessAirMarshalSettingsByNetwork) | **Get** /organizations/{organizationId}/wireless/airMarshal/settings/byNetwork | Returns the current Air Marshal settings for this network
[**UpdateNetworkWirelessAirMarshalRule**](AirMarshalAPI.md#UpdateNetworkWirelessAirMarshalRule) | **Put** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Update a rule
[**UpdateNetworkWirelessAirMarshalSettings**](AirMarshalAPI.md#UpdateNetworkWirelessAirMarshalSettings) | **Put** /networks/{networkId}/wireless/airMarshal/settings | Updates Air Marshal settings.



## CreateNetworkWirelessAirMarshalRule

> CreateNetworkWirelessAirMarshalRule201Response CreateNetworkWirelessAirMarshalRule(ctx, networkId).CreateNetworkWirelessAirMarshalRuleRequest(createNetworkWirelessAirMarshalRuleRequest).Execute()

Creates a new rule



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
	createNetworkWirelessAirMarshalRuleRequest := *openapiclient.NewCreateNetworkWirelessAirMarshalRuleRequest("Type_example", *openapiclient.NewCreateNetworkWirelessAirMarshalRuleRequestMatch()) // CreateNetworkWirelessAirMarshalRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.CreateNetworkWirelessAirMarshalRule(context.Background(), networkId).CreateNetworkWirelessAirMarshalRuleRequest(createNetworkWirelessAirMarshalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.CreateNetworkWirelessAirMarshalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWirelessAirMarshalRule`: CreateNetworkWirelessAirMarshalRule201Response
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.CreateNetworkWirelessAirMarshalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessAirMarshalRuleRequest** | [**CreateNetworkWirelessAirMarshalRuleRequest**](CreateNetworkWirelessAirMarshalRuleRequest.md) |  | 

### Return type

[**CreateNetworkWirelessAirMarshalRule201Response**](CreateNetworkWirelessAirMarshalRule201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWirelessAirMarshalRule

> DeleteNetworkWirelessAirMarshalRule(ctx, networkId, ruleId).Execute()

Delete an Air Marshal rule.



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
	ruleId := "ruleId_example" // string | Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AirMarshalAPI.DeleteNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.DeleteNetworkWirelessAirMarshalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


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


## GetNetworkWirelessAirMarshal

> []map[string]interface{} GetNetworkWirelessAirMarshal(ctx, networkId).T0(t0).Timespan(timespan).Execute()

List Air Marshal scan results from a network



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
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.GetNetworkWirelessAirMarshal(context.Background(), networkId).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.GetNetworkWirelessAirMarshal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessAirMarshal`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.GetNetworkWirelessAirMarshal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessAirMarshalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 

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


## GetOrganizationWirelessAirMarshalRules

> GetOrganizationWirelessAirMarshalRules200Response GetOrganizationWirelessAirMarshalRules(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the current Air Marshal rules for this organization



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
	networkIds := []string{"Inner_example"} // []string | (optional) The set of network IDs to include. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.GetOrganizationWirelessAirMarshalRules(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.GetOrganizationWirelessAirMarshalRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessAirMarshalRules`: GetOrganizationWirelessAirMarshalRules200Response
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.GetOrganizationWirelessAirMarshalRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessAirMarshalRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | (optional) The set of network IDs to include. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationWirelessAirMarshalRules200Response**](GetOrganizationWirelessAirMarshalRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessAirMarshalSettingsByNetwork

> GetOrganizationWirelessAirMarshalSettingsByNetwork200Response GetOrganizationWirelessAirMarshalSettingsByNetwork(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the current Air Marshal settings for this network



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
	networkIds := []string{"Inner_example"} // []string | The network IDs to include in the result set. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.GetOrganizationWirelessAirMarshalSettingsByNetwork(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.GetOrganizationWirelessAirMarshalSettingsByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessAirMarshalSettingsByNetwork`: GetOrganizationWirelessAirMarshalSettingsByNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.GetOrganizationWirelessAirMarshalSettingsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessAirMarshalSettingsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | The network IDs to include in the result set. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationWirelessAirMarshalSettingsByNetwork200Response**](GetOrganizationWirelessAirMarshalSettingsByNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAirMarshalRule

> CreateNetworkWirelessAirMarshalRule201Response UpdateNetworkWirelessAirMarshalRule(ctx, networkId, ruleId).UpdateNetworkWirelessAirMarshalRuleRequest(updateNetworkWirelessAirMarshalRuleRequest).Execute()

Update a rule



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
	ruleId := "ruleId_example" // string | Rule ID
	updateNetworkWirelessAirMarshalRuleRequest := *openapiclient.NewUpdateNetworkWirelessAirMarshalRuleRequest() // UpdateNetworkWirelessAirMarshalRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.UpdateNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).UpdateNetworkWirelessAirMarshalRuleRequest(updateNetworkWirelessAirMarshalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.UpdateNetworkWirelessAirMarshalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessAirMarshalRule`: CreateNetworkWirelessAirMarshalRule201Response
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.UpdateNetworkWirelessAirMarshalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessAirMarshalRuleRequest** | [**UpdateNetworkWirelessAirMarshalRuleRequest**](UpdateNetworkWirelessAirMarshalRuleRequest.md) |  | 

### Return type

[**CreateNetworkWirelessAirMarshalRule201Response**](CreateNetworkWirelessAirMarshalRule201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAirMarshalSettings

> UpdateNetworkWirelessAirMarshalSettings200Response UpdateNetworkWirelessAirMarshalSettings(ctx, networkId).UpdateNetworkWirelessAirMarshalSettingsRequest(updateNetworkWirelessAirMarshalSettingsRequest).Execute()

Updates Air Marshal settings.



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
	updateNetworkWirelessAirMarshalSettingsRequest := *openapiclient.NewUpdateNetworkWirelessAirMarshalSettingsRequest("DefaultPolicy_example") // UpdateNetworkWirelessAirMarshalSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirMarshalAPI.UpdateNetworkWirelessAirMarshalSettings(context.Background(), networkId).UpdateNetworkWirelessAirMarshalSettingsRequest(updateNetworkWirelessAirMarshalSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirMarshalAPI.UpdateNetworkWirelessAirMarshalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessAirMarshalSettings`: UpdateNetworkWirelessAirMarshalSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AirMarshalAPI.UpdateNetworkWirelessAirMarshalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAirMarshalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessAirMarshalSettingsRequest** | [**UpdateNetworkWirelessAirMarshalSettingsRequest**](UpdateNetworkWirelessAirMarshalSettingsRequest.md) |  | 

### Return type

[**UpdateNetworkWirelessAirMarshalSettings200Response**](UpdateNetworkWirelessAirMarshalSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


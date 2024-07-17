# \RulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessAirMarshalRule**](RulesAPI.md#CreateNetworkWirelessAirMarshalRule) | **Post** /networks/{networkId}/wireless/airMarshal/rules | Creates a new rule
[**DeleteNetworkWirelessAirMarshalRule**](RulesAPI.md#DeleteNetworkWirelessAirMarshalRule) | **Delete** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Delete an Air Marshal rule.
[**GetNetworkApplianceTrafficShapingRules**](RulesAPI.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for an MX network
[**GetNetworkWirelessSsidTrafficShapingRules**](RulesAPI.md#GetNetworkWirelessSsidTrafficShapingRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Display the traffic shaping settings for a SSID on an MR network
[**GetOrganizationWirelessAirMarshalRules**](RulesAPI.md#GetOrganizationWirelessAirMarshalRules) | **Get** /organizations/{organizationId}/wireless/airMarshal/rules | Returns the current Air Marshal rules for this organization
[**UpdateNetworkApplianceTrafficShapingRules**](RulesAPI.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for an MX network
[**UpdateNetworkWirelessAirMarshalRule**](RulesAPI.md#UpdateNetworkWirelessAirMarshalRule) | **Put** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Update a rule
[**UpdateNetworkWirelessSsidTrafficShapingRules**](RulesAPI.md#UpdateNetworkWirelessSsidTrafficShapingRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Update the traffic shaping rules for an SSID on an MR network.



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
	resp, r, err := apiClient.RulesAPI.CreateNetworkWirelessAirMarshalRule(context.Background(), networkId).CreateNetworkWirelessAirMarshalRuleRequest(createNetworkWirelessAirMarshalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateNetworkWirelessAirMarshalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWirelessAirMarshalRule`: CreateNetworkWirelessAirMarshalRule201Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateNetworkWirelessAirMarshalRule`: %v\n", resp)
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
	r, err := apiClient.RulesAPI.DeleteNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteNetworkWirelessAirMarshalRule``: %v\n", err)
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


## GetNetworkApplianceTrafficShapingRules

> map[string]interface{} GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for an MX network



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
	resp, r, err := apiClient.RulesAPI.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetNetworkApplianceTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidTrafficShapingRules

> GetNetworkWirelessSsidTrafficShapingRules200Response GetNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).Execute()

Display the traffic shaping settings for a SSID on an MR network



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
	number := "number_example" // string | Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidTrafficShapingRules`: GetNetworkWirelessSsidTrafficShapingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidTrafficShapingRules200Response**](GetNetworkWirelessSsidTrafficShapingRules200Response.md)

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
	resp, r, err := apiClient.RulesAPI.GetOrganizationWirelessAirMarshalRules(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetOrganizationWirelessAirMarshalRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessAirMarshalRules`: GetOrganizationWirelessAirMarshalRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetOrganizationWirelessAirMarshalRules`: %v\n", resp)
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


## UpdateNetworkApplianceTrafficShapingRules

> map[string]interface{} UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()

Update the traffic shaping settings rules for an MX network



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
	updateNetworkApplianceTrafficShapingRulesRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingRulesRequest() // UpdateNetworkApplianceTrafficShapingRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRulesRequest** | [**UpdateNetworkApplianceTrafficShapingRulesRequest**](UpdateNetworkApplianceTrafficShapingRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.RulesAPI.UpdateNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).UpdateNetworkWirelessAirMarshalRuleRequest(updateNetworkWirelessAirMarshalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateNetworkWirelessAirMarshalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessAirMarshalRule`: CreateNetworkWirelessAirMarshalRule201Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateNetworkWirelessAirMarshalRule`: %v\n", resp)
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


## UpdateNetworkWirelessSsidTrafficShapingRules

> GetNetworkWirelessSsidTrafficShapingRules200Response UpdateNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).UpdateNetworkWirelessSsidTrafficShapingRulesRequest(updateNetworkWirelessSsidTrafficShapingRulesRequest).Execute()

Update the traffic shaping rules for an SSID on an MR network.



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
	number := "number_example" // string | Number
	updateNetworkWirelessSsidTrafficShapingRulesRequest := *openapiclient.NewUpdateNetworkWirelessSsidTrafficShapingRulesRequest() // UpdateNetworkWirelessSsidTrafficShapingRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidTrafficShapingRulesRequest(updateNetworkWirelessSsidTrafficShapingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidTrafficShapingRules`: GetNetworkWirelessSsidTrafficShapingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidTrafficShapingRulesRequest** | [**UpdateNetworkWirelessSsidTrafficShapingRulesRequest**](UpdateNetworkWirelessSsidTrafficShapingRulesRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidTrafficShapingRules200Response**](GetNetworkWirelessSsidTrafficShapingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


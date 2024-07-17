# \QosRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchQosRule**](QosRulesAPI.md#CreateNetworkSwitchQosRule) | **Post** /networks/{networkId}/switch/qosRules | Add a quality of service rule
[**DeleteNetworkSwitchQosRule**](QosRulesAPI.md#DeleteNetworkSwitchQosRule) | **Delete** /networks/{networkId}/switch/qosRules/{qosRuleId} | Delete a quality of service rule
[**GetNetworkSwitchQosRule**](QosRulesAPI.md#GetNetworkSwitchQosRule) | **Get** /networks/{networkId}/switch/qosRules/{qosRuleId} | Return a quality of service rule
[**GetNetworkSwitchQosRules**](QosRulesAPI.md#GetNetworkSwitchQosRules) | **Get** /networks/{networkId}/switch/qosRules | List quality of service rules
[**GetNetworkSwitchQosRulesOrder**](QosRulesAPI.md#GetNetworkSwitchQosRulesOrder) | **Get** /networks/{networkId}/switch/qosRules/order | Return the quality of service rule IDs by order in which they will be processed by the switch
[**UpdateNetworkSwitchQosRule**](QosRulesAPI.md#UpdateNetworkSwitchQosRule) | **Put** /networks/{networkId}/switch/qosRules/{qosRuleId} | Update a quality of service rule
[**UpdateNetworkSwitchQosRulesOrder**](QosRulesAPI.md#UpdateNetworkSwitchQosRulesOrder) | **Put** /networks/{networkId}/switch/qosRules/order | Update the order in which the rules should be processed by the switch



## CreateNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner CreateNetworkSwitchQosRule(ctx, networkId).CreateNetworkSwitchQosRuleRequest(createNetworkSwitchQosRuleRequest).Execute()

Add a quality of service rule



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
	createNetworkSwitchQosRuleRequest := *openapiclient.NewCreateNetworkSwitchQosRuleRequest(int32(123)) // CreateNetworkSwitchQosRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QosRulesAPI.CreateNetworkSwitchQosRule(context.Background(), networkId).CreateNetworkSwitchQosRuleRequest(createNetworkSwitchQosRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.CreateNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.CreateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchQosRuleRequest** | [**CreateNetworkSwitchQosRuleRequest**](CreateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchQosRule

> DeleteNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Delete a quality of service rule



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
	qosRuleId := "qosRuleId_example" // string | Qos rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QosRulesAPI.DeleteNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.DeleteNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchQosRuleRequest struct via the builder pattern


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


## GetNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner GetNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Return a quality of service rule



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
	qosRuleId := "qosRuleId_example" // string | Qos rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QosRulesAPI.GetNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.GetNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.GetNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRules

> []GetNetworkSwitchQosRules200ResponseInner GetNetworkSwitchQosRules(ctx, networkId).Execute()

List quality of service rules



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
	resp, r, err := apiClient.QosRulesAPI.GetNetworkSwitchQosRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.GetNetworkSwitchQosRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRules`: []GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.GetNetworkSwitchQosRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response GetNetworkSwitchQosRulesOrder(ctx, networkId).Execute()

Return the quality of service rule IDs by order in which they will be processed by the switch



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
	resp, r, err := apiClient.QosRulesAPI.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.GetNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.GetNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner UpdateNetworkSwitchQosRule(ctx, networkId, qosRuleId).UpdateNetworkSwitchQosRuleRequest(updateNetworkSwitchQosRuleRequest).Execute()

Update a quality of service rule



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
	qosRuleId := "qosRuleId_example" // string | Qos rule ID
	updateNetworkSwitchQosRuleRequest := *openapiclient.NewUpdateNetworkSwitchQosRuleRequest() // UpdateNetworkSwitchQosRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QosRulesAPI.UpdateNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).UpdateNetworkSwitchQosRuleRequest(updateNetworkSwitchQosRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.UpdateNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.UpdateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchQosRuleRequest** | [**UpdateNetworkSwitchQosRuleRequest**](UpdateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response UpdateNetworkSwitchQosRulesOrder(ctx, networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()

Update the order in which the rules should be processed by the switch



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
	updateNetworkSwitchQosRulesOrderRequest := *openapiclient.NewUpdateNetworkSwitchQosRulesOrderRequest([]string{"RuleIds_example"}) // UpdateNetworkSwitchQosRulesOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QosRulesAPI.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QosRulesAPI.UpdateNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `QosRulesAPI.UpdateNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchQosRulesOrderRequest** | [**UpdateNetworkSwitchQosRulesOrderRequest**](UpdateNetworkSwitchQosRulesOrderRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


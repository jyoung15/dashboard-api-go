# \PayloadTemplatesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWebhooksPayloadTemplate**](PayloadTemplatesAPI.md#CreateNetworkWebhooksPayloadTemplate) | **Post** /networks/{networkId}/webhooks/payloadTemplates | Create a webhook payload template for a network
[**DeleteNetworkWebhooksPayloadTemplate**](PayloadTemplatesAPI.md#DeleteNetworkWebhooksPayloadTemplate) | **Delete** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Destroy a webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplate**](PayloadTemplatesAPI.md#GetNetworkWebhooksPayloadTemplate) | **Get** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Get the webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplates**](PayloadTemplatesAPI.md#GetNetworkWebhooksPayloadTemplates) | **Get** /networks/{networkId}/webhooks/payloadTemplates | List the webhook payload templates for a network
[**UpdateNetworkWebhooksPayloadTemplate**](PayloadTemplatesAPI.md#UpdateNetworkWebhooksPayloadTemplate) | **Put** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Update a webhook payload template for a network



## CreateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner CreateNetworkWebhooksPayloadTemplate(ctx, networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()

Create a webhook payload template for a network



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
	createNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewCreateNetworkWebhooksPayloadTemplateRequest("Name_example") // CreateNetworkWebhooksPayloadTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayloadTemplatesAPI.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesAPI.CreateNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesAPI.CreateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksPayloadTemplateRequest** | [**CreateNetworkWebhooksPayloadTemplateRequest**](CreateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWebhooksPayloadTemplate

> DeleteNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Destroy a webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PayloadTemplatesAPI.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesAPI.DeleteNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


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


## GetNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Get the webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplates

> []GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplates(ctx, networkId).Execute()

List the webhook payload templates for a network



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
	resp, r, err := apiClient.PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksPayloadTemplates`: []GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesAPI.GetNetworkWebhooksPayloadTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner UpdateNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()

Update a webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID
	updateNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewUpdateNetworkWebhooksPayloadTemplateRequest() // UpdateNetworkWebhooksPayloadTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayloadTemplatesAPI.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesAPI.UpdateNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesAPI.UpdateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksPayloadTemplateRequest** | [**UpdateNetworkWebhooksPayloadTemplateRequest**](UpdateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


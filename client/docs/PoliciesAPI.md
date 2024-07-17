# \PoliciesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdaptivePolicyPolicy**](PoliciesAPI.md#CreateOrganizationAdaptivePolicyPolicy) | **Post** /organizations/{organizationId}/adaptivePolicy/policies | Add an Adaptive Policy
[**DeleteOrganizationAdaptivePolicyPolicy**](PoliciesAPI.md#DeleteOrganizationAdaptivePolicyPolicy) | **Delete** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Delete an Adaptive Policy
[**GetNetworkPoliciesByClient**](PoliciesAPI.md#GetNetworkPoliciesByClient) | **Get** /networks/{networkId}/policies/byClient | Get policies for all clients with policies
[**GetOrganizationAdaptivePolicyPolicies**](PoliciesAPI.md#GetOrganizationAdaptivePolicyPolicies) | **Get** /organizations/{organizationId}/adaptivePolicy/policies | List adaptive policies in an organization
[**GetOrganizationAdaptivePolicyPolicy**](PoliciesAPI.md#GetOrganizationAdaptivePolicyPolicy) | **Get** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Return an adaptive policy
[**GetOrganizationSmSentryPoliciesAssignmentsByNetwork**](PoliciesAPI.md#GetOrganizationSmSentryPoliciesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork | List the Sentry Policies for an organization ordered in ascending order of priority
[**UpdateOrganizationAdaptivePolicyPolicy**](PoliciesAPI.md#UpdateOrganizationAdaptivePolicyPolicy) | **Put** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Update an Adaptive Policy
[**UpdateOrganizationSmSentryPoliciesAssignments**](PoliciesAPI.md#UpdateOrganizationSmSentryPoliciesAssignments) | **Put** /organizations/{organizationId}/sm/sentry/policies/assignments | Update an Organizations Sentry Policies using the provided list



## CreateOrganizationAdaptivePolicyPolicy

> GetOrganizationAdaptivePolicyPolicies200ResponseInner CreateOrganizationAdaptivePolicyPolicy(ctx, organizationId).CreateOrganizationAdaptivePolicyPolicyRequest(createOrganizationAdaptivePolicyPolicyRequest).Execute()

Add an Adaptive Policy



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
	createOrganizationAdaptivePolicyPolicyRequest := *openapiclient.NewCreateOrganizationAdaptivePolicyPolicyRequest(*openapiclient.NewCreateOrganizationAdaptivePolicyPolicyRequestSourceGroup(), *openapiclient.NewCreateOrganizationAdaptivePolicyPolicyRequestDestinationGroup()) // CreateOrganizationAdaptivePolicyPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).CreateOrganizationAdaptivePolicyPolicyRequest(createOrganizationAdaptivePolicyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.CreateOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.CreateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyPolicyRequest** | [**CreateOrganizationAdaptivePolicyPolicyRequest**](CreateOrganizationAdaptivePolicyPolicyRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyPolicies200ResponseInner**](GetOrganizationAdaptivePolicyPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdaptivePolicyPolicy

> DeleteOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Delete an Adaptive Policy



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.DeleteOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


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


## GetNetworkPoliciesByClient

> []GetNetworkPoliciesByClient200ResponseInner GetNetworkPoliciesByClient(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()

Get policies for all clients with policies



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.GetNetworkPoliciesByClient(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetNetworkPoliciesByClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoliciesByClient`: []GetNetworkPoliciesByClient200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetNetworkPoliciesByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoliciesByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetNetworkPoliciesByClient200ResponseInner**](GetNetworkPoliciesByClient200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyPolicies

> []GetOrganizationAdaptivePolicyPolicies200ResponseInner GetOrganizationAdaptivePolicyPolicies(ctx, organizationId).Execute()

List adaptive policies in an organization



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
	resp, r, err := apiClient.PoliciesAPI.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetOrganizationAdaptivePolicyPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyPolicies`: []GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetOrganizationAdaptivePolicyPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationAdaptivePolicyPolicies200ResponseInner**](GetOrganizationAdaptivePolicyPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyPolicy

> GetOrganizationAdaptivePolicyPolicies200ResponseInner GetOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Return an adaptive policy



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationAdaptivePolicyPolicies200ResponseInner**](GetOrganizationAdaptivePolicyPolicies200ResponseInner.md)

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
	resp, r, err := apiClient.PoliciesAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: %v\n", resp)
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


## UpdateOrganizationAdaptivePolicyPolicy

> GetOrganizationAdaptivePolicyPolicies200ResponseInner UpdateOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyPolicyRequest(updateOrganizationAdaptivePolicyPolicyRequest).Execute()

Update an Adaptive Policy



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
	id := "id_example" // string | ID
	updateOrganizationAdaptivePolicyPolicyRequest := *openapiclient.NewUpdateOrganizationAdaptivePolicyPolicyRequest() // UpdateOrganizationAdaptivePolicyPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyPolicyRequest(updateOrganizationAdaptivePolicyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.UpdateOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.UpdateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyPolicyRequest** | [**UpdateOrganizationAdaptivePolicyPolicyRequest**](UpdateOrganizationAdaptivePolicyPolicyRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyPolicies200ResponseInner**](GetOrganizationAdaptivePolicyPolicies200ResponseInner.md)

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
	resp, r, err := apiClient.PoliciesAPI.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).UpdateOrganizationSmSentryPoliciesAssignmentsRequest(updateOrganizationSmSentryPoliciesAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.UpdateOrganizationSmSentryPoliciesAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSmSentryPoliciesAssignments`: UpdateOrganizationSmSentryPoliciesAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.UpdateOrganizationSmSentryPoliciesAssignments`: %v\n", resp)
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


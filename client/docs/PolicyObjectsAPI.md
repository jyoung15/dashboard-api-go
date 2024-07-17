# \PolicyObjectsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPolicyObject**](PolicyObjectsAPI.md#CreateOrganizationPolicyObject) | **Post** /organizations/{organizationId}/policyObjects | Creates a new Policy Object.
[**CreateOrganizationPolicyObjectsGroup**](PolicyObjectsAPI.md#CreateOrganizationPolicyObjectsGroup) | **Post** /organizations/{organizationId}/policyObjects/groups | Creates a new Policy Object Group.
[**DeleteOrganizationPolicyObject**](PolicyObjectsAPI.md#DeleteOrganizationPolicyObject) | **Delete** /organizations/{organizationId}/policyObjects/{policyObjectId} | Deletes a Policy Object.
[**DeleteOrganizationPolicyObjectsGroup**](PolicyObjectsAPI.md#DeleteOrganizationPolicyObjectsGroup) | **Delete** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Deletes a Policy Object Group.
[**GetOrganizationPolicyObject**](PolicyObjectsAPI.md#GetOrganizationPolicyObject) | **Get** /organizations/{organizationId}/policyObjects/{policyObjectId} | Shows details of a Policy Object.
[**GetOrganizationPolicyObjects**](PolicyObjectsAPI.md#GetOrganizationPolicyObjects) | **Get** /organizations/{organizationId}/policyObjects | Lists Policy Objects belonging to the organization.
[**GetOrganizationPolicyObjectsGroup**](PolicyObjectsAPI.md#GetOrganizationPolicyObjectsGroup) | **Get** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Shows details of a Policy Object Group.
[**GetOrganizationPolicyObjectsGroups**](PolicyObjectsAPI.md#GetOrganizationPolicyObjectsGroups) | **Get** /organizations/{organizationId}/policyObjects/groups | Lists Policy Object Groups belonging to the organization.
[**UpdateOrganizationPolicyObject**](PolicyObjectsAPI.md#UpdateOrganizationPolicyObject) | **Put** /organizations/{organizationId}/policyObjects/{policyObjectId} | Updates a Policy Object.
[**UpdateOrganizationPolicyObjectsGroup**](PolicyObjectsAPI.md#UpdateOrganizationPolicyObjectsGroup) | **Put** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Updates a Policy Object Group.



## CreateOrganizationPolicyObject

> GetOrganizationPolicyObjects200Response CreateOrganizationPolicyObject(ctx, organizationId).CreateOrganizationPolicyObjectRequest(createOrganizationPolicyObjectRequest).Execute()

Creates a new Policy Object.



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
	createOrganizationPolicyObjectRequest := *openapiclient.NewCreateOrganizationPolicyObjectRequest("Name_example", "Category_example", "Type_example") // CreateOrganizationPolicyObjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.CreateOrganizationPolicyObject(context.Background(), organizationId).CreateOrganizationPolicyObjectRequest(createOrganizationPolicyObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.CreateOrganizationPolicyObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationPolicyObject`: GetOrganizationPolicyObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.CreateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObjectRequest** | [**CreateOrganizationPolicyObjectRequest**](CreateOrganizationPolicyObjectRequest.md) |  | 

### Return type

[**GetOrganizationPolicyObjects200Response**](GetOrganizationPolicyObjects200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPolicyObjectsGroup

> GetOrganizationPolicyObjectsGroups200Response CreateOrganizationPolicyObjectsGroup(ctx, organizationId).CreateOrganizationPolicyObjectsGroupRequest(createOrganizationPolicyObjectsGroupRequest).Execute()

Creates a new Policy Object Group.



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
	createOrganizationPolicyObjectsGroupRequest := *openapiclient.NewCreateOrganizationPolicyObjectsGroupRequest("Name_example") // CreateOrganizationPolicyObjectsGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.CreateOrganizationPolicyObjectsGroup(context.Background(), organizationId).CreateOrganizationPolicyObjectsGroupRequest(createOrganizationPolicyObjectsGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.CreateOrganizationPolicyObjectsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationPolicyObjectsGroup`: GetOrganizationPolicyObjectsGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.CreateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObjectsGroupRequest** | [**CreateOrganizationPolicyObjectsGroupRequest**](CreateOrganizationPolicyObjectsGroupRequest.md) |  | 

### Return type

[**GetOrganizationPolicyObjectsGroups200Response**](GetOrganizationPolicyObjectsGroups200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPolicyObject

> DeleteOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Deletes a Policy Object.



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
	policyObjectId := "policyObjectId_example" // string | Policy object ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyObjectsAPI.DeleteOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.DeleteOrganizationPolicyObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectRequest struct via the builder pattern


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


## DeleteOrganizationPolicyObjectsGroup

> DeleteOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Deletes a Policy Object Group.



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
	policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyObjectsAPI.DeleteOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.DeleteOrganizationPolicyObjectsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectsGroupRequest struct via the builder pattern


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


## GetOrganizationPolicyObject

> GetOrganizationPolicyObjects200Response GetOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Shows details of a Policy Object.



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
	policyObjectId := "policyObjectId_example" // string | Policy object ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.GetOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.GetOrganizationPolicyObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationPolicyObject`: GetOrganizationPolicyObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.GetOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationPolicyObjects200Response**](GetOrganizationPolicyObjects200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjects

> GetOrganizationPolicyObjects200Response GetOrganizationPolicyObjects(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Objects belonging to the organization.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.GetOrganizationPolicyObjects(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.GetOrganizationPolicyObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationPolicyObjects`: GetOrganizationPolicyObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.GetOrganizationPolicyObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationPolicyObjects200Response**](GetOrganizationPolicyObjects200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroup

> GetOrganizationPolicyObjectsGroups200Response GetOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Shows details of a Policy Object Group.



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
	policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.GetOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.GetOrganizationPolicyObjectsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationPolicyObjectsGroup`: GetOrganizationPolicyObjectsGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.GetOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationPolicyObjectsGroups200Response**](GetOrganizationPolicyObjectsGroups200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroups

> GetOrganizationPolicyObjectsGroups200Response GetOrganizationPolicyObjectsGroups(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Object Groups belonging to the organization.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.GetOrganizationPolicyObjectsGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.GetOrganizationPolicyObjectsGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationPolicyObjectsGroups`: GetOrganizationPolicyObjectsGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.GetOrganizationPolicyObjectsGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationPolicyObjectsGroups200Response**](GetOrganizationPolicyObjectsGroups200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObject

> GetOrganizationPolicyObjects200Response UpdateOrganizationPolicyObject(ctx, organizationId, policyObjectId).UpdateOrganizationPolicyObjectRequest(updateOrganizationPolicyObjectRequest).Execute()

Updates a Policy Object.



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
	policyObjectId := "policyObjectId_example" // string | Policy object ID
	updateOrganizationPolicyObjectRequest := *openapiclient.NewUpdateOrganizationPolicyObjectRequest() // UpdateOrganizationPolicyObjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.UpdateOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).UpdateOrganizationPolicyObjectRequest(updateOrganizationPolicyObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.UpdateOrganizationPolicyObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationPolicyObject`: GetOrganizationPolicyObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.UpdateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObjectRequest** | [**UpdateOrganizationPolicyObjectRequest**](UpdateOrganizationPolicyObjectRequest.md) |  | 

### Return type

[**GetOrganizationPolicyObjects200Response**](GetOrganizationPolicyObjects200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObjectsGroup

> GetOrganizationPolicyObjectsGroups200Response UpdateOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroupRequest(updateOrganizationPolicyObjectsGroupRequest).Execute()

Updates a Policy Object Group.



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
	policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID
	updateOrganizationPolicyObjectsGroupRequest := *openapiclient.NewUpdateOrganizationPolicyObjectsGroupRequest() // UpdateOrganizationPolicyObjectsGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyObjectsAPI.UpdateOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroupRequest(updateOrganizationPolicyObjectsGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsAPI.UpdateOrganizationPolicyObjectsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationPolicyObjectsGroup`: GetOrganizationPolicyObjectsGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsAPI.UpdateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObjectsGroupRequest** | [**UpdateOrganizationPolicyObjectsGroupRequest**](UpdateOrganizationPolicyObjectsGroupRequest.md) |  | 

### Return type

[**GetOrganizationPolicyObjectsGroups200Response**](GetOrganizationPolicyObjectsGroups200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


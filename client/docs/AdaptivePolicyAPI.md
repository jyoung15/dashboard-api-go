# \AdaptivePolicyAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdaptivePolicyAcl**](AdaptivePolicyAPI.md#CreateOrganizationAdaptivePolicyAcl) | **Post** /organizations/{organizationId}/adaptivePolicy/acls | Creates new adaptive policy ACL
[**CreateOrganizationAdaptivePolicyGroup**](AdaptivePolicyAPI.md#CreateOrganizationAdaptivePolicyGroup) | **Post** /organizations/{organizationId}/adaptivePolicy/groups | Creates a new adaptive policy group
[**CreateOrganizationAdaptivePolicyPolicy**](AdaptivePolicyAPI.md#CreateOrganizationAdaptivePolicyPolicy) | **Post** /organizations/{organizationId}/adaptivePolicy/policies | Add an Adaptive Policy
[**DeleteOrganizationAdaptivePolicyAcl**](AdaptivePolicyAPI.md#DeleteOrganizationAdaptivePolicyAcl) | **Delete** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Deletes the specified adaptive policy ACL
[**DeleteOrganizationAdaptivePolicyGroup**](AdaptivePolicyAPI.md#DeleteOrganizationAdaptivePolicyGroup) | **Delete** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Deletes the specified adaptive policy group and any associated policies and references
[**DeleteOrganizationAdaptivePolicyPolicy**](AdaptivePolicyAPI.md#DeleteOrganizationAdaptivePolicyPolicy) | **Delete** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Delete an Adaptive Policy
[**GetOrganizationAdaptivePolicyAcl**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyAcl) | **Get** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Returns the adaptive policy ACL information
[**GetOrganizationAdaptivePolicyAcls**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyAcls) | **Get** /organizations/{organizationId}/adaptivePolicy/acls | List adaptive policy ACLs in a organization
[**GetOrganizationAdaptivePolicyGroup**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyGroup) | **Get** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Returns an adaptive policy group
[**GetOrganizationAdaptivePolicyGroups**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyGroups) | **Get** /organizations/{organizationId}/adaptivePolicy/groups | List adaptive policy groups in a organization
[**GetOrganizationAdaptivePolicyOverview**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyOverview) | **Get** /organizations/{organizationId}/adaptivePolicy/overview | Returns adaptive policy aggregate statistics for an organization
[**GetOrganizationAdaptivePolicyPolicies**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyPolicies) | **Get** /organizations/{organizationId}/adaptivePolicy/policies | List adaptive policies in an organization
[**GetOrganizationAdaptivePolicyPolicy**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicyPolicy) | **Get** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Return an adaptive policy
[**GetOrganizationAdaptivePolicySettings**](AdaptivePolicyAPI.md#GetOrganizationAdaptivePolicySettings) | **Get** /organizations/{organizationId}/adaptivePolicy/settings | Returns global adaptive policy settings in an organization
[**UpdateOrganizationAdaptivePolicyAcl**](AdaptivePolicyAPI.md#UpdateOrganizationAdaptivePolicyAcl) | **Put** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Updates an adaptive policy ACL
[**UpdateOrganizationAdaptivePolicyGroup**](AdaptivePolicyAPI.md#UpdateOrganizationAdaptivePolicyGroup) | **Put** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Updates an adaptive policy group
[**UpdateOrganizationAdaptivePolicyPolicy**](AdaptivePolicyAPI.md#UpdateOrganizationAdaptivePolicyPolicy) | **Put** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Update an Adaptive Policy
[**UpdateOrganizationAdaptivePolicySettings**](AdaptivePolicyAPI.md#UpdateOrganizationAdaptivePolicySettings) | **Put** /organizations/{organizationId}/adaptivePolicy/settings | Update global adaptive policy settings



## CreateOrganizationAdaptivePolicyAcl

> GetOrganizationAdaptivePolicyAcls200ResponseInner CreateOrganizationAdaptivePolicyAcl(ctx, organizationId).CreateOrganizationAdaptivePolicyAclRequest(createOrganizationAdaptivePolicyAclRequest).Execute()

Creates new adaptive policy ACL



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
	createOrganizationAdaptivePolicyAclRequest := *openapiclient.NewCreateOrganizationAdaptivePolicyAclRequest("Name_example", []openapiclient.CreateOrganizationAdaptivePolicyAclRequestRulesInner{*openapiclient.NewCreateOrganizationAdaptivePolicyAclRequestRulesInner("Policy_example", "Protocol_example")}, "IpVersion_example") // CreateOrganizationAdaptivePolicyAclRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyAcl(context.Background(), organizationId).CreateOrganizationAdaptivePolicyAclRequest(createOrganizationAdaptivePolicyAclRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAdaptivePolicyAcl`: GetOrganizationAdaptivePolicyAcls200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyAclRequest** | [**CreateOrganizationAdaptivePolicyAclRequest**](CreateOrganizationAdaptivePolicyAclRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyAcls200ResponseInner**](GetOrganizationAdaptivePolicyAcls200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAdaptivePolicyGroup

> GetOrganizationAdaptivePolicyGroups200ResponseInner CreateOrganizationAdaptivePolicyGroup(ctx, organizationId).CreateOrganizationAdaptivePolicyGroupRequest(createOrganizationAdaptivePolicyGroupRequest).Execute()

Creates a new adaptive policy group



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
	createOrganizationAdaptivePolicyGroupRequest := *openapiclient.NewCreateOrganizationAdaptivePolicyGroupRequest("Name_example", int32(123)) // CreateOrganizationAdaptivePolicyGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyGroup(context.Background(), organizationId).CreateOrganizationAdaptivePolicyGroupRequest(createOrganizationAdaptivePolicyGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAdaptivePolicyGroup`: GetOrganizationAdaptivePolicyGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyGroupRequest** | [**CreateOrganizationAdaptivePolicyGroupRequest**](CreateOrganizationAdaptivePolicyGroupRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyGroups200ResponseInner**](GetOrganizationAdaptivePolicyGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).CreateOrganizationAdaptivePolicyPolicyRequest(createOrganizationAdaptivePolicyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.CreateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
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


## DeleteOrganizationAdaptivePolicyAcl

> DeleteOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Deletes the specified adaptive policy ACL



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
	aclId := "aclId_example" // string | Acl ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyAclRequest struct via the builder pattern


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


## DeleteOrganizationAdaptivePolicyGroup

> DeleteOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Deletes the specified adaptive policy group and any associated policies and references



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
	r, err := apiClient.AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


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
	r, err := apiClient.AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.DeleteOrganizationAdaptivePolicyPolicy``: %v\n", err)
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


## GetOrganizationAdaptivePolicyAcl

> GetOrganizationAdaptivePolicyAcls200ResponseInner GetOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Returns the adaptive policy ACL information



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
	aclId := "aclId_example" // string | Acl ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyAcl`: GetOrganizationAdaptivePolicyAcls200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationAdaptivePolicyAcls200ResponseInner**](GetOrganizationAdaptivePolicyAcls200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyAcls

> []GetOrganizationAdaptivePolicyAcls200ResponseInner GetOrganizationAdaptivePolicyAcls(ctx, organizationId).Execute()

List adaptive policy ACLs in a organization



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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcls(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyAcls`: []GetOrganizationAdaptivePolicyAcls200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationAdaptivePolicyAcls200ResponseInner**](GetOrganizationAdaptivePolicyAcls200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyGroup

> GetOrganizationAdaptivePolicyGroups200ResponseInner GetOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Returns an adaptive policy group



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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyGroup`: GetOrganizationAdaptivePolicyGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationAdaptivePolicyGroups200ResponseInner**](GetOrganizationAdaptivePolicyGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyGroups

> []GetOrganizationAdaptivePolicyGroups200ResponseInner GetOrganizationAdaptivePolicyGroups(ctx, organizationId).Execute()

List adaptive policy groups in a organization



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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroups(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyGroups`: []GetOrganizationAdaptivePolicyGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationAdaptivePolicyGroups200ResponseInner**](GetOrganizationAdaptivePolicyGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyOverview

> GetOrganizationAdaptivePolicyOverview200Response GetOrganizationAdaptivePolicyOverview(ctx, organizationId).Execute()

Returns adaptive policy aggregate statistics for an organization



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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyOverview(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyOverview`: GetOrganizationAdaptivePolicyOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationAdaptivePolicyOverview200Response**](GetOrganizationAdaptivePolicyOverview200Response.md)

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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyPolicies`: []GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicies`: %v\n", resp)
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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicyPolicy`: %v\n", resp)
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


## GetOrganizationAdaptivePolicySettings

> GetOrganizationAdaptivePolicySettings200Response GetOrganizationAdaptivePolicySettings(ctx, organizationId).Execute()

Returns global adaptive policy settings in an organization



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
	resp, r, err := apiClient.AdaptivePolicyAPI.GetOrganizationAdaptivePolicySettings(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.GetOrganizationAdaptivePolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAdaptivePolicySettings`: GetOrganizationAdaptivePolicySettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.GetOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationAdaptivePolicySettings200Response**](GetOrganizationAdaptivePolicySettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyAcl

> GetOrganizationAdaptivePolicyAcls200ResponseInner UpdateOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).UpdateOrganizationAdaptivePolicyAclRequest(updateOrganizationAdaptivePolicyAclRequest).Execute()

Updates an adaptive policy ACL



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
	aclId := "aclId_example" // string | Acl ID
	updateOrganizationAdaptivePolicyAclRequest := *openapiclient.NewUpdateOrganizationAdaptivePolicyAclRequest() // UpdateOrganizationAdaptivePolicyAclRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).UpdateOrganizationAdaptivePolicyAclRequest(updateOrganizationAdaptivePolicyAclRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAdaptivePolicyAcl`: GetOrganizationAdaptivePolicyAcls200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyAclRequest** | [**UpdateOrganizationAdaptivePolicyAclRequest**](UpdateOrganizationAdaptivePolicyAclRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyAcls200ResponseInner**](GetOrganizationAdaptivePolicyAcls200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyGroup

> GetOrganizationAdaptivePolicyGroups200ResponseInner UpdateOrganizationAdaptivePolicyGroup(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyGroupRequest(updateOrganizationAdaptivePolicyGroupRequest).Execute()

Updates an adaptive policy group



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
	updateOrganizationAdaptivePolicyGroupRequest := *openapiclient.NewUpdateOrganizationAdaptivePolicyGroupRequest() // UpdateOrganizationAdaptivePolicyGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyGroupRequest(updateOrganizationAdaptivePolicyGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAdaptivePolicyGroup`: GetOrganizationAdaptivePolicyGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyGroupRequest** | [**UpdateOrganizationAdaptivePolicyGroupRequest**](UpdateOrganizationAdaptivePolicyGroupRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicyGroups200ResponseInner**](GetOrganizationAdaptivePolicyGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyPolicyRequest(updateOrganizationAdaptivePolicyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAdaptivePolicyPolicy`: GetOrganizationAdaptivePolicyPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
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


## UpdateOrganizationAdaptivePolicySettings

> GetOrganizationAdaptivePolicySettings200Response UpdateOrganizationAdaptivePolicySettings(ctx, organizationId).UpdateOrganizationAdaptivePolicySettingsRequest(updateOrganizationAdaptivePolicySettingsRequest).Execute()

Update global adaptive policy settings



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
	updateOrganizationAdaptivePolicySettingsRequest := *openapiclient.NewUpdateOrganizationAdaptivePolicySettingsRequest() // UpdateOrganizationAdaptivePolicySettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicySettings(context.Background(), organizationId).UpdateOrganizationAdaptivePolicySettingsRequest(updateOrganizationAdaptivePolicySettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationAdaptivePolicySettings`: GetOrganizationAdaptivePolicySettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyAPI.UpdateOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationAdaptivePolicySettingsRequest** | [**UpdateOrganizationAdaptivePolicySettingsRequest**](UpdateOrganizationAdaptivePolicySettingsRequest.md) |  | 

### Return type

[**GetOrganizationAdaptivePolicySettings200Response**](GetOrganizationAdaptivePolicySettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


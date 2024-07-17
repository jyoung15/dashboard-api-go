# \ConfigTemplatesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationConfigTemplate**](ConfigTemplatesAPI.md#CreateOrganizationConfigTemplate) | **Post** /organizations/{organizationId}/configTemplates | Create a new configuration template
[**DeleteOrganizationConfigTemplate**](ConfigTemplatesAPI.md#DeleteOrganizationConfigTemplate) | **Delete** /organizations/{organizationId}/configTemplates/{configTemplateId} | Remove a configuration template
[**GetOrganizationConfigTemplate**](ConfigTemplatesAPI.md#GetOrganizationConfigTemplate) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId} | Return a single configuration template
[**GetOrganizationConfigTemplateSwitchProfilePort**](ConfigTemplatesAPI.md#GetOrganizationConfigTemplateSwitchProfilePort) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Return a switch template port
[**GetOrganizationConfigTemplateSwitchProfilePorts**](ConfigTemplatesAPI.md#GetOrganizationConfigTemplateSwitchProfilePorts) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports | Return all the ports of a switch template
[**GetOrganizationConfigTemplateSwitchProfiles**](ConfigTemplatesAPI.md#GetOrganizationConfigTemplateSwitchProfiles) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles | List the switch templates for your switch template configuration
[**GetOrganizationConfigTemplates**](ConfigTemplatesAPI.md#GetOrganizationConfigTemplates) | **Get** /organizations/{organizationId}/configTemplates | List the configuration templates for this organization
[**UpdateOrganizationConfigTemplate**](ConfigTemplatesAPI.md#UpdateOrganizationConfigTemplate) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId} | Update a configuration template
[**UpdateOrganizationConfigTemplateSwitchProfilePort**](ConfigTemplatesAPI.md#UpdateOrganizationConfigTemplateSwitchProfilePort) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Update a switch template port



## CreateOrganizationConfigTemplate

> GetOrganizationConfigTemplates200ResponseInner CreateOrganizationConfigTemplate(ctx, organizationId).CreateOrganizationConfigTemplateRequest(createOrganizationConfigTemplateRequest).Execute()

Create a new configuration template



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
	createOrganizationConfigTemplateRequest := *openapiclient.NewCreateOrganizationConfigTemplateRequest("Name_example") // CreateOrganizationConfigTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.CreateOrganizationConfigTemplate(context.Background(), organizationId).CreateOrganizationConfigTemplateRequest(createOrganizationConfigTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.CreateOrganizationConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationConfigTemplate`: GetOrganizationConfigTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.CreateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationConfigTemplateRequest** | [**CreateOrganizationConfigTemplateRequest**](CreateOrganizationConfigTemplateRequest.md) |  | 

### Return type

[**GetOrganizationConfigTemplates200ResponseInner**](GetOrganizationConfigTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationConfigTemplate

> DeleteOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Remove a configuration template



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
	configTemplateId := "configTemplateId_example" // string | Config template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigTemplatesAPI.DeleteOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.DeleteOrganizationConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationConfigTemplateRequest struct via the builder pattern


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


## GetOrganizationConfigTemplate

> GetOrganizationConfigTemplates200ResponseInner GetOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Return a single configuration template



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
	configTemplateId := "configTemplateId_example" // string | Config template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.GetOrganizationConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplate`: GetOrganizationConfigTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.GetOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationConfigTemplates200ResponseInner**](GetOrganizationConfigTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).Execute()

Return a switch template port



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
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePorts

> []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePorts(ctx, organizationId, configTemplateId, profileId).Execute()

Return all the ports of a switch template



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
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePorts`: []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfiles

> []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner GetOrganizationConfigTemplateSwitchProfiles(ctx, organizationId, configTemplateId).Execute()

List the switch templates for your switch template configuration



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
	configTemplateId := "configTemplateId_example" // string | Config template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfiles`: []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetOrganizationConfigTemplateSwitchProfiles200ResponseInner**](GetOrganizationConfigTemplateSwitchProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplates

> []GetOrganizationConfigTemplates200ResponseInner GetOrganizationConfigTemplates(ctx, organizationId).Execute()

List the configuration templates for this organization



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
	resp, r, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplates(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.GetOrganizationConfigTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplates`: []GetOrganizationConfigTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.GetOrganizationConfigTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationConfigTemplates200ResponseInner**](GetOrganizationConfigTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplate

> GetOrganizationConfigTemplates200ResponseInner UpdateOrganizationConfigTemplate(ctx, organizationId, configTemplateId).UpdateOrganizationConfigTemplateRequest(updateOrganizationConfigTemplateRequest).Execute()

Update a configuration template



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
	configTemplateId := "configTemplateId_example" // string | Config template ID
	updateOrganizationConfigTemplateRequest := *openapiclient.NewUpdateOrganizationConfigTemplateRequest() // UpdateOrganizationConfigTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.UpdateOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).UpdateOrganizationConfigTemplateRequest(updateOrganizationConfigTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.UpdateOrganizationConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationConfigTemplate`: GetOrganizationConfigTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.UpdateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationConfigTemplateRequest** | [**UpdateOrganizationConfigTemplateRequest**](UpdateOrganizationConfigTemplateRequest.md) |  | 

### Return type

[**GetOrganizationConfigTemplates200ResponseInner**](GetOrganizationConfigTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner UpdateOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()

Update a switch template port



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
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID
	portId := "portId_example" // string | Port ID
	updateOrganizationConfigTemplateSwitchProfilePortRequest := *openapiclient.NewUpdateOrganizationConfigTemplateSwitchProfilePortRequest() // UpdateOrganizationConfigTemplateSwitchProfilePortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigTemplatesAPI.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesAPI.UpdateOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesAPI.UpdateOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateOrganizationConfigTemplateSwitchProfilePortRequest** | [**UpdateOrganizationConfigTemplateSwitchProfilePortRequest**](UpdateOrganizationConfigTemplateSwitchProfilePortRequest.md) |  | 

### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


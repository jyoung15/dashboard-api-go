# \LicensesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationLicensesSeats**](LicensesAPI.md#AssignOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/assignSeats | Assign SM seats to a network
[**GetOrganizationLicense**](LicensesAPI.md#GetOrganizationLicense) | **Get** /organizations/{organizationId}/licenses/{licenseId} | Display a license
[**GetOrganizationLicenses**](LicensesAPI.md#GetOrganizationLicenses) | **Get** /organizations/{organizationId}/licenses | List the licenses for an organization
[**GetOrganizationLicensesOverview**](LicensesAPI.md#GetOrganizationLicensesOverview) | **Get** /organizations/{organizationId}/licenses/overview | Return an overview of the license state for an organization
[**GetOrganizationLicensingCotermLicenses**](LicensesAPI.md#GetOrganizationLicensingCotermLicenses) | **Get** /organizations/{organizationId}/licensing/coterm/licenses | List the licenses in a coterm organization
[**MoveOrganizationLicenses**](LicensesAPI.md#MoveOrganizationLicenses) | **Post** /organizations/{organizationId}/licenses/move | Move licenses to another organization
[**MoveOrganizationLicensesSeats**](LicensesAPI.md#MoveOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/moveSeats | Move SM seats to another organization
[**MoveOrganizationLicensingCotermLicenses**](LicensesAPI.md#MoveOrganizationLicensingCotermLicenses) | **Post** /organizations/{organizationId}/licensing/coterm/licenses/move | Moves a license to a different organization (coterm only)
[**RenewOrganizationLicensesSeats**](LicensesAPI.md#RenewOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/renewSeats | Renew SM seats of a license
[**UpdateOrganizationLicense**](LicensesAPI.md#UpdateOrganizationLicense) | **Put** /organizations/{organizationId}/licenses/{licenseId} | Update a license



## AssignOrganizationLicensesSeats

> AssignOrganizationLicensesSeats200Response AssignOrganizationLicensesSeats(ctx, organizationId).AssignOrganizationLicensesSeatsRequest(assignOrganizationLicensesSeatsRequest).Execute()

Assign SM seats to a network



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
	assignOrganizationLicensesSeatsRequest := *openapiclient.NewAssignOrganizationLicensesSeatsRequest("LicenseId_example", "NetworkId_example", int32(123)) // AssignOrganizationLicensesSeatsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.AssignOrganizationLicensesSeats(context.Background(), organizationId).AssignOrganizationLicensesSeatsRequest(assignOrganizationLicensesSeatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.AssignOrganizationLicensesSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignOrganizationLicensesSeats`: AssignOrganizationLicensesSeats200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.AssignOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationLicensesSeatsRequest** | [**AssignOrganizationLicensesSeatsRequest**](AssignOrganizationLicensesSeatsRequest.md) |  | 

### Return type

[**AssignOrganizationLicensesSeats200Response**](AssignOrganizationLicensesSeats200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicense

> GetOrganizationLicenses200ResponseInner GetOrganizationLicense(ctx, organizationId, licenseId).Execute()

Display a license



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
	licenseId := "licenseId_example" // string | License ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.GetOrganizationLicense(context.Background(), organizationId, licenseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetOrganizationLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLicense`: GetOrganizationLicenses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationLicenses200ResponseInner**](GetOrganizationLicenses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicenses

> []GetOrganizationLicenses200ResponseInner GetOrganizationLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()

List the licenses for an organization



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
	deviceSerial := "deviceSerial_example" // string | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. (optional)
	networkId := "networkId_example" // string | Filter the licenses to those assigned in a particular network (optional)
	state := "state_example" // string | Filter the licenses to those in a particular state. Can be one of 'active', 'expired', 'expiring', 'recentlyQueued', 'unused' or 'unusedActive' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.GetOrganizationLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetOrganizationLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLicenses`: []GetOrganizationLicenses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **deviceSerial** | **string** | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. | 
 **networkId** | **string** | Filter the licenses to those assigned in a particular network | 
 **state** | **string** | Filter the licenses to those in a particular state. Can be one of &#39;active&#39;, &#39;expired&#39;, &#39;expiring&#39;, &#39;recentlyQueued&#39;, &#39;unused&#39; or &#39;unusedActive&#39; | 

### Return type

[**[]GetOrganizationLicenses200ResponseInner**](GetOrganizationLicenses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensesOverview

> GetOrganizationLicensesOverview200Response GetOrganizationLicensesOverview(ctx, organizationId).Execute()

Return an overview of the license state for an organization



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
	resp, r, err := apiClient.LicensesAPI.GetOrganizationLicensesOverview(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetOrganizationLicensesOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLicensesOverview`: GetOrganizationLicensesOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetOrganizationLicensesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationLicensesOverview200Response**](GetOrganizationLicensesOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensingCotermLicenses

> []GetOrganizationLicensingCotermLicenses200ResponseInner GetOrganizationLicensingCotermLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()

List the licenses in a coterm organization



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
	invalidated := true // bool | Filter for licenses that are invalidated (optional)
	expired := true // bool | Filter for licenses that are expired (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.GetOrganizationLicensingCotermLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.GetOrganizationLicensingCotermLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLicensingCotermLicenses`: []GetOrganizationLicensingCotermLicenses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.GetOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **invalidated** | **bool** | Filter for licenses that are invalidated | 
 **expired** | **bool** | Filter for licenses that are expired | 

### Return type

[**[]GetOrganizationLicensingCotermLicenses200ResponseInner**](GetOrganizationLicensingCotermLicenses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicenses

> MoveOrganizationLicenses200Response MoveOrganizationLicenses(ctx, organizationId).MoveOrganizationLicensesRequest(moveOrganizationLicensesRequest).Execute()

Move licenses to another organization



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
	moveOrganizationLicensesRequest := *openapiclient.NewMoveOrganizationLicensesRequest("DestOrganizationId_example", []string{"LicenseIds_example"}) // MoveOrganizationLicensesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.MoveOrganizationLicenses(context.Background(), organizationId).MoveOrganizationLicensesRequest(moveOrganizationLicensesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.MoveOrganizationLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveOrganizationLicenses`: MoveOrganizationLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.MoveOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensesRequest** | [**MoveOrganizationLicensesRequest**](MoveOrganizationLicensesRequest.md) |  | 

### Return type

[**MoveOrganizationLicenses200Response**](MoveOrganizationLicenses200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensesSeats

> MoveOrganizationLicensesSeats200Response MoveOrganizationLicensesSeats(ctx, organizationId).MoveOrganizationLicensesSeatsRequest(moveOrganizationLicensesSeatsRequest).Execute()

Move SM seats to another organization



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
	moveOrganizationLicensesSeatsRequest := *openapiclient.NewMoveOrganizationLicensesSeatsRequest("DestOrganizationId_example", "LicenseId_example", int32(123)) // MoveOrganizationLicensesSeatsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.MoveOrganizationLicensesSeats(context.Background(), organizationId).MoveOrganizationLicensesSeatsRequest(moveOrganizationLicensesSeatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.MoveOrganizationLicensesSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveOrganizationLicensesSeats`: MoveOrganizationLicensesSeats200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.MoveOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensesSeatsRequest** | [**MoveOrganizationLicensesSeatsRequest**](MoveOrganizationLicensesSeatsRequest.md) |  | 

### Return type

[**MoveOrganizationLicensesSeats200Response**](MoveOrganizationLicensesSeats200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensingCotermLicenses

> MoveOrganizationLicensingCotermLicenses200Response MoveOrganizationLicensingCotermLicenses(ctx, organizationId).MoveOrganizationLicensingCotermLicensesRequest(moveOrganizationLicensingCotermLicensesRequest).Execute()

Moves a license to a different organization (coterm only)



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
	moveOrganizationLicensingCotermLicensesRequest := *openapiclient.NewMoveOrganizationLicensingCotermLicensesRequest(*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestDestination(), []openapiclient.MoveOrganizationLicensingCotermLicensesRequestLicensesInner{*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestLicensesInner("Key_example", []openapiclient.MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner("Model_example", int32(123))})}) // MoveOrganizationLicensingCotermLicensesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.MoveOrganizationLicensingCotermLicenses(context.Background(), organizationId).MoveOrganizationLicensingCotermLicensesRequest(moveOrganizationLicensingCotermLicensesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.MoveOrganizationLicensingCotermLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveOrganizationLicensingCotermLicenses`: MoveOrganizationLicensingCotermLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.MoveOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensingCotermLicensesRequest** | [**MoveOrganizationLicensingCotermLicensesRequest**](MoveOrganizationLicensingCotermLicensesRequest.md) |  | 

### Return type

[**MoveOrganizationLicensingCotermLicenses200Response**](MoveOrganizationLicensingCotermLicenses200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewOrganizationLicensesSeats

> AssignOrganizationLicensesSeats200Response RenewOrganizationLicensesSeats(ctx, organizationId).RenewOrganizationLicensesSeatsRequest(renewOrganizationLicensesSeatsRequest).Execute()

Renew SM seats of a license



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
	renewOrganizationLicensesSeatsRequest := *openapiclient.NewRenewOrganizationLicensesSeatsRequest("LicenseIdToRenew_example", "UnusedLicenseId_example") // RenewOrganizationLicensesSeatsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.RenewOrganizationLicensesSeats(context.Background(), organizationId).RenewOrganizationLicensesSeatsRequest(renewOrganizationLicensesSeatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.RenewOrganizationLicensesSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewOrganizationLicensesSeats`: AssignOrganizationLicensesSeats200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.RenewOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewOrganizationLicensesSeatsRequest** | [**RenewOrganizationLicensesSeatsRequest**](RenewOrganizationLicensesSeatsRequest.md) |  | 

### Return type

[**AssignOrganizationLicensesSeats200Response**](AssignOrganizationLicensesSeats200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLicense

> GetOrganizationLicenses200ResponseInner UpdateOrganizationLicense(ctx, organizationId, licenseId).UpdateOrganizationLicenseRequest(updateOrganizationLicenseRequest).Execute()

Update a license



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
	licenseId := "licenseId_example" // string | License ID
	updateOrganizationLicenseRequest := *openapiclient.NewUpdateOrganizationLicenseRequest() // UpdateOrganizationLicenseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.UpdateOrganizationLicense(context.Background(), organizationId, licenseId).UpdateOrganizationLicenseRequest(updateOrganizationLicenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.UpdateOrganizationLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationLicense`: GetOrganizationLicenses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.UpdateOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationLicenseRequest** | [**UpdateOrganizationLicenseRequest**](UpdateOrganizationLicenseRequest.md) |  | 

### Return type

[**GetOrganizationLicenses200ResponseInner**](GetOrganizationLicenses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


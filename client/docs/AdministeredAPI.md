# \AdministeredAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministeredIdentitiesMe**](AdministeredAPI.md#GetAdministeredIdentitiesMe) | **Get** /administered/identities/me | Returns the identity of the current user.



## GetAdministeredIdentitiesMe

> GetAdministeredIdentitiesMe200Response GetAdministeredIdentitiesMe(ctx).Execute()

Returns the identity of the current user.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministeredAPI.GetAdministeredIdentitiesMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministeredAPI.GetAdministeredIdentitiesMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdministeredIdentitiesMe`: GetAdministeredIdentitiesMe200Response
	fmt.Fprintf(os.Stdout, "Response from `AdministeredAPI.GetAdministeredIdentitiesMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredIdentitiesMeRequest struct via the builder pattern


### Return type

[**GetAdministeredIdentitiesMe200Response**](GetAdministeredIdentitiesMe200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


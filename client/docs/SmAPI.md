# \SmAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckinNetworkSmDevices**](SmAPI.md#CheckinNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/checkin | Force check-in a set of devices
[**CreateNetworkSmBypassActivationLockAttempt**](SmAPI.md#CreateNetworkSmBypassActivationLockAttempt) | **Post** /networks/{networkId}/sm/bypassActivationLockAttempts | Bypass activation lock attempt
[**CreateNetworkSmTargetGroup**](SmAPI.md#CreateNetworkSmTargetGroup) | **Post** /networks/{networkId}/sm/targetGroups | Add a target group
[**CreateOrganizationSmAdminsRole**](SmAPI.md#CreateOrganizationSmAdminsRole) | **Post** /organizations/{organizationId}/sm/admins/roles | Create a Limited Access Role
[**DeleteNetworkSmTargetGroup**](SmAPI.md#DeleteNetworkSmTargetGroup) | **Delete** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Delete a target group from a network
[**DeleteNetworkSmUserAccessDevice**](SmAPI.md#DeleteNetworkSmUserAccessDevice) | **Delete** /networks/{networkId}/sm/userAccessDevices/{userAccessDeviceId} | Delete a User Access Device
[**DeleteOrganizationSmAdminsRole**](SmAPI.md#DeleteOrganizationSmAdminsRole) | **Delete** /organizations/{organizationId}/sm/admins/roles/{roleId} | Delete a Limited Access Role
[**GetNetworkSmBypassActivationLockAttempt**](SmAPI.md#GetNetworkSmBypassActivationLockAttempt) | **Get** /networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId} | Bypass activation lock attempt status
[**GetNetworkSmDeviceCellularUsageHistory**](SmAPI.md#GetNetworkSmDeviceCellularUsageHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory | Return the client&#39;s daily cellular data usage history
[**GetNetworkSmDeviceCerts**](SmAPI.md#GetNetworkSmDeviceCerts) | **Get** /networks/{networkId}/sm/devices/{deviceId}/certs | List the certs on a device
[**GetNetworkSmDeviceConnectivity**](SmAPI.md#GetNetworkSmDeviceConnectivity) | **Get** /networks/{networkId}/sm/devices/{deviceId}/connectivity | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
[**GetNetworkSmDeviceDesktopLogs**](SmAPI.md#GetNetworkSmDeviceDesktopLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/desktopLogs | Return historical records of various Systems Manager network connection details for desktop devices.
[**GetNetworkSmDeviceDeviceCommandLogs**](SmAPI.md#GetNetworkSmDeviceDeviceCommandLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceCommandLogs | Return historical records of commands sent to Systems Manager devices
[**GetNetworkSmDeviceDeviceProfiles**](SmAPI.md#GetNetworkSmDeviceDeviceProfiles) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles | Get the installed profiles associated with a device
[**GetNetworkSmDeviceNetworkAdapters**](SmAPI.md#GetNetworkSmDeviceNetworkAdapters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/networkAdapters | List the network adapters of a device
[**GetNetworkSmDevicePerformanceHistory**](SmAPI.md#GetNetworkSmDevicePerformanceHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/performanceHistory | Return historical records of various Systems Manager client metrics for desktop devices.
[**GetNetworkSmDeviceRestrictions**](SmAPI.md#GetNetworkSmDeviceRestrictions) | **Get** /networks/{networkId}/sm/devices/{deviceId}/restrictions | List the restrictions on a device
[**GetNetworkSmDeviceSecurityCenters**](SmAPI.md#GetNetworkSmDeviceSecurityCenters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/securityCenters | List the security centers on a device
[**GetNetworkSmDeviceSoftwares**](SmAPI.md#GetNetworkSmDeviceSoftwares) | **Get** /networks/{networkId}/sm/devices/{deviceId}/softwares | Get a list of softwares associated with a device
[**GetNetworkSmDeviceWlanLists**](SmAPI.md#GetNetworkSmDeviceWlanLists) | **Get** /networks/{networkId}/sm/devices/{deviceId}/wlanLists | List the saved SSID names on a device
[**GetNetworkSmDevices**](SmAPI.md#GetNetworkSmDevices) | **Get** /networks/{networkId}/sm/devices | List the devices enrolled in an SM network with various specified fields and filters
[**GetNetworkSmProfiles**](SmAPI.md#GetNetworkSmProfiles) | **Get** /networks/{networkId}/sm/profiles | List all profiles in a network
[**GetNetworkSmTargetGroup**](SmAPI.md#GetNetworkSmTargetGroup) | **Get** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Return a target group
[**GetNetworkSmTargetGroups**](SmAPI.md#GetNetworkSmTargetGroups) | **Get** /networks/{networkId}/sm/targetGroups | List the target groups in this network
[**GetNetworkSmTrustedAccessConfigs**](SmAPI.md#GetNetworkSmTrustedAccessConfigs) | **Get** /networks/{networkId}/sm/trustedAccessConfigs | List Trusted Access Configs
[**GetNetworkSmUserAccessDevices**](SmAPI.md#GetNetworkSmUserAccessDevices) | **Get** /networks/{networkId}/sm/userAccessDevices | List User Access Devices and its Trusted Access Connections
[**GetNetworkSmUserDeviceProfiles**](SmAPI.md#GetNetworkSmUserDeviceProfiles) | **Get** /networks/{networkId}/sm/users/{userId}/deviceProfiles | Get the profiles associated with a user
[**GetNetworkSmUserSoftwares**](SmAPI.md#GetNetworkSmUserSoftwares) | **Get** /networks/{networkId}/sm/users/{userId}/softwares | Get a list of softwares associated with a user
[**GetNetworkSmUsers**](SmAPI.md#GetNetworkSmUsers) | **Get** /networks/{networkId}/sm/users | List the owners in an SM network with various specified fields and filters
[**GetOrganizationSmAdminsRole**](SmAPI.md#GetOrganizationSmAdminsRole) | **Get** /organizations/{organizationId}/sm/admins/roles/{roleId} | Return a Limited Access Role
[**GetOrganizationSmAdminsRoles**](SmAPI.md#GetOrganizationSmAdminsRoles) | **Get** /organizations/{organizationId}/sm/admins/roles | List the Limited Access Roles for an organization
[**GetOrganizationSmApnsCert**](SmAPI.md#GetOrganizationSmApnsCert) | **Get** /organizations/{organizationId}/sm/apnsCert | Get the organization&#39;s APNS certificate
[**GetOrganizationSmSentryPoliciesAssignmentsByNetwork**](SmAPI.md#GetOrganizationSmSentryPoliciesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork | List the Sentry Policies for an organization ordered in ascending order of priority
[**GetOrganizationSmVppAccount**](SmAPI.md#GetOrganizationSmVppAccount) | **Get** /organizations/{organizationId}/sm/vppAccounts/{vppAccountId} | Get a hash containing the unparsed token of the VPP account with the given ID
[**GetOrganizationSmVppAccounts**](SmAPI.md#GetOrganizationSmVppAccounts) | **Get** /organizations/{organizationId}/sm/vppAccounts | List the VPP accounts in the organization
[**InstallNetworkSmDeviceApps**](SmAPI.md#InstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/installApps | Install applications on a device
[**LockNetworkSmDevices**](SmAPI.md#LockNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/lock | Lock a set of devices
[**ModifyNetworkSmDevicesTags**](SmAPI.md#ModifyNetworkSmDevicesTags) | **Post** /networks/{networkId}/sm/devices/modifyTags | Add, delete, or update the tags of a set of devices
[**MoveNetworkSmDevices**](SmAPI.md#MoveNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/move | Move a set of devices to a new network
[**RebootNetworkSmDevices**](SmAPI.md#RebootNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/reboot | Reboot a set of endpoints
[**RefreshNetworkSmDeviceDetails**](SmAPI.md#RefreshNetworkSmDeviceDetails) | **Post** /networks/{networkId}/sm/devices/{deviceId}/refreshDetails | Refresh the details of a device
[**ShutdownNetworkSmDevices**](SmAPI.md#ShutdownNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/shutdown | Shutdown a set of endpoints
[**UnenrollNetworkSmDevice**](SmAPI.md#UnenrollNetworkSmDevice) | **Post** /networks/{networkId}/sm/devices/{deviceId}/unenroll | Unenroll a device
[**UninstallNetworkSmDeviceApps**](SmAPI.md#UninstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/uninstallApps | Uninstall applications on a device
[**UpdateNetworkSmDevicesFields**](SmAPI.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device
[**UpdateNetworkSmTargetGroup**](SmAPI.md#UpdateNetworkSmTargetGroup) | **Put** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Update a target group
[**UpdateOrganizationSmAdminsRole**](SmAPI.md#UpdateOrganizationSmAdminsRole) | **Put** /organizations/{organizationId}/sm/admins/roles/{roleId} | Update a Limited Access Role
[**UpdateOrganizationSmSentryPoliciesAssignments**](SmAPI.md#UpdateOrganizationSmSentryPoliciesAssignments) | **Put** /organizations/{organizationId}/sm/sentry/policies/assignments | Update an Organizations Sentry Policies using the provided list
[**WipeNetworkSmDevices**](SmAPI.md#WipeNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/wipe | Wipe a device



## CheckinNetworkSmDevices

> CheckinNetworkSmDevices200Response CheckinNetworkSmDevices(ctx, networkId).CheckinNetworkSmDevicesRequest(checkinNetworkSmDevicesRequest).Execute()

Force check-in a set of devices



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
	checkinNetworkSmDevicesRequest := *openapiclient.NewCheckinNetworkSmDevicesRequest() // CheckinNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.CheckinNetworkSmDevices(context.Background(), networkId).CheckinNetworkSmDevicesRequest(checkinNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.CheckinNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckinNetworkSmDevices`: CheckinNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.CheckinNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckinNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkinNetworkSmDevicesRequest** | [**CheckinNetworkSmDevicesRequest**](CheckinNetworkSmDevicesRequest.md) |  | 

### Return type

[**CheckinNetworkSmDevices200Response**](CheckinNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSmBypassActivationLockAttempt

> map[string]interface{} CreateNetworkSmBypassActivationLockAttempt(ctx, networkId).CreateNetworkSmBypassActivationLockAttemptRequest(createNetworkSmBypassActivationLockAttemptRequest).Execute()

Bypass activation lock attempt



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
	createNetworkSmBypassActivationLockAttemptRequest := *openapiclient.NewCreateNetworkSmBypassActivationLockAttemptRequest([]string{"Ids_example"}) // CreateNetworkSmBypassActivationLockAttemptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.CreateNetworkSmBypassActivationLockAttempt(context.Background(), networkId).CreateNetworkSmBypassActivationLockAttemptRequest(createNetworkSmBypassActivationLockAttemptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.CreateNetworkSmBypassActivationLockAttempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSmBypassActivationLockAttempt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.CreateNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSmBypassActivationLockAttemptRequest** | [**CreateNetworkSmBypassActivationLockAttemptRequest**](CreateNetworkSmBypassActivationLockAttemptRequest.md) |  | 

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


## CreateNetworkSmTargetGroup

> GetNetworkSmTargetGroups200ResponseInner CreateNetworkSmTargetGroup(ctx, networkId).CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest).Execute()

Add a target group



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
	createNetworkSmTargetGroupRequest := *openapiclient.NewCreateNetworkSmTargetGroupRequest() // CreateNetworkSmTargetGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.CreateNetworkSmTargetGroup(context.Background(), networkId).CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.CreateNetworkSmTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSmTargetGroup`: GetNetworkSmTargetGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.CreateNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSmTargetGroupRequest** | [**CreateNetworkSmTargetGroupRequest**](CreateNetworkSmTargetGroupRequest.md) |  | 

### Return type

[**GetNetworkSmTargetGroups200ResponseInner**](GetNetworkSmTargetGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner CreateOrganizationSmAdminsRole(ctx, organizationId).CreateOrganizationSmAdminsRoleRequest(createOrganizationSmAdminsRoleRequest).Execute()

Create a Limited Access Role



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
	createOrganizationSmAdminsRoleRequest := *openapiclient.NewCreateOrganizationSmAdminsRoleRequest("Name_example") // CreateOrganizationSmAdminsRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.CreateOrganizationSmAdminsRole(context.Background(), organizationId).CreateOrganizationSmAdminsRoleRequest(createOrganizationSmAdminsRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.CreateOrganizationSmAdminsRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.CreateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSmAdminsRoleRequest** | [**CreateOrganizationSmAdminsRoleRequest**](CreateOrganizationSmAdminsRoleRequest.md) |  | 

### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSmTargetGroup

> DeleteNetworkSmTargetGroup(ctx, networkId, targetGroupId).Execute()

Delete a target group from a network



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
	targetGroupId := "targetGroupId_example" // string | Target group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmAPI.DeleteNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.DeleteNetworkSmTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSmTargetGroupRequest struct via the builder pattern


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


## DeleteNetworkSmUserAccessDevice

> DeleteNetworkSmUserAccessDevice(ctx, networkId, userAccessDeviceId).Execute()

Delete a User Access Device



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
	userAccessDeviceId := "userAccessDeviceId_example" // string | User access device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmAPI.DeleteNetworkSmUserAccessDevice(context.Background(), networkId, userAccessDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.DeleteNetworkSmUserAccessDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userAccessDeviceId** | **string** | User access device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSmUserAccessDeviceRequest struct via the builder pattern


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


## DeleteOrganizationSmAdminsRole

> DeleteOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Delete a Limited Access Role



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmAPI.DeleteOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.DeleteOrganizationSmAdminsRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSmAdminsRoleRequest struct via the builder pattern


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


## GetNetworkSmBypassActivationLockAttempt

> map[string]interface{} GetNetworkSmBypassActivationLockAttempt(ctx, networkId, attemptId).Execute()

Bypass activation lock attempt status



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
	attemptId := "attemptId_example" // string | Attempt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmBypassActivationLockAttempt(context.Background(), networkId, attemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmBypassActivationLockAttempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmBypassActivationLockAttempt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**attemptId** | **string** | Attempt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


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


## GetNetworkSmDeviceCellularUsageHistory

> []GetNetworkSmDeviceCellularUsageHistory200ResponseInner GetNetworkSmDeviceCellularUsageHistory(ctx, networkId, deviceId).Execute()

Return the client's daily cellular data usage history



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceCellularUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceCellularUsageHistory`: []GetNetworkSmDeviceCellularUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceCellularUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCellularUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCellularUsageHistory200ResponseInner**](GetNetworkSmDeviceCellularUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceCerts

> []GetNetworkSmDeviceCerts200ResponseInner GetNetworkSmDeviceCerts(ctx, networkId, deviceId).Execute()

List the certs on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceCerts(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceCerts`: []GetNetworkSmDeviceCerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCerts200ResponseInner**](GetNetworkSmDeviceCerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceConnectivity

> []GetNetworkSmDeviceConnectivity200ResponseInner GetNetworkSmDeviceConnectivity(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns historical connectivity data (whether a device is regularly checking in to Dashboard).



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceConnectivity(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceConnectivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceConnectivity`: []GetNetworkSmDeviceConnectivity200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceConnectivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceConnectivity200ResponseInner**](GetNetworkSmDeviceConnectivity200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDesktopLogs

> []GetNetworkSmDeviceDesktopLogs200ResponseInner GetNetworkSmDeviceDesktopLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager network connection details for desktop devices.



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceDesktopLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceDesktopLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDesktopLogs`: []GetNetworkSmDeviceDesktopLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceDesktopLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDesktopLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceDesktopLogs200ResponseInner**](GetNetworkSmDeviceDesktopLogs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceCommandLogs

> []GetNetworkSmDeviceDeviceCommandLogs200ResponseInner GetNetworkSmDeviceDeviceCommandLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of commands sent to Systems Manager devices



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceDeviceCommandLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceDeviceCommandLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDeviceCommandLogs`: []GetNetworkSmDeviceDeviceCommandLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceDeviceCommandLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceCommandLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceDeviceCommandLogs200ResponseInner**](GetNetworkSmDeviceDeviceCommandLogs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmDeviceDeviceProfiles(ctx, networkId, deviceId).Execute()

Get the installed profiles associated with a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceNetworkAdapters

> []GetNetworkSmDeviceNetworkAdapters200ResponseInner GetNetworkSmDeviceNetworkAdapters(ctx, networkId, deviceId).Execute()

List the network adapters of a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceNetworkAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceNetworkAdapters`: []GetNetworkSmDeviceNetworkAdapters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceNetworkAdapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceNetworkAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceNetworkAdapters200ResponseInner**](GetNetworkSmDeviceNetworkAdapters200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevicePerformanceHistory

> []GetNetworkSmDevicePerformanceHistory200ResponseInner GetNetworkSmDevicePerformanceHistory(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager client metrics for desktop devices.



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDevicePerformanceHistory(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDevicePerformanceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDevicePerformanceHistory`: []GetNetworkSmDevicePerformanceHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDevicePerformanceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicePerformanceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDevicePerformanceHistory200ResponseInner**](GetNetworkSmDevicePerformanceHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceRestrictions

> GetNetworkSmDeviceRestrictions200Response GetNetworkSmDeviceRestrictions(ctx, networkId, deviceId).Execute()

List the restrictions on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceRestrictions(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceRestrictions`: GetNetworkSmDeviceRestrictions200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSmDeviceRestrictions200Response**](GetNetworkSmDeviceRestrictions200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSecurityCenters

> []GetNetworkSmDeviceSecurityCenters200ResponseInner GetNetworkSmDeviceSecurityCenters(ctx, networkId, deviceId).Execute()

List the security centers on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceSecurityCenters(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceSecurityCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceSecurityCenters`: []GetNetworkSmDeviceSecurityCenters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceSecurityCenters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSecurityCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSecurityCenters200ResponseInner**](GetNetworkSmDeviceSecurityCenters200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmDeviceSoftwares(ctx, networkId, deviceId).Execute()

Get a list of softwares associated with a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceSoftwares(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceSoftwares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceWlanLists

> []GetNetworkSmDeviceWlanLists200ResponseInner GetNetworkSmDeviceWlanLists(ctx, networkId, deviceId).Execute()

List the saved SSID names on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDeviceWlanLists(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDeviceWlanLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceWlanLists`: []GetNetworkSmDeviceWlanLists200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDeviceWlanLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceWlanListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceWlanLists200ResponseInner**](GetNetworkSmDeviceWlanLists200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevices

> []GetNetworkSmDevices200ResponseInner GetNetworkSmDevices(ctx, networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the devices enrolled in an SM network with various specified fields and filters



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
	fields := []string{"Inner_example"} // []string | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. (optional)
	wifiMacs := []string{"Inner_example"} // []string | Filter devices by wifi mac(s). (optional)
	serials := []string{"Inner_example"} // []string | Filter devices by serial(s). (optional)
	ids := []string{"Inner_example"} // []string | Filter devices by id(s). (optional)
	uuids := []string{"Inner_example"} // []string | Filter devices by uuid(s). (optional)
	systemTypes := []string{"Inner_example"} // []string | Filter devices by system type(s). (optional)
	scope := []string{"Inner_example"} // []string | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmDevices(context.Background(), networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDevices`: []GetNetworkSmDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. | 
 **wifiMacs** | **[]string** | Filter devices by wifi mac(s). | 
 **serials** | **[]string** | Filter devices by serial(s). | 
 **ids** | **[]string** | Filter devices by id(s). | 
 **uuids** | **[]string** | Filter devices by uuid(s). | 
 **systemTypes** | **[]string** | Filter devices by system type(s). | 
 **scope** | **[]string** | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDevices200ResponseInner**](GetNetworkSmDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmProfiles

> []GetNetworkSmProfiles200ResponseInner GetNetworkSmProfiles(ctx, networkId).PayloadTypes(payloadTypes).Execute()

List all profiles in a network



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
	payloadTypes := []string{"Inner_example"} // []string | Filter by payload types (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmProfiles(context.Background(), networkId).PayloadTypes(payloadTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmProfiles`: []GetNetworkSmProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payloadTypes** | **[]string** | Filter by payload types | 

### Return type

[**[]GetNetworkSmProfiles200ResponseInner**](GetNetworkSmProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmTargetGroup

> GetNetworkSmTargetGroups200ResponseInner GetNetworkSmTargetGroup(ctx, networkId, targetGroupId).WithDetails(withDetails).Execute()

Return a target group



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
	targetGroupId := "targetGroupId_example" // string | Target group ID
	withDetails := true // bool | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).WithDetails(withDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmTargetGroup`: GetNetworkSmTargetGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDetails** | **bool** | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response | 

### Return type

[**GetNetworkSmTargetGroups200ResponseInner**](GetNetworkSmTargetGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmTargetGroups

> []GetNetworkSmTargetGroups200ResponseInner GetNetworkSmTargetGroups(ctx, networkId).WithDetails(withDetails).Execute()

List the target groups in this network



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
	withDetails := true // bool | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmTargetGroups(context.Background(), networkId).WithDetails(withDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmTargetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmTargetGroups`: []GetNetworkSmTargetGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmTargetGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmTargetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDetails** | **bool** | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response | 

### Return type

[**[]GetNetworkSmTargetGroups200ResponseInner**](GetNetworkSmTargetGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmTrustedAccessConfigs

> []GetNetworkSmTrustedAccessConfigs200ResponseInner GetNetworkSmTrustedAccessConfigs(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Trusted Access Configs



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmTrustedAccessConfigs(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmTrustedAccessConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmTrustedAccessConfigs`: []GetNetworkSmTrustedAccessConfigs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmTrustedAccessConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmTrustedAccessConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmTrustedAccessConfigs200ResponseInner**](GetNetworkSmTrustedAccessConfigs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserAccessDevices

> []GetNetworkSmUserAccessDevices200ResponseInner GetNetworkSmUserAccessDevices(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List User Access Devices and its Trusted Access Connections



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmUserAccessDevices(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmUserAccessDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserAccessDevices`: []GetNetworkSmUserAccessDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmUserAccessDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserAccessDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmUserAccessDevices200ResponseInner**](GetNetworkSmUserAccessDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmUserDeviceProfiles(ctx, networkId, userId).Execute()

Get the profiles associated with a user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmUserDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmUserDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmUserSoftwares(ctx, networkId, userId).Execute()

Get a list of softwares associated with a user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmUserSoftwares(context.Background(), networkId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmUserSoftwares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmUserSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUsers

> []GetNetworkSmUsers200ResponseInner GetNetworkSmUsers(ctx, networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()

List the owners in an SM network with various specified fields and filters



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
	ids := []string{"Inner_example"} // []string | Filter users by id(s). (optional)
	usernames := []string{"Inner_example"} // []string | Filter users by username(s). (optional)
	emails := []string{"Inner_example"} // []string | Filter users by email(s). (optional)
	scope := []string{"Inner_example"} // []string | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetNetworkSmUsers(context.Background(), networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetNetworkSmUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUsers`: []GetNetworkSmUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetNetworkSmUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Filter users by id(s). | 
 **usernames** | **[]string** | Filter users by username(s). | 
 **emails** | **[]string** | Filter users by email(s). | 
 **scope** | **[]string** | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. | 

### Return type

[**[]GetNetworkSmUsers200ResponseInner**](GetNetworkSmUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner GetOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Return a Limited Access Role



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
	roleId := "roleId_example" // string | Role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmAdminsRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRoles

> GetOrganizationSmAdminsRoles200Response GetOrganizationSmAdminsRoles(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the Limited Access Roles for an organization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetOrganizationSmAdminsRoles(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmAdminsRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmAdminsRoles`: GetOrganizationSmAdminsRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmAdminsRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationSmAdminsRoles200Response**](GetOrganizationSmAdminsRoles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmApnsCert

> GetOrganizationSmApnsCert200Response GetOrganizationSmApnsCert(ctx, organizationId).Execute()

Get the organization's APNS certificate



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
	resp, r, err := apiClient.SmAPI.GetOrganizationSmApnsCert(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmApnsCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmApnsCert`: GetOrganizationSmApnsCert200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmApnsCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmApnsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationSmApnsCert200Response**](GetOrganizationSmApnsCert200Response.md)

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
	resp, r, err := apiClient.SmAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: []GetOrganizationSmSentryPoliciesAssignmentsByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: %v\n", resp)
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


## GetOrganizationSmVppAccount

> GetOrganizationSmVppAccounts200ResponseInner GetOrganizationSmVppAccount(ctx, organizationId, vppAccountId).Execute()

Get a hash containing the unparsed token of the VPP account with the given ID



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
	vppAccountId := "vppAccountId_example" // string | Vpp account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.GetOrganizationSmVppAccount(context.Background(), organizationId, vppAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmVppAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmVppAccount`: GetOrganizationSmVppAccounts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmVppAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**vppAccountId** | **string** | Vpp account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmVppAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSmVppAccounts200ResponseInner**](GetOrganizationSmVppAccounts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmVppAccounts

> []GetOrganizationSmVppAccounts200ResponseInner GetOrganizationSmVppAccounts(ctx, organizationId).Execute()

List the VPP accounts in the organization



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
	resp, r, err := apiClient.SmAPI.GetOrganizationSmVppAccounts(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.GetOrganizationSmVppAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSmVppAccounts`: []GetOrganizationSmVppAccounts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.GetOrganizationSmVppAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmVppAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationSmVppAccounts200ResponseInner**](GetOrganizationSmVppAccounts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallNetworkSmDeviceApps

> map[string]interface{} InstallNetworkSmDeviceApps(ctx, networkId, deviceId).InstallNetworkSmDeviceAppsRequest(installNetworkSmDeviceAppsRequest).Execute()

Install applications on a device



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
	deviceId := "deviceId_example" // string | Device ID
	installNetworkSmDeviceAppsRequest := *openapiclient.NewInstallNetworkSmDeviceAppsRequest([]string{"AppIds_example"}) // InstallNetworkSmDeviceAppsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.InstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).InstallNetworkSmDeviceAppsRequest(installNetworkSmDeviceAppsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.InstallNetworkSmDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallNetworkSmDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.InstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installNetworkSmDeviceAppsRequest** | [**InstallNetworkSmDeviceAppsRequest**](InstallNetworkSmDeviceAppsRequest.md) |  | 

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


## LockNetworkSmDevices

> CheckinNetworkSmDevices200Response LockNetworkSmDevices(ctx, networkId).LockNetworkSmDevicesRequest(lockNetworkSmDevicesRequest).Execute()

Lock a set of devices



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
	lockNetworkSmDevicesRequest := *openapiclient.NewLockNetworkSmDevicesRequest() // LockNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.LockNetworkSmDevices(context.Background(), networkId).LockNetworkSmDevicesRequest(lockNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.LockNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockNetworkSmDevices`: CheckinNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.LockNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockNetworkSmDevicesRequest** | [**LockNetworkSmDevicesRequest**](LockNetworkSmDevicesRequest.md) |  | 

### Return type

[**CheckinNetworkSmDevices200Response**](CheckinNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNetworkSmDevicesTags

> []ModifyNetworkSmDevicesTags200ResponseInner ModifyNetworkSmDevicesTags(ctx, networkId).ModifyNetworkSmDevicesTagsRequest(modifyNetworkSmDevicesTagsRequest).Execute()

Add, delete, or update the tags of a set of devices



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
	modifyNetworkSmDevicesTagsRequest := *openapiclient.NewModifyNetworkSmDevicesTagsRequest([]string{"Tags_example"}, "UpdateAction_example") // ModifyNetworkSmDevicesTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.ModifyNetworkSmDevicesTags(context.Background(), networkId).ModifyNetworkSmDevicesTagsRequest(modifyNetworkSmDevicesTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.ModifyNetworkSmDevicesTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNetworkSmDevicesTags`: []ModifyNetworkSmDevicesTags200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.ModifyNetworkSmDevicesTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNetworkSmDevicesTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyNetworkSmDevicesTagsRequest** | [**ModifyNetworkSmDevicesTagsRequest**](ModifyNetworkSmDevicesTagsRequest.md) |  | 

### Return type

[**[]ModifyNetworkSmDevicesTags200ResponseInner**](ModifyNetworkSmDevicesTags200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveNetworkSmDevices

> MoveNetworkSmDevices200Response MoveNetworkSmDevices(ctx, networkId).MoveNetworkSmDevicesRequest(moveNetworkSmDevicesRequest).Execute()

Move a set of devices to a new network



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
	moveNetworkSmDevicesRequest := *openapiclient.NewMoveNetworkSmDevicesRequest("NewNetwork_example") // MoveNetworkSmDevicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.MoveNetworkSmDevices(context.Background(), networkId).MoveNetworkSmDevicesRequest(moveNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.MoveNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveNetworkSmDevices`: MoveNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.MoveNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveNetworkSmDevicesRequest** | [**MoveNetworkSmDevicesRequest**](MoveNetworkSmDevicesRequest.md) |  | 

### Return type

[**MoveNetworkSmDevices200Response**](MoveNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootNetworkSmDevices

> RebootNetworkSmDevices200Response RebootNetworkSmDevices(ctx, networkId).RebootNetworkSmDevicesRequest(rebootNetworkSmDevicesRequest).Execute()

Reboot a set of endpoints



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
	rebootNetworkSmDevicesRequest := *openapiclient.NewRebootNetworkSmDevicesRequest() // RebootNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.RebootNetworkSmDevices(context.Background(), networkId).RebootNetworkSmDevicesRequest(rebootNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.RebootNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootNetworkSmDevices`: RebootNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.RebootNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rebootNetworkSmDevicesRequest** | [**RebootNetworkSmDevicesRequest**](RebootNetworkSmDevicesRequest.md) |  | 

### Return type

[**RebootNetworkSmDevices200Response**](RebootNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshNetworkSmDeviceDetails

> map[string]interface{} RefreshNetworkSmDeviceDetails(ctx, networkId, deviceId).Execute()

Refresh the details of a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.RefreshNetworkSmDeviceDetails(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.RefreshNetworkSmDeviceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshNetworkSmDeviceDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.RefreshNetworkSmDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshNetworkSmDeviceDetailsRequest struct via the builder pattern


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


## ShutdownNetworkSmDevices

> RebootNetworkSmDevices200Response ShutdownNetworkSmDevices(ctx, networkId).ShutdownNetworkSmDevicesRequest(shutdownNetworkSmDevicesRequest).Execute()

Shutdown a set of endpoints



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
	shutdownNetworkSmDevicesRequest := *openapiclient.NewShutdownNetworkSmDevicesRequest() // ShutdownNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.ShutdownNetworkSmDevices(context.Background(), networkId).ShutdownNetworkSmDevicesRequest(shutdownNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.ShutdownNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShutdownNetworkSmDevices`: RebootNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.ShutdownNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shutdownNetworkSmDevicesRequest** | [**ShutdownNetworkSmDevicesRequest**](ShutdownNetworkSmDevicesRequest.md) |  | 

### Return type

[**RebootNetworkSmDevices200Response**](RebootNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnenrollNetworkSmDevice

> UnenrollNetworkSmDevice200Response UnenrollNetworkSmDevice(ctx, networkId, deviceId).Execute()

Unenroll a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.UnenrollNetworkSmDevice(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UnenrollNetworkSmDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnenrollNetworkSmDevice`: UnenrollNetworkSmDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UnenrollNetworkSmDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollNetworkSmDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UnenrollNetworkSmDevice200Response**](UnenrollNetworkSmDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallNetworkSmDeviceApps

> map[string]interface{} UninstallNetworkSmDeviceApps(ctx, networkId, deviceId).UninstallNetworkSmDeviceAppsRequest(uninstallNetworkSmDeviceAppsRequest).Execute()

Uninstall applications on a device



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
	deviceId := "deviceId_example" // string | Device ID
	uninstallNetworkSmDeviceAppsRequest := *openapiclient.NewUninstallNetworkSmDeviceAppsRequest([]string{"AppIds_example"}) // UninstallNetworkSmDeviceAppsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.UninstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).UninstallNetworkSmDeviceAppsRequest(uninstallNetworkSmDeviceAppsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UninstallNetworkSmDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallNetworkSmDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UninstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uninstallNetworkSmDeviceAppsRequest** | [**UninstallNetworkSmDeviceAppsRequest**](UninstallNetworkSmDeviceAppsRequest.md) |  | 

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


## UpdateNetworkSmDevicesFields

> []UpdateNetworkSmDevicesFields200ResponseInner UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()

Modify the fields of a device



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
	updateNetworkSmDevicesFieldsRequest := *openapiclient.NewUpdateNetworkSmDevicesFieldsRequest(*openapiclient.NewUpdateNetworkSmDevicesFieldsRequestDeviceFields()) // UpdateNetworkSmDevicesFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UpdateNetworkSmDevicesFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSmDevicesFields`: []UpdateNetworkSmDevicesFields200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFieldsRequest** | [**UpdateNetworkSmDevicesFieldsRequest**](UpdateNetworkSmDevicesFieldsRequest.md) |  | 

### Return type

[**[]UpdateNetworkSmDevicesFields200ResponseInner**](UpdateNetworkSmDevicesFields200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSmTargetGroup

> GetNetworkSmTargetGroups200ResponseInner UpdateNetworkSmTargetGroup(ctx, networkId, targetGroupId).CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest).Execute()

Update a target group



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
	targetGroupId := "targetGroupId_example" // string | Target group ID
	createNetworkSmTargetGroupRequest := *openapiclient.NewCreateNetworkSmTargetGroupRequest() // CreateNetworkSmTargetGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.UpdateNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).CreateNetworkSmTargetGroupRequest(createNetworkSmTargetGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UpdateNetworkSmTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSmTargetGroup`: GetNetworkSmTargetGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UpdateNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSmTargetGroupRequest** | [**CreateNetworkSmTargetGroupRequest**](CreateNetworkSmTargetGroupRequest.md) |  | 

### Return type

[**GetNetworkSmTargetGroups200ResponseInner**](GetNetworkSmTargetGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner UpdateOrganizationSmAdminsRole(ctx, organizationId, roleId).UpdateOrganizationSmAdminsRoleRequest(updateOrganizationSmAdminsRoleRequest).Execute()

Update a Limited Access Role



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
	roleId := "roleId_example" // string | Role ID
	updateOrganizationSmAdminsRoleRequest := *openapiclient.NewUpdateOrganizationSmAdminsRoleRequest() // UpdateOrganizationSmAdminsRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.UpdateOrganizationSmAdminsRole(context.Background(), organizationId, roleId).UpdateOrganizationSmAdminsRoleRequest(updateOrganizationSmAdminsRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UpdateOrganizationSmAdminsRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UpdateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSmAdminsRoleRequest** | [**UpdateOrganizationSmAdminsRoleRequest**](UpdateOrganizationSmAdminsRoleRequest.md) |  | 

### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

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
	resp, r, err := apiClient.SmAPI.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).UpdateOrganizationSmSentryPoliciesAssignmentsRequest(updateOrganizationSmSentryPoliciesAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.UpdateOrganizationSmSentryPoliciesAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSmSentryPoliciesAssignments`: UpdateOrganizationSmSentryPoliciesAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.UpdateOrganizationSmSentryPoliciesAssignments`: %v\n", resp)
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


## WipeNetworkSmDevices

> WipeNetworkSmDevices200Response WipeNetworkSmDevices(ctx, networkId).WipeNetworkSmDevicesRequest(wipeNetworkSmDevicesRequest).Execute()

Wipe a device



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
	wipeNetworkSmDevicesRequest := *openapiclient.NewWipeNetworkSmDevicesRequest() // WipeNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmAPI.WipeNetworkSmDevices(context.Background(), networkId).WipeNetworkSmDevicesRequest(wipeNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmAPI.WipeNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WipeNetworkSmDevices`: WipeNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SmAPI.WipeNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWipeNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wipeNetworkSmDevicesRequest** | [**WipeNetworkSmDevicesRequest**](WipeNetworkSmDevicesRequest.md) |  | 

### Return type

[**WipeNetworkSmDevices200Response**](WipeNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


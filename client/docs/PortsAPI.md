# \PortsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNetworkWirelessEthernetPortsProfiles**](PortsAPI.md#AssignNetworkWirelessEthernetPortsProfiles) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/assign | Assign AP port profile to list of APs
[**CreateNetworkWirelessEthernetPortsProfile**](PortsAPI.md#CreateNetworkWirelessEthernetPortsProfile) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles | Create an AP port profile
[**CycleDeviceSwitchPorts**](PortsAPI.md#CycleDeviceSwitchPorts) | **Post** /devices/{serial}/switch/ports/cycle | Cycle a set of switch ports
[**DeleteNetworkWirelessEthernetPortsProfile**](PortsAPI.md#DeleteNetworkWirelessEthernetPortsProfile) | **Delete** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Delete an AP port profile
[**GetDeviceSwitchPort**](PortsAPI.md#GetDeviceSwitchPort) | **Get** /devices/{serial}/switch/ports/{portId} | Return a switch port
[**GetDeviceSwitchPorts**](PortsAPI.md#GetDeviceSwitchPorts) | **Get** /devices/{serial}/switch/ports | List the switch ports for a switch
[**GetDeviceSwitchPortsStatuses**](PortsAPI.md#GetDeviceSwitchPortsStatuses) | **Get** /devices/{serial}/switch/ports/statuses | Return the status for all the ports of a switch
[**GetDeviceSwitchPortsStatusesPackets**](PortsAPI.md#GetDeviceSwitchPortsStatusesPackets) | **Get** /devices/{serial}/switch/ports/statuses/packets | Return the packet counters for all the ports of a switch
[**GetNetworkAppliancePort**](PortsAPI.md#GetNetworkAppliancePort) | **Get** /networks/{networkId}/appliance/ports/{portId} | Return per-port VLAN settings for a single MX port.
[**GetNetworkAppliancePorts**](PortsAPI.md#GetNetworkAppliancePorts) | **Get** /networks/{networkId}/appliance/ports | List per-port VLAN settings for all ports of a MX.
[**GetNetworkWirelessEthernetPortsProfile**](PortsAPI.md#GetNetworkWirelessEthernetPortsProfile) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Show the AP port profile by ID for this network
[**GetNetworkWirelessEthernetPortsProfiles**](PortsAPI.md#GetNetworkWirelessEthernetPortsProfiles) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles | List the AP port profiles for this network
[**GetOrganizationConfigTemplateSwitchProfilePort**](PortsAPI.md#GetOrganizationConfigTemplateSwitchProfilePort) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Return a switch template port
[**GetOrganizationConfigTemplateSwitchProfilePorts**](PortsAPI.md#GetOrganizationConfigTemplateSwitchProfilePorts) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports | Return all the ports of a switch template
[**GetOrganizationSwitchPortsBySwitch**](PortsAPI.md#GetOrganizationSwitchPortsBySwitch) | **Get** /organizations/{organizationId}/switch/ports/bySwitch | List the switchports in an organization by switch
[**GetOrganizationSwitchPortsOverview**](PortsAPI.md#GetOrganizationSwitchPortsOverview) | **Get** /organizations/{organizationId}/switch/ports/overview | Returns the counts of all active ports for the requested timespan, grouped by speed
[**SetNetworkWirelessEthernetPortsProfilesDefault**](PortsAPI.md#SetNetworkWirelessEthernetPortsProfilesDefault) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/setDefault | Set the AP port profile to be default for this network
[**UpdateDeviceSwitchPort**](PortsAPI.md#UpdateDeviceSwitchPort) | **Put** /devices/{serial}/switch/ports/{portId} | Update a switch port
[**UpdateNetworkAppliancePort**](PortsAPI.md#UpdateNetworkAppliancePort) | **Put** /networks/{networkId}/appliance/ports/{portId} | Update the per-port VLAN settings for a single MX port.
[**UpdateNetworkWirelessEthernetPortsProfile**](PortsAPI.md#UpdateNetworkWirelessEthernetPortsProfile) | **Put** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Update the AP port profile by ID for this network
[**UpdateOrganizationConfigTemplateSwitchProfilePort**](PortsAPI.md#UpdateOrganizationConfigTemplateSwitchProfilePort) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Update a switch template port



## AssignNetworkWirelessEthernetPortsProfiles

> AssignNetworkWirelessEthernetPortsProfiles201Response AssignNetworkWirelessEthernetPortsProfiles(ctx, networkId).AssignNetworkWirelessEthernetPortsProfilesRequest(assignNetworkWirelessEthernetPortsProfilesRequest).Execute()

Assign AP port profile to list of APs



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
	assignNetworkWirelessEthernetPortsProfilesRequest := *openapiclient.NewAssignNetworkWirelessEthernetPortsProfilesRequest([]string{"Serials_example"}, "ProfileId_example") // AssignNetworkWirelessEthernetPortsProfilesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.AssignNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).AssignNetworkWirelessEthernetPortsProfilesRequest(assignNetworkWirelessEthernetPortsProfilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.AssignNetworkWirelessEthernetPortsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignNetworkWirelessEthernetPortsProfiles`: AssignNetworkWirelessEthernetPortsProfiles201Response
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.AssignNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignNetworkWirelessEthernetPortsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignNetworkWirelessEthernetPortsProfilesRequest** | [**AssignNetworkWirelessEthernetPortsProfilesRequest**](AssignNetworkWirelessEthernetPortsProfilesRequest.md) |  | 

### Return type

[**AssignNetworkWirelessEthernetPortsProfiles201Response**](AssignNetworkWirelessEthernetPortsProfiles201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner CreateNetworkWirelessEthernetPortsProfile(ctx, networkId).CreateNetworkWirelessEthernetPortsProfileRequest(createNetworkWirelessEthernetPortsProfileRequest).Execute()

Create an AP port profile



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
	createNetworkWirelessEthernetPortsProfileRequest := *openapiclient.NewCreateNetworkWirelessEthernetPortsProfileRequest("Name_example", []openapiclient.CreateNetworkWirelessEthernetPortsProfileRequestPortsInner{*openapiclient.NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner("Name_example")}) // CreateNetworkWirelessEthernetPortsProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.CreateNetworkWirelessEthernetPortsProfile(context.Background(), networkId).CreateNetworkWirelessEthernetPortsProfileRequest(createNetworkWirelessEthernetPortsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.CreateNetworkWirelessEthernetPortsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.CreateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessEthernetPortsProfileRequest** | [**CreateNetworkWirelessEthernetPortsProfileRequest**](CreateNetworkWirelessEthernetPortsProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CycleDeviceSwitchPorts

> CycleDeviceSwitchPorts200Response CycleDeviceSwitchPorts(ctx, serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()

Cycle a set of switch ports



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
	serial := "serial_example" // string | Serial
	cycleDeviceSwitchPortsRequest := *openapiclient.NewCycleDeviceSwitchPortsRequest([]string{"Ports_example"}) // CycleDeviceSwitchPortsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.CycleDeviceSwitchPorts(context.Background(), serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.CycleDeviceSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CycleDeviceSwitchPorts`: CycleDeviceSwitchPorts200Response
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.CycleDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCycleDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cycleDeviceSwitchPortsRequest** | [**CycleDeviceSwitchPortsRequest**](CycleDeviceSwitchPortsRequest.md) |  | 

### Return type

[**CycleDeviceSwitchPorts200Response**](CycleDeviceSwitchPorts200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWirelessEthernetPortsProfile

> DeleteNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).Execute()

Delete an AP port profile



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
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortsAPI.DeleteNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.DeleteNetworkWirelessEthernetPortsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


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


## GetDeviceSwitchPort

> GetDeviceSwitchPorts200ResponseInner GetDeviceSwitchPort(ctx, serial, portId).Execute()

Return a switch port



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
	serial := "serial_example" // string | Serial
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetDeviceSwitchPort(context.Background(), serial, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetDeviceSwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPort`: GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetDeviceSwitchPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPorts

> []GetDeviceSwitchPorts200ResponseInner GetDeviceSwitchPorts(ctx, serial).Execute()

List the switch ports for a switch



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetDeviceSwitchPorts(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetDeviceSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPorts`: []GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPortsStatuses

> []GetDeviceSwitchPortsStatuses200ResponseInner GetDeviceSwitchPortsStatuses(ctx, serial).T0(t0).Timespan(timespan).Execute()

Return the status for all the ports of a switch



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
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetDeviceSwitchPortsStatuses(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetDeviceSwitchPortsStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPortsStatuses`: []GetDeviceSwitchPortsStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetDeviceSwitchPortsStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetDeviceSwitchPortsStatuses200ResponseInner**](GetDeviceSwitchPortsStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPortsStatusesPackets

> []GetDeviceSwitchPortsStatusesPackets200ResponseInner GetDeviceSwitchPortsStatusesPackets(ctx, serial).T0(t0).Timespan(timespan).Execute()

Return the packet counters for all the ports of a switch



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
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetDeviceSwitchPortsStatusesPackets(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetDeviceSwitchPortsStatusesPackets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPortsStatusesPackets`: []GetDeviceSwitchPortsStatusesPackets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetDeviceSwitchPortsStatusesPackets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsStatusesPacketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. | 

### Return type

[**[]GetDeviceSwitchPortsStatusesPackets200ResponseInner**](GetDeviceSwitchPortsStatusesPackets200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePort(ctx, networkId, portId).Execute()

Return per-port VLAN settings for a single MX port.



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
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetNetworkAppliancePort(context.Background(), networkId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetNetworkAppliancePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePorts

> []GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePorts(ctx, networkId).Execute()

List per-port VLAN settings for all ports of a MX.



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
	resp, r, err := apiClient.PortsAPI.GetNetworkAppliancePorts(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetNetworkAppliancePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePorts`: []GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetNetworkAppliancePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner GetNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).Execute()

Show the AP port profile by ID for this network



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
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetNetworkWirelessEthernetPortsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessEthernetPortsProfiles

> []GetNetworkWirelessEthernetPortsProfiles200ResponseInner GetNetworkWirelessEthernetPortsProfiles(ctx, networkId).Execute()

List the AP port profiles for this network



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
	resp, r, err := apiClient.PortsAPI.GetNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetNetworkWirelessEthernetPortsProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessEthernetPortsProfiles`: []GetNetworkWirelessEthernetPortsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessEthernetPortsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

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
	resp, r, err := apiClient.PortsAPI.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
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
	resp, r, err := apiClient.PortsAPI.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetOrganizationConfigTemplateSwitchProfilePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePorts`: []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetOrganizationConfigTemplateSwitchProfilePorts`: %v\n", resp)
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


## GetOrganizationSwitchPortsBySwitch

> []GetOrganizationSwitchPortsBySwitch200ResponseInner GetOrganizationSwitchPortsBySwitch(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()

List the switchports in an organization by switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	configurationUpdatedAfter := time.Now() // time.Time | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. (optional)
	mac := "mac_example" // string | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. (optional)
	macs := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided MAC addresses. (optional)
	name := "name_example" // string | Optional parameter to filter items to switches with names that contain the search term or are an exact match. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches in one of the provided networks. (optional)
	portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. (optional)
	serial := "serial_example" // string | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided serials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetOrganizationSwitchPortsBySwitch(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetOrganizationSwitchPortsBySwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSwitchPortsBySwitch`: []GetOrganizationSwitchPortsBySwitch200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetOrganizationSwitchPortsBySwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsBySwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **time.Time** | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. | 
 **mac** | **string** | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. | 
 **macs** | **[]string** | Optional parameter to filter items to switches that have one of the provided MAC addresses. | 
 **name** | **string** | Optional parameter to filter items to switches with names that contain the search term or are an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter items to switches in one of the provided networks. | 
 **portProfileIds** | **[]string** | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. | 
 **serial** | **string** | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. | 
 **serials** | **[]string** | Optional parameter to filter items to switches that have one of the provided serials. | 

### Return type

[**[]GetOrganizationSwitchPortsBySwitch200ResponseInner**](GetOrganizationSwitchPortsBySwitch200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsOverview

> GetOrganizationSwitchPortsOverview200Response GetOrganizationSwitchPortsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns the counts of all active ports for the requested timespan, grouped by speed



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.GetOrganizationSwitchPortsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.GetOrganizationSwitchPortsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSwitchPortsOverview`: GetOrganizationSwitchPortsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.GetOrganizationSwitchPortsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**GetOrganizationSwitchPortsOverview200Response**](GetOrganizationSwitchPortsOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNetworkWirelessEthernetPortsProfilesDefault

> SetNetworkWirelessEthernetPortsProfilesDefault200Response SetNetworkWirelessEthernetPortsProfilesDefault(ctx, networkId).SetNetworkWirelessEthernetPortsProfilesDefaultRequest(setNetworkWirelessEthernetPortsProfilesDefaultRequest).Execute()

Set the AP port profile to be default for this network



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
	setNetworkWirelessEthernetPortsProfilesDefaultRequest := *openapiclient.NewSetNetworkWirelessEthernetPortsProfilesDefaultRequest("ProfileId_example") // SetNetworkWirelessEthernetPortsProfilesDefaultRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.SetNetworkWirelessEthernetPortsProfilesDefault(context.Background(), networkId).SetNetworkWirelessEthernetPortsProfilesDefaultRequest(setNetworkWirelessEthernetPortsProfilesDefaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.SetNetworkWirelessEthernetPortsProfilesDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNetworkWirelessEthernetPortsProfilesDefault`: SetNetworkWirelessEthernetPortsProfilesDefault200Response
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.SetNetworkWirelessEthernetPortsProfilesDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkWirelessEthernetPortsProfilesDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNetworkWirelessEthernetPortsProfilesDefaultRequest** | [**SetNetworkWirelessEthernetPortsProfilesDefaultRequest**](SetNetworkWirelessEthernetPortsProfilesDefaultRequest.md) |  | 

### Return type

[**SetNetworkWirelessEthernetPortsProfilesDefault200Response**](SetNetworkWirelessEthernetPortsProfilesDefault200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchPort

> GetDeviceSwitchPorts200ResponseInner UpdateDeviceSwitchPort(ctx, serial, portId).UpdateDeviceSwitchPortRequest(updateDeviceSwitchPortRequest).Execute()

Update a switch port



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
	serial := "serial_example" // string | Serial
	portId := "portId_example" // string | Port ID
	updateDeviceSwitchPortRequest := *openapiclient.NewUpdateDeviceSwitchPortRequest() // UpdateDeviceSwitchPortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.UpdateDeviceSwitchPort(context.Background(), serial, portId).UpdateDeviceSwitchPortRequest(updateDeviceSwitchPortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.UpdateDeviceSwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchPort`: GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.UpdateDeviceSwitchPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchPortRequest** | [**UpdateDeviceSwitchPortRequest**](UpdateDeviceSwitchPortRequest.md) |  | 

### Return type

[**GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner UpdateNetworkAppliancePort(ctx, networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()

Update the per-port VLAN settings for a single MX port.



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
	portId := "portId_example" // string | Port ID
	updateNetworkAppliancePortRequest := *openapiclient.NewUpdateNetworkAppliancePortRequest() // UpdateNetworkAppliancePortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.UpdateNetworkAppliancePort(context.Background(), networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.UpdateNetworkAppliancePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.UpdateNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePortRequest** | [**UpdateNetworkAppliancePortRequest**](UpdateNetworkAppliancePortRequest.md) |  | 

### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner UpdateNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).UpdateNetworkWirelessEthernetPortsProfileRequest(updateNetworkWirelessEthernetPortsProfileRequest).Execute()

Update the AP port profile by ID for this network



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
	profileId := "profileId_example" // string | Profile ID
	updateNetworkWirelessEthernetPortsProfileRequest := *openapiclient.NewUpdateNetworkWirelessEthernetPortsProfileRequest() // UpdateNetworkWirelessEthernetPortsProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortsAPI.UpdateNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).UpdateNetworkWirelessEthernetPortsProfileRequest(updateNetworkWirelessEthernetPortsProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.UpdateNetworkWirelessEthernetPortsProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.UpdateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessEthernetPortsProfileRequest** | [**UpdateNetworkWirelessEthernetPortsProfileRequest**](UpdateNetworkWirelessEthernetPortsProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

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
	resp, r, err := apiClient.PortsAPI.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.UpdateOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortsAPI.UpdateOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
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


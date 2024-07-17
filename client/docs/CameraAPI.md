# \CameraAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraQualityRetentionProfile**](CameraAPI.md#CreateNetworkCameraQualityRetentionProfile) | **Post** /networks/{networkId}/camera/qualityRetentionProfiles | Creates new quality retention profile for this network.
[**CreateNetworkCameraWirelessProfile**](CameraAPI.md#CreateNetworkCameraWirelessProfile) | **Post** /networks/{networkId}/camera/wirelessProfiles | Creates a new camera wireless profile for this network.
[**CreateOrganizationCameraCustomAnalyticsArtifact**](CameraAPI.md#CreateOrganizationCameraCustomAnalyticsArtifact) | **Post** /organizations/{organizationId}/camera/customAnalytics/artifacts | Create custom analytics artifact
[**CreateOrganizationCameraRole**](CameraAPI.md#CreateOrganizationCameraRole) | **Post** /organizations/{organizationId}/camera/roles | Creates new role for this organization.
[**DeleteNetworkCameraQualityRetentionProfile**](CameraAPI.md#DeleteNetworkCameraQualityRetentionProfile) | **Delete** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Delete an existing quality retention profile for this network.
[**DeleteNetworkCameraWirelessProfile**](CameraAPI.md#DeleteNetworkCameraWirelessProfile) | **Delete** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Delete an existing camera wireless profile for this network.
[**DeleteOrganizationCameraCustomAnalyticsArtifact**](CameraAPI.md#DeleteOrganizationCameraCustomAnalyticsArtifact) | **Delete** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Delete Custom Analytics Artifact
[**DeleteOrganizationCameraRole**](CameraAPI.md#DeleteOrganizationCameraRole) | **Delete** /organizations/{organizationId}/camera/roles/{roleId} | Delete an existing role for this organization.
[**GenerateDeviceCameraSnapshot**](CameraAPI.md#GenerateDeviceCameraSnapshot) | **Post** /devices/{serial}/camera/generateSnapshot | Generate a snapshot of what the camera sees at the specified time and return a link to that image.
[**GetDeviceCameraAnalyticsLive**](CameraAPI.md#GetDeviceCameraAnalyticsLive) | **Get** /devices/{serial}/camera/analytics/live | Returns live state from camera analytics zones
[**GetDeviceCameraAnalyticsOverview**](CameraAPI.md#GetDeviceCameraAnalyticsOverview) | **Get** /devices/{serial}/camera/analytics/overview | Returns an overview of aggregate analytics data for a timespan
[**GetDeviceCameraAnalyticsRecent**](CameraAPI.md#GetDeviceCameraAnalyticsRecent) | **Get** /devices/{serial}/camera/analytics/recent | Returns most recent record for analytics zones
[**GetDeviceCameraAnalyticsZoneHistory**](CameraAPI.md#GetDeviceCameraAnalyticsZoneHistory) | **Get** /devices/{serial}/camera/analytics/zones/{zoneId}/history | Return historical records for analytic zones
[**GetDeviceCameraAnalyticsZones**](CameraAPI.md#GetDeviceCameraAnalyticsZones) | **Get** /devices/{serial}/camera/analytics/zones | Returns all configured analytic zones for this camera
[**GetDeviceCameraCustomAnalytics**](CameraAPI.md#GetDeviceCameraCustomAnalytics) | **Get** /devices/{serial}/camera/customAnalytics | Return custom analytics settings for a camera
[**GetDeviceCameraQualityAndRetention**](CameraAPI.md#GetDeviceCameraQualityAndRetention) | **Get** /devices/{serial}/camera/qualityAndRetention | Returns quality and retention settings for the given camera
[**GetDeviceCameraSense**](CameraAPI.md#GetDeviceCameraSense) | **Get** /devices/{serial}/camera/sense | Returns sense settings for a given camera
[**GetDeviceCameraSenseObjectDetectionModels**](CameraAPI.md#GetDeviceCameraSenseObjectDetectionModels) | **Get** /devices/{serial}/camera/sense/objectDetectionModels | Returns the MV Sense object detection model list for the given camera
[**GetDeviceCameraVideoLink**](CameraAPI.md#GetDeviceCameraVideoLink) | **Get** /devices/{serial}/camera/videoLink | Returns video link to the specified camera
[**GetDeviceCameraVideoSettings**](CameraAPI.md#GetDeviceCameraVideoSettings) | **Get** /devices/{serial}/camera/video/settings | Returns video settings for the given camera
[**GetDeviceCameraWirelessProfiles**](CameraAPI.md#GetDeviceCameraWirelessProfiles) | **Get** /devices/{serial}/camera/wirelessProfiles | Returns wireless profile assigned to the given camera
[**GetNetworkCameraQualityRetentionProfile**](CameraAPI.md#GetNetworkCameraQualityRetentionProfile) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Retrieve a single quality retention profile
[**GetNetworkCameraQualityRetentionProfiles**](CameraAPI.md#GetNetworkCameraQualityRetentionProfiles) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles | List the quality retention profiles for this network
[**GetNetworkCameraSchedules**](CameraAPI.md#GetNetworkCameraSchedules) | **Get** /networks/{networkId}/camera/schedules | Returns a list of all camera recording schedules.
[**GetNetworkCameraWirelessProfile**](CameraAPI.md#GetNetworkCameraWirelessProfile) | **Get** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Retrieve a single camera wireless profile.
[**GetNetworkCameraWirelessProfiles**](CameraAPI.md#GetNetworkCameraWirelessProfiles) | **Get** /networks/{networkId}/camera/wirelessProfiles | List the camera wireless profiles for this network.
[**GetOrganizationCameraBoundariesAreasByDevice**](CameraAPI.md#GetOrganizationCameraBoundariesAreasByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/areas/byDevice | Returns all configured area boundaries of cameras
[**GetOrganizationCameraBoundariesLinesByDevice**](CameraAPI.md#GetOrganizationCameraBoundariesLinesByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/lines/byDevice | Returns all configured crossingline boundaries of cameras
[**GetOrganizationCameraCustomAnalyticsArtifact**](CameraAPI.md#GetOrganizationCameraCustomAnalyticsArtifact) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Get Custom Analytics Artifact
[**GetOrganizationCameraCustomAnalyticsArtifacts**](CameraAPI.md#GetOrganizationCameraCustomAnalyticsArtifacts) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts | List Custom Analytics Artifacts
[**GetOrganizationCameraDetectionsHistoryByBoundaryByInterval**](CameraAPI.md#GetOrganizationCameraDetectionsHistoryByBoundaryByInterval) | **Get** /organizations/{organizationId}/camera/detections/history/byBoundary/byInterval | Returns analytics data for timespans
[**GetOrganizationCameraOnboardingStatuses**](CameraAPI.md#GetOrganizationCameraOnboardingStatuses) | **Get** /organizations/{organizationId}/camera/onboarding/statuses | Fetch onboarding status of cameras
[**GetOrganizationCameraPermission**](CameraAPI.md#GetOrganizationCameraPermission) | **Get** /organizations/{organizationId}/camera/permissions/{permissionScopeId} | Retrieve a single permission scope
[**GetOrganizationCameraPermissions**](CameraAPI.md#GetOrganizationCameraPermissions) | **Get** /organizations/{organizationId}/camera/permissions | List the permissions scopes for this organization
[**GetOrganizationCameraRole**](CameraAPI.md#GetOrganizationCameraRole) | **Get** /organizations/{organizationId}/camera/roles/{roleId} | Retrieve a single role.
[**GetOrganizationCameraRoles**](CameraAPI.md#GetOrganizationCameraRoles) | **Get** /organizations/{organizationId}/camera/roles | List all the roles in this organization
[**UpdateDeviceCameraCustomAnalytics**](CameraAPI.md#UpdateDeviceCameraCustomAnalytics) | **Put** /devices/{serial}/camera/customAnalytics | Update custom analytics settings for a camera
[**UpdateDeviceCameraQualityAndRetention**](CameraAPI.md#UpdateDeviceCameraQualityAndRetention) | **Put** /devices/{serial}/camera/qualityAndRetention | Update quality and retention settings for the given camera
[**UpdateDeviceCameraSense**](CameraAPI.md#UpdateDeviceCameraSense) | **Put** /devices/{serial}/camera/sense | Update sense settings for the given camera
[**UpdateDeviceCameraVideoSettings**](CameraAPI.md#UpdateDeviceCameraVideoSettings) | **Put** /devices/{serial}/camera/video/settings | Update video settings for the given camera
[**UpdateDeviceCameraWirelessProfiles**](CameraAPI.md#UpdateDeviceCameraWirelessProfiles) | **Put** /devices/{serial}/camera/wirelessProfiles | Assign wireless profiles to the given camera
[**UpdateNetworkCameraQualityRetentionProfile**](CameraAPI.md#UpdateNetworkCameraQualityRetentionProfile) | **Put** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Update an existing quality retention profile for this network.
[**UpdateNetworkCameraWirelessProfile**](CameraAPI.md#UpdateNetworkCameraWirelessProfile) | **Put** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Update an existing camera wireless profile in this network.
[**UpdateOrganizationCameraOnboardingStatuses**](CameraAPI.md#UpdateOrganizationCameraOnboardingStatuses) | **Put** /organizations/{organizationId}/camera/onboarding/statuses | Notify that credential handoff to camera has completed
[**UpdateOrganizationCameraRole**](CameraAPI.md#UpdateOrganizationCameraRole) | **Put** /organizations/{organizationId}/camera/roles/{roleId} | Update an existing role in this organization.



## CreateNetworkCameraQualityRetentionProfile

> map[string]interface{} CreateNetworkCameraQualityRetentionProfile(ctx, networkId).CreateNetworkCameraQualityRetentionProfileRequest(createNetworkCameraQualityRetentionProfileRequest).Execute()

Creates new quality retention profile for this network.



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
	createNetworkCameraQualityRetentionProfileRequest := *openapiclient.NewCreateNetworkCameraQualityRetentionProfileRequest("Name_example") // CreateNetworkCameraQualityRetentionProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.CreateNetworkCameraQualityRetentionProfile(context.Background(), networkId).CreateNetworkCameraQualityRetentionProfileRequest(createNetworkCameraQualityRetentionProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CreateNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.CreateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraQualityRetentionProfileRequest** | [**CreateNetworkCameraQualityRetentionProfileRequest**](CreateNetworkCameraQualityRetentionProfileRequest.md) |  | 

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


## CreateNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner CreateNetworkCameraWirelessProfile(ctx, networkId).CreateNetworkCameraWirelessProfileRequest(createNetworkCameraWirelessProfileRequest).Execute()

Creates a new camera wireless profile for this network.



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
	createNetworkCameraWirelessProfileRequest := *openapiclient.NewCreateNetworkCameraWirelessProfileRequest("Name_example", *openapiclient.NewCreateNetworkCameraWirelessProfileRequestSsid()) // CreateNetworkCameraWirelessProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.CreateNetworkCameraWirelessProfile(context.Background(), networkId).CreateNetworkCameraWirelessProfileRequest(createNetworkCameraWirelessProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CreateNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.CreateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraWirelessProfileRequest** | [**CreateNetworkCameraWirelessProfileRequest**](CreateNetworkCameraWirelessProfileRequest.md) |  | 

### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationCameraCustomAnalyticsArtifact

> CreateOrganizationCameraCustomAnalyticsArtifact201Response CreateOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId).CreateOrganizationCameraCustomAnalyticsArtifactRequest(createOrganizationCameraCustomAnalyticsArtifactRequest).Execute()

Create custom analytics artifact



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
	createOrganizationCameraCustomAnalyticsArtifactRequest := *openapiclient.NewCreateOrganizationCameraCustomAnalyticsArtifactRequest() // CreateOrganizationCameraCustomAnalyticsArtifactRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.CreateOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId).CreateOrganizationCameraCustomAnalyticsArtifactRequest(createOrganizationCameraCustomAnalyticsArtifactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CreateOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationCameraCustomAnalyticsArtifact`: CreateOrganizationCameraCustomAnalyticsArtifact201Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.CreateOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraCustomAnalyticsArtifactRequest** | [**CreateOrganizationCameraCustomAnalyticsArtifactRequest**](CreateOrganizationCameraCustomAnalyticsArtifactRequest.md) |  | 

### Return type

[**CreateOrganizationCameraCustomAnalyticsArtifact201Response**](CreateOrganizationCameraCustomAnalyticsArtifact201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationCameraRole

> map[string]interface{} CreateOrganizationCameraRole(ctx, organizationId).CreateOrganizationCameraRoleRequest(createOrganizationCameraRoleRequest).Execute()

Creates new role for this organization.



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
	createOrganizationCameraRoleRequest := *openapiclient.NewCreateOrganizationCameraRoleRequest("Name_example") // CreateOrganizationCameraRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.CreateOrganizationCameraRole(context.Background(), organizationId).CreateOrganizationCameraRoleRequest(createOrganizationCameraRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CreateOrganizationCameraRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationCameraRole`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.CreateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraRoleRequest** | [**CreateOrganizationCameraRoleRequest**](CreateOrganizationCameraRoleRequest.md) |  | 

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


## DeleteNetworkCameraQualityRetentionProfile

> DeleteNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Delete an existing quality retention profile for this network.



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CameraAPI.DeleteNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.DeleteNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


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


## DeleteNetworkCameraWirelessProfile

> DeleteNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Delete an existing camera wireless profile for this network.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CameraAPI.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.DeleteNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraWirelessProfileRequest struct via the builder pattern


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


## DeleteOrganizationCameraCustomAnalyticsArtifact

> DeleteOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Delete Custom Analytics Artifact



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
	artifactId := "artifactId_example" // string | Artifact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CameraAPI.DeleteOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.DeleteOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


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


## DeleteOrganizationCameraRole

> DeleteOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Delete an existing role for this organization.



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
	r, err := apiClient.CameraAPI.DeleteOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.DeleteOrganizationCameraRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraRoleRequest struct via the builder pattern


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


## GenerateDeviceCameraSnapshot

> map[string]interface{} GenerateDeviceCameraSnapshot(ctx, serial).GenerateDeviceCameraSnapshotRequest(generateDeviceCameraSnapshotRequest).Execute()

Generate a snapshot of what the camera sees at the specified time and return a link to that image.



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
	generateDeviceCameraSnapshotRequest := *openapiclient.NewGenerateDeviceCameraSnapshotRequest() // GenerateDeviceCameraSnapshotRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GenerateDeviceCameraSnapshot(context.Background(), serial).GenerateDeviceCameraSnapshotRequest(generateDeviceCameraSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GenerateDeviceCameraSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDeviceCameraSnapshot`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GenerateDeviceCameraSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceCameraSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateDeviceCameraSnapshotRequest** | [**GenerateDeviceCameraSnapshotRequest**](GenerateDeviceCameraSnapshotRequest.md) |  | 

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


## GetDeviceCameraAnalyticsLive

> GetDeviceCameraAnalyticsLive200Response GetDeviceCameraAnalyticsLive(ctx, serial).Execute()

Returns live state from camera analytics zones



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraAnalyticsLive(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraAnalyticsLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsLive`: GetDeviceCameraAnalyticsLive200Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraAnalyticsLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraAnalyticsLive200Response**](GetDeviceCameraAnalyticsLive200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsOverview

> []GetDeviceCameraAnalyticsOverview200ResponseInner GetDeviceCameraAnalyticsOverview(ctx, serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()

Returns an overview of aggregate analytics data for a timespan



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. (optional)
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraAnalyticsOverview(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraAnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsOverview`: []GetDeviceCameraAnalyticsOverview200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsOverview200ResponseInner**](GetDeviceCameraAnalyticsOverview200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsRecent

> []GetDeviceCameraAnalyticsRecent200ResponseInner GetDeviceCameraAnalyticsRecent(ctx, serial).ObjectType(objectType).Execute()

Returns most recent record for analytics zones



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
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraAnalyticsRecent(context.Background(), serial).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraAnalyticsRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsRecent`: []GetDeviceCameraAnalyticsRecent200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraAnalyticsRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsRecent200ResponseInner**](GetDeviceCameraAnalyticsRecent200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsZoneHistory

> []GetDeviceCameraAnalyticsZoneHistory200ResponseInner GetDeviceCameraAnalyticsZoneHistory(ctx, serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()

Return historical records for analytic zones



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
	zoneId := "zoneId_example" // string | Zone ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. (optional)
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraAnalyticsZoneHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsZoneHistory`: []GetDeviceCameraAnalyticsZoneHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraAnalyticsZoneHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**zoneId** | **string** | Zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZoneHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsZoneHistory200ResponseInner**](GetDeviceCameraAnalyticsZoneHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsZones

> []GetDeviceCameraAnalyticsZones200ResponseInner GetDeviceCameraAnalyticsZones(ctx, serial).Execute()

Returns all configured analytic zones for this camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraAnalyticsZones(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraAnalyticsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsZones`: []GetDeviceCameraAnalyticsZones200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraAnalyticsZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceCameraAnalyticsZones200ResponseInner**](GetDeviceCameraAnalyticsZones200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraCustomAnalytics

> GetDeviceCameraCustomAnalytics200Response GetDeviceCameraCustomAnalytics(ctx, serial).Execute()

Return custom analytics settings for a camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraCustomAnalytics(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraCustomAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraCustomAnalytics`: GetDeviceCameraCustomAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraCustomAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraCustomAnalytics200Response**](GetDeviceCameraCustomAnalytics200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraQualityAndRetention

> map[string]interface{} GetDeviceCameraQualityAndRetention(ctx, serial).Execute()

Returns quality and retention settings for the given camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraQualityAndRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraQualityAndRetention`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraQualityAndRetentionRequest struct via the builder pattern


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


## GetDeviceCameraSense

> map[string]interface{} GetDeviceCameraSense(ctx, serial).Execute()

Returns sense settings for a given camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraSense(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraSense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraSense`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseRequest struct via the builder pattern


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


## GetDeviceCameraSenseObjectDetectionModels

> []map[string]interface{} GetDeviceCameraSenseObjectDetectionModels(ctx, serial).Execute()

Returns the MV Sense object detection model list for the given camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraSenseObjectDetectionModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraSenseObjectDetectionModels`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraSenseObjectDetectionModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseObjectDetectionModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraVideoLink

> map[string]interface{} GetDeviceCameraVideoLink(ctx, serial).Timestamp(timestamp).Execute()

Returns video link to the specified camera



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
	serial := "serial_example" // string | Serial
	timestamp := time.Now() // time.Time | [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraVideoLink(context.Background(), serial).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraVideoLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraVideoLink`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraVideoLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraVideoLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **time.Time** | [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time. | 

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


## GetDeviceCameraVideoSettings

> GetDeviceCameraVideoSettings200Response GetDeviceCameraVideoSettings(ctx, serial).Execute()

Returns video settings for the given camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraVideoSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraVideoSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraVideoSettings`: GetDeviceCameraVideoSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraVideoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraVideoSettings200Response**](GetDeviceCameraVideoSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraWirelessProfiles

> map[string]interface{} GetDeviceCameraWirelessProfiles(ctx, serial).Execute()

Returns wireless profile assigned to the given camera



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
	resp, r, err := apiClient.CameraAPI.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetDeviceCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraWirelessProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraWirelessProfilesRequest struct via the builder pattern


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


## GetNetworkCameraQualityRetentionProfile

> map[string]interface{} GetNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Retrieve a single quality retention profile



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


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


## GetNetworkCameraQualityRetentionProfiles

> []map[string]interface{} GetNetworkCameraQualityRetentionProfiles(ctx, networkId).Execute()

List the quality retention profiles for this network



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
	resp, r, err := apiClient.CameraAPI.GetNetworkCameraQualityRetentionProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetNetworkCameraQualityRetentionProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraQualityRetentionProfiles`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetNetworkCameraQualityRetentionProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraSchedules

> []GetNetworkCameraSchedules200ResponseInner GetNetworkCameraSchedules(ctx, networkId).Execute()

Returns a list of all camera recording schedules.



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
	resp, r, err := apiClient.CameraAPI.GetNetworkCameraSchedules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetNetworkCameraSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraSchedules`: []GetNetworkCameraSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetNetworkCameraSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkCameraSchedules200ResponseInner**](GetNetworkCameraSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner GetNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Retrieve a single camera wireless profile.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraWirelessProfiles

> []GetNetworkCameraWirelessProfiles200ResponseInner GetNetworkCameraWirelessProfiles(ctx, networkId).Execute()

List the camera wireless profiles for this network.



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
	resp, r, err := apiClient.CameraAPI.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetNetworkCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraWirelessProfiles`: []GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetNetworkCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraBoundariesAreasByDevice

> []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner GetOrganizationCameraBoundariesAreasByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured area boundaries of cameras



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
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraBoundariesAreasByDevice(context.Background(), organizationId).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraBoundariesAreasByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraBoundariesAreasByDevice`: []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraBoundariesAreasByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesAreasByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesAreasByDevice200ResponseInner**](GetOrganizationCameraBoundariesAreasByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraBoundariesLinesByDevice

> []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner GetOrganizationCameraBoundariesLinesByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured crossingline boundaries of cameras



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
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraBoundariesLinesByDevice(context.Background(), organizationId).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraBoundariesLinesByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraBoundariesLinesByDevice`: []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraBoundariesLinesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesLinesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesLinesByDevice200ResponseInner**](GetOrganizationCameraBoundariesLinesByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraCustomAnalyticsArtifact

> GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner GetOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Get Custom Analytics Artifact



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
	artifactId := "artifactId_example" // string | Artifact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraCustomAnalyticsArtifact`: GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner**](GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraCustomAnalyticsArtifacts

> []GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner GetOrganizationCameraCustomAnalyticsArtifacts(ctx, organizationId).Execute()

List Custom Analytics Artifacts



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
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraCustomAnalyticsArtifacts(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraCustomAnalyticsArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraCustomAnalyticsArtifacts`: []GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraCustomAnalyticsArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner**](GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraDetectionsHistoryByBoundaryByInterval

> []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(ctx, organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()

Returns analytics data for timespans



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
	boundaryIds := []string{"Inner_example"} // []string | A list of boundary ids. The returned cameras will be filtered to only include these ids.
	ranges := []openapiclient.GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner{*openapiclient.NewGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner("StartTime_example", "EndTime_example", int32(123))} // []GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner | A list of time ranges with intervals
	duration := int32(56) // int32 | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. (optional)
	boundaryTypes := []string{"BoundaryTypes_example"} // []string | The detection types. Defaults to 'person'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(context.Background(), organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boundaryIds** | **[]string** | A list of boundary ids. The returned cameras will be filtered to only include these ids. | 
 **ranges** | [**[]GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner.md) | A list of time ranges with intervals | 
 **duration** | **int32** | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. | 
 **boundaryTypes** | **[]string** | The detection types. Defaults to &#39;person&#39;. | 

### Return type

[**[]GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraOnboardingStatuses

> []map[string]interface{} GetOrganizationCameraOnboardingStatuses(ctx, organizationId).Serials(serials).NetworkIds(networkIds).Execute()

Fetch onboarding status of cameras



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
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)
	networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned cameras will be filtered to only include these networks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Serials(serials).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraOnboardingStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraOnboardingStatuses`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 
 **networkIds** | **[]string** | A list of network IDs. The returned cameras will be filtered to only include these networks. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraPermission

> GetOrganizationCameraPermissions200ResponseInner GetOrganizationCameraPermission(ctx, organizationId, permissionScopeId).Execute()

Retrieve a single permission scope



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
	permissionScopeId := "permissionScopeId_example" // string | Permission scope ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraPermission(context.Background(), organizationId, permissionScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraPermission`: GetOrganizationCameraPermissions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**permissionScopeId** | **string** | Permission scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationCameraPermissions200ResponseInner**](GetOrganizationCameraPermissions200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraPermissions

> []GetOrganizationCameraPermissions200ResponseInner GetOrganizationCameraPermissions(ctx, organizationId).Execute()

List the permissions scopes for this organization



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
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraPermissions(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraPermissions`: []GetOrganizationCameraPermissions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationCameraPermissions200ResponseInner**](GetOrganizationCameraPermissions200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraRole

> map[string]interface{} GetOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Retrieve a single role.



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
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraRole`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRoleRequest struct via the builder pattern


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


## GetOrganizationCameraRoles

> []map[string]interface{} GetOrganizationCameraRoles(ctx, organizationId).Execute()

List all the roles in this organization



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
	resp, r, err := apiClient.CameraAPI.GetOrganizationCameraRoles(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.GetOrganizationCameraRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraRoles`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.GetOrganizationCameraRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraCustomAnalytics

> GetDeviceCameraCustomAnalytics200Response UpdateDeviceCameraCustomAnalytics(ctx, serial).UpdateDeviceCameraCustomAnalyticsRequest(updateDeviceCameraCustomAnalyticsRequest).Execute()

Update custom analytics settings for a camera



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
	updateDeviceCameraCustomAnalyticsRequest := *openapiclient.NewUpdateDeviceCameraCustomAnalyticsRequest() // UpdateDeviceCameraCustomAnalyticsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateDeviceCameraCustomAnalytics(context.Background(), serial).UpdateDeviceCameraCustomAnalyticsRequest(updateDeviceCameraCustomAnalyticsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateDeviceCameraCustomAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraCustomAnalytics`: GetDeviceCameraCustomAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraCustomAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraCustomAnalyticsRequest** | [**UpdateDeviceCameraCustomAnalyticsRequest**](UpdateDeviceCameraCustomAnalyticsRequest.md) |  | 

### Return type

[**GetDeviceCameraCustomAnalytics200Response**](GetDeviceCameraCustomAnalytics200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraQualityAndRetention

> map[string]interface{} UpdateDeviceCameraQualityAndRetention(ctx, serial).UpdateDeviceCameraQualityAndRetentionRequest(updateDeviceCameraQualityAndRetentionRequest).Execute()

Update quality and retention settings for the given camera



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
	updateDeviceCameraQualityAndRetentionRequest := *openapiclient.NewUpdateDeviceCameraQualityAndRetentionRequest() // UpdateDeviceCameraQualityAndRetentionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).UpdateDeviceCameraQualityAndRetentionRequest(updateDeviceCameraQualityAndRetentionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateDeviceCameraQualityAndRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraQualityAndRetention`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraQualityAndRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraQualityAndRetentionRequest** | [**UpdateDeviceCameraQualityAndRetentionRequest**](UpdateDeviceCameraQualityAndRetentionRequest.md) |  | 

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


## UpdateDeviceCameraSense

> map[string]interface{} UpdateDeviceCameraSense(ctx, serial).UpdateDeviceCameraSenseRequest(updateDeviceCameraSenseRequest).Execute()

Update sense settings for the given camera



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
	updateDeviceCameraSenseRequest := *openapiclient.NewUpdateDeviceCameraSenseRequest() // UpdateDeviceCameraSenseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateDeviceCameraSense(context.Background(), serial).UpdateDeviceCameraSenseRequest(updateDeviceCameraSenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateDeviceCameraSense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraSense`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraSenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraSenseRequest** | [**UpdateDeviceCameraSenseRequest**](UpdateDeviceCameraSenseRequest.md) |  | 

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


## UpdateDeviceCameraVideoSettings

> GetDeviceCameraVideoSettings200Response UpdateDeviceCameraVideoSettings(ctx, serial).UpdateDeviceCameraVideoSettingsRequest(updateDeviceCameraVideoSettingsRequest).Execute()

Update video settings for the given camera



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
	updateDeviceCameraVideoSettingsRequest := *openapiclient.NewUpdateDeviceCameraVideoSettingsRequest() // UpdateDeviceCameraVideoSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateDeviceCameraVideoSettings(context.Background(), serial).UpdateDeviceCameraVideoSettingsRequest(updateDeviceCameraVideoSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateDeviceCameraVideoSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraVideoSettings`: GetDeviceCameraVideoSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraVideoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraVideoSettingsRequest** | [**UpdateDeviceCameraVideoSettingsRequest**](UpdateDeviceCameraVideoSettingsRequest.md) |  | 

### Return type

[**GetDeviceCameraVideoSettings200Response**](GetDeviceCameraVideoSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraWirelessProfiles

> map[string]interface{} UpdateDeviceCameraWirelessProfiles(ctx, serial).UpdateDeviceCameraWirelessProfilesRequest(updateDeviceCameraWirelessProfilesRequest).Execute()

Assign wireless profiles to the given camera



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
	updateDeviceCameraWirelessProfilesRequest := *openapiclient.NewUpdateDeviceCameraWirelessProfilesRequest(*openapiclient.NewUpdateDeviceCameraWirelessProfilesRequestIds()) // UpdateDeviceCameraWirelessProfilesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).UpdateDeviceCameraWirelessProfilesRequest(updateDeviceCameraWirelessProfilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateDeviceCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraWirelessProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraWirelessProfilesRequest** | [**UpdateDeviceCameraWirelessProfilesRequest**](UpdateDeviceCameraWirelessProfilesRequest.md) |  | 

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


## UpdateNetworkCameraQualityRetentionProfile

> map[string]interface{} UpdateNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfileRequest(updateNetworkCameraQualityRetentionProfileRequest).Execute()

Update an existing quality retention profile for this network.



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID
	updateNetworkCameraQualityRetentionProfileRequest := *openapiclient.NewUpdateNetworkCameraQualityRetentionProfileRequest() // UpdateNetworkCameraQualityRetentionProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfileRequest(updateNetworkCameraQualityRetentionProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraQualityRetentionProfileRequest** | [**UpdateNetworkCameraQualityRetentionProfileRequest**](UpdateNetworkCameraQualityRetentionProfileRequest.md) |  | 

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


## UpdateNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner UpdateNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfileRequest(updateNetworkCameraWirelessProfileRequest).Execute()

Update an existing camera wireless profile in this network.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID
	updateNetworkCameraWirelessProfileRequest := *openapiclient.NewUpdateNetworkCameraWirelessProfileRequest() // UpdateNetworkCameraWirelessProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfileRequest(updateNetworkCameraWirelessProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraWirelessProfileRequest** | [**UpdateNetworkCameraWirelessProfileRequest**](UpdateNetworkCameraWirelessProfileRequest.md) |  | 

### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCameraOnboardingStatuses

> map[string]interface{} UpdateOrganizationCameraOnboardingStatuses(ctx, organizationId).UpdateOrganizationCameraOnboardingStatusesRequest(updateOrganizationCameraOnboardingStatusesRequest).Execute()

Notify that credential handoff to camera has completed



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
	updateOrganizationCameraOnboardingStatusesRequest := *openapiclient.NewUpdateOrganizationCameraOnboardingStatusesRequest() // UpdateOrganizationCameraOnboardingStatusesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).UpdateOrganizationCameraOnboardingStatusesRequest(updateOrganizationCameraOnboardingStatusesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateOrganizationCameraOnboardingStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationCameraOnboardingStatuses`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationCameraOnboardingStatusesRequest** | [**UpdateOrganizationCameraOnboardingStatusesRequest**](UpdateOrganizationCameraOnboardingStatusesRequest.md) |  | 

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


## UpdateOrganizationCameraRole

> map[string]interface{} UpdateOrganizationCameraRole(ctx, organizationId, roleId).UpdateOrganizationCameraRoleRequest(updateOrganizationCameraRoleRequest).Execute()

Update an existing role in this organization.



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
	updateOrganizationCameraRoleRequest := *openapiclient.NewUpdateOrganizationCameraRoleRequest() // UpdateOrganizationCameraRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.UpdateOrganizationCameraRole(context.Background(), organizationId, roleId).UpdateOrganizationCameraRoleRequest(updateOrganizationCameraRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.UpdateOrganizationCameraRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationCameraRole`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.UpdateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCameraRoleRequest** | [**UpdateOrganizationCameraRoleRequest**](UpdateOrganizationCameraRoleRequest.md) |  | 

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


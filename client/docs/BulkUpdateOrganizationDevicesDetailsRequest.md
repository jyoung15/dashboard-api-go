# BulkUpdateOrganizationDevicesDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | **[]string** | A list of serials of devices to update | 
**Details** | [**[]BulkUpdateOrganizationDevicesDetailsRequestDetailsInner**](BulkUpdateOrganizationDevicesDetailsRequestDetailsInner.md) | An array of details | 

## Methods

### NewBulkUpdateOrganizationDevicesDetailsRequest

`func NewBulkUpdateOrganizationDevicesDetailsRequest(serials []string, details []BulkUpdateOrganizationDevicesDetailsRequestDetailsInner, ) *BulkUpdateOrganizationDevicesDetailsRequest`

NewBulkUpdateOrganizationDevicesDetailsRequest instantiates a new BulkUpdateOrganizationDevicesDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateOrganizationDevicesDetailsRequestWithDefaults

`func NewBulkUpdateOrganizationDevicesDetailsRequestWithDefaults() *BulkUpdateOrganizationDevicesDetailsRequest`

NewBulkUpdateOrganizationDevicesDetailsRequestWithDefaults instantiates a new BulkUpdateOrganizationDevicesDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetDetails

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) GetDetails() []BulkUpdateOrganizationDevicesDetailsRequestDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) GetDetailsOk() (*[]BulkUpdateOrganizationDevicesDetailsRequestDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkUpdateOrganizationDevicesDetailsRequest) SetDetails(v []BulkUpdateOrganizationDevicesDetailsRequestDetailsInner)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



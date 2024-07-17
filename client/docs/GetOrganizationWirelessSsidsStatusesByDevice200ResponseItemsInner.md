# GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for device. | [optional] 
**Name** | Pointer to **string** | Name of device. | [optional] 
**Network** | Pointer to [**GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerNetwork**](GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerNetwork.md) |  | [optional] 
**BasicServiceSets** | Pointer to [**[]GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerBasicServiceSetsInner**](GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerBasicServiceSetsInner.md) | Status information for wireless access points. | [optional] 

## Methods

### NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner

`func NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner() *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner`

NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerWithDefaults

`func NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerWithDefaults() *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner`

NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerWithDefaults instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetNetwork() GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetNetworkOk() (*GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) SetNetwork(v GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetBasicServiceSets

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetBasicServiceSets() []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerBasicServiceSetsInner`

GetBasicServiceSets returns the BasicServiceSets field if non-nil, zero value otherwise.

### GetBasicServiceSetsOk

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) GetBasicServiceSetsOk() (*[]GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerBasicServiceSetsInner, bool)`

GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicServiceSets

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) SetBasicServiceSets(v []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInnerBasicServiceSetsInner)`

SetBasicServiceSets sets BasicServiceSets field to given value.

### HasBasicServiceSets

`func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) HasBasicServiceSets() bool`

HasBasicServiceSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



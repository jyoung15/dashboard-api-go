# GetOrganizationApplianceUplinkStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Serial** | Pointer to **string** | The uplink serial | [optional] 
**Model** | Pointer to **string** | The uplink model | [optional] 
**LastReportedAt** | Pointer to **time.Time** | Last reported time for the device | [optional] 
**HighAvailability** | Pointer to [**GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability**](GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability.md) |  | [optional] 
**Uplinks** | Pointer to [**[]GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner**](GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner.md) | Uplinks | [optional] 

## Methods

### NewGetOrganizationApplianceUplinkStatuses200ResponseInner

`func NewGetOrganizationApplianceUplinkStatuses200ResponseInner() *GetOrganizationApplianceUplinkStatuses200ResponseInner`

NewGetOrganizationApplianceUplinkStatuses200ResponseInner instantiates a new GetOrganizationApplianceUplinkStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApplianceUplinkStatuses200ResponseInnerWithDefaults

`func NewGetOrganizationApplianceUplinkStatuses200ResponseInnerWithDefaults() *GetOrganizationApplianceUplinkStatuses200ResponseInner`

NewGetOrganizationApplianceUplinkStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationApplianceUplinkStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetHighAvailability

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetHighAvailability() GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetHighAvailabilityOk() (*GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetHighAvailability(v GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetUplinks

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetUplinks() []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetUplinksOk() (*[]GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetUplinks(v []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



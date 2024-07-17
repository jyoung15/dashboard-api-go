# GetOrganizationSwitchPortsOverview200ResponseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of ports | [optional] 
**ByStatus** | Pointer to [**GetOrganizationSwitchPortsOverview200ResponseCountsByStatus**](GetOrganizationSwitchPortsOverview200ResponseCountsByStatus.md) |  | [optional] 

## Methods

### NewGetOrganizationSwitchPortsOverview200ResponseCounts

`func NewGetOrganizationSwitchPortsOverview200ResponseCounts() *GetOrganizationSwitchPortsOverview200ResponseCounts`

NewGetOrganizationSwitchPortsOverview200ResponseCounts instantiates a new GetOrganizationSwitchPortsOverview200ResponseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSwitchPortsOverview200ResponseCountsWithDefaults

`func NewGetOrganizationSwitchPortsOverview200ResponseCountsWithDefaults() *GetOrganizationSwitchPortsOverview200ResponseCounts`

NewGetOrganizationSwitchPortsOverview200ResponseCountsWithDefaults instantiates a new GetOrganizationSwitchPortsOverview200ResponseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) GetByStatus() GetOrganizationSwitchPortsOverview200ResponseCountsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) GetByStatusOk() (*GetOrganizationSwitchPortsOverview200ResponseCountsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) SetByStatus(v GetOrganizationSwitchPortsOverview200ResponseCountsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *GetOrganizationSwitchPortsOverview200ResponseCounts) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



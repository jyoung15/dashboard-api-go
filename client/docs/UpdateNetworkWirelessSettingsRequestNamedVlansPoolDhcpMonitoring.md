# UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP. | [optional] 
**Duration** | Pointer to **int32** | The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool. | [optional] 

## Methods

### NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring

`func NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring() *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring`

NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoringWithDefaults

`func NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoringWithDefaults() *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring`

NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoringWithDefaults instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDuration

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



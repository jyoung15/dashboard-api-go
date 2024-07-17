# GetNetworkWirelessClientConnectivityEvents200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OccurredAt** | Pointer to **time.Time** | Timestamp at which the event occurred | [optional] 
**Band** | Pointer to **string** | Wireless band the event occurred on | [optional] 
**SsidNumber** | Pointer to **int32** | Number of the SSID the event occurred in | [optional] 
**Type** | Pointer to **string** | Event type | [optional] 
**Subtype** | Pointer to **string** | Event subtype | [optional] 
**Severity** | Pointer to **string** | Event severity | [optional] 
**DurationMs** | Pointer to **int32** | Duration of the event in milliseconds | [optional] 
**Channel** | Pointer to **int32** | Wireless channel the event occurred over | [optional] 
**Rssi** | Pointer to **int32** | RSSI recorded at the time of the event | [optional] 
**EventData** | Pointer to **map[string]interface{}** | Additional information relevant to the given event. Properties vary based on event type. | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device the event occurred for | [optional] 

## Methods

### NewGetNetworkWirelessClientConnectivityEvents200ResponseInner

`func NewGetNetworkWirelessClientConnectivityEvents200ResponseInner() *GetNetworkWirelessClientConnectivityEvents200ResponseInner`

NewGetNetworkWirelessClientConnectivityEvents200ResponseInner instantiates a new GetNetworkWirelessClientConnectivityEvents200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessClientConnectivityEvents200ResponseInnerWithDefaults

`func NewGetNetworkWirelessClientConnectivityEvents200ResponseInnerWithDefaults() *GetNetworkWirelessClientConnectivityEvents200ResponseInner`

NewGetNetworkWirelessClientConnectivityEvents200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessClientConnectivityEvents200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurredAt

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetBand

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetSsidNumber

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSeverity

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDurationMs

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetChannel

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRssi

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetEventData

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetEventData() map[string]interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetEventDataOk() (*map[string]interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetEventData(v map[string]interface{})`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *GetNetworkWirelessClientConnectivityEvents200ResponseInner) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



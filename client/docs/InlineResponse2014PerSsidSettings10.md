# InlineResponse2014PerSsidSettings10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of SSID | [optional] 
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of this SSID. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**BandOperationMode** | Pointer to **string** | Choice between &#39;dual&#39;, &#39;2.4ghz&#39; or &#39;5ghz&#39;. | [optional] 
**BandSteeringEnabled** | Pointer to **bool** | Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false. | [optional] 

## Methods

### NewInlineResponse2014PerSsidSettings10

`func NewInlineResponse2014PerSsidSettings10() *InlineResponse2014PerSsidSettings10`

NewInlineResponse2014PerSsidSettings10 instantiates a new InlineResponse2014PerSsidSettings10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2014PerSsidSettings10WithDefaults

`func NewInlineResponse2014PerSsidSettings10WithDefaults() *InlineResponse2014PerSsidSettings10`

NewInlineResponse2014PerSsidSettings10WithDefaults instantiates a new InlineResponse2014PerSsidSettings10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2014PerSsidSettings10) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2014PerSsidSettings10) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2014PerSsidSettings10) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2014PerSsidSettings10) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse2014PerSsidSettings10) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse2014PerSsidSettings10) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse2014PerSsidSettings10) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse2014PerSsidSettings10) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetBandOperationMode

`func (o *InlineResponse2014PerSsidSettings10) GetBandOperationMode() string`

GetBandOperationMode returns the BandOperationMode field if non-nil, zero value otherwise.

### GetBandOperationModeOk

`func (o *InlineResponse2014PerSsidSettings10) GetBandOperationModeOk() (*string, bool)`

GetBandOperationModeOk returns a tuple with the BandOperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandOperationMode

`func (o *InlineResponse2014PerSsidSettings10) SetBandOperationMode(v string)`

SetBandOperationMode sets BandOperationMode field to given value.

### HasBandOperationMode

`func (o *InlineResponse2014PerSsidSettings10) HasBandOperationMode() bool`

HasBandOperationMode returns a boolean if a field has been set.

### GetBandSteeringEnabled

`func (o *InlineResponse2014PerSsidSettings10) GetBandSteeringEnabled() bool`

GetBandSteeringEnabled returns the BandSteeringEnabled field if non-nil, zero value otherwise.

### GetBandSteeringEnabledOk

`func (o *InlineResponse2014PerSsidSettings10) GetBandSteeringEnabledOk() (*bool, bool)`

GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringEnabled

`func (o *InlineResponse2014PerSsidSettings10) SetBandSteeringEnabled(v bool)`

SetBandSteeringEnabled sets BandSteeringEnabled field to given value.

### HasBandSteeringEnabled

`func (o *InlineResponse2014PerSsidSettings10) HasBandSteeringEnabled() bool`

HasBandSteeringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


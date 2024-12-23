# CreateNetworkWirelessRfProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new profile. Must be unique. This param is required on creation. | 
**ClientBalancingEnabled** | Pointer to **bool** | Steers client to best available access point. Can be either true or false. Defaults to true. | [optional] 
**MinBitrateType** | Pointer to **string** | Minimum bitrate can be set to either &#39;band&#39; or &#39;ssid&#39;. Defaults to band. | [optional] 
**BandSelectionType** | **string** | Band selection can be set to either &#39;ssid&#39; or &#39;ap&#39;. This param is required on creation. | 
**ApBandSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestApBandSettings**](CreateNetworkWirelessRfProfileRequestApBandSettings.md) |  | [optional] 
**TwoFourGhzSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings**](GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings.md) |  | [optional] 
**FiveGhzSettings** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings**](GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings.md) |  | [optional] 
**SixGhzSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestSixGhzSettings**](CreateNetworkWirelessRfProfileRequestSixGhzSettings.md) |  | [optional] 
**Transmission** | Pointer to [**GetNetworkWirelessRfProfiles200ResponseTransmission**](GetNetworkWirelessRfProfiles200ResponseTransmission.md) |  | [optional] 
**PerSsidSettings** | Pointer to [**CreateNetworkWirelessRfProfileRequestPerSsidSettings**](CreateNetworkWirelessRfProfileRequestPerSsidSettings.md) |  | [optional] 
**FlexRadios** | Pointer to [**CreateNetworkWirelessRfProfileRequestFlexRadios**](CreateNetworkWirelessRfProfileRequestFlexRadios.md) |  | [optional] 

## Methods

### NewCreateNetworkWirelessRfProfileRequest

`func NewCreateNetworkWirelessRfProfileRequest(name string, bandSelectionType string, ) *CreateNetworkWirelessRfProfileRequest`

NewCreateNetworkWirelessRfProfileRequest instantiates a new CreateNetworkWirelessRfProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessRfProfileRequestWithDefaults

`func NewCreateNetworkWirelessRfProfileRequestWithDefaults() *CreateNetworkWirelessRfProfileRequest`

NewCreateNetworkWirelessRfProfileRequestWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkWirelessRfProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkWirelessRfProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfileRequest) GetClientBalancingEnabled() bool`

GetClientBalancingEnabled returns the ClientBalancingEnabled field if non-nil, zero value otherwise.

### GetClientBalancingEnabledOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetClientBalancingEnabledOk() (*bool, bool)`

GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfileRequest) SetClientBalancingEnabled(v bool)`

SetClientBalancingEnabled sets ClientBalancingEnabled field to given value.

### HasClientBalancingEnabled

`func (o *CreateNetworkWirelessRfProfileRequest) HasClientBalancingEnabled() bool`

HasClientBalancingEnabled returns a boolean if a field has been set.

### GetMinBitrateType

`func (o *CreateNetworkWirelessRfProfileRequest) GetMinBitrateType() string`

GetMinBitrateType returns the MinBitrateType field if non-nil, zero value otherwise.

### GetMinBitrateTypeOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetMinBitrateTypeOk() (*string, bool)`

GetMinBitrateTypeOk returns a tuple with the MinBitrateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrateType

`func (o *CreateNetworkWirelessRfProfileRequest) SetMinBitrateType(v string)`

SetMinBitrateType sets MinBitrateType field to given value.

### HasMinBitrateType

`func (o *CreateNetworkWirelessRfProfileRequest) HasMinBitrateType() bool`

HasMinBitrateType returns a boolean if a field has been set.

### GetBandSelectionType

`func (o *CreateNetworkWirelessRfProfileRequest) GetBandSelectionType() string`

GetBandSelectionType returns the BandSelectionType field if non-nil, zero value otherwise.

### GetBandSelectionTypeOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetBandSelectionTypeOk() (*string, bool)`

GetBandSelectionTypeOk returns a tuple with the BandSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSelectionType

`func (o *CreateNetworkWirelessRfProfileRequest) SetBandSelectionType(v string)`

SetBandSelectionType sets BandSelectionType field to given value.


### GetApBandSettings

`func (o *CreateNetworkWirelessRfProfileRequest) GetApBandSettings() CreateNetworkWirelessRfProfileRequestApBandSettings`

GetApBandSettings returns the ApBandSettings field if non-nil, zero value otherwise.

### GetApBandSettingsOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetApBandSettingsOk() (*CreateNetworkWirelessRfProfileRequestApBandSettings, bool)`

GetApBandSettingsOk returns a tuple with the ApBandSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApBandSettings

`func (o *CreateNetworkWirelessRfProfileRequest) SetApBandSettings(v CreateNetworkWirelessRfProfileRequestApBandSettings)`

SetApBandSettings sets ApBandSettings field to given value.

### HasApBandSettings

`func (o *CreateNetworkWirelessRfProfileRequest) HasApBandSettings() bool`

HasApBandSettings returns a boolean if a field has been set.

### GetTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettings() GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings`

GetTwoFourGhzSettings returns the TwoFourGhzSettings field if non-nil, zero value otherwise.

### GetTwoFourGhzSettingsOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings, bool)`

GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) SetTwoFourGhzSettings(v GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings)`

SetTwoFourGhzSettings sets TwoFourGhzSettings field to given value.

### HasTwoFourGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) HasTwoFourGhzSettings() bool`

HasTwoFourGhzSettings returns a boolean if a field has been set.

### GetFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) GetFiveGhzSettings() GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings`

GetFiveGhzSettings returns the FiveGhzSettings field if non-nil, zero value otherwise.

### GetFiveGhzSettingsOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetFiveGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings, bool)`

GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) SetFiveGhzSettings(v GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings)`

SetFiveGhzSettings sets FiveGhzSettings field to given value.

### HasFiveGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) HasFiveGhzSettings() bool`

HasFiveGhzSettings returns a boolean if a field has been set.

### GetSixGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) GetSixGhzSettings() CreateNetworkWirelessRfProfileRequestSixGhzSettings`

GetSixGhzSettings returns the SixGhzSettings field if non-nil, zero value otherwise.

### GetSixGhzSettingsOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetSixGhzSettingsOk() (*CreateNetworkWirelessRfProfileRequestSixGhzSettings, bool)`

GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSixGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) SetSixGhzSettings(v CreateNetworkWirelessRfProfileRequestSixGhzSettings)`

SetSixGhzSettings sets SixGhzSettings field to given value.

### HasSixGhzSettings

`func (o *CreateNetworkWirelessRfProfileRequest) HasSixGhzSettings() bool`

HasSixGhzSettings returns a boolean if a field has been set.

### GetTransmission

`func (o *CreateNetworkWirelessRfProfileRequest) GetTransmission() GetNetworkWirelessRfProfiles200ResponseTransmission`

GetTransmission returns the Transmission field if non-nil, zero value otherwise.

### GetTransmissionOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetTransmissionOk() (*GetNetworkWirelessRfProfiles200ResponseTransmission, bool)`

GetTransmissionOk returns a tuple with the Transmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmission

`func (o *CreateNetworkWirelessRfProfileRequest) SetTransmission(v GetNetworkWirelessRfProfiles200ResponseTransmission)`

SetTransmission sets Transmission field to given value.

### HasTransmission

`func (o *CreateNetworkWirelessRfProfileRequest) HasTransmission() bool`

HasTransmission returns a boolean if a field has been set.

### GetPerSsidSettings

`func (o *CreateNetworkWirelessRfProfileRequest) GetPerSsidSettings() CreateNetworkWirelessRfProfileRequestPerSsidSettings`

GetPerSsidSettings returns the PerSsidSettings field if non-nil, zero value otherwise.

### GetPerSsidSettingsOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkWirelessRfProfileRequestPerSsidSettings, bool)`

GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSsidSettings

`func (o *CreateNetworkWirelessRfProfileRequest) SetPerSsidSettings(v CreateNetworkWirelessRfProfileRequestPerSsidSettings)`

SetPerSsidSettings sets PerSsidSettings field to given value.

### HasPerSsidSettings

`func (o *CreateNetworkWirelessRfProfileRequest) HasPerSsidSettings() bool`

HasPerSsidSettings returns a boolean if a field has been set.

### GetFlexRadios

`func (o *CreateNetworkWirelessRfProfileRequest) GetFlexRadios() CreateNetworkWirelessRfProfileRequestFlexRadios`

GetFlexRadios returns the FlexRadios field if non-nil, zero value otherwise.

### GetFlexRadiosOk

`func (o *CreateNetworkWirelessRfProfileRequest) GetFlexRadiosOk() (*CreateNetworkWirelessRfProfileRequestFlexRadios, bool)`

GetFlexRadiosOk returns a tuple with the FlexRadios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexRadios

`func (o *CreateNetworkWirelessRfProfileRequest) SetFlexRadios(v CreateNetworkWirelessRfProfileRequestFlexRadios)`

SetFlexRadios sets FlexRadios field to given value.

### HasFlexRadios

`func (o *CreateNetworkWirelessRfProfileRequest) HasFlexRadios() bool`

HasFlexRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequest{}

// UpdateNetworkWirelessSsidSplashSettingsRequest struct for UpdateNetworkWirelessSsidSplashSettingsRequest
type UpdateNetworkWirelessSsidSplashSettingsRequest struct {
	// [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see 'useSplashUrl'
	SplashUrl *string `json:"splashUrl,omitempty"`
	// [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID's access control settings, it may not be possible to use the custom splash URL.
	UseSplashUrl *bool `json:"useSplashUrl,omitempty"`
	// Splash timeout in minutes. This will determine how often users will see the splash page.
	SplashTimeout *int32 `json:"splashTimeout,omitempty"`
	// The custom redirect URL where the users will go after the splash page.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.
	UseRedirectUrl *bool `json:"useRedirectUrl,omitempty"`
	// The welcome message for the users on the splash page.
	WelcomeMessage *string `json:"welcomeMessage,omitempty"`
	SplashLogo *UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo `json:"splashLogo,omitempty"`
	SplashImage *UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage `json:"splashImage,omitempty"`
	SplashPrepaidFront *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`
	// How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	BlockAllTrafficBeforeSignOn *bool `json:"blockAllTrafficBeforeSignOn,omitempty"`
	// How login attempts should be handled when the controller is unreachable. Can be either 'open', 'restricted', or 'default'.
	ControllerDisconnectionBehavior *string `json:"controllerDisconnectionBehavior,omitempty"`
	// Whether or not to allow simultaneous logins from different devices.
	AllowSimultaneousLogins *bool `json:"allowSimultaneousLogins,omitempty"`
	GuestSponsorship *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship `json:"guestSponsorship,omitempty"`
	Billing *UpdateNetworkWirelessSsidSplashSettingsRequestBilling `json:"billing,omitempty"`
	SentryEnrollment *UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment `json:"sentryEnrollment,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequest instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequest() *UpdateNetworkWirelessSsidSplashSettingsRequest {
	this := UpdateNetworkWirelessSsidSplashSettingsRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequest {
	this := UpdateNetworkWirelessSsidSplashSettingsRequest{}
	return &this
}

// GetSplashUrl returns the SplashUrl field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashUrl() string {
	if o == nil || IsNil(o.SplashUrl) {
		var ret string
		return ret
	}
	return *o.SplashUrl
}

// GetSplashUrlOk returns a tuple with the SplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SplashUrl) {
		return nil, false
	}
	return o.SplashUrl, true
}

// HasSplashUrl returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashUrl() bool {
	if o != nil && !IsNil(o.SplashUrl) {
		return true
	}

	return false
}

// SetSplashUrl gets a reference to the given string and assigns it to the SplashUrl field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashUrl(v string) {
	o.SplashUrl = &v
}

// GetUseSplashUrl returns the UseSplashUrl field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseSplashUrl() bool {
	if o == nil || IsNil(o.UseSplashUrl) {
		var ret bool
		return ret
	}
	return *o.UseSplashUrl
}

// GetUseSplashUrlOk returns a tuple with the UseSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseSplashUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSplashUrl) {
		return nil, false
	}
	return o.UseSplashUrl, true
}

// HasUseSplashUrl returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasUseSplashUrl() bool {
	if o != nil && !IsNil(o.UseSplashUrl) {
		return true
	}

	return false
}

// SetUseSplashUrl gets a reference to the given bool and assigns it to the UseSplashUrl field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetUseSplashUrl(v bool) {
	o.UseSplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashTimeout() int32 {
	if o == nil || IsNil(o.SplashTimeout) {
		var ret int32
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SplashTimeout) {
		return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashTimeout() bool {
	if o != nil && !IsNil(o.SplashTimeout) {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given int32 and assigns it to the SplashTimeout field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashTimeout(v int32) {
	o.SplashTimeout = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetUseRedirectUrl returns the UseRedirectUrl field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseRedirectUrl() bool {
	if o == nil || IsNil(o.UseRedirectUrl) {
		var ret bool
		return ret
	}
	return *o.UseRedirectUrl
}

// GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetUseRedirectUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseRedirectUrl) {
		return nil, false
	}
	return o.UseRedirectUrl, true
}

// HasUseRedirectUrl returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasUseRedirectUrl() bool {
	if o != nil && !IsNil(o.UseRedirectUrl) {
		return true
	}

	return false
}

// SetUseRedirectUrl gets a reference to the given bool and assigns it to the UseRedirectUrl field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetUseRedirectUrl(v bool) {
	o.UseRedirectUrl = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetWelcomeMessage() string {
	if o == nil || IsNil(o.WelcomeMessage) {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || IsNil(o.WelcomeMessage) {
		return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasWelcomeMessage() bool {
	if o != nil && !IsNil(o.WelcomeMessage) {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

// GetSplashLogo returns the SplashLogo field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashLogo() UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo {
	if o == nil || IsNil(o.SplashLogo) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo
		return ret
	}
	return *o.SplashLogo
}

// GetSplashLogoOk returns a tuple with the SplashLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashLogoOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo, bool) {
	if o == nil || IsNil(o.SplashLogo) {
		return nil, false
	}
	return o.SplashLogo, true
}

// HasSplashLogo returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashLogo() bool {
	if o != nil && !IsNil(o.SplashLogo) {
		return true
	}

	return false
}

// SetSplashLogo gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo and assigns it to the SplashLogo field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashLogo(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashLogo) {
	o.SplashLogo = &v
}

// GetSplashImage returns the SplashImage field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashImage() UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage {
	if o == nil || IsNil(o.SplashImage) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage
		return ret
	}
	return *o.SplashImage
}

// GetSplashImageOk returns a tuple with the SplashImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashImageOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage, bool) {
	if o == nil || IsNil(o.SplashImage) {
		return nil, false
	}
	return o.SplashImage, true
}

// HasSplashImage returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashImage() bool {
	if o != nil && !IsNil(o.SplashImage) {
		return true
	}

	return false
}

// SetSplashImage gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage and assigns it to the SplashImage field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashImage(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashImage) {
	o.SplashImage = &v
}

// GetSplashPrepaidFront returns the SplashPrepaidFront field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashPrepaidFront() UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront {
	if o == nil || IsNil(o.SplashPrepaidFront) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront
		return ret
	}
	return *o.SplashPrepaidFront
}

// GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSplashPrepaidFrontOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront, bool) {
	if o == nil || IsNil(o.SplashPrepaidFront) {
		return nil, false
	}
	return o.SplashPrepaidFront, true
}

// HasSplashPrepaidFront returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSplashPrepaidFront() bool {
	if o != nil && !IsNil(o.SplashPrepaidFront) {
		return true
	}

	return false
}

// SetSplashPrepaidFront gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront and assigns it to the SplashPrepaidFront field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSplashPrepaidFront(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) {
	o.SplashPrepaidFront = &v
}

// GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBlockAllTrafficBeforeSignOn() bool {
	if o == nil || IsNil(o.BlockAllTrafficBeforeSignOn) {
		var ret bool
		return ret
	}
	return *o.BlockAllTrafficBeforeSignOn
}

// GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockAllTrafficBeforeSignOn) {
		return nil, false
	}
	return o.BlockAllTrafficBeforeSignOn, true
}

// HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasBlockAllTrafficBeforeSignOn() bool {
	if o != nil && !IsNil(o.BlockAllTrafficBeforeSignOn) {
		return true
	}

	return false
}

// SetBlockAllTrafficBeforeSignOn gets a reference to the given bool and assigns it to the BlockAllTrafficBeforeSignOn field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetBlockAllTrafficBeforeSignOn(v bool) {
	o.BlockAllTrafficBeforeSignOn = &v
}

// GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetControllerDisconnectionBehavior() string {
	if o == nil || IsNil(o.ControllerDisconnectionBehavior) {
		var ret string
		return ret
	}
	return *o.ControllerDisconnectionBehavior
}

// GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetControllerDisconnectionBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerDisconnectionBehavior) {
		return nil, false
	}
	return o.ControllerDisconnectionBehavior, true
}

// HasControllerDisconnectionBehavior returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasControllerDisconnectionBehavior() bool {
	if o != nil && !IsNil(o.ControllerDisconnectionBehavior) {
		return true
	}

	return false
}

// SetControllerDisconnectionBehavior gets a reference to the given string and assigns it to the ControllerDisconnectionBehavior field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetControllerDisconnectionBehavior(v string) {
	o.ControllerDisconnectionBehavior = &v
}

// GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetAllowSimultaneousLogins() bool {
	if o == nil || IsNil(o.AllowSimultaneousLogins) {
		var ret bool
		return ret
	}
	return *o.AllowSimultaneousLogins
}

// GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetAllowSimultaneousLoginsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSimultaneousLogins) {
		return nil, false
	}
	return o.AllowSimultaneousLogins, true
}

// HasAllowSimultaneousLogins returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasAllowSimultaneousLogins() bool {
	if o != nil && !IsNil(o.AllowSimultaneousLogins) {
		return true
	}

	return false
}

// SetAllowSimultaneousLogins gets a reference to the given bool and assigns it to the AllowSimultaneousLogins field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetAllowSimultaneousLogins(v bool) {
	o.AllowSimultaneousLogins = &v
}

// GetGuestSponsorship returns the GuestSponsorship field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetGuestSponsorship() UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship {
	if o == nil || IsNil(o.GuestSponsorship) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship
		return ret
	}
	return *o.GuestSponsorship
}

// GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetGuestSponsorshipOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship, bool) {
	if o == nil || IsNil(o.GuestSponsorship) {
		return nil, false
	}
	return o.GuestSponsorship, true
}

// HasGuestSponsorship returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasGuestSponsorship() bool {
	if o != nil && !IsNil(o.GuestSponsorship) {
		return true
	}

	return false
}

// SetGuestSponsorship gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship and assigns it to the GuestSponsorship field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetGuestSponsorship(v UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) {
	o.GuestSponsorship = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBilling() UpdateNetworkWirelessSsidSplashSettingsRequestBilling {
	if o == nil || IsNil(o.Billing) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetBillingOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestBilling, bool) {
	if o == nil || IsNil(o.Billing) {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasBilling() bool {
	if o != nil && !IsNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestBilling and assigns it to the Billing field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetBilling(v UpdateNetworkWirelessSsidSplashSettingsRequestBilling) {
	o.Billing = &v
}

// GetSentryEnrollment returns the SentryEnrollment field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSentryEnrollment() UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment {
	if o == nil || IsNil(o.SentryEnrollment) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment
		return ret
	}
	return *o.SentryEnrollment
}

// GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) GetSentryEnrollmentOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment, bool) {
	if o == nil || IsNil(o.SentryEnrollment) {
		return nil, false
	}
	return o.SentryEnrollment, true
}

// HasSentryEnrollment returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) HasSentryEnrollment() bool {
	if o != nil && !IsNil(o.SentryEnrollment) {
		return true
	}

	return false
}

// SetSentryEnrollment gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment and assigns it to the SentryEnrollment field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequest) SetSentryEnrollment(v UpdateNetworkWirelessSsidSplashSettingsRequestSentryEnrollment) {
	o.SentryEnrollment = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SplashUrl) {
		toSerialize["splashUrl"] = o.SplashUrl
	}
	if !IsNil(o.UseSplashUrl) {
		toSerialize["useSplashUrl"] = o.UseSplashUrl
	}
	if !IsNil(o.SplashTimeout) {
		toSerialize["splashTimeout"] = o.SplashTimeout
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.UseRedirectUrl) {
		toSerialize["useRedirectUrl"] = o.UseRedirectUrl
	}
	if !IsNil(o.WelcomeMessage) {
		toSerialize["welcomeMessage"] = o.WelcomeMessage
	}
	if !IsNil(o.SplashLogo) {
		toSerialize["splashLogo"] = o.SplashLogo
	}
	if !IsNil(o.SplashImage) {
		toSerialize["splashImage"] = o.SplashImage
	}
	if !IsNil(o.SplashPrepaidFront) {
		toSerialize["splashPrepaidFront"] = o.SplashPrepaidFront
	}
	if !IsNil(o.BlockAllTrafficBeforeSignOn) {
		toSerialize["blockAllTrafficBeforeSignOn"] = o.BlockAllTrafficBeforeSignOn
	}
	if !IsNil(o.ControllerDisconnectionBehavior) {
		toSerialize["controllerDisconnectionBehavior"] = o.ControllerDisconnectionBehavior
	}
	if !IsNil(o.AllowSimultaneousLogins) {
		toSerialize["allowSimultaneousLogins"] = o.AllowSimultaneousLogins
	}
	if !IsNil(o.GuestSponsorship) {
		toSerialize["guestSponsorship"] = o.GuestSponsorship
	}
	if !IsNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}
	if !IsNil(o.SentryEnrollment) {
		toSerialize["sentryEnrollment"] = o.SentryEnrollment
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequest struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequest) Get() *UpdateNetworkWirelessSsidSplashSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequest) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequest(val *UpdateNetworkWirelessSsidSplashSettingsRequest) *NullableUpdateNetworkWirelessSsidSplashSettingsRequest {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



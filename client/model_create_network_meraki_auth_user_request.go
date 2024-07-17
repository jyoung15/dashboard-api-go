/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkMerakiAuthUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkMerakiAuthUserRequest{}

// CreateNetworkMerakiAuthUserRequest struct for CreateNetworkMerakiAuthUserRequest
type CreateNetworkMerakiAuthUserRequest struct {
	// Email address of the user
	Email string `json:"email"`
	// Name of the user. Only required If the user is not a Dashboard administrator.
	Name *string `json:"name,omitempty"`
	// The password for this user account. Only required If the user is not a Dashboard administrator.
	Password *string `json:"password,omitempty"`
	// Authorization type for user. Can be 'Guest' or '802.1X' for wireless networks, or 'Client VPN' for MX networks. Defaults to '802.1X'.
	AccountType *string `json:"accountType,omitempty"`
	// Whether or not Meraki should email the password to user. Default is false.
	EmailPasswordToUser *bool `json:"emailPasswordToUser,omitempty"`
	// Whether or not the user is a Dashboard administrator.
	IsAdmin *bool `json:"isAdmin,omitempty"`
	// Authorization zones and expiration dates for the user.
	Authorizations []CreateNetworkMerakiAuthUserRequestAuthorizationsInner `json:"authorizations"`
}

type _CreateNetworkMerakiAuthUserRequest CreateNetworkMerakiAuthUserRequest

// NewCreateNetworkMerakiAuthUserRequest instantiates a new CreateNetworkMerakiAuthUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkMerakiAuthUserRequest(email string, authorizations []CreateNetworkMerakiAuthUserRequestAuthorizationsInner) *CreateNetworkMerakiAuthUserRequest {
	this := CreateNetworkMerakiAuthUserRequest{}
	this.Email = email
	var accountType string = "802.1X"
	this.AccountType = &accountType
	this.Authorizations = authorizations
	return &this
}

// NewCreateNetworkMerakiAuthUserRequestWithDefaults instantiates a new CreateNetworkMerakiAuthUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkMerakiAuthUserRequestWithDefaults() *CreateNetworkMerakiAuthUserRequest {
	this := CreateNetworkMerakiAuthUserRequest{}
	var accountType string = "802.1X"
	this.AccountType = &accountType
	return &this
}

// GetEmail returns the Email field value
func (o *CreateNetworkMerakiAuthUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateNetworkMerakiAuthUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkMerakiAuthUserRequest) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateNetworkMerakiAuthUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequest) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequest) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *CreateNetworkMerakiAuthUserRequest) SetAccountType(v string) {
	o.AccountType = &v
}

// GetEmailPasswordToUser returns the EmailPasswordToUser field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequest) GetEmailPasswordToUser() bool {
	if o == nil || IsNil(o.EmailPasswordToUser) {
		var ret bool
		return ret
	}
	return *o.EmailPasswordToUser
}

// GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetEmailPasswordToUserOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailPasswordToUser) {
		return nil, false
	}
	return o.EmailPasswordToUser, true
}

// HasEmailPasswordToUser returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequest) HasEmailPasswordToUser() bool {
	if o != nil && !IsNil(o.EmailPasswordToUser) {
		return true
	}

	return false
}

// SetEmailPasswordToUser gets a reference to the given bool and assigns it to the EmailPasswordToUser field.
func (o *CreateNetworkMerakiAuthUserRequest) SetEmailPasswordToUser(v bool) {
	o.EmailPasswordToUser = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequest) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequest) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *CreateNetworkMerakiAuthUserRequest) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetAuthorizations returns the Authorizations field value
func (o *CreateNetworkMerakiAuthUserRequest) GetAuthorizations() []CreateNetworkMerakiAuthUserRequestAuthorizationsInner {
	if o == nil {
		var ret []CreateNetworkMerakiAuthUserRequestAuthorizationsInner
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequest) GetAuthorizationsOk() ([]CreateNetworkMerakiAuthUserRequestAuthorizationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *CreateNetworkMerakiAuthUserRequest) SetAuthorizations(v []CreateNetworkMerakiAuthUserRequestAuthorizationsInner) {
	o.Authorizations = v
}

func (o CreateNetworkMerakiAuthUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkMerakiAuthUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.EmailPasswordToUser) {
		toSerialize["emailPasswordToUser"] = o.EmailPasswordToUser
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	toSerialize["authorizations"] = o.Authorizations
	return toSerialize, nil
}

func (o *CreateNetworkMerakiAuthUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"authorizations",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNetworkMerakiAuthUserRequest := _CreateNetworkMerakiAuthUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkMerakiAuthUserRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkMerakiAuthUserRequest(varCreateNetworkMerakiAuthUserRequest)

	return err
}

type NullableCreateNetworkMerakiAuthUserRequest struct {
	value *CreateNetworkMerakiAuthUserRequest
	isSet bool
}

func (v NullableCreateNetworkMerakiAuthUserRequest) Get() *CreateNetworkMerakiAuthUserRequest {
	return v.value
}

func (v *NullableCreateNetworkMerakiAuthUserRequest) Set(val *CreateNetworkMerakiAuthUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkMerakiAuthUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkMerakiAuthUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkMerakiAuthUserRequest(val *CreateNetworkMerakiAuthUserRequest) *NullableCreateNetworkMerakiAuthUserRequest {
	return &NullableCreateNetworkMerakiAuthUserRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkMerakiAuthUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkMerakiAuthUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



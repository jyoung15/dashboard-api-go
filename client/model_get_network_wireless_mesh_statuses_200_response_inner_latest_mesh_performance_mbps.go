/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps - Average Mbps.
type GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps struct {
	Int32 *int32
	String *string
}

// int32AsGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps is a convenience function that returns int32 wrapped in GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps
func Int32AsGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps(v *int32) GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps {
	return GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps{
		Int32: v,
	}
}

// stringAsGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps is a convenience function that returns string wrapped in GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps
func StringAsGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps(v *string) GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps {
	return GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps struct {
	value *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps
	isSet bool
}

func (v NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) Get() *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps {
	return v.value
}

func (v *NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) Set(val *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps(val *GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) *NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps {
	return &NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformanceMbps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



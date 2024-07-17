# GetDeviceLldpCdp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceMac** | Pointer to **string** | Source MAC address | [optional] 
**Ports** | Pointer to [**map[string]GetDeviceLldpCdp200ResponsePortsValue**](GetDeviceLldpCdp200ResponsePortsValue.md) | Mapping of ports to lldp and/or cdp information | [optional] 

## Methods

### NewGetDeviceLldpCdp200Response

`func NewGetDeviceLldpCdp200Response() *GetDeviceLldpCdp200Response`

NewGetDeviceLldpCdp200Response instantiates a new GetDeviceLldpCdp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceLldpCdp200ResponseWithDefaults

`func NewGetDeviceLldpCdp200ResponseWithDefaults() *GetDeviceLldpCdp200Response`

NewGetDeviceLldpCdp200ResponseWithDefaults instantiates a new GetDeviceLldpCdp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceMac

`func (o *GetDeviceLldpCdp200Response) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *GetDeviceLldpCdp200Response) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *GetDeviceLldpCdp200Response) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *GetDeviceLldpCdp200Response) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetPorts

`func (o *GetDeviceLldpCdp200Response) GetPorts() map[string]GetDeviceLldpCdp200ResponsePortsValue`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetDeviceLldpCdp200Response) GetPortsOk() (*map[string]GetDeviceLldpCdp200ResponsePortsValue, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetDeviceLldpCdp200Response) SetPorts(v map[string]GetDeviceLldpCdp200ResponsePortsValue)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetDeviceLldpCdp200Response) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
Meraki Dashboard API

Testing DeviceProfilesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/meraki/dashboard-api-go/client"
)

func Test_client_DeviceProfilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceProfilesAPIService GetNetworkSmDeviceDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.DeviceProfilesAPI.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceProfilesAPIService GetNetworkSmUserDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var userId string

		resp, httpRes, err := apiClient.DeviceProfilesAPI.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

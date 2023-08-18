/*
Meraki Dashboard API

Testing WirelessProfilesAPIService

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

func Test_client_WirelessProfilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WirelessProfilesAPIService CreateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.CreateNetworkCameraWirelessProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService DeleteNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		httpRes, err := apiClient.WirelessProfilesAPI.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService GetDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService GetNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService GetNetworkCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService UpdateDeviceCameraWirelessProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WirelessProfilesAPIService UpdateNetworkCameraWirelessProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var wirelessProfileId string

		resp, httpRes, err := apiClient.WirelessProfilesAPI.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

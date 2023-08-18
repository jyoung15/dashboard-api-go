/*
Meraki Dashboard API

Testing FirmwareUpgradesAPIService

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

func Test_client_FirmwareUpgradesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirmwareUpgradesAPIService CreateNetworkFirmwareUpgradesRollback", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService CreateNetworkFirmwareUpgradesStagedEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService CreateNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService DeferNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService DeleteNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		httpRes, err := apiClient.FirmwareUpgradesAPI.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService GetNetworkFirmwareUpgrades", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.GetNetworkFirmwareUpgrades(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService GetNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService GetNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService GetNetworkFirmwareUpgradesStagedGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService GetNetworkFirmwareUpgradesStagedStages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService RollbacksNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService UpdateNetworkFirmwareUpgrades", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.UpdateNetworkFirmwareUpgrades(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService UpdateNetworkFirmwareUpgradesStagedEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService UpdateNetworkFirmwareUpgradesStagedGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var groupId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareUpgradesAPIService UpdateNetworkFirmwareUpgradesStagedStages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.FirmwareUpgradesAPI.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

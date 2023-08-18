/*
Meraki Dashboard API

Testing InventoryAPIService

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

func Test_client_InventoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InventoryAPIService ClaimIntoOrganizationInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.ClaimIntoOrganizationInventory(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService CreateOrganizationInventoryOnboardingCloudMonitoringImport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService CreateOrganizationInventoryOnboardingCloudMonitoringPrepare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService GetOrganizationInventoryDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var serial string

		resp, httpRes, err := apiClient.InventoryAPI.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService GetOrganizationInventoryDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.GetOrganizationInventoryDevices(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService GetOrganizationInventoryOnboardingCloudMonitoringImports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService GetOrganizationInventoryOnboardingCloudMonitoringNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryAPIService ReleaseFromOrganizationInventory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InventoryAPI.ReleaseFromOrganizationInventory(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

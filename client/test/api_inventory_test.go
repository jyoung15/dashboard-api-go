/*
Meraki Dashboard API

Testing InventoryApiService

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

func Test_client_InventoryApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test InventoryApiService ClaimIntoOrganizationInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.ClaimIntoOrganizationInventory(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService CreateOrganizationInventoryOnboardingCloudMonitoringImport", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService CreateOrganizationInventoryOnboardingCloudMonitoringPrepare", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryDevice", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var serial string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryDevices(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryOnboardingCloudMonitoringImports", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryOnboardingCloudMonitoringNetworks", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService ReleaseFromOrganizationInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.ReleaseFromOrganizationInventory(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

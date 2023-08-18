/*
Meraki Dashboard API

Testing TrafficShapingAPIService

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

func Test_client_TrafficShapingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TrafficShapingAPIService CreateNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService DeleteNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		httpRes, err := apiClient.TrafficShapingAPI.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShaping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShaping(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShapingCustomPerformanceClasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShapingUplinkBandwidth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkApplianceTrafficShapingUplinkSelection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkTrafficShapingApplicationCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkTrafficShapingApplicationCategories(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkTrafficShapingDscpTaggingOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkTrafficShapingDscpTaggingOptions(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService GetNetworkWirelessSsidTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.TrafficShapingAPI.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkApplianceTrafficShaping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkApplianceTrafficShaping(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkApplianceTrafficShapingCustomPerformanceClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var customPerformanceClassId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkApplianceTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkApplianceTrafficShapingUplinkBandwidth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkApplianceTrafficShapingUplinkSelection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TrafficShapingAPIService UpdateNetworkWirelessSsidTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.TrafficShapingAPI.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

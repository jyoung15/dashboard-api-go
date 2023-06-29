/*
Meraki Dashboard API

Testing DhcpServerPolicyApiService

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

func Test_client_DhcpServerPolicyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DhcpServerPolicyApiService CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var trustedServerId string

		httpRes, err := apiClient.DhcpServerPolicyApi.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService GetNetworkSwitchDhcpServerPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService UpdateNetworkSwitchDhcpServerPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DhcpServerPolicyApiService UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var trustedServerId string

		resp, httpRes, err := apiClient.DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

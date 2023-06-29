/*
Meraki Dashboard API

Testing EapOverrideApiService

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

func Test_client_EapOverrideApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EapOverrideApiService GetNetworkWirelessSsidEapOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.EapOverrideApi.GetNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EapOverrideApiService UpdateNetworkWirelessSsidEapOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.EapOverrideApi.UpdateNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

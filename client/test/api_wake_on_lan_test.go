/*
Meraki Dashboard API

Testing WakeOnLanAPIService

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

func Test_client_WakeOnLanAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WakeOnLanAPIService CreateDeviceLiveToolsWakeOnLan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.WakeOnLanAPI.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WakeOnLanAPIService GetDeviceLiveToolsWakeOnLan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var wakeOnLanId string

		resp, httpRes, err := apiClient.WakeOnLanAPI.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

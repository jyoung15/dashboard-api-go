/*
Meraki Dashboard API

Testing CommandsAPIService

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

func Test_client_CommandsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CommandsAPIService CreateDeviceSensorCommand", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CommandsAPI.CreateDeviceSensorCommand(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommandsAPIService GetDeviceSensorCommand", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string
		var commandId string

		resp, httpRes, err := apiClient.CommandsAPI.GetDeviceSensorCommand(context.Background(), serial, commandId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommandsAPIService GetDeviceSensorCommands", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.CommandsAPI.GetDeviceSensorCommands(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

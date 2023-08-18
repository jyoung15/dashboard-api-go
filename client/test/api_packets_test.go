/*
Meraki Dashboard API

Testing PacketsAPIService

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

func Test_client_PacketsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PacketsAPIService GetDeviceSwitchPortsStatusesPackets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.PacketsAPI.GetDeviceSwitchPortsStatusesPackets(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

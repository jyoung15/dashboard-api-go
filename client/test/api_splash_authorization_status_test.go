/*
Meraki Dashboard API

Testing SplashAuthorizationStatusAPIService

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

func Test_client_SplashAuthorizationStatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SplashAuthorizationStatusAPIService GetNetworkClientSplashAuthorizationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var clientId string

		resp, httpRes, err := apiClient.SplashAuthorizationStatusAPI.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SplashAuthorizationStatusAPIService UpdateNetworkClientSplashAuthorizationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var clientId string

		resp, httpRes, err := apiClient.SplashAuthorizationStatusAPI.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

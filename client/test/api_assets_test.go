/*
Meraki Dashboard API

Testing AssetsAPIService

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

func Test_client_AssetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssetsAPIService CreateOrganizationSplashThemeAsset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var themeIdentifier string

		resp, httpRes, err := apiClient.AssetsAPI.CreateOrganizationSplashThemeAsset(context.Background(), organizationId, themeIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService DeleteOrganizationSplashAsset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var id string

		httpRes, err := apiClient.AssetsAPI.DeleteOrganizationSplashAsset(context.Background(), organizationId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService GetOrganizationSplashAsset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var id string

		resp, httpRes, err := apiClient.AssetsAPI.GetOrganizationSplashAsset(context.Background(), organizationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

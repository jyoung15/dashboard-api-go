/*
Meraki Dashboard API

Testing ApplicationsApiService

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

func Test_client_ApplicationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsApiService GetNetworkInsightApplicationHealthByTime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsApi.GetNetworkInsightApplicationHealthByTime(context.Background(), networkId, applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsApiService GetOrganizationInsightApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationsApi.GetOrganizationInsightApplications(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

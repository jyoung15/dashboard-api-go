/*
Meraki Dashboard API

Testing CategoriesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_client_CategoriesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CategoriesApiService GetNetworkApplianceContentFilteringCategories", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.CategoriesApi.GetNetworkApplianceContentFilteringCategories(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
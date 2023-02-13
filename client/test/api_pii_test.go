/*
Meraki Dashboard API

Testing PiiApiService

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

func Test_client_PiiApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PiiApiService CreateNetworkPiiRequest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.PiiApi.CreateNetworkPiiRequest(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService DeleteNetworkPiiRequest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var requestId string

        resp, httpRes, err := apiClient.PiiApi.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService GetNetworkPiiPiiKeys", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.PiiApi.GetNetworkPiiPiiKeys(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService GetNetworkPiiRequest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var requestId string

        resp, httpRes, err := apiClient.PiiApi.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService GetNetworkPiiRequests", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.PiiApi.GetNetworkPiiRequests(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService GetNetworkPiiSmDevicesForKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.PiiApi.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PiiApiService GetNetworkPiiSmOwnersForKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.PiiApi.GetNetworkPiiSmOwnersForKey(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
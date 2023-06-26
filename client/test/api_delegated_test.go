/*
Meraki Dashboard API

Testing DelegatedApiService

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

func Test_client_DelegatedApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DelegatedApiService CreateNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.DelegatedApi.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService DeleteNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.DelegatedApi.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService GetDeviceAppliancePrefixesDelegated", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.DelegatedApi.GetDeviceAppliancePrefixesDelegated(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService GetDeviceAppliancePrefixesDelegatedVlanAssignments", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.DelegatedApi.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService GetNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.DelegatedApi.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService GetNetworkAppliancePrefixesDelegatedStatics", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.DelegatedApi.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DelegatedApiService UpdateNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.DelegatedApi.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

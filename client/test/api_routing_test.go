/*
Meraki Dashboard API

Testing RoutingApiService

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

func Test_client_RoutingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RoutingApiService CreateDeviceSwitchRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.RoutingApi.CreateDeviceSwitchRoutingInterface(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService CreateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.RoutingApi.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService CreateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService CreateNetworkSwitchStackRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.RoutingApi.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService CreateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.RoutingApi.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService DeleteDeviceSwitchRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService DeleteDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService DeleteNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var rendezvousPointId string

        resp, httpRes, err := apiClient.RoutingApi.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService DeleteNetworkSwitchStackRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService DeleteNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetDeviceSwitchRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetDeviceSwitchRoutingInterfaceDhcp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetDeviceSwitchRoutingInterfaces", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetDeviceSwitchRoutingStaticRoutes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.RoutingApi.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchRoutingMulticast", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var rendezvousPointId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchRoutingMulticastRendezvousPoints", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchRoutingOspf", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchStackRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchStackRoutingInterfaceDhcp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchStackRoutingInterfaces", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService GetNetworkSwitchStackRoutingStaticRoutes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateDeviceSwitchRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateDeviceSwitchRoutingInterfaceDhcp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchRoutingMulticast", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchRoutingMulticastRendezvousPoint", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var rendezvousPointId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchRoutingOspf", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchStackRoutingInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchStackRoutingInterfaceDhcp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var interfaceId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RoutingApiService UpdateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

/*
Meraki Dashboard API

Testing ThirdPartyVPNPeersAPIService

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

func Test_client_ThirdPartyVPNPeersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThirdPartyVPNPeersAPIService GetOrganizationApplianceVpnThirdPartyVPNPeers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ThirdPartyVPNPeersAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThirdPartyVPNPeersAPIService UpdateOrganizationApplianceVpnThirdPartyVPNPeers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ThirdPartyVPNPeersAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

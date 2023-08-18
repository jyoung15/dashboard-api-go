/*
Meraki Dashboard API

Testing ConfigTemplatesAPIService

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

func Test_client_ConfigTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigTemplatesAPIService CreateOrganizationConfigTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.CreateOrganizationConfigTemplate(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService DeleteOrganizationConfigTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		httpRes, err := apiClient.ConfigTemplatesAPI.DeleteOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService GetOrganizationConfigTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService GetOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService GetOrganizationConfigTemplateSwitchProfilePorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService GetOrganizationConfigTemplateSwitchProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService GetOrganizationConfigTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.GetOrganizationConfigTemplates(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService UpdateOrganizationConfigTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.UpdateOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigTemplatesAPIService UpdateOrganizationConfigTemplateSwitchProfilePort", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var configTemplateId string
		var profileId string
		var portId string

		resp, httpRes, err := apiClient.ConfigTemplatesAPI.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

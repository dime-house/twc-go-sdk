/*
Timeweb Cloud API

Testing DedicatedServersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/dime-house/twc-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_DedicatedServersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DedicatedServersAPIService CreateDedicatedServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedServersAPI.CreateDedicatedServer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService DeleteDedicatedServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dedicatedId int32

		httpRes, err := apiClient.DedicatedServersAPI.DeleteDedicatedServer(context.Background(), dedicatedId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService GetDedicatedServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dedicatedId int32

		resp, httpRes, err := apiClient.DedicatedServersAPI.GetDedicatedServer(context.Background(), dedicatedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService GetDedicatedServerPresetAdditionalServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var presetId int32

		resp, httpRes, err := apiClient.DedicatedServersAPI.GetDedicatedServerPresetAdditionalServices(context.Background(), presetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService GetDedicatedServers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedServersAPI.GetDedicatedServers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService GetDedicatedServersPresets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DedicatedServersAPI.GetDedicatedServersPresets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DedicatedServersAPIService UpdateDedicatedServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dedicatedId int32

		resp, httpRes, err := apiClient.DedicatedServersAPI.UpdateDedicatedServer(context.Background(), dedicatedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

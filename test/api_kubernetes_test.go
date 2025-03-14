/*
Timeweb Cloud API

Testing KubernetesAPIService

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

func Test_openapi_KubernetesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubernetesAPIService CreateCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.CreateCluster(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService CreateClusterNodeGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.CreateClusterNodeGroup(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService DeleteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.DeleteCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService DeleteClusterNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var nodeId int32

		httpRes, err := apiClient.KubernetesAPI.DeleteClusterNode(context.Background(), clusterId, nodeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService DeleteClusterNodeGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var groupId int32

		httpRes, err := apiClient.KubernetesAPI.DeleteClusterNodeGroup(context.Background(), clusterId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterKubeconfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterKubeconfig(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterNodeGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var groupId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterNodeGroup(context.Background(), clusterId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterNodeGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterNodeGroups(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterNodes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterNodes(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterNodesFromGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var groupId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterNodesFromGroup(context.Background(), clusterId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusterResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusterResources(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetClusters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.GetClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetK8SNetworkDrivers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.GetK8SNetworkDrivers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetK8SVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.GetK8SVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService GetKubernetesPresets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.GetKubernetesPresets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService IncreaseCountOfNodesInGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var groupId int32

		resp, httpRes, err := apiClient.KubernetesAPI.IncreaseCountOfNodesInGroup(context.Background(), clusterId, groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService ReduceCountOfNodesInGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32
		var groupId int32

		httpRes, err := apiClient.KubernetesAPI.ReduceCountOfNodesInGroup(context.Background(), clusterId, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService UpdateCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId int32

		resp, httpRes, err := apiClient.KubernetesAPI.UpdateCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RecoveryPointsRecommendedForMoveClient contains the methods for the RecoveryPointsRecommendedForMove group.
// Don't use this type directly, use NewRecoveryPointsRecommendedForMoveClient() instead.
type RecoveryPointsRecommendedForMoveClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRecoveryPointsRecommendedForMoveClient creates a new instance of RecoveryPointsRecommendedForMoveClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoveryPointsRecommendedForMoveClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoveryPointsRecommendedForMoveClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RecoveryPointsRecommendedForMoveClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Lists the recovery points recommended for move to another tier
//
// Generated from API version 2022-10-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - List Recovery points Recommended for Move Request
//   - options - RecoveryPointsRecommendedForMoveClientListOptions contains the optional parameters for the RecoveryPointsRecommendedForMoveClient.NewListPager
//     method.
func (client *RecoveryPointsRecommendedForMoveClient) NewListPager(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest, options *RecoveryPointsRecommendedForMoveClientListOptions) *runtime.Pager[RecoveryPointsRecommendedForMoveClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoveryPointsRecommendedForMoveClientListResponse]{
		More: func(page RecoveryPointsRecommendedForMoveClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoveryPointsRecommendedForMoveClientListResponse) (RecoveryPointsRecommendedForMoveClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, parameters, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RecoveryPointsRecommendedForMoveClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RecoveryPointsRecommendedForMoveClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RecoveryPointsRecommendedForMoveClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RecoveryPointsRecommendedForMoveClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest, options *RecoveryPointsRecommendedForMoveClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPointsRecommendedForMove"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listHandleResponse handles the List response.
func (client *RecoveryPointsRecommendedForMoveClient) listHandleResponse(resp *http.Response) (RecoveryPointsRecommendedForMoveClientListResponse, error) {
	result := RecoveryPointsRecommendedForMoveClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointResourceList); err != nil {
		return RecoveryPointsRecommendedForMoveClientListResponse{}, err
	}
	return result, nil
}

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RecoveryPointsRecommendedForMoveClient contains the methods for the RecoveryPointsRecommendedForMove group.
// Don't use this type directly, use NewRecoveryPointsRecommendedForMoveClient() instead.
type RecoveryPointsRecommendedForMoveClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecoveryPointsRecommendedForMoveClient creates a new instance of RecoveryPointsRecommendedForMoveClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoveryPointsRecommendedForMoveClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoveryPointsRecommendedForMoveClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoveryPointsRecommendedForMoveClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists the recovery points recommended for move to another tier
//
// Generated from API version 2025-02-01
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
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecoveryPointsRecommendedForMoveClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, parameters, options)
			}, nil)
			if err != nil {
				return RecoveryPointsRecommendedForMoveClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RecoveryPointsRecommendedForMoveClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters ListRecoveryPointsRecommendedForMoveRequest, _ *RecoveryPointsRecommendedForMoveClientListOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecoveryPointsRecommendedForMoveClient) listHandleResponse(resp *http.Response) (RecoveryPointsRecommendedForMoveClientListResponse, error) {
	result := RecoveryPointsRecommendedForMoveClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointResourceList); err != nil {
		return RecoveryPointsRecommendedForMoveClientListResponse{}, err
	}
	return result, nil
}

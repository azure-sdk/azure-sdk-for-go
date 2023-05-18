//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// ManagedDatabaseRecommendedSensitivityLabelsClient contains the methods for the ManagedDatabaseRecommendedSensitivityLabels group.
// Don't use this type directly, use NewManagedDatabaseRecommendedSensitivityLabelsClient() instead.
type ManagedDatabaseRecommendedSensitivityLabelsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseRecommendedSensitivityLabelsClient creates a new instance of ManagedDatabaseRecommendedSensitivityLabelsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseRecommendedSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseRecommendedSensitivityLabelsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedDatabaseRecommendedSensitivityLabelsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseRecommendedSensitivityLabelsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Update - Update recommended sensitivity labels states of a given database using an operations batch.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedDatabaseRecommendedSensitivityLabelsClientUpdateOptions contains the optional parameters for the ManagedDatabaseRecommendedSensitivityLabelsClient.Update
//     method.
func (client *ManagedDatabaseRecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *ManagedDatabaseRecommendedSensitivityLabelsClientUpdateOptions) (ManagedDatabaseRecommendedSensitivityLabelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
	if err != nil {
		return ManagedDatabaseRecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseRecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseRecommendedSensitivityLabelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedDatabaseRecommendedSensitivityLabelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedDatabaseRecommendedSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *ManagedDatabaseRecommendedSensitivityLabelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/recommendedSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

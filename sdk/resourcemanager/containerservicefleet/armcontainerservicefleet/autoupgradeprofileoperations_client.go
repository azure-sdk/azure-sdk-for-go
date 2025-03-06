// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

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

// AutoUpgradeProfileOperationsClient contains the methods for the AutoUpgradeProfileOperations group.
// Don't use this type directly, use NewAutoUpgradeProfileOperationsClient() instead.
type AutoUpgradeProfileOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAutoUpgradeProfileOperationsClient creates a new instance of AutoUpgradeProfileOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAutoUpgradeProfileOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AutoUpgradeProfileOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AutoUpgradeProfileOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginGenerateUpdateRun - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - autoUpgradeProfileName - The name of the AutoUpgradeProfile resource.
//   - options - AutoUpgradeProfileOperationsClientBeginGenerateUpdateRunOptions contains the optional parameters for the AutoUpgradeProfileOperationsClient.BeginGenerateUpdateRun
//     method.
func (client *AutoUpgradeProfileOperationsClient) BeginGenerateUpdateRun(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfileOperationsClientBeginGenerateUpdateRunOptions) (*runtime.Poller[AutoUpgradeProfileOperationsClientGenerateUpdateRunResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateUpdateRun(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AutoUpgradeProfileOperationsClientGenerateUpdateRunResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AutoUpgradeProfileOperationsClientGenerateUpdateRunResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GenerateUpdateRun - A long-running resource action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *AutoUpgradeProfileOperationsClient) generateUpdateRun(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, options *AutoUpgradeProfileOperationsClientBeginGenerateUpdateRunOptions) (*http.Response, error) {
	var err error
	const operationName = "AutoUpgradeProfileOperationsClient.BeginGenerateUpdateRun"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateUpdateRunCreateRequest(ctx, resourceGroupName, fleetName, autoUpgradeProfileName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// generateUpdateRunCreateRequest creates the GenerateUpdateRun request.
func (client *AutoUpgradeProfileOperationsClient) generateUpdateRunCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, autoUpgradeProfileName string, _ *AutoUpgradeProfileOperationsClientBeginGenerateUpdateRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/autoUpgradeProfiles/{autoUpgradeProfileName}/generateUpdateRun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if autoUpgradeProfileName == "" {
		return nil, errors.New("parameter autoUpgradeProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{autoUpgradeProfileName}", url.PathEscape(autoUpgradeProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

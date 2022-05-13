//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ValidateOperationClient contains the methods for the ValidateOperation group.
// Don't use this type directly, use NewValidateOperationClient() instead.
type ValidateOperationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewValidateOperationClient creates a new instance of ValidateOperationClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewValidateOperationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ValidateOperationClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ValidateOperationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginTrigger - Validate operation for specified backed up item in the form of an asynchronous operation. Returns tracking
// headers which can be tracked using GetValidateOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - resource validate operation request
// options - ValidateOperationClientBeginTriggerOptions contains the optional parameters for the ValidateOperationClient.BeginTrigger
// method.
func (client *ValidateOperationClient) BeginTrigger(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateOperationRequestClassification, options *ValidateOperationClientBeginTriggerOptions) (*armruntime.Poller[ValidateOperationClientTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.trigger(ctx, vaultName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[ValidateOperationClientTriggerResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[ValidateOperationClientTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// Trigger - Validate operation for specified backed up item in the form of an asynchronous operation. Returns tracking headers
// which can be tracked using GetValidateOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ValidateOperationClient) trigger(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateOperationRequestClassification, options *ValidateOperationClientBeginTriggerOptions) (*http.Response, error) {
	req, err := client.triggerCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *ValidateOperationClient) triggerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateOperationRequestClassification, options *ValidateOperationClientBeginTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupTriggerValidateOperation"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

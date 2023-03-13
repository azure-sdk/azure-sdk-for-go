//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// ManagementGroupGovernanceRulesExecuteStatusClient contains the methods for the ManagementGroupGovernanceRulesExecuteStatus group.
// Don't use this type directly, use NewManagementGroupGovernanceRulesExecuteStatusClient() instead.
type ManagementGroupGovernanceRulesExecuteStatusClient struct {
	host              string
	managementGroupID string
	pl                runtime.Pipeline
}

// NewManagementGroupGovernanceRulesExecuteStatusClient creates a new instance of ManagementGroupGovernanceRulesExecuteStatusClient with the specified values.
//   - managementGroupID - Azure Management Group ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementGroupGovernanceRulesExecuteStatusClient(managementGroupID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementGroupGovernanceRulesExecuteStatusClient, error) {
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
	client := &ManagementGroupGovernanceRulesExecuteStatusClient{
		managementGroupID: managementGroupID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Get a specific governance rule execution status for the requested scope by ruleId and operationId
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - ruleID - The governance rule key - unique key for the standard governance rule (GUID)
//   - operationID - The governance rule execution key - unique key for the execution of governance rule
//   - options - ManagementGroupGovernanceRulesExecuteStatusClientGetOptions contains the optional parameters for the ManagementGroupGovernanceRulesExecuteStatusClient.Get
//     method.
func (client *ManagementGroupGovernanceRulesExecuteStatusClient) Get(ctx context.Context, ruleID string, operationID string, options *ManagementGroupGovernanceRulesExecuteStatusClientGetOptions) (ManagementGroupGovernanceRulesExecuteStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, ruleID, operationID, options)
	if err != nil {
		return ManagementGroupGovernanceRulesExecuteStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupGovernanceRulesExecuteStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return ManagementGroupGovernanceRulesExecuteStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagementGroupGovernanceRulesExecuteStatusClient) getCreateRequest(ctx context.Context, ruleID string, operationID string, options *ManagementGroupGovernanceRulesExecuteStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Security/governanceRules/{ruleId}/execute/operationResults/{operationId}"
	if client.managementGroupID == "" {
		return nil, errors.New("parameter client.managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(client.managementGroupID))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementGroupGovernanceRulesExecuteStatusClient) getHandleResponse(resp *http.Response) (ManagementGroupGovernanceRulesExecuteStatusClientGetResponse, error) {
	result := ManagementGroupGovernanceRulesExecuteStatusClientGetResponse{}
	if val := resp.Header.Get("location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExecuteRuleStatus); err != nil {
		return ManagementGroupGovernanceRulesExecuteStatusClientGetResponse{}, err
	}
	return result, nil
}

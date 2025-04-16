// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// StatisticsClient contains the methods for the Statistics group.
// Don't use this type directly, use NewStatisticsClient() instead.
type StatisticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStatisticsClient creates a new instance of StatisticsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStatisticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StatisticsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StatisticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByAutomationAccountPager - Retrieve the statistics for the account.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - StatisticsClientListByAutomationAccountOptions contains the optional parameters for the StatisticsClient.NewListByAutomationAccountPager
//     method.
func (client *StatisticsClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *StatisticsClientListByAutomationAccountOptions) *runtime.Pager[StatisticsClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[StatisticsClientListByAutomationAccountResponse]{
		More: func(page StatisticsClientListByAutomationAccountResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *StatisticsClientListByAutomationAccountResponse) (StatisticsClientListByAutomationAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StatisticsClient.NewListByAutomationAccountPager")
			req, err := client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			if err != nil {
				return StatisticsClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StatisticsClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StatisticsClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *StatisticsClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *StatisticsClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/statistics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *StatisticsClient) listByAutomationAccountHandleResponse(resp *http.Response) (StatisticsClientListByAutomationAccountResponse, error) {
	result := StatisticsClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StatisticsListResult); err != nil {
		return StatisticsClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

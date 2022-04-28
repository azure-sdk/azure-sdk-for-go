//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

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

// HCIReportsClient contains the methods for the HCIReports group.
// Don't use this type directly, use NewHCIReportsClient() instead.
type HCIReportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHCIReportsClient creates a new instance of HCIReportsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHCIReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HCIReportsClient, error) {
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
	client := &HCIReportsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get information about a report associated with a configuration profile assignment run
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Arc machine.
// configurationProfileAssignmentName - The configuration profile assignment name.
// reportName - The report name.
// options - HCIReportsClientGetOptions contains the optional parameters for the HCIReportsClient.Get method.
func (client *HCIReportsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string, options *HCIReportsClientGetOptions) (HCIReportsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, reportName, options)
	if err != nil {
		return HCIReportsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HCIReportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HCIReportsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HCIReportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string, options *HCIReportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports/{reportName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HCIReportsClient) getHandleResponse(resp *http.Response) (HCIReportsClientGetResponse, error) {
	result := HCIReportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Report); err != nil {
		return HCIReportsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConfigurationProfileAssignmentsPager - Retrieve a list of reports within a given configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Arc machine.
// configurationProfileAssignmentName - The configuration profile assignment name.
// options - HCIReportsClientListByConfigurationProfileAssignmentsOptions contains the optional parameters for the HCIReportsClient.ListByConfigurationProfileAssignments
// method.
func (client *HCIReportsClient) NewListByConfigurationProfileAssignmentsPager(resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *HCIReportsClientListByConfigurationProfileAssignmentsOptions) *runtime.Pager[HCIReportsClientListByConfigurationProfileAssignmentsResponse] {
	return runtime.NewPager(runtime.PageProcessor[HCIReportsClientListByConfigurationProfileAssignmentsResponse]{
		More: func(page HCIReportsClientListByConfigurationProfileAssignmentsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *HCIReportsClientListByConfigurationProfileAssignmentsResponse) (HCIReportsClientListByConfigurationProfileAssignmentsResponse, error) {
			req, err := client.listByConfigurationProfileAssignmentsCreateRequest(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, options)
			if err != nil {
				return HCIReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HCIReportsClientListByConfigurationProfileAssignmentsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HCIReportsClientListByConfigurationProfileAssignmentsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationProfileAssignmentsHandleResponse(resp)
		},
	})
}

// listByConfigurationProfileAssignmentsCreateRequest creates the ListByConfigurationProfileAssignments request.
func (client *HCIReportsClient) listByConfigurationProfileAssignmentsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *HCIReportsClientListByConfigurationProfileAssignmentsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}/reports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByConfigurationProfileAssignmentsHandleResponse handles the ListByConfigurationProfileAssignments response.
func (client *HCIReportsClient) listByConfigurationProfileAssignmentsHandleResponse(resp *http.Response) (HCIReportsClientListByConfigurationProfileAssignmentsResponse, error) {
	result := HCIReportsClientListByConfigurationProfileAssignmentsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportList); err != nil {
		return HCIReportsClientListByConfigurationProfileAssignmentsResponse{}, err
	}
	return result, nil
}

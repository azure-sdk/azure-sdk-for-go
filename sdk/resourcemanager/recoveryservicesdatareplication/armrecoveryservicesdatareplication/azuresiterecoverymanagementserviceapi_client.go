//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

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

// AzureSiteRecoveryManagementServiceAPIClient contains the methods for the AzureSiteRecoveryManagementServiceAPI group.
// Don't use this type directly, use NewAzureSiteRecoveryManagementServiceAPIClient() instead.
type AzureSiteRecoveryManagementServiceAPIClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureSiteRecoveryManagementServiceAPIClient creates a new instance of AzureSiteRecoveryManagementServiceAPIClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureSiteRecoveryManagementServiceAPIClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureSiteRecoveryManagementServiceAPIClient, error) {
	cl, err := arm.NewClient(moduleName+".AzureSiteRecoveryManagementServiceAPIClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureSiteRecoveryManagementServiceAPIClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks the resource name availability.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - location - The name of the Azure region.
//   - options - AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityOptions contains the optional parameters for
//     the AzureSiteRecoveryManagementServiceAPIClient.CheckNameAvailability method.
func (client *AzureSiteRecoveryManagementServiceAPIClient) CheckNameAvailability(ctx context.Context, location string, options *AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityOptions) (AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse, error) {
	var err error
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, options)
	if err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *AzureSiteRecoveryManagementServiceAPIClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, options *AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataReplication/locations/{location}/checkNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *AzureSiteRecoveryManagementServiceAPIClient) checkNameAvailabilityHandleResponse(resp *http.Response) (AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse, error) {
	result := AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponseModel); err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// DeploymentPreflight - Performs resource deployment validation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentID - Deployment Id.
//   - options - AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightOptions contains the optional parameters for the
//     AzureSiteRecoveryManagementServiceAPIClient.DeploymentPreflight method.
func (client *AzureSiteRecoveryManagementServiceAPIClient) DeploymentPreflight(ctx context.Context, resourceGroupName string, deploymentID string, options *AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightOptions) (AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse, error) {
	var err error
	req, err := client.deploymentPreflightCreateRequest(ctx, resourceGroupName, deploymentID, options)
	if err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse{}, err
	}
	resp, err := client.deploymentPreflightHandleResponse(httpResp)
	return resp, err
}

// deploymentPreflightCreateRequest creates the DeploymentPreflight request.
func (client *AzureSiteRecoveryManagementServiceAPIClient) deploymentPreflightCreateRequest(ctx context.Context, resourceGroupName string, deploymentID string, options *AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/deployments/{deploymentId}/preflight"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentID == "" {
		return nil, errors.New("parameter deploymentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentId}", url.PathEscape(deploymentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// deploymentPreflightHandleResponse handles the DeploymentPreflight response.
func (client *AzureSiteRecoveryManagementServiceAPIClient) deploymentPreflightHandleResponse(resp *http.Response) (AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse, error) {
	result := AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentPreflightModel); err != nil {
		return AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightResponse{}, err
	}
	return result, nil
}

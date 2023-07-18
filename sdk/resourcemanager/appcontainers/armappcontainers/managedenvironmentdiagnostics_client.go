//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// ManagedEnvironmentDiagnosticsClient contains the methods for the ManagedEnvironmentDiagnostics group.
// Don't use this type directly, use NewManagedEnvironmentDiagnosticsClient() instead.
type ManagedEnvironmentDiagnosticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedEnvironmentDiagnosticsClient creates a new instance of ManagedEnvironmentDiagnosticsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedEnvironmentDiagnosticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedEnvironmentDiagnosticsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedEnvironmentDiagnosticsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedEnvironmentDiagnosticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetDetector - Get the diagnostics data for a Managed Environment used to host container apps.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Environment.
//   - detectorName - Name of the Managed Environment detector.
//   - options - ManagedEnvironmentDiagnosticsClientGetDetectorOptions contains the optional parameters for the ManagedEnvironmentDiagnosticsClient.GetDetector
//     method.
func (client *ManagedEnvironmentDiagnosticsClient) GetDetector(ctx context.Context, resourceGroupName string, environmentName string, detectorName string, options *ManagedEnvironmentDiagnosticsClientGetDetectorOptions) (ManagedEnvironmentDiagnosticsClientGetDetectorResponse, error) {
	req, err := client.getDetectorCreateRequest(ctx, resourceGroupName, environmentName, detectorName, options)
	if err != nil {
		return ManagedEnvironmentDiagnosticsClientGetDetectorResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedEnvironmentDiagnosticsClientGetDetectorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentDiagnosticsClientGetDetectorResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDetectorHandleResponse(resp)
}

// getDetectorCreateRequest creates the GetDetector request.
func (client *ManagedEnvironmentDiagnosticsClient) getDetectorCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, detectorName string, options *ManagedEnvironmentDiagnosticsClientGetDetectorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/detectors/{detectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if detectorName == "" {
		return nil, errors.New("parameter detectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{detectorName}", url.PathEscape(detectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDetectorHandleResponse handles the GetDetector response.
func (client *ManagedEnvironmentDiagnosticsClient) getDetectorHandleResponse(resp *http.Response) (ManagedEnvironmentDiagnosticsClientGetDetectorResponse, error) {
	result := ManagedEnvironmentDiagnosticsClientGetDetectorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Diagnostics); err != nil {
		return ManagedEnvironmentDiagnosticsClientGetDetectorResponse{}, err
	}
	return result, nil
}

// ListDetectors - Get the list of diagnostics for a Managed Environment used to host container apps.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Environment.
//   - options - ManagedEnvironmentDiagnosticsClientListDetectorsOptions contains the optional parameters for the ManagedEnvironmentDiagnosticsClient.ListDetectors
//     method.
func (client *ManagedEnvironmentDiagnosticsClient) ListDetectors(ctx context.Context, resourceGroupName string, environmentName string, options *ManagedEnvironmentDiagnosticsClientListDetectorsOptions) (ManagedEnvironmentDiagnosticsClientListDetectorsResponse, error) {
	req, err := client.listDetectorsCreateRequest(ctx, resourceGroupName, environmentName, options)
	if err != nil {
		return ManagedEnvironmentDiagnosticsClientListDetectorsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedEnvironmentDiagnosticsClientListDetectorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentDiagnosticsClientListDetectorsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listDetectorsHandleResponse(resp)
}

// listDetectorsCreateRequest creates the ListDetectors request.
func (client *ManagedEnvironmentDiagnosticsClient) listDetectorsCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *ManagedEnvironmentDiagnosticsClientListDetectorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/detectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDetectorsHandleResponse handles the ListDetectors response.
func (client *ManagedEnvironmentDiagnosticsClient) listDetectorsHandleResponse(resp *http.Response) (ManagedEnvironmentDiagnosticsClientListDetectorsResponse, error) {
	result := ManagedEnvironmentDiagnosticsClientListDetectorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticsCollection); err != nil {
		return ManagedEnvironmentDiagnosticsClientListDetectorsResponse{}, err
	}
	return result, nil
}

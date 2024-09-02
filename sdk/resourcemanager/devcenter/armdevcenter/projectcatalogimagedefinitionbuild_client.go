//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

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

// ProjectCatalogImageDefinitionBuildClient contains the methods for the ProjectCatalogImageDefinitionBuild group.
// Don't use this type directly, use NewProjectCatalogImageDefinitionBuildClient() instead.
type ProjectCatalogImageDefinitionBuildClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProjectCatalogImageDefinitionBuildClient creates a new instance of ProjectCatalogImageDefinitionBuildClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProjectCatalogImageDefinitionBuildClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectCatalogImageDefinitionBuildClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProjectCatalogImageDefinitionBuildClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Cancels the specified build for an image definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - catalogName - The name of the Catalog.
//   - imageDefinitionName - The name of the Image Definition.
//   - buildName - The ID of the Image Definition Build.
//   - options - ProjectCatalogImageDefinitionBuildClientBeginCancelOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.BeginCancel
//     method.
func (client *ProjectCatalogImageDefinitionBuildClient) BeginCancel(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientBeginCancelOptions) (*runtime.Poller[ProjectCatalogImageDefinitionBuildClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, projectName, catalogName, imageDefinitionName, buildName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProjectCatalogImageDefinitionBuildClientCancelResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProjectCatalogImageDefinitionBuildClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cancel - Cancels the specified build for an image definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
func (client *ProjectCatalogImageDefinitionBuildClient) cancel(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "ProjectCatalogImageDefinitionBuildClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, projectName, catalogName, imageDefinitionName, buildName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ProjectCatalogImageDefinitionBuildClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/catalogs/{catalogName}/imageDefinitions/{imageDefinitionName}/builds/{buildName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if imageDefinitionName == "" {
		return nil, errors.New("parameter imageDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageDefinitionName}", url.PathEscape(imageDefinitionName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a build for a specified image definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - catalogName - The name of the Catalog.
//   - imageDefinitionName - The name of the Image Definition.
//   - buildName - The ID of the Image Definition Build.
//   - options - ProjectCatalogImageDefinitionBuildClientGetOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.Get
//     method.
func (client *ProjectCatalogImageDefinitionBuildClient) Get(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientGetOptions) (ProjectCatalogImageDefinitionBuildClientGetResponse, error) {
	var err error
	const operationName = "ProjectCatalogImageDefinitionBuildClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, catalogName, imageDefinitionName, buildName, options)
	if err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectCatalogImageDefinitionBuildClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProjectCatalogImageDefinitionBuildClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/catalogs/{catalogName}/imageDefinitions/{imageDefinitionName}/builds/{buildName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if imageDefinitionName == "" {
		return nil, errors.New("parameter imageDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageDefinitionName}", url.PathEscape(imageDefinitionName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectCatalogImageDefinitionBuildClient) getHandleResponse(resp *http.Response) (ProjectCatalogImageDefinitionBuildClientGetResponse, error) {
	result := ProjectCatalogImageDefinitionBuildClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageDefinitionBuild); err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetResponse{}, err
	}
	return result, nil
}

// GetBuildDetails - Gets Build details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - catalogName - The name of the Catalog.
//   - imageDefinitionName - The name of the Image Definition.
//   - buildName - The ID of the Image Definition Build.
//   - options - ProjectCatalogImageDefinitionBuildClientGetBuildDetailsOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.GetBuildDetails
//     method.
func (client *ProjectCatalogImageDefinitionBuildClient) GetBuildDetails(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientGetBuildDetailsOptions) (ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse, error) {
	var err error
	const operationName = "ProjectCatalogImageDefinitionBuildClient.GetBuildDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBuildDetailsCreateRequest(ctx, resourceGroupName, projectName, catalogName, imageDefinitionName, buildName, options)
	if err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{}, err
	}
	resp, err := client.getBuildDetailsHandleResponse(httpResp)
	return resp, err
}

// getBuildDetailsCreateRequest creates the GetBuildDetails request.
func (client *ProjectCatalogImageDefinitionBuildClient) getBuildDetailsCreateRequest(ctx context.Context, resourceGroupName string, projectName string, catalogName string, imageDefinitionName string, buildName string, options *ProjectCatalogImageDefinitionBuildClientGetBuildDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/catalogs/{catalogName}/imageDefinitions/{imageDefinitionName}/builds/{buildName}/getBuildDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if imageDefinitionName == "" {
		return nil, errors.New("parameter imageDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageDefinitionName}", url.PathEscape(imageDefinitionName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuildDetailsHandleResponse handles the GetBuildDetails response.
func (client *ProjectCatalogImageDefinitionBuildClient) getBuildDetailsHandleResponse(resp *http.Response) (ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse, error) {
	result := ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageDefinitionBuildDetails); err != nil {
		return ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{}, err
	}
	return result, nil
}

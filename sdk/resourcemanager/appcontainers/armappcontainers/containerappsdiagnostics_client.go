//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ContainerAppsDiagnosticsClient contains the methods for the ContainerAppsDiagnostics group.
// Don't use this type directly, use NewContainerAppsDiagnosticsClient() instead.
type ContainerAppsDiagnosticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsDiagnosticsClient creates a new instance of ContainerAppsDiagnosticsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsDiagnosticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsDiagnosticsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsDiagnosticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetDetector - Get a diagnostics result of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - detectorName - Name of the Container App Detector.
//   - options - ContainerAppsDiagnosticsClientGetDetectorOptions contains the optional parameters for the ContainerAppsDiagnosticsClient.GetDetector
//     method.
func (client *ContainerAppsDiagnosticsClient) GetDetector(ctx context.Context, resourceGroupName string, containerAppName string, detectorName string, options *ContainerAppsDiagnosticsClientGetDetectorOptions) (ContainerAppsDiagnosticsClientGetDetectorResponse, error) {
	var err error
	const operationName = "ContainerAppsDiagnosticsClient.GetDetector"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDetectorCreateRequest(ctx, resourceGroupName, containerAppName, detectorName, options)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetDetectorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetDetectorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsDiagnosticsClientGetDetectorResponse{}, err
	}
	resp, err := client.getDetectorHandleResponse(httpResp)
	return resp, err
}

// getDetectorCreateRequest creates the GetDetector request.
func (client *ContainerAppsDiagnosticsClient) getDetectorCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, detectorName string, options *ContainerAppsDiagnosticsClientGetDetectorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/detectors/{detectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if detectorName == "" {
		return nil, errors.New("parameter detectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{detectorName}", url.PathEscape(detectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDetectorHandleResponse handles the GetDetector response.
func (client *ContainerAppsDiagnosticsClient) getDetectorHandleResponse(resp *http.Response) (ContainerAppsDiagnosticsClientGetDetectorResponse, error) {
	result := ContainerAppsDiagnosticsClientGetDetectorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Diagnostics); err != nil {
		return ContainerAppsDiagnosticsClientGetDetectorResponse{}, err
	}
	return result, nil
}

// GetRevision - Get a revision of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - revisionName - Name of the Container App Revision.
//   - options - ContainerAppsDiagnosticsClientGetRevisionOptions contains the optional parameters for the ContainerAppsDiagnosticsClient.GetRevision
//     method.
func (client *ContainerAppsDiagnosticsClient) GetRevision(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *ContainerAppsDiagnosticsClientGetRevisionOptions) (ContainerAppsDiagnosticsClientGetRevisionResponse, error) {
	var err error
	const operationName = "ContainerAppsDiagnosticsClient.GetRevision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRevisionCreateRequest(ctx, resourceGroupName, containerAppName, revisionName, options)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetRevisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetRevisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsDiagnosticsClientGetRevisionResponse{}, err
	}
	resp, err := client.getRevisionHandleResponse(httpResp)
	return resp, err
}

// getRevisionCreateRequest creates the GetRevision request.
func (client *ContainerAppsDiagnosticsClient) getRevisionCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *ContainerAppsDiagnosticsClientGetRevisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/detectorProperties/revisionsApi/revisions/{revisionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if revisionName == "" {
		return nil, errors.New("parameter revisionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{revisionName}", url.PathEscape(revisionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRevisionHandleResponse handles the GetRevision response.
func (client *ContainerAppsDiagnosticsClient) getRevisionHandleResponse(resp *http.Response) (ContainerAppsDiagnosticsClientGetRevisionResponse, error) {
	result := ContainerAppsDiagnosticsClientGetRevisionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Revision); err != nil {
		return ContainerAppsDiagnosticsClientGetRevisionResponse{}, err
	}
	return result, nil
}

// GetRoot - Get the properties of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - options - ContainerAppsDiagnosticsClientGetRootOptions contains the optional parameters for the ContainerAppsDiagnosticsClient.GetRoot
//     method.
func (client *ContainerAppsDiagnosticsClient) GetRoot(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientGetRootOptions) (ContainerAppsDiagnosticsClientGetRootResponse, error) {
	var err error
	const operationName = "ContainerAppsDiagnosticsClient.GetRoot"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getRootCreateRequest(ctx, resourceGroupName, containerAppName, options)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetRootResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsDiagnosticsClientGetRootResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsDiagnosticsClientGetRootResponse{}, err
	}
	resp, err := client.getRootHandleResponse(httpResp)
	return resp, err
}

// getRootCreateRequest creates the GetRoot request.
func (client *ContainerAppsDiagnosticsClient) getRootCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientGetRootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/detectorProperties/rootApi/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRootHandleResponse handles the GetRoot response.
func (client *ContainerAppsDiagnosticsClient) getRootHandleResponse(resp *http.Response) (ContainerAppsDiagnosticsClientGetRootResponse, error) {
	result := ContainerAppsDiagnosticsClientGetRootResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerApp); err != nil {
		return ContainerAppsDiagnosticsClientGetRootResponse{}, err
	}
	return result, nil
}

// NewListDetectorsPager - Get the list of diagnostics for a given Container App.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App for which detector info is needed.
//   - options - ContainerAppsDiagnosticsClientListDetectorsOptions contains the optional parameters for the ContainerAppsDiagnosticsClient.NewListDetectorsPager
//     method.
func (client *ContainerAppsDiagnosticsClient) NewListDetectorsPager(resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientListDetectorsOptions) *runtime.Pager[ContainerAppsDiagnosticsClientListDetectorsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsDiagnosticsClientListDetectorsResponse]{
		More: func(page ContainerAppsDiagnosticsClientListDetectorsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsDiagnosticsClientListDetectorsResponse) (ContainerAppsDiagnosticsClientListDetectorsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsDiagnosticsClient.NewListDetectorsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listDetectorsCreateRequest(ctx, resourceGroupName, containerAppName, options)
			}, nil)
			if err != nil {
				return ContainerAppsDiagnosticsClientListDetectorsResponse{}, err
			}
			return client.listDetectorsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listDetectorsCreateRequest creates the ListDetectors request.
func (client *ContainerAppsDiagnosticsClient) listDetectorsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientListDetectorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/detectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDetectorsHandleResponse handles the ListDetectors response.
func (client *ContainerAppsDiagnosticsClient) listDetectorsHandleResponse(resp *http.Response) (ContainerAppsDiagnosticsClientListDetectorsResponse, error) {
	result := ContainerAppsDiagnosticsClientListDetectorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticsCollection); err != nil {
		return ContainerAppsDiagnosticsClientListDetectorsResponse{}, err
	}
	return result, nil
}

// NewListRevisionsPager - Get the Revisions for a given Container App.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App for which Revisions are needed.
//   - options - ContainerAppsDiagnosticsClientListRevisionsOptions contains the optional parameters for the ContainerAppsDiagnosticsClient.NewListRevisionsPager
//     method.
func (client *ContainerAppsDiagnosticsClient) NewListRevisionsPager(resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientListRevisionsOptions) *runtime.Pager[ContainerAppsDiagnosticsClientListRevisionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsDiagnosticsClientListRevisionsResponse]{
		More: func(page ContainerAppsDiagnosticsClientListRevisionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsDiagnosticsClientListRevisionsResponse) (ContainerAppsDiagnosticsClientListRevisionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsDiagnosticsClient.NewListRevisionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listRevisionsCreateRequest(ctx, resourceGroupName, containerAppName, options)
			}, nil)
			if err != nil {
				return ContainerAppsDiagnosticsClientListRevisionsResponse{}, err
			}
			return client.listRevisionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listRevisionsCreateRequest creates the ListRevisions request.
func (client *ContainerAppsDiagnosticsClient) listRevisionsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsDiagnosticsClientListRevisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/detectorProperties/revisionsApi/revisions/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRevisionsHandleResponse handles the ListRevisions response.
func (client *ContainerAppsDiagnosticsClient) listRevisionsHandleResponse(resp *http.Response) (ContainerAppsDiagnosticsClientListRevisionsResponse, error) {
	result := ContainerAppsDiagnosticsClientListRevisionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RevisionCollection); err != nil {
		return ContainerAppsDiagnosticsClientListRevisionsResponse{}, err
	}
	return result, nil
}

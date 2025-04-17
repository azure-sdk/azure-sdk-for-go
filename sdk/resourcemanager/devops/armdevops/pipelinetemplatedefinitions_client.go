// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevops

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PipelineTemplateDefinitionsClient contains the methods for the PipelineTemplateDefinitions group.
// Don't use this type directly, use NewPipelineTemplateDefinitionsClient() instead.
type PipelineTemplateDefinitionsClient struct {
	internal *arm.Client
}

// NewPipelineTemplateDefinitionsClient creates a new instance of PipelineTemplateDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelineTemplateDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineTemplateDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelineTemplateDefinitionsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists all pipeline templates which can be used to configure an Azure Pipeline.
//
// Generated from API version 2019-07-01-preview
//   - options - PipelineTemplateDefinitionsClientListOptions contains the optional parameters for the PipelineTemplateDefinitionsClient.NewListPager
//     method.
func (client *PipelineTemplateDefinitionsClient) NewListPager(options *PipelineTemplateDefinitionsClientListOptions) *runtime.Pager[PipelineTemplateDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineTemplateDefinitionsClientListResponse]{
		More: func(page PipelineTemplateDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineTemplateDefinitionsClientListResponse) (PipelineTemplateDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PipelineTemplateDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PipelineTemplateDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PipelineTemplateDefinitionsClient) listCreateRequest(ctx context.Context, _ *PipelineTemplateDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DevOps/pipelineTemplateDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineTemplateDefinitionsClient) listHandleResponse(resp *http.Response) (PipelineTemplateDefinitionsClientListResponse, error) {
	result := PipelineTemplateDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTemplateDefinitionListResult); err != nil {
		return PipelineTemplateDefinitionsClientListResponse{}, err
	}
	return result, nil
}

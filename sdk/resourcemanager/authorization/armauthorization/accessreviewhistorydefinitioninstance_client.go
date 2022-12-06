//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// AccessReviewHistoryDefinitionInstanceClient contains the methods for the AccessReviewHistoryDefinitionInstance group.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionInstanceClient() instead.
type AccessReviewHistoryDefinitionInstanceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessReviewHistoryDefinitionInstanceClient creates a new instance of AccessReviewHistoryDefinitionInstanceClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewHistoryDefinitionInstanceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewHistoryDefinitionInstanceClient, error) {
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
	client := &AccessReviewHistoryDefinitionInstanceClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GenerateDownloadURI - Generates a uri which can be used to retrieve review history data. This URI has a TTL of 1 day and
// can be retrieved by fetching the accessReviewHistoryDefinition object.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// historyDefinitionID - The id of the access review history definition.
// instanceID - The id of the access review history definition instance to generate a URI for.
// options - AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions contains the optional parameters for the
// AccessReviewHistoryDefinitionInstanceClient.GenerateDownloadURI method.
func (client *AccessReviewHistoryDefinitionInstanceClient) GenerateDownloadURI(ctx context.Context, historyDefinitionID string, instanceID string, options *AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions) (AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse, error) {
	req, err := client.generateDownloadURICreateRequest(ctx, historyDefinitionID, instanceID, options)
	if err != nil {
		return AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateDownloadURIHandleResponse(resp)
}

// generateDownloadURICreateRequest creates the GenerateDownloadURI request.
func (client *AccessReviewHistoryDefinitionInstanceClient) generateDownloadURICreateRequest(ctx context.Context, historyDefinitionID string, instanceID string, options *AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}/instances/{instanceId}/generateDownloadUri"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateDownloadURIHandleResponse handles the GenerateDownloadURI response.
func (client *AccessReviewHistoryDefinitionInstanceClient) generateDownloadURIHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse, error) {
	result := AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryInstance); err != nil {
		return AccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	return result, nil
}

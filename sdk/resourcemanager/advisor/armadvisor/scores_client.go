//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armadvisor

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

// ScoresClient contains the methods for the AdvisorScores group.
// Don't use this type directly, use NewScoresClient() instead.
type ScoresClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewScoresClient creates a new instance of ScoresClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScoresClient, error) {
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
	client := &ScoresClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the advisor score.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - name - The scope of Advisor score entity.
//   - options - ScoresClientGetOptions contains the optional parameters for the ScoresClient.Get method.
func (client *ScoresClient) Get(ctx context.Context, name string, options *ScoresClientGetOptions) (ScoresClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return ScoresClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScoresClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScoresClient) getCreateRequest(ctx context.Context, name string, options *ScoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/advisorScore/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScoresClient) getHandleResponse(resp *http.Response) (ScoresClientGetResponse, error) {
	result := ScoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScoreEntityForAdvisor); err != nil {
		return ScoresClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of advisor scores.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - options - ScoresClientListOptions contains the optional parameters for the ScoresClient.List method.
func (client *ScoresClient) List(ctx context.Context, options *ScoresClientListOptions) (ScoresClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return ScoresClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScoresClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScoresClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ScoresClient) listCreateRequest(ctx context.Context, options *ScoresClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/advisorScore"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScoresClient) listHandleResponse(resp *http.Response) (ScoresClientListResponse, error) {
	result := ScoresClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScoreResponse); err != nil {
		return ScoresClientListResponse{}, err
	}
	return result, nil
}

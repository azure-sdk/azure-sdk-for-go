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

// APICollectionOnboardingClient contains the methods for the APICollectionOnboarding group.
// Don't use this type directly, use NewAPICollectionOnboardingClient() instead.
type APICollectionOnboardingClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPICollectionOnboardingClient creates a new instance of APICollectionOnboardingClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPICollectionOnboardingClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APICollectionOnboardingClient, error) {
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
	client := &APICollectionOnboardingClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Onboard an Azure API Management API to Defender for APIs. The system will start monitoring the operations within
// the Azure Management API for intrusive behaviors and provide alerts for attacks that
// have been detected.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-20-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiCollectionID - A string representing the apiCollections resource within the Microsoft.Security provider namespace. This
//     string matches the Azure API Management API name.
//   - options - APICollectionOnboardingClientCreateOptions contains the optional parameters for the APICollectionOnboardingClient.Create
//     method.
func (client *APICollectionOnboardingClient) Create(ctx context.Context, resourceGroupName string, serviceName string, apiCollectionID string, options *APICollectionOnboardingClientCreateOptions) (APICollectionOnboardingClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, serviceName, apiCollectionID, options)
	if err != nil {
		return APICollectionOnboardingClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APICollectionOnboardingClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APICollectionOnboardingClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *APICollectionOnboardingClient) createCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiCollectionID string, options *APICollectionOnboardingClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/providers/Microsoft.Security/apiCollections/{apiCollectionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiCollectionID == "" {
		return nil, errors.New("parameter apiCollectionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiCollectionId}", url.PathEscape(apiCollectionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *APICollectionOnboardingClient) createHandleResponse(resp *http.Response) (APICollectionOnboardingClientCreateResponse, error) {
	result := APICollectionOnboardingClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollectionResponse); err != nil {
		return APICollectionOnboardingClientCreateResponse{}, err
	}
	return result, nil
}

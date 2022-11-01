//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

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

// FleetMembersClient contains the methods for the FleetMembers group.
// Don't use this type directly, use NewFleetMembersClient() instead.
type FleetMembersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFleetMembersClient creates a new instance of FleetMembersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFleetMembersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FleetMembersClient, error) {
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
	client := &FleetMembersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - A member contains a reference to an existing Kubernetes cluster. Creating a member makes the referenced
// cluster join the Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// fleetMemberName - The name of the Fleet member resource.
// parameters - The Fleet member to create or update.
// options - FleetMembersClientBeginCreateOrUpdateOptions contains the optional parameters for the FleetMembersClient.BeginCreateOrUpdate
// method.
func (client *FleetMembersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, parameters FleetMember, options *FleetMembersClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetMembersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, fleetMemberName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FleetMembersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FleetMembersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - A member contains a reference to an existing Kubernetes cluster. Creating a member makes the referenced
// cluster join the Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-02-preview
func (client *FleetMembersClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, parameters FleetMember, options *FleetMembersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FleetMembersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, parameters FleetMember, options *FleetMembersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deleting a Fleet member results in the member cluster leaving fleet. The Member azure resource is deleted
// upon success. The underlying cluster is not deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// fleetMemberName - The name of the Fleet member resource.
// options - FleetMembersClientBeginDeleteOptions contains the optional parameters for the FleetMembersClient.BeginDelete
// method.
func (client *FleetMembersClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*runtime.Poller[FleetMembersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, fleetMemberName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FleetMembersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FleetMembersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deleting a Fleet member results in the member cluster leaving fleet. The Member azure resource is deleted upon
// success. The underlying cluster is not deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-02-preview
func (client *FleetMembersClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FleetMembersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Fleet member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// fleetMemberName - The name of the Fleet member resource.
// options - FleetMembersClientGetOptions contains the optional parameters for the FleetMembersClient.Get method.
func (client *FleetMembersClient) Get(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientGetOptions) (FleetMembersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, options)
	if err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FleetMembersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FleetMembersClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FleetMembersClient) getHandleResponse(resp *http.Response) (FleetMembersClientGetResponse, error) {
	result := FleetMembersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetMember); err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFleetPager - Lists the members of a fleet.
// Generated from API version 2022-10-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// options - FleetMembersClientListByFleetOptions contains the optional parameters for the FleetMembersClient.ListByFleet
// method.
func (client *FleetMembersClient) NewListByFleetPager(resourceGroupName string, fleetName string, options *FleetMembersClientListByFleetOptions) *runtime.Pager[FleetMembersClientListByFleetResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetMembersClientListByFleetResponse]{
		More: func(page FleetMembersClientListByFleetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetMembersClientListByFleetResponse) (FleetMembersClientListByFleetResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFleetCreateRequest(ctx, resourceGroupName, fleetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FleetMembersClientListByFleetResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FleetMembersClientListByFleetResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FleetMembersClientListByFleetResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFleetHandleResponse(resp)
		},
	})
}

// listByFleetCreateRequest creates the ListByFleet request.
func (client *FleetMembersClient) listByFleetCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetMembersClientListByFleetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFleetHandleResponse handles the ListByFleet response.
func (client *FleetMembersClient) listByFleetHandleResponse(resp *http.Response) (FleetMembersClientListByFleetResponse, error) {
	result := FleetMembersClientListByFleetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetMembersListResult); err != nil {
		return FleetMembersClientListByFleetResponse{}, err
	}
	return result, nil
}

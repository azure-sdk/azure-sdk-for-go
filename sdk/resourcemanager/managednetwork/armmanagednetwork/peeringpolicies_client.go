// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PeeringPoliciesClient contains the methods for the ManagedNetworkPeeringPolicies group.
// Don't use this type directly, use NewPeeringPoliciesClient() instead.
type PeeringPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPeeringPoliciesClient creates a new instance of PeeringPoliciesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPeeringPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PeeringPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PeeringPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The Put ManagedNetworkPeeringPolicies operation creates/updates a new Managed Network Peering Policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkPeeringPolicyName - The name of the Managed Network Peering Policy.
//   - managedNetworkPolicy - Parameters supplied to create/update a Managed Network Peering Policy
//   - options - PeeringPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the PeeringPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *PeeringPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, managedNetworkPolicy PeeringPolicy, options *PeeringPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PeeringPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName, managedNetworkPolicy, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PeeringPoliciesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PeeringPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The Put ManagedNetworkPeeringPolicies operation creates/updates a new Managed Network Peering Policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *PeeringPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, managedNetworkPolicy PeeringPolicy, options *PeeringPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PeeringPoliciesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName, managedNetworkPolicy, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PeeringPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, managedNetworkPolicy PeeringPolicy, _ *PeeringPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkPeeringPolicyName == "" {
		return nil, errors.New("parameter managedNetworkPeeringPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkPeeringPolicyName}", url.PathEscape(managedNetworkPeeringPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, managedNetworkPolicy); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The Delete ManagedNetworkPeeringPolicies operation deletes a Managed Network Peering Policy, specified by
// the resource group, Managed Network name, and peering policy name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkPeeringPolicyName - The name of the Managed Network Peering Policy.
//   - options - PeeringPoliciesClientBeginDeleteOptions contains the optional parameters for the PeeringPoliciesClient.BeginDelete
//     method.
func (client *PeeringPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, options *PeeringPoliciesClientBeginDeleteOptions) (*runtime.Poller[PeeringPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PeeringPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PeeringPoliciesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The Delete ManagedNetworkPeeringPolicies operation deletes a Managed Network Peering Policy, specified by the
// resource group, Managed Network name, and peering policy name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *PeeringPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, options *PeeringPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PeeringPoliciesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PeeringPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, _ *PeeringPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkPeeringPolicyName == "" {
		return nil, errors.New("parameter managedNetworkPeeringPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkPeeringPolicyName}", url.PathEscape(managedNetworkPeeringPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The Get ManagedNetworkPeeringPolicies operation gets a Managed Network Peering Policy resource, specified by the
// resource group, Managed Network name, and peering policy name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkPeeringPolicyName - The name of the Managed Network Peering Policy.
//   - options - PeeringPoliciesClientGetOptions contains the optional parameters for the PeeringPoliciesClient.Get method.
func (client *PeeringPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, options *PeeringPoliciesClientGetOptions) (PeeringPoliciesClientGetResponse, error) {
	var err error
	const operationName = "PeeringPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName, options)
	if err != nil {
		return PeeringPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PeeringPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PeeringPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PeeringPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string, _ *PeeringPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkPeeringPolicyName == "" {
		return nil, errors.New("parameter managedNetworkPeeringPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkPeeringPolicyName}", url.PathEscape(managedNetworkPeeringPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeeringPoliciesClient) getHandleResponse(resp *http.Response) (PeeringPoliciesClientGetResponse, error) {
	result := PeeringPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeeringPolicy); err != nil {
		return PeeringPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedNetworkPager - The ListByManagedNetwork PeeringPolicies operation retrieves all the Managed Network Peering
// Policies in a specified Managed Network, in a paginated format.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - options - PeeringPoliciesClientListByManagedNetworkOptions contains the optional parameters for the PeeringPoliciesClient.NewListByManagedNetworkPager
//     method.
func (client *PeeringPoliciesClient) NewListByManagedNetworkPager(resourceGroupName string, managedNetworkName string, options *PeeringPoliciesClientListByManagedNetworkOptions) *runtime.Pager[PeeringPoliciesClientListByManagedNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[PeeringPoliciesClientListByManagedNetworkResponse]{
		More: func(page PeeringPoliciesClientListByManagedNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PeeringPoliciesClientListByManagedNetworkResponse) (PeeringPoliciesClientListByManagedNetworkResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PeeringPoliciesClient.NewListByManagedNetworkPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedNetworkCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
			}, nil)
			if err != nil {
				return PeeringPoliciesClientListByManagedNetworkResponse{}, err
			}
			return client.listByManagedNetworkHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedNetworkCreateRequest creates the ListByManagedNetwork request.
func (client *PeeringPoliciesClient) listByManagedNetworkCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *PeeringPoliciesClientListByManagedNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedNetworkHandleResponse handles the ListByManagedNetwork response.
func (client *PeeringPoliciesClient) listByManagedNetworkHandleResponse(resp *http.Response) (PeeringPoliciesClientListByManagedNetworkResponse, error) {
	result := PeeringPoliciesClientListByManagedNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeeringPolicyListResult); err != nil {
		return PeeringPoliciesClientListByManagedNetworkResponse{}, err
	}
	return result, nil
}

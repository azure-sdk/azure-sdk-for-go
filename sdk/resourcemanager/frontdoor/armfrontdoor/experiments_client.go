//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// ExperimentsClient contains the methods for the Experiments group.
// Don't use this type directly, use NewExperimentsClient() instead.
type ExperimentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExperimentsClient creates a new instance of ExperimentsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExperimentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExperimentsClient, error) {
	cl, err := arm.NewClient(moduleName+".ExperimentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExperimentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - experimentName - The Experiment identifier associated with the Experiment
//   - parameters - The Experiment resource
//   - options - ExperimentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.BeginCreateOrUpdate
//     method.
func (client *ExperimentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExperimentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, profileName, experimentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ExperimentsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ExperimentsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *ExperimentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, profileName, experimentName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExperimentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment, options *ExperimentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - experimentName - The Experiment identifier associated with the Experiment
//   - options - ExperimentsClientBeginDeleteOptions contains the optional parameters for the ExperimentsClient.BeginDelete method.
func (client *ExperimentsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*runtime.Poller[ExperimentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, experimentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ExperimentsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ExperimentsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *ExperimentsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, experimentName, options)
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
func (client *ExperimentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Experiment by ExperimentName
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - experimentName - The Experiment identifier associated with the Experiment
//   - options - ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
func (client *ExperimentsClient) Get(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientGetOptions) (ExperimentsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, experimentName, options)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExperimentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExperimentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *ExperimentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExperimentsClient) getHandleResponse(resp *http.Response) (ExperimentsClientGetResponse, error) {
	result := ExperimentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProfilePager - Gets a list of Experiments
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - options - ExperimentsClientListByProfileOptions contains the optional parameters for the ExperimentsClient.NewListByProfilePager
//     method.
func (client *ExperimentsClient) NewListByProfilePager(resourceGroupName string, profileName string, options *ExperimentsClientListByProfileOptions) *runtime.Pager[ExperimentsClientListByProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListByProfileResponse]{
		More: func(page ExperimentsClientListByProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListByProfileResponse) (ExperimentsClientListByProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExperimentsClientListByProfileResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExperimentsClientListByProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExperimentsClientListByProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProfileHandleResponse(resp)
		},
	})
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *ExperimentsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ExperimentsClientListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *ExperimentsClient) listByProfileHandleResponse(resp *http.Response) (ExperimentsClientListByProfileResponse, error) {
	result := ExperimentsClientListByProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentList); err != nil {
		return ExperimentsClientListByProfileResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - experimentName - The Experiment identifier associated with the Experiment
//   - parameters - The Experiment Update Model
//   - options - ExperimentsClientBeginUpdateOptions contains the optional parameters for the ExperimentsClient.BeginUpdate method.
func (client *ExperimentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (*runtime.Poller[ExperimentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, experimentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ExperimentsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ExperimentsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an Experiment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *ExperimentsClient) update(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, experimentName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ExperimentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel, options *ExperimentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos

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

// ExperimentsClient contains the methods for the Experiments group.
// Don't use this type directly, use NewExperimentsClient() instead.
type ExperimentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExperimentsClient creates a new instance of ExperimentsClient with the specified values.
//   - subscriptionID - GUID that represents an Azure subscription ID.
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

// Cancel - Cancel a running Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientCancelOptions contains the optional parameters for the ExperimentsClient.Cancel method.
func (client *ExperimentsClient) Cancel(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientCancelOptions) (ExperimentsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsClientCancelResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ExperimentsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *ExperimentsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *ExperimentsClient) cancelHandleResponse(resp *http.Response) (ExperimentsClientCancelResponse, error) {
	result := ExperimentsClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentCancelOperationResult); err != nil {
		return ExperimentsClientCancelResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Create or update a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - experiment - Experiment resource to be created or updated.
//   - options - ExperimentsClientCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.CreateOrUpdate
//     method.
func (client *ExperimentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, experimentName string, experiment Experiment, options *ExperimentsClientCreateOrUpdateOptions) (ExperimentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, experimentName, experiment, options)
	if err != nil {
		return ExperimentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExperimentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, experiment Experiment, options *ExperimentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, experiment)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExperimentsClient) createOrUpdateHandleResponse(resp *http.Response) (ExperimentsClientCreateOrUpdateResponse, error) {
	result := ExperimentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientDeleteOptions contains the optional parameters for the ExperimentsClient.Delete method.
func (client *ExperimentsClient) Delete(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientDeleteOptions) (ExperimentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ExperimentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ExperimentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExperimentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
func (client *ExperimentsClient) Get(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientGetOptions) (ExperimentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExperimentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
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

// GetExecutionDetails - Get an execution detail of a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - executionDetailsID - GUID that represents a Experiment execution detail.
//   - options - ExperimentsClientGetExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.GetExecutionDetails
//     method.
func (client *ExperimentsClient) GetExecutionDetails(ctx context.Context, resourceGroupName string, experimentName string, executionDetailsID string, options *ExperimentsClientGetExecutionDetailsOptions) (ExperimentsClientGetExecutionDetailsResponse, error) {
	req, err := client.getExecutionDetailsCreateRequest(ctx, resourceGroupName, experimentName, executionDetailsID, options)
	if err != nil {
		return ExperimentsClientGetExecutionDetailsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetExecutionDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientGetExecutionDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getExecutionDetailsHandleResponse(resp)
}

// getExecutionDetailsCreateRequest creates the GetExecutionDetails request.
func (client *ExperimentsClient) getExecutionDetailsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, executionDetailsID string, options *ExperimentsClientGetExecutionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executionDetails/{executionDetailsId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if executionDetailsID == "" {
		return nil, errors.New("parameter executionDetailsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{executionDetailsId}", url.PathEscape(executionDetailsID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getExecutionDetailsHandleResponse handles the GetExecutionDetails response.
func (client *ExperimentsClient) getExecutionDetailsHandleResponse(resp *http.Response) (ExperimentsClientGetExecutionDetailsResponse, error) {
	result := ExperimentsClientGetExecutionDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionDetails); err != nil {
		return ExperimentsClientGetExecutionDetailsResponse{}, err
	}
	return result, nil
}

// GetStatus - Get a status of a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - statusID - GUID that represents a Experiment status.
//   - options - ExperimentsClientGetStatusOptions contains the optional parameters for the ExperimentsClient.GetStatus method.
func (client *ExperimentsClient) GetStatus(ctx context.Context, resourceGroupName string, experimentName string, statusID string, options *ExperimentsClientGetStatusOptions) (ExperimentsClientGetStatusResponse, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, experimentName, statusID, options)
	if err != nil {
		return ExperimentsClientGetStatusResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientGetStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStatusHandleResponse(resp)
}

// getStatusCreateRequest creates the GetStatus request.
func (client *ExperimentsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, statusID string, options *ExperimentsClientGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/statuses/{statusId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	if statusID == "" {
		return nil, errors.New("parameter statusID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{statusId}", url.PathEscape(statusID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *ExperimentsClient) getStatusHandleResponse(resp *http.Response) (ExperimentsClientGetStatusResponse, error) {
	result := ExperimentsClientGetStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStatus); err != nil {
		return ExperimentsClientGetStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Experiment resources in a resource group.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - options - ExperimentsClientListOptions contains the optional parameters for the ExperimentsClient.NewListPager method.
func (client *ExperimentsClient) NewListPager(resourceGroupName string, options *ExperimentsClientListOptions) *runtime.Pager[ExperimentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListResponse]{
		More: func(page ExperimentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListResponse) (ExperimentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExperimentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExperimentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExperimentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExperimentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ExperimentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExperimentsClient) listHandleResponse(resp *http.Response) (ExperimentsClientListResponse, error) {
	result := ExperimentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Get a list of Experiment resources in a subscription.
//
// Generated from API version 2023-04-15-preview
//   - options - ExperimentsClientListAllOptions contains the optional parameters for the ExperimentsClient.NewListAllPager method.
func (client *ExperimentsClient) NewListAllPager(options *ExperimentsClientListAllOptions) *runtime.Pager[ExperimentsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListAllResponse]{
		More: func(page ExperimentsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListAllResponse) (ExperimentsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExperimentsClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExperimentsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExperimentsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *ExperimentsClient) listAllCreateRequest(ctx context.Context, options *ExperimentsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/experiments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	if options != nil && options.Running != nil {
		reqQP.Set("running", strconv.FormatBool(*options.Running))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ExperimentsClient) listAllHandleResponse(resp *http.Response) (ExperimentsClientListAllResponse, error) {
	result := ExperimentsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentListResult); err != nil {
		return ExperimentsClientListAllResponse{}, err
	}
	return result, nil
}

// NewListAllStatusesPager - Get a list of statuses of a Experiment resource.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientListAllStatusesOptions contains the optional parameters for the ExperimentsClient.NewListAllStatusesPager
//     method.
func (client *ExperimentsClient) NewListAllStatusesPager(resourceGroupName string, experimentName string, options *ExperimentsClientListAllStatusesOptions) *runtime.Pager[ExperimentsClientListAllStatusesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListAllStatusesResponse]{
		More: func(page ExperimentsClientListAllStatusesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListAllStatusesResponse) (ExperimentsClientListAllStatusesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllStatusesCreateRequest(ctx, resourceGroupName, experimentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExperimentsClientListAllStatusesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExperimentsClientListAllStatusesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExperimentsClientListAllStatusesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllStatusesHandleResponse(resp)
		},
	})
}

// listAllStatusesCreateRequest creates the ListAllStatuses request.
func (client *ExperimentsClient) listAllStatusesCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientListAllStatusesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/statuses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllStatusesHandleResponse handles the ListAllStatuses response.
func (client *ExperimentsClient) listAllStatusesHandleResponse(resp *http.Response) (ExperimentsClientListAllStatusesResponse, error) {
	result := ExperimentsClientListAllStatusesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStatusListResult); err != nil {
		return ExperimentsClientListAllStatusesResponse{}, err
	}
	return result, nil
}

// NewListExecutionDetailsPager - Get a list of execution details of a Experiment resource.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientListExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.NewListExecutionDetailsPager
//     method.
func (client *ExperimentsClient) NewListExecutionDetailsPager(resourceGroupName string, experimentName string, options *ExperimentsClientListExecutionDetailsOptions) *runtime.Pager[ExperimentsClientListExecutionDetailsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExperimentsClientListExecutionDetailsResponse]{
		More: func(page ExperimentsClientListExecutionDetailsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExperimentsClientListExecutionDetailsResponse) (ExperimentsClientListExecutionDetailsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listExecutionDetailsCreateRequest(ctx, resourceGroupName, experimentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExperimentsClientListExecutionDetailsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExperimentsClientListExecutionDetailsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExperimentsClientListExecutionDetailsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listExecutionDetailsHandleResponse(resp)
		},
	})
}

// listExecutionDetailsCreateRequest creates the ListExecutionDetails request.
func (client *ExperimentsClient) listExecutionDetailsCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientListExecutionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/executionDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listExecutionDetailsHandleResponse handles the ListExecutionDetails response.
func (client *ExperimentsClient) listExecutionDetailsHandleResponse(resp *http.Response) (ExperimentsClientListExecutionDetailsResponse, error) {
	result := ExperimentsClientListExecutionDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentExecutionDetailsListResult); err != nil {
		return ExperimentsClientListExecutionDetailsResponse{}, err
	}
	return result, nil
}

// Start - Start a Experiment resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - options - ExperimentsClientStartOptions contains the optional parameters for the ExperimentsClient.Start method.
func (client *ExperimentsClient) Start(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientStartOptions) (ExperimentsClientStartResponse, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, experimentName, options)
	if err != nil {
		return ExperimentsClientStartResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientStartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ExperimentsClientStartResponse{}, runtime.NewResponseError(resp)
	}
	return client.startHandleResponse(resp)
}

// startCreateRequest creates the Start request.
func (client *ExperimentsClient) startCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, options *ExperimentsClientStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// startHandleResponse handles the Start response.
func (client *ExperimentsClient) startHandleResponse(resp *http.Response) (ExperimentsClientStartResponse, error) {
	result := ExperimentsClientStartResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExperimentStartOperationResult); err != nil {
		return ExperimentsClientStartResponse{}, err
	}
	return result, nil
}

// Update - The operation to update an experiment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - experimentName - String that represents a Experiment resource name.
//   - experiment - Parameters supplied to the Update experiment operation.
//   - options - ExperimentsClientUpdateOptions contains the optional parameters for the ExperimentsClient.Update method.
func (client *ExperimentsClient) Update(ctx context.Context, resourceGroupName string, experimentName string, experiment ExperimentUpdate, options *ExperimentsClientUpdateOptions) (ExperimentsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, experimentName, experiment, options)
	if err != nil {
		return ExperimentsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExperimentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExperimentsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ExperimentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, experimentName string, experiment ExperimentUpdate, options *ExperimentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments/{experimentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if experimentName == "" {
		return nil, errors.New("parameter experimentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{experimentName}", url.PathEscape(experimentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, experiment)
}

// updateHandleResponse handles the Update response.
func (client *ExperimentsClient) updateHandleResponse(resp *http.Response) (ExperimentsClientUpdateResponse, error) {
	result := ExperimentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Experiment); err != nil {
		return ExperimentsClientUpdateResponse{}, err
	}
	return result, nil
}

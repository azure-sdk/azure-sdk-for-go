// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

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

// OracleSubscriptionsClient contains the methods for the OracleSubscriptions group.
// Don't use this type directly, use NewOracleSubscriptionsClient() instead.
type OracleSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOracleSubscriptionsClient creates a new instance of OracleSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOracleSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OracleSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OracleSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAddAzureSubscriptions - Add Azure Subscriptions
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - body - The content of the action request
//   - options - OracleSubscriptionsClientBeginAddAzureSubscriptionsOptions contains the optional parameters for the OracleSubscriptionsClient.BeginAddAzureSubscriptions
//     method.
func (client *OracleSubscriptionsClient) BeginAddAzureSubscriptions(ctx context.Context, body AzureSubscriptions, options *OracleSubscriptionsClientBeginAddAzureSubscriptionsOptions) (*runtime.Poller[OracleSubscriptionsClientAddAzureSubscriptionsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.addAzureSubscriptions(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientAddAzureSubscriptionsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientAddAzureSubscriptionsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// AddAzureSubscriptions - Add Azure Subscriptions
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) addAzureSubscriptions(ctx context.Context, body AzureSubscriptions, options *OracleSubscriptionsClientBeginAddAzureSubscriptionsOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginAddAzureSubscriptions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addAzureSubscriptionsCreateRequest(ctx, body, options)
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

// addAzureSubscriptionsCreateRequest creates the AddAzureSubscriptions request.
func (client *OracleSubscriptionsClient) addAzureSubscriptionsCreateRequest(ctx context.Context, body AzureSubscriptions, _ *OracleSubscriptionsClientBeginAddAzureSubscriptionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default/addAzureSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Create a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resource - Resource create parameters.
//   - options - OracleSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the OracleSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *OracleSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resource OracleSubscription, options *OracleSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[OracleSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) createOrUpdate(ctx context.Context, resource OracleSubscription, options *OracleSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resource, options)
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
func (client *OracleSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resource OracleSubscription, _ *OracleSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientBeginDeleteOptions contains the optional parameters for the OracleSubscriptionsClient.BeginDelete
//     method.
func (client *OracleSubscriptionsClient) BeginDelete(ctx context.Context, options *OracleSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[OracleSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) deleteOperation(ctx context.Context, options *OracleSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OracleSubscriptionsClient) deleteCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientGetOptions contains the optional parameters for the OracleSubscriptionsClient.Get method.
func (client *OracleSubscriptionsClient) Get(ctx context.Context, options *OracleSubscriptionsClientGetOptions) (OracleSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return OracleSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OracleSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OracleSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OracleSubscriptionsClient) getCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OracleSubscriptionsClient) getHandleResponse(resp *http.Response) (OracleSubscriptionsClientGetResponse, error) {
	result := OracleSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OracleSubscription); err != nil {
		return OracleSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// BeginListActivationLinks - List Activation Links
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientBeginListActivationLinksOptions contains the optional parameters for the OracleSubscriptionsClient.BeginListActivationLinks
//     method.
func (client *OracleSubscriptionsClient) BeginListActivationLinks(ctx context.Context, options *OracleSubscriptionsClientBeginListActivationLinksOptions) (*runtime.Poller[OracleSubscriptionsClientListActivationLinksResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listActivationLinks(ctx, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientListActivationLinksResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientListActivationLinksResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListActivationLinks - List Activation Links
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) listActivationLinks(ctx context.Context, options *OracleSubscriptionsClientBeginListActivationLinksOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginListActivationLinks"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listActivationLinksCreateRequest(ctx, options)
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

// listActivationLinksCreateRequest creates the ListActivationLinks request.
func (client *OracleSubscriptionsClient) listActivationLinksCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientBeginListActivationLinksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default/listActivationLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListBySubscriptionPager - List OracleSubscription resources by subscription ID
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientListBySubscriptionOptions contains the optional parameters for the OracleSubscriptionsClient.NewListBySubscriptionPager
//     method.
func (client *OracleSubscriptionsClient) NewListBySubscriptionPager(options *OracleSubscriptionsClientListBySubscriptionOptions) *runtime.Pager[OracleSubscriptionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OracleSubscriptionsClientListBySubscriptionResponse]{
		More: func(page OracleSubscriptionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OracleSubscriptionsClientListBySubscriptionResponse) (OracleSubscriptionsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OracleSubscriptionsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OracleSubscriptionsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OracleSubscriptionsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OracleSubscriptionsClient) listBySubscriptionHandleResponse(resp *http.Response) (OracleSubscriptionsClientListBySubscriptionResponse, error) {
	result := OracleSubscriptionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OracleSubscriptionListResult); err != nil {
		return OracleSubscriptionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginListCloudAccountDetails - List Cloud Account Details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientBeginListCloudAccountDetailsOptions contains the optional parameters for the OracleSubscriptionsClient.BeginListCloudAccountDetails
//     method.
func (client *OracleSubscriptionsClient) BeginListCloudAccountDetails(ctx context.Context, options *OracleSubscriptionsClientBeginListCloudAccountDetailsOptions) (*runtime.Poller[OracleSubscriptionsClientListCloudAccountDetailsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listCloudAccountDetails(ctx, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientListCloudAccountDetailsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientListCloudAccountDetailsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListCloudAccountDetails - List Cloud Account Details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) listCloudAccountDetails(ctx context.Context, options *OracleSubscriptionsClientBeginListCloudAccountDetailsOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginListCloudAccountDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCloudAccountDetailsCreateRequest(ctx, options)
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

// listCloudAccountDetailsCreateRequest creates the ListCloudAccountDetails request.
func (client *OracleSubscriptionsClient) listCloudAccountDetailsCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientBeginListCloudAccountDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default/listCloudAccountDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginListSaasSubscriptionDetails - List Saas Subscription Details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - options - OracleSubscriptionsClientBeginListSaasSubscriptionDetailsOptions contains the optional parameters for the OracleSubscriptionsClient.BeginListSaasSubscriptionDetails
//     method.
func (client *OracleSubscriptionsClient) BeginListSaasSubscriptionDetails(ctx context.Context, options *OracleSubscriptionsClientBeginListSaasSubscriptionDetailsOptions) (*runtime.Poller[OracleSubscriptionsClientListSaasSubscriptionDetailsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listSaasSubscriptionDetails(ctx, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientListSaasSubscriptionDetailsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientListSaasSubscriptionDetailsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ListSaasSubscriptionDetails - List Saas Subscription Details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) listSaasSubscriptionDetails(ctx context.Context, options *OracleSubscriptionsClientBeginListSaasSubscriptionDetailsOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginListSaasSubscriptionDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSaasSubscriptionDetailsCreateRequest(ctx, options)
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

// listSaasSubscriptionDetailsCreateRequest creates the ListSaasSubscriptionDetails request.
func (client *OracleSubscriptionsClient) listSaasSubscriptionDetailsCreateRequest(ctx context.Context, _ *OracleSubscriptionsClientBeginListSaasSubscriptionDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default/listSaasSubscriptionDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - properties - The resource properties to be updated.
//   - options - OracleSubscriptionsClientBeginUpdateOptions contains the optional parameters for the OracleSubscriptionsClient.BeginUpdate
//     method.
func (client *OracleSubscriptionsClient) BeginUpdate(ctx context.Context, properties OracleSubscriptionUpdate, options *OracleSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[OracleSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OracleSubscriptionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OracleSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a OracleSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *OracleSubscriptionsClient) update(ctx context.Context, properties OracleSubscriptionUpdate, options *OracleSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OracleSubscriptionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, properties, options)
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
func (client *OracleSubscriptionsClient) updateCreateRequest(ctx context.Context, properties OracleSubscriptionUpdate, _ *OracleSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/oracleSubscriptions/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

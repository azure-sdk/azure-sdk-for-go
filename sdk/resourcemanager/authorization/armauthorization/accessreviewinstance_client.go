// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstanceClient contains the methods for the AccessReviewInstance group.
// Don't use this type directly, use NewAccessReviewInstanceClient() instead.
type AccessReviewInstanceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewInstanceClient creates a new instance of AccessReviewInstanceClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewInstanceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewInstanceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewInstanceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AcceptRecommendations - An action to accept recommendations for decision in an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceClientAcceptRecommendationsOptions contains the optional parameters for the AccessReviewInstanceClient.AcceptRecommendations
//     method.
func (client *AccessReviewInstanceClient) AcceptRecommendations(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientAcceptRecommendationsOptions) (AccessReviewInstanceClientAcceptRecommendationsResponse, error) {
	var err error
	const operationName = "AccessReviewInstanceClient.AcceptRecommendations"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.acceptRecommendationsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstanceClientAcceptRecommendationsResponse{}, err
	}
	return AccessReviewInstanceClientAcceptRecommendationsResponse{}, nil
}

// acceptRecommendationsCreateRequest creates the AcceptRecommendations request.
func (client *AccessReviewInstanceClient) acceptRecommendationsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, _ *AccessReviewInstanceClientAcceptRecommendationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/acceptRecommendations"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ApplyDecisions - An action to apply all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceClientApplyDecisionsOptions contains the optional parameters for the AccessReviewInstanceClient.ApplyDecisions
//     method.
func (client *AccessReviewInstanceClient) ApplyDecisions(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientApplyDecisionsOptions) (AccessReviewInstanceClientApplyDecisionsResponse, error) {
	var err error
	const operationName = "AccessReviewInstanceClient.ApplyDecisions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.applyDecisionsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	return AccessReviewInstanceClientApplyDecisionsResponse{}, nil
}

// applyDecisionsCreateRequest creates the ApplyDecisions request.
func (client *AccessReviewInstanceClient) applyDecisionsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, _ *AccessReviewInstanceClientApplyDecisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/applyDecisions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ResetDecisions - An action to reset all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceClientResetDecisionsOptions contains the optional parameters for the AccessReviewInstanceClient.ResetDecisions
//     method.
func (client *AccessReviewInstanceClient) ResetDecisions(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientResetDecisionsOptions) (AccessReviewInstanceClientResetDecisionsResponse, error) {
	var err error
	const operationName = "AccessReviewInstanceClient.ResetDecisions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetDecisionsCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	return AccessReviewInstanceClientResetDecisionsResponse{}, nil
}

// resetDecisionsCreateRequest creates the ResetDecisions request.
func (client *AccessReviewInstanceClient) resetDecisionsCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, _ *AccessReviewInstanceClientResetDecisionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/resetDecisions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// SendReminders - An action to send reminders for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceClientSendRemindersOptions contains the optional parameters for the AccessReviewInstanceClient.SendReminders
//     method.
func (client *AccessReviewInstanceClient) SendReminders(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientSendRemindersOptions) (AccessReviewInstanceClientSendRemindersResponse, error) {
	var err error
	const operationName = "AccessReviewInstanceClient.SendReminders"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendRemindersCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientSendRemindersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstanceClientSendRemindersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstanceClientSendRemindersResponse{}, err
	}
	return AccessReviewInstanceClientSendRemindersResponse{}, nil
}

// sendRemindersCreateRequest creates the SendReminders request.
func (client *AccessReviewInstanceClient) sendRemindersCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, _ *AccessReviewInstanceClientSendRemindersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/sendReminders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - An action to stop an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceClientStopOptions contains the optional parameters for the AccessReviewInstanceClient.Stop
//     method.
func (client *AccessReviewInstanceClient) Stop(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceClientStopOptions) (AccessReviewInstanceClientStopResponse, error) {
	var err error
	const operationName = "AccessReviewInstanceClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstanceClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewInstanceClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewInstanceClientStopResponse{}, err
	}
	return AccessReviewInstanceClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *AccessReviewInstanceClient) stopCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, _ *AccessReviewInstanceClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

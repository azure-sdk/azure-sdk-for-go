// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

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

// TicketsNoSubscriptionClient contains the methods for the SupportTicketsNoSubscription group.
// Don't use this type directly, use NewTicketsNoSubscriptionClient() instead.
type TicketsNoSubscriptionClient struct {
	internal *arm.Client
}

// NewTicketsNoSubscriptionClient creates a new instance of TicketsNoSubscriptionClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTicketsNoSubscriptionClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TicketsNoSubscriptionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TicketsNoSubscriptionClient{
		internal: cl,
	}
	return client, nil
}

// CheckNameAvailability - Check the availability of a resource name. This API should be used to check the uniqueness of the
// name for support ticket creation for the selected subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - checkNameAvailabilityInput - Input to check.
//   - options - TicketsNoSubscriptionClientCheckNameAvailabilityOptions contains the optional parameters for the TicketsNoSubscriptionClient.CheckNameAvailability
//     method.
func (client *TicketsNoSubscriptionClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *TicketsNoSubscriptionClientCheckNameAvailabilityOptions) (TicketsNoSubscriptionClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "TicketsNoSubscriptionClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityInput, options)
	if err != nil {
		return TicketsNoSubscriptionClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsNoSubscriptionClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsNoSubscriptionClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *TicketsNoSubscriptionClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, _ *TicketsNoSubscriptionClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *TicketsNoSubscriptionClient) checkNameAvailabilityHandleResponse(resp *http.Response) (TicketsNoSubscriptionClientCheckNameAvailabilityResponse, error) {
	result := TicketsNoSubscriptionClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return TicketsNoSubscriptionClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates a new support ticket for Billing, and Subscription Management issues. Learn the prerequisites [https://aka.ms/supportAPI]
// required to create a support ticket.
// Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required
// for support ticket creation.
// Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the Manage
// support ticket [https://portal.azure.com/#blade/MicrosoftAzure
// Support/HelpAndSupportBlade/managesupportrequest] page in the Azure portal, select the support ticket, and use the file
// upload control to add a new file.
// Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure
// support engineer working on your ticket will reach out to you for consent if your
// issue requires gathering diagnostic information from your Azure resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - supportTicketName - Support ticket name.
//   - createSupportTicketParameters - Support ticket request payload.
//   - options - TicketsNoSubscriptionClientBeginCreateOptions contains the optional parameters for the TicketsNoSubscriptionClient.BeginCreate
//     method.
func (client *TicketsNoSubscriptionClient) BeginCreate(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, options *TicketsNoSubscriptionClientBeginCreateOptions) (*runtime.Poller[TicketsNoSubscriptionClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, supportTicketName, createSupportTicketParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TicketsNoSubscriptionClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TicketsNoSubscriptionClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a new support ticket for Billing, and Subscription Management issues. Learn the prerequisites [https://aka.ms/supportAPI]
// required to create a support ticket.
// Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required
// for support ticket creation.
// Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the Manage
// support ticket [https://portal.azure.com/#blade/MicrosoftAzure
// Support/HelpAndSupportBlade/managesupportrequest] page in the Azure portal, select the support ticket, and use the file
// upload control to add a new file.
// Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure
// support engineer working on your ticket will reach out to you for consent if your
// issue requires gathering diagnostic information from your Azure resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *TicketsNoSubscriptionClient) create(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, options *TicketsNoSubscriptionClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "TicketsNoSubscriptionClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, supportTicketName, createSupportTicketParameters, options)
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

// createCreateRequest creates the Create request.
func (client *TicketsNoSubscriptionClient) createCreateRequest(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, _ *TicketsNoSubscriptionClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createSupportTicketParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets details for a specific support ticket. Support ticket data is available for 18 months after ticket creation.
// If a ticket was created more than 18 months ago, a request for data might cause an
// error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - supportTicketName - Support ticket name.
//   - options - TicketsNoSubscriptionClientGetOptions contains the optional parameters for the TicketsNoSubscriptionClient.Get
//     method.
func (client *TicketsNoSubscriptionClient) Get(ctx context.Context, supportTicketName string, options *TicketsNoSubscriptionClientGetOptions) (TicketsNoSubscriptionClientGetResponse, error) {
	var err error
	const operationName = "TicketsNoSubscriptionClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, supportTicketName, options)
	if err != nil {
		return TicketsNoSubscriptionClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsNoSubscriptionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsNoSubscriptionClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TicketsNoSubscriptionClient) getCreateRequest(ctx context.Context, supportTicketName string, _ *TicketsNoSubscriptionClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TicketsNoSubscriptionClient) getHandleResponse(resp *http.Response) (TicketsNoSubscriptionClientGetResponse, error) {
	result := TicketsNoSubscriptionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketDetails); err != nil {
		return TicketsNoSubscriptionClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the support tickets.
// You can also filter the support tickets by Status, CreatedDate, , ServiceId, and ProblemClassificationId using the $filter
// parameter. Output will be a paged result with nextLink, using which you can
// retrieve the next set of support tickets.
// Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago,
// a request for data might cause an error.
//
// Generated from API version 2023-06-01-preview
//   - options - TicketsNoSubscriptionClientListOptions contains the optional parameters for the TicketsNoSubscriptionClient.NewListPager
//     method.
func (client *TicketsNoSubscriptionClient) NewListPager(options *TicketsNoSubscriptionClientListOptions) *runtime.Pager[TicketsNoSubscriptionClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TicketsNoSubscriptionClientListResponse]{
		More: func(page TicketsNoSubscriptionClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TicketsNoSubscriptionClientListResponse) (TicketsNoSubscriptionClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TicketsNoSubscriptionClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TicketsNoSubscriptionClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TicketsNoSubscriptionClient) listCreateRequest(ctx context.Context, options *TicketsNoSubscriptionClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/supportTickets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TicketsNoSubscriptionClient) listHandleResponse(resp *http.Response) (TicketsNoSubscriptionClientListResponse, error) {
	result := TicketsNoSubscriptionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketsListResult); err != nil {
		return TicketsNoSubscriptionClientListResponse{}, err
	}
	return result, nil
}

// Update - This API allows you to update the severity level, ticket status, and your contact information in the support ticket.
// Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer.
// In such a case, contact your support engineer to request severity update by
// adding a new communication using the Communications API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - supportTicketName - Support ticket name.
//   - updateSupportTicket - UpdateSupportTicket object.
//   - options - TicketsNoSubscriptionClientUpdateOptions contains the optional parameters for the TicketsNoSubscriptionClient.Update
//     method.
func (client *TicketsNoSubscriptionClient) Update(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket, options *TicketsNoSubscriptionClientUpdateOptions) (TicketsNoSubscriptionClientUpdateResponse, error) {
	var err error
	const operationName = "TicketsNoSubscriptionClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, supportTicketName, updateSupportTicket, options)
	if err != nil {
		return TicketsNoSubscriptionClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsNoSubscriptionClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsNoSubscriptionClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TicketsNoSubscriptionClient) updateCreateRequest(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket, _ *TicketsNoSubscriptionClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateSupportTicket); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TicketsNoSubscriptionClient) updateHandleResponse(resp *http.Response) (TicketsNoSubscriptionClientUpdateResponse, error) {
	result := TicketsNoSubscriptionClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketDetails); err != nil {
		return TicketsNoSubscriptionClientUpdateResponse{}, err
	}
	return result, nil
}

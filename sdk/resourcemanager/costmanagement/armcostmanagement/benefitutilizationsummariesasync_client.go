//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement

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

// BenefitUtilizationSummariesAsyncClient contains the methods for the BenefitUtilizationSummariesAsync group.
// Don't use this type directly, use NewBenefitUtilizationSummariesAsyncClient() instead.
type BenefitUtilizationSummariesAsyncClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBenefitUtilizationSummariesAsyncClient creates a new instance of BenefitUtilizationSummariesAsyncClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBenefitUtilizationSummariesAsyncClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BenefitUtilizationSummariesAsyncClient, error) {
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
	client := &BenefitUtilizationSummariesAsyncClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// GetOperationStatusBillingAccountScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// billingAccountID - Enrollment ID (Legacy BillingAccount ID)
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusBillingAccountScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusBillingAccountScope(ctx context.Context, billingAccountID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse, error) {
	req, err := client.getOperationStatusBillingAccountScopeCreateRequest(ctx, billingAccountID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusBillingAccountScopeHandleResponse(resp)
}

// getOperationStatusBillingAccountScopeCreateRequest creates the GetOperationStatusBillingAccountScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusBillingAccountScopeCreateRequest(ctx context.Context, billingAccountID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/benefitUtilizationSummariesOperationResults/{operationId}"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusBillingAccountScopeHandleResponse handles the GetOperationStatusBillingAccountScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusBillingAccountScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingAccountScopeResponse{}, err
	}
	return result, nil
}

// GetOperationStatusBillingProfileScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// billingAccountID - BillingAccount ID
// billingProfileID - BillingProfile ID
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusBillingProfileScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusBillingProfileScope(ctx context.Context, billingAccountID string, billingProfileID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse, error) {
	req, err := client.getOperationStatusBillingProfileScopeCreateRequest(ctx, billingAccountID, billingProfileID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusBillingProfileScopeHandleResponse(resp)
}

// getOperationStatusBillingProfileScopeCreateRequest creates the GetOperationStatusBillingProfileScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusBillingProfileScopeCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.CostManagement/benefitsUtilizationSummariesOperationResults/{operationId}"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusBillingProfileScopeHandleResponse handles the GetOperationStatusBillingProfileScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusBillingProfileScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusBillingProfileScopeResponse{}, err
	}
	return result, nil
}

// GetOperationStatusReservationOrderScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// reservationOrderID - Reservation Order Id
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusReservationOrderScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusReservationOrderScope(ctx context.Context, reservationOrderID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse, error) {
	req, err := client.getOperationStatusReservationOrderScopeCreateRequest(ctx, reservationOrderID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusReservationOrderScopeHandleResponse(resp)
}

// getOperationStatusReservationOrderScopeCreateRequest creates the GetOperationStatusReservationOrderScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusReservationOrderScopeCreateRequest(ctx context.Context, reservationOrderID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.CostManagement/benefitUtilizationSummariesOperationResults/{operationId}"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusReservationOrderScopeHandleResponse handles the GetOperationStatusReservationOrderScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusReservationOrderScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationOrderScopeResponse{}, err
	}
	return result, nil
}

// GetOperationStatusReservationScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// reservationOrderID - Reservation Order Id
// reservationID - Reservation Id
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusReservationScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusReservationScope(ctx context.Context, reservationOrderID string, reservationID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse, error) {
	req, err := client.getOperationStatusReservationScopeCreateRequest(ctx, reservationOrderID, reservationID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusReservationScopeHandleResponse(resp)
}

// getOperationStatusReservationScopeCreateRequest creates the GetOperationStatusReservationScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusReservationScopeCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.CostManagement/benefitsUtilizationSummariesOperationResults/{operationId}"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusReservationScopeHandleResponse handles the GetOperationStatusReservationScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusReservationScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusReservationScopeResponse{}, err
	}
	return result, nil
}

// GetOperationStatusSavingsPlanOrderScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// savingsPlanOrderID - Savings Plan Order Id
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusSavingsPlanOrderScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusSavingsPlanOrderScope(ctx context.Context, savingsPlanOrderID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse, error) {
	req, err := client.getOperationStatusSavingsPlanOrderScopeCreateRequest(ctx, savingsPlanOrderID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusSavingsPlanOrderScopeHandleResponse(resp)
}

// getOperationStatusSavingsPlanOrderScopeCreateRequest creates the GetOperationStatusSavingsPlanOrderScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusSavingsPlanOrderScopeCreateRequest(ctx context.Context, savingsPlanOrderID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/providers/Microsoft.CostManagement/benefitUtilizationSummariesOperationResults/{operationId}"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusSavingsPlanOrderScopeHandleResponse handles the GetOperationStatusSavingsPlanOrderScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusSavingsPlanOrderScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanOrderScopeResponse{}, err
	}
	return result, nil
}

// GetOperationStatusSavingsPlanScope - Gets status of benefit utilization summaries report.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-15-preview
// savingsPlanOrderID - Savings Plan Order Id
// savingsPlanID - Savings Plan Id
// operationID - Id of the operation that is generating the benefits utilization summaries report.
// options - BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeOptions contains the optional parameters
// for the BenefitUtilizationSummariesAsyncClient.GetOperationStatusSavingsPlanScope method.
func (client *BenefitUtilizationSummariesAsyncClient) GetOperationStatusSavingsPlanScope(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeOptions) (BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse, error) {
	req, err := client.getOperationStatusSavingsPlanScopeCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, operationID, options)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusSavingsPlanScopeHandleResponse(resp)
}

// getOperationStatusSavingsPlanScopeCreateRequest creates the GetOperationStatusSavingsPlanScope request.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusSavingsPlanScopeCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, operationID string, options *BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}/providers/Microsoft.CostManagement/benefitsUtilizationSummariesOperationResults/{operationId}"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusSavingsPlanScopeHandleResponse handles the GetOperationStatusSavingsPlanScope response.
func (client *BenefitUtilizationSummariesAsyncClient) getOperationStatusSavingsPlanScopeHandleResponse(resp *http.Response) (BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse, error) {
	result := BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitsUtilizationSummariesOperationStatus); err != nil {
		return BenefitUtilizationSummariesAsyncClientGetOperationStatusSavingsPlanScopeResponse{}, err
	}
	return result, nil
}

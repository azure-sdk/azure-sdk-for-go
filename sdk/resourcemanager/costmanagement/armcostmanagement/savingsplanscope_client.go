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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SavingsPlanScopeClient contains the methods for the SavingsPlanScope group.
// Don't use this type directly, use NewSavingsPlanScopeClient() instead.
type SavingsPlanScopeClient struct {
	internal *arm.Client
}

// NewSavingsPlanScopeClient creates a new instance of SavingsPlanScopeClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSavingsPlanScopeClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SavingsPlanScopeClient, error) {
	cl, err := arm.NewClient(moduleName+".SavingsPlanScopeClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SavingsPlanScopeClient{
		internal: cl,
	}
	return client, nil
}

// BeginGenerateBenefitUtilizationSummariesReportAsync - Triggers generation of a benefit utilization summaries report for
// the provided savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - savingsPlanOrderID - Savings plan order ID.
//   - savingsPlanID - Savings plan ID.
//   - benefitUtilizationSummariesRequest - Async Benefit Utilization Summary report to be created.
//   - options - SavingsPlanScopeClientBeginGenerateBenefitUtilizationSummariesReportAsyncOptions contains the optional parameters
//     for the SavingsPlanScopeClient.BeginGenerateBenefitUtilizationSummariesReportAsync method.
func (client *SavingsPlanScopeClient) BeginGenerateBenefitUtilizationSummariesReportAsync(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, benefitUtilizationSummariesRequest BenefitUtilizationSummariesRequest, options *SavingsPlanScopeClientBeginGenerateBenefitUtilizationSummariesReportAsyncOptions) (*runtime.Poller[SavingsPlanScopeClientGenerateBenefitUtilizationSummariesReportAsyncResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateBenefitUtilizationSummariesReportAsync(ctx, savingsPlanOrderID, savingsPlanID, benefitUtilizationSummariesRequest, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SavingsPlanScopeClientGenerateBenefitUtilizationSummariesReportAsyncResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SavingsPlanScopeClientGenerateBenefitUtilizationSummariesReportAsyncResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GenerateBenefitUtilizationSummariesReportAsync - Triggers generation of a benefit utilization summaries report for the
// provided savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *SavingsPlanScopeClient) generateBenefitUtilizationSummariesReportAsync(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, benefitUtilizationSummariesRequest BenefitUtilizationSummariesRequest, options *SavingsPlanScopeClientBeginGenerateBenefitUtilizationSummariesReportAsyncOptions) (*http.Response, error) {
	req, err := client.generateBenefitUtilizationSummariesReportAsyncCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, benefitUtilizationSummariesRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// generateBenefitUtilizationSummariesReportAsyncCreateRequest creates the GenerateBenefitUtilizationSummariesReportAsync request.
func (client *SavingsPlanScopeClient) generateBenefitUtilizationSummariesReportAsyncCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, benefitUtilizationSummariesRequest BenefitUtilizationSummariesRequest, options *SavingsPlanScopeClientBeginGenerateBenefitUtilizationSummariesReportAsyncOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}/providers/Microsoft.CostManagement/generateBenefitUtilizationSummariesReport"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, benefitUtilizationSummariesRequest)
}

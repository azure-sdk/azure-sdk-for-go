//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ReservationRecommendationDetailsClient contains the methods for the ReservationRecommendationDetails group.
// Don't use this type directly, use NewReservationRecommendationDetailsClient() instead.
type ReservationRecommendationDetailsClient struct {
	internal *arm.Client
}

// NewReservationRecommendationDetailsClient creates a new instance of ReservationRecommendationDetailsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReservationRecommendationDetailsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationRecommendationDetailsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReservationRecommendationDetailsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationRecommendationDetailsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Details of a reservation recommendation for what-if analysis of reserved instances.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceScope - The scope associated with reservation recommendation details operations. This includes '/subscriptions/{subscriptionId}/'
//     for subscription scope,
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resource group scope, /providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
//     for BillingAccount scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile
//     scope
//   - scope - Scope of the reservation.
//   - region - Used to select the region the recommendation should be generated for.
//   - term - Specify length of reservation recommendation term.
//   - lookBackPeriod - Filter the time period on which reservation recommendation results are based.
//   - product - Filter the products for which reservation recommendation results are generated. Examples: StandardDS1v2 (for
//     VM), PremiumSSDManagedDisksP30 (for Managed Disks)
//   - options - ReservationRecommendationDetailsClientGetOptions contains the optional parameters for the ReservationRecommendationDetailsClient.Get
//     method.
func (client *ReservationRecommendationDetailsClient) Get(ctx context.Context, resourceScope string, scope Scope, region string, term Term, lookBackPeriod LookBackPeriod, product string, options *ReservationRecommendationDetailsClientGetOptions) (ReservationRecommendationDetailsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceScope, scope, region, term, lookBackPeriod, product, options)
	if err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReservationRecommendationDetailsClient) getCreateRequest(ctx context.Context, resourceScope string, scope Scope, region string, term Term, lookBackPeriod LookBackPeriod, product string, options *ReservationRecommendationDetailsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Consumption/reservationRecommendationDetails"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	reqQP.Set("scope", string(scope))
	reqQP.Set("region", region)
	reqQP.Set("term", string(term))
	reqQP.Set("lookBackPeriod", string(lookBackPeriod))
	reqQP.Set("product", product)
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReservationRecommendationDetailsClient) getHandleResponse(resp *http.Response) (ReservationRecommendationDetailsClientGetResponse, error) {
	result := ReservationRecommendationDetailsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationRecommendationDetailsModel); err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	return result, nil
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits

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

// SavingsPlanOrderClient contains the methods for the SavingsPlanOrder group.
// Don't use this type directly, use NewSavingsPlanOrderClient() instead.
type SavingsPlanOrderClient struct {
	host   string
	expand *string
	pl     runtime.Pipeline
}

// NewSavingsPlanOrderClient creates a new instance of SavingsPlanOrderClient with the specified values.
//   - expand - May be used to expand the detail information of some properties.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSavingsPlanOrderClient(expand *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SavingsPlanOrderClient, error) {
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
	client := &SavingsPlanOrderClient{
		expand: expand,
		host:   ep,
		pl:     pl,
	}
	return client, nil
}

// Elevate - Elevate as owner on savings plan order based on billing permissions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - savingsPlanOrderID - Order ID of the savings plan
//   - options - SavingsPlanOrderClientElevateOptions contains the optional parameters for the SavingsPlanOrderClient.Elevate
//     method.
func (client *SavingsPlanOrderClient) Elevate(ctx context.Context, savingsPlanOrderID string, options *SavingsPlanOrderClientElevateOptions) (SavingsPlanOrderClientElevateResponse, error) {
	req, err := client.elevateCreateRequest(ctx, savingsPlanOrderID, options)
	if err != nil {
		return SavingsPlanOrderClientElevateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanOrderClientElevateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavingsPlanOrderClientElevateResponse{}, runtime.NewResponseError(resp)
	}
	return client.elevateHandleResponse(resp)
}

// elevateCreateRequest creates the Elevate request.
func (client *SavingsPlanOrderClient) elevateCreateRequest(ctx context.Context, savingsPlanOrderID string, options *SavingsPlanOrderClientElevateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/elevate"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// elevateHandleResponse handles the Elevate response.
func (client *SavingsPlanOrderClient) elevateHandleResponse(resp *http.Response) (SavingsPlanOrderClientElevateResponse, error) {
	result := SavingsPlanOrderClientElevateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentEntity); err != nil {
		return SavingsPlanOrderClientElevateResponse{}, err
	}
	return result, nil
}

// Get - Get a savings plan order.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - savingsPlanOrderID - Order ID of the savings plan
//   - options - SavingsPlanOrderClientGetOptions contains the optional parameters for the SavingsPlanOrderClient.Get method.
func (client *SavingsPlanOrderClient) Get(ctx context.Context, savingsPlanOrderID string, options *SavingsPlanOrderClientGetOptions) (SavingsPlanOrderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, savingsPlanOrderID, options)
	if err != nil {
		return SavingsPlanOrderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanOrderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavingsPlanOrderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SavingsPlanOrderClient) getCreateRequest(ctx context.Context, savingsPlanOrderID string, options *SavingsPlanOrderClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if client.expand != nil {
		reqQP.Set("$expand", *client.expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavingsPlanOrderClient) getHandleResponse(resp *http.Response) (SavingsPlanOrderClientGetResponse, error) {
	result := SavingsPlanOrderClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanOrderModel); err != nil {
		return SavingsPlanOrderClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Savings plan orders.
//
// Generated from API version 2022-11-01
//   - options - SavingsPlanOrderClientListOptions contains the optional parameters for the SavingsPlanOrderClient.NewListPager
//     method.
func (client *SavingsPlanOrderClient) NewListPager(options *SavingsPlanOrderClientListOptions) *runtime.Pager[SavingsPlanOrderClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SavingsPlanOrderClientListResponse]{
		More: func(page SavingsPlanOrderClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SavingsPlanOrderClientListResponse) (SavingsPlanOrderClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SavingsPlanOrderClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SavingsPlanOrderClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SavingsPlanOrderClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SavingsPlanOrderClient) listCreateRequest(ctx context.Context, options *SavingsPlanOrderClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SavingsPlanOrderClient) listHandleResponse(resp *http.Response) (SavingsPlanOrderClientListResponse, error) {
	result := SavingsPlanOrderClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanOrderModelList); err != nil {
		return SavingsPlanOrderClientListResponse{}, err
	}
	return result, nil
}

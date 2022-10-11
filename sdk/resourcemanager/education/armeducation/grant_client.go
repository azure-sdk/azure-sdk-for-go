//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

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

// GrantClient contains the methods for the Grant group.
// Don't use this type directly, use NewGrantClient() instead.
type GrantClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGrantClient creates a new instance of GrantClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGrantClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GrantClient, error) {
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
	client := &GrantClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Renewal - Request grant renewal
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// options - GrantClientRenewalOptions contains the optional parameters for the GrantClient.Renewal method.
func (client *GrantClient) Renewal(ctx context.Context, billingAccountName string, billingProfileName string, options *GrantClientRenewalOptions) (GrantClientRenewalResponse, error) {
	req, err := client.renewalCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return GrantClientRenewalResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GrantClientRenewalResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return GrantClientRenewalResponse{}, runtime.NewResponseError(resp)
	}
	return client.renewalHandleResponse(resp)
}

// renewalCreateRequest creates the Renewal request.
func (client *GrantClient) renewalCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *GrantClientRenewalOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Education/grants/default/renew"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// renewalHandleResponse handles the Renewal response.
func (client *GrantClient) renewalHandleResponse(resp *http.Response) (GrantClientRenewalResponse, error) {
	result := GrantClientRenewalResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		result.RetryAfter = &val
	}
	return result, nil
}

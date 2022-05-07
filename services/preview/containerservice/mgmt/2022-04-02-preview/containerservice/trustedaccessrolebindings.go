package containerservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TrustedAccessRoleBindingsClient is the the Container Service Client.
type TrustedAccessRoleBindingsClient struct {
	BaseClient
}

// NewTrustedAccessRoleBindingsClient creates an instance of the TrustedAccessRoleBindingsClient client.
func NewTrustedAccessRoleBindingsClient(subscriptionID string) TrustedAccessRoleBindingsClient {
	return NewTrustedAccessRoleBindingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTrustedAccessRoleBindingsClientWithBaseURI creates an instance of the TrustedAccessRoleBindingsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewTrustedAccessRoleBindingsClientWithBaseURI(baseURI string, subscriptionID string) TrustedAccessRoleBindingsClient {
	return TrustedAccessRoleBindingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// trustedAccessRoleBindingName - the name of trusted access role binding.
// trustedAccessRoleBinding - a trusted access role binding
func (client TrustedAccessRoleBindingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding) (result TrustedAccessRoleBinding, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrustedAccessRoleBindingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}},
		{TargetValue: trustedAccessRoleBindingName,
			Constraints: []validation.Constraint{{Target: "trustedAccessRoleBindingName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "trustedAccessRoleBindingName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: trustedAccessRoleBinding,
			Constraints: []validation.Constraint{{Target: "trustedAccessRoleBinding.TrustedAccessRoleBindingProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "trustedAccessRoleBinding.TrustedAccessRoleBindingProperties.SourceResourceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "trustedAccessRoleBinding.TrustedAccessRoleBindingProperties.Roles", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("containerservice.TrustedAccessRoleBindingsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, trustedAccessRoleBinding)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TrustedAccessRoleBindingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"resourceName":                 autorest.Encode("path", resourceName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"trustedAccessRoleBindingName": autorest.Encode("path", trustedAccessRoleBindingName),
	}

	const APIVersion = "2022-04-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}", pathParameters),
		autorest.WithJSON(trustedAccessRoleBinding),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedAccessRoleBindingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TrustedAccessRoleBindingsClient) CreateOrUpdateResponder(resp *http.Response) (result TrustedAccessRoleBinding, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// trustedAccessRoleBindingName - the name of trusted access role binding.
func (client TrustedAccessRoleBindingsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrustedAccessRoleBindingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}},
		{TargetValue: trustedAccessRoleBindingName,
			Constraints: []validation.Constraint{{Target: "trustedAccessRoleBindingName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "trustedAccessRoleBindingName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.TrustedAccessRoleBindingsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TrustedAccessRoleBindingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"resourceName":                 autorest.Encode("path", resourceName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"trustedAccessRoleBindingName": autorest.Encode("path", trustedAccessRoleBindingName),
	}

	const APIVersion = "2022-04-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedAccessRoleBindingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TrustedAccessRoleBindingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// trustedAccessRoleBindingName - the name of trusted access role binding.
func (client TrustedAccessRoleBindingsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string) (result TrustedAccessRoleBinding, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrustedAccessRoleBindingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}},
		{TargetValue: trustedAccessRoleBindingName,
			Constraints: []validation.Constraint{{Target: "trustedAccessRoleBindingName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "trustedAccessRoleBindingName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.TrustedAccessRoleBindingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TrustedAccessRoleBindingsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"resourceName":                 autorest.Encode("path", resourceName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
		"trustedAccessRoleBindingName": autorest.Encode("path", trustedAccessRoleBindingName),
	}

	const APIVersion = "2022-04-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedAccessRoleBindingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TrustedAccessRoleBindingsClient) GetResponder(resp *http.Response) (result TrustedAccessRoleBinding, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
func (client TrustedAccessRoleBindingsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result TrustedAccessRoleBindingListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrustedAccessRoleBindingsClient.List")
		defer func() {
			sc := -1
			if result.tarblr.Response.Response != nil {
				sc = result.tarblr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.TrustedAccessRoleBindingsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tarblr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "List", resp, "Failure sending request")
		return
	}

	result.tarblr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tarblr.hasNextLink() && result.tarblr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TrustedAccessRoleBindingsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-04-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedAccessRoleBindingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TrustedAccessRoleBindingsClient) ListResponder(resp *http.Response) (result TrustedAccessRoleBindingListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TrustedAccessRoleBindingsClient) listNextResults(ctx context.Context, lastResults TrustedAccessRoleBindingListResult) (result TrustedAccessRoleBindingListResult, err error) {
	req, err := lastResults.trustedAccessRoleBindingListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.TrustedAccessRoleBindingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TrustedAccessRoleBindingsClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result TrustedAccessRoleBindingListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TrustedAccessRoleBindingsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, resourceName)
	return
}

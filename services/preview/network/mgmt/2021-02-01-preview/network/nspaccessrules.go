package network

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

// NspAccessRulesClient is the network Client
type NspAccessRulesClient struct {
	BaseClient
}

// NewNspAccessRulesClient creates an instance of the NspAccessRulesClient client.
func NewNspAccessRulesClient(subscriptionID string) NspAccessRulesClient {
	return NewNspAccessRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNspAccessRulesClientWithBaseURI creates an instance of the NspAccessRulesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNspAccessRulesClientWithBaseURI(baseURI string, subscriptionID string) NspAccessRulesClient {
	return NspAccessRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a network access rule.
// Parameters:
// parameters - parameters that hold the NspAccessRule resource to be created/updated.
// resourceGroupName - the name of the resource group.
// networkSecurityPerimeterName - the name of the network security perimeter.
// profileName - the name of the NSP profile.
// accessRuleName - the name of the NSP access rule.
func (client NspAccessRulesClient) CreateOrUpdate(ctx context.Context, parameters NspAccessRule, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (result NspAccessRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NspAccessRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, networkSecurityPerimeterName, profileName, accessRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NspAccessRulesClient) CreateOrUpdatePreparer(ctx context.Context, parameters NspAccessRule, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessRuleName":               autorest.Encode("path", accessRuleName),
		"networkSecurityPerimeterName": autorest.Encode("path", networkSecurityPerimeterName),
		"profileName":                  autorest.Encode("path", profileName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/profiles/{profileName}/accessRules/{accessRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NspAccessRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client NspAccessRulesClient) CreateOrUpdateResponder(resp *http.Response) (result NspAccessRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an NSP access rule.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityPerimeterName - the name of the network security perimeter.
// profileName - the name of the NSP profile.
// accessRuleName - the name of the NSP access rule.
func (client NspAccessRulesClient) Delete(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NspAccessRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkSecurityPerimeterName, profileName, accessRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NspAccessRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessRuleName":               autorest.Encode("path", accessRuleName),
		"networkSecurityPerimeterName": autorest.Encode("path", networkSecurityPerimeterName),
		"profileName":                  autorest.Encode("path", profileName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/profiles/{profileName}/accessRules/{accessRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NspAccessRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NspAccessRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified NSP access rule by name.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityPerimeterName - the name of the network security perimeter.
// profileName - the name of the NSP profile.
// accessRuleName - the name of the NSP access rule.
func (client NspAccessRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (result NspAccessRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NspAccessRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkSecurityPerimeterName, profileName, accessRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client NspAccessRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessRuleName":               autorest.Encode("path", accessRuleName),
		"networkSecurityPerimeterName": autorest.Encode("path", networkSecurityPerimeterName),
		"profileName":                  autorest.Encode("path", profileName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/profiles/{profileName}/accessRules/{accessRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NspAccessRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NspAccessRulesClient) GetResponder(resp *http.Response) (result NspAccessRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the NSP access rules in the specified NSP profile.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityPerimeterName - the name of the network security perimeter.
// profileName - the name of the NSP profile.
// top - an optional query parameter which specifies the maximum number of records to be returned by the
// server.
// skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
// specifies a starting point to use for subsequent calls.
func (client NspAccessRulesClient) List(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, top *int32, skipToken string) (result NspAccessRuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NspAccessRulesClient.List")
		defer func() {
			sc := -1
			if result.narlr.Response.Response != nil {
				sc = result.narlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.NspAccessRulesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkSecurityPerimeterName, profileName, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.narlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.narlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.narlr.hasNextLink() && result.narlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client NspAccessRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityPerimeterName": autorest.Encode("path", networkSecurityPerimeterName),
		"profileName":                  autorest.Encode("path", profileName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/profiles/{profileName}/accessRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NspAccessRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NspAccessRulesClient) ListResponder(resp *http.Response) (result NspAccessRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client NspAccessRulesClient) listNextResults(ctx context.Context, lastResults NspAccessRuleListResult) (result NspAccessRuleListResult, err error) {
	req, err := lastResults.nspAccessRuleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NspAccessRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client NspAccessRulesClient) ListComplete(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, top *int32, skipToken string) (result NspAccessRuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NspAccessRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkSecurityPerimeterName, profileName, top, skipToken)
	return
}

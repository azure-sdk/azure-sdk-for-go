package eventhub

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

// ApplicationGroupClient is the client for the ApplicationGroup methods of the Eventhub service.
type ApplicationGroupClient struct {
	BaseClient
}

// NewApplicationGroupClient creates an instance of the ApplicationGroupClient client.
func NewApplicationGroupClient(subscriptionID string) ApplicationGroupClient {
	return NewApplicationGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGroupClientWithBaseURI creates an instance of the ApplicationGroupClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewApplicationGroupClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGroupClient {
	return ApplicationGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateApplicationGroup creates or updates an ApplicationGroup for a Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// applicationGroupName - the Application Group name
// parameters - the ApplicationGroup.
func (client ApplicationGroupClient) CreateOrUpdateApplicationGroup(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string, parameters ApplicationGroup) (result ApplicationGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupClient.CreateOrUpdateApplicationGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationGroupProperties.ClientAppGroupIdentifier", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("eventhub.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", err.Error())
	}

	req, err := client.CreateOrUpdateApplicationGroupPreparer(ctx, resourceGroupName, namespaceName, applicationGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateApplicationGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateApplicationGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdateApplicationGroupPreparer prepares the CreateOrUpdateApplicationGroup request.
func (client ApplicationGroupClient) CreateOrUpdateApplicationGroupPreparer(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string, parameters ApplicationGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"namespaceName":        autorest.Encode("path", namespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateApplicationGroupSender sends the CreateOrUpdateApplicationGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupClient) CreateOrUpdateApplicationGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateApplicationGroupResponder handles the response to the CreateOrUpdateApplicationGroup request. The method always
// closes the http.Response Body.
func (client ApplicationGroupClient) CreateOrUpdateApplicationGroupResponder(resp *http.Response) (result ApplicationGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an ApplicationGroup for a Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// applicationGroupName - the Application Group name
func (client ApplicationGroupClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ApplicationGroupClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, namespaceName, applicationGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGroupClient) DeletePreparer(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"namespaceName":        autorest.Encode("path", namespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an ApplicationGroup for a Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// applicationGroupName - the Application Group name
func (client ApplicationGroupClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string) (result ApplicationGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ApplicationGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, namespaceName, applicationGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"namespaceName":        autorest.Encode("path", namespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGroupClient) GetResponder(resp *http.Response) (result ApplicationGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByNamespace gets a list of application groups for a Namespace.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
func (client ApplicationGroupClient) ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result ApplicationGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupClient.ListByNamespace")
		defer func() {
			sc := -1
			if result.aglr.Response.Response != nil {
				sc = result.aglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventhub.ApplicationGroupClient", "ListByNamespace", err.Error())
	}

	result.fn = client.listByNamespaceNextResults
	req, err := client.ListByNamespacePreparer(ctx, resourceGroupName, namespaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "ListByNamespace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByNamespaceSender(req)
	if err != nil {
		result.aglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "ListByNamespace", resp, "Failure sending request")
		return
	}

	result.aglr, err = client.ListByNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "ListByNamespace", resp, "Failure responding to request")
		return
	}
	if result.aglr.hasNextLink() && result.aglr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByNamespacePreparer prepares the ListByNamespace request.
func (client ApplicationGroupClient) ListByNamespacePreparer(ctx context.Context, resourceGroupName string, namespaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/applicationGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByNamespaceSender sends the ListByNamespace request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupClient) ListByNamespaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByNamespaceResponder handles the response to the ListByNamespace request. The method always
// closes the http.Response Body.
func (client ApplicationGroupClient) ListByNamespaceResponder(resp *http.Response) (result ApplicationGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByNamespaceNextResults retrieves the next set of results, if any.
func (client ApplicationGroupClient) listByNamespaceNextResults(ctx context.Context, lastResults ApplicationGroupListResult) (result ApplicationGroupListResult, err error) {
	req, err := lastResults.applicationGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "listByNamespaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByNamespaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "listByNamespaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ApplicationGroupClient", "listByNamespaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByNamespaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGroupClient) ListByNamespaceComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result ApplicationGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupClient.ListByNamespace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByNamespace(ctx, resourceGroupName, namespaceName)
	return
}

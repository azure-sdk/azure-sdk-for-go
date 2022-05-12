package kubernetesconfiguration

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

// LocationExtensionTypesClient is the kubernetesConfiguration Client
type LocationExtensionTypesClient struct {
	BaseClient
}

// NewLocationExtensionTypesClient creates an instance of the LocationExtensionTypesClient client.
func NewLocationExtensionTypesClient(subscriptionID string) LocationExtensionTypesClient {
	return NewLocationExtensionTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationExtensionTypesClientWithBaseURI creates an instance of the LocationExtensionTypesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewLocationExtensionTypesClientWithBaseURI(baseURI string, subscriptionID string) LocationExtensionTypesClient {
	return LocationExtensionTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list all Extension Types
// Parameters:
// location - extension location
func (client LocationExtensionTypesClient) List(ctx context.Context, location string) (result ExtensionTypeListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationExtensionTypesClient.List")
		defer func() {
			sc := -1
			if result.etl.Response.Response != nil {
				sc = result.etl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.LocationExtensionTypesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.etl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "List", resp, "Failure sending request")
		return
	}

	result.etl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.etl.hasNextLink() && result.etl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LocationExtensionTypesClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.KubernetesConfiguration/locations/{location}/extensionTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocationExtensionTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocationExtensionTypesClient) ListResponder(resp *http.Response) (result ExtensionTypeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LocationExtensionTypesClient) listNextResults(ctx context.Context, lastResults ExtensionTypeList) (result ExtensionTypeList, err error) {
	req, err := lastResults.extensionTypeListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.LocationExtensionTypesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationExtensionTypesClient) ListComplete(ctx context.Context, location string) (result ExtensionTypeListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationExtensionTypesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}

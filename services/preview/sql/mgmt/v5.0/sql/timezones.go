package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TimeZonesClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type TimeZonesClient struct {
	BaseClient
}

// NewTimeZonesClient creates an instance of the TimeZonesClient client.
func NewTimeZonesClient(subscriptionID string) TimeZonesClient {
	return NewTimeZonesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTimeZonesClientWithBaseURI creates an instance of the TimeZonesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTimeZonesClientWithBaseURI(baseURI string, subscriptionID string) TimeZonesClient {
	return TimeZonesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a managed instance time zone.
func (client TimeZonesClient) Get(ctx context.Context, locationName string, timeZoneID string) (result TimeZone, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TimeZonesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, locationName, timeZoneID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TimeZonesClient) GetPreparer(ctx context.Context, locationName string, timeZoneID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"timeZoneId":     autorest.Encode("path", timeZoneID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones/{timeZoneId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TimeZonesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TimeZonesClient) GetResponder(resp *http.Response) (result TimeZone, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByLocation gets a list of managed instance time zones by location.
func (client TimeZonesClient) ListByLocation(ctx context.Context, locationName string) (result TimeZoneListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TimeZonesClient.ListByLocation")
		defer func() {
			sc := -1
			if result.tzlr.Response.Response != nil {
				sc = result.tzlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByLocationNextResults
	req, err := client.ListByLocationPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.tzlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result.tzlr, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "ListByLocation", resp, "Failure responding to request")
		return
	}
	if result.tzlr.hasNextLink() && result.tzlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client TimeZonesClient) ListByLocationPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client TimeZonesClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client TimeZonesClient) ListByLocationResponder(resp *http.Response) (result TimeZoneListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLocationNextResults retrieves the next set of results, if any.
func (client TimeZonesClient) listByLocationNextResults(ctx context.Context, lastResults TimeZoneListResult) (result TimeZoneListResult, err error) {
	req, err := lastResults.timeZoneListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.TimeZonesClient", "listByLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.TimeZonesClient", "listByLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.TimeZonesClient", "listByLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client TimeZonesClient) ListByLocationComplete(ctx context.Context, locationName string) (result TimeZoneListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TimeZonesClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLocation(ctx, locationName)
	return
}

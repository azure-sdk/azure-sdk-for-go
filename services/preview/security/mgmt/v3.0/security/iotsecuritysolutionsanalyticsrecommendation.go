package security

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

// IotSecuritySolutionsAnalyticsRecommendationClient is the API spec for Microsoft.Security (Azure Security Center)
// resource provider
type IotSecuritySolutionsAnalyticsRecommendationClient struct {
	BaseClient
}

// NewIotSecuritySolutionsAnalyticsRecommendationClient creates an instance of the
// IotSecuritySolutionsAnalyticsRecommendationClient client.
func NewIotSecuritySolutionsAnalyticsRecommendationClient(subscriptionID string) IotSecuritySolutionsAnalyticsRecommendationClient {
	return NewIotSecuritySolutionsAnalyticsRecommendationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIotSecuritySolutionsAnalyticsRecommendationClientWithBaseURI creates an instance of the
// IotSecuritySolutionsAnalyticsRecommendationClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIotSecuritySolutionsAnalyticsRecommendationClientWithBaseURI(baseURI string, subscriptionID string) IotSecuritySolutionsAnalyticsRecommendationClient {
	return IotSecuritySolutionsAnalyticsRecommendationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get use this method to get the aggregated security analytics recommendation of yours IoT Security solution. This
// aggregation is performed by recommendation name.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the name of the IoT Security solution.
// aggregatedRecommendationName - name of the recommendation aggregated for this query.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string) (result IoTSecurityAggregatedRecommendation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotSecuritySolutionsAnalyticsRecommendationClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotSecuritySolutionsAnalyticsRecommendationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, solutionName, aggregatedRecommendationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) GetPreparer(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"aggregatedRecommendationName": autorest.Encode("path", aggregatedRecommendationName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"solutionName":                 autorest.Encode("path", solutionName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations/{aggregatedRecommendationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) GetResponder(resp *http.Response) (result IoTSecurityAggregatedRecommendation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List use this method to get the list of aggregated security analytics recommendations of yours IoT Security
// solution.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the name of the IoT Security solution.
// top - number of results to retrieve.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) List(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result IoTSecurityAggregatedRecommendationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotSecuritySolutionsAnalyticsRecommendationClient.List")
		defer func() {
			sc := -1
			if result.itsarl.Response.Response != nil {
				sc = result.itsarl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IotSecuritySolutionsAnalyticsRecommendationClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, solutionName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itsarl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "List", resp, "Failure sending request")
		return
	}

	result.itsarl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "List", resp, "Failure responding to request")
		return
	}
	if result.itsarl.hasNextLink() && result.itsarl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) ListPreparer(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"solutionName":      autorest.Encode("path", solutionName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) ListResponder(resp *http.Response) (result IoTSecurityAggregatedRecommendationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) listNextResults(ctx context.Context, lastResults IoTSecurityAggregatedRecommendationList) (result IoTSecurityAggregatedRecommendationList, err error) {
	req, err := lastResults.ioTSecurityAggregatedRecommendationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionsAnalyticsRecommendationClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IotSecuritySolutionsAnalyticsRecommendationClient) ListComplete(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result IoTSecurityAggregatedRecommendationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotSecuritySolutionsAnalyticsRecommendationClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, solutionName, top)
	return
}

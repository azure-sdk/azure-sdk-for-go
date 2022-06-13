package datacollaboration

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

// PipelineRunsClient is the creates a Microsoft.DataCollaboration management client.
type PipelineRunsClient struct {
	BaseClient
}

// NewPipelineRunsClient creates an instance of the PipelineRunsClient client.
func NewPipelineRunsClient(subscriptionID string) PipelineRunsClient {
	return NewPipelineRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelineRunsClientWithBaseURI creates an instance of the PipelineRunsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPipelineRunsClientWithBaseURI(baseURI string, subscriptionID string) PipelineRunsClient {
	return PipelineRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel request to cancel a pipeline run.
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// pipelineRunName - the name of the pipeline run.
func (client PipelineRunsClient) Cancel(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (result PipelineRunsCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.Cancel")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, workspaceName, pipelineRunName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client PipelineRunsClient) CancelPreparer(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineRunName":   autorest.Encode("path", pipelineRunName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelineRuns/{pipelineRunName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) CancelSender(req *http.Request) (future PipelineRunsCancelFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) CancelResponder(resp *http.Response) (result PipelineRun, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a pipeline run from a workspace
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// pipelineRunName - the name of the pipeline run.
func (client PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (result PipelineRun, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, pipelineRunName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelineRunsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineRunName":   autorest.Encode("path", pipelineRunName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelineRuns/{pipelineRunName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) GetResponder(resp *http.Response) (result PipelineRun, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace list pipelines runs in a workspace
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// skipToken - continuation token
// filter - filters the results using OData syntax.
// orderby - sorts the results using OData syntax.
func (client PipelineRunsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result PipelineRunListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.prl.Response.Response != nil {
				sc = result.prl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName, skipToken, filter, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.prl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.prl, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.prl.hasNextLink() && result.prl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client PipelineRunsClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelineRuns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) ListByWorkspaceResponder(resp *http.Response) (result PipelineRunList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client PipelineRunsClient) listByWorkspaceNextResults(ctx context.Context, lastResults PipelineRunList) (result PipelineRunList, err error) {
	req, err := lastResults.pipelineRunListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelineRunsClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result PipelineRunListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName, skipToken, filter, orderby)
	return
}

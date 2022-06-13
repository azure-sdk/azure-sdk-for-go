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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ParticipantsClient is the creates a Microsoft.DataCollaboration management client.
type ParticipantsClient struct {
	BaseClient
}

// NewParticipantsClient creates an instance of the ParticipantsClient client.
func NewParticipantsClient(subscriptionID string) ParticipantsClient {
	return NewParticipantsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewParticipantsClientWithBaseURI creates an instance of the ParticipantsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewParticipantsClientWithBaseURI(baseURI string, subscriptionID string) ParticipantsClient {
	return ParticipantsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Participant in a Proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// participantName - the name of the participant.
// participant - the new participant information.
func (client ParticipantsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string, participant Participant) (result Participant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ParticipantsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: participant,
			Constraints: []validation.Constraint{{Target: "participant.ParticipantProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "participant.ParticipantProperties.PlaceholderInfo", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "participant.ParticipantProperties.PlaceholderInfo.DisplayName", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("datacollaboration.ParticipantsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, proposalName, participantName, participant)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ParticipantsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string, participant Participant) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"participantName":   autorest.Encode("path", participantName),
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/participants/{participantName}", pathParameters),
		autorest.WithJSON(participant),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ParticipantsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ParticipantsClient) CreateOrUpdateResponder(resp *http.Response) (result Participant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Participant in a Proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// participantName - the name of the participant.
func (client ParticipantsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (result ParticipantsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ParticipantsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, proposalName, participantName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ParticipantsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"participantName":   autorest.Encode("path", participantName),
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/participants/{participantName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ParticipantsClient) DeleteSender(req *http.Request) (future ParticipantsDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ParticipantsClient) DeleteResponder(resp *http.Response) (result OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a Participant in a Proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// participantName - the name of the participant.
func (client ParticipantsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (result Participant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ParticipantsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, proposalName, participantName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ParticipantsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"participantName":   autorest.Encode("path", participantName),
		"proposalName":      autorest.Encode("path", proposalName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/participants/{participantName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ParticipantsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ParticipantsClient) GetResponder(resp *http.Response) (result Participant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProposal list Participants in a Proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// skipToken - continuation token
func (client ParticipantsClient) ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result ParticipantListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ParticipantsClient.ListByProposal")
		defer func() {
			sc := -1
			if result.pl.Response.Response != nil {
				sc = result.pl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByProposalNextResults
	req, err := client.ListByProposalPreparer(ctx, resourceGroupName, workspaceName, proposalName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "ListByProposal", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProposalSender(req)
	if err != nil {
		result.pl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "ListByProposal", resp, "Failure sending request")
		return
	}

	result.pl, err = client.ListByProposalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "ListByProposal", resp, "Failure responding to request")
		return
	}
	if result.pl.hasNextLink() && result.pl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByProposalPreparer prepares the ListByProposal request.
func (client ParticipantsClient) ListByProposalPreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"proposalName":      autorest.Encode("path", proposalName),
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

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/participants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProposalSender sends the ListByProposal request. The method will close the
// http.Response Body if it receives an error.
func (client ParticipantsClient) ListByProposalSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProposalResponder handles the response to the ListByProposal request. The method always
// closes the http.Response Body.
func (client ParticipantsClient) ListByProposalResponder(resp *http.Response) (result ParticipantList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProposalNextResults retrieves the next set of results, if any.
func (client ParticipantsClient) listByProposalNextResults(ctx context.Context, lastResults ParticipantList) (result ParticipantList, err error) {
	req, err := lastResults.participantListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "listByProposalNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProposalSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "listByProposalNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProposalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.ParticipantsClient", "listByProposalNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProposalComplete enumerates all values, automatically crossing page boundaries as required.
func (client ParticipantsClient) ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result ParticipantListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ParticipantsClient.ListByProposal")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProposal(ctx, resourceGroupName, workspaceName, proposalName, skipToken)
	return
}

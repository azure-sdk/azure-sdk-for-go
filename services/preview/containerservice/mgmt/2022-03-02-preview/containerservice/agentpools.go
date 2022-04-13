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

// AgentPoolsClient is the the Container Service Client.
type AgentPoolsClient struct {
	BaseClient
}

// NewAgentPoolsClient creates an instance of the AgentPoolsClient client.
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return NewAgentPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAgentPoolsClientWithBaseURI creates an instance of the AgentPoolsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return AgentPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// agentPoolName - the name of the agent pool.
// parameters - the agent pool to create or update.
func (client AgentPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool) (result AgentPoolsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
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
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ManagedClusterAgentPoolProfileProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ManagedClusterAgentPoolProfileProperties.KubeletConfig", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ManagedClusterAgentPoolProfileProperties.KubeletConfig.ContainerLogMaxFiles", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ManagedClusterAgentPoolProfileProperties.KubeletConfig.ContainerLogMaxFiles", Name: validation.InclusiveMinimum, Rule: int64(2), Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("containerservice.AgentPoolsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, agentPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AgentPoolsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) CreateOrUpdateSender(req *http.Request) (future AgentPoolsCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// agentPoolName - the name of the agent pool.
// ignorePodDisruptionBudget - ignore-pod-disruption-budget=true to delete those pods on a node without
// considering Pod Disruption Budget
func (client AgentPoolsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, ignorePodDisruptionBudget *bool) (result AgentPoolsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
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
		return result, validation.NewError("containerservice.AgentPoolsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, agentPoolName, ignorePodDisruptionBudget)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AgentPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, ignorePodDisruptionBudget *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if ignorePodDisruptionBudget != nil {
		queryParameters["ignore-pod-disruption-budget"] = autorest.Encode("query", *ignorePodDisruptionBudget)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) DeleteSender(req *http.Request) (future AgentPoolsDeleteFuture, err error) {
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
func (client AgentPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result AgentPool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Get")
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
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.AgentPoolsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AgentPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) GetResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAvailableAgentPoolVersions see [supported Kubernetes
// versions](https://docs.microsoft.com/azure/aks/supported-kubernetes-versions) for more details about the version
// lifecycle.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
func (client AgentPoolsClient) GetAvailableAgentPoolVersions(ctx context.Context, resourceGroupName string, resourceName string) (result AgentPoolAvailableVersions, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.GetAvailableAgentPoolVersions")
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
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.AgentPoolsClient", "GetAvailableAgentPoolVersions", err.Error())
	}

	req, err := client.GetAvailableAgentPoolVersionsPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetAvailableAgentPoolVersions", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAvailableAgentPoolVersionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetAvailableAgentPoolVersions", resp, "Failure sending request")
		return
	}

	result, err = client.GetAvailableAgentPoolVersionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetAvailableAgentPoolVersions", resp, "Failure responding to request")
		return
	}

	return
}

// GetAvailableAgentPoolVersionsPreparer prepares the GetAvailableAgentPoolVersions request.
func (client AgentPoolsClient) GetAvailableAgentPoolVersionsPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableAgentPoolVersions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAvailableAgentPoolVersionsSender sends the GetAvailableAgentPoolVersions request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) GetAvailableAgentPoolVersionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAvailableAgentPoolVersionsResponder handles the response to the GetAvailableAgentPoolVersions request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) GetAvailableAgentPoolVersionsResponder(resp *http.Response) (result AgentPoolAvailableVersions, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUpgradeProfile sends the get upgrade profile request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result AgentPoolUpgradeProfile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.GetUpgradeProfile")
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
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]$|^[a-zA-Z0-9][-_a-zA-Z0-9]{0,61}[a-zA-Z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerservice.AgentPoolsClient", "GetUpgradeProfile", err.Error())
	}

	req, err := client.GetUpgradeProfilePreparer(ctx, resourceGroupName, resourceName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetUpgradeProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUpgradeProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetUpgradeProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetUpgradeProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "GetUpgradeProfile", resp, "Failure responding to request")
		return
	}

	return
}

// GetUpgradeProfilePreparer prepares the GetUpgradeProfile request.
func (client AgentPoolsClient) GetUpgradeProfilePreparer(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUpgradeProfileSender sends the GetUpgradeProfile request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) GetUpgradeProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetUpgradeProfileResponder handles the response to the GetUpgradeProfile request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) GetUpgradeProfileResponder(resp *http.Response) (result AgentPoolUpgradeProfile, err error) {
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
func (client AgentPoolsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result AgentPoolListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.List")
		defer func() {
			sc := -1
			if result.aplr.Response.Response != nil {
				sc = result.aplr.Response.Response.StatusCode
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
		return result, validation.NewError("containerservice.AgentPoolsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.aplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.aplr.hasNextLink() && result.aplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AgentPoolsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) ListResponder(resp *http.Response) (result AgentPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AgentPoolsClient) listNextResults(ctx context.Context, lastResults AgentPoolListResult) (result AgentPoolListResult, err error) {
	req, err := lastResults.agentPoolListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AgentPoolsClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result AgentPoolListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.List")
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

// UpgradeNodeImageVersion upgrading the node image version of an agent pool applies the newest OS and runtime updates
// to the nodes. AKS provides one new image per week with the latest updates. For more details on node image versions,
// see: https://docs.microsoft.com/azure/aks/node-image-upgrade
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the managed cluster resource.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) UpgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result AgentPoolsUpgradeNodeImageVersionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.UpgradeNodeImageVersion")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
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
		return result, validation.NewError("containerservice.AgentPoolsClient", "UpgradeNodeImageVersion", err.Error())
	}

	req, err := client.UpgradeNodeImageVersionPreparer(ctx, resourceGroupName, resourceName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "UpgradeNodeImageVersion", nil, "Failure preparing request")
		return
	}

	result, err = client.UpgradeNodeImageVersionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.AgentPoolsClient", "UpgradeNodeImageVersion", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpgradeNodeImageVersionPreparer prepares the UpgradeNodeImageVersion request.
func (client AgentPoolsClient) UpgradeNodeImageVersionPreparer(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-03-02-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeNodeImageVersion", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpgradeNodeImageVersionSender sends the UpgradeNodeImageVersion request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) UpgradeNodeImageVersionSender(req *http.Request) (future AgentPoolsUpgradeNodeImageVersionFuture, err error) {
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

// UpgradeNodeImageVersionResponder handles the response to the UpgradeNodeImageVersion request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) UpgradeNodeImageVersionResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

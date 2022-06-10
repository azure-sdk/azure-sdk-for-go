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

// FluxConfigOperationStatusClient is the kubernetesConfiguration Client
type FluxConfigOperationStatusClient struct {
	BaseClient
}

// NewFluxConfigOperationStatusClient creates an instance of the FluxConfigOperationStatusClient client.
func NewFluxConfigOperationStatusClient(subscriptionID string) FluxConfigOperationStatusClient {
	return NewFluxConfigOperationStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFluxConfigOperationStatusClientWithBaseURI creates an instance of the FluxConfigOperationStatusClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewFluxConfigOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) FluxConfigOperationStatusClient {
	return FluxConfigOperationStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get Async Operation status
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// clusterRp - the Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes,
// Microsoft.HybridContainerService.
// clusterResourceName - the Kubernetes cluster resource name - i.e. managedClusters, connectedClusters,
// provisionedClusters.
// clusterName - the name of the kubernetes cluster.
// fluxConfigurationName - name of the Flux Configuration.
// operationID - operation Id
func (client FluxConfigOperationStatusClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, operationID string) (result OperationStatusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FluxConfigOperationStatusClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetesconfiguration.FluxConfigOperationStatusClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigOperationStatusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigOperationStatusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.FluxConfigOperationStatusClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FluxConfigOperationStatusClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, fluxConfigurationName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":           autorest.Encode("path", clusterName),
		"clusterResourceName":   autorest.Encode("path", clusterResourceName),
		"clusterRp":             autorest.Encode("path", clusterRp),
		"fluxConfigurationName": autorest.Encode("path", fluxConfigurationName),
		"operationId":           autorest.Encode("path", operationID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/{fluxConfigurationName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FluxConfigOperationStatusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FluxConfigOperationStatusClient) GetResponder(resp *http.Response) (result OperationStatusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

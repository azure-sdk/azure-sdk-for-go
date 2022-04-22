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

// SecurityUserConfigurationsClient is the network Client
type SecurityUserConfigurationsClient struct {
	BaseClient
}

// NewSecurityUserConfigurationsClient creates an instance of the SecurityUserConfigurationsClient client.
func NewSecurityUserConfigurationsClient(subscriptionID string) SecurityUserConfigurationsClient {
	return NewSecurityUserConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityUserConfigurationsClientWithBaseURI creates an instance of the SecurityUserConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSecurityUserConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SecurityUserConfigurationsClient {
	return SecurityUserConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a network manager security user configuration.
// Parameters:
// securityUserConfiguration - the security user configuration to create or update
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// configurationName - the name of the network manager security Configuration.
func (client SecurityUserConfigurationsClient) CreateOrUpdate(ctx context.Context, securityUserConfiguration SecurityConfiguration, resourceGroupName string, networkManagerName string, configurationName string) (result SecurityConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityUserConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, securityUserConfiguration, resourceGroupName, networkManagerName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityUserConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, securityUserConfiguration SecurityConfiguration, resourceGroupName string, networkManagerName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName":  autorest.Encode("path", configurationName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	securityUserConfiguration.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}", pathParameters),
		autorest.WithJSON(securityUserConfiguration),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityUserConfigurationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityUserConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a network manager security user configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// configurationName - the name of the network manager security Configuration.
func (client SecurityUserConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityUserConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityUserConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName":  autorest.Encode("path", configurationName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityUserConfigurationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityUserConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves a network manager security user configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// configurationName - the name of the network manager security Configuration.
func (client SecurityUserConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (result SecurityConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityUserConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityUserConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName":  autorest.Encode("path", configurationName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityUserConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityUserConfigurationsClient) GetResponder(resp *http.Response) (result SecurityConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the network manager security user configurations in a network manager, in a paginated format.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// top - an optional query parameter which specifies the maximum number of records to be returned by the
// server.
// skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
// specifies a starting point to use for subsequent calls.
func (client SecurityUserConfigurationsClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result SecurityConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityUserConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.sclr.Response.Response != nil {
				sc = result.sclr.Response.Response.StatusCode
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
		return result, validation.NewError("network.SecurityUserConfigurationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkManagerName, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.sclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sclr.hasNextLink() && result.sclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityUserConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityUserConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityUserConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityUserConfigurationsClient) ListResponder(resp *http.Response) (result SecurityConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SecurityUserConfigurationsClient) listNextResults(ctx context.Context, lastResults SecurityConfigurationListResult) (result SecurityConfigurationListResult, err error) {
	req, err := lastResults.securityConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityUserConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityUserConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result SecurityConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityUserConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkManagerName, top, skipToken)
	return
}

package guestconfiguration

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

// AssignmentReportsVMSSClient is the guest Configuration Client
type AssignmentReportsVMSSClient struct {
	BaseClient
}

// NewAssignmentReportsVMSSClient creates an instance of the AssignmentReportsVMSSClient client.
func NewAssignmentReportsVMSSClient(subscriptionID string) AssignmentReportsVMSSClient {
	return NewAssignmentReportsVMSSClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssignmentReportsVMSSClientWithBaseURI creates an instance of the AssignmentReportsVMSSClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAssignmentReportsVMSSClientWithBaseURI(baseURI string, subscriptionID string) AssignmentReportsVMSSClient {
	return AssignmentReportsVMSSClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a report for the VMSS guest configuration assignment, by reportId.
// Parameters:
// resourceGroupName - the resource group name.
// vmssName - the name of the virtual machine scale set.
// name - the guest configuration assignment name.
// ID - the GUID for the guest configuration assignment report.
func (client AssignmentReportsVMSSClient) Get(ctx context.Context, resourceGroupName string, vmssName string, name string, ID string) (result AssignmentReportType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentReportsVMSSClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("guestconfiguration.AssignmentReportsVMSSClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, vmssName, name, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentReportsVMSSClient) GetPreparer(ctx context.Context, resourceGroupName string, vmssName string, name string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmssName":          autorest.Encode("path", vmssName),
	}

	const APIVersion = "2022-01-25"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}/reports/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentReportsVMSSClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentReportsVMSSClient) GetResponder(resp *http.Response) (result AssignmentReportType, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all reports for the VMSS guest configuration assignment, latest report first.
// Parameters:
// resourceGroupName - the resource group name.
// vmssName - the name of the virtual machine scale set.
// name - the guest configuration assignment name.
func (client AssignmentReportsVMSSClient) List(ctx context.Context, resourceGroupName string, vmssName string, name string) (result AssignmentReportList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentReportsVMSSClient.List")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("guestconfiguration.AssignmentReportsVMSSClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, vmssName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "guestconfiguration.AssignmentReportsVMSSClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentReportsVMSSClient) ListPreparer(ctx context.Context, resourceGroupName string, vmssName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmssName":          autorest.Encode("path", vmssName),
	}

	const APIVersion = "2022-01-25"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentReportsVMSSClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentReportsVMSSClient) ListResponder(resp *http.Response) (result AssignmentReportList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

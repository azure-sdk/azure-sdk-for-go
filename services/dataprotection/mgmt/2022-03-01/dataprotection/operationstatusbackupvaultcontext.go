package dataprotection

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

// OperationStatusBackupVaultContextClient is the open API 2.0 Specs for Azure Data Protection service
type OperationStatusBackupVaultContextClient struct {
	BaseClient
}

// NewOperationStatusBackupVaultContextClient creates an instance of the OperationStatusBackupVaultContextClient
// client.
func NewOperationStatusBackupVaultContextClient(subscriptionID string) OperationStatusBackupVaultContextClient {
	return NewOperationStatusBackupVaultContextClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationStatusBackupVaultContextClientWithBaseURI creates an instance of the
// OperationStatusBackupVaultContextClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOperationStatusBackupVaultContextClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusBackupVaultContextClient {
	return OperationStatusBackupVaultContextClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
// vaultName - the name of the backup vault.
func (client OperationStatusBackupVaultContextClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result OperationResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationStatusBackupVaultContextClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vaultName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.OperationStatusBackupVaultContextClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.OperationStatusBackupVaultContextClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.OperationStatusBackupVaultContextClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationStatusBackupVaultContextClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/operationStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationStatusBackupVaultContextClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationStatusBackupVaultContextClient) GetResponder(resp *http.Response) (result OperationResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

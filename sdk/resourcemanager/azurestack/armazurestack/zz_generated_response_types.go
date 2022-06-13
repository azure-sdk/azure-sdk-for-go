//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

// CloudManifestFileClientGetResponse contains the response from method CloudManifestFileClient.Get.
type CloudManifestFileClientGetResponse struct {
	CloudManifestFileResponse
}

// CloudManifestFileClientListResponse contains the response from method CloudManifestFileClient.List.
type CloudManifestFileClientListResponse struct {
	CloudManifestFileResponse
}

// CustomerSubscriptionsClientCreateResponse contains the response from method CustomerSubscriptionsClient.Create.
type CustomerSubscriptionsClientCreateResponse struct {
	CustomerSubscription
}

// CustomerSubscriptionsClientDeleteResponse contains the response from method CustomerSubscriptionsClient.Delete.
type CustomerSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// CustomerSubscriptionsClientGetResponse contains the response from method CustomerSubscriptionsClient.Get.
type CustomerSubscriptionsClientGetResponse struct {
	CustomerSubscription
}

// CustomerSubscriptionsClientListResponse contains the response from method CustomerSubscriptionsClient.List.
type CustomerSubscriptionsClientListResponse struct {
	CustomerSubscriptionList
}

// LinkedSubscriptionsClientCreateOrUpdateResponse contains the response from method LinkedSubscriptionsClient.CreateOrUpdate.
type LinkedSubscriptionsClientCreateOrUpdateResponse struct {
	LinkedSubscription
}

// LinkedSubscriptionsClientDeleteResponse contains the response from method LinkedSubscriptionsClient.Delete.
type LinkedSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// LinkedSubscriptionsClientGetResponse contains the response from method LinkedSubscriptionsClient.Get.
type LinkedSubscriptionsClientGetResponse struct {
	LinkedSubscription
}

// LinkedSubscriptionsClientListByResourceGroupResponse contains the response from method LinkedSubscriptionsClient.ListByResourceGroup.
type LinkedSubscriptionsClientListByResourceGroupResponse struct {
	LinkedSubscriptionsList
}

// LinkedSubscriptionsClientListBySubscriptionResponse contains the response from method LinkedSubscriptionsClient.ListBySubscription.
type LinkedSubscriptionsClientListBySubscriptionResponse struct {
	LinkedSubscriptionsList
}

// LinkedSubscriptionsClientUpdateResponse contains the response from method LinkedSubscriptionsClient.Update.
type LinkedSubscriptionsClientUpdateResponse struct {
	LinkedSubscription
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationList
}

// ProductsClientGetProductResponse contains the response from method ProductsClient.GetProduct.
type ProductsClientGetProductResponse struct {
	Product
}

// ProductsClientGetProductsResponse contains the response from method ProductsClient.GetProducts.
type ProductsClientGetProductsResponse struct {
	ProductList
}

// ProductsClientGetResponse contains the response from method ProductsClient.Get.
type ProductsClientGetResponse struct {
	Product
}

// ProductsClientListDetailsResponse contains the response from method ProductsClient.ListDetails.
type ProductsClientListDetailsResponse struct {
	ExtendedProduct
}

// ProductsClientListResponse contains the response from method ProductsClient.List.
type ProductsClientListResponse struct {
	ProductList
}

// ProductsClientUploadLogResponse contains the response from method ProductsClient.UploadLog.
type ProductsClientUploadLogResponse struct {
	ProductLog
}

// RegistrationsClientCreateOrUpdateResponse contains the response from method RegistrationsClient.CreateOrUpdate.
type RegistrationsClientCreateOrUpdateResponse struct {
	Registration
}

// RegistrationsClientDeleteResponse contains the response from method RegistrationsClient.Delete.
type RegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// RegistrationsClientEnableRemoteManagementResponse contains the response from method RegistrationsClient.EnableRemoteManagement.
type RegistrationsClientEnableRemoteManagementResponse struct {
	// placeholder for future response values
}

// RegistrationsClientGetActivationKeyResponse contains the response from method RegistrationsClient.GetActivationKey.
type RegistrationsClientGetActivationKeyResponse struct {
	ActivationKeyResult
}

// RegistrationsClientGetResponse contains the response from method RegistrationsClient.Get.
type RegistrationsClientGetResponse struct {
	Registration
}

// RegistrationsClientListBySubscriptionResponse contains the response from method RegistrationsClient.ListBySubscription.
type RegistrationsClientListBySubscriptionResponse struct {
	RegistrationList
}

// RegistrationsClientListResponse contains the response from method RegistrationsClient.List.
type RegistrationsClientListResponse struct {
	RegistrationList
}

// RegistrationsClientUpdateResponse contains the response from method RegistrationsClient.Update.
type RegistrationsClientUpdateResponse struct {
	Registration
}

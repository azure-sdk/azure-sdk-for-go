//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armedgeorder

// AddressesClientCreateResponse contains the response from method AddressesClient.BeginCreate.
type AddressesClientCreateResponse struct {
	AddressResource
}

// AddressesClientDeleteResponse contains the response from method AddressesClient.BeginDelete.
type AddressesClientDeleteResponse struct {
	// placeholder for future response values
}

// AddressesClientGetResponse contains the response from method AddressesClient.Get.
type AddressesClientGetResponse struct {
	AddressResource
}

// AddressesClientListByResourceGroupResponse contains the response from method AddressesClient.NewListByResourceGroupPager.
type AddressesClientListByResourceGroupResponse struct {
	AddressResourceList
}

// AddressesClientListBySubscriptionResponse contains the response from method AddressesClient.NewListBySubscriptionPager.
type AddressesClientListBySubscriptionResponse struct {
	AddressResourceList
}

// AddressesClientUpdateResponse contains the response from method AddressesClient.BeginUpdate.
type AddressesClientUpdateResponse struct {
	AddressResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// OrderItemsClientCancelResponse contains the response from method OrderItemsClient.Cancel.
type OrderItemsClientCancelResponse struct {
	// placeholder for future response values
}

// OrderItemsClientCreateResponse contains the response from method OrderItemsClient.BeginCreate.
type OrderItemsClientCreateResponse struct {
	OrderItemResource
}

// OrderItemsClientDeleteResponse contains the response from method OrderItemsClient.BeginDelete.
type OrderItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// OrderItemsClientGetResponse contains the response from method OrderItemsClient.Get.
type OrderItemsClientGetResponse struct {
	OrderItemResource
}

// OrderItemsClientListByResourceGroupResponse contains the response from method OrderItemsClient.NewListByResourceGroupPager.
type OrderItemsClientListByResourceGroupResponse struct {
	OrderItemResourceList
}

// OrderItemsClientListBySubscriptionResponse contains the response from method OrderItemsClient.NewListBySubscriptionPager.
type OrderItemsClientListBySubscriptionResponse struct {
	OrderItemResourceList
}

// OrderItemsClientReturnResponse contains the response from method OrderItemsClient.BeginReturn.
type OrderItemsClientReturnResponse struct {
	// placeholder for future response values
}

// OrderItemsClientUpdateResponse contains the response from method OrderItemsClient.BeginUpdate.
type OrderItemsClientUpdateResponse struct {
	OrderItemResource
}

// OrdersClientGetResponse contains the response from method OrdersClient.Get.
type OrdersClientGetResponse struct {
	OrderResource
}

// OrdersClientListByResourceGroupResponse contains the response from method OrdersClient.NewListByResourceGroupPager.
type OrdersClientListByResourceGroupResponse struct {
	OrderResourceList
}

// OrdersClientListBySubscriptionResponse contains the response from method OrdersClient.NewListBySubscriptionPager.
type OrdersClientListBySubscriptionResponse struct {
	OrderResourceList
}

// ProductsAndConfigurationsClientListConfigurationsResponse contains the response from method ProductsAndConfigurationsClient.NewListConfigurationsPager.
type ProductsAndConfigurationsClientListConfigurationsResponse struct {
	Configurations
}

// ProductsAndConfigurationsClientListProductFamiliesMetadataResponse contains the response from method ProductsAndConfigurationsClient.NewListProductFamiliesMetadataPager.
type ProductsAndConfigurationsClientListProductFamiliesMetadataResponse struct {
	ProductFamiliesMetadata
}

// ProductsAndConfigurationsClientListProductFamiliesResponse contains the response from method ProductsAndConfigurationsClient.NewListProductFamiliesPager.
type ProductsAndConfigurationsClientListProductFamiliesResponse struct {
	ProductFamilies
}

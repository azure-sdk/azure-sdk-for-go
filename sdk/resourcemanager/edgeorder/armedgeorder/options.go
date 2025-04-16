// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

// AddressesClientBeginCreateOptions contains the optional parameters for the AddressesClient.BeginCreate method.
type AddressesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AddressesClientBeginDeleteOptions contains the optional parameters for the AddressesClient.BeginDelete method.
type AddressesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AddressesClientBeginUpdateOptions contains the optional parameters for the AddressesClient.BeginUpdate method.
type AddressesClientBeginUpdateOptions struct {
	// Defines the If-Match condition. The patch will be performed only if the ETag of the job on the server matches this value.
	IfMatch *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AddressesClientGetOptions contains the optional parameters for the AddressesClient.Get method.
type AddressesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AddressesClientListByResourceGroupOptions contains the optional parameters for the AddressesClient.NewListByResourceGroupPager
// method.
type AddressesClientListByResourceGroupOptions struct {
	// $filter is supported to filter based on shipping address properties. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of addresses, which provides the next page in the list of addresses.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// AddressesClientListBySubscriptionOptions contains the optional parameters for the AddressesClient.NewListBySubscriptionPager
// method.
type AddressesClientListBySubscriptionOptions struct {
	// $filter is supported to filter based on shipping address properties. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of addresses, which provides the next page in the list of addresses.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrderItemsClientBeginCreateOptions contains the optional parameters for the OrderItemsClient.BeginCreate method.
type OrderItemsClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OrderItemsClientBeginDeleteOptions contains the optional parameters for the OrderItemsClient.BeginDelete method.
type OrderItemsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OrderItemsClientBeginReturnOptions contains the optional parameters for the OrderItemsClient.BeginReturn method.
type OrderItemsClientBeginReturnOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OrderItemsClientBeginUpdateOptions contains the optional parameters for the OrderItemsClient.BeginUpdate method.
type OrderItemsClientBeginUpdateOptions struct {
	// Defines the If-Match condition. The patch will be performed only if the ETag of the order on the server matches this value.
	IfMatch *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OrderItemsClientCancelOptions contains the optional parameters for the OrderItemsClient.Cancel method.
type OrderItemsClientCancelOptions struct {
	// placeholder for future optional parameters
}

// OrderItemsClientGetOptions contains the optional parameters for the OrderItemsClient.Get method.
type OrderItemsClientGetOptions struct {
	// $expand is supported on parent device details, device details, forward shipping details and reverse shipping details parameters.
	// Each of these can be provided as a comma separated list. Parent Device
	// Details for order item provides details on the devices of the product, Device Details for order item provides details on
	// the devices of the child configurations of the product, Forward and Reverse
	// Shipping details provide forward and reverse shipping details respectively.
	Expand *string
}

// OrderItemsClientListByResourceGroupOptions contains the optional parameters for the OrderItemsClient.NewListByResourceGroupPager
// method.
type OrderItemsClientListByResourceGroupOptions struct {
	// $expand is supported on parent device details, device details, forward shipping details and reverse shipping details parameters.
	// Each of these can be provided as a comma separated list. Parent Device
	// Details for order item provides details on the devices of the product, Device Details for order item provides details on
	// the devices of the child configurations of the product, Forward and Reverse
	// Shipping details provide forward and reverse shipping details respectively.
	Expand *string

	// $filter is supported to filter based on order id and order Item Type. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of order items, which provides the next page in the list of order items.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// OrderItemsClientListBySubscriptionOptions contains the optional parameters for the OrderItemsClient.NewListBySubscriptionPager
// method.
type OrderItemsClientListBySubscriptionOptions struct {
	// $expand is supported on parent device details, device details, forward shipping details and reverse shipping details parameters.
	// Each of these can be provided as a comma separated list. Parent Device
	// Details for order item provides details on the devices of the product, Device Details for order item provides details on
	// the devices of the child configurations of the product, Forward and Reverse
	// Shipping details provide forward and reverse shipping details respectively.
	Expand *string

	// $filter is supported to filter based on order id and order Item Type. Filter supports only equals operation.
	Filter *string

	// $skipToken is supported on Get list of order items, which provides the next page in the list of order items.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// OrdersClientGetOptions contains the optional parameters for the OrdersClient.Get method.
type OrdersClientGetOptions struct {
	// placeholder for future optional parameters
}

// OrdersClientListByResourceGroupOptions contains the optional parameters for the OrdersClient.NewListByResourceGroupPager
// method.
type OrdersClientListByResourceGroupOptions struct {
	// $skipToken is supported on Get list of orders, which provides the next page in the list of orders.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// OrdersClientListBySubscriptionOptions contains the optional parameters for the OrdersClient.NewListBySubscriptionPager
// method.
type OrdersClientListBySubscriptionOptions struct {
	// $skipToken is supported on Get list of orders, which provides the next page in the list of orders.
	SkipToken *string

	// $top is supported on fetching list of resources. $top=10 means that the first 10 items in the list will be returned to
	// the API caller.
	Top *int32
}

// ProductsAndConfigurationsClientListConfigurationsOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListConfigurationsPager
// method.
type ProductsAndConfigurationsClientListConfigurationsOptions struct {
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken *string
}

// ProductsAndConfigurationsClientListProductFamiliesMetadataOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListProductFamiliesMetadataPager
// method.
type ProductsAndConfigurationsClientListProductFamiliesMetadataOptions struct {
	// $skipToken is supported on list of product families metadata, which provides the next page in the list of product families
	// metadata.
	SkipToken *string
}

// ProductsAndConfigurationsClientListProductFamiliesOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListProductFamiliesPager
// method.
type ProductsAndConfigurationsClientListProductFamiliesOptions struct {
	// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
	Expand *string

	// $skipToken is supported on list of product families, which provides the next page in the list of product families.
	SkipToken *string
}

# Release History

## 2.0.0-beta.1 (2023-07-11)
### Breaking Changes

- Function `*ClientFactory.NewManagementClient` has been removed
- Function `NewManagementClient` has been removed
- Function `*ManagementClient.CancelOrderItem` has been removed
- Function `*ManagementClient.BeginCreateAddress` has been removed
- Function `*ManagementClient.BeginCreateOrderItem` has been removed
- Function `*ManagementClient.BeginDeleteAddressByName` has been removed
- Function `*ManagementClient.BeginDeleteOrderItemByName` has been removed
- Function `*ManagementClient.GetAddressByName` has been removed
- Function `*ManagementClient.GetOrderByName` has been removed
- Function `*ManagementClient.GetOrderItemByName` has been removed
- Function `*ManagementClient.NewListAddressesAtResourceGroupLevelPager` has been removed
- Function `*ManagementClient.NewListAddressesAtSubscriptionLevelPager` has been removed
- Function `*ManagementClient.NewListConfigurationsPager` has been removed
- Function `*ManagementClient.NewListOperationsPager` has been removed
- Function `*ManagementClient.NewListOrderAtResourceGroupLevelPager` has been removed
- Function `*ManagementClient.NewListOrderAtSubscriptionLevelPager` has been removed
- Function `*ManagementClient.NewListOrderItemsAtResourceGroupLevelPager` has been removed
- Function `*ManagementClient.NewListOrderItemsAtSubscriptionLevelPager` has been removed
- Function `*ManagementClient.NewListProductFamiliesMetadataPager` has been removed
- Function `*ManagementClient.NewListProductFamiliesPager` has been removed
- Function `*ManagementClient.BeginReturnOrderItem` has been removed
- Function `*ManagementClient.BeginUpdateAddress` has been removed
- Function `*ManagementClient.BeginUpdateOrderItem` has been removed
- Struct `ConfigurationFilters` has been removed
- Struct `ShippingDetails` has been removed
- Field `ConfigurationFilters` of struct `ConfigurationsRequest` has been removed
- Field `ManagementRpDetails` of struct `OrderItemDetails` has been removed
- Field `Count`, `DeviceDetails` of struct `ProductDetails` has been removed

### Features Added

- New value `AvailabilityStageDiscoverable` added to enum type `AvailabilityStage`
- New value `LinkTypeDiscoverable` added to enum type `LinkType`
- New enum type `ChildConfigurationType` with values `ChildConfigurationTypeAdditionalConfiguration`, `ChildConfigurationTypeDeviceConfiguration`
- New enum type `FulfillmentType` with values `FulfillmentTypeExternal`, `FulfillmentTypeMicrosoft`
- New enum type `IdentificationType` with values `IdentificationTypeNotSupported`, `IdentificationTypeSerialNumber`
- New enum type `OrderMode` with values `OrderModeDefault`, `OrderModeDoNotFulfill`
- New function `NewAddressesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AddressesClient, error)`
- New function `*AddressesClient.BeginCreate(context.Context, string, string, AddressResource, *AddressesClientBeginCreateOptions) (*runtime.Poller[AddressesClientCreateResponse], error)`
- New function `*AddressesClient.BeginDelete(context.Context, string, string, *AddressesClientBeginDeleteOptions) (*runtime.Poller[AddressesClientDeleteResponse], error)`
- New function `*AddressesClient.Get(context.Context, string, string, *AddressesClientGetOptions) (AddressesClientGetResponse, error)`
- New function `*AddressesClient.NewListByResourceGroupPager(string, *AddressesClientListByResourceGroupOptions) *runtime.Pager[AddressesClientListByResourceGroupResponse]`
- New function `*AddressesClient.NewListBySubscriptionPager(*AddressesClientListBySubscriptionOptions) *runtime.Pager[AddressesClientListBySubscriptionResponse]`
- New function `*AddressesClient.BeginUpdate(context.Context, string, string, AddressUpdateParameter, *AddressesClientBeginUpdateOptions) (*runtime.Poller[AddressesClientUpdateResponse], error)`
- New function `*ClientFactory.NewAddressesClient() *AddressesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewOrderItemsClient() *OrderItemsClient`
- New function `*ClientFactory.NewOrdersClient() *OrdersClient`
- New function `*ClientFactory.NewProductsAndConfigurationsClient() *ProductsAndConfigurationsClient`
- New function `NewOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationsClient, error)`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `NewOrderItemsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OrderItemsClient, error)`
- New function `*OrderItemsClient.Cancel(context.Context, string, string, CancellationReason, *OrderItemsClientCancelOptions) (OrderItemsClientCancelResponse, error)`
- New function `*OrderItemsClient.BeginCreate(context.Context, string, string, OrderItemResource, *OrderItemsClientBeginCreateOptions) (*runtime.Poller[OrderItemsClientCreateResponse], error)`
- New function `*OrderItemsClient.BeginDelete(context.Context, string, string, *OrderItemsClientBeginDeleteOptions) (*runtime.Poller[OrderItemsClientDeleteResponse], error)`
- New function `*OrderItemsClient.Get(context.Context, string, string, *OrderItemsClientGetOptions) (OrderItemsClientGetResponse, error)`
- New function `*OrderItemsClient.NewListByResourceGroupPager(string, *OrderItemsClientListByResourceGroupOptions) *runtime.Pager[OrderItemsClientListByResourceGroupResponse]`
- New function `*OrderItemsClient.NewListBySubscriptionPager(*OrderItemsClientListBySubscriptionOptions) *runtime.Pager[OrderItemsClientListBySubscriptionResponse]`
- New function `*OrderItemsClient.BeginReturn(context.Context, string, string, ReturnOrderItemDetails, *OrderItemsClientBeginReturnOptions) (*runtime.Poller[OrderItemsClientReturnResponse], error)`
- New function `*OrderItemsClient.BeginUpdate(context.Context, string, string, OrderItemUpdateParameter, *OrderItemsClientBeginUpdateOptions) (*runtime.Poller[OrderItemsClientUpdateResponse], error)`
- New function `NewOrdersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OrdersClient, error)`
- New function `*OrdersClient.Get(context.Context, string, string, string, *OrdersClientGetOptions) (OrdersClientGetResponse, error)`
- New function `*OrdersClient.NewListByResourceGroupPager(string, *OrdersClientListByResourceGroupOptions) *runtime.Pager[OrdersClientListByResourceGroupResponse]`
- New function `*OrdersClient.NewListBySubscriptionPager(*OrdersClientListBySubscriptionOptions) *runtime.Pager[OrdersClientListBySubscriptionResponse]`
- New function `NewProductsAndConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ProductsAndConfigurationsClient, error)`
- New function `*ProductsAndConfigurationsClient.NewListConfigurationsPager(ConfigurationsRequest, *ProductsAndConfigurationsClientListConfigurationsOptions) *runtime.Pager[ProductsAndConfigurationsClientListConfigurationsResponse]`
- New function `*ProductsAndConfigurationsClient.NewListProductFamiliesMetadataPager(*ProductsAndConfigurationsClientListProductFamiliesMetadataOptions) *runtime.Pager[ProductsAndConfigurationsClientListProductFamiliesMetadataResponse]`
- New function `*ProductsAndConfigurationsClient.NewListProductFamiliesPager(ProductFamiliesRequest, *ProductsAndConfigurationsClientListProductFamiliesOptions) *runtime.Pager[ProductsAndConfigurationsClientListProductFamiliesResponse]`
- New struct `AdditionalConfiguration`
- New struct `CategoryInformation`
- New struct `ChildConfiguration`
- New struct `ChildConfigurationFilter`
- New struct `ChildConfigurationProperties`
- New struct `ConfigurationDeviceDetails`
- New struct `ConfigurationFilter`
- New struct `GroupedChildConfigurations`
- New struct `ResourceMoveRequest`
- New struct `UserAssignedIdentity`
- New field `FulfilledBy` in struct `BasicInformation`
- New field `FulfilledBy` in struct `CommonProperties`
- New field `ChildConfigurationTypes`, `FulfilledBy`, `GroupedChildConfigurations` in struct `ConfigurationProperties`
- New field `ConfigurationFilter` in struct `ConfigurationsRequest`
- New field `OrderItemMode` in struct `OrderItemDetails`
- New field `Identity` in struct `OrderItemResource`
- New field `Identity` in struct `OrderItemUpdateParameter`
- New field `OrderMode` in struct `OrderProperties`
- New field `ChildConfigurationDeviceDetails`, `IdentificationType`, `OptInAdditionalConfigurations`, `ParentDeviceDetails` in struct `ProductDetails`
- New field `FulfilledBy` in struct `ProductFamilyProperties`
- New field `FulfilledBy` in struct `ProductLineProperties`
- New field `FulfilledBy` in struct `ProductProperties`
- New field `UserAssignedIdentities` in struct `ResourceIdentity`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
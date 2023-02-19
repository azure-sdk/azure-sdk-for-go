# Release History

## 2.0.0-beta.1 (2023-02-19)
### Breaking Changes

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
- Struct `ManagementClient` has been removed
- Struct `ResourceIdentity` has been removed
- Struct `ShippingDetails` has been removed
- Field `ConfigurationFilters` of struct `ConfigurationsRequest` has been removed
- Field `ManagementRpDetails` of struct `OrderItemDetails` has been removed
- Field `Count` of struct `ProductDetails` has been removed
- Field `DeviceDetails` of struct `ProductDetails` has been removed

### Features Added

- New value `AvailabilityStageDiscoverable` added to type alias `AvailabilityStage`
- New value `LinkTypeDiscoverable` added to type alias `LinkType`
- New type alias `ChildConfigurationType` with values `ChildConfigurationTypeAdditionalConfiguration`, `ChildConfigurationTypeDeviceConfiguration`
- New type alias `FulfillmentType` with values `FulfillmentTypeExternal`, `FulfillmentTypeMicrosoft`
- New type alias `IdentificationType` with values `IdentificationTypeNotSupported`, `IdentificationTypeSerialNumber`
- New type alias `OrderMode` with values `OrderModeDefault`, `OrderModeDoNotFulfill`
- New function `NewAddressesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AddressesClient, error)`
- New function `*AddressesClient.BeginCreate(context.Context, string, string, AddressResource, *AddressesClientBeginCreateOptions) (*runtime.Poller[AddressesClientCreateResponse], error)`
- New function `*AddressesClient.BeginDelete(context.Context, string, string, *AddressesClientBeginDeleteOptions) (*runtime.Poller[AddressesClientDeleteResponse], error)`
- New function `*AddressesClient.Get(context.Context, string, string, *AddressesClientGetOptions) (AddressesClientGetResponse, error)`
- New function `*AddressesClient.NewListByResourceGroupPager(string, *AddressesClientListByResourceGroupOptions) *runtime.Pager[AddressesClientListByResourceGroupResponse]`
- New function `*AddressesClient.NewListBySubscriptionPager(*AddressesClientListBySubscriptionOptions) *runtime.Pager[AddressesClientListBySubscriptionResponse]`
- New function `*AddressesClient.BeginUpdate(context.Context, string, string, AddressUpdateParameter, *AddressesClientBeginUpdateOptions) (*runtime.Poller[AddressesClientUpdateResponse], error)`
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
- New struct `AddressesClient`
- New struct `CategoryInformation`
- New struct `ChildConfiguration`
- New struct `ChildConfigurationFilter`
- New struct `ChildConfigurationProperties`
- New struct `ConfigurationDeviceDetails`
- New struct `ConfigurationFilter`
- New struct `GroupedChildConfigurations`
- New struct `OperationsClient`
- New struct `OrderItemsClient`
- New struct `OrdersClient`
- New struct `ProductsAndConfigurationsClient`
- New struct `ResourceMoveRequest`
- New field `FulfilledBy` in struct `BasicInformation`
- New field `FulfilledBy` in struct `CommonProperties`
- New field `ChildConfigurationTypes` in struct `ConfigurationProperties`
- New field `FulfilledBy` in struct `ConfigurationProperties`
- New field `GroupedChildConfigurations` in struct `ConfigurationProperties`
- New field `ConfigurationFilter` in struct `ConfigurationsRequest`
- New field `OrderItemMode` in struct `OrderItemDetails`
- New field `OrderMode` in struct `OrderProperties`
- New field `ChildConfigurationDeviceDetails` in struct `ProductDetails`
- New field `IdentificationType` in struct `ProductDetails`
- New field `OptInAdditionalConfigurations` in struct `ProductDetails`
- New field `ParentDeviceDetails` in struct `ProductDetails`
- New field `FulfilledBy` in struct `ProductFamilyProperties`
- New field `FulfilledBy` in struct `ProductLineProperties`
- New field `FulfilledBy` in struct `ProductProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
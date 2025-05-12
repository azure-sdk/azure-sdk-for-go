# Release History

## 2.0.0 (2025-05-12)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `*ErrorAdditionalInfoInfo`
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
- Struct `BasicInformation` has been removed
- Struct `CommonProperties` has been removed
- Struct `ConfigurationFilters` has been removed
- Struct `ErrorResponse` has been removed
- Struct `ProxyResource` has been removed
- Struct `Resource` has been removed
- Struct `ShippingDetails` has been removed
- Struct `TrackedResource` has been removed
- Field `ConfigurationFilters` of struct `ConfigurationsRequest` has been removed
- Field `ManagementRpDetails` of struct `OrderItemDetails` has been removed
- Field `Count`, `DeviceDetails` of struct `ProductDetails` has been removed

### Features Added

- New value `AvailabilityStageDiscoverable` added to enum type `AvailabilityStage`
- New value `LinkTypeDiscoverable` added to enum type `LinkType`
- New value `OrderItemTypeExternal` added to enum type `OrderItemType`
- New value `StageNameReadyToSetup` added to enum type `StageName`
- New enum type `AddressClassification` with values `AddressClassificationShipping`, `AddressClassificationSite`
- New enum type `AutoProvisioningStatus` with values `AutoProvisioningStatusDisabled`, `AutoProvisioningStatusEnabled`
- New enum type `ChildConfigurationType` with values `ChildConfigurationTypeAdditionalConfiguration`, `ChildConfigurationTypeDeviceConfiguration`
- New enum type `DevicePresenceVerificationStatus` with values `DevicePresenceVerificationStatusCompleted`, `DevicePresenceVerificationStatusNotInitiated`
- New enum type `FulfillmentType` with values `FulfillmentTypeExternal`, `FulfillmentTypeMicrosoft`
- New enum type `IdentificationType` with values `IdentificationTypeNotSupported`, `IdentificationTypeSerialNumber`
- New enum type `OrderMode` with values `OrderModeDefault`, `OrderModeDoNotFulfill`
- New enum type `ProvisioningState` with values `ProvisioningStateCanceled`, `ProvisioningStateCreating`, `ProvisioningStateFailed`, `ProvisioningStateSucceeded`
- New enum type `ProvisioningSupport` with values `ProvisioningSupportCloudBased`, `ProvisioningSupportManual`
- New enum type `TermCommitmentType` with values `TermCommitmentTypeNone`, `TermCommitmentTypeTimed`, `TermCommitmentTypeTrial`
- New function `NewAddressResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AddressResourcesClient, error)`
- New function `*AddressResourcesClient.BeginCreate(context.Context, string, string, AddressResource, *AddressResourcesClientBeginCreateOptions) (*runtime.Poller[AddressResourcesClientCreateResponse], error)`
- New function `*AddressResourcesClient.BeginDelete(context.Context, string, string, *AddressResourcesClientBeginDeleteOptions) (*runtime.Poller[AddressResourcesClientDeleteResponse], error)`
- New function `*AddressResourcesClient.Get(context.Context, string, string, *AddressResourcesClientGetOptions) (AddressResourcesClientGetResponse, error)`
- New function `*AddressResourcesClient.NewListByResourceGroupPager(string, *AddressResourcesClientListByResourceGroupOptions) *runtime.Pager[AddressResourcesClientListByResourceGroupResponse]`
- New function `*AddressResourcesClient.NewListBySubscriptionPager(*AddressResourcesClientListBySubscriptionOptions) *runtime.Pager[AddressResourcesClientListBySubscriptionResponse]`
- New function `*AddressResourcesClient.BeginUpdate(context.Context, string, string, AddressUpdateParameter, *AddressResourcesClientBeginUpdateOptions) (*runtime.Poller[AddressResourcesClientUpdateResponse], error)`
- New function `*ClientFactory.NewAddressResourcesClient() *AddressResourcesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewOrderItemResourcesClient() *OrderItemResourcesClient`
- New function `*ClientFactory.NewOrderResourcesClient() *OrderResourcesClient`
- New function `*ClientFactory.NewOrdersOperationGroupClient() *OrdersOperationGroupClient`
- New function `*ClientFactory.NewProductsAndConfigurationsOperationGroupClient() *ProductsAndConfigurationsOperationGroupClient`
- New function `NewOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationsClient, error)`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `NewOrderItemResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OrderItemResourcesClient, error)`
- New function `*OrderItemResourcesClient.Cancel(context.Context, string, string, CancellationReason, *OrderItemResourcesClientCancelOptions) (OrderItemResourcesClientCancelResponse, error)`
- New function `*OrderItemResourcesClient.BeginCreate(context.Context, string, string, OrderItemResource, *OrderItemResourcesClientBeginCreateOptions) (*runtime.Poller[OrderItemResourcesClientCreateResponse], error)`
- New function `*OrderItemResourcesClient.BeginDelete(context.Context, string, string, *OrderItemResourcesClientBeginDeleteOptions) (*runtime.Poller[OrderItemResourcesClientDeleteResponse], error)`
- New function `*OrderItemResourcesClient.Get(context.Context, string, string, *OrderItemResourcesClientGetOptions) (OrderItemResourcesClientGetResponse, error)`
- New function `*OrderItemResourcesClient.NewListByResourceGroupPager(string, *OrderItemResourcesClientListByResourceGroupOptions) *runtime.Pager[OrderItemResourcesClientListByResourceGroupResponse]`
- New function `*OrderItemResourcesClient.NewListBySubscriptionPager(*OrderItemResourcesClientListBySubscriptionOptions) *runtime.Pager[OrderItemResourcesClientListBySubscriptionResponse]`
- New function `*OrderItemResourcesClient.BeginReturnOrderItem(context.Context, string, string, ReturnOrderItemDetails, *OrderItemResourcesClientBeginReturnOrderItemOptions) (*runtime.Poller[OrderItemResourcesClientReturnOrderItemResponse], error)`
- New function `*OrderItemResourcesClient.BeginUpdate(context.Context, string, string, OrderItemUpdateParameter, *OrderItemResourcesClientBeginUpdateOptions) (*runtime.Poller[OrderItemResourcesClientUpdateResponse], error)`
- New function `NewOrderResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OrderResourcesClient, error)`
- New function `*OrderResourcesClient.Get(context.Context, string, string, string, *OrderResourcesClientGetOptions) (OrderResourcesClientGetResponse, error)`
- New function `NewOrdersOperationGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OrdersOperationGroupClient, error)`
- New function `*OrdersOperationGroupClient.NewListByResourceGroupPager(string, *OrdersOperationGroupClientListByResourceGroupOptions) *runtime.Pager[OrdersOperationGroupClientListByResourceGroupResponse]`
- New function `*OrdersOperationGroupClient.NewListBySubscriptionPager(*OrdersOperationGroupClientListBySubscriptionOptions) *runtime.Pager[OrdersOperationGroupClientListBySubscriptionResponse]`
- New function `NewProductsAndConfigurationsOperationGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ProductsAndConfigurationsOperationGroupClient, error)`
- New function `*ProductsAndConfigurationsOperationGroupClient.NewListConfigurationsPager(ConfigurationsRequest, *ProductsAndConfigurationsOperationGroupClientListConfigurationsOptions) *runtime.Pager[ProductsAndConfigurationsOperationGroupClientListConfigurationsResponse]`
- New function `*ProductsAndConfigurationsOperationGroupClient.NewListProductFamiliesMetadataPager(*ProductsAndConfigurationsOperationGroupClientListProductFamiliesMetadataOptions) *runtime.Pager[ProductsAndConfigurationsOperationGroupClientListProductFamiliesMetadataResponse]`
- New function `*ProductsAndConfigurationsOperationGroupClient.NewListProductFamiliesPager(ProductFamiliesRequest, *ProductsAndConfigurationsOperationGroupClientListProductFamiliesOptions) *runtime.Pager[ProductsAndConfigurationsOperationGroupClientListProductFamiliesResponse]`
- New struct `AdditionalConfiguration`
- New struct `CategoryInformation`
- New struct `ChildConfiguration`
- New struct `ChildConfigurationFilter`
- New struct `ChildConfigurationProperties`
- New struct `ConfigurationDeviceDetails`
- New struct `ConfigurationFilter`
- New struct `DevicePresenceVerificationDetails`
- New struct `ErrorAdditionalInfoInfo`
- New struct `GroupedChildConfigurations`
- New struct `OkResponse`
- New struct `OrderItemDetailsUpdateParameter`
- New struct `ProductDetailsUpdateParameter`
- New struct `ProvisioningDetails`
- New struct `SiteDetails`
- New struct `TermCommitmentInformation`
- New struct `TermCommitmentPreferences`
- New struct `TermTypeDetails`
- New struct `UserAssignedIdentity`
- New field `AddressClassification`, `ProvisioningState` in struct `AddressProperties`
- New field `TermTypeDetails` in struct `BillingMeterDetails`
- New field `ChildConfigurationTypes`, `FulfilledBy`, `GroupedChildConfigurations`, `ProvisioningSupport`, `SupportedTermCommitmentDurations` in struct `ConfigurationProperties`
- New field `ConfigurationFilter` in struct `ConfigurationsRequest`
- New field `DisplaySerialNumber`, `ProvisioningDetails`, `ProvisioningSupport` in struct `DeviceDetails`
- New field `ConfigurationIDDisplayName` in struct `HierarchyInformation`
- New field `OrderItemMode`, `SiteDetails` in struct `OrderItemDetails`
- New field `ProvisioningState` in struct `OrderItemProperties`
- New field `Identity` in struct `OrderItemResource`
- New field `Identity` in struct `OrderItemUpdateParameter`
- New field `OrderItemDetails` in struct `OrderItemUpdateProperties`
- New field `OrderMode` in struct `OrderProperties`
- New field `TermCommitmentPreferences` in struct `Preferences`
- New field `ChildConfigurationDeviceDetails`, `IdentificationType`, `OptInAdditionalConfigurations`, `ParentDeviceDetails`, `ParentProvisioningDetails`, `TermCommitmentInformation` in struct `ProductDetails`
- New field `FulfilledBy` in struct `ProductFamilyProperties`
- New field `FulfilledBy` in struct `ProductLineProperties`
- New field `FulfilledBy` in struct `ProductProperties`
- New field `UserAssignedIdentities` in struct `ResourceIdentity`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


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
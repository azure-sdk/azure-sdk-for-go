# Release History

## 2.0.0 (2022-06-18)
### Breaking Changes

- Type of `CertificateOrderProperties.AppServiceCertificateNotRenewableReasons` has been changed from `[]*AppServiceCertificateOrderPropertiesAppServiceCertificateNotRenewableReasonsItem` to `[]*ResourceNotRenewableReason`
- Type of `CertificateOrderPatchResourceProperties.AppServiceCertificateNotRenewableReasons` has been changed from `[]*AppServiceCertificateOrderPatchResourcePropertiesAppServiceCertificateNotRenewableReasonsItem` to `[]*ResourceNotRenewableReason`
- Type of `ProviderClientGetAvailableStacksOnPremOptions.OSTypeSelected` has been changed from `*Enum20` to `*Enum19`
- Type of `ProviderClientGetAvailableStacksOptions.OSTypeSelected` has been changed from `*Enum15` to `*Enum14`
- Type of `DomainPatchResourceProperties.DomainNotRenewableReasons` has been changed from `[]*DomainPatchResourcePropertiesDomainNotRenewableReasonsItem` to `[]*ResourceNotRenewableReason`
- Type of `ProviderClientGetWebAppStacksForLocationOptions.StackOsType` has been changed from `*Enum18` to `*Enum17`
- Type of `ProviderClientGetFunctionAppStacksForLocationOptions.StackOsType` has been changed from `*Enum17` to `*Enum16`
- Type of `DomainProperties.DomainNotRenewableReasons` has been changed from `[]*DomainPropertiesDomainNotRenewableReasonsItem` to `[]*ResourceNotRenewableReason`
- Type of `ProviderClientGetWebAppStacksOptions.StackOsType` has been changed from `*Enum19` to `*Enum18`
- Type of `ProviderClientGetFunctionAppStacksOptions.StackOsType` has been changed from `*Enum16` to `*Enum15`
- Const `Enum20WindowsFunctions` has been removed
- Const `DomainPropertiesDomainNotRenewableReasonsItemRegistrationStatusNotSupportedForRenewal` has been removed
- Const `AppServiceCertificateOrderPropertiesAppServiceCertificateNotRenewableReasonsItemSubscriptionNotActive` has been removed
- Const `DomainPropertiesDomainNotRenewableReasonsItemSubscriptionNotActive` has been removed
- Const `DomainPropertiesDomainNotRenewableReasonsItemExpirationNotInRenewalTimeRange` has been removed
- Const `AppServiceCertificateOrderPropertiesAppServiceCertificateNotRenewableReasonsItemExpirationNotInRenewalTimeRange` has been removed
- Const `Enum15WindowsFunctions` has been removed
- Const `AppServiceCertificateOrderPatchResourcePropertiesAppServiceCertificateNotRenewableReasonsItemSubscriptionNotActive` has been removed
- Const `DomainPatchResourcePropertiesDomainNotRenewableReasonsItemSubscriptionNotActive` has been removed
- Const `AppServiceCertificateOrderPatchResourcePropertiesAppServiceCertificateNotRenewableReasonsItemExpirationNotInRenewalTimeRange` has been removed
- Const `Enum15LinuxFunctions` has been removed
- Const `Enum20All` has been removed
- Const `AppServiceCertificateOrderPropertiesAppServiceCertificateNotRenewableReasonsItemRegistrationStatusNotSupportedForRenewal` has been removed
- Const `Enum20LinuxFunctions` has been removed
- Const `DomainPatchResourcePropertiesDomainNotRenewableReasonsItemExpirationNotInRenewalTimeRange` has been removed
- Const `Enum20Linux` has been removed
- Const `AppServiceCertificateOrderPatchResourcePropertiesAppServiceCertificateNotRenewableReasonsItemRegistrationStatusNotSupportedForRenewal` has been removed
- Const `DomainPatchResourcePropertiesDomainNotRenewableReasonsItemRegistrationStatusNotSupportedForRenewal` has been removed
- Const `Enum20Windows` has been removed
- Function `PossibleDomainPatchResourcePropertiesDomainNotRenewableReasonsItemValues` has been removed
- Function `PossibleAppServiceCertificateOrderPropertiesAppServiceCertificateNotRenewableReasonsItemValues` has been removed
- Function `PossibleDomainPropertiesDomainNotRenewableReasonsItemValues` has been removed
- Function `PossibleEnum20Values` has been removed
- Function `PossibleAppServiceCertificateOrderPatchResourcePropertiesAppServiceCertificateNotRenewableReasonsItemValues` has been removed
- Struct `CertificateEmailProperties` has been removed
- Struct `CertificateOrderActionProperties` has been removed
- Field `Kind` of struct `CertificateEmail` has been removed
- Field `Properties` of struct `CertificateEmail` has been removed
- Field `ID` of struct `CertificateEmail` has been removed
- Field `Name` of struct `CertificateEmail` has been removed
- Field `Type` of struct `CertificateEmail` has been removed
- Field `Properties` of struct `CertificateOrderAction` has been removed
- Field `ID` of struct `CertificateOrderAction` has been removed
- Field `Name` of struct `CertificateOrderAction` has been removed
- Field `Type` of struct `CertificateOrderAction` has been removed
- Field `Kind` of struct `CertificateOrderAction` has been removed

### Features Added

- New field `LinkedBackends` in struct `StaticSiteBuildARMResourceProperties`
- New field `NumberOfWorkers` in struct `PlanProperties`
- New field `VnetContentShareEnabled` in struct `SiteProperties`
- New field `PublicNetworkAccess` in struct `SiteProperties`
- New field `VnetImagePullEnabled` in struct `SiteProperties`
- New field `VnetRouteAllEnabled` in struct `SiteProperties`
- New field `StorageAccountRequired` in struct `TriggeredWebJobProperties`
- New field `PublicNetworkAccess` in struct `TriggeredWebJobProperties`
- New field `NumberOfWorkers` in struct `PlanPatchResourceProperties`
- New field `Hostname` in struct `WebSiteManagementClientListCustomHostNameSitesOptions`
- New field `FtpEnabled` in struct `AseV3NetworkingConfigurationProperties`
- New field `InboundIPAddressOverride` in struct `AseV3NetworkingConfigurationProperties`
- New field `RemoteDebugEnabled` in struct `AseV3NetworkingConfigurationProperties`
- New field `UpgradeAvailability` in struct `Environment`
- New field `CustomDNSSuffixConfiguration` in struct `Environment`
- New field `NetworkingConfiguration` in struct `Environment`
- New field `UpgradePreference` in struct `Environment`
- New field `ActionType` in struct `CertificateOrderAction`
- New field `CreatedAt` in struct `CertificateOrderAction`
- New field `EmailID` in struct `CertificateEmail`
- New field `TimeStamp` in struct `CertificateEmail`
- New field `LinkedBackends` in struct `StaticSite`
- New field `PublicNetworkAccess` in struct `StaticSite`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
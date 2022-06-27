# Release History

## 0.2.0 (2022-06-27)
### Features Added

- New const `StatusWaitingForCloudOperator`
- New const `SSHKeyTypeSSHCustomerUser`
- New const `StatusUpgradeFailed`
- New const `StatusConnecting`
- New const `StatusOffline`
- New const `StatusPreUpgrade`
- New const `StatusUpgradePrerequisitesCompleted`
- New const `ProviderKubeVirt`
- New const `StatusPreparingForUpgrade`
- New const `StatusPostUpgrade`
- New const `StatusNone`
- New const `StatusUpdatingCluster`
- New const `StatusUpdatingCloudOperator`
- New const `StatusUpgradeComplete`
- New const `ProviderOpenStack`
- New const `StatusUpgradeClusterExtensionFailedToDelete`
- New const `AccessProfileTypeClusterCustomerUser`
- New const `StatusUpdatingCAPI`
- New function `*AppliancesClient.GetUpgradeGraph(context.Context, string, string, string, *AppliancesClientGetUpgradeGraphOptions) (AppliancesClientGetUpgradeGraphResponse, error)`
- New function `*AppliancesClient.ListClusterCustomerUserCredential(context.Context, string, string, *AppliancesClientListClusterCustomerUserCredentialOptions) (AppliancesClientListClusterCustomerUserCredentialResponse, error)`
- New function `PossibleSSHKeyTypeValues() []SSHKeyType`
- New struct `ApplianceListClusterCustomerUserCredentialResults`
- New struct `AppliancesClientGetUpgradeGraphOptions`
- New struct `AppliancesClientGetUpgradeGraphResponse`
- New struct `AppliancesClientListClusterCustomerUserCredentialOptions`
- New struct `AppliancesClientListClusterCustomerUserCredentialResponse`
- New struct `SSHKey`
- New struct `SupportedVersion`
- New struct `SupportedVersionCatalogVersion`
- New struct `SupportedVersionCatalogVersionData`
- New struct `SupportedVersionMetadata`
- New struct `UpgradeGraph`
- New struct `UpgradeGraphProperties`


## 0.1.0 (2022-06-10)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
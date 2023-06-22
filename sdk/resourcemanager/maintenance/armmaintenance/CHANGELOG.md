# Release History

## 1.2.0-beta.1 (2023-06-22)
### Features Added

- New value `MaintenanceScopeResource` added to enum type `MaintenanceScope`
- New enum type `RebootOptions` with values `RebootOptionsAlways`, `RebootOptionsIfRequired`, `RebootOptionsNever`
- New enum type `TaskScope` with values `TaskScopeGlobal`, `TaskScopeResource`
- New function `*ClientFactory.NewConfigurationAssignmentsWithinSubscriptionClient() *ConfigurationAssignmentsWithinSubscriptionClient`
- New function `*ConfigurationAssignmentsClient.Get(context.Context, string, string, string, string, string, *ConfigurationAssignmentsClientGetOptions) (ConfigurationAssignmentsClientGetResponse, error)`
- New function `*ConfigurationAssignmentsClient.GetParent(context.Context, string, string, string, string, string, string, string, *ConfigurationAssignmentsClientGetParentOptions) (ConfigurationAssignmentsClientGetParentResponse, error)`
- New function `NewConfigurationAssignmentsWithinSubscriptionClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConfigurationAssignmentsWithinSubscriptionClient, error)`
- New function `*ConfigurationAssignmentsWithinSubscriptionClient.NewListPager(*ConfigurationAssignmentsWithinSubscriptionClientListOptions) *runtime.Pager[ConfigurationAssignmentsWithinSubscriptionClientListResponse]`
- New struct `ConfigurationOverrides`
- New struct `InputLinuxParameters`
- New struct `InputPatchConfiguration`
- New struct `InputWindowsParameters`
- New struct `OverrideProperties`
- New struct `SoftwareUpdateConfigurationTasks`
- New struct `TaskProperties`
- New field `InstallPatches`, `Overrides` in struct `ConfigurationProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-04-07)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
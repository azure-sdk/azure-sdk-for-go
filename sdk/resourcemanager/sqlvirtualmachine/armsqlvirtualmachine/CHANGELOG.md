# Release History

## 0.6.0 (2022-06-06)
### Breaking Changes

- Type of `ResourceIdentity.Type` has been changed from `*IdentityType` to `*IdentityTypeWithNone`
- Type of `AutoBackupSettings.DaysOfWeek` has been changed from `[]*DaysOfWeek` to `[]*AutoBackupDaysOfWeek`
- Type of `Schedule.DayOfWeek` has been changed from `*DayOfWeek` to `*AssessmentDayOfWeek`
- Const `CreatedByTypeApplication` has been removed
- Const `IdentityTypeNone` has been removed
- Const `DaysOfWeekWednesday` has been removed
- Const `CreatedByTypeManagedIdentity` has been removed
- Const `CreatedByTypeUser` has been removed
- Const `DaysOfWeekMonday` has been removed
- Const `DaysOfWeekSaturday` has been removed
- Const `CreatedByTypeKey` has been removed
- Const `DaysOfWeekFriday` has been removed
- Const `DaysOfWeekTuesday` has been removed
- Const `DaysOfWeekThursday` has been removed
- Const `IdentityTypeSystemAssigned` has been removed
- Const `DaysOfWeekSunday` has been removed
- Function `*timeRFC3339.UnmarshalText` has been removed
- Function `timeRFC3339.MarshalText` has been removed
- Function `PossibleIdentityTypeValues` has been removed
- Function `*timeRFC3339.Parse` has been removed
- Function `PossibleDaysOfWeekValues` has been removed
- Function `PossibleCreatedByTypeValues` has been removed
- Struct `SystemData` has been removed
- Field `SystemData` of struct `AvailabilityGroupListener` has been removed
- Field `SystemData` of struct `SQLVirtualMachine` has been removed
- Field `SystemData` of struct `Group` has been removed

### Features Added

- New const `DayOfWeekEveryday`
- New const `AssessmentDayOfWeekTuesday`
- New const `IdentityTypeWithNoneNone`
- New const `IdentityTypeWithNoneSystemAssigned`
- New const `AutoBackupDaysOfWeekWednesday`
- New const `AutoBackupDaysOfWeekThursday`
- New const `AutoBackupDaysOfWeekTuesday`
- New const `AssessmentDayOfWeekWednesday`
- New const `AutoBackupDaysOfWeekMonday`
- New const `AssessmentDayOfWeekThursday`
- New const `AutoBackupDaysOfWeekFriday`
- New const `AssessmentDayOfWeekFriday`
- New const `ClusterSubnetTypeSingleSubnet`
- New const `AssessmentDayOfWeekSunday`
- New const `ClusterSubnetTypeMultiSubnet`
- New const `AssessmentDayOfWeekMonday`
- New const `AssessmentDayOfWeekSaturday`
- New const `AutoBackupDaysOfWeekSaturday`
- New const `AutoBackupDaysOfWeekSunday`
- New function `PossibleIdentityTypeWithNoneValues() []IdentityTypeWithNone`
- New function `PossibleAssessmentDayOfWeekValues() []AssessmentDayOfWeek`
- New function `PossibleAutoBackupDaysOfWeekValues() []AutoBackupDaysOfWeek`
- New function `PossibleClusterSubnetTypeValues() []ClusterSubnetType`
- New struct `MultiSubnetIPConfiguration`
- New field `IsLPIMEnabled` in struct `SQLInstanceSettings`
- New field `IsIFIEnabled` in struct `SQLInstanceSettings`
- New field `PersistFolderPath` in struct `SQLTempDbSettings`
- New field `PersistFolder` in struct `SQLTempDbSettings`
- New field `ClusterSubnetType` in struct `WsfcDomainProfile`
- New field `MultiSubnetIPConfigurations` in struct `AvailabilityGroupListenerProperties`
- New field `WsfcStaticIP` in struct `Properties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
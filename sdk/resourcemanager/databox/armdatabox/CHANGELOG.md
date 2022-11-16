# Release History

## 2.0.0 (2022-11-16)
### Breaking Changes

- Field `AccountName` of struct `DiskGranularCopyLogDetails` has been removed

### Features Added

- New const `ReverseTransportPreferenceEditStatusEnabled`
- New const `ReverseShippingDetailsEditStatusEnabled`
- New const `ReverseShippingDetailsEditStatusNotSupported`
- New const `ReverseShippingDetailsEditStatusDisabled`
- New const `ReverseTransportPreferenceEditStatusNotSupported`
- New const `ReverseTransportPreferenceEditStatusDisabled`
- New type alias `ReverseTransportPreferenceEditStatus`
- New type alias `ReverseShippingDetailsEditStatus`
- New function `PossibleReverseShippingDetailsEditStatusValues() []ReverseShippingDetailsEditStatus`
- New function `PossibleReverseTransportPreferenceEditStatusValues() []ReverseTransportPreferenceEditStatus`
- New struct `ContactInfo`
- New struct `ReverseShippingDetails`
- New field `AccountID` in struct `DiskGranularCopyLogDetails`
- New field `CountriesWithinCommerceBoundary` in struct `SKUProperties`
- New field `ReverseShippingDetails` in struct `JobDetails`
- New field `ReverseTransportPreferences` in struct `Preferences`
- New field `ReverseShippingDetails` in struct `DiskJobDetails`
- New field `GranularCopyLogDetails` in struct `DiskJobDetails`
- New field `ReverseShippingDetails` in struct `HeavyJobDetails`
- New field `ReverseShippingDetails` in struct `CommonJobDetails`
- New field `ReverseShippingDetails` in struct `UpdateJobDetails`
- New field `Preferences` in struct `UpdateJobDetails`
- New field `IsUpdated` in struct `TransportPreferences`
- New field `ReverseTransportPreferenceUpdate` in struct `JobProperties`
- New field `ReverseShippingDetailsUpdate` in struct `JobProperties`
- New field `ReverseShippingDetails` in struct `CustomerDiskJobDetails`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
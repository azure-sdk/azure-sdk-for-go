# Release History

## 2.0.0 (2023-03-30)
### Breaking Changes

- Field `AccountName` of struct `DiskGranularCopyLogDetails` has been removed

### Features Added

- New enum type `HardwareEncryption` with values `HardwareEncryptionDisabled`, `HardwareEncryptionEnabled`
- New enum type `ReverseShippingDetailsEditStatus` with values `ReverseShippingDetailsEditStatusDisabled`, `ReverseShippingDetailsEditStatusEnabled`, `ReverseShippingDetailsEditStatusNotSupported`
- New enum type `ReverseTransportPreferenceEditStatus` with values `ReverseTransportPreferenceEditStatusDisabled`, `ReverseTransportPreferenceEditStatusEnabled`, `ReverseTransportPreferenceEditStatusNotSupported`
- New struct `ContactInfo`
- New struct `ReverseShippingDetails`
- New field `ReverseShippingDetails` in struct `CommonJobDetails`
- New field `Actions` in struct `CopyProgress`
- New field `Error` in struct `CopyProgress`
- New field `Actions` in struct `CustomerDiskCopyProgress`
- New field `Error` in struct `CustomerDiskCopyProgress`
- New field `ReverseShippingDetails` in struct `CustomerDiskJobDetails`
- New field `Actions` in struct `DiskCopyProgress`
- New field `Error` in struct `DiskCopyProgress`
- New field `AccountID` in struct `DiskGranularCopyLogDetails`
- New field `Actions` in struct `DiskGranularCopyProgress`
- New field `Error` in struct `DiskGranularCopyProgress`
- New field `GranularCopyLogDetails` in struct `DiskJobDetails`
- New field `ReverseShippingDetails` in struct `DiskJobDetails`
- New field `HardwareEncryption` in struct `EncryptionPreferences`
- New field `Actions` in struct `GranularCopyProgress`
- New field `Error` in struct `GranularCopyProgress`
- New field `ReverseShippingDetails` in struct `HeavyJobDetails`
- New field `ReverseShippingDetails` in struct `JobDetails`
- New field `ReverseShippingDetailsUpdate` in struct `JobProperties`
- New field `ReverseTransportPreferenceUpdate` in struct `JobProperties`
- New field `SerialNumberCustomerResolutionMap` in struct `MitigateJobRequest`
- New field `ReverseTransportPreferences` in struct `Preferences`
- New field `CountriesWithinCommerceBoundary` in struct `SKUProperties`
- New field `SkipAddressValidation` in struct `ShippingAddress`
- New field `TaxIdentificationNumber` in struct `ShippingAddress`
- New field `IsUpdated` in struct `TransportPreferences`
- New field `Preferences` in struct `UpdateJobDetails`
- New field `ReverseShippingDetails` in struct `UpdateJobDetails`


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
# Release History

## 2.0.0 (2025-04-29)
### Breaking Changes

- Function `*RecordSetsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, RecordType, string, RecordSet, *RecordSetsClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, RecordType, RecordSet, *RecordSetsClientCreateOrUpdateOptions)`
- Function `*RecordSetsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, RecordType, string, *RecordSetsClientDeleteOptions)` to `(context.Context, string, string, string, RecordType, *RecordSetsClientDeleteOptions)`
- Function `*RecordSetsClient.Get` parameter(s) have been changed from `(context.Context, string, string, RecordType, string, *RecordSetsClientGetOptions)` to `(context.Context, string, string, string, RecordType, *RecordSetsClientGetOptions)`
- Function `*RecordSetsClient.Update` parameter(s) have been changed from `(context.Context, string, string, RecordType, string, RecordSet, *RecordSetsClientUpdateOptions)` to `(context.Context, string, string, string, RecordType, RecordSet, *RecordSetsClientUpdateOptions)`

### Features Added

- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `SystemData`
- New field `SystemData` in struct `PrivateZone`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `RecordSet`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`
- New field `SystemData` in struct `VirtualNetworkLink`


## 1.3.0 (2024-09-27)
### Features Added

- New enum type `ResolutionPolicy` with values `ResolutionPolicyDefault`, `ResolutionPolicyNxDomainRedirect`
- New field `ResolutionPolicy` in struct `VirtualNetworkLinkProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
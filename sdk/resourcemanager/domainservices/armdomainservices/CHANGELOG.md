# Release History

## 2.0.0 (2023-02-19)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New type alias `ChannelBinding` with values `ChannelBindingDisabled`, `ChannelBindingEnabled`
- New type alias `LdapSigning` with values `LdapSigningDisabled`, `LdapSigningEnabled`
- New type alias `SyncScope` with values `SyncScopeAll`, `SyncScopeCloudOnly`
- New field `ChannelBinding` in struct `DomainSecuritySettings`
- New field `LdapSigning` in struct `DomainSecuritySettings`
- New field `SyncApplicationID` in struct `DomainServiceProperties`
- New field `SyncScope` in struct `DomainServiceProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
# Release History

## 2.0.0 (2022-07-28)
### Breaking Changes

- Function `*CertificatesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *CertificatesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, Certificate, *CertificatesClientCreateOrUpdateOptions)`
- Struct `CustomHostnameAnalysisResultProperties` has been removed
- Field `Type` of struct `CustomHostnameAnalysisResult` has been removed
- Field `Properties` of struct `CustomHostnameAnalysisResult` has been removed
- Field `ID` of struct `CustomHostnameAnalysisResult` has been removed
- Field `Name` of struct `CustomHostnameAnalysisResult` has been removed
- Field `SystemData` of struct `CustomHostnameAnalysisResult` has been removed
- Field `CertificateEnvelope` of struct `CertificatesClientCreateOrUpdateOptions` has been removed

### Features Added

- New struct `CustomHostnameAnalysisResultCustomDomainVerificationFailureInfo`
- New struct `CustomHostnameAnalysisResultCustomDomainVerificationFailureInfoDetailsItem`
- New field `TxtRecords` in struct `CustomHostnameAnalysisResult`
- New field `AlternateTxtRecords` in struct `CustomHostnameAnalysisResult`
- New field `IsHostnameAlreadyVerified` in struct `CustomHostnameAnalysisResult`
- New field `ARecords` in struct `CustomHostnameAnalysisResult`
- New field `ConflictingContainerAppResourceID` in struct `CustomHostnameAnalysisResult`
- New field `CustomDomainVerificationFailureInfo` in struct `CustomHostnameAnalysisResult`
- New field `HasConflictOnManagedEnvironment` in struct `CustomHostnameAnalysisResult`
- New field `HostName` in struct `CustomHostnameAnalysisResult`
- New field `AlternateCNameRecords` in struct `CustomHostnameAnalysisResult`
- New field `CustomDomainVerificationTest` in struct `CustomHostnameAnalysisResult`
- New field `CNameRecords` in struct `CustomHostnameAnalysisResult`


## 1.0.0 (2022-05-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
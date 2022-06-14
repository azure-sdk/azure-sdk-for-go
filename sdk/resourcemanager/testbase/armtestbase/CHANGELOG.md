# Release History

## 0.6.0 (2022-06-14)
### Features Added

- New field `StderrLogFileName` in struct `ScriptExecutionResult`
- New field `StdoutLogFileName` in struct `ScriptExecutionResult`
- New field `BaselineOSs` in struct `TargetOSInfo`
- New field `PackageTags` in struct `TestSummaryProperties`
- New field `ValidationResultID` in struct `Test`
- New field `InteropMediaVersion` in struct `TestResultProperties`
- New field `InteropMediaType` in struct `TestResultProperties`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
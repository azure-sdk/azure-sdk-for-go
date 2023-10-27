# Release History

## 2.0.0-beta.3 (2023-10-27)
### Breaking Changes

- Field `InnerError` of struct `WorkbookErrorDefinition` has been removed

### Features Added

- New value `WebTestKindStandard` added to enum type `WebTestKind`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `NewOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationsClient, error)`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New struct `ErrorFieldContract`
- New struct `HeaderField`
- New struct `WebTestPropertiesRequest`
- New struct `WebTestPropertiesValidationRules`
- New struct `WebTestPropertiesValidationRulesContentValidation`
- New field `Details` in struct `ErrorResponse`
- New field `Request`, `ValidationRules` in struct `WebTestProperties`
- New field `Innererror` in struct `WorkbookErrorDefinition`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 1.1.0 (2023-04-06)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-06-02)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
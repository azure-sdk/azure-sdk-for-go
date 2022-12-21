# Release History

## 1.1.0-beta.2 (2022-12-21)
### Features Added

- New type alias `AggregateErrors` with values `AggregateErrorsFalse`, `AggregateErrorsTrue`
- New type alias `MedtechMappingValidationCategory` with values `MedtechMappingValidationCategoryFHIRTRANSFORMATION`, `MedtechMappingValidationCategoryNORMALIZATION`
- New type alias `MedtechMappingValidationErrorLevel` with values `MedtechMappingValidationErrorLevelERROR`, `MedtechMappingValidationErrorLevelWARN`
- New function `*ServicesClient.ValidateMedtechMappings(context.Context, ValidateMedtechMappingsParameters, *ServicesClientValidateMedtechMappingsOptions) (ServicesClientValidateMedtechMappingsResponse, error)`
- New struct `CorsConfiguration`
- New struct `MedtechMappingValidationError`
- New struct `MedtechMappingValidationLineInfo`
- New struct `MedtechMeasurement`
- New struct `MedtechMeasurementProperty`
- New struct `ValidateMedtechMappingsDeviceResult`
- New struct `ValidateMedtechMappingsParameters`
- New struct `ValidateMedtechMappingsResult`
- New struct `ValidateMedtechMappingsResultTemplateResult`
- New field `CorsConfiguration` in struct `DicomServiceProperties`
- New field `EnableRegionalMdmAccount` in struct `MetricSpecification`
- New field `IsInternal` in struct `MetricSpecification`
- New field `MetricFilterPattern` in struct `MetricSpecification`
- New field `ResourceIDDimensionNameOverride` in struct `MetricSpecification`
- New field `SourceMdmAccount` in struct `MetricSpecification`


## 1.1.0-beta.1 (2022-05-19)
### Features Added

- New struct `FhirServiceImportConfiguration`
- New struct `ServiceImportConfigurationInfo`
- New field `ImportConfiguration` in struct `ServicesProperties`
- New field `ImportConfiguration` in struct `FhirServiceProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
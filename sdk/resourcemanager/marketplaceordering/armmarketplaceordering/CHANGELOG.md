# Release History

## 2.0.0 (2023-03-08)
### Breaking Changes

- Field `AgreementTerms` of struct `MarketplaceAgreementsClientCancelResponse` has been removed
- Field `AgreementTerms` of struct `MarketplaceAgreementsClientGetAgreementResponse` has been removed
- Field `AgreementTermsArray` of struct `MarketplaceAgreementsClientListResponse` has been removed
- Field `AgreementTerms` of struct `MarketplaceAgreementsClientSignResponse` has been removed

### Features Added

- New type alias `State` with values `StateActive`, `StateCanceled`
- New struct `OldAgreementProperties`
- New struct `OldAgreementTerms`
- New struct `OldAgreementTermsList`
- New anonymous field `OldAgreementTerms` in struct `MarketplaceAgreementsClientCancelResponse`
- New anonymous field `OldAgreementTerms` in struct `MarketplaceAgreementsClientGetAgreementResponse`
- New anonymous field `OldAgreementTermsList` in struct `MarketplaceAgreementsClientListResponse`
- New anonymous field `OldAgreementTerms` in struct `MarketplaceAgreementsClientSignResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
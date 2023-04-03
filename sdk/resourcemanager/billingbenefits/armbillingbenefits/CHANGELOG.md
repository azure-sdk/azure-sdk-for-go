# Release History

## 2.0.0 (2023-04-03)
### Breaking Changes

- Function `NewSavingsPlanClient` parameter(s) have been changed from `(*string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`
- Function `NewSavingsPlanOrderClient` parameter(s) have been changed from `(*string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`

### Features Added

- New function `NewClientFactory(azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewRPClient() *RPClient`
- New function `*ClientFactory.NewReservationOrderAliasClient() *ReservationOrderAliasClient`
- New function `*ClientFactory.NewSavingsPlanClient() *SavingsPlanClient`
- New function `*ClientFactory.NewSavingsPlanOrderAliasClient() *SavingsPlanOrderAliasClient`
- New function `*ClientFactory.NewSavingsPlanOrderClient() *SavingsPlanOrderClient`
- New struct `ClientFactory`
- New field `Expand` in struct `SavingsPlanClientGetOptions`
- New field `Expand` in struct `SavingsPlanOrderClientGetOptions`


## 1.0.0 (2022-12-23)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
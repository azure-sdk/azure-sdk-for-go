# Release History

## 0.6.0 (2022-05-31)
### Features Added

- New const `EmailChannelAuthMethodGraph`
- New const `EmailChannelAuthMethodPassword`
- New function `*OutlookChannel.GetChannel() *Channel`
- New function `PossibleEmailChannelAuthMethodValues() []EmailChannelAuthMethod`
- New struct `CreateEmailSignInURLResponse`
- New struct `CreateEmailSignInURLResponseProperties`
- New struct `EmailClientCreateSignInURLOptions`
- New struct `EmailClientCreateSignInURLResponse`
- New struct `OutlookChannel`
- New field `AuthMethod` in struct `EmailChannelProperties`
- New field `MagicCode` in struct `EmailChannelProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
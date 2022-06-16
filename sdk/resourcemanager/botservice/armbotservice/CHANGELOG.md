# Release History

## 0.6.0 (2022-06-16)
### Breaking Changes

- Field `CallingWebHook` of struct `MsTeamsChannelProperties` has been removed

### Features Added

- New field `TenantID` in struct `BotProperties`
- New field `IsEndpointParametersEnabled` in struct `DirectLineSite`
- New field `IsTokenEnabled` in struct `DirectLineSite`
- New field `IsWebchatPreviewEnabled` in struct `DirectLineSite`
- New field `ETag` in struct `DirectLineSite`
- New field `IsNoStorageEnabled` in struct `DirectLineSite`
- New field `AppID` in struct `DirectLineSite`
- New field `IsDetailedLoggingEnabled` in struct `DirectLineSite`
- New field `CallingWebhook` in struct `MsTeamsChannelProperties`
- New field `CognitiveServiceResourceID` in struct `DirectLineSpeechChannelProperties`
- New field `IsEndpointParametersEnabled` in struct `Site`
- New field `IsDetailedLoggingEnabled` in struct `Site`
- New field `AppID` in struct `Site`
- New field `IsNoStorageEnabled` in struct `Site`
- New field `AuthMethod` in struct `EmailChannelProperties`
- New field `TrustedOrigins` in struct `WebChatSite`
- New field `IsNoStorageEnabled` in struct `WebChatSite`
- New field `IsEndpointParametersEnabled` in struct `WebChatSite`
- New field `AppID` in struct `WebChatSite`
- New field `IsBlockUserUploadEnabled` in struct `WebChatSite`
- New field `IsV1Enabled` in struct `WebChatSite`
- New field `IsSecureSiteEnabled` in struct `WebChatSite`
- New field `IsTokenEnabled` in struct `WebChatSite`
- New field `IsV3Enabled` in struct `WebChatSite`
- New field `ETag` in struct `WebChatSite`
- New field `IsDetailedLoggingEnabled` in struct `WebChatSite`
- New field `ExtensionKey1` in struct `DirectLineChannelProperties`
- New field `ExtensionKey2` in struct `DirectLineChannelProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. BotProperties.IsIsolated
1. WebChatSite.EnablePreview

## Additive Changes

### New Constants

1. MsaAppType.MsaAppTypeMultiTenant
1. MsaAppType.MsaAppTypeSingleTenant
1. MsaAppType.MsaAppTypeUserAssignedMSI
1. PublicNetworkAccess.PublicNetworkAccessDisabled
1. PublicNetworkAccess.PublicNetworkAccessEnabled

### New Funcs

1. PossibleMsaAppTypeValues() []MsaAppType
1. PossiblePublicNetworkAccessValues() []PublicNetworkAccess

### Struct Changes

#### New Structs

1. ServiceProviderParameterMetadata
1. ServiceProviderParameterMetadataConstraints

#### New Struct Fields

1. AlexaChannel.Etag
1. AlexaChannel.ProvisioningState
1. Bot.Zones
1. BotChannel.Zones
1. BotProperties.AllSettings
1. BotProperties.AppPasswordHint
1. BotProperties.CmekEncryptionStatus
1. BotProperties.DisableLocalAuth
1. BotProperties.IsDeveloperAppInsightsAPIKeySet
1. BotProperties.IsStreamingSupported
1. BotProperties.ManifestURL
1. BotProperties.MigrationToken
1. BotProperties.MsaAppMSIResourceID
1. BotProperties.MsaAppTenantID
1. BotProperties.MsaAppType
1. BotProperties.OpenWithHint
1. BotProperties.Parameters
1. BotProperties.ProvisioningState
1. BotProperties.PublicNetworkAccess
1. BotProperties.PublishingCredentials
1. Channel.Etag
1. ConnectionSetting.Zones
1. ConnectionSettingProperties.ProvisioningState
1. DirectLineChannel.Etag
1. DirectLineChannel.Location
1. DirectLineChannel.ProvisioningState
1. DirectLineChannelProperties.DirectLineEmbedCode
1. DirectLineSite.IsBlockUserUploadEnabled
1. DirectLineSpeechChannel.Etag
1. EmailChannel.Etag
1. FacebookChannel.Etag
1. KikChannel.Etag
1. LineChannel.Etag
1. MsTeamsChannel.Etag
1. MsTeamsChannel.Location
1. MsTeamsChannel.ProvisioningState
1. MsTeamsChannelProperties.AcceptedTerms
1. MsTeamsChannelProperties.DeploymentEnvironment
1. MsTeamsChannelProperties.IncomingCallRoute
1. Resource.Zones
1. ServiceProviderParameter.Metadata
1. SkypeChannel.Etag
1. SkypeChannelProperties.IncomingCallRoute
1. SlackChannel.Etag
1. SmsChannel.Etag
1. TelegramChannel.Etag
1. WebChatChannel.Etag
1. WebChatChannel.Location
1. WebChatChannel.ProvisioningState
1. WebChatSite.IsWebchatPreviewEnabled

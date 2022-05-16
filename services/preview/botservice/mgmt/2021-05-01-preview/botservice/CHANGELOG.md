# Unreleased

## Breaking Changes

### Removed Constants

1. ChannelName.ChannelNameOutlookChannel

### Removed Funcs

1. *ListChannelWithKeysResponse.UnmarshalJSON([]byte) error
1. ListChannelWithKeysResponse.MarshalJSON() ([]byte, error)
1. Site.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. ChannelSettings
1. ListChannelWithKeysResponse
1. Site

#### Removed Struct Fields

1. AlexaChannel.Location
1. BotProperties.StorageResourceID
1. Channel.Location
1. Channel.ProvisioningState
1. ConnectionSettingProperties.ID
1. ConnectionSettingProperties.Name
1. DirectLineSpeechChannel.Location
1. EmailChannel.Location
1. KikChannel.Location
1. KikChannel.ProvisioningState
1. LineChannel.Location
1. LineChannel.ProvisioningState
1. SkypeChannel.Location
1. SkypeChannel.ProvisioningState
1. SlackChannel.ProvisioningState
1. SmsChannel.Location
1. SmsChannel.ProvisioningState
1. TelegramChannel.Location

### Signature Changes

#### Funcs

1. ChannelsClient.ListWithKeys
	- Returns
		- From: ListChannelWithKeysResponse, error
		- To: BotChannel, error
1. ChannelsClient.ListWithKeysResponder
	- Returns
		- From: ListChannelWithKeysResponse, error
		- To: BotChannel, error

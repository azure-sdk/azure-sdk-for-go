# Unreleased

## Additive Changes

### New Constants

1. ChannelNameBasicChannel.ChannelNameBasicChannelChannelNameOutlookChannel

### New Funcs

1. AlexaChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. Channel.AsOutlookChannel() (*OutlookChannel, bool)
1. CreateEmailSignInURLResponse.MarshalJSON() ([]byte, error)
1. DirectLineChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. DirectLineSpeechChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. EmailChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. EmailClient.CreateSignInURL(context.Context, string, string) (CreateEmailSignInURLResponse, error)
1. EmailClient.CreateSignInURLPreparer(context.Context, string, string) (*http.Request, error)
1. EmailClient.CreateSignInURLResponder(*http.Response) (CreateEmailSignInURLResponse, error)
1. EmailClient.CreateSignInURLSender(*http.Request) (*http.Response, error)
1. FacebookChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. KikChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. LineChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. MsTeamsChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. NewEmailClient(string) EmailClient
1. NewEmailClientWithBaseURI(string, string) EmailClient
1. OutlookChannel.AsAlexaChannel() (*AlexaChannel, bool)
1. OutlookChannel.AsBasicChannel() (BasicChannel, bool)
1. OutlookChannel.AsChannel() (*Channel, bool)
1. OutlookChannel.AsDirectLineChannel() (*DirectLineChannel, bool)
1. OutlookChannel.AsDirectLineSpeechChannel() (*DirectLineSpeechChannel, bool)
1. OutlookChannel.AsEmailChannel() (*EmailChannel, bool)
1. OutlookChannel.AsFacebookChannel() (*FacebookChannel, bool)
1. OutlookChannel.AsKikChannel() (*KikChannel, bool)
1. OutlookChannel.AsLineChannel() (*LineChannel, bool)
1. OutlookChannel.AsMsTeamsChannel() (*MsTeamsChannel, bool)
1. OutlookChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. OutlookChannel.AsSkypeChannel() (*SkypeChannel, bool)
1. OutlookChannel.AsSlackChannel() (*SlackChannel, bool)
1. OutlookChannel.AsSmsChannel() (*SmsChannel, bool)
1. OutlookChannel.AsTelegramChannel() (*TelegramChannel, bool)
1. OutlookChannel.AsWebChatChannel() (*WebChatChannel, bool)
1. OutlookChannel.MarshalJSON() ([]byte, error)
1. SkypeChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. SlackChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. SmsChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. TelegramChannel.AsOutlookChannel() (*OutlookChannel, bool)
1. WebChatChannel.AsOutlookChannel() (*OutlookChannel, bool)

### Struct Changes

#### New Structs

1. CreateEmailSignInURLResponse
1. CreateEmailSignInURLResponseProperties
1. EmailClient
1. OutlookChannel

#### New Struct Fields

1. EmailChannelProperties.AuthMethod
1. EmailChannelProperties.MagicCode

# Release History

## 1.1.0 (2022-06-08)
### Features Added

- New const `JobErrorCodeIdentityUnsupported`
- New const `AV1VideoProfileMain`
- New const `EncoderNamedPresetDDGoodQualityAudio`
- New const `JobErrorCategoryAccount`
- New const `EncoderNamedPresetAV1AdaptiveStreaming`
- New const `AV1ComplexityBalanced`
- New const `AV1VideoProfileAuto`
- New const `EncoderNamedPresetAV1SingleBitrate4K`
- New const `EncoderNamedPresetAV1SingleBitrate1080P`
- New const `EncoderNamedPresetAV1SingleBitrate720P`
- New const `AV1ComplexityQuality`
- New const `EncoderNamedPresetAV1ContentAwareEncoding`
- New const `AV1ComplexitySpeed`
- New function `PossibleAV1VideoProfileValues() []AV1VideoProfile`
- New function `*DDAudio.GetAudio() *Audio`
- New function `*AV1Video.GetVideo() *Video`
- New function `*DDAudio.GetCodec() *Codec`
- New function `PossibleAV1ComplexityValues() []AV1Complexity`
- New function `*AV1Video.GetCodec() *Codec`
- New struct `AV1Layer`
- New struct `AV1Video`
- New struct `AV1VideoLayer`
- New struct `DDAudio`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
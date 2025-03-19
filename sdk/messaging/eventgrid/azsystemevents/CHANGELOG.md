# Release History

## 1.0.0 (2025-03-19)
### Breaking Changes

- Type of `ACSMessageDeliveryStatusUpdatedEventData.Error` has been changed from `*Error` to `*InternalACSMessageChannelEventError`
- Type of `ACSMessageReceivedEventData.Error` has been changed from `*Error` to `*InternalACSMessageChannelEventError`
- Type of `ACSRouterJobClassificationFailedEventData.Errors` has been changed from `[]*Error` to `[]InternalACSRouterCommunicationError`

### Features Added

- New struct `ACSChatEventBaseProperties`
- New struct `ACSChatEventInThreadBaseProperties`
- New struct `ACSChatMessageEventBaseProperties`
- New struct `ACSChatMessageEventInThreadBaseProperties`
- New struct `ACSChatThreadEventBaseProperties`
- New struct `ACSChatThreadEventInThreadBaseProperties`
- New struct `ACSMessageEventData`
- New struct `ACSRouterEventData`
- New struct `ACSRouterJobEventData`
- New struct `ACSRouterWorkerEventData`
- New struct `ACSSMSEventBaseProperties`
- New struct `AVSClusterEventData`
- New struct `AVSPrivateCloudEventData`
- New struct `AVSScriptExecutionEventData`
- New struct `AppConfigurationSnapshotEventData`
- New struct `ContainerRegistryArtifactEventData`
- New struct `ContainerRegistryEventData`
- New struct `ContainerServiceClusterSupportEventData`
- New struct `ContainerServiceNodePoolRollingEventData`
- New struct `DeviceConnectionStateEventProperties`
- New struct `DeviceLifeCycleEventProperties`
- New struct `DeviceTelemetryEventProperties`
- New struct `EventGridMQTTClientEventData`
- New struct `InternalACSMessageChannelEventError`
- New struct `InternalACSRouterCommunicationError`
- New struct `MapsGeofenceEventProperties`
- New struct `ResourceNotificationsResourceDeletedEventData`
- New struct `ResourceNotificationsResourceUpdatedEventData`


## 0.6.2 (Unreleased)

### Features Added

### Breaking Changes

### Bugs Fixed

### Other Changes

## 0.6.1 (2025-02-20)

### Features Added

The following fields have been added:
- ACSMessageReceivedEventData.MessageType
- ACSMessageReceivedEventData.MessageID
- ACSMessageReceivedEventData.Reaction

## 0.6.0 (2025-01-21)

### Features Added

The following fields have been added:
- ACSEmailDeliveryReportReceivedEventData.InternetMessageID
- ACSEmailDeliveryReportStatusDetails.RecipientMailServerHostName
- ACSSmsReceivedEventData.SegmentCount

### Breaking Changes

- Azure Media services has been retired - all related system events have been removed. See the [retirement announcement](https://azure.microsoft.com/updates?id=retirement-notice-azure-media-services-is-being-retired-on-30-june-2024) for more details.
- The following types have been renamed to correct incorrect casing: ACSSMSDeliveryAttemptProperties, ACSSMSDeliveryReportReceivedEventData, ACSSMSReceivedEventData, TypeACSSMSDeliveryReportReceived and TypeACSSMSReceived.

## 0.5.0 (2024-11-19)

### Features Added

- A new system event has been added: `ResourceNotificationsContainerServiceEventResourcesScheduledEventData`
- A new field: `StorageLifecyclePolicyCompletedEventData.PolicyRunSummary`

### Breaking Changes

- Models, that were not system events, have been removed. Any fields from those types are incorporated into their corresponding system event type.

### Bugs Fixed

- ACSMessageDeliveryStatusUpdatedEventData.Error has been corrected to use an exported type.

## 0.4.3 (2024-10-14)

### Features Added

- New field has been added to ACSIncomingCallEventData: OnBehalfOfCallee.

## 0.4.2 (2024-09-19)

### Features Added

- A new field has been added to StorageLifecyclePolicyCompletedEventData:
  - TierToColdSummary

## 0.4.1 (2024-08-20)

### Features Added

- New fields have been added:
  - StorageBlobCreatedEventData: AccessTier
  - StorageBlobTierChangedEventData: AccessTier and PreviousTier

### Breaking Changes

- Models that were not system events (ex: ACSChatMessageEventInThreadBaseProperties), or referenced by system events, have been removed.

## 0.4.0 (2024-06-11)

### Breaking Changes

- `Type` has been removed, making it simpler to compare the EventGridEvent.Type and CloudEvent.Type values against
  our provided constants.

- The following models have had 'Advanced' removed from their name:
  - ACSMessageButtonContent
  - ACSMessageContext
  - ACSMessageDeliveryStatusUpdatedEventData
  - ACSMessageEventData
  - ACSMessageInteractiveButtonReplyContent
  - ACSMessageInteractiveContent
  - ACSMessageInteractiveListReplyContent
  - ACSMessageMediaContent
  - ACSMessageReceivedEventData

## 0.3.0 (2024-04-03)

### Features Added

- Added events ACSRouterWorkerUpdatedEventData and ACSAdvancedMessageDeliveryStatusUpdatedEventData. (PR#22638)

### Breaking Changes

Field and type renames:

- Globally, types and fields named ChannelType has been renamed to ChannelKind
- ACS events and constants have been changed to use an all-caps name (ex: AcsEmailDeliveryReportStatusDetails -> ACSEmailDeliveryReportStatusDetails).
- ACSAdvancedMessageContext.ID -> MessageID
- ACSAdvancedMessageReceivedEventData
  - .Media -> MediaContent
  - .Interactive -> InteractiveContent

## 0.2.0 (2024-03-14)

### Features Added

- Added API Center system events under their official names.

### Breaking Changes

- Events have been renamed:
  - APIDefinitionAddedEventData renamed to APICenterAPIDefinitionAddedEventData
  - APIDefinitionUpdatedEventData renamed to APICenterAPIDefinitionUpdatedEventData

## 0.1.0 (2024-03-05)

### Features Added

- Initial preview for Event Grid system events.

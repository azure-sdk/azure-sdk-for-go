# Release History

## 2.0.0 (2025-04-15)
### Breaking Changes

- Function `*ReceiverClient.AcknowledgeEvents` parameter(s) have been changed from `(context.Context, []string, *AcknowledgeEventsOptions)` to `(context.Context, string, string, []string, *ReceiverClientAcknowledgeEventsOptions)`
- Function `*ReceiverClient.AcknowledgeEvents` return value(s) have been changed from `(AcknowledgeEventsResponse, error)` to `(ReceiverClientAcknowledgeEventsResponse, error)`
- Function `*ReceiverClient.ReceiveEvents` parameter(s) have been changed from `(context.Context, *ReceiveEventsOptions)` to `(context.Context, string, string, *ReceiverClientReceiveEventsOptions)`
- Function `*ReceiverClient.ReceiveEvents` return value(s) have been changed from `(ReceiveEventsResponse, error)` to `(ReceiverClientReceiveEventsResponse, error)`
- Function `*ReceiverClient.RejectEvents` parameter(s) have been changed from `(context.Context, []string, *RejectEventsOptions)` to `(context.Context, string, string, []string, *ReceiverClientRejectEventsOptions)`
- Function `*ReceiverClient.RejectEvents` return value(s) have been changed from `(RejectEventsResponse, error)` to `(ReceiverClientRejectEventsResponse, error)`
- Function `*ReceiverClient.ReleaseEvents` parameter(s) have been changed from `(context.Context, []string, *ReleaseEventsOptions)` to `(context.Context, string, string, []string, *ReceiverClientReleaseEventsOptions)`
- Function `*ReceiverClient.ReleaseEvents` return value(s) have been changed from `(ReleaseEventsResponse, error)` to `(ReceiverClientReleaseEventsResponse, error)`
- Function `*ReceiverClient.RenewEventLocks` parameter(s) have been changed from `(context.Context, []string, *RenewEventLocksOptions)` to `(context.Context, string, string, []string, *ReceiverClientRenewEventLocksOptions)`
- Function `*ReceiverClient.RenewEventLocks` return value(s) have been changed from `(RenewEventLocksResponse, error)` to `(ReceiverClientRenewEventLocksResponse, error)`
- Function `*SenderClient.SendEvent` parameter(s) have been changed from `(context.Context, *messaging.CloudEvent, *SendEventOptions)` to `(context.Context, string, CloudEvent, *SenderClientSendEventOptions)`
- Function `*SenderClient.SendEvent` return value(s) have been changed from `(SendEventResponse, error)` to `(SenderClientSendEventResponse, error)`
- Function `*SenderClient.SendEvents` parameter(s) have been changed from `(context.Context, []*messaging.CloudEvent, *SendEventsOptions)` to `(context.Context, string, []CloudEvent, *SenderClientSendEventsOptions)`
- Function `*SenderClient.SendEvents` return value(s) have been changed from `(SendEventsResponse, error)` to `(SenderClientSendEventsResponse, error)`
- Type of `ReceiveDetails.Event` has been changed from `messaging.CloudEvent` to `*CloudEvent`
- Function `NewReceiverClient` has been removed
- Function `NewReceiverClientWithSharedKeyCredential` has been removed
- Function `NewSenderClient` has been removed
- Function `NewSenderClientWithSharedKeyCredential` has been removed
- Struct `AcknowledgeEventsOptions` has been removed
- Struct `AcknowledgeEventsResponse` has been removed
- Struct `AcknowledgeEventsResult` has been removed
- Struct `ReceiveEventsOptions` has been removed
- Struct `ReceiveEventsResponse` has been removed
- Struct `ReceiveEventsResult` has been removed
- Struct `ReceiverClientOptions` has been removed
- Struct `RejectEventsOptions` has been removed
- Struct `RejectEventsResponse` has been removed
- Struct `RejectEventsResult` has been removed
- Struct `ReleaseEventsOptions` has been removed
- Struct `ReleaseEventsResponse` has been removed
- Struct `ReleaseEventsResult` has been removed
- Struct `RenewEventLocksOptions` has been removed
- Struct `RenewEventLocksResponse` has been removed
- Struct `RenewEventLocksResult` has been removed
- Struct `SendEventOptions` has been removed
- Struct `SendEventResponse` has been removed
- Struct `SendEventsOptions` has been removed
- Struct `SendEventsResponse` has been removed
- Struct `SenderClientOptions` has been removed

### Features Added

- New struct `AcknowledgeResult`
- New struct `CloudEvent`
- New struct `InnerError`
- New struct `PublishResult`
- New struct `ReceiveResult`
- New struct `ReceiverClientAcknowledgeEventsOptions`
- New struct `ReceiverClientAcknowledgeEventsResponse`
- New struct `ReceiverClientReceiveEventsOptions`
- New struct `ReceiverClientReceiveEventsResponse`
- New struct `ReceiverClientRejectEventsOptions`
- New struct `ReceiverClientRejectEventsResponse`
- New struct `ReceiverClientReleaseEventsOptions`
- New struct `ReceiverClientReleaseEventsResponse`
- New struct `ReceiverClientRenewEventLocksOptions`
- New struct `ReceiverClientRenewEventLocksResponse`
- New struct `RejectResult`
- New struct `ReleaseResult`
- New struct `RenewLocksResult`
- New struct `SenderClientSendEventOptions`
- New struct `SenderClientSendEventResponse`
- New struct `SenderClientSendEventsOptions`
- New struct `SenderClientSendEventsResponse`
- New field `Details`, `Innererror`, `Target` in struct `Error`


## 1.0.1 (Unreleased)

### Features Added

### Breaking Changes

### Bugs Fixed

### Other Changes

## 1.0.0 (2024-06-11)

### Features Added

- First stable release of the aznamespaces package targeted at API version `2024-06-01`.

### Breaking Changes

- Sending and receiving operations have been moved to separate clients (SenderClient and ReceiverClient).
- Method names have been shortened from <Verb>CloudEvent(s) to <Verb>Event(s)
- LockTokens for AcknowledgeEvents, RejectEvents and ReleaseEvents are now a positional argument, instead of optional.
- Topic and subscription name are now set at the Client level, as part of `NewSenderClient` or `NewReceiverClient`.

## 0.4.1 (2024-03-05)

### Breaking Changes

- This module has been moved from its previous location in `azeventgrid` to this location (`github.com/Azure/azure-sdk-for-go/sdk/messaging/eventgrid/aznamespaces`).

## 0.4.0 (2023-11-27)

### Features Added

- New functionality for Event Grid namespaces: 
  - Client.PublishCloudEvent can be used to publish a single `messaging.CloudEvent`.
  - Client.RenewCloudEventLocks can extend the lock time for a set of events.
  - Client.ReleaseCloudEvents (via ReleaseCloudEventsOptions.ReleaseDelayInSeconds) can release an event with a 
    server-side delay, allowing the message to remain unavailable for a configured period of time.

### Breaking Changes

- FailedLockToken, included in the response for settlement functions, has an `Error` field, which contains the data previously
  in `ErrorDescription` and `ErrorCode`.
- Settlement functions (AcknowledgeCloudEvents, ReleaseCloudEvents, RejectCloudEvents) take lock tokens as a parameter.

## 0.3.0 (2023-10-17)

### Breaking Changes

- Client constructors that take a `key string` parameter for a credential now require an `*azcore.KeyCredential` or `*azcore.SASCredential`.

## 0.2.0 (2023-09-12)

### Features Added

- The publisher client for Event Grid topics has been added as a sub-package under `publisher`.

### Other Changes

- Documentation and examples added for Event Grid namespace client.

## 0.1.0 (2023-07-11)

### Features Added

- Initial preview for the Event Grid package for Event Grid Namespaces

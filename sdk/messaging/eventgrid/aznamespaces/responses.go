// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

// ReceiverClientAcknowledgeEventsResponse contains the response from method ReceiverClient.AcknowledgeEvents.
type ReceiverClientAcknowledgeEventsResponse struct {
	// The result of the Acknowledge operation.
	AcknowledgeResult
}

// ReceiverClientReceiveEventsResponse contains the response from method ReceiverClient.ReceiveEvents.
type ReceiverClientReceiveEventsResponse struct {
	// Details of the Receive operation response.
	ReceiveResult
}

// ReceiverClientRejectEventsResponse contains the response from method ReceiverClient.RejectEvents.
type ReceiverClientRejectEventsResponse struct {
	// The result of the Reject operation.
	RejectResult
}

// ReceiverClientReleaseEventsResponse contains the response from method ReceiverClient.ReleaseEvents.
type ReceiverClientReleaseEventsResponse struct {
	// The result of the Release operation.
	ReleaseResult
}

// ReceiverClientRenewEventLocksResponse contains the response from method ReceiverClient.RenewEventLocks.
type ReceiverClientRenewEventLocksResponse struct {
	// The result of the RenewLock operation.
	RenewLocksResult
}

// SenderClientSendEventResponse contains the response from method SenderClient.SendEvent.
type SenderClientSendEventResponse struct {
	// The result of the Publish operation.
	PublishResult
}

// SenderClientSendEventsResponse contains the response from method SenderClient.SendEvents.
type SenderClientSendEventsResponse struct {
	// The result of the Publish operation.
	PublishResult
}

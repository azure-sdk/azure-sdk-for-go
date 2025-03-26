// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

// ReceiverClientAcknowledgeEventsOptions contains the optional parameters for the ReceiverClient.AcknowledgeEvents method.
type ReceiverClientAcknowledgeEventsOptions struct {
	// placeholder for future optional parameters
}

// ReceiverClientReceiveEventsOptions contains the optional parameters for the ReceiverClient.ReceiveEvents method.
type ReceiverClientReceiveEventsOptions struct {
	// Max Events count to be received. Minimum value is 1, while maximum value is 100 events. If not specified, the default value
	// is 1.
	MaxEvents *int32

	// Max wait time value for receive operation in Seconds. It is the time in seconds that the server approximately waits for
	// the availability of an event and responds to the request. If an event is available, the broker responds immediately to
	// the client. Minimum value is 10 seconds, while maximum value is 120 seconds. If not specified, the default value is 60
	// seconds.
	MaxWaitTime *int32
}

// ReceiverClientRejectEventsOptions contains the optional parameters for the ReceiverClient.RejectEvents method.
type ReceiverClientRejectEventsOptions struct {
	// placeholder for future optional parameters
}

// ReceiverClientReleaseEventsOptions contains the optional parameters for the ReceiverClient.ReleaseEvents method.
type ReceiverClientReleaseEventsOptions struct {
	// Release cloud events with the specified delay in seconds.
	ReleaseDelayInSeconds *ReleaseDelay
}

// ReceiverClientRenewEventLocksOptions contains the optional parameters for the ReceiverClient.RenewEventLocks method.
type ReceiverClientRenewEventLocksOptions struct {
	// placeholder for future optional parameters
}

// SenderClientSendEventOptions contains the optional parameters for the SenderClient.SendEvent method.
type SenderClientSendEventOptions struct {
	// placeholder for future optional parameters
}

// SenderClientSendEventsOptions contains the optional parameters for the SenderClient.SendEvents method.
type SenderClientSendEventsOptions struct {
	// placeholder for future optional parameters
}

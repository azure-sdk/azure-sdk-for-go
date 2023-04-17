//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

const (
	moduleName    = "armmonitor"
	moduleVersion = "v0.10.0"
)

// ReceiverStatus - Indicates the status of the receiver. Receivers that are not Enabled will not receive any communications.
type ReceiverStatus string

const (
	ReceiverStatusNotSpecified ReceiverStatus = "NotSpecified"
	ReceiverStatusEnabled      ReceiverStatus = "Enabled"
	ReceiverStatusDisabled     ReceiverStatus = "Disabled"
)

// PossibleReceiverStatusValues returns the possible values for the ReceiverStatus const type.
func PossibleReceiverStatusValues() []ReceiverStatus {
	return []ReceiverStatus{
		ReceiverStatusNotSpecified,
		ReceiverStatusEnabled,
		ReceiverStatusDisabled,
	}
}

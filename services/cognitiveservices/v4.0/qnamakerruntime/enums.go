package qnamakerruntime

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ErrorCodeType enumerates the values for error code type.
type ErrorCodeType string

const (
	// BadArgument ...
	BadArgument ErrorCodeType = "BadArgument"
	// EndpointKeysError ...
	EndpointKeysError ErrorCodeType = "EndpointKeysError"
	// ExtractionFailure ...
	ExtractionFailure ErrorCodeType = "ExtractionFailure"
	// Forbidden ...
	Forbidden ErrorCodeType = "Forbidden"
	// KbNotFound ...
	KbNotFound ErrorCodeType = "KbNotFound"
	// NotFound ...
	NotFound ErrorCodeType = "NotFound"
	// OperationNotFound ...
	OperationNotFound ErrorCodeType = "OperationNotFound"
	// QnaRuntimeError ...
	QnaRuntimeError ErrorCodeType = "QnaRuntimeError"
	// QuotaExceeded ...
	QuotaExceeded ErrorCodeType = "QuotaExceeded"
	// ServiceError ...
	ServiceError ErrorCodeType = "ServiceError"
	// SKULimitExceeded ...
	SKULimitExceeded ErrorCodeType = "SKULimitExceeded"
	// Unauthorized ...
	Unauthorized ErrorCodeType = "Unauthorized"
	// Unspecified ...
	Unspecified ErrorCodeType = "Unspecified"
	// ValidationFailure ...
	ValidationFailure ErrorCodeType = "ValidationFailure"
)

// PossibleErrorCodeTypeValues returns an array of possible values for the ErrorCodeType const type.
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return []ErrorCodeType{BadArgument, EndpointKeysError, ExtractionFailure, Forbidden, KbNotFound, NotFound, OperationNotFound, QnaRuntimeError, QuotaExceeded, ServiceError, SKULimitExceeded, Unauthorized, Unspecified, ValidationFailure}
}

// StrictFiltersCompoundOperationType enumerates the values for strict filters compound operation type.
type StrictFiltersCompoundOperationType string

const (
	// AND ...
	AND StrictFiltersCompoundOperationType = "AND"
	// OR ...
	OR StrictFiltersCompoundOperationType = "OR"
)

// PossibleStrictFiltersCompoundOperationTypeValues returns an array of possible values for the StrictFiltersCompoundOperationType const type.
func PossibleStrictFiltersCompoundOperationTypeValues() []StrictFiltersCompoundOperationType {
	return []StrictFiltersCompoundOperationType{AND, OR}
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armedgeorder

const (
	moduleName    = "armedgeorder"
	moduleVersion = "v2.0.0-beta.1"
)

// ActionStatusEnum - Describes whether the order item is deletable or not.
type ActionStatusEnum string

const (
	// ActionStatusEnumAllowed - Allowed flag.
	ActionStatusEnumAllowed ActionStatusEnum = "Allowed"
	// ActionStatusEnumNotAllowed - Not Allowed flag.
	ActionStatusEnumNotAllowed ActionStatusEnum = "NotAllowed"
)

// PossibleActionStatusEnumValues returns the possible values for the ActionStatusEnum const type.
func PossibleActionStatusEnumValues() []ActionStatusEnum {
	return []ActionStatusEnum{
		ActionStatusEnumAllowed,
		ActionStatusEnumNotAllowed,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AddressType - Type of address.
type AddressType string

const (
	// AddressTypeCommercial - Commercial Address.
	AddressTypeCommercial AddressType = "Commercial"
	// AddressTypeNone - Address type not known.
	AddressTypeNone AddressType = "None"
	// AddressTypeResidential - Residential Address.
	AddressTypeResidential AddressType = "Residential"
)

// PossibleAddressTypeValues returns the possible values for the AddressType const type.
func PossibleAddressTypeValues() []AddressType {
	return []AddressType{
		AddressTypeCommercial,
		AddressTypeNone,
		AddressTypeResidential,
	}
}

// AddressValidationStatus - Status of address validation.
type AddressValidationStatus string

const (
	// AddressValidationStatusAmbiguous - Address provided is ambiguous, please choose one of the alternate addresses returned.
	AddressValidationStatusAmbiguous AddressValidationStatus = "Ambiguous"
	// AddressValidationStatusInvalid - Address provided is invalid or not supported.
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	// AddressValidationStatusValid - Address provided is valid.
	AddressValidationStatusValid AddressValidationStatus = "Valid"
)

// PossibleAddressValidationStatusValues returns the possible values for the AddressValidationStatus const type.
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return []AddressValidationStatus{
		AddressValidationStatusAmbiguous,
		AddressValidationStatusInvalid,
		AddressValidationStatusValid,
	}
}

// AvailabilityStage - Current availability stage of the product.
type AvailabilityStage string

const (
	// AvailabilityStageAvailable - Product is available.
	AvailabilityStageAvailable AvailabilityStage = "Available"
	// AvailabilityStageComingSoon - Product is coming soon.
	AvailabilityStageComingSoon AvailabilityStage = "ComingSoon"
	// AvailabilityStageDeprecated - Product is deprecated.
	AvailabilityStageDeprecated AvailabilityStage = "Deprecated"
	// AvailabilityStageDiscoverable - Product is not available in our service but can be discovered from other sources.
	AvailabilityStageDiscoverable AvailabilityStage = "Discoverable"
	// AvailabilityStagePreview - Product is in preview.
	AvailabilityStagePreview AvailabilityStage = "Preview"
	// AvailabilityStageSignup - Product is available only on signup.
	AvailabilityStageSignup AvailabilityStage = "Signup"
	// AvailabilityStageUnavailable - Product is not available.
	AvailabilityStageUnavailable AvailabilityStage = "Unavailable"
)

// PossibleAvailabilityStageValues returns the possible values for the AvailabilityStage const type.
func PossibleAvailabilityStageValues() []AvailabilityStage {
	return []AvailabilityStage{
		AvailabilityStageAvailable,
		AvailabilityStageComingSoon,
		AvailabilityStageDeprecated,
		AvailabilityStageDiscoverable,
		AvailabilityStagePreview,
		AvailabilityStageSignup,
		AvailabilityStageUnavailable,
	}
}

// BillingType - Represents billing type.
type BillingType string

const (
	// BillingTypePav2 - PaV2 billing.
	BillingTypePav2 BillingType = "Pav2"
	// BillingTypePurchase - Purchase billing.
	BillingTypePurchase BillingType = "Purchase"
)

// PossibleBillingTypeValues returns the possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{
		BillingTypePav2,
		BillingTypePurchase,
	}
}

// ChargingType - Charging type.
type ChargingType string

const (
	// ChargingTypePerDevice - Per device charging type.
	ChargingTypePerDevice ChargingType = "PerDevice"
	// ChargingTypePerOrder - Per order charging type.
	ChargingTypePerOrder ChargingType = "PerOrder"
)

// PossibleChargingTypeValues returns the possible values for the ChargingType const type.
func PossibleChargingTypeValues() []ChargingType {
	return []ChargingType{
		ChargingTypePerDevice,
		ChargingTypePerOrder,
	}
}

type ChildConfigurationType string

const (
	// ChildConfigurationTypeAdditionalConfiguration - Child configuration is an additional configuration.
	ChildConfigurationTypeAdditionalConfiguration ChildConfigurationType = "AdditionalConfiguration"
	// ChildConfigurationTypeDeviceConfiguration - Child configuration is a device configuration.
	ChildConfigurationTypeDeviceConfiguration ChildConfigurationType = "DeviceConfiguration"
)

// PossibleChildConfigurationTypeValues returns the possible values for the ChildConfigurationType const type.
func PossibleChildConfigurationTypeValues() []ChildConfigurationType {
	return []ChildConfigurationType{
		ChildConfigurationTypeAdditionalConfiguration,
		ChildConfigurationTypeDeviceConfiguration,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DescriptionType - Type of description.
type DescriptionType string

const (
	// DescriptionTypeBase - Base description.
	DescriptionTypeBase DescriptionType = "Base"
)

// PossibleDescriptionTypeValues returns the possible values for the DescriptionType const type.
func PossibleDescriptionTypeValues() []DescriptionType {
	return []DescriptionType{
		DescriptionTypeBase,
	}
}

// DisabledReason - Reason why the product is disabled.
type DisabledReason string

const (
	// DisabledReasonCountry - Not available in the requested country.
	DisabledReasonCountry DisabledReason = "Country"
	// DisabledReasonFeature - Required features are not enabled.
	DisabledReasonFeature DisabledReason = "Feature"
	// DisabledReasonNoSubscriptionInfo - Subscription has not registered to Microsoft.DataBox and Service does not have the subscription
	// notification.
	DisabledReasonNoSubscriptionInfo DisabledReason = "NoSubscriptionInfo"
	// DisabledReasonNone - Not disabled.
	DisabledReasonNone DisabledReason = "None"
	// DisabledReasonNotAvailable - The product is not yet available.
	DisabledReasonNotAvailable DisabledReason = "NotAvailable"
	// DisabledReasonOfferType - Subscription does not have required offer types.
	DisabledReasonOfferType DisabledReason = "OfferType"
	// DisabledReasonOutOfStock - The product is out of stock.
	DisabledReasonOutOfStock DisabledReason = "OutOfStock"
	// DisabledReasonRegion - Not available to push data to the requested Azure region.
	DisabledReasonRegion DisabledReason = "Region"
)

// PossibleDisabledReasonValues returns the possible values for the DisabledReason const type.
func PossibleDisabledReasonValues() []DisabledReason {
	return []DisabledReason{
		DisabledReasonCountry,
		DisabledReasonFeature,
		DisabledReasonNoSubscriptionInfo,
		DisabledReasonNone,
		DisabledReasonNotAvailable,
		DisabledReasonOfferType,
		DisabledReasonOutOfStock,
		DisabledReasonRegion,
	}
}

// DoubleEncryptionStatus - Double encryption status as entered by the customer. It is compulsory to give this parameter if
// the 'Deny' or 'Disabled' policy is configured.
type DoubleEncryptionStatus string

const (
	// DoubleEncryptionStatusDisabled - Double encryption is disabled.
	DoubleEncryptionStatusDisabled DoubleEncryptionStatus = "Disabled"
	// DoubleEncryptionStatusEnabled - Double encryption is enabled.
	DoubleEncryptionStatusEnabled DoubleEncryptionStatus = "Enabled"
)

// PossibleDoubleEncryptionStatusValues returns the possible values for the DoubleEncryptionStatus const type.
func PossibleDoubleEncryptionStatusValues() []DoubleEncryptionStatus {
	return []DoubleEncryptionStatus{
		DoubleEncryptionStatusDisabled,
		DoubleEncryptionStatusEnabled,
	}
}

// FulfillmentType - The entity responsible for fulfillment of the item at the given hierarchy level.
type FulfillmentType string

const (
	// FulfillmentTypeExternal - The fulfillment (the whole journey of the product offering) is handled by external third party
	// entities.
	FulfillmentTypeExternal FulfillmentType = "External"
	// FulfillmentTypeMicrosoft - The fulfillment (the whole journey of the product offering) is handled by microsoft.
	FulfillmentTypeMicrosoft FulfillmentType = "Microsoft"
)

// PossibleFulfillmentTypeValues returns the possible values for the FulfillmentType const type.
func PossibleFulfillmentTypeValues() []FulfillmentType {
	return []FulfillmentType{
		FulfillmentTypeExternal,
		FulfillmentTypeMicrosoft,
	}
}

// IdentificationType - Identification type of the configuration.
type IdentificationType string

const (
	// IdentificationTypeNotSupported - Product does not have any explicit identifier.
	IdentificationTypeNotSupported IdentificationType = "NotSupported"
	// IdentificationTypeSerialNumber - Product is identifiable by serial number.
	IdentificationTypeSerialNumber IdentificationType = "SerialNumber"
)

// PossibleIdentificationTypeValues returns the possible values for the IdentificationType const type.
func PossibleIdentificationTypeValues() []IdentificationType {
	return []IdentificationType{
		IdentificationTypeNotSupported,
		IdentificationTypeSerialNumber,
	}
}

// ImageType - Type of the image.
type ImageType string

const (
	// ImageTypeBulletImage - Bullet image.
	ImageTypeBulletImage ImageType = "BulletImage"
	// ImageTypeGenericImage - Generic image.
	ImageTypeGenericImage ImageType = "GenericImage"
	// ImageTypeMainImage - Main image.
	ImageTypeMainImage ImageType = "MainImage"
)

// PossibleImageTypeValues returns the possible values for the ImageType const type.
func PossibleImageTypeValues() []ImageType {
	return []ImageType{
		ImageTypeBulletImage,
		ImageTypeGenericImage,
		ImageTypeMainImage,
	}
}

// LengthHeightUnit - Unit for the dimensions of length, height and width.
type LengthHeightUnit string

const (
	// LengthHeightUnitCM - Centimeter.
	LengthHeightUnitCM LengthHeightUnit = "CM"
	// LengthHeightUnitIN - Inch, applicable for West US.
	LengthHeightUnitIN LengthHeightUnit = "IN"
)

// PossibleLengthHeightUnitValues returns the possible values for the LengthHeightUnit const type.
func PossibleLengthHeightUnitValues() []LengthHeightUnit {
	return []LengthHeightUnit{
		LengthHeightUnitCM,
		LengthHeightUnitIN,
	}
}

// LinkType - Type of link.
type LinkType string

const (
	// LinkTypeDiscoverable - Link to order the product from another source and not from Azure Edge Hardware Center.
	LinkTypeDiscoverable LinkType = "Discoverable"
	// LinkTypeDocumentation - Link to product documentation.
	LinkTypeDocumentation LinkType = "Documentation"
	// LinkTypeGeneric - Generic link.
	LinkTypeGeneric LinkType = "Generic"
	// LinkTypeKnowMore - Link to know more.
	LinkTypeKnowMore LinkType = "KnowMore"
	// LinkTypeSignUp - Link to sign up for products.
	LinkTypeSignUp LinkType = "SignUp"
	// LinkTypeSpecification - Link to product specification.
	LinkTypeSpecification LinkType = "Specification"
	// LinkTypeTermsAndConditions - Terms and conditions link.
	LinkTypeTermsAndConditions LinkType = "TermsAndConditions"
)

// PossibleLinkTypeValues returns the possible values for the LinkType const type.
func PossibleLinkTypeValues() []LinkType {
	return []LinkType{
		LinkTypeDiscoverable,
		LinkTypeDocumentation,
		LinkTypeGeneric,
		LinkTypeKnowMore,
		LinkTypeSignUp,
		LinkTypeSpecification,
		LinkTypeTermsAndConditions,
	}
}

// MeteringType - Represents Metering type (eg one-time or recurrent).
type MeteringType string

const (
	// MeteringTypeAdhoc - Adhoc billing.
	MeteringTypeAdhoc MeteringType = "Adhoc"
	// MeteringTypeOneTime - One time billing.
	MeteringTypeOneTime MeteringType = "OneTime"
	// MeteringTypeRecurring - Recurring billing.
	MeteringTypeRecurring MeteringType = "Recurring"
)

// PossibleMeteringTypeValues returns the possible values for the MeteringType const type.
func PossibleMeteringTypeValues() []MeteringType {
	return []MeteringType{
		MeteringTypeAdhoc,
		MeteringTypeOneTime,
		MeteringTypeRecurring,
	}
}

// NotificationStageName - Name of the stage.
type NotificationStageName string

const (
	// NotificationStageNameDelivered - Notification at order item delivered to customer.
	NotificationStageNameDelivered NotificationStageName = "Delivered"
	// NotificationStageNameShipped - Notification at order item shipped from microsoft datacenter.
	NotificationStageNameShipped NotificationStageName = "Shipped"
)

// PossibleNotificationStageNameValues returns the possible values for the NotificationStageName const type.
func PossibleNotificationStageNameValues() []NotificationStageName {
	return []NotificationStageName{
		NotificationStageNameDelivered,
		NotificationStageNameShipped,
	}
}

// OrderItemCancellationEnum - Describes whether the order item is cancellable or not.
type OrderItemCancellationEnum string

const (
	// OrderItemCancellationEnumCancellable - Order item can be cancelled without fee.
	OrderItemCancellationEnumCancellable OrderItemCancellationEnum = "Cancellable"
	// OrderItemCancellationEnumCancellableWithFee - Order item can be cancelled with fee.
	OrderItemCancellationEnumCancellableWithFee OrderItemCancellationEnum = "CancellableWithFee"
	// OrderItemCancellationEnumNotCancellable - Order item not cancellable.
	OrderItemCancellationEnumNotCancellable OrderItemCancellationEnum = "NotCancellable"
)

// PossibleOrderItemCancellationEnumValues returns the possible values for the OrderItemCancellationEnum const type.
func PossibleOrderItemCancellationEnumValues() []OrderItemCancellationEnum {
	return []OrderItemCancellationEnum{
		OrderItemCancellationEnumCancellable,
		OrderItemCancellationEnumCancellableWithFee,
		OrderItemCancellationEnumNotCancellable,
	}
}

// OrderItemReturnEnum - Describes whether the order item is returnable or not.
type OrderItemReturnEnum string

const (
	// OrderItemReturnEnumNotReturnable - Order item not returnable.
	OrderItemReturnEnumNotReturnable OrderItemReturnEnum = "NotReturnable"
	// OrderItemReturnEnumReturnable - Order item can be returned without fee.
	OrderItemReturnEnumReturnable OrderItemReturnEnum = "Returnable"
	// OrderItemReturnEnumReturnableWithFee - Order item can be returned with fee.
	OrderItemReturnEnumReturnableWithFee OrderItemReturnEnum = "ReturnableWithFee"
)

// PossibleOrderItemReturnEnumValues returns the possible values for the OrderItemReturnEnum const type.
func PossibleOrderItemReturnEnumValues() []OrderItemReturnEnum {
	return []OrderItemReturnEnum{
		OrderItemReturnEnumNotReturnable,
		OrderItemReturnEnumReturnable,
		OrderItemReturnEnumReturnableWithFee,
	}
}

// OrderItemType - Order item type.
type OrderItemType string

const (
	// OrderItemTypePurchase - Purchase OrderItem.
	OrderItemTypePurchase OrderItemType = "Purchase"
	// OrderItemTypeRental - Rental OrderItem.
	OrderItemTypeRental OrderItemType = "Rental"
)

// PossibleOrderItemTypeValues returns the possible values for the OrderItemType const type.
func PossibleOrderItemTypeValues() []OrderItemType {
	return []OrderItemType{
		OrderItemTypePurchase,
		OrderItemTypeRental,
	}
}

// OrderMode - Defines the mode of the Order item.
type OrderMode string

const (
	// OrderModeDefault - Default Order mode.
	OrderModeDefault OrderMode = "Default"
	// OrderModeDoNotFulfill - Mode in which the Order will not be fulfilled.
	OrderModeDoNotFulfill OrderMode = "DoNotFulfill"
)

// PossibleOrderModeValues returns the possible values for the OrderMode const type.
func PossibleOrderModeValues() []OrderMode {
	return []OrderMode{
		OrderModeDefault,
		OrderModeDoNotFulfill,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// StageName - Stage name.
type StageName string

const (
	// StageNameCancelled - Order has been cancelled.
	StageNameCancelled StageName = "Cancelled"
	// StageNameConfirmed - Order is confirmed.
	StageNameConfirmed StageName = "Confirmed"
	// StageNameDelivered - Order is delivered to customer.
	StageNameDelivered StageName = "Delivered"
	// StageNameInReview - Order is currently in draft mode and can still be cancelled.
	StageNameInReview StageName = "InReview"
	// StageNameInUse - Order is in use at customer site.
	StageNameInUse StageName = "InUse"
	// StageNamePlaced - Currently in draft mode and can still be cancelled.
	StageNamePlaced StageName = "Placed"
	// StageNameReadyToShip - Order is ready to ship.
	StageNameReadyToShip StageName = "ReadyToShip"
	// StageNameReturnCompleted - Return has now completed.
	StageNameReturnCompleted StageName = "ReturnCompleted"
	// StageNameReturnInitiated - Return has been initiated by customer.
	StageNameReturnInitiated StageName = "ReturnInitiated"
	// StageNameReturnPickedUp - Order is in transit from customer to Microsoft.
	StageNameReturnPickedUp StageName = "ReturnPickedUp"
	// StageNameReturnedToMicrosoft - Order has been received back to Microsoft.
	StageNameReturnedToMicrosoft StageName = "ReturnedToMicrosoft"
	// StageNameShipped - Order is in transit to customer.
	StageNameShipped StageName = "Shipped"
)

// PossibleStageNameValues returns the possible values for the StageName const type.
func PossibleStageNameValues() []StageName {
	return []StageName{
		StageNameCancelled,
		StageNameConfirmed,
		StageNameDelivered,
		StageNameInReview,
		StageNameInUse,
		StageNamePlaced,
		StageNameReadyToShip,
		StageNameReturnCompleted,
		StageNameReturnInitiated,
		StageNameReturnPickedUp,
		StageNameReturnedToMicrosoft,
		StageNameShipped,
	}
}

// StageStatus - Stage status.
type StageStatus string

const (
	// StageStatusCancelled - Stage has been cancelled.
	StageStatusCancelled StageStatus = "Cancelled"
	// StageStatusCancelling - Stage is cancelling.
	StageStatusCancelling StageStatus = "Cancelling"
	// StageStatusFailed - Stage has failed.
	StageStatusFailed StageStatus = "Failed"
	// StageStatusInProgress - Stage is in progress.
	StageStatusInProgress StageStatus = "InProgress"
	// StageStatusNone - No status available yet.
	StageStatusNone StageStatus = "None"
	// StageStatusSucceeded - Stage has succeeded.
	StageStatusSucceeded StageStatus = "Succeeded"
)

// PossibleStageStatusValues returns the possible values for the StageStatus const type.
func PossibleStageStatusValues() []StageStatus {
	return []StageStatus{
		StageStatusCancelled,
		StageStatusCancelling,
		StageStatusFailed,
		StageStatusInProgress,
		StageStatusNone,
		StageStatusSucceeded,
	}
}

// SupportedFilterTypes - Type of product filter.
type SupportedFilterTypes string

const (
	// SupportedFilterTypesDoubleEncryptionStatus - Double encryption status.
	SupportedFilterTypesDoubleEncryptionStatus SupportedFilterTypes = "DoubleEncryptionStatus"
	// SupportedFilterTypesShipToCountries - Ship to country.
	SupportedFilterTypesShipToCountries SupportedFilterTypes = "ShipToCountries"
)

// PossibleSupportedFilterTypesValues returns the possible values for the SupportedFilterTypes const type.
func PossibleSupportedFilterTypesValues() []SupportedFilterTypes {
	return []SupportedFilterTypes{
		SupportedFilterTypesDoubleEncryptionStatus,
		SupportedFilterTypesShipToCountries,
	}
}

// TermCommitmentType - Term Commitment Type
type TermCommitmentType string

const (
	// TermCommitmentTypeNone - Pay as you go Term Commitment Model.
	TermCommitmentTypeNone TermCommitmentType = "None"
	// TermCommitmentTypeTimed - Time based Term Commitment Model.
	TermCommitmentTypeTimed TermCommitmentType = "Timed"
	// TermCommitmentTypeTrial - Trial Term Commitment Model.
	TermCommitmentTypeTrial TermCommitmentType = "Trial"
)

// PossibleTermCommitmentTypeValues returns the possible values for the TermCommitmentType const type.
func PossibleTermCommitmentTypeValues() []TermCommitmentType {
	return []TermCommitmentType{
		TermCommitmentTypeNone,
		TermCommitmentTypeTimed,
		TermCommitmentTypeTrial,
	}
}

// TransportShipmentTypes - Indicates Shipment Logistics type that the customer preferred.
type TransportShipmentTypes string

const (
	// TransportShipmentTypesCustomerManaged - Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged TransportShipmentTypes = "CustomerManaged"
	// TransportShipmentTypesMicrosoftManaged - Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged TransportShipmentTypes = "MicrosoftManaged"
)

// PossibleTransportShipmentTypesValues returns the possible values for the TransportShipmentTypes const type.
func PossibleTransportShipmentTypesValues() []TransportShipmentTypes {
	return []TransportShipmentTypes{
		TransportShipmentTypesCustomerManaged,
		TransportShipmentTypesMicrosoftManaged,
	}
}

// WeightMeasurementUnit - Unit for the dimensions of weight.
type WeightMeasurementUnit string

const (
	// WeightMeasurementUnitKGS - Kilograms.
	WeightMeasurementUnitKGS WeightMeasurementUnit = "KGS"
	// WeightMeasurementUnitLBS - Pounds.
	WeightMeasurementUnitLBS WeightMeasurementUnit = "LBS"
)

// PossibleWeightMeasurementUnitValues returns the possible values for the WeightMeasurementUnit const type.
func PossibleWeightMeasurementUnitValues() []WeightMeasurementUnit {
	return []WeightMeasurementUnit{
		WeightMeasurementUnitKGS,
		WeightMeasurementUnitLBS,
	}
}

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
	moduleVersion = "v2.0.0-beta.2"
)

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

// ActivationState - State of security domain activation
type ActivationState string

const (
	ActivationStateActive       ActivationState = "Active"
	ActivationStateFailed       ActivationState = "Failed"
	ActivationStateNotActivated ActivationState = "NotActivated"
	ActivationStateNotDefined   ActivationState = "NotDefined"
	ActivationStateUnknown      ActivationState = "Unknown"
)

// PossibleActivationStateValues returns the possible values for the ActivationState const type.
func PossibleActivationStateValues() []ActivationState {
	return []ActivationState{
		ActivationStateActive,
		ActivationStateFailed,
		ActivationStateNotActivated,
		ActivationStateNotDefined,
		ActivationStateUnknown,
	}
}

// AutoGeneratedDomainNameLabelScope - The Cloud HSM Cluster's auto-generated Domain Name Label Scope
type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScopeNoReuse            AutoGeneratedDomainNameLabelScope = "NoReuse"
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse  AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	AutoGeneratedDomainNameLabelScopeTenantReuse        AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeNoReuse,
		AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
}

// BackupRestoreOperationStatus - Status of the backup/restore operation
type BackupRestoreOperationStatus string

const (
	BackupRestoreOperationStatusCancelled  BackupRestoreOperationStatus = "Cancelled"
	BackupRestoreOperationStatusFailed     BackupRestoreOperationStatus = "Failed"
	BackupRestoreOperationStatusInProgress BackupRestoreOperationStatus = "InProgress"
	BackupRestoreOperationStatusSucceeded  BackupRestoreOperationStatus = "Succeeded"
)

// PossibleBackupRestoreOperationStatusValues returns the possible values for the BackupRestoreOperationStatus const type.
func PossibleBackupRestoreOperationStatusValues() []BackupRestoreOperationStatus {
	return []BackupRestoreOperationStatus{
		BackupRestoreOperationStatusCancelled,
		BackupRestoreOperationStatusFailed,
		BackupRestoreOperationStatusInProgress,
		BackupRestoreOperationStatusSucceeded,
	}
}

// CloudHsmClusterSKUFamily - Sku family of the Cloud HSM Cluster
type CloudHsmClusterSKUFamily string

const (
	CloudHsmClusterSKUFamilyB CloudHsmClusterSKUFamily = "B"
)

// PossibleCloudHsmClusterSKUFamilyValues returns the possible values for the CloudHsmClusterSKUFamily const type.
func PossibleCloudHsmClusterSKUFamilyValues() []CloudHsmClusterSKUFamily {
	return []CloudHsmClusterSKUFamily{
		CloudHsmClusterSKUFamilyB,
	}
}

// CloudHsmClusterSKUName - Sku name of the Cloud HSM Cluster
type CloudHsmClusterSKUName string

const (
	CloudHsmClusterSKUNameStandardB1  CloudHsmClusterSKUName = "Standard_B1"
	CloudHsmClusterSKUNameStandardB10 CloudHsmClusterSKUName = "Standard B10"
)

// PossibleCloudHsmClusterSKUNameValues returns the possible values for the CloudHsmClusterSKUName const type.
func PossibleCloudHsmClusterSKUNameValues() []CloudHsmClusterSKUName {
	return []CloudHsmClusterSKUName{
		CloudHsmClusterSKUNameStandardB1,
		CloudHsmClusterSKUNameStandardB10,
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

// IdentityType - The type of identity.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// JSONWebKeyType - Provisioning state.
type JSONWebKeyType string

const (
	// JSONWebKeyTypeAllocating - A device is currently being allocated for the dedicated HSM resource.
	JSONWebKeyTypeAllocating JSONWebKeyType = "Allocating"
	// JSONWebKeyTypeCheckingQuota - Validating the subscription has sufficient quota to allocate a dedicated HSM device.
	JSONWebKeyTypeCheckingQuota JSONWebKeyType = "CheckingQuota"
	// JSONWebKeyTypeConnecting - The dedicated HSM is being connected to the virtual network.
	JSONWebKeyTypeConnecting JSONWebKeyType = "Connecting"
	// JSONWebKeyTypeDeleting - The dedicated HSM is currently being deleted.
	JSONWebKeyTypeDeleting JSONWebKeyType = "Deleting"
	// JSONWebKeyTypeFailed - Provisioning of the dedicated HSM has failed.
	JSONWebKeyTypeFailed JSONWebKeyType = "Failed"
	// JSONWebKeyTypeProvisioning - The dedicated HSM is currently being provisioned.
	JSONWebKeyTypeProvisioning JSONWebKeyType = "Provisioning"
	// JSONWebKeyTypeSucceeded - The dedicated HSM has been fully provisioned.
	JSONWebKeyTypeSucceeded JSONWebKeyType = "Succeeded"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeAllocating,
		JSONWebKeyTypeCheckingQuota,
		JSONWebKeyTypeConnecting,
		JSONWebKeyTypeDeleting,
		JSONWebKeyTypeFailed,
		JSONWebKeyTypeProvisioning,
		JSONWebKeyTypeSucceeded,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
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

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCanceled      PrivateEndpointConnectionProvisioningState = "Canceled"
	PrivateEndpointConnectionProvisioningStateCreating      PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting      PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed        PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateInternalError PrivateEndpointConnectionProvisioningState = "InternalError"
	PrivateEndpointConnectionProvisioningStateSucceeded     PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating      PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCanceled,
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateInternalError,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - The Cloud HSM Cluster's provisioningState
type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
	}
}

// PublicNetworkAccess - The Cloud HSM Cluster public network access
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
	}
}

// SKUName - SKU of the dedicated HSM
type SKUName string

const (
	// SKUNamePayShield10KLMK1CPS250 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 250 calls per second.
	SKUNamePayShield10KLMK1CPS250 SKUName = "payShield10K_LMK1_CPS250"
	// SKUNamePayShield10KLMK1CPS2500 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 2500 calls per second.
	SKUNamePayShield10KLMK1CPS2500 SKUName = "payShield10K_LMK1_CPS2500"
	// SKUNamePayShield10KLMK1CPS60 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 1 local master key which supports up to 60 calls per second.
	SKUNamePayShield10KLMK1CPS60 SKUName = "payShield10K_LMK1_CPS60"
	// SKUNamePayShield10KLMK2CPS250 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 250 calls per second.
	SKUNamePayShield10KLMK2CPS250 SKUName = "payShield10K_LMK2_CPS250"
	// SKUNamePayShield10KLMK2CPS2500 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 2500 calls per second.
	SKUNamePayShield10KLMK2CPS2500 SKUName = "payShield10K_LMK2_CPS2500"
	// SKUNamePayShield10KLMK2CPS60 - The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device
	// with 2 local master keys which supports up to 60 calls per second.
	SKUNamePayShield10KLMK2CPS60 SKUName = "payShield10K_LMK2_CPS60"
	// SKUNameSafeNetLunaNetworkHSMA790 - The dedicated HSM is a Safenet Luna Network HSM A790 device.
	SKUNameSafeNetLunaNetworkHSMA790 SKUName = "SafeNet Luna Network HSM A790"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePayShield10KLMK1CPS250,
		SKUNamePayShield10KLMK1CPS2500,
		SKUNamePayShield10KLMK1CPS60,
		SKUNamePayShield10KLMK2CPS250,
		SKUNamePayShield10KLMK2CPS2500,
		SKUNamePayShield10KLMK2CPS60,
		SKUNameSafeNetLunaNetworkHSMA790,
	}
}

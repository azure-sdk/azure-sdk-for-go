//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage

const (
	moduleName    = "armstorage"
	moduleVersion = "v2.0.0"
)

// AccessTier - Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium'
// access tier is the default value for premium block blobs storage account type and it cannot
// be changed for the premium block blobs storage account type.
type AccessTier string

const (
	AccessTierHot     AccessTier = "Hot"
	AccessTierCool    AccessTier = "Cool"
	AccessTierPremium AccessTier = "Premium"
)

// PossibleAccessTierValues returns the possible values for the AccessTier const type.
func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{
		AccessTierHot,
		AccessTierCool,
		AccessTierPremium,
	}
}

// AccountImmutabilityPolicyState - The ImmutabilityPolicy state defines the mode of the policy. Disabled state disables the
// policy, Unlocked state allows increase and decrease of immutability retention time and also allows toggling
// allowProtectedAppendWrites property, Locked state only allows the increase of the immutability retention time. A policy
// can only be created in a Disabled or Unlocked state and can be toggled between
// the two states. Only a policy in an Unlocked state can transition to a Locked state which cannot be reverted.
type AccountImmutabilityPolicyState string

const (
	AccountImmutabilityPolicyStateDisabled AccountImmutabilityPolicyState = "Disabled"
	AccountImmutabilityPolicyStateLocked   AccountImmutabilityPolicyState = "Locked"
	AccountImmutabilityPolicyStateUnlocked AccountImmutabilityPolicyState = "Unlocked"
)

// PossibleAccountImmutabilityPolicyStateValues returns the possible values for the AccountImmutabilityPolicyState const type.
func PossibleAccountImmutabilityPolicyStateValues() []AccountImmutabilityPolicyState {
	return []AccountImmutabilityPolicyState{
		AccountImmutabilityPolicyStateDisabled,
		AccountImmutabilityPolicyStateLocked,
		AccountImmutabilityPolicyStateUnlocked,
	}
}

// AccountStatus - Gets the status indicating whether the primary location of the storage account is available or unavailable.
type AccountStatus string

const (
	AccountStatusAvailable   AccountStatus = "available"
	AccountStatusUnavailable AccountStatus = "unavailable"
)

// PossibleAccountStatusValues returns the possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{
		AccountStatusAvailable,
		AccountStatusUnavailable,
	}
}

// AccountType - Specifies the Active Directory account type for Azure Storage.
type AccountType string

const (
	AccountTypeComputer AccountType = "Computer"
	AccountTypeUser     AccountType = "User"
)

// PossibleAccountTypeValues returns the possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{
		AccountTypeComputer,
		AccountTypeUser,
	}
}

// AllowedCopyScope - Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
type AllowedCopyScope string

const (
	AllowedCopyScopeAAD         AllowedCopyScope = "AAD"
	AllowedCopyScopePrivateLink AllowedCopyScope = "PrivateLink"
)

// PossibleAllowedCopyScopeValues returns the possible values for the AllowedCopyScope const type.
func PossibleAllowedCopyScopeValues() []AllowedCopyScope {
	return []AllowedCopyScope{
		AllowedCopyScopeAAD,
		AllowedCopyScopePrivateLink,
	}
}

type AllowedMethods string

const (
	AllowedMethodsDELETE  AllowedMethods = "DELETE"
	AllowedMethodsGET     AllowedMethods = "GET"
	AllowedMethodsHEAD    AllowedMethods = "HEAD"
	AllowedMethodsMERGE   AllowedMethods = "MERGE"
	AllowedMethodsOPTIONS AllowedMethods = "OPTIONS"
	AllowedMethodsPATCH   AllowedMethods = "PATCH"
	AllowedMethodsPOST    AllowedMethods = "POST"
	AllowedMethodsPUT     AllowedMethods = "PUT"
)

// PossibleAllowedMethodsValues returns the possible values for the AllowedMethods const type.
func PossibleAllowedMethodsValues() []AllowedMethods {
	return []AllowedMethods{
		AllowedMethodsDELETE,
		AllowedMethodsGET,
		AllowedMethodsHEAD,
		AllowedMethodsMERGE,
		AllowedMethodsOPTIONS,
		AllowedMethodsPATCH,
		AllowedMethodsPOST,
		AllowedMethodsPUT,
	}
}

type BlobInventoryPolicyName string

const (
	BlobInventoryPolicyNameDefault BlobInventoryPolicyName = "default"
)

// PossibleBlobInventoryPolicyNameValues returns the possible values for the BlobInventoryPolicyName const type.
func PossibleBlobInventoryPolicyNameValues() []BlobInventoryPolicyName {
	return []BlobInventoryPolicyName{
		BlobInventoryPolicyNameDefault,
	}
}

// BlobRestoreProgressStatus - The status of blob restore progress. Possible values are: - InProgress: Indicates that blob
// restore is ongoing. - Complete: Indicates that blob restore has been completed successfully. - Failed:
// Indicates that blob restore is failed.
type BlobRestoreProgressStatus string

const (
	BlobRestoreProgressStatusComplete   BlobRestoreProgressStatus = "Complete"
	BlobRestoreProgressStatusFailed     BlobRestoreProgressStatus = "Failed"
	BlobRestoreProgressStatusInProgress BlobRestoreProgressStatus = "InProgress"
)

// PossibleBlobRestoreProgressStatusValues returns the possible values for the BlobRestoreProgressStatus const type.
func PossibleBlobRestoreProgressStatusValues() []BlobRestoreProgressStatus {
	return []BlobRestoreProgressStatus{
		BlobRestoreProgressStatusComplete,
		BlobRestoreProgressStatusFailed,
		BlobRestoreProgressStatusInProgress,
	}
}

// Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
// Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none
// of those traffics.
type Bypass string

const (
	BypassAzureServices Bypass = "AzureServices"
	BypassLogging       Bypass = "Logging"
	BypassMetrics       Bypass = "Metrics"
	BypassNone          Bypass = "None"
)

// PossibleBypassValues returns the possible values for the Bypass const type.
func PossibleBypassValues() []Bypass {
	return []Bypass{
		BypassAzureServices,
		BypassLogging,
		BypassMetrics,
		BypassNone,
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

// DNSEndpointType - Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts
// in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL
// will have an alphanumeric DNS Zone identifier.
type DNSEndpointType string

const (
	DNSEndpointTypeAzureDNSZone DNSEndpointType = "AzureDnsZone"
	DNSEndpointTypeStandard     DNSEndpointType = "Standard"
)

// PossibleDNSEndpointTypeValues returns the possible values for the DNSEndpointType const type.
func PossibleDNSEndpointTypeValues() []DNSEndpointType {
	return []DNSEndpointType{
		DNSEndpointTypeAzureDNSZone,
		DNSEndpointTypeStandard,
	}
}

// DefaultAction - Specifies the default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// DefaultSharePermission - Default share permission for users using Kerberos authentication if RBAC role is not assigned.
type DefaultSharePermission string

const (
	DefaultSharePermissionNone                                       DefaultSharePermission = "None"
	DefaultSharePermissionStorageFileDataSmbShareContributor         DefaultSharePermission = "StorageFileDataSmbShareContributor"
	DefaultSharePermissionStorageFileDataSmbShareElevatedContributor DefaultSharePermission = "StorageFileDataSmbShareElevatedContributor"
	DefaultSharePermissionStorageFileDataSmbShareReader              DefaultSharePermission = "StorageFileDataSmbShareReader"
)

// PossibleDefaultSharePermissionValues returns the possible values for the DefaultSharePermission const type.
func PossibleDefaultSharePermissionValues() []DefaultSharePermission {
	return []DefaultSharePermission{
		DefaultSharePermissionNone,
		DefaultSharePermissionStorageFileDataSmbShareContributor,
		DefaultSharePermissionStorageFileDataSmbShareElevatedContributor,
		DefaultSharePermissionStorageFileDataSmbShareReader,
	}
}

// DirectoryServiceOptions - Indicates the directory service used. Note that this enum may be extended in the future.
type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsAADDS   DirectoryServiceOptions = "AADDS"
	DirectoryServiceOptionsAADKERB DirectoryServiceOptions = "AADKERB"
	DirectoryServiceOptionsAD      DirectoryServiceOptions = "AD"
	DirectoryServiceOptionsNone    DirectoryServiceOptions = "None"
)

// PossibleDirectoryServiceOptionsValues returns the possible values for the DirectoryServiceOptions const type.
func PossibleDirectoryServiceOptionsValues() []DirectoryServiceOptions {
	return []DirectoryServiceOptions{
		DirectoryServiceOptionsAADDS,
		DirectoryServiceOptionsAADKERB,
		DirectoryServiceOptionsAD,
		DirectoryServiceOptionsNone,
	}
}

// EnabledProtocols - The authentication protocol that is used for the file share. Can only be specified when creating a share.
type EnabledProtocols string

const (
	EnabledProtocolsNFS EnabledProtocols = "NFS"
	EnabledProtocolsSMB EnabledProtocols = "SMB"
)

// PossibleEnabledProtocolsValues returns the possible values for the EnabledProtocols const type.
func PossibleEnabledProtocolsValues() []EnabledProtocols {
	return []EnabledProtocols{
		EnabledProtocolsNFS,
		EnabledProtocolsSMB,
	}
}

// EncryptionScopeSource - The provider for the encryption scope. Possible values (case-insensitive): Microsoft.Storage, Microsoft.KeyVault.
type EncryptionScopeSource string

const (
	EncryptionScopeSourceMicrosoftKeyVault EncryptionScopeSource = "Microsoft.KeyVault"
	EncryptionScopeSourceMicrosoftStorage  EncryptionScopeSource = "Microsoft.Storage"
)

// PossibleEncryptionScopeSourceValues returns the possible values for the EncryptionScopeSource const type.
func PossibleEncryptionScopeSourceValues() []EncryptionScopeSource {
	return []EncryptionScopeSource{
		EncryptionScopeSourceMicrosoftKeyVault,
		EncryptionScopeSourceMicrosoftStorage,
	}
}

// EncryptionScopeState - The state of the encryption scope. Possible values (case-insensitive): Enabled, Disabled.
type EncryptionScopeState string

const (
	EncryptionScopeStateDisabled EncryptionScopeState = "Disabled"
	EncryptionScopeStateEnabled  EncryptionScopeState = "Enabled"
)

// PossibleEncryptionScopeStateValues returns the possible values for the EncryptionScopeState const type.
func PossibleEncryptionScopeStateValues() []EncryptionScopeState {
	return []EncryptionScopeState{
		EncryptionScopeStateDisabled,
		EncryptionScopeStateEnabled,
	}
}

// ExpirationAction - The SAS expiration action. Can only be Log.
type ExpirationAction string

const (
	ExpirationActionLog ExpirationAction = "Log"
)

// PossibleExpirationActionValues returns the possible values for the ExpirationAction const type.
func PossibleExpirationActionValues() []ExpirationAction {
	return []ExpirationAction{
		ExpirationActionLog,
	}
}

// ExtendedLocationTypes - The type of extendedLocation.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone ExtendedLocationTypes = "EdgeZone"
)

// PossibleExtendedLocationTypesValues returns the possible values for the ExtendedLocationTypes const type.
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return []ExtendedLocationTypes{
		ExtendedLocationTypesEdgeZone,
	}
}

// Format - This is a required field, it specifies the format for the inventory files.
type Format string

const (
	FormatCSV     Format = "Csv"
	FormatParquet Format = "Parquet"
)

// PossibleFormatValues returns the possible values for the Format const type.
func PossibleFormatValues() []Format {
	return []Format{
		FormatCSV,
		FormatParquet,
	}
}

// GeoReplicationStatus - The status of the secondary location. Possible values are: - Live: Indicates that the secondary
// location is active and operational. - Bootstrap: Indicates initial synchronization from the primary
// location to the secondary location is in progress.This typically occurs when replication is first enabled. - Unavailable:
// Indicates that the secondary location is temporarily unavailable.
type GeoReplicationStatus string

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = "Bootstrap"
	GeoReplicationStatusLive        GeoReplicationStatus = "Live"
	GeoReplicationStatusUnavailable GeoReplicationStatus = "Unavailable"
)

// PossibleGeoReplicationStatusValues returns the possible values for the GeoReplicationStatus const type.
func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return []GeoReplicationStatus{
		GeoReplicationStatusBootstrap,
		GeoReplicationStatusLive,
		GeoReplicationStatusUnavailable,
	}
}

// HTTPProtocol - The protocol permitted for a request made with the account SAS.
type HTTPProtocol string

const (
	HTTPProtocolHTTPSHTTP HTTPProtocol = "https,http"
	HTTPProtocolHTTPS     HTTPProtocol = "https"
)

// PossibleHTTPProtocolValues returns the possible values for the HTTPProtocol const type.
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return []HTTPProtocol{
		HTTPProtocolHTTPSHTTP,
		HTTPProtocolHTTPS,
	}
}

// IdentityType - The identity type.
type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned,UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

// ImmutabilityPolicyState - The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
type ImmutabilityPolicyState string

const (
	ImmutabilityPolicyStateLocked   ImmutabilityPolicyState = "Locked"
	ImmutabilityPolicyStateUnlocked ImmutabilityPolicyState = "Unlocked"
)

// PossibleImmutabilityPolicyStateValues returns the possible values for the ImmutabilityPolicyState const type.
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return []ImmutabilityPolicyState{
		ImmutabilityPolicyStateLocked,
		ImmutabilityPolicyStateUnlocked,
	}
}

// ImmutabilityPolicyUpdateType - The ImmutabilityPolicy update type of a blob container, possible values include: put, lock
// and extend.
type ImmutabilityPolicyUpdateType string

const (
	ImmutabilityPolicyUpdateTypeExtend ImmutabilityPolicyUpdateType = "extend"
	ImmutabilityPolicyUpdateTypeLock   ImmutabilityPolicyUpdateType = "lock"
	ImmutabilityPolicyUpdateTypePut    ImmutabilityPolicyUpdateType = "put"
)

// PossibleImmutabilityPolicyUpdateTypeValues returns the possible values for the ImmutabilityPolicyUpdateType const type.
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return []ImmutabilityPolicyUpdateType{
		ImmutabilityPolicyUpdateTypeExtend,
		ImmutabilityPolicyUpdateTypeLock,
		ImmutabilityPolicyUpdateTypePut,
	}
}

// InventoryRuleType - The valid value is Inventory
type InventoryRuleType string

const (
	InventoryRuleTypeInventory InventoryRuleType = "Inventory"
)

// PossibleInventoryRuleTypeValues returns the possible values for the InventoryRuleType const type.
func PossibleInventoryRuleTypeValues() []InventoryRuleType {
	return []InventoryRuleType{
		InventoryRuleTypeInventory,
	}
}

// KeyPermission - Permissions for the key -- read-only or full permissions.
type KeyPermission string

const (
	KeyPermissionRead KeyPermission = "Read"
	KeyPermissionFull KeyPermission = "Full"
)

// PossibleKeyPermissionValues returns the possible values for the KeyPermission const type.
func PossibleKeyPermissionValues() []KeyPermission {
	return []KeyPermission{
		KeyPermissionRead,
		KeyPermissionFull,
	}
}

// KeySource - The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Storage, Microsoft.Keyvault
type KeySource string

const (
	KeySourceMicrosoftKeyvault KeySource = "Microsoft.Keyvault"
	KeySourceMicrosoftStorage  KeySource = "Microsoft.Storage"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftKeyvault,
		KeySourceMicrosoftStorage,
	}
}

// KeyType - Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
// encryption key will be used. 'Service' key type implies that a default service key is used.
type KeyType string

const (
	KeyTypeAccount KeyType = "Account"
	KeyTypeService KeyType = "Service"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypeAccount,
		KeyTypeService,
	}
}

// Kind - Indicates the type of storage account.
type Kind string

const (
	KindBlobStorage      Kind = "BlobStorage"
	KindBlockBlobStorage Kind = "BlockBlobStorage"
	KindFileStorage      Kind = "FileStorage"
	KindStorage          Kind = "Storage"
	KindStorageV2        Kind = "StorageV2"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindBlobStorage,
		KindBlockBlobStorage,
		KindFileStorage,
		KindStorage,
		KindStorageV2,
	}
}

// LargeFileSharesState - Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled LargeFileSharesState = "Disabled"
	LargeFileSharesStateEnabled  LargeFileSharesState = "Enabled"
)

// PossibleLargeFileSharesStateValues returns the possible values for the LargeFileSharesState const type.
func PossibleLargeFileSharesStateValues() []LargeFileSharesState {
	return []LargeFileSharesState{
		LargeFileSharesStateDisabled,
		LargeFileSharesStateEnabled,
	}
}

// LeaseContainerRequest - Specifies the lease action. Can be one of the available actions.
type LeaseContainerRequest string

const (
	LeaseContainerRequestAcquire LeaseContainerRequest = "Acquire"
	LeaseContainerRequestBreak   LeaseContainerRequest = "Break"
	LeaseContainerRequestChange  LeaseContainerRequest = "Change"
	LeaseContainerRequestRelease LeaseContainerRequest = "Release"
	LeaseContainerRequestRenew   LeaseContainerRequest = "Renew"
)

// PossibleLeaseContainerRequestValues returns the possible values for the LeaseContainerRequest const type.
func PossibleLeaseContainerRequestValues() []LeaseContainerRequest {
	return []LeaseContainerRequest{
		LeaseContainerRequestAcquire,
		LeaseContainerRequestBreak,
		LeaseContainerRequestChange,
		LeaseContainerRequestRelease,
		LeaseContainerRequestRenew,
	}
}

// LeaseDuration - Specifies whether the lease on a container is of infinite or fixed duration, only when the container is
// leased.
type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

// PossibleLeaseDurationValues returns the possible values for the LeaseDuration const type.
func PossibleLeaseDurationValues() []LeaseDuration {
	return []LeaseDuration{
		LeaseDurationFixed,
		LeaseDurationInfinite,
	}
}

// LeaseShareAction - Specifies the lease action. Can be one of the available actions.
type LeaseShareAction string

const (
	LeaseShareActionAcquire LeaseShareAction = "Acquire"
	LeaseShareActionBreak   LeaseShareAction = "Break"
	LeaseShareActionChange  LeaseShareAction = "Change"
	LeaseShareActionRelease LeaseShareAction = "Release"
	LeaseShareActionRenew   LeaseShareAction = "Renew"
)

// PossibleLeaseShareActionValues returns the possible values for the LeaseShareAction const type.
func PossibleLeaseShareActionValues() []LeaseShareAction {
	return []LeaseShareAction{
		LeaseShareActionAcquire,
		LeaseShareActionBreak,
		LeaseShareActionChange,
		LeaseShareActionRelease,
		LeaseShareActionRenew,
	}
}

// LeaseState - Lease state of the container.
type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

// PossibleLeaseStateValues returns the possible values for the LeaseState const type.
func PossibleLeaseStateValues() []LeaseState {
	return []LeaseState{
		LeaseStateAvailable,
		LeaseStateBreaking,
		LeaseStateBroken,
		LeaseStateExpired,
		LeaseStateLeased,
	}
}

// LeaseStatus - The lease status of the container.
type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

// PossibleLeaseStatusValues returns the possible values for the LeaseStatus const type.
func PossibleLeaseStatusValues() []LeaseStatus {
	return []LeaseStatus{
		LeaseStatusLocked,
		LeaseStatusUnlocked,
	}
}

type ListContainersInclude string

const (
	ListContainersIncludeDeleted ListContainersInclude = "deleted"
)

// PossibleListContainersIncludeValues returns the possible values for the ListContainersInclude const type.
func PossibleListContainersIncludeValues() []ListContainersInclude {
	return []ListContainersInclude{
		ListContainersIncludeDeleted,
	}
}

type ManagementPolicyName string

const (
	ManagementPolicyNameDefault ManagementPolicyName = "default"
)

// PossibleManagementPolicyNameValues returns the possible values for the ManagementPolicyName const type.
func PossibleManagementPolicyNameValues() []ManagementPolicyName {
	return []ManagementPolicyName{
		ManagementPolicyNameDefault,
	}
}

// MigrationState - This property denotes the container level immutability to object level immutability migration state.
type MigrationState string

const (
	MigrationStateCompleted  MigrationState = "Completed"
	MigrationStateInProgress MigrationState = "InProgress"
)

// PossibleMigrationStateValues returns the possible values for the MigrationState const type.
func PossibleMigrationStateValues() []MigrationState {
	return []MigrationState{
		MigrationStateCompleted,
		MigrationStateInProgress,
	}
}

// MinimumTLSVersion - Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
// 1.0 for this property.
type MinimumTLSVersion string

const (
	MinimumTLSVersionTLS10 MinimumTLSVersion = "TLS1_0"
	MinimumTLSVersionTLS11 MinimumTLSVersion = "TLS1_1"
	MinimumTLSVersionTLS12 MinimumTLSVersion = "TLS1_2"
)

// PossibleMinimumTLSVersionValues returns the possible values for the MinimumTLSVersion const type.
func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return []MinimumTLSVersion{
		MinimumTLSVersionTLS10,
		MinimumTLSVersionTLS11,
		MinimumTLSVersionTLS12,
	}
}

// Name - Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
type Name string

const (
	NameAccessTimeTracking Name = "AccessTimeTracking"
)

// PossibleNameValues returns the possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{
		NameAccessTimeTracking,
	}
}

// ObjectType - This is a required field. This field specifies the scope of the inventory created either at the blob or container
// level.
type ObjectType string

const (
	ObjectTypeBlob      ObjectType = "Blob"
	ObjectTypeContainer ObjectType = "Container"
)

// PossibleObjectTypeValues returns the possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{
		ObjectTypeBlob,
		ObjectTypeContainer,
	}
}

// Permissions - The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List
// (l), Add (a), Create (c), Update (u) and Process (p).
type Permissions string

const (
	PermissionsA Permissions = "a"
	PermissionsC Permissions = "c"
	PermissionsD Permissions = "d"
	PermissionsL Permissions = "l"
	PermissionsP Permissions = "p"
	PermissionsR Permissions = "r"
	PermissionsU Permissions = "u"
	PermissionsW Permissions = "w"
)

// PossiblePermissionsValues returns the possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{
		PermissionsA,
		PermissionsC,
		PermissionsD,
		PermissionsL,
		PermissionsP,
		PermissionsR,
		PermissionsU,
		PermissionsW,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
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

// ProvisioningState - Gets the status of the storage account at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateResolvingDNS ProvisioningState = "ResolvingDNS"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateResolvingDNS,
		ProvisioningStateSucceeded,
	}
}

// PublicAccess - Specifies whether data in the container may be accessed publicly and the level of access.
type PublicAccess string

const (
	PublicAccessContainer PublicAccess = "Container"
	PublicAccessBlob      PublicAccess = "Blob"
	PublicAccessNone      PublicAccess = "None"
)

// PossiblePublicAccessValues returns the possible values for the PublicAccess const type.
func PossiblePublicAccessValues() []PublicAccess {
	return []PublicAccess{
		PublicAccessContainer,
		PublicAccessBlob,
		PublicAccessNone,
	}
}

// PublicNetworkAccess - Allow or disallow public network access to Storage Account. Value is optional but if passed in, must
// be 'Enabled' or 'Disabled'.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// Reason - Gets the reason that a storage account name could not be used. The Reason element is only returned if NameAvailable
// is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

// ReasonCode - The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". Quota Id
// is set when the SKU has requiredQuotas parameter as the subscription does not belong to that
// quota. The "NotAvailableForSubscription" is related to capacity at DC.
type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaID                     ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns the possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeQuotaID,
	}
}

// RootSquashType - The property is for NFS share only. The default is NoRootSquash.
type RootSquashType string

const (
	RootSquashTypeAllSquash    RootSquashType = "AllSquash"
	RootSquashTypeNoRootSquash RootSquashType = "NoRootSquash"
	RootSquashTypeRootSquash   RootSquashType = "RootSquash"
)

// PossibleRootSquashTypeValues returns the possible values for the RootSquashType const type.
func PossibleRootSquashTypeValues() []RootSquashType {
	return []RootSquashType{
		RootSquashTypeAllSquash,
		RootSquashTypeNoRootSquash,
		RootSquashTypeRootSquash,
	}
}

// RoutingChoice - Routing Choice defines the kind of network routing opted by the user.
type RoutingChoice string

const (
	RoutingChoiceInternetRouting  RoutingChoice = "InternetRouting"
	RoutingChoiceMicrosoftRouting RoutingChoice = "MicrosoftRouting"
)

// PossibleRoutingChoiceValues returns the possible values for the RoutingChoice const type.
func PossibleRoutingChoiceValues() []RoutingChoice {
	return []RoutingChoice{
		RoutingChoiceInternetRouting,
		RoutingChoiceMicrosoftRouting,
	}
}

// RuleType - The valid value is Lifecycle
type RuleType string

const (
	RuleTypeLifecycle RuleType = "Lifecycle"
)

// PossibleRuleTypeValues returns the possible values for the RuleType const type.
func PossibleRuleTypeValues() []RuleType {
	return []RuleType{
		RuleTypeLifecycle,
	}
}

// SKUConversionStatus - This property indicates the current sku conversion status.
type SKUConversionStatus string

const (
	SKUConversionStatusFailed     SKUConversionStatus = "Failed"
	SKUConversionStatusInProgress SKUConversionStatus = "InProgress"
	SKUConversionStatusSucceeded  SKUConversionStatus = "Succeeded"
)

// PossibleSKUConversionStatusValues returns the possible values for the SKUConversionStatus const type.
func PossibleSKUConversionStatusValues() []SKUConversionStatus {
	return []SKUConversionStatus{
		SKUConversionStatusFailed,
		SKUConversionStatusInProgress,
		SKUConversionStatusSucceeded,
	}
}

// SKUName - The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
// accountType.
type SKUName string

const (
	SKUNamePremiumLRS     SKUName = "Premium_LRS"
	SKUNamePremiumZRS     SKUName = "Premium_ZRS"
	SKUNameStandardGRS    SKUName = "Standard_GRS"
	SKUNameStandardGZRS   SKUName = "Standard_GZRS"
	SKUNameStandardLRS    SKUName = "Standard_LRS"
	SKUNameStandardRAGRS  SKUName = "Standard_RAGRS"
	SKUNameStandardRAGZRS SKUName = "Standard_RAGZRS"
	SKUNameStandardZRS    SKUName = "Standard_ZRS"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremiumLRS,
		SKUNamePremiumZRS,
		SKUNameStandardGRS,
		SKUNameStandardGZRS,
		SKUNameStandardLRS,
		SKUNameStandardRAGRS,
		SKUNameStandardRAGZRS,
		SKUNameStandardZRS,
	}
}

// SKUTier - The SKU tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierStandard,
		SKUTierPremium,
	}
}

// Schedule - This is a required field. This field is used to schedule an inventory formation.
type Schedule string

const (
	ScheduleDaily  Schedule = "Daily"
	ScheduleWeekly Schedule = "Weekly"
)

// PossibleScheduleValues returns the possible values for the Schedule const type.
func PossibleScheduleValues() []Schedule {
	return []Schedule{
		ScheduleDaily,
		ScheduleWeekly,
	}
}

// Services - The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t),
// File (f).
type Services string

const (
	ServicesB Services = "b"
	ServicesF Services = "f"
	ServicesQ Services = "q"
	ServicesT Services = "t"
)

// PossibleServicesValues returns the possible values for the Services const type.
func PossibleServicesValues() []Services {
	return []Services{
		ServicesB,
		ServicesF,
		ServicesQ,
		ServicesT,
	}
}

// ShareAccessTier - Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot,
// and Cool. FileStorage account can choose Premium.
type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierPremium              ShareAccessTier = "Premium"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

// PossibleShareAccessTierValues returns the possible values for the ShareAccessTier const type.
func PossibleShareAccessTierValues() []ShareAccessTier {
	return []ShareAccessTier{
		ShareAccessTierCool,
		ShareAccessTierHot,
		ShareAccessTierPremium,
		ShareAccessTierTransactionOptimized,
	}
}

// SignedResource - The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c),
// File (f), Share (s).
type SignedResource string

const (
	SignedResourceB SignedResource = "b"
	SignedResourceC SignedResource = "c"
	SignedResourceF SignedResource = "f"
	SignedResourceS SignedResource = "s"
)

// PossibleSignedResourceValues returns the possible values for the SignedResource const type.
func PossibleSignedResourceValues() []SignedResource {
	return []SignedResource{
		SignedResourceB,
		SignedResourceC,
		SignedResourceF,
		SignedResourceS,
	}
}

// SignedResourceTypes - The signed resource types that are accessible with the account SAS. Service (s): Access to service-level
// APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs
// for blobs, queue messages, table entities, and files.
type SignedResourceTypes string

const (
	SignedResourceTypesC SignedResourceTypes = "c"
	SignedResourceTypesO SignedResourceTypes = "o"
	SignedResourceTypesS SignedResourceTypes = "s"
)

// PossibleSignedResourceTypesValues returns the possible values for the SignedResourceTypes const type.
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return []SignedResourceTypes{
		SignedResourceTypesC,
		SignedResourceTypesO,
		SignedResourceTypesS,
	}
}

// State - Gets the state of virtual network rule.
type State string

const (
	StateDeprovisioning       State = "Deprovisioning"
	StateFailed               State = "Failed"
	StateNetworkSourceDeleted State = "NetworkSourceDeleted"
	StateProvisioning         State = "Provisioning"
	StateSucceeded            State = "Succeeded"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDeprovisioning,
		StateFailed,
		StateNetworkSourceDeleted,
		StateProvisioning,
		StateSucceeded,
	}
}

type StorageAccountExpand string

const (
	StorageAccountExpandGeoReplicationStats StorageAccountExpand = "geoReplicationStats"
	StorageAccountExpandBlobRestoreStatus   StorageAccountExpand = "blobRestoreStatus"
)

// PossibleStorageAccountExpandValues returns the possible values for the StorageAccountExpand const type.
func PossibleStorageAccountExpandValues() []StorageAccountExpand {
	return []StorageAccountExpand{
		StorageAccountExpandGeoReplicationStats,
		StorageAccountExpandBlobRestoreStatus,
	}
}

// UsageUnit - Gets the unit of measurement.
type UsageUnit string

const (
	UsageUnitCount           UsageUnit = "Count"
	UsageUnitBytes           UsageUnit = "Bytes"
	UsageUnitSeconds         UsageUnit = "Seconds"
	UsageUnitPercent         UsageUnit = "Percent"
	UsageUnitCountsPerSecond UsageUnit = "CountsPerSecond"
	UsageUnitBytesPerSecond  UsageUnit = "BytesPerSecond"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
		UsageUnitBytes,
		UsageUnitSeconds,
		UsageUnitPercent,
		UsageUnitCountsPerSecond,
		UsageUnitBytesPerSecond,
	}
}

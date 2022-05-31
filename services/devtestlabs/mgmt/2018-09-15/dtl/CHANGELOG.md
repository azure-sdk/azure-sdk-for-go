# Unreleased

## Breaking Changes

### Removed Funcs

1. *ApplicableScheduleFragment.UnmarshalJSON([]byte) error
1. *ArtifactSourceFragment.UnmarshalJSON([]byte) error
1. *CustomImageFragment.UnmarshalJSON([]byte) error
1. *DiskFragment.UnmarshalJSON([]byte) error
1. *EnvironmentFragment.UnmarshalJSON([]byte) error
1. *FormulaFragment.UnmarshalJSON([]byte) error
1. *LabFragment.UnmarshalJSON([]byte) error
1. *LabVirtualMachineCreationParameterFragment.UnmarshalJSON([]byte) error
1. *LabVirtualMachineFragment.UnmarshalJSON([]byte) error
1. *NotificationChannelFragment.UnmarshalJSON([]byte) error
1. *PolicyFragment.UnmarshalJSON([]byte) error
1. *ScheduleCreationParameterFragment.UnmarshalJSON([]byte) error
1. *ScheduleFragment.UnmarshalJSON([]byte) error
1. *SecretFragment.UnmarshalJSON([]byte) error
1. *ServiceFabricFragment.UnmarshalJSON([]byte) error
1. *UserFragment.UnmarshalJSON([]byte) error
1. *VirtualNetworkFragment.UnmarshalJSON([]byte) error
1. LabPropertiesFragment.MarshalJSON() ([]byte, error)
1. LabVirtualMachineCreationParameterFragment.MarshalJSON() ([]byte, error)
1. ScheduleCreationParameterFragment.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. ApplicableSchedulePropertiesFragment
1. ArmTemplateParameterPropertiesFragment
1. ArtifactDeploymentStatusPropertiesFragment
1. ArtifactInstallPropertiesFragment
1. ArtifactParameterPropertiesFragment
1. ArtifactSourcePropertiesFragment
1. AttachNewDataDiskOptionsFragment
1. BulkCreationParametersFragment
1. ComputeDataDiskFragment
1. ComputeVMInstanceViewStatusFragment
1. ComputeVMPropertiesFragment
1. CustomImagePropertiesCustomFragment
1. CustomImagePropertiesFragment
1. CustomImagePropertiesFromPlanFragment
1. CustomImagePropertiesFromVMFragment
1. DataDiskPropertiesFragment
1. DataDiskStorageTypeInfoFragment
1. DayDetailsFragment
1. DiskPropertiesFragment
1. EnvironmentDeploymentPropertiesFragment
1. EnvironmentPropertiesFragment
1. EventFragment
1. ExternalSubnetFragment
1. FormulaPropertiesFragment
1. FormulaPropertiesFromVMFragment
1. GalleryImageReferenceFragment
1. HourDetailsFragment
1. InboundNatRuleFragment
1. LabAnnouncementPropertiesFragment
1. LabPropertiesFragment
1. LabSupportPropertiesFragment
1. LabVirtualMachineCreationParameterFragment
1. LabVirtualMachineCreationParameterPropertiesFragment
1. LabVirtualMachinePropertiesFragment
1. LinuxOsInfoFragment
1. NetworkInterfacePropertiesFragment
1. NotificationChannelPropertiesFragment
1. NotificationSettingsFragment
1. PolicyPropertiesFragment
1. PortFragment
1. ScheduleCreationParameterFragment
1. ScheduleCreationParameterPropertiesFragment
1. SchedulePropertiesFragment
1. SecretPropertiesFragment
1. ServiceFabricPropertiesFragment
1. SharedPublicIPAddressConfigurationFragment
1. SubnetFragment
1. SubnetOverrideFragment
1. SubnetSharedPublicIPAddressConfigurationFragment
1. UserIdentityFragment
1. UserPropertiesFragment
1. UserSecretStoreFragment
1. VirtualNetworkPropertiesFragment
1. WeekDetailsFragment
1. WindowsOsInfoFragment

#### Removed Struct Fields

1. ApplicableScheduleFragment.*ApplicableSchedulePropertiesFragment
1. ArtifactSourceFragment.*ArtifactSourcePropertiesFragment
1. CustomImageFragment.*CustomImagePropertiesFragment
1. DiskFragment.*DiskPropertiesFragment
1. EnvironmentFragment.*EnvironmentPropertiesFragment
1. FormulaFragment.*FormulaPropertiesFragment
1. LabFragment.*LabPropertiesFragment
1. LabVirtualMachineCreationParameterProperties.ArtifactDeploymentStatus
1. LabVirtualMachineCreationParameterProperties.ComputeID
1. LabVirtualMachineCreationParameterProperties.CreatedByUser
1. LabVirtualMachineCreationParameterProperties.CreatedByUserID
1. LabVirtualMachineCreationParameterProperties.Fqdn
1. LabVirtualMachineCreationParameterProperties.LastKnownPowerState
1. LabVirtualMachineCreationParameterProperties.OsType
1. LabVirtualMachineCreationParameterProperties.VirtualMachineCreationSource
1. LabVirtualMachineFragment.*LabVirtualMachinePropertiesFragment
1. NotificationChannelFragment.*NotificationChannelPropertiesFragment
1. PolicyFragment.*PolicyPropertiesFragment
1. ScheduleFragment.*SchedulePropertiesFragment
1. SecretFragment.*SecretPropertiesFragment
1. ServiceFabricFragment.*ServiceFabricPropertiesFragment
1. UserFragment.*UserPropertiesFragment
1. VirtualNetworkFragment.*VirtualNetworkPropertiesFragment

### Signature Changes

#### Struct Fields

1. IdentityProperties.Type changed type from *string to ManagedIdentityType

## Additive Changes

### New Constants

1. HTTPStatusCode.Ambiguous
1. HTTPStatusCode.Found
1. HTTPStatusCode.Moved
1. HTTPStatusCode.RedirectKeepVerb
1. HTTPStatusCode.RedirectMethod
1. ManagedIdentityType.ManagedIdentityTypeNone
1. ManagedIdentityType.ManagedIdentityTypeSystemAssigned
1. ManagedIdentityType.ManagedIdentityTypeSystemAssignedUserAssigned
1. ManagedIdentityType.ManagedIdentityTypeUserAssigned
1. SourceControlType.StorageAccount

### New Funcs

1. PossibleManagedIdentityTypeValues() []ManagedIdentityType

### Struct Changes

#### New Structs

1. ServiceRunnerList

#### New Struct Fields

1. DiskProperties.StorageAccountID

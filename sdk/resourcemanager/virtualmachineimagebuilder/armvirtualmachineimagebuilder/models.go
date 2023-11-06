//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import "time"

// DistributeVersioner - Describes how to generate new x.y.z version number for distribution.
type DistributeVersioner struct {
	// REQUIRED; Version numbering scheme to be used.
	Scheme *string
}

// GetDistributeVersioner implements the DistributeVersionerClassification interface for type DistributeVersioner.
func (d *DistributeVersioner) GetDistributeVersioner() *DistributeVersioner { return d }

// DistributeVersionerLatest - Generates version number that will be latest based on existing version numbers.
type DistributeVersionerLatest struct {
	// REQUIRED; Version numbering scheme to be used.
	Scheme *string

	// Major version for the generated version number. Determine what is "latest" based on versions with this value as the major
	// version. -1 is equivalent to leaving it unset.
	Major *int32
}

// GetDistributeVersioner implements the DistributeVersionerClassification interface for type DistributeVersionerLatest.
func (d *DistributeVersionerLatest) GetDistributeVersioner() *DistributeVersioner {
	return &DistributeVersioner{
		Scheme: d.Scheme,
	}
}

// DistributeVersionerSource - Generates version number based on version number of source image
type DistributeVersionerSource struct {
	// REQUIRED; Version numbering scheme to be used.
	Scheme *string
}

// GetDistributeVersioner implements the DistributeVersionerClassification interface for type DistributeVersionerSource.
func (d *DistributeVersionerSource) GetDistributeVersioner() *DistributeVersioner {
	return &DistributeVersioner{
		Scheme: d.Scheme,
	}
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// ImageTemplate - Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
type ImageTemplate struct {
	// REQUIRED; The identity of the image template, if configured.
	Identity *ImageTemplateIdentity

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The properties of the image template
	Properties *ImageTemplateProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ImageTemplateCustomizer - Describes a unit of image customization
type ImageTemplateCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// Friendly Name to provide context on what this customization step does
	Name *string
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateCustomizer.
func (i *ImageTemplateCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer { return i }

// ImageTemplateDistributor - Generic distribution object
type ImageTemplateDistributor struct {
	// REQUIRED; The name to be used for the associated RunOutput.
	RunOutputName *string

	// REQUIRED; Type of distribution.
	Type *string

	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]*string
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateDistributor.
func (i *ImageTemplateDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor { return i }

// ImageTemplateFileCustomizer - Uploads files to VMs (Linux, Windows). Corresponds to Packer file provisioner
type ImageTemplateFileCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// The absolute path to a file (with nested directory structures already created) where the file (from sourceUri) will be
	// uploaded to in the VM
	Destination *string

	// Friendly Name to provide context on what this customization step does
	Name *string

	// SHA256 checksum of the file provided in the sourceUri field above
	SHA256Checksum *string

	// The URI of the file to be uploaded for customizing the VM. It can be a github link, SAS URI for Azure Storage, etc
	SourceURI *string
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateFileCustomizer.
func (i *ImageTemplateFileCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateFileValidator - Uploads files required for validation to VMs (Linux, Windows). Corresponds to Packer file
// provisioner
type ImageTemplateFileValidator struct {
	// REQUIRED; The type of validation you want to use on the Image. For example, "Shell" can be shell validation
	Type *string

	// The absolute path to a file (with nested directory structures already created) where the file (from sourceUri) will be
	// uploaded to in the VM
	Destination *string

	// Friendly Name to provide context on what this validation step does
	Name *string

	// SHA256 checksum of the file provided in the sourceUri field above
	SHA256Checksum *string

	// The URI of the file to be uploaded to the VM for validation. It can be a github link, Azure Storage URI (authorized or
	// SAS), etc
	SourceURI *string
}

// GetImageTemplateInVMValidator implements the ImageTemplateInVMValidatorClassification interface for type ImageTemplateFileValidator.
func (i *ImageTemplateFileValidator) GetImageTemplateInVMValidator() *ImageTemplateInVMValidator {
	return &ImageTemplateInVMValidator{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateIdentity - Identity for the image template.
type ImageTemplateIdentity struct {
	// The type of identity used for the image template. The type 'None' will remove any identities from the image template.
	Type *ResourceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity
}

// ImageTemplateInVMValidator - Describes a unit of in-VM validation of image
type ImageTemplateInVMValidator struct {
	// REQUIRED; The type of validation you want to use on the Image. For example, "Shell" can be shell validation
	Type *string

	// Friendly Name to provide context on what this validation step does
	Name *string
}

// GetImageTemplateInVMValidator implements the ImageTemplateInVMValidatorClassification interface for type ImageTemplateInVMValidator.
func (i *ImageTemplateInVMValidator) GetImageTemplateInVMValidator() *ImageTemplateInVMValidator {
	return i
}

// ImageTemplateLastRunStatus - Describes the latest status of running an image template
type ImageTemplateLastRunStatus struct {
	// End time of the last run (UTC)
	EndTime *time.Time

	// Verbose information about the last run state
	Message *string

	// State of the last run
	RunState *RunState

	// Sub-state of the last run
	RunSubState *RunSubState

	// Start time of the last run (UTC)
	StartTime *time.Time
}

// ImageTemplateListResult - The result of List image templates operation
type ImageTemplateListResult struct {
	// The continuation token.
	NextLink *string

	// An array of image templates
	Value []*ImageTemplate
}

// ImageTemplateManagedImageDistributor - Distribute as a Managed Disk Image.
type ImageTemplateManagedImageDistributor struct {
	// REQUIRED; Resource Id of the Managed Disk Image
	ImageID *string

	// REQUIRED; Azure location for the image, should match if image already exists
	Location *string

	// REQUIRED; The name to be used for the associated RunOutput.
	RunOutputName *string

	// REQUIRED; Type of distribution.
	Type *string

	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]*string
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateManagedImageDistributor.
func (i *ImageTemplateManagedImageDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		ArtifactTags:  i.ArtifactTags,
		RunOutputName: i.RunOutputName,
		Type:          i.Type,
	}
}

// ImageTemplateManagedImageSource - Describes an image source that is a managed image in customer subscription. This image
// must reside in the same subscription and region as the Image Builder template.
type ImageTemplateManagedImageSource struct {
	// REQUIRED; ARM resource id of the managed image in customer subscription
	ImageID *string

	// REQUIRED; Specifies the type of source image you want to start with.
	Type *string
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateManagedImageSource.
func (i *ImageTemplateManagedImageSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// ImageTemplatePlatformImageSource - Describes an image source from Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages].
type ImageTemplatePlatformImageSource struct {
	// REQUIRED; Specifies the type of source image you want to start with.
	Type *string

	// Image offer from the Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages].
	Offer *string

	// Optional configuration of purchase plan for platform image.
	PlanInfo *PlatformImagePurchasePlan

	// Image Publisher in Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages].
	Publisher *string

	// Image sku from the Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages].
	SKU *string

	// Image version from the Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages]. If
	// 'latest' is specified here, the version is evaluated when the image build takes
	// place, not when the template is submitted.
	Version *string

	// READ-ONLY; Image version from the Azure Gallery Images [https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineimages].
	// This readonly field differs from 'version', only if the value specified in
	// 'version' field is 'latest'.
	ExactVersion *string
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplatePlatformImageSource.
func (i *ImageTemplatePlatformImageSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// ImageTemplatePowerShellCustomizer - Runs the specified PowerShell on the VM (Windows). Corresponds to Packer powershell
// provisioner. Exactly one of 'scriptUri' or 'inline' can be specified.
type ImageTemplatePowerShellCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// Array of PowerShell commands to execute
	Inline []*string

	// Friendly Name to provide context on what this customization step does
	Name *string

	// If specified, the PowerShell script will be run with elevated privileges using the Local System user. Can only be true
	// when the runElevated field above is set to true.
	RunAsSystem *bool

	// If specified, the PowerShell script will be run with elevated privileges
	RunElevated *bool

	// SHA256 checksum of the power shell script provided in the scriptUri field above
	SHA256Checksum *string

	// URI of the PowerShell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	ScriptURI *string

	// Valid exit codes for the PowerShell script. [Default: 0]
	ValidExitCodes []*int32
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplatePowerShellCustomizer.
func (i *ImageTemplatePowerShellCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplatePowerShellValidator - Runs the specified PowerShell script during the validation phase (Windows). Corresponds
// to Packer powershell provisioner. Exactly one of 'scriptUri' or 'inline' can be specified.
type ImageTemplatePowerShellValidator struct {
	// REQUIRED; The type of validation you want to use on the Image. For example, "Shell" can be shell validation
	Type *string

	// Array of PowerShell commands to execute
	Inline []*string

	// Friendly Name to provide context on what this validation step does
	Name *string

	// If specified, the PowerShell script will be run with elevated privileges using the Local System user. Can only be true
	// when the runElevated field above is set to true.
	RunAsSystem *bool

	// If specified, the PowerShell script will be run with elevated privileges
	RunElevated *bool

	// SHA256 checksum of the power shell script provided in the scriptUri field above
	SHA256Checksum *string

	// URI of the PowerShell script to be run for validation. It can be a github link, Azure Storage URI, etc
	ScriptURI *string

	// Valid exit codes for the PowerShell script. [Default: 0]
	ValidExitCodes []*int32
}

// GetImageTemplateInVMValidator implements the ImageTemplateInVMValidatorClassification interface for type ImageTemplatePowerShellValidator.
func (i *ImageTemplatePowerShellValidator) GetImageTemplateInVMValidator() *ImageTemplateInVMValidator {
	return &ImageTemplateInVMValidator{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateProperties - Describes the properties of an image template
type ImageTemplateProperties struct {
	// REQUIRED; The distribution targets where the image output needs to go to.
	Distribute []ImageTemplateDistributorClassification

	// REQUIRED; Specifies the properties used to describe the source image.
	Source ImageTemplateSourceClassification

	// Maximum duration to wait while building the image template (includes all customizations, optimization, validations, and
	// distributions). Omit or specify 0 to use the default (4 hours).
	BuildTimeoutInMinutes *int32

	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize []ImageTemplateCustomizerClassification

	// Error handling options upon a build failure
	ErrorHandling *ImageTemplatePropertiesErrorHandling

	// Specifies optimization to be performed on image.
	Optimize *ImageTemplatePropertiesOptimize

	// The staging resource group id in the same subscription as the image template that will be used to build the image. If this
	// field is empty, a resource group with a random name will be created. If the
	// resource group specified in this field doesn't exist, it will be created with the same name. If the resource group specified
	// exists, it must be empty and in the same region as the image template. The
	// resource group created will be deleted during template deletion if this field is empty or the resource group specified
	// doesn't exist, but if the resource group specified exists the resources created
	// in the resource group will be deleted during template deletion and the resource group itself will remain.
	StagingResourceGroup *string

	// Describes how virtual machine is set up to build images
	VMProfile *ImageTemplateVMProfile

	// Configuration options and list of validations to be performed on the resulting image.
	Validate *ImageTemplatePropertiesValidate

	// READ-ONLY; The staging resource group id in the same subscription as the image template that will be used to build the
	// image. This read-only field differs from 'stagingResourceGroup' only if the value specified
	// in the 'stagingResourceGroup' field is empty.
	ExactStagingResourceGroup *string

	// READ-ONLY; State of 'run' that is currently executing or was last executed.
	LastRunStatus *ImageTemplateLastRunStatus

	// READ-ONLY; Provisioning error, if any
	ProvisioningError *ProvisioningError

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ProvisioningState
}

// ImageTemplatePropertiesErrorHandling - Error handling options upon a build failure
type ImageTemplatePropertiesErrorHandling struct {
	// If there is a customizer error and this field is set to 'cleanup', the build VM and associated network resources will be
	// cleaned up. This is the default behavior. If there is a customizer error and
	// this field is set to 'abort', the build VM will be preserved.
	OnCustomizerError *OnBuildError

	// If there is a validation error and this field is set to 'cleanup', the build VM and associated network resources will be
	// cleaned up. This is the default behavior. If there is a validation error and
	// this field is set to 'abort', the build VM will be preserved.
	OnValidationError *OnBuildError
}

// ImageTemplatePropertiesOptimize - Specifies optimization to be performed on image.
type ImageTemplatePropertiesOptimize struct {
	// Optimization is applied on the image for a faster VM boot.
	VMBoot *ImageTemplatePropertiesOptimizeVMBoot
}

// ImageTemplatePropertiesOptimizeVMBoot - Optimization is applied on the image for a faster VM boot.
type ImageTemplatePropertiesOptimizeVMBoot struct {
	// Enabling this field will improve VM boot time by optimizing the final customized image output.
	State *VMBootOptimizationState
}

// ImageTemplatePropertiesValidate - Configuration options and list of validations to be performed on the resulting image.
type ImageTemplatePropertiesValidate struct {
	// If validation fails and this field is set to false, output image(s) will not be distributed. This is the default behavior.
	// If validation fails and this field is set to true, output image(s) will still
	// be distributed. Please use this option with caution as it may result in bad images being distributed for use. In either
	// case (true or false), the end to end image run will be reported as having failed
	// in case of a validation failure. [Note: This field has no effect if validation succeeds.]
	ContinueDistributeOnFailure *bool

	// List of validations to be performed.
	InVMValidations []ImageTemplateInVMValidatorClassification

	// If this field is set to true, the image specified in the 'source' section will directly be validated. No separate build
	// will be run to generate and then validate a customized image.
	SourceValidationOnly *bool
}

// ImageTemplateRestartCustomizer - Reboots a VM and waits for it to come back online (Windows). Corresponds to Packer windows-restart
// provisioner
type ImageTemplateRestartCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// Friendly Name to provide context on what this customization step does
	Name *string

	// Command to check if restart succeeded [Default: '']
	RestartCheckCommand *string

	// Command to execute the restart [Default: 'shutdown /r /f /t 0 /c "packer restart"']
	RestartCommand *string

	// Restart timeout specified as a string of magnitude and unit, e.g. '5m' (5 minutes) or '2h' (2 hours) [Default: '5m']
	RestartTimeout *string
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateRestartCustomizer.
func (i *ImageTemplateRestartCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateSharedImageDistributor - Distribute via Azure Compute Gallery.
type ImageTemplateSharedImageDistributor struct {
	// REQUIRED; Resource Id of the Azure Compute Gallery image
	GalleryImageID *string

	// REQUIRED; The name to be used for the associated RunOutput.
	RunOutputName *string

	// REQUIRED; Type of distribution.
	Type *string

	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]*string

	// Flag that indicates whether created image version should be excluded from latest. Omit to use the default (false).
	ExcludeFromLatest *bool

	// [Deprecated] A list of regions that the image will be replicated to. This list can be specified only if targetRegions is
	// not specified. This field is deprecated - use targetRegions instead.
	ReplicationRegions []*string

	// [Deprecated] Storage account type to be used to store the shared image. Omit to use the default (Standard_LRS). This field
	// can be specified only if replicationRegions is specified. This field is
	// deprecated - use targetRegions instead.
	StorageAccountType *SharedImageStorageAccountType

	// The target regions where the distributed Image Version is going to be replicated to. This object supersedes replicationRegions
	// and can be specified only if replicationRegions is not specified.
	TargetRegions []*TargetRegion

	// Describes how to generate new x.y.z version number for distribution.
	Versioning DistributeVersionerClassification
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateSharedImageDistributor.
func (i *ImageTemplateSharedImageDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		ArtifactTags:  i.ArtifactTags,
		RunOutputName: i.RunOutputName,
		Type:          i.Type,
	}
}

// ImageTemplateSharedImageVersionSource - Describes an image source that is an image version in an Azure Compute Gallery
// or a Direct Shared Gallery.
type ImageTemplateSharedImageVersionSource struct {
	// REQUIRED; ARM resource id of the image version. When image version name is 'latest', the version is evaluated when the
	// image build takes place.
	ImageVersionID *string

	// REQUIRED; Specifies the type of source image you want to start with.
	Type *string

	// READ-ONLY; Exact ARM resource id of the image version. This readonly field differs from the image version Id in 'imageVersionId'
	// only if the version name specified in 'imageVersionId' field is 'latest'.
	ExactVersion *string
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateSharedImageVersionSource.
func (i *ImageTemplateSharedImageVersionSource) GetImageTemplateSource() *ImageTemplateSource {
	return &ImageTemplateSource{
		Type: i.Type,
	}
}

// ImageTemplateShellCustomizer - Runs a shell script during the customization phase (Linux). Corresponds to Packer shell
// provisioner. Exactly one of 'scriptUri' or 'inline' can be specified.
type ImageTemplateShellCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// Array of shell commands to execute
	Inline []*string

	// Friendly Name to provide context on what this customization step does
	Name *string

	// SHA256 checksum of the shell script provided in the scriptUri field
	SHA256Checksum *string

	// URI of the shell script to be run for customizing. It can be a github link, SAS URI for Azure Storage, etc
	ScriptURI *string
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateShellCustomizer.
func (i *ImageTemplateShellCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateShellValidator - Runs the specified shell script during the validation phase (Linux). Corresponds to Packer
// shell provisioner. Exactly one of 'scriptUri' or 'inline' can be specified.
type ImageTemplateShellValidator struct {
	// REQUIRED; The type of validation you want to use on the Image. For example, "Shell" can be shell validation
	Type *string

	// Array of shell commands to execute
	Inline []*string

	// Friendly Name to provide context on what this validation step does
	Name *string

	// SHA256 checksum of the shell script provided in the scriptUri field
	SHA256Checksum *string

	// URI of the shell script to be run for validation. It can be a github link, Azure Storage URI, etc
	ScriptURI *string
}

// GetImageTemplateInVMValidator implements the ImageTemplateInVMValidatorClassification interface for type ImageTemplateShellValidator.
func (i *ImageTemplateShellValidator) GetImageTemplateInVMValidator() *ImageTemplateInVMValidator {
	return &ImageTemplateInVMValidator{
		Name: i.Name,
		Type: i.Type,
	}
}

// ImageTemplateSource - Describes a virtual machine image source for building, customizing and distributing
type ImageTemplateSource struct {
	// REQUIRED; Specifies the type of source image you want to start with.
	Type *string
}

// GetImageTemplateSource implements the ImageTemplateSourceClassification interface for type ImageTemplateSource.
func (i *ImageTemplateSource) GetImageTemplateSource() *ImageTemplateSource { return i }

// ImageTemplateUpdateParameters - Parameters for updating an image template.
type ImageTemplateUpdateParameters struct {
	// The identity of the image template, if configured.
	Identity *ImageTemplateIdentity

	// Parameters for updating an image template.
	Properties *ImageTemplateUpdateParametersProperties

	// The user-specified tags associated with the image template.
	Tags map[string]*string
}

// ImageTemplateUpdateParametersProperties - Parameters for updating an image template.
type ImageTemplateUpdateParametersProperties struct {
	// The distribution targets where the image output needs to go to.
	Distribute []ImageTemplateDistributorClassification
}

// ImageTemplateVMProfile - Describes the virtual machines used to build and validate images
type ImageTemplateVMProfile struct {
	// Size of the OS disk in GB. Omit or specify 0 to use Azure's default OS disk size.
	OSDiskSizeGB *int32

	// Optional array of resource IDs of user assigned managed identities to be configured on the build VM and validation VM.
	// This may include the identity of the image template.
	UserAssignedIdentities []*string

	// Size of the virtual machine used to build, customize and capture images. Omit or specify empty string to use the default
	// (StandardD1v2 for Gen1 images and StandardD2dsv4 for Gen2 images).
	VMSize *string

	// Optional configuration of the virtual network to use to deploy the build VM and validation VM in. Omit if no specific virtual
	// network needs to be used.
	VnetConfig *VirtualNetworkConfig
}

// ImageTemplateVhdDistributor - Distribute via VHD in a storage account.
type ImageTemplateVhdDistributor struct {
	// REQUIRED; The name to be used for the associated RunOutput.
	RunOutputName *string

	// REQUIRED; Type of distribution.
	Type *string

	// Tags that will be applied to the artifact once it has been created/updated by the distributor.
	ArtifactTags map[string]*string

	// Optional Azure Storage URI for the distributed VHD blob. Omit to use the default (empty string) in which case VHD would
	// be published to the storage account in the staging resource group.
	URI *string
}

// GetImageTemplateDistributor implements the ImageTemplateDistributorClassification interface for type ImageTemplateVhdDistributor.
func (i *ImageTemplateVhdDistributor) GetImageTemplateDistributor() *ImageTemplateDistributor {
	return &ImageTemplateDistributor{
		ArtifactTags:  i.ArtifactTags,
		RunOutputName: i.RunOutputName,
		Type:          i.Type,
	}
}

// ImageTemplateWindowsUpdateCustomizer - Installs Windows Updates. Corresponds to Packer Windows Update Provisioner (https://github.com/rgl/packer-provisioner-windows-update)
type ImageTemplateWindowsUpdateCustomizer struct {
	// REQUIRED; The type of customization tool you want to use on the Image. For example, "Shell" can be shell customizer
	Type *string

	// Array of filters to select updates to apply. Omit or specify empty array to use the default (no filter). Refer to above
	// link for examples and detailed description of this field.
	Filters []*string

	// Friendly Name to provide context on what this customization step does
	Name *string

	// Criteria to search updates. Omit or specify empty string to use the default (search all). Refer to above link for examples
	// and detailed description of this field.
	SearchCriteria *string

	// Maximum number of updates to apply at a time. Omit or specify 0 to use the default (1000)
	UpdateLimit *int32
}

// GetImageTemplateCustomizer implements the ImageTemplateCustomizerClassification interface for type ImageTemplateWindowsUpdateCustomizer.
func (i *ImageTemplateWindowsUpdateCustomizer) GetImageTemplateCustomizer() *ImageTemplateCustomizer {
	return &ImageTemplateCustomizer{
		Name: i.Name,
		Type: i.Type,
	}
}

// Operation - A REST API operation
type Operation struct {
	// The object that describes the operation.
	Display *OperationDisplay

	// The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool

	// This is of the format {provider}/{resource}/{operation}
	Name *string

	// The intended executor of the operation.
	Origin *string

	// Properties of the operation.
	Properties any
}

// OperationDisplay - The object that describes the operation.
type OperationDisplay struct {
	// The friendly name of the operation
	Description *string

	// For example: read, write, delete, or listKeys/action
	Operation *string

	// Friendly name of the resource provider.
	Provider *string

	// The resource type on which the operation is performed.
	Resource *string
}

// OperationListResult - Result of the request to list REST API operations. It contains a list of operations and a URL nextLink
// to get the next set of results.
type OperationListResult struct {
	// The URL to get the next set of operation list results if there are any.
	NextLink *string

	// The list of operations supported by the resource provider.
	Value []*Operation
}

// PlatformImagePurchasePlan - Purchase plan configuration for platform image.
type PlatformImagePurchasePlan struct {
	// REQUIRED; Name of the purchase plan.
	PlanName *string

	// REQUIRED; Product of the purchase plan.
	PlanProduct *string

	// REQUIRED; Publisher of the purchase plan.
	PlanPublisher *string
}

// ProvisioningError - Describes the error happened when create or update an image template
type ProvisioningError struct {
	// Verbose error message about the provisioning failure
	Message *string

	// Error code of the provisioning failure
	ProvisioningErrorCode *ProvisioningErrorCode
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RunOutput - Represents an output that was created by running an image template.
type RunOutput struct {
	// The properties of the run output
	Properties *RunOutputProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RunOutputCollection - The result of List run outputs operation
type RunOutputCollection struct {
	// The continuation token.
	NextLink *string

	// An array of run outputs
	Value []*RunOutput
}

// RunOutputProperties - Describes the properties of a run output
type RunOutputProperties struct {
	// The resource id of the artifact.
	ArtifactID *string

	// The location URI of the artifact.
	ArtifactURI *string

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ProvisioningState
}

// SourceImageTriggerProperties - Properties of SourceImage kind of trigger
type SourceImageTriggerProperties struct {
	// REQUIRED; The kind of trigger.
	Kind *string

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ProvisioningState

	// READ-ONLY; Trigger status
	Status *TriggerStatus
}

// GetTriggerProperties implements the TriggerPropertiesClassification interface for type SourceImageTriggerProperties.
func (s *SourceImageTriggerProperties) GetTriggerProperties() *TriggerProperties {
	return &TriggerProperties{
		Kind:              s.Kind,
		ProvisioningState: s.ProvisioningState,
		Status:            s.Status,
	}
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TargetRegion - Describes the target region information.
type TargetRegion struct {
	// REQUIRED; The name of the region.
	Name *string

	// The number of replicas of the Image Version to be created in this region. Omit to use the default (1).
	ReplicaCount *int32

	// Specifies the storage account type to be used to store the image in this region. Omit to use the default (Standard_LRS).
	StorageAccountType *SharedImageStorageAccountType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Trigger - Represents a trigger that can invoke an image template build.
type Trigger struct {
	// The properties of a trigger
	Properties TriggerPropertiesClassification

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TriggerCollection - The result of List triggers operation
type TriggerCollection struct {
	// REQUIRED; An array of triggers
	Value []*Trigger

	// The continuation token.
	NextLink *string
}

// TriggerProperties - Describes the properties of a trigger
type TriggerProperties struct {
	// REQUIRED; The kind of trigger.
	Kind *string

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ProvisioningState

	// READ-ONLY; Trigger status
	Status *TriggerStatus
}

// GetTriggerProperties implements the TriggerPropertiesClassification interface for type TriggerProperties.
func (t *TriggerProperties) GetTriggerProperties() *TriggerProperties { return t }

// TriggerStatus - Describes the status of a trigger
type TriggerStatus struct {
	// READ-ONLY; The status code.
	Code *string

	// READ-ONLY; The detailed status message, including for alerts and error messages.
	Message *string

	// READ-ONLY; The time of the status.
	Time *time.Time
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VirtualNetworkConfig - Virtual Network configuration.
type VirtualNetworkConfig struct {
	// Size of the proxy virtual machine used to pass traffic to the build VM and validation VM. Omit or specify empty string
	// to use the default (StandardA1v2).
	ProxyVMSize *string

	// Resource id of a pre-existing subnet.
	SubnetID *string
}

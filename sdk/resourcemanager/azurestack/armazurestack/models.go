//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

// ActivationKeyResult - The resource containing the Azure Stack activation key.
type ActivationKeyResult struct {
	// Azure Stack activation key.
	ActivationKey *string
}

// CloudManifestFileDeploymentData - Cloud specific manifest data for AzureStack deployment.
type CloudManifestFileDeploymentData struct {
	// Signing verification public key.
	CustomCloudVerificationKey *string

	// Environment endpoints.
	CustomEnvironmentEndpoints *CloudManifestFileEnvironmentEndpoints

	// Dsms external certificates.
	ExternalDsmsCertificates *string
}

// CloudManifestFileEnvironmentEndpoints - Cloud specific environment endpoints for AzureStack deployment.
type CloudManifestFileEnvironmentEndpoints struct {
	// ARM endpoint.
	CustomCloudArmEndpoint *string

	// Dsms endpoint.
	ExternalDsmsEndpoint *string
}

// CloudManifestFileProperties - Cloud specific manifest JSON properties.
type CloudManifestFileProperties struct {
	// Cloud specific manifest data.
	DeploymentData *CloudManifestFileDeploymentData

	// Signature of the cloud specific manifest data.
	Signature *string
}

// CloudManifestFileResponse - Cloud specific manifest GET response.
type CloudManifestFileResponse struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string

	// Cloud specific manifest data.
	Properties *CloudManifestFileProperties

	// READ-ONLY; ID of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of Resource.
	Type *string
}

// Compatibility - Product compatibility
type Compatibility struct {
	// Full error message if any compatibility issues are found
	Description *string

	// Tells if product is compatible with current device
	IsCompatible *bool

	// List of all issues found
	Issues []*CompatibilityIssue

	// Short error message if any compatibility issues are found
	Message *string
}

// CustomerSubscription - Customer subscription.
type CustomerSubscription struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string

	// Customer subscription properties.
	Properties *CustomerSubscriptionProperties

	// READ-ONLY; ID of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of Resource.
	Type *string
}

// CustomerSubscriptionList - Pageable list of customer subscriptions.
type CustomerSubscriptionList struct {
	// URI to the next page.
	NextLink *string

	// List of customer subscriptions.
	Value []*CustomerSubscription
}

// CustomerSubscriptionProperties - Customer subscription properties.
type CustomerSubscriptionProperties struct {
	// Tenant Id.
	TenantID *string
}

// DataDiskImage - Data disk image.
type DataDiskImage struct {
	// READ-ONLY; The LUN.
	Lun *int32

	// READ-ONLY; SAS key for source blob.
	SourceBlobSasURI *string
}

// DeploymentLicenseRequest - Request details for generating a deployment license.
type DeploymentLicenseRequest struct {
	// Signing verification public key version.
	VerificationVersion *string
}

// DeploymentLicenseResponse - A license that can be used to deploy an Azure Stack device.
type DeploymentLicenseResponse struct {
	// Signature of the license chain.
	Signature *string

	// A license chain that can be used to temporarily activate an Azure Stack device.
	TemporaryLicenseChain []*string
}

// DeviceConfiguration - Device Configuration.
type DeviceConfiguration struct {
	// READ-ONLY; Version of the device.
	DeviceVersion *string

	// READ-ONLY; Identity system of the device.
	IdentitySystem *Category
}

// Display - Contains the localized display information for this particular operation or action.
type Display struct {
	// The localized, friendly description for the operation. The description will be displayed to the user. It should be thorough
	// and concise for used in both tooltips and detailed views.
	Description *string

	// The localized, friendly name for the operation. Use the name as it will displayed to the user.
	Operation *string

	// The localized, friendly version of the resource provider name.
	Provider *string

	// The localized, friendly version of the resource type related to this action or operation; the resource type should match
	// the public documentation for the resource provider.
	Resource *string
}

// ExtendedProduct - Extended description about the product required for installing it into Azure Stack.
type ExtendedProduct struct {
	// READ-ONLY; The URI to the .azpkg file that provides information required for showing product in the gallery.
	GalleryPackageBlobSasURI *string

	// READ-ONLY; Specifies the kind of the product (virtualMachine or virtualMachineExtension).
	ProductKind *string

	// READ-ONLY; Specifies additional properties describing the product.
	Properties *ExtendedProductProperties
}

// ExtendedProductProperties - Product information.
type ExtendedProductProperties struct {
	// READ-ONLY; Specifies kind of compute role included in the package.
	ComputeRole *ComputeRole

	// READ-ONLY; List of attached data disks.
	DataDiskImages []*DataDiskImage

	// READ-ONLY; Specifies if product is a Virtual Machine Extension.
	IsSystemExtension *bool

	// READ-ONLY; OS disk image used by product.
	OSDiskImage *OsDiskImage

	// READ-ONLY; Specifies a download location where content can be downloaded from.
	SourceBlob *URI

	// READ-ONLY; Indicates if specified product supports multiple extensions.
	SupportMultipleExtensions *bool

	// READ-ONLY; Specifies operating system used by the product.
	VMOsType *OperatingSystem

	// READ-ONLY; Indicates if virtual machine Scale Set is enabled in the specified product.
	VMScaleSetEnabled *bool

	// READ-ONLY; Specifies product version.
	Version *string
}

// IconUris - Links to product icons.
type IconUris struct {
	// URI to hero icon.
	Hero *string

	// URI to large icon.
	Large *string

	// URI to medium icon.
	Medium *string

	// URI to small icon.
	Small *string

	// URI to wide icon.
	Wide *string
}

// MarketplaceProductLogUpdate - Update details for product log.
type MarketplaceProductLogUpdate struct {
	// READ-ONLY; Error details related to operation.
	Details *string

	// READ-ONLY; Error related to the operation.
	Error *string

	// READ-ONLY; Operation to log.
	Operation *string

	// READ-ONLY; Operation status to log.
	Status *string
}

// Operation - Describes the supported REST operation.
type Operation struct {
	// Contains the localized display information for this particular operation or action.
	Display *Display

	// The name of the operation being performed on this particular object.
	Name *string

	// The intended executor of the operation.
	Origin *string
}

// OperationList - List of Operations
type OperationList struct {
	// URI to the next page of operations.
	NextLink *string

	// Array of operations
	Value []*Operation
}

// OsDiskImage - OS disk image.
type OsDiskImage struct {
	// READ-ONLY; OS operating system type.
	OperatingSystem *OperatingSystem

	// READ-ONLY; SAS key for source blob.
	SourceBlobSasURI *string
}

// Product information.
type Product struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string

	// Properties of the product resource.
	Properties *ProductNestedProperties

	// READ-ONLY; ID of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of Resource.
	Type *string
}

// ProductLink - Link with additional information about a product.
type ProductLink struct {
	// The description of the link.
	DisplayName *string

	// The URI corresponding to the link.
	URI *string
}

// ProductList - Pageable list of products.
type ProductList struct {
	// URI to the next page.
	NextLink *string

	// List of products.
	Value []*Product
}

// ProductLog - Product action log.
type ProductLog struct {
	// READ-ONLY; Operation error details.
	Details *string

	// READ-ONLY; Operation end datetime.
	EndDate *string

	// READ-ONLY; Operation error data.
	Error *string

	// READ-ONLY; Log ID.
	ID *string

	// READ-ONLY; Logged operation.
	Operation *string

	// READ-ONLY; Logged product ID.
	ProductID *string

	// READ-ONLY; Logged registration name.
	RegistrationName *string

	// READ-ONLY; Logged resource group name.
	ResourceGroupName *string

	// READ-ONLY; Operation start datetime.
	StartDate *string

	// READ-ONLY; Operation status.
	Status *string

	// READ-ONLY; Logged subscription ID.
	SubscriptionID *string
}

// ProductNestedProperties - Properties portion of the product resource.
type ProductNestedProperties struct {
	// The part number used for billing purposes.
	BillingPartNumber *string

	// Product compatibility with current device.
	Compatibility *Compatibility

	// The description of the product.
	Description *string

	// The display name of the product.
	DisplayName *string

	// The identifier of the gallery item corresponding to the product.
	GalleryItemIdentity *string

	// Additional links available for this product.
	IconUris *IconUris

	// The legal terms.
	LegalTerms *string

	// Additional links available for this product.
	Links []*ProductLink

	// The offer representing the product.
	Offer *string

	// The version of the product offer.
	OfferVersion *string

	// The length of product content.
	PayloadLength *int64

	// The privacy policy.
	PrivacyPolicy *string

	// The kind of the product (virtualMachine or virtualMachineExtension)
	ProductKind *string

	// Additional properties for the product.
	ProductProperties *ProductProperties

	// The user-friendly name of the product publisher.
	PublisherDisplayName *string

	// Publisher identifier.
	PublisherIdentifier *string

	// The product SKU.
	SKU *string

	// The type of the Virtual Machine Extension.
	VMExtensionType *string
}

// ProductProperties - Additional properties of the product
type ProductProperties struct {
	// The version.
	Version *string
}

// Registration information.
type Registration struct {
	// REQUIRED; Location of the resource.
	Location *Location

	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string

	// Registration resource.
	Properties *RegistrationProperties

	// Custom tags for the resource.
	Tags map[string]*string

	// READ-ONLY; ID of the resource.
	ID *string

	// READ-ONLY; Name of the resource.
	Name *string

	// READ-ONLY; Type of Resource.
	Type *string
}

// RegistrationList - Pageable list of registrations.
type RegistrationList struct {
	// URI to the next page.
	NextLink *string

	// List of Registrations
	Value []*Registration
}

// RegistrationParameter - Registration resource
type RegistrationParameter struct {
	// REQUIRED; Location of the resource.
	Location *Location

	// REQUIRED; Properties of the Azure Stack registration resource
	Properties *RegistrationParameterProperties
}

// RegistrationParameterProperties - Properties of the Azure Stack registration resource
type RegistrationParameterProperties struct {
	// REQUIRED; The token identifying registered Azure Stack
	RegistrationToken *string
}

// RegistrationProperties - Properties portion of the registration resource.
type RegistrationProperties struct {
	// Specifies the billing mode for the Azure Stack registration.
	BillingModel *string

	// The identifier of the registered Azure Stack.
	CloudID *string

	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectID *string
}

// URI - The URI.
type URI struct {
	// READ-ONLY; The URI.
	URI *string
}

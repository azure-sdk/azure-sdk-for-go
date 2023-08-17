//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armazuresphere

// CatalogsClientBeginCreateOrUpdateOptions contains the optional parameters for the CatalogsClient.BeginCreateOrUpdate method.
type CatalogsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientBeginDeleteOptions contains the optional parameters for the CatalogsClient.BeginDelete method.
type CatalogsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientCountDevicesOptions contains the optional parameters for the CatalogsClient.CountDevices method.
type CatalogsClientCountDevicesOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientGetOptions contains the optional parameters for the CatalogsClient.Get method.
type CatalogsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientListByResourceGroupOptions contains the optional parameters for the CatalogsClient.NewListByResourceGroupPager
// method.
type CatalogsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientListBySubscriptionOptions contains the optional parameters for the CatalogsClient.NewListBySubscriptionPager
// method.
type CatalogsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientListDeploymentsOptions contains the optional parameters for the CatalogsClient.NewListDeploymentsPager method.
type CatalogsClientListDeploymentsOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// CatalogsClientListDeviceGroupsOptions contains the optional parameters for the CatalogsClient.NewListDeviceGroupsPager
// method.
type CatalogsClientListDeviceGroupsOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// CatalogsClientListDeviceInsightsOptions contains the optional parameters for the CatalogsClient.NewListDeviceInsightsPager
// method.
type CatalogsClientListDeviceInsightsOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// CatalogsClientListDevicesOptions contains the optional parameters for the CatalogsClient.NewListDevicesPager method.
type CatalogsClientListDevicesOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// CatalogsClientUpdateOptions contains the optional parameters for the CatalogsClient.Update method.
type CatalogsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
type CertificatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CertificatesClientListByCatalogOptions contains the optional parameters for the CertificatesClient.NewListByCatalogPager
// method.
type CertificatesClientListByCatalogOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// CertificatesClientRetrieveCertChainOptions contains the optional parameters for the CertificatesClient.RetrieveCertChain
// method.
type CertificatesClientRetrieveCertChainOptions struct {
	// placeholder for future optional parameters
}

// CertificatesClientRetrieveProofOfPossessionNonceOptions contains the optional parameters for the CertificatesClient.RetrieveProofOfPossessionNonce
// method.
type CertificatesClientRetrieveProofOfPossessionNonceOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeploymentsClient.BeginCreateOrUpdate
// method.
type DeploymentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientBeginDeleteOptions contains the optional parameters for the DeploymentsClient.BeginDelete method.
type DeploymentsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientGetOptions contains the optional parameters for the DeploymentsClient.Get method.
type DeploymentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientListByDeviceGroupOptions contains the optional parameters for the DeploymentsClient.NewListByDeviceGroupPager
// method.
type DeploymentsClientListByDeviceGroupOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// DeviceGroupsClientBeginClaimDevicesOptions contains the optional parameters for the DeviceGroupsClient.BeginClaimDevices
// method.
type DeviceGroupsClientBeginClaimDevicesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeviceGroupsClient.BeginCreateOrUpdate
// method.
type DeviceGroupsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceGroupsClientBeginDeleteOptions contains the optional parameters for the DeviceGroupsClient.BeginDelete method.
type DeviceGroupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceGroupsClientBeginUpdateOptions contains the optional parameters for the DeviceGroupsClient.BeginUpdate method.
type DeviceGroupsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeviceGroupsClientCountDevicesOptions contains the optional parameters for the DeviceGroupsClient.CountDevices method.
type DeviceGroupsClientCountDevicesOptions struct {
	// placeholder for future optional parameters
}

// DeviceGroupsClientGetOptions contains the optional parameters for the DeviceGroupsClient.Get method.
type DeviceGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DeviceGroupsClientListByProductOptions contains the optional parameters for the DeviceGroupsClient.NewListByProductPager
// method.
type DeviceGroupsClientListByProductOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// DevicesClientBeginCreateOrUpdateOptions contains the optional parameters for the DevicesClient.BeginCreateOrUpdate method.
type DevicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginDeleteOptions contains the optional parameters for the DevicesClient.BeginDelete method.
type DevicesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginGenerateCapabilityImageOptions contains the optional parameters for the DevicesClient.BeginGenerateCapabilityImage
// method.
type DevicesClientBeginGenerateCapabilityImageOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientBeginUpdateOptions contains the optional parameters for the DevicesClient.BeginUpdate method.
type DevicesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
type DevicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListByDeviceGroupOptions contains the optional parameters for the DevicesClient.NewListByDeviceGroupPager
// method.
type DevicesClientListByDeviceGroupOptions struct {
	// placeholder for future optional parameters
}

// ImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the ImagesClient.BeginCreateOrUpdate method.
type ImagesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ImagesClientBeginDeleteOptions contains the optional parameters for the ImagesClient.BeginDelete method.
type ImagesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ImagesClientGetOptions contains the optional parameters for the ImagesClient.Get method.
type ImagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ImagesClientListByCatalogOptions contains the optional parameters for the ImagesClient.NewListByCatalogPager method.
type ImagesClientListByCatalogOptions struct {
	// Filter the result list using the given expression
	Filter *string

	// The maximum number of result items per page.
	Maxpagesize *int32

	// The number of result items to skip.
	Skip *int32

	// The number of result items to return.
	Top *int32
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProductsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProductsClient.BeginCreateOrUpdate method.
type ProductsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProductsClientBeginDeleteOptions contains the optional parameters for the ProductsClient.BeginDelete method.
type ProductsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProductsClientBeginUpdateOptions contains the optional parameters for the ProductsClient.BeginUpdate method.
type ProductsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProductsClientCountDevicesOptions contains the optional parameters for the ProductsClient.CountDevices method.
type ProductsClientCountDevicesOptions struct {
	// placeholder for future optional parameters
}

// ProductsClientGenerateDefaultDeviceGroupsOptions contains the optional parameters for the ProductsClient.NewGenerateDefaultDeviceGroupsPager
// method.
type ProductsClientGenerateDefaultDeviceGroupsOptions struct {
	// placeholder for future optional parameters
}

// ProductsClientGetOptions contains the optional parameters for the ProductsClient.Get method.
type ProductsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProductsClientListByCatalogOptions contains the optional parameters for the ProductsClient.NewListByCatalogPager method.
type ProductsClientListByCatalogOptions struct {
	// placeholder for future optional parameters
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewCommunityGalleriesClient() *CommunityGalleriesClient {
	subClient, _ := NewCommunityGalleriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCommunityGalleryImagesClient() *CommunityGalleryImagesClient {
	subClient, _ := NewCommunityGalleryImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCommunityGalleryImageVersionsClient() *CommunityGalleryImageVersionsClient {
	subClient, _ := NewCommunityGalleryImageVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGalleriesClient() *GalleriesClient {
	subClient, _ := NewGalleriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGalleryImagesClient() *GalleryImagesClient {
	subClient, _ := NewGalleryImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGalleryImageVersionsClient() *GalleryImageVersionsClient {
	subClient, _ := NewGalleryImageVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGalleryApplicationsClient() *GalleryApplicationsClient {
	subClient, _ := NewGalleryApplicationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGalleryApplicationVersionsClient() *GalleryApplicationVersionsClient {
	subClient, _ := NewGalleryApplicationVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGallerySharingProfileClient() *GallerySharingProfileClient {
	subClient, _ := NewGallerySharingProfileClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSharedGalleriesClient() *SharedGalleriesClient {
	subClient, _ := NewSharedGalleriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSharedGalleryImagesClient() *SharedGalleryImagesClient {
	subClient, _ := NewSharedGalleryImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSharedGalleryImageVersionsClient() *SharedGalleryImageVersionsClient {
	subClient, _ := NewSharedGalleryImageVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/CommunityGallery_Create.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createACommunityGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
			SharingProfile: &armcompute.SharingProfile{
				CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
					Eula:             to.Ptr("eula"),
					PublicNamePrefix: to.Ptr("PirPublic"),
					PublisherContact: to.Ptr("pir@microsoft.com"),
					PublisherURI:     to.Ptr("uri"),
				},
				Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
	// 				Eula: to.Ptr("eula"),
	// 				PublicNamePrefix: to.Ptr("PirPublic"),
	// 				PublisherContact: to.Ptr("pir@microsoft.com"),
	// 				PublisherURI: to.Ptr("uri"),
	// 			},
	// 			Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Create_WithSharingProfile.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryWithSharingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
			SharingProfile: &armcompute.SharingProfile{
				Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesGroups),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesGroups),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Create_SoftDeletionEnabled.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryWithSoftDeletionEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
			SoftDeletePolicy: &armcompute.SoftDeletePolicy{
				IsSoftDeleteEnabled: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SoftDeletePolicy: &armcompute.SoftDeletePolicy{
	// 			IsSoftDeleteEnabled: to.Ptr(true),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Create.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createOrUpdateASimpleGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Update.json
func ExampleGalleriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.GalleryUpdate{
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/CommunityGallery_Get.json
func ExampleGalleriesClient_Get_getACommunityGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: nil,
		Expand: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/communityGalleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
	// 				CommunityGalleryEnabled: to.Ptr(true),
	// 				Eula: to.Ptr("eula"),
	// 				PublicNames: []*string{
	// 					to.Ptr("GalleryPublicName")},
	// 					PublisherContact: to.Ptr("pir@microsoft.com"),
	// 					PublisherURI: to.Ptr("uri"),
	// 				},
	// 				Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
	// 			},
	// 			SharingStatus: &armcompute.SharingStatus{
	// 				AggregatedState: to.Ptr(armcompute.SharingStateSucceeded),
	// 				Summary: []*armcompute.RegionalSharingStatus{
	// 					{
	// 						Region: to.Ptr("westus"),
	// 						State: to.Ptr(armcompute.SharingStateSucceeded),
	// 						Details: to.Ptr(""),
	// 				}},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Get_WithExpandSharingProfileGroups.json
func ExampleGalleriesClient_Get_getAGalleryWithExpandSharingProfileGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: nil,
		Expand: to.Ptr(armcompute.GalleryExpandParamsSharingProfileGroups),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			Groups: []*armcompute.SharingProfileGroup{
	// 				{
	// 					Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
	// 					IDs: []*string{
	// 						to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
	// 						to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
	// 					},
	// 					{
	// 						Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
	// 						IDs: []*string{
	// 							to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
	// 					}},
	// 					Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesGroups),
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Get_WithSelectPermissions.json
func ExampleGalleriesClient_Get_getAGalleryWithSelectPermissions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: to.Ptr(armcompute.SelectPermissionsPermissions),
		Expand: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			Groups: []*armcompute.SharingProfileGroup{
	// 				{
	// 					Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
	// 					IDs: []*string{
	// 						to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
	// 						to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
	// 					},
	// 					{
	// 						Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
	// 						IDs: []*string{
	// 							to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
	// 					}},
	// 					Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesGroups),
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Get.json
func ExampleGalleriesClient_Get_getAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: nil,
		Expand: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_Delete.json
func ExampleGalleriesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginDelete(ctx, "myResourceGroup", "myGalleryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_ListByResourceGroup.json
func ExampleGalleriesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleriesClient().NewListByResourceGroupPager("myResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.GalleryList = armcompute.GalleryList{
		// 	Value: []*armcompute.Gallery{
		// 		{
		// 			Name: to.Ptr("myGalleryName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryProperties{
		// 				Description: to.Ptr("This is the gallery description."),
		// 				Identifier: &armcompute.GalleryIdentifier{
		// 					UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/Gallery_ListBySubscription.json
func ExampleGalleriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleriesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.GalleryList = armcompute.GalleryList{
		// 	Value: []*armcompute.Gallery{
		// 		{
		// 			Name: to.Ptr("myGalleryName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryProperties{
		// 				Description: to.Ptr("This is the gallery description."),
		// 				Identifier: &armcompute.GalleryIdentifier{
		// 					UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
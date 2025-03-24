// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

// ClientCreateOrUpdateOptions contains the optional parameters for the Client.CreateOrUpdate method.
type ClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteOptions contains the optional parameters for the Client.Delete method.
type ClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ClientGetBuiltInOptions contains the optional parameters for the Client.GetBuiltIn method.
type ClientGetBuiltInOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientListBuiltInsOptions contains the optional parameters for the Client.NewListBuiltInsPager method.
type ClientListBuiltInsOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientListByResourceGroupOptions contains the optional parameters for the Client.NewListByResourceGroupPager method.
type ClientListByResourceGroupOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientListBySubscriptionOptions contains the optional parameters for the Client.NewListBySubscriptionPager method.
type ClientListBySubscriptionOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientUpdateOptions contains the optional parameters for the Client.Update method.
type ClientUpdateOptions struct {
	// Template Spec resource with the tags to be updated.
	TemplateSpec *TemplateSpecUpdateModel
}

// TemplateSpecVersionsClientCreateOrUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.CreateOrUpdate
// method.
type TemplateSpecVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientDeleteOptions contains the optional parameters for the TemplateSpecVersionsClient.Delete method.
type TemplateSpecVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientGetBuiltInOptions contains the optional parameters for the TemplateSpecVersionsClient.GetBuiltIn
// method.
type TemplateSpecVersionsClientGetBuiltInOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientGetOptions contains the optional parameters for the TemplateSpecVersionsClient.Get method.
type TemplateSpecVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientListBuiltInsOptions contains the optional parameters for the TemplateSpecVersionsClient.NewListBuiltInsPager
// method.
type TemplateSpecVersionsClientListBuiltInsOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientListOptions contains the optional parameters for the TemplateSpecVersionsClient.NewListPager
// method.
type TemplateSpecVersionsClientListOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.Update method.
type TemplateSpecVersionsClientUpdateOptions struct {
	// Template Spec Version resource with the tags to be updated.
	TemplateSpecVersionUpdateModel *TemplateSpecVersionUpdateModel
}

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package azmanager

// CreateOrUpdateWidgetResponse contains the response from method WidgetManagerWidgetsClient.BeginCreateOrUpdateWidget.
type CreateOrUpdateWidgetResponse struct {
	// A widget.
	WidgetSuite
}

// DeleteWidgetResponse contains the response from method WidgetManagerWidgetsClient.BeginDeleteWidget.
type DeleteWidgetResponse struct {
	// A widget.
	WidgetSuite
}

// GetWidgetOperationStatusResponse contains the response from method WidgetManagerWidgetsClient.GetWidgetOperationStatus.
type GetWidgetOperationStatusResponse struct {
	// Provides status details for long running operations.
	ResourceOperationStatusWidgetSuiteWidgetSuiteError
}

// GetWidgetResponse contains the response from method WidgetManagerWidgetsClient.GetWidget.
type GetWidgetResponse struct {
	// A widget.
	WidgetSuite
}

// ListWidgetsResponse contains the response from method WidgetManagerWidgetsClient.NewListWidgetsPager.
type ListWidgetsResponse struct {
	// Paged collection of WidgetSuite items
	PagedWidgetSuite
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginCreateOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginCreate
// method.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginDeleteOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginDelete
// method.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientBeginDeleteOptions struct {
	// Optional force delete boolean. If this is provided as true, migration will be deleted even if active.
	Force *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbRUMongoClient.NewGetForScopePager
// method.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeOptions struct {
	// placeholder for future optional parameters
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientGetOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbRUMongoClient.Get
// method.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginCreate
// method.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginDelete
// method.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientBeginDeleteOptions struct {
	// Optional force delete boolean. If this is provided as true, migration will be deleted even if active.
	Force *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.NewGetForScopePager
// method.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeOptions struct {
	// placeholder for future optional parameters
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetOptions contains the optional parameters for the DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.Get
// method.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabaseMigrationsSQLDbClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginCancel
// method.
type DatabaseMigrationsSQLDbClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginCreateOrUpdate
// method.
type DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLDbClientBeginDeleteOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginDelete
// method.
type DatabaseMigrationsSQLDbClientBeginDeleteOptions struct {
	// Optional force delete boolean. If this is provided as true, migration will be deleted even if active.
	Force *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLDbClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.Get method.
type DatabaseMigrationsSQLDbClientGetOptions struct {
	// Complete migration details be included in the response.
	Expand *string

	// Optional migration operation ID. If this is provided, then details of migration operation for that ID are retrieved. If
	// not provided (default), then details related to most recent or current operation
	// are retrieved.
	MigrationOperationID *string
}

// DatabaseMigrationsSQLMiClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCancel
// method.
type DatabaseMigrationsSQLMiClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate
// method.
type DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLMiClientBeginCutoverOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCutover
// method.
type DatabaseMigrationsSQLMiClientBeginCutoverOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLMiClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.Get method.
type DatabaseMigrationsSQLMiClientGetOptions struct {
	// Complete migration details be included in the response.
	Expand *string

	// Optional migration operation ID. If this is provided, then details of migration operation for that ID are retrieved. If
	// not provided (default), then details related to most recent or current operation
	// are retrieved.
	MigrationOperationID *string
}

// DatabaseMigrationsSQLVMClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLVMClient.BeginCancel
// method.
type DatabaseMigrationsSQLVMClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLVMClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLVMClient.BeginCreateOrUpdate
// method.
type DatabaseMigrationsSQLVMClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLVMClientBeginCutoverOptions contains the optional parameters for the DatabaseMigrationsSQLVMClient.BeginCutover
// method.
type DatabaseMigrationsSQLVMClientBeginCutoverOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabaseMigrationsSQLVMClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLVMClient.Get method.
type DatabaseMigrationsSQLVMClientGetOptions struct {
	// Complete migration details be included in the response.
	Expand *string

	// Optional migration operation ID. If this is provided, then details of migration operation for that ID are retrieved. If
	// not provided (default), then details related to most recent or current operation
	// are retrieved.
	MigrationOperationID *string
}

// FilesClientCreateOrUpdateOptions contains the optional parameters for the FilesClient.CreateOrUpdate method.
type FilesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// FilesClientDeleteOptions contains the optional parameters for the FilesClient.Delete method.
type FilesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FilesClientGetOptions contains the optional parameters for the FilesClient.Get method.
type FilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FilesClientListOptions contains the optional parameters for the FilesClient.NewListPager method.
type FilesClientListOptions struct {
	// placeholder for future optional parameters
}

// FilesClientReadOptions contains the optional parameters for the FilesClient.Read method.
type FilesClientReadOptions struct {
	// placeholder for future optional parameters
}

// FilesClientReadWriteOptions contains the optional parameters for the FilesClient.ReadWrite method.
type FilesClientReadWriteOptions struct {
	// placeholder for future optional parameters
}

// FilesClientUpdateOptions contains the optional parameters for the FilesClient.Update method.
type FilesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// MigrationServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the MigrationServicesClient.BeginCreateOrUpdate
// method.
type MigrationServicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MigrationServicesClientBeginDeleteOptions contains the optional parameters for the MigrationServicesClient.BeginDelete
// method.
type MigrationServicesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MigrationServicesClientBeginUpdateOptions contains the optional parameters for the MigrationServicesClient.BeginUpdate
// method.
type MigrationServicesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MigrationServicesClientGetOptions contains the optional parameters for the MigrationServicesClient.Get method.
type MigrationServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// MigrationServicesClientListByResourceGroupOptions contains the optional parameters for the MigrationServicesClient.NewListByResourceGroupPager
// method.
type MigrationServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MigrationServicesClientListBySubscriptionOptions contains the optional parameters for the MigrationServicesClient.NewListBySubscriptionPager
// method.
type MigrationServicesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MigrationServicesClientListMigrationsOptions contains the optional parameters for the MigrationServicesClient.NewListMigrationsPager
// method.
type MigrationServicesClientListMigrationsOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.CreateOrUpdate method.
type ProjectsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientDeleteOptions contains the optional parameters for the ProjectsClient.Delete method.
type ProjectsClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
type ProjectsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientListOptions contains the optional parameters for the ProjectsClient.NewListPager method.
type ProjectsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
type ProjectsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ResourceSKUsClientListSKUsOptions contains the optional parameters for the ResourceSKUsClient.NewListSKUsPager method.
type ResourceSKUsClientListSKUsOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLMigrationServicesClient.BeginCreateOrUpdate
// method.
type SQLMigrationServicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLMigrationServicesClientBeginDeleteOptions contains the optional parameters for the SQLMigrationServicesClient.BeginDelete
// method.
type SQLMigrationServicesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLMigrationServicesClientBeginUpdateOptions contains the optional parameters for the SQLMigrationServicesClient.BeginUpdate
// method.
type SQLMigrationServicesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SQLMigrationServicesClientDeleteNodeOptions contains the optional parameters for the SQLMigrationServicesClient.DeleteNode
// method.
type SQLMigrationServicesClientDeleteNodeOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientGetOptions contains the optional parameters for the SQLMigrationServicesClient.Get method.
type SQLMigrationServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientListAuthKeysOptions contains the optional parameters for the SQLMigrationServicesClient.ListAuthKeys
// method.
type SQLMigrationServicesClientListAuthKeysOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientListByResourceGroupOptions contains the optional parameters for the SQLMigrationServicesClient.NewListByResourceGroupPager
// method.
type SQLMigrationServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientListBySubscriptionOptions contains the optional parameters for the SQLMigrationServicesClient.NewListBySubscriptionPager
// method.
type SQLMigrationServicesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientListMigrationsOptions contains the optional parameters for the SQLMigrationServicesClient.NewListMigrationsPager
// method.
type SQLMigrationServicesClientListMigrationsOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientListMonitoringDataOptions contains the optional parameters for the SQLMigrationServicesClient.ListMonitoringData
// method.
type SQLMigrationServicesClientListMonitoringDataOptions struct {
	// placeholder for future optional parameters
}

// SQLMigrationServicesClientRegenerateAuthKeysOptions contains the optional parameters for the SQLMigrationServicesClient.RegenerateAuthKeys
// method.
type SQLMigrationServicesClientRegenerateAuthKeysOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientCancelOptions contains the optional parameters for the ServiceTasksClient.Cancel method.
type ServiceTasksClientCancelOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientCreateOrUpdateOptions contains the optional parameters for the ServiceTasksClient.CreateOrUpdate method.
type ServiceTasksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ServiceTasksClientDeleteOptions contains the optional parameters for the ServiceTasksClient.Delete method.
type ServiceTasksClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// ServiceTasksClientGetOptions contains the optional parameters for the ServiceTasksClient.Get method.
type ServiceTasksClientGetOptions struct {
	// Expand the response
	Expand *string
}

// ServiceTasksClientListOptions contains the optional parameters for the ServiceTasksClient.NewListPager method.
type ServiceTasksClientListOptions struct {
	// Filter tasks by task type
	TaskType *string
}

// ServiceTasksClientUpdateOptions contains the optional parameters for the ServiceTasksClient.Update method.
type ServiceTasksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate method.
type ServicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
type ServicesClientBeginDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginStartOptions contains the optional parameters for the ServicesClient.BeginStart method.
type ServicesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginStopOptions contains the optional parameters for the ServicesClient.BeginStop method.
type ServicesClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientBeginUpdateOptions contains the optional parameters for the ServicesClient.BeginUpdate method.
type ServicesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientCheckChildrenNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckChildrenNameAvailability
// method.
type ServicesClientCheckChildrenNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
type ServicesClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCheckStatusOptions contains the optional parameters for the ServicesClient.CheckStatus method.
type ServicesClientCheckStatusOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
// method.
type ServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
type ServicesClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListSKUsOptions contains the optional parameters for the ServicesClient.NewListSKUsPager method.
type ServicesClientListSKUsOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCancelOptions contains the optional parameters for the TasksClient.Cancel method.
type TasksClientCancelOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCommandOptions contains the optional parameters for the TasksClient.Command method.
type TasksClientCommandOptions struct {
	// placeholder for future optional parameters
}

// TasksClientCreateOrUpdateOptions contains the optional parameters for the TasksClient.CreateOrUpdate method.
type TasksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TasksClientDeleteOptions contains the optional parameters for the TasksClient.Delete method.
type TasksClientDeleteOptions struct {
	// Delete the resource even if it contains running tasks
	DeleteRunningTasks *bool
}

// TasksClientGetOptions contains the optional parameters for the TasksClient.Get method.
type TasksClientGetOptions struct {
	// Expand the response
	Expand *string
}

// TasksClientListOptions contains the optional parameters for the TasksClient.NewListPager method.
type TasksClientListOptions struct {
	// Filter tasks by task type
	TaskType *string
}

// TasksClientUpdateOptions contains the optional parameters for the TasksClient.Update method.
type TasksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
type UsagesClientListOptions struct {
	// placeholder for future optional parameters
}

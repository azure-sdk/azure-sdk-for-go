//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

// BackupPolicyClassification provides polymorphic access to related types.
// Call the interface's GetBackupPolicy() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BackupPolicy, *ContinuousModeBackupPolicy, *PeriodicModeBackupPolicy
type BackupPolicyClassification interface {
	// GetBackupPolicy returns the BackupPolicy content of the underlying type.
	GetBackupPolicy() *BackupPolicy
}

// BaseCosmosDataTransferDataSourceSinkClassification provides polymorphic access to related types.
// Call the interface's GetBaseCosmosDataTransferDataSourceSink() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BaseCosmosDataTransferDataSourceSink, *CassandraDataTransferDataSourceSink, *MongoDataTransferDataSourceSink, *SQLDataTransferDataSourceSink
type BaseCosmosDataTransferDataSourceSinkClassification interface {
	DataTransferDataSourceSinkClassification
	// GetBaseCosmosDataTransferDataSourceSink returns the BaseCosmosDataTransferDataSourceSink content of the underlying type.
	GetBaseCosmosDataTransferDataSourceSink() *BaseCosmosDataTransferDataSourceSink
}

// DataTransferDataSourceSinkClassification provides polymorphic access to related types.
// Call the interface's GetDataTransferDataSourceSink() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureBlobDataTransferDataSourceSink, *BaseCosmosDataTransferDataSourceSink, *CassandraDataTransferDataSourceSink, *DataTransferDataSourceSink,
// - *MongoDataTransferDataSourceSink, *SQLDataTransferDataSourceSink
type DataTransferDataSourceSinkClassification interface {
	// GetDataTransferDataSourceSink returns the DataTransferDataSourceSink content of the underlying type.
	GetDataTransferDataSourceSink() *DataTransferDataSourceSink
}

// ServiceResourceCreateUpdatePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetServiceResourceCreateUpdateProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataTransferServiceResourceCreateUpdateParameters, *GraphAPIComputeServiceResourceCreateUpdateParameters, *MaterializedViewsBuilderServiceResourceCreateUpdateParameters,
// - *SQLDedicatedGatewayServiceResourceCreateUpdateParameters, *ServiceResourceCreateUpdateProperties
type ServiceResourceCreateUpdatePropertiesClassification interface {
	// GetServiceResourceCreateUpdateProperties returns the ServiceResourceCreateUpdateProperties content of the underlying type.
	GetServiceResourceCreateUpdateProperties() *ServiceResourceCreateUpdateProperties
}

// ServiceResourcePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetServiceResourceProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataTransferServiceResourceProperties, *GraphAPIComputeServiceResourceProperties, *MaterializedViewsBuilderServiceResourceProperties,
// - *SQLDedicatedGatewayServiceResourceProperties, *ServiceResourceProperties
type ServiceResourcePropertiesClassification interface {
	// GetServiceResourceProperties returns the ServiceResourceProperties content of the underlying type.
	GetServiceResourceProperties() *ServiceResourceProperties
}

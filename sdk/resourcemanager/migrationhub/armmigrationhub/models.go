//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationhub

import "time"

type AADAppDetails struct {
	ApplicationID *string
	TenantID      *string
}

// AssessmentDetails - Assessment properties that can be shared by various publishers.
type AssessmentDetails struct {
	// Gets or sets the id of the assessment done on the machine.
	AssessmentID *string

	// Gets or sets the BIOS ID of the machine.
	BiosID *string

	// Gets or sets the time the message was enqueued.
	EnqueueTime *string

	// Gets or sets the ISV specific extended information.
	ExtendedInfo map[string]*string

	// Gets or sets the fabric type.
	FabricType *string

	// Gets or sets the FQDN of the machine.
	Fqdn *string

	// Gets the relative URL to get to this AssessmentDetails resource.
	ID *string

	// Gets or sets the list of IP addresses of the machine. IP addresses could be IP V4 or IP V6.
	IPAddresses []*string

	// Gets or sets the time of the last modification of the machine details.
	LastUpdatedTime *time.Time

	// Gets or sets the list of MAC addresses of the machine.
	MacAddresses []*string

	// Gets or sets the unique identifier of the machine.
	MachineID *string

	// Gets or sets the unique identifier of the virtual machine manager(vCenter/VMM).
	MachineManagerID *string

	// Gets or sets the name of the machine.
	MachineName *string

	// Gets or sets the name of the solution that sent the data.
	SolutionName *string

	// Gets or sets the target storage type.
	TargetStorageType map[string]*string

	// Gets or sets the target VM location.
	TargetVMLocation *string

	// Gets or sets the target VM size.
	TargetVMSize *string
}

// Database REST resource.
type Database struct {
	// Gets or sets the relative URL to get to this REST resource.
	ID *string

	// Gets or sets the name of this REST resource.
	Name *string

	// Gets or sets the properties of the database.
	Properties *DatabaseProperties

	// READ-ONLY; Gets the type of this REST resource.
	Type *string
}

// DatabaseAssessmentDetails - Assessment properties that can be shared by various publishers.
type DatabaseAssessmentDetails struct {
	// Gets or sets the database assessment scope/Id.
	AssessmentID *string

	// Gets or sets the assessed target database type.
	AssessmentTargetType *string

	// Gets or sets the number of breaking changes found.
	BreakingChangesCount *int32

	// Gets or sets the compatibility level of the database.
	CompatibilityLevel *string

	// Gets or sets the database name.
	DatabaseName *string

	// Gets or sets the database size.
	DatabaseSizeInMB *string

	// Gets or sets the time the message was enqueued.
	EnqueueTime *string

	// Gets or sets the extended properties of the database.
	ExtendedInfo map[string]*string

	// Gets the relative URL to get to this DatabaseAssessmentDetails resource.
	ID *string

	// Gets or sets the database server instance Id.
	InstanceID *string

	// Gets or sets a value indicating whether the database is ready for migration.
	IsReadyForMigration *bool

	// Gets or sets the time when the database was last assessed.
	LastAssessedTime *time.Time

	// Gets or sets the time of the last modification of the database details.
	LastUpdatedTime *time.Time

	// Gets or sets the number of blocking changes found.
	MigrationBlockersCount *int32

	// Gets or sets the name of the solution that sent the data.
	SolutionName *string
}

// DatabaseCollection - Collection of databases.
type DatabaseCollection struct {
	// Gets or sets the value of nextLink.
	NextLink *string

	// Gets or sets the databases.
	Value []*Database
}

// DatabaseInstance REST resource.
type DatabaseInstance struct {
	// Gets or sets the relative URL to get to this REST resource.
	ID *string

	// Gets or sets the name of this REST resource.
	Name *string

	// Gets or sets the properties of the machine.
	Properties *DatabaseInstanceProperties

	// READ-ONLY; Gets the type of this REST resource.
	Type *string
}

// DatabaseInstanceCollection - Collection of database instances.
type DatabaseInstanceCollection struct {
	// Gets or sets the value of nextLink.
	NextLink *string

	// Gets or sets the database instances.
	Value []*DatabaseInstance
}

// DatabaseInstanceDiscoveryDetails - Discovery properties that can be shared by various publishers.
type DatabaseInstanceDiscoveryDetails struct {
	// Gets or sets the time the message was enqueued.
	EnqueueTime *string

	// Gets or sets the extended properties of the database server.
	ExtendedInfo map[string]*string

	// Gets or sets the host name of the database server.
	HostName *string

	// Gets the relative URL to get to this DatabaseInstanceDiscoveryDetails resource.
	ID *string

	// Gets or sets the IP addresses of the database server. IP addresses could be IP V4 or IP V6.
	IPAddress *string

	// Gets or sets the database instance Id.
	InstanceID *string

	// Gets or sets the database instance name.
	InstanceName *string

	// Gets or sets the database instance type.
	InstanceType *string

	// Gets or sets the database instance version.
	InstanceVersion *string

	// Gets or sets the time of the last modification of the database instance details.
	LastUpdatedTime *time.Time

	// Gets or sets the port number of the database server.
	PortNumber *int32

	// Gets or sets the name of the solution that sent the data.
	SolutionName *string
}

// DatabaseInstanceProperties - Properties of the database instance resource.
type DatabaseInstanceProperties struct {
	// Gets or sets the assessment details of the database instance published by various sources.
	DiscoveryData []*DatabaseInstanceDiscoveryDetails

	// Gets or sets the time of the last modification of the database.
	LastUpdatedTime *time.Time

	// Gets or sets the database instances summary per solution. The key of dictionary is the solution name and value is the corresponding
	// database instance summary object.
	Summary map[string]*DatabaseInstanceSummary
}

// DatabaseInstanceSummary - Class representing the database instance summary object.
type DatabaseInstanceSummary struct {
	// Gets or sets the count of databases assessed.
	DatabasesAssessedCount *int32

	// Gets or sets the count of databases ready for migration.
	MigrationReadyCount *int32
}

// DatabaseProperties - Properties of the database resource.
type DatabaseProperties struct {
	// Gets or sets the assessment details of the database published by various sources.
	AssessmentData []*DatabaseAssessmentDetails

	// Gets or sets the time of the last modification of the database.
	LastUpdatedTime *time.Time
}

// DiscoveryDetails - Discovery properties that can be published by various ISVs.
type DiscoveryDetails struct {
	// Gets or sets the BIOS ID of the machine.
	BiosID *string

	// Gets or sets the time the message was enqueued.
	EnqueueTime *string

	// Gets or sets the ISV specific extended information.
	ExtendedInfo map[string]*string

	// Gets or sets the fabric type.
	FabricType *string

	// Gets or sets the FQDN of the machine.
	Fqdn *string

	// Gets the relative URL to get to this DiscoveryDetails resource.
	ID *string

	// Gets or sets the list of IP addresses of the machine. IP addresses could be IP V4 or IP V6.
	IPAddresses []*string

	// Gets or sets the time of the last modification of the machine details.
	LastUpdatedTime *time.Time

	// Gets or sets the list of MAC addresses of the machine.
	MacAddresses []*string

	// Gets or sets the unique identifier of the machine.
	MachineID *string

	// Gets or sets the unique identifier of the virtual machine manager(vCenter/VMM).
	MachineManagerID *string

	// Gets or sets the name of the machine.
	MachineName *string

	// Gets or sets the OS name.
	OSName *string

	// Gets or sets the OS type.
	OSType *string

	// Gets or sets the OS version.
	OSVersion *string

	// Gets or sets the name of the solution that sent the data.
	SolutionName *string
}

// EventCollection - Collection of events.
type EventCollection struct {
	// Gets or sets the value of nextLink.
	NextLink *string

	// Gets or sets the machines.
	Value []*MigrateEvent
}

// GroupConnectivityInformation - Defines Private link service group connectivity.
type GroupConnectivityInformation struct {
	CustomerVisibleFqdns        []*string
	GroupID                     *string
	ID                          *string
	InternalFqdn                *string
	MemberName                  *string
	PrivateLinkServiceArmRegion *string
	RedirectMapID               *string
}

// IPConfiguration - Defines Private link IP configuration.
type IPConfiguration struct {
	GroupID          *string
	ID               *string
	LinkIdentifier   *string
	MemberName       *string
	PrivateIPAddress *string
}

// Machine REST resource.
type Machine struct {
	// Gets or sets the relative URL to get to this REST resource.
	ID *string

	// Gets or sets the name of this REST resource.
	Name *string

	// Gets or sets the properties of the machine.
	Properties *MachineProperties

	// READ-ONLY; Gets the type of this REST resource.
	Type *string
}

// MachineCollection - Collection of machines.
type MachineCollection struct {
	// Gets or sets the value of nextLink.
	NextLink *string

	// Gets or sets the machines.
	Value []*Machine
}

// MachineProperties - Properties of the machine resource.
type MachineProperties struct {
	// Gets or sets the assessment details of the machine published by various sources.
	AssessmentData []*AssessmentDetails

	// Gets or sets the discovery details of the machine published by various sources.
	DiscoveryData []*DiscoveryDetails

	// Gets or sets the time of the last modification of the machine.
	LastUpdatedTime *time.Time

	// Gets or sets the migration details of the machine published by various sources.
	MigrationData []*MigrationDetails
}

// MigrateEvent REST resource.
type MigrateEvent struct {
	// Gets or sets the relative URL to get to this REST resource.
	ID *string

	// Gets or sets the name of this REST resource.
	Name *string

	// Gets or sets the properties of the event.
	Properties *MigrateEventProperties

	// READ-ONLY; Gets the type of this REST resource.
	Type *string
}

// MigrateEventProperties - Properties of the error resource.
type MigrateEventProperties struct {
	// Gets or sets the client request Id of the payload for which the event is being reported.
	ClientRequestID *string

	// Gets or sets the error code.
	ErrorCode *string

	// Gets or sets the error message.
	ErrorMessage *string

	// Gets or sets the possible causes for the error.
	PossibleCauses *string

	// Gets or sets the recommendation for the error.
	Recommendation *string

	// Gets or sets the solution for which the error is being reported.
	Solution *string

	// READ-ONLY; Gets the Instance type.
	InstanceType *string
}

// MigrateProject - Migrate project.
type MigrateProject struct {
	// For optimistic concurrency control.
	ETag *string

	// Azure location in which project is created.
	Location *string

	// Properties of a migrate project.
	Properties *MigrateProjectProperties

	// READ-ONLY; Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{projectName}
	ID *string

	// READ-ONLY; Name of the project.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Type of the object = [Microsoft.Migrate/migrateProjects].
	Type *string
}

// MigrateProjectProperties - Properties of a migrate project.
type MigrateProjectProperties struct {
	// Gets or sets the state of public network access.
	PublicNetworkAccess *PublicNetworkAccess

	// Service endpoint.
	ServiceEndpoint *string

	// Utility storage account id.
	UtilityStorageAccountID *string

	// READ-ONLY; Last summary refresh time.
	LastSummaryRefreshedTime *time.Time

	// READ-ONLY; Gets the private endpoint connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; Refresh summary state.
	RefreshSummaryState *RefreshSummaryState

	// READ-ONLY; Register tools inside project.
	RegisteredTools []*Items

	// READ-ONLY; Project summary.
	Summary map[string]*ProjectSummary
}

// MigrationDetails - Migration properties that can be shared by various publishers.
type MigrationDetails struct {
	// Gets or sets the BIOS ID of the machine.
	BiosID *string

	// Gets or sets the time the message was enqueued.
	EnqueueTime *string

	// Gets or sets the ISV specific extended information.
	ExtendedInfo map[string]*string

	// Gets or sets the fabric type.
	FabricType *string

	// Gets or sets the FQDN of the machine.
	Fqdn *string

	// Gets the relative URL to get to this MigrationDetails resource.
	ID *string

	// Gets or sets the list of IP addresses of the machine. IP addresses could be IP V4 or IP V6.
	IPAddresses []*string

	// Gets or sets the time of the last modification of the machine details.
	LastUpdatedTime *time.Time

	// Gets or sets the list of MAC addresses of the machine.
	MacAddresses []*string

	// Gets or sets the unique identifier of the machine.
	MachineID *string

	// Gets or sets the unique identifier of the virtual machine manager(vCenter/VMM).
	MachineManagerID *string

	// Gets or sets the name of the machine.
	MachineName *string

	// Gets or sets the phase of migration of the machine.
	MigrationPhase *string

	// Gets or sets a value indicating whether migration was tested on the machine.
	MigrationTested *bool

	// Gets or sets the progress percentage of migration on the machine.
	ReplicationProgressPercentage *int32

	// Gets or sets the name of the solution that sent the data.
	SolutionName *string

	// Gets or sets the ARM id the migrated VM.
	TargetVMArmID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// PrivateEndpointConnection - REST model used to encapsulate the user visible state of a PrivateEndpoint.
type PrivateEndpointConnection struct {
	// Gets the properties of the object.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Gets the tag for optimistic concurrency control.
	ETag *string

	// READ-ONLY; Relative URL to get this Sites.
	ID *string

	// READ-ONLY; Gets the name of the resource.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Gets the resource type.
	Type *string
}

// PrivateEndpointConnectionCollection - Collection of PrivateLink resources.
type PrivateEndpointConnectionCollection struct {
	// READ-ONLY; Gets the value of next link.
	NextLink *string

	// READ-ONLY; Gets the list of machines.
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of a private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// Gets the properties of the object.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// READ-ONLY; Defines resource ID of a private endpoint connection.
	PrivateEndpoint *ResourceID

	// READ-ONLY; Provisioning state.
	ProvisioningState *ProvisioningState
}

// PrivateEndpointConnectionProxy - Defines Private endpoint proxy resource.
type PrivateEndpointConnectionProxy struct {
	ETag *string

	// Properties of a private endpoint connection proxy.
	Properties *PrivateEndpointConnectionProxyProperties

	// READ-ONLY
	ID *string

	// READ-ONLY
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY
	Type *string
}

// PrivateEndpointConnectionProxyCollection - Collection of PrivateLink proxy resources.
type PrivateEndpointConnectionProxyCollection struct {
	// READ-ONLY; Gets the value of next link.
	NextLink *string

	// READ-ONLY; Gets the list of PrivateLink proxy resources.
	Value []*PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxyProperties - Properties of a private endpoint connection proxy.
type PrivateEndpointConnectionProxyProperties struct {
	// READ-ONLY; Defines Private endpoint additional details.
	RemotePrivateEndpoint *PrivateEndpointDetails

	// READ-ONLY
	Status *PrivateEndpointConnectionProxyPropertiesStatus
}

// PrivateEndpointDetails - Defines Private endpoint additional details.
type PrivateEndpointDetails struct {
	ConnectionDetails                   []*IPConfiguration
	ID                                  *string
	ManualPrivateLinkServiceConnections []*PrivateLinkServiceConnection
	PrivateLinkServiceConnections       []*PrivateLinkServiceConnection
	PrivateLinkServiceProxies           []*PrivateLinkServiceProxy
}

// PrivateLinkResource - Private link resource.
type PrivateLinkResource struct {
	// READ-ONLY; Relative URL to get this Sites.
	ID *string

	// READ-ONLY; Gets the name of the resource.
	Name *string

	// READ-ONLY; Gets nested properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Gets the resource type.
	Type *string
}

// PrivateLinkResourceCollection - Collection of private link resources.
type PrivateLinkResourceCollection struct {
	// READ-ONLY; Value of next link.
	NextLink *string

	// READ-ONLY; List of private links.
	Value []*PrivateLinkResource
}

// PrivateLinkResourceProperties - Properties of private link resource.
type PrivateLinkResourceProperties struct {
	// Group id.
	GroupID *string

	// Required members.
	RequiredMembers []*string

	// Required zone names.
	RequiredZoneNames []*string
}

// PrivateLinkServiceConnection - Defines Private link service connection.
type PrivateLinkServiceConnection struct {
	GroupIDs       []*string
	ID             *string
	Name           *string
	RequestMessage *string
}

// PrivateLinkServiceConnectionState - Private endpoint connection state.
type PrivateLinkServiceConnectionState struct {
	// Action required.
	ActionsRequired *string

	// Description of the object.
	Description *string

	// Private link connection state.
	Status *PrivateLinkServiceConnectionStateStatus
}

// PrivateLinkServiceProxy - Defines Private link service proxy.
type PrivateLinkServiceProxy struct {
	GroupConnectivityInformation []*GroupConnectivityInformation
	ID                           *string

	// Defines resource ID of a private endpoint connection.
	RemotePrivateEndpointConnection *ResourceID

	// Private endpoint connection state.
	RemotePrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState
}

// ProjectResultList - List of projects.
type ProjectResultList struct {
	NextLink *string

	// List of projects.
	Value []*MigrateProject
}

// ProjectSummary - Project summary.
type ProjectSummary struct {
	// Extended summary.
	ExtendedSummary map[string]*string

	// Last summary refresh time.
	LastSummaryRefreshedTime *time.Time

	// Refresh summary state.
	RefreshSummaryState *RefreshSummaryState

	// READ-ONLY; Instance type.
	InstanceType *string
}

// RefreshSummaryInput - Class representing the refresh summary input.
type RefreshSummaryInput struct {
	// Gets or sets the goal for which summary needs to be refreshed.
	Goal *Goal
}

// RefreshSummaryResult - Class representing the refresh summary status of the migrate project.
type RefreshSummaryResult struct {
	// Gets or sets a value indicating whether the migrate project summary is refreshed.
	IsRefreshed *bool
}

// RegisterToolInput - Class representing the register tool input.
type RegisterToolInput struct {
	// Gets or sets the tool to be registered.
	Tool *Tool
}

type RegistrationDetailsInput struct {
	ApplicationDetails *AADAppDetails

	// Gets or sets the tool to be registered.
	Tool *Tool
}

type RegistrationDetailsResponse struct {
	OneTimeKey      *string
	ServiceEndpoint *string
}

// RegistrationResult - Class representing the registration status of a tool with the migrate project.
type RegistrationResult struct {
	// Gets or sets a value indicating whether the tool is registered or not.
	IsRegistered *bool
}

// ResourceID - Defines resource ID of a private endpoint connection.
type ResourceID struct {
	// READ-ONLY
	ID *string
}

// Solution REST Resource.
type Solution struct {
	// Gets or sets the ETAG for optimistic concurrency control.
	Etag *string

	// Gets or sets the properties of the solution.
	Properties *SolutionProperties

	// READ-ONLY; Gets the relative URL to get to this REST resource.
	ID *string

	// READ-ONLY; Gets the name of this REST resource.
	Name *string

	// READ-ONLY; Gets the type of this REST resource.
	Type *string
}

// SolutionConfig - Class representing the config for the solution in the migrate project.
type SolutionConfig struct {
	// Gets or sets the publisher sas uri for the solution.
	PublisherSasURI *string
}

// SolutionDetails - Class representing the details of the solution.
type SolutionDetails struct {
	// Gets or sets the count of assessments reported by the solution.
	AssessmentCount *int32

	// Gets or sets the extended details reported by the solution.
	ExtendedDetails map[string]*string

	// Gets or sets the count of groups reported by the solution.
	GroupCount *int32
}

// SolutionProperties - Class for solution properties.
type SolutionProperties struct {
	// Gets or sets the cleanup state of the solution.
	CleanupState *CleanupState

	// Gets or sets the details of the solution.
	Details *SolutionDetails

	// Gets or sets the goal of the solution.
	Goal *Goal

	// Gets or sets the purpose of the solution.
	Purpose *Purpose

	// Gets or sets the current status of the solution.
	Status *Status

	// Gets or sets the summary of the solution.
	Summary *SolutionSummary

	// Gets or sets the tool being used in the solution.
	Tool *Tool
}

// SolutionSummary - The solution summary class.
type SolutionSummary struct {
	// READ-ONLY; Gets the Instance type.
	InstanceType *string
}

// SolutionsCollection - Collection of solutions.
type SolutionsCollection struct {
	// Gets or sets the value of next link.
	NextLink *string

	// Gets or sets the list of solutions.
	Value []*Solution
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

// VirtualDesktopUser - Class representing virtual desktop user.
type VirtualDesktopUser struct {
	ID   *string
	Name *string

	// Properties of class representing virtual desktop user.
	Properties *VirtualDesktopUserProperties

	// READ-ONLY
	Type *string
}

// VirtualDesktopUserAssessmentDetails - Details of assessment data of virtual desktop user.
type VirtualDesktopUserAssessmentDetails struct {
	ActiveWeeklyHours     *int32
	AssessmentID          *string
	City                  *string
	Country               *string
	CriticalApplications  []*string
	DevicesUsed           []*string
	EgressBandwidthWeekly *float64
	EnqueueTime           *string

	// Dictionary of
	ExtendedInfo map[string]*string

	// Gets the relative URL to get to this VirtualDesktopUserAssessmentDetails resource.
	ID                     *string
	IsReadyForMigration    *bool
	LastUpdatedTime        *time.Time
	MultiUserWindows10     *bool
	OSUsed                 []*string
	Persona                *string
	SolutionName           *string
	State                  *string
	TargetAzureVMSize      *string
	TargetLocation         *string
	TargetStorageType      *string
	TotalApplicationsCount *int32
	UserAccount            *string
	UserExperienceScore    *float64
	UserID                 *string
	UserName               *string
	Virtualization         *string
	Windows7               *bool
}

// VirtualDesktopUserCollection - Collection of virtual desktop users.
type VirtualDesktopUserCollection struct {
	NextLink *string
	Value    []*VirtualDesktopUser
}

// VirtualDesktopUserProperties - Properties of class representing virtual desktop user.
type VirtualDesktopUserProperties struct {
	AssessmentData  []*VirtualDesktopUserAssessmentDetails
	LastUpdatedTime *time.Time
}

// WebServer - Class representing a web server.
type WebServer struct {
	ID   *string
	Name *string

	// Properties of class representing web server.
	Properties *WebServerProperties

	// READ-ONLY
	Type *string
}

// WebServerCollection - Collection of web servers.
type WebServerCollection struct {
	NextLink *string
	Value    []*WebServer
}

// WebServerDiscoveryDetails - Details of discovery data of web server.
type WebServerDiscoveryDetails struct {
	CPUCores    *int32
	EnqueueTime *string

	// Dictionary of
	ExtendedInfo map[string]*string

	// Gets the relative URL to get to this WebServerDiscoveryDetails resource.
	ID               *string
	LastUpdatedTime  *time.Time
	MemoryInMb       *string
	OSName           *string
	OSVersion        *string
	PortList         []*int32
	SolutionName     *string
	WebServerID      *string
	WebServerName    *string
	WebServerType    *string
	WebServerVersion *string
}

// WebServerProperties - Properties of class representing web server.
type WebServerProperties struct {
	DiscoveryData   []*WebServerDiscoveryDetails
	LastUpdatedTime *time.Time

	// Dictionary of
	Summary map[string]*WebServerSummary
}

// WebServerSummary - Summary of information representing web server.
type WebServerSummary struct {
	AssessedCount     *int32
	DiscoveredCount   *int32
	MigratedCount     *int32
	MigratingCount    *int32
	ReadyForMigration *int32
}

// WebSite - Class representing a web site.
type WebSite struct {
	ID   *string
	Name *string

	// Properties of class representing web site.
	Properties *WebSiteProperties

	// READ-ONLY
	Type *string
}

// WebSiteAssessmentDetails - Details of assessment data of web site.
type WebSiteAssessmentDetails struct {
	AssessmentID         *string
	AssessmentTargetType *string
	EnqueueTime          *string
	ErrorList            []*string

	// Dictionary of
	ExtendedInfo     map[string]*string
	Framework        *string
	FrameworkVersion *string

	// Gets the relative URL to get to this WebSiteAssessmentDetails resource.
	ID                     *string
	IsReadyForMigration    *bool
	LastUpdatedTime        *time.Time
	MigrationBlockersCount *int32
	Port                   *int32
	SolutionName           *string
	SuccessList            []*string
	WarningList            []*string
	WebServerID            *string
	WebServerType          *string
	WebSiteName            *string
}

// WebSiteCollection - Collection of web sites.
type WebSiteCollection struct {
	NextLink *string
	Value    []*WebSite
}

// WebSiteDiscoveryDetails - Details of discovery data of web site.
type WebSiteDiscoveryDetails struct {
	EnqueueTime *string

	// Dictionary of
	ExtendedInfo map[string]*string

	// Gets the relative URL to get to this WebSiteDiscoveryDetails resource.
	ID              *string
	LastUpdatedTime *time.Time
	Port            *int32
	SolutionName    *string
	WebServerID     *string
	WebServerType   *string
	WebSiteName     *string
}

// WebSiteMigrationDetails - Details of migration data of web site.
type WebSiteMigrationDetails struct {
	EnqueueTime *string

	// Dictionary of
	ExtendedInfo map[string]*string

	// Gets the relative URL to get to this WebSiteMigrationDetails resource.
	ID                    *string
	LastUpdatedTime       *time.Time
	MigrationPhase        *string
	Port                  *int32
	ProgressPercentage    *int32
	SolutionName          *string
	TargetAppServiceArmID *string
	WebServerID           *string
	WebServerType         *string
	WebSiteName           *string
}

// WebSiteProperties - Properties of class representing web site.
type WebSiteProperties struct {
	AssessmentData  []*WebSiteAssessmentDetails
	DiscoveryData   []*WebSiteDiscoveryDetails
	LastUpdatedTime *time.Time
	MigrationData   []*WebSiteMigrationDetails
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcegraph

import "time"

// ClientResourceChangeDetailsOptions contains the optional parameters for the Client.ResourceChangeDetails method.
type ClientResourceChangeDetailsOptions struct {
	// placeholder for future optional parameters
}

// ClientResourceChangesOptions contains the optional parameters for the Client.ResourceChanges method.
type ClientResourceChangesOptions struct {
	// placeholder for future optional parameters
}

// ClientResourcesHistoryOptions contains the optional parameters for the Client.ResourcesHistory method.
type ClientResourcesHistoryOptions struct {
	// placeholder for future optional parameters
}

// ClientResourcesOptions contains the optional parameters for the Client.Resources method.
type ClientResourcesOptions struct {
	// placeholder for future optional parameters
}

// Column - Query result column descriptor.
type Column struct {
	// REQUIRED; Column name.
	Name *string

	// REQUIRED; Column data type.
	Type *ColumnDataType
}

// DateTimeInterval - An interval in time specifying the date and time for the inclusive start and exclusive end, i.e. [start,
// end).
type DateTimeInterval struct {
	// REQUIRED; A datetime indicating the exclusive/open end of the time interval, i.e. [start,end). Specifying an end that occurs
	// chronologically before start will result in an error.
	End *time.Time

	// REQUIRED; A datetime indicating the inclusive/closed start of the time interval, i.e. [start, end). Specifying a start
	// that occurs chronologically after end will result in an error.
	Start *time.Time
}

// Error details.
type Error struct {
	// REQUIRED; Error code identifying the specific error.
	Code *string

	// REQUIRED; A human readable error message.
	Message *string

	// Error details
	Details []*ErrorDetails
}

// ErrorDetails - Error details.
type ErrorDetails struct {
	// REQUIRED; Error code identifying the specific error.
	Code *string

	// REQUIRED; A human readable error message.
	Message *string

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any
}

// ErrorResponse - An error response from the API.
type ErrorResponse struct {
	// REQUIRED; Error information.
	Error *Error
}

// FacetClassification provides polymorphic access to related types.
// Call the interface's GetFacet() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Facet, *FacetError, *FacetResult
type FacetClassification interface {
	// GetFacet returns the Facet content of the underlying type.
	GetFacet() *Facet
}

// Facet - A facet containing additional statistics on the response of a query. Can be either FacetResult or FacetError.
type Facet struct {
	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string

	// REQUIRED; Result type
	ResultType *string
}

// GetFacet implements the FacetClassification interface for type Facet.
func (f *Facet) GetFacet() *Facet { return f }

// FacetError - A facet whose execution resulted in an error.
type FacetError struct {
	// REQUIRED; An array containing detected facet errors with details.
	Errors []*ErrorDetails

	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string

	// REQUIRED; Result type
	ResultType *string
}

// GetFacet implements the FacetClassification interface for type FacetError.
func (f *FacetError) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// FacetRequest - A request to compute additional statistics (facets) over the query results.
type FacetRequest struct {
	// REQUIRED; The column or list of columns to summarize by
	Expression *string

	// The options for facet evaluation
	Options *FacetRequestOptions
}

// FacetRequestOptions - The options for facet evaluation
type FacetRequestOptions struct {
	// Specifies the filter condition for the 'where' clause which will be run on main query's result, just before the actual
	// faceting.
	Filter *string

	// The column name or query expression to sort on. Defaults to count if not present.
	SortBy *string

	// The sorting order by the selected column (count by default).
	SortOrder *FacetSortOrder

	// The maximum number of facet rows that should be returned.
	Top *int32
}

// FacetResult - Successfully executed facet containing additional statistics on the response of a query.
type FacetResult struct {
	// REQUIRED; Number of records returned in the facet response.
	Count *int32

	// REQUIRED; A JObject array or Table containing the desired facets. Only present if the facet is valid.
	Data any

	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string

	// REQUIRED; Result type
	ResultType *string

	// REQUIRED; Number of total records in the facet results.
	TotalRecords *int64
}

// GetFacet implements the FacetClassification interface for type FacetResult.
func (f *FacetResult) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// Operation - Resource Graph REST API operation definition.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string

	// The origin of operations.
	Origin *string
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Description for the operation.
	Description *string

	// Type of operation: get, read, delete, etc.
	Operation *string

	// Service provider: Microsoft Resource Graph.
	Provider *string

	// Resource on which the operation is performed etc.
	Resource *string
}

// OperationListResult - Result of the request to list Resource Graph operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	// List of Resource Graph operations supported by the Resource Graph resource provider.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// QueryRequest - Describes a query to be executed.
type QueryRequest struct {
	// REQUIRED; The resources query.
	Query *string

	// An array of facet requests to be computed against the query result.
	Facets []*FacetRequest

	// Azure management groups against which to execute the query. Example: [ 'mg1', 'mg2' ]
	ManagementGroups []*string

	// The query evaluation options
	Options *QueryRequestOptions

	// Azure subscriptions against which to execute the query.
	Subscriptions []*string
}

// QueryRequestOptions - The options for query evaluation
type QueryRequestOptions struct {
	// Only applicable for tenant and management group level queries to decide whether to allow partial scopes for result in case
	// the number of subscriptions exceed allowed limits.
	AllowPartialScopes *bool

	// Defines what level of authorization resources should be returned based on the which subscriptions and management groups
	// are passed as scopes.
	AuthorizationScopeFilter *AuthorizationScopeFilter

	// Defines in which format query result returned.
	ResultFormat *ResultFormat

	// The number of rows to skip from the beginning of the results. Overrides the next page offset when $skipToken property is
	// present.
	Skip *int32

	// Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string

	// The maximum number of rows that the query should return. Overrides the page size when $skipToken property is present.
	Top *int32
}

// QueryResponse - Query result.
type QueryResponse struct {
	// REQUIRED; Number of records returned in the current response. In the case of paging, this is the number of records in the
	// current page.
	Count *int64

	// REQUIRED; Query output in JObject array or Table format.
	Data any

	// REQUIRED; Indicates whether the query results are truncated.
	ResultTruncated *ResultTruncated

	// REQUIRED; Number of total records matching the query.
	TotalRecords *int64

	// Query facets.
	Facets []FacetClassification

	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current
	// request) to retrieve the next page of data.
	SkipToken *string
}

// ResourceChangeData - Data on a specific change, represented by a pair of before and after resource snapshots.
type ResourceChangeData struct {
	// REQUIRED; The snapshot after the change.
	AfterSnapshot *ResourceChangeDataAfterSnapshot

	// REQUIRED; The snapshot before the change.
	BeforeSnapshot *ResourceChangeDataBeforeSnapshot

	// REQUIRED; The change ID. Valid and unique within the specified resource only.
	ChangeID *string

	// The change type for snapshot. PropertyChanges will be provided in case of Update change type
	ChangeType *ChangeType

	// An array of resource property change
	PropertyChanges []*ResourcePropertyChange

	// The resource for a change.
	ResourceID *string
}

// ResourceChangeDataAfterSnapshot - The snapshot after the change.
type ResourceChangeDataAfterSnapshot struct {
	// REQUIRED; The time when the snapshot was created. The snapshot timestamp provides an approximation as to when a modification
	// to a resource was detected. There can be a difference between the actual modification
	// time and the detection time. This is due to differences in how operations that modify a resource are processed, versus
	// how operation that record resource snapshots are processed.
	Timestamp *time.Time

	// The resource snapshot content (in resourceChangeDetails response only).
	Content any

	// The ID of the snapshot.
	SnapshotID *string
}

// ResourceChangeDataBeforeSnapshot - The snapshot before the change.
type ResourceChangeDataBeforeSnapshot struct {
	// REQUIRED; The time when the snapshot was created. The snapshot timestamp provides an approximation as to when a modification
	// to a resource was detected. There can be a difference between the actual modification
	// time and the detection time. This is due to differences in how operations that modify a resource are processed, versus
	// how operation that record resource snapshots are processed.
	Timestamp *time.Time

	// The resource snapshot content (in resourceChangeDetails response only).
	Content any

	// The ID of the snapshot.
	SnapshotID *string
}

// ResourceChangeDetailsRequestParameters - The parameters for a specific change details request.
type ResourceChangeDetailsRequestParameters struct {
	// REQUIRED; Specifies the list of change IDs for a change details request.
	ChangeIDs []*string

	// REQUIRED; Specifies the list of resources for a change details request.
	ResourceIDs []*string
}

// ResourceChangeList - A list of changes associated with a resource over a specific time interval.
type ResourceChangeList struct {
	// The pageable value returned by the operation, i.e. a list of changes to the resource.
	// * The list is ordered from the most recent changes to the least recent changes.
	// * This list will be empty if there were no changes during the requested interval.
	// * The Before snapshot timestamp value of the oldest change can be outside of the specified time interval.
	Changes []*ResourceChangeData

	// Skip token that encodes the skip information while executing the current request
	SkipToken any
}

// ResourceChangesRequestParameters - The parameters for a specific changes request.
type ResourceChangesRequestParameters struct {
	// REQUIRED; Specifies the date and time interval for a changes request.
	Interval *ResourceChangesRequestParametersInterval

	// The flag if set to true will fetch property changes
	FetchPropertyChanges *bool

	// The flag if set to true will fetch change snapshots
	FetchSnapshots *bool

	// Specifies the list of resources for a changes request.
	ResourceIDs []*string

	// Acts as the continuation token for paged responses.
	SkipToken *string

	// The subscription id of resources to query the changes from.
	SubscriptionID *string

	// The table name to query resources from.
	Table *string

	// The maximum number of changes the client can accept in a paged response.
	Top *int32
}

// ResourceChangesRequestParametersInterval - Specifies the date and time interval for a changes request.
type ResourceChangesRequestParametersInterval struct {
	// REQUIRED; A datetime indicating the exclusive/open end of the time interval, i.e. [start,end). Specifying an end that occurs
	// chronologically before start will result in an error.
	End *time.Time

	// REQUIRED; A datetime indicating the inclusive/closed start of the time interval, i.e. [start, end). Specifying a start
	// that occurs chronologically after end will result in an error.
	Start *time.Time
}

// ResourcePropertyChange - The resource property change
type ResourcePropertyChange struct {
	// REQUIRED; The change category.
	ChangeCategory *ChangeCategory

	// REQUIRED; The property change Type
	PropertyChangeType *PropertyChangeType

	// REQUIRED; The property name
	PropertyName *string

	// The property value in after snapshot
	AfterValue *string

	// The property value in before snapshot
	BeforeValue *string
}

// ResourceSnapshotData - Data on a specific resource snapshot.
type ResourceSnapshotData struct {
	// REQUIRED; The time when the snapshot was created. The snapshot timestamp provides an approximation as to when a modification
	// to a resource was detected. There can be a difference between the actual modification
	// time and the detection time. This is due to differences in how operations that modify a resource are processed, versus
	// how operation that record resource snapshots are processed.
	Timestamp *time.Time

	// The resource snapshot content (in resourceChangeDetails response only).
	Content any

	// The ID of the snapshot.
	SnapshotID *string
}

// ResourcesHistoryRequest - Describes a history request to be executed.
type ResourcesHistoryRequest struct {
	// Azure management groups against which to execute the query. Example: [ 'mg1', 'mg2' ]
	ManagementGroups []*string

	// The history request evaluation options
	Options *ResourcesHistoryRequestOptions

	// The resources query.
	Query *string

	// Azure subscriptions against which to execute the query.
	Subscriptions []*string
}

// ResourcesHistoryRequestOptions - The options for history request evaluation
type ResourcesHistoryRequestOptions struct {
	// The time interval used to fetch history.
	Interval *DateTimeInterval

	// Defines in which format query result returned.
	ResultFormat *ResultFormat

	// The number of rows to skip from the beginning of the results. Overrides the next page offset when $skipToken property is
	// present.
	Skip *int32

	// Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string

	// The maximum number of rows that the query should return. Overrides the page size when $skipToken property is present.
	Top *int32
}

// Table - Query output in tabular format.
type Table struct {
	// REQUIRED; Query result column descriptors.
	Columns []*Column

	// REQUIRED; Query result rows.
	Rows [][]any
}

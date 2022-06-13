package datacollaborationapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/datacollaboration/mgmt/2020-05-04-preview/datacollaboration"
)

// ConsumerInvitationsClientAPI contains the set of methods on the ConsumerInvitationsClient type.
type ConsumerInvitationsClientAPI interface {
	Get(ctx context.Context, location string, invitationID string) (result datacollaboration.ConsumerInvitation, err error)
	ListInvitations(ctx context.Context, skipToken string) (result datacollaboration.ConsumerInvitationListPage, err error)
	ListInvitationsComplete(ctx context.Context, skipToken string) (result datacollaboration.ConsumerInvitationListIterator, err error)
	RejectInvitation(ctx context.Context, location string, invitationID string) (result datacollaboration.ConsumerInvitation, err error)
}

var _ ConsumerInvitationsClientAPI = (*datacollaboration.ConsumerInvitationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result datacollaboration.OperationListPage, err error)
	ListComplete(ctx context.Context) (result datacollaboration.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*datacollaboration.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspace datacollaboration.Workspace) (result datacollaboration.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result datacollaboration.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result datacollaboration.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result datacollaboration.WorkspaceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string) (result datacollaboration.WorkspaceListIterator, err error)
	ListBySubscription(ctx context.Context, skipToken string) (result datacollaboration.WorkspaceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skipToken string) (result datacollaboration.WorkspaceListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, workspaceUpdateParameters datacollaboration.WorkspaceUpdateParameters) (result datacollaboration.Workspace, err error)
}

var _ WorkspacesClientAPI = (*datacollaboration.WorkspacesClient)(nil)

// ConstrainedResourcesClientAPI contains the set of methods on the ConstrainedResourcesClient type.
type ConstrainedResourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, constrainedResourceName string, constrainedResource datacollaboration.BasicConstrainedResource, force *bool) (result datacollaboration.ConstrainedResourcesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, constrainedResourceName string) (result datacollaboration.ConstrainedResourcesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, constrainedResourceName string) (result datacollaboration.ConstrainedResourceModel, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ConstrainedResourceListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ConstrainedResourceListIterator, err error)
}

var _ ConstrainedResourcesClientAPI = (*datacollaboration.ConstrainedResourcesClient)(nil)

// DataAssetsClientAPI contains the set of methods on the DataAssetsClient type.
type DataAssetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, dataAsset datacollaboration.DataAsset) (result datacollaboration.DataAsset, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string) (result datacollaboration.DataAssetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string) (result datacollaboration.DataAsset, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.DataAssetListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.DataAssetListIterator, err error)
}

var _ DataAssetsClientAPI = (*datacollaboration.DataAssetsClient)(nil)

// DataSetsClientAPI contains the set of methods on the DataSetsClient type.
type DataSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, dataSetCategory datacollaboration.DataSetCategoryType, dataSet datacollaboration.BasicDataSet) (result datacollaboration.DataSetModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, dataSetCategory datacollaboration.DataSetCategoryType) (result datacollaboration.DataSetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, dataSetCategory datacollaboration.DataSetCategoryType) (result datacollaboration.DataSetModel, err error)
	ListByDataAsset(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, skipToken string) (result datacollaboration.DataSetListPage, err error)
	ListByDataAssetComplete(ctx context.Context, resourceGroupName string, workspaceName string, dataAssetName string, skipToken string) (result datacollaboration.DataSetListIterator, err error)
}

var _ DataSetsClientAPI = (*datacollaboration.DataSetsClient)(nil)

// ManagedPrivateEndpointsClientAPI contains the set of methods on the ManagedPrivateEndpointsClient type.
type ManagedPrivateEndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, managedPrivateEndpointName string, managedPrivateEndpoint datacollaboration.ManagedPrivateEndpoint) (result datacollaboration.ManagedPrivateEndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, managedPrivateEndpointName string) (result datacollaboration.ManagedPrivateEndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, managedPrivateEndpointName string) (result datacollaboration.ManagedPrivateEndpoint, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ManagedPrivateEndpointListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ManagedPrivateEndpointListIterator, err error)
}

var _ ManagedPrivateEndpointsClientAPI = (*datacollaboration.ManagedPrivateEndpointsClient)(nil)

// PipelineRunsClientAPI contains the set of methods on the PipelineRunsClient type.
type PipelineRunsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (result datacollaboration.PipelineRunsCancelFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string) (result datacollaboration.PipelineRun, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.PipelineRunListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.PipelineRunListIterator, err error)
}

var _ PipelineRunsClientAPI = (*datacollaboration.PipelineRunsClient)(nil)

// PipelineStepRunsClientAPI contains the set of methods on the PipelineStepRunsClient type.
type PipelineStepRunsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string, pipelineStepRunName string) (result datacollaboration.PipelineStepRun, err error)
	ListByPipelineRun(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string, skipToken string) (result datacollaboration.PipelineStepRunListPage, err error)
	ListByPipelineRunComplete(ctx context.Context, resourceGroupName string, workspaceName string, pipelineRunName string, skipToken string) (result datacollaboration.PipelineStepRunListIterator, err error)
}

var _ PipelineStepRunsClientAPI = (*datacollaboration.PipelineStepRunsClient)(nil)

// PipelinesClientAPI contains the set of methods on the PipelinesClient type.
type PipelinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipeline datacollaboration.Pipeline) (result datacollaboration.Pipeline, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string) (result datacollaboration.PipelinesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string) (result datacollaboration.Pipeline, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.PipelineListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.PipelineListIterator, err error)
	Run(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineRunParameters datacollaboration.PipelineRunParameters) (result datacollaboration.PipelinesRunFuture, err error)
}

var _ PipelinesClientAPI = (*datacollaboration.PipelinesClient)(nil)

// PipelineStepsClientAPI contains the set of methods on the PipelineStepsClient type.
type PipelineStepsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineStepName string, pipelineStep datacollaboration.BasicPipelineStep) (result datacollaboration.PipelineStepModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineStepName string) (result datacollaboration.PipelineStepsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineStepName string) (result datacollaboration.PipelineStepModel, err error)
	ListByPipeline(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, skipToken string) (result datacollaboration.PipelineStepListPage, err error)
	ListByPipelineComplete(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, skipToken string) (result datacollaboration.PipelineStepListIterator, err error)
}

var _ PipelineStepsClientAPI = (*datacollaboration.PipelineStepsClient)(nil)

// ProposalsClientAPI contains the set of methods on the ProposalsClient type.
type ProposalsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, proposal datacollaboration.Proposal) (result datacollaboration.ProposalsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string) (result datacollaboration.ProposalsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string) (result datacollaboration.Proposal, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ProposalListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ProposalListIterator, err error)
	Revoke(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string) (result datacollaboration.Proposal, err error)
	Sign(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, proposalSignature datacollaboration.ProposalSignature) (result datacollaboration.ProposalsSignFuture, err error)
}

var _ ProposalsClientAPI = (*datacollaboration.ProposalsClient)(nil)

// DataAssetReferencesClientAPI contains the set of methods on the DataAssetReferencesClient type.
type DataAssetReferencesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string, dataAssetReference datacollaboration.DataAssetReference) (result datacollaboration.DataAssetReference, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.DataAssetReferencesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.DataAssetReference, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.DataAssetReferenceListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.DataAssetReferenceListIterator, err error)
	Resolve(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.DataAsset, err error)
}

var _ DataAssetReferencesClientAPI = (*datacollaboration.DataAssetReferencesClient)(nil)

// EntitlementsClientAPI contains the set of methods on the EntitlementsClient type.
type EntitlementsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, entitlement datacollaboration.Entitlement) (result datacollaboration.Entitlement, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string) (result datacollaboration.EntitlementsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string) (result datacollaboration.Entitlement, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string, filter string, orderby string) (result datacollaboration.EntitlementListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string, filter string, orderby string) (result datacollaboration.EntitlementListIterator, err error)
}

var _ EntitlementsClientAPI = (*datacollaboration.EntitlementsClient)(nil)

// ConstraintsClientAPI contains the set of methods on the ConstraintsClient type.
type ConstraintsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, constraintName string, constraint datacollaboration.BasicConstraint) (result datacollaboration.ConstraintModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, constraintName string) (result datacollaboration.ConstraintsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, constraintName string) (result datacollaboration.ConstraintModel, err error)
	ListByEntitlement(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, skipToken string, filter string, orderby string) (result datacollaboration.ConstraintListPage, err error)
	ListByEntitlementComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, skipToken string, filter string, orderby string) (result datacollaboration.ConstraintListIterator, err error)
}

var _ ConstraintsClientAPI = (*datacollaboration.ConstraintsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, policyName string, policy datacollaboration.BasicPolicy) (result datacollaboration.PolicyModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, policyName string) (result datacollaboration.PoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, policyName string) (result datacollaboration.PolicyModel, err error)
	ListByEntitlement(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, skipToken string, filter string, orderby string) (result datacollaboration.PolicyListPage, err error)
	ListByEntitlementComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, entitlementName string, skipToken string, filter string, orderby string) (result datacollaboration.PolicyListIterator, err error)
}

var _ PoliciesClientAPI = (*datacollaboration.PoliciesClient)(nil)

// InvitationsClientAPI contains the set of methods on the InvitationsClient type.
type InvitationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string, invitation datacollaboration.Invitation) (result datacollaboration.Invitation, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (result datacollaboration.OperationResponse, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (result datacollaboration.Invitation, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.InvitationListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.InvitationListIterator, err error)
}

var _ InvitationsClientAPI = (*datacollaboration.InvitationsClient)(nil)

// ParticipantsClientAPI contains the set of methods on the ParticipantsClient type.
type ParticipantsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string, participant datacollaboration.Participant) (result datacollaboration.Participant, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (result datacollaboration.ParticipantsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, participantName string) (result datacollaboration.Participant, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.ParticipantListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.ParticipantListIterator, err error)
}

var _ ParticipantsClientAPI = (*datacollaboration.ParticipantsClient)(nil)

// ScriptReferencesClientAPI contains the set of methods on the ScriptReferencesClient type.
type ScriptReferencesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string, scriptReference datacollaboration.ScriptReference) (result datacollaboration.ScriptReference, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.ScriptReferencesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.ScriptReference, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.ScriptReferenceListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.ScriptReferenceListIterator, err error)
	Resolve(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.ScriptModel, err error)
}

var _ ScriptReferencesClientAPI = (*datacollaboration.ScriptReferencesClient)(nil)

// VirtualOutputReferencesClientAPI contains the set of methods on the VirtualOutputReferencesClient type.
type VirtualOutputReferencesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, referenceName string) (result datacollaboration.VirtualOutputReference, err error)
	ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.VirtualOutputReferenceListPage, err error)
	ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result datacollaboration.VirtualOutputReferenceListIterator, err error)
}

var _ VirtualOutputReferencesClientAPI = (*datacollaboration.VirtualOutputReferencesClient)(nil)

// ResourceReferencesClientAPI contains the set of methods on the ResourceReferencesClient type.
type ResourceReferencesClientAPI interface {
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ResourceReferenceListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ResourceReferenceListIterator, err error)
}

var _ ResourceReferencesClientAPI = (*datacollaboration.ResourceReferencesClient)(nil)

// ScriptsClientAPI contains the set of methods on the ScriptsClient type.
type ScriptsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string, script datacollaboration.BasicScript) (result datacollaboration.ScriptModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string) (result datacollaboration.ScriptsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string) (result datacollaboration.ScriptModel, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ScriptListPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, skipToken string, filter string, orderby string) (result datacollaboration.ScriptListIterator, err error)
}

var _ ScriptsClientAPI = (*datacollaboration.ScriptsClient)(nil)

// ScriptRevisionsClientAPI contains the set of methods on the ScriptRevisionsClient type.
type ScriptRevisionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string, revision int32) (result datacollaboration.ScriptModel, err error)
	ListByScript(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string, skipToken string) (result datacollaboration.ScriptListPage, err error)
	ListByScriptComplete(ctx context.Context, resourceGroupName string, workspaceName string, scriptName string, skipToken string) (result datacollaboration.ScriptListIterator, err error)
}

var _ ScriptRevisionsClientAPI = (*datacollaboration.ScriptRevisionsClient)(nil)

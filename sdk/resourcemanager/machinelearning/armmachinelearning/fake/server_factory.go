//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armmachinelearning.ClientFactory type.
type ServerFactory struct {
	BatchDeploymentsServer              BatchDeploymentsServer
	BatchEndpointsServer                BatchEndpointsServer
	CapacityReservationGroupsServer     CapacityReservationGroupsServer
	CodeContainersServer                CodeContainersServer
	CodeVersionsServer                  CodeVersionsServer
	ComponentContainersServer           ComponentContainersServer
	ComponentVersionsServer             ComponentVersionsServer
	ComputeServer                       ComputeServer
	DataContainersServer                DataContainersServer
	DataVersionsServer                  DataVersionsServer
	DatastoresServer                    DatastoresServer
	EnvironmentContainersServer         EnvironmentContainersServer
	EnvironmentVersionsServer           EnvironmentVersionsServer
	FeaturesServer                      FeaturesServer
	FeaturesetContainersServer          FeaturesetContainersServer
	FeaturesetVersionsServer            FeaturesetVersionsServer
	FeaturestoreEntityContainersServer  FeaturestoreEntityContainersServer
	FeaturestoreEntityVersionsServer    FeaturestoreEntityVersionsServer
	IndexesServer                       IndexesServer
	IndexesVersionsServer               IndexesVersionsServer
	InferenceEndpointsServer            InferenceEndpointsServer
	InferenceGroupsServer               InferenceGroupsServer
	InferencePoolsServer                InferencePoolsServer
	JobsServer                          JobsServer
	LabelingJobsServer                  LabelingJobsServer
	ManagedNetworkProvisionsServer      ManagedNetworkProvisionsServer
	ManagedNetworkSettingsRuleServer    ManagedNetworkSettingsRuleServer
	ModelContainersServer               ModelContainersServer
	ModelVersionsServer                 ModelVersionsServer
	OnlineDeploymentsServer             OnlineDeploymentsServer
	OnlineEndpointsServer               OnlineEndpointsServer
	OperationsServer                    OperationsServer
	PrivateEndpointConnectionsServer    PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer          PrivateLinkResourcesServer
	QuotasServer                        QuotasServer
	RegistriesServer                    RegistriesServer
	RegistryCodeContainersServer        RegistryCodeContainersServer
	RegistryCodeVersionsServer          RegistryCodeVersionsServer
	RegistryComponentContainersServer   RegistryComponentContainersServer
	RegistryComponentVersionsServer     RegistryComponentVersionsServer
	RegistryDataContainersServer        RegistryDataContainersServer
	RegistryDataReferencesServer        RegistryDataReferencesServer
	RegistryDataVersionsServer          RegistryDataVersionsServer
	RegistryEnvironmentContainersServer RegistryEnvironmentContainersServer
	RegistryEnvironmentVersionsServer   RegistryEnvironmentVersionsServer
	RegistryModelContainersServer       RegistryModelContainersServer
	RegistryModelVersionsServer         RegistryModelVersionsServer
	SchedulesServer                     SchedulesServer
	ServerlessEndpointsServer           ServerlessEndpointsServer
	UsagesServer                        UsagesServer
	VirtualMachineSizesServer           VirtualMachineSizesServer
	WorkspaceConnectionsServer          WorkspaceConnectionsServer
	WorkspaceFeaturesServer             WorkspaceFeaturesServer
	WorkspacesServer                    WorkspacesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmachinelearning.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmachinelearning.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                   *ServerFactory
	trMu                                  sync.Mutex
	trBatchDeploymentsServer              *BatchDeploymentsServerTransport
	trBatchEndpointsServer                *BatchEndpointsServerTransport
	trCapacityReservationGroupsServer     *CapacityReservationGroupsServerTransport
	trCodeContainersServer                *CodeContainersServerTransport
	trCodeVersionsServer                  *CodeVersionsServerTransport
	trComponentContainersServer           *ComponentContainersServerTransport
	trComponentVersionsServer             *ComponentVersionsServerTransport
	trComputeServer                       *ComputeServerTransport
	trDataContainersServer                *DataContainersServerTransport
	trDataVersionsServer                  *DataVersionsServerTransport
	trDatastoresServer                    *DatastoresServerTransport
	trEnvironmentContainersServer         *EnvironmentContainersServerTransport
	trEnvironmentVersionsServer           *EnvironmentVersionsServerTransport
	trFeaturesServer                      *FeaturesServerTransport
	trFeaturesetContainersServer          *FeaturesetContainersServerTransport
	trFeaturesetVersionsServer            *FeaturesetVersionsServerTransport
	trFeaturestoreEntityContainersServer  *FeaturestoreEntityContainersServerTransport
	trFeaturestoreEntityVersionsServer    *FeaturestoreEntityVersionsServerTransport
	trIndexesServer                       *IndexesServerTransport
	trIndexesVersionsServer               *IndexesVersionsServerTransport
	trInferenceEndpointsServer            *InferenceEndpointsServerTransport
	trInferenceGroupsServer               *InferenceGroupsServerTransport
	trInferencePoolsServer                *InferencePoolsServerTransport
	trJobsServer                          *JobsServerTransport
	trLabelingJobsServer                  *LabelingJobsServerTransport
	trManagedNetworkProvisionsServer      *ManagedNetworkProvisionsServerTransport
	trManagedNetworkSettingsRuleServer    *ManagedNetworkSettingsRuleServerTransport
	trModelContainersServer               *ModelContainersServerTransport
	trModelVersionsServer                 *ModelVersionsServerTransport
	trOnlineDeploymentsServer             *OnlineDeploymentsServerTransport
	trOnlineEndpointsServer               *OnlineEndpointsServerTransport
	trOperationsServer                    *OperationsServerTransport
	trPrivateEndpointConnectionsServer    *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer          *PrivateLinkResourcesServerTransport
	trQuotasServer                        *QuotasServerTransport
	trRegistriesServer                    *RegistriesServerTransport
	trRegistryCodeContainersServer        *RegistryCodeContainersServerTransport
	trRegistryCodeVersionsServer          *RegistryCodeVersionsServerTransport
	trRegistryComponentContainersServer   *RegistryComponentContainersServerTransport
	trRegistryComponentVersionsServer     *RegistryComponentVersionsServerTransport
	trRegistryDataContainersServer        *RegistryDataContainersServerTransport
	trRegistryDataReferencesServer        *RegistryDataReferencesServerTransport
	trRegistryDataVersionsServer          *RegistryDataVersionsServerTransport
	trRegistryEnvironmentContainersServer *RegistryEnvironmentContainersServerTransport
	trRegistryEnvironmentVersionsServer   *RegistryEnvironmentVersionsServerTransport
	trRegistryModelContainersServer       *RegistryModelContainersServerTransport
	trRegistryModelVersionsServer         *RegistryModelVersionsServerTransport
	trSchedulesServer                     *SchedulesServerTransport
	trServerlessEndpointsServer           *ServerlessEndpointsServerTransport
	trUsagesServer                        *UsagesServerTransport
	trVirtualMachineSizesServer           *VirtualMachineSizesServerTransport
	trWorkspaceConnectionsServer          *WorkspaceConnectionsServerTransport
	trWorkspaceFeaturesServer             *WorkspaceFeaturesServerTransport
	trWorkspacesServer                    *WorkspacesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "BatchDeploymentsClient":
		initServer(s, &s.trBatchDeploymentsServer, func() *BatchDeploymentsServerTransport {
			return NewBatchDeploymentsServerTransport(&s.srv.BatchDeploymentsServer)
		})
		resp, err = s.trBatchDeploymentsServer.Do(req)
	case "BatchEndpointsClient":
		initServer(s, &s.trBatchEndpointsServer, func() *BatchEndpointsServerTransport {
			return NewBatchEndpointsServerTransport(&s.srv.BatchEndpointsServer)
		})
		resp, err = s.trBatchEndpointsServer.Do(req)
	case "CapacityReservationGroupsClient":
		initServer(s, &s.trCapacityReservationGroupsServer, func() *CapacityReservationGroupsServerTransport {
			return NewCapacityReservationGroupsServerTransport(&s.srv.CapacityReservationGroupsServer)
		})
		resp, err = s.trCapacityReservationGroupsServer.Do(req)
	case "CodeContainersClient":
		initServer(s, &s.trCodeContainersServer, func() *CodeContainersServerTransport {
			return NewCodeContainersServerTransport(&s.srv.CodeContainersServer)
		})
		resp, err = s.trCodeContainersServer.Do(req)
	case "CodeVersionsClient":
		initServer(s, &s.trCodeVersionsServer, func() *CodeVersionsServerTransport { return NewCodeVersionsServerTransport(&s.srv.CodeVersionsServer) })
		resp, err = s.trCodeVersionsServer.Do(req)
	case "ComponentContainersClient":
		initServer(s, &s.trComponentContainersServer, func() *ComponentContainersServerTransport {
			return NewComponentContainersServerTransport(&s.srv.ComponentContainersServer)
		})
		resp, err = s.trComponentContainersServer.Do(req)
	case "ComponentVersionsClient":
		initServer(s, &s.trComponentVersionsServer, func() *ComponentVersionsServerTransport {
			return NewComponentVersionsServerTransport(&s.srv.ComponentVersionsServer)
		})
		resp, err = s.trComponentVersionsServer.Do(req)
	case "ComputeClient":
		initServer(s, &s.trComputeServer, func() *ComputeServerTransport { return NewComputeServerTransport(&s.srv.ComputeServer) })
		resp, err = s.trComputeServer.Do(req)
	case "DataContainersClient":
		initServer(s, &s.trDataContainersServer, func() *DataContainersServerTransport {
			return NewDataContainersServerTransport(&s.srv.DataContainersServer)
		})
		resp, err = s.trDataContainersServer.Do(req)
	case "DataVersionsClient":
		initServer(s, &s.trDataVersionsServer, func() *DataVersionsServerTransport { return NewDataVersionsServerTransport(&s.srv.DataVersionsServer) })
		resp, err = s.trDataVersionsServer.Do(req)
	case "DatastoresClient":
		initServer(s, &s.trDatastoresServer, func() *DatastoresServerTransport { return NewDatastoresServerTransport(&s.srv.DatastoresServer) })
		resp, err = s.trDatastoresServer.Do(req)
	case "EnvironmentContainersClient":
		initServer(s, &s.trEnvironmentContainersServer, func() *EnvironmentContainersServerTransport {
			return NewEnvironmentContainersServerTransport(&s.srv.EnvironmentContainersServer)
		})
		resp, err = s.trEnvironmentContainersServer.Do(req)
	case "EnvironmentVersionsClient":
		initServer(s, &s.trEnvironmentVersionsServer, func() *EnvironmentVersionsServerTransport {
			return NewEnvironmentVersionsServerTransport(&s.srv.EnvironmentVersionsServer)
		})
		resp, err = s.trEnvironmentVersionsServer.Do(req)
	case "FeaturesClient":
		initServer(s, &s.trFeaturesServer, func() *FeaturesServerTransport { return NewFeaturesServerTransport(&s.srv.FeaturesServer) })
		resp, err = s.trFeaturesServer.Do(req)
	case "FeaturesetContainersClient":
		initServer(s, &s.trFeaturesetContainersServer, func() *FeaturesetContainersServerTransport {
			return NewFeaturesetContainersServerTransport(&s.srv.FeaturesetContainersServer)
		})
		resp, err = s.trFeaturesetContainersServer.Do(req)
	case "FeaturesetVersionsClient":
		initServer(s, &s.trFeaturesetVersionsServer, func() *FeaturesetVersionsServerTransport {
			return NewFeaturesetVersionsServerTransport(&s.srv.FeaturesetVersionsServer)
		})
		resp, err = s.trFeaturesetVersionsServer.Do(req)
	case "FeaturestoreEntityContainersClient":
		initServer(s, &s.trFeaturestoreEntityContainersServer, func() *FeaturestoreEntityContainersServerTransport {
			return NewFeaturestoreEntityContainersServerTransport(&s.srv.FeaturestoreEntityContainersServer)
		})
		resp, err = s.trFeaturestoreEntityContainersServer.Do(req)
	case "FeaturestoreEntityVersionsClient":
		initServer(s, &s.trFeaturestoreEntityVersionsServer, func() *FeaturestoreEntityVersionsServerTransport {
			return NewFeaturestoreEntityVersionsServerTransport(&s.srv.FeaturestoreEntityVersionsServer)
		})
		resp, err = s.trFeaturestoreEntityVersionsServer.Do(req)
	case "IndexesClient":
		initServer(s, &s.trIndexesServer, func() *IndexesServerTransport { return NewIndexesServerTransport(&s.srv.IndexesServer) })
		resp, err = s.trIndexesServer.Do(req)
	case "IndexesVersionsClient":
		initServer(s, &s.trIndexesVersionsServer, func() *IndexesVersionsServerTransport {
			return NewIndexesVersionsServerTransport(&s.srv.IndexesVersionsServer)
		})
		resp, err = s.trIndexesVersionsServer.Do(req)
	case "InferenceEndpointsClient":
		initServer(s, &s.trInferenceEndpointsServer, func() *InferenceEndpointsServerTransport {
			return NewInferenceEndpointsServerTransport(&s.srv.InferenceEndpointsServer)
		})
		resp, err = s.trInferenceEndpointsServer.Do(req)
	case "InferenceGroupsClient":
		initServer(s, &s.trInferenceGroupsServer, func() *InferenceGroupsServerTransport {
			return NewInferenceGroupsServerTransport(&s.srv.InferenceGroupsServer)
		})
		resp, err = s.trInferenceGroupsServer.Do(req)
	case "InferencePoolsClient":
		initServer(s, &s.trInferencePoolsServer, func() *InferencePoolsServerTransport {
			return NewInferencePoolsServerTransport(&s.srv.InferencePoolsServer)
		})
		resp, err = s.trInferencePoolsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "LabelingJobsClient":
		initServer(s, &s.trLabelingJobsServer, func() *LabelingJobsServerTransport { return NewLabelingJobsServerTransport(&s.srv.LabelingJobsServer) })
		resp, err = s.trLabelingJobsServer.Do(req)
	case "ManagedNetworkProvisionsClient":
		initServer(s, &s.trManagedNetworkProvisionsServer, func() *ManagedNetworkProvisionsServerTransport {
			return NewManagedNetworkProvisionsServerTransport(&s.srv.ManagedNetworkProvisionsServer)
		})
		resp, err = s.trManagedNetworkProvisionsServer.Do(req)
	case "ManagedNetworkSettingsRuleClient":
		initServer(s, &s.trManagedNetworkSettingsRuleServer, func() *ManagedNetworkSettingsRuleServerTransport {
			return NewManagedNetworkSettingsRuleServerTransport(&s.srv.ManagedNetworkSettingsRuleServer)
		})
		resp, err = s.trManagedNetworkSettingsRuleServer.Do(req)
	case "ModelContainersClient":
		initServer(s, &s.trModelContainersServer, func() *ModelContainersServerTransport {
			return NewModelContainersServerTransport(&s.srv.ModelContainersServer)
		})
		resp, err = s.trModelContainersServer.Do(req)
	case "ModelVersionsClient":
		initServer(s, &s.trModelVersionsServer, func() *ModelVersionsServerTransport {
			return NewModelVersionsServerTransport(&s.srv.ModelVersionsServer)
		})
		resp, err = s.trModelVersionsServer.Do(req)
	case "OnlineDeploymentsClient":
		initServer(s, &s.trOnlineDeploymentsServer, func() *OnlineDeploymentsServerTransport {
			return NewOnlineDeploymentsServerTransport(&s.srv.OnlineDeploymentsServer)
		})
		resp, err = s.trOnlineDeploymentsServer.Do(req)
	case "OnlineEndpointsClient":
		initServer(s, &s.trOnlineEndpointsServer, func() *OnlineEndpointsServerTransport {
			return NewOnlineEndpointsServerTransport(&s.srv.OnlineEndpointsServer)
		})
		resp, err = s.trOnlineEndpointsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "QuotasClient":
		initServer(s, &s.trQuotasServer, func() *QuotasServerTransport { return NewQuotasServerTransport(&s.srv.QuotasServer) })
		resp, err = s.trQuotasServer.Do(req)
	case "RegistriesClient":
		initServer(s, &s.trRegistriesServer, func() *RegistriesServerTransport { return NewRegistriesServerTransport(&s.srv.RegistriesServer) })
		resp, err = s.trRegistriesServer.Do(req)
	case "RegistryCodeContainersClient":
		initServer(s, &s.trRegistryCodeContainersServer, func() *RegistryCodeContainersServerTransport {
			return NewRegistryCodeContainersServerTransport(&s.srv.RegistryCodeContainersServer)
		})
		resp, err = s.trRegistryCodeContainersServer.Do(req)
	case "RegistryCodeVersionsClient":
		initServer(s, &s.trRegistryCodeVersionsServer, func() *RegistryCodeVersionsServerTransport {
			return NewRegistryCodeVersionsServerTransport(&s.srv.RegistryCodeVersionsServer)
		})
		resp, err = s.trRegistryCodeVersionsServer.Do(req)
	case "RegistryComponentContainersClient":
		initServer(s, &s.trRegistryComponentContainersServer, func() *RegistryComponentContainersServerTransport {
			return NewRegistryComponentContainersServerTransport(&s.srv.RegistryComponentContainersServer)
		})
		resp, err = s.trRegistryComponentContainersServer.Do(req)
	case "RegistryComponentVersionsClient":
		initServer(s, &s.trRegistryComponentVersionsServer, func() *RegistryComponentVersionsServerTransport {
			return NewRegistryComponentVersionsServerTransport(&s.srv.RegistryComponentVersionsServer)
		})
		resp, err = s.trRegistryComponentVersionsServer.Do(req)
	case "RegistryDataContainersClient":
		initServer(s, &s.trRegistryDataContainersServer, func() *RegistryDataContainersServerTransport {
			return NewRegistryDataContainersServerTransport(&s.srv.RegistryDataContainersServer)
		})
		resp, err = s.trRegistryDataContainersServer.Do(req)
	case "RegistryDataReferencesClient":
		initServer(s, &s.trRegistryDataReferencesServer, func() *RegistryDataReferencesServerTransport {
			return NewRegistryDataReferencesServerTransport(&s.srv.RegistryDataReferencesServer)
		})
		resp, err = s.trRegistryDataReferencesServer.Do(req)
	case "RegistryDataVersionsClient":
		initServer(s, &s.trRegistryDataVersionsServer, func() *RegistryDataVersionsServerTransport {
			return NewRegistryDataVersionsServerTransport(&s.srv.RegistryDataVersionsServer)
		})
		resp, err = s.trRegistryDataVersionsServer.Do(req)
	case "RegistryEnvironmentContainersClient":
		initServer(s, &s.trRegistryEnvironmentContainersServer, func() *RegistryEnvironmentContainersServerTransport {
			return NewRegistryEnvironmentContainersServerTransport(&s.srv.RegistryEnvironmentContainersServer)
		})
		resp, err = s.trRegistryEnvironmentContainersServer.Do(req)
	case "RegistryEnvironmentVersionsClient":
		initServer(s, &s.trRegistryEnvironmentVersionsServer, func() *RegistryEnvironmentVersionsServerTransport {
			return NewRegistryEnvironmentVersionsServerTransport(&s.srv.RegistryEnvironmentVersionsServer)
		})
		resp, err = s.trRegistryEnvironmentVersionsServer.Do(req)
	case "RegistryModelContainersClient":
		initServer(s, &s.trRegistryModelContainersServer, func() *RegistryModelContainersServerTransport {
			return NewRegistryModelContainersServerTransport(&s.srv.RegistryModelContainersServer)
		})
		resp, err = s.trRegistryModelContainersServer.Do(req)
	case "RegistryModelVersionsClient":
		initServer(s, &s.trRegistryModelVersionsServer, func() *RegistryModelVersionsServerTransport {
			return NewRegistryModelVersionsServerTransport(&s.srv.RegistryModelVersionsServer)
		})
		resp, err = s.trRegistryModelVersionsServer.Do(req)
	case "SchedulesClient":
		initServer(s, &s.trSchedulesServer, func() *SchedulesServerTransport { return NewSchedulesServerTransport(&s.srv.SchedulesServer) })
		resp, err = s.trSchedulesServer.Do(req)
	case "ServerlessEndpointsClient":
		initServer(s, &s.trServerlessEndpointsServer, func() *ServerlessEndpointsServerTransport {
			return NewServerlessEndpointsServerTransport(&s.srv.ServerlessEndpointsServer)
		})
		resp, err = s.trServerlessEndpointsServer.Do(req)
	case "UsagesClient":
		initServer(s, &s.trUsagesServer, func() *UsagesServerTransport { return NewUsagesServerTransport(&s.srv.UsagesServer) })
		resp, err = s.trUsagesServer.Do(req)
	case "VirtualMachineSizesClient":
		initServer(s, &s.trVirtualMachineSizesServer, func() *VirtualMachineSizesServerTransport {
			return NewVirtualMachineSizesServerTransport(&s.srv.VirtualMachineSizesServer)
		})
		resp, err = s.trVirtualMachineSizesServer.Do(req)
	case "WorkspaceConnectionsClient":
		initServer(s, &s.trWorkspaceConnectionsServer, func() *WorkspaceConnectionsServerTransport {
			return NewWorkspaceConnectionsServerTransport(&s.srv.WorkspaceConnectionsServer)
		})
		resp, err = s.trWorkspaceConnectionsServer.Do(req)
	case "WorkspaceFeaturesClient":
		initServer(s, &s.trWorkspaceFeaturesServer, func() *WorkspaceFeaturesServerTransport {
			return NewWorkspaceFeaturesServerTransport(&s.srv.WorkspaceFeaturesServer)
		})
		resp, err = s.trWorkspaceFeaturesServer.Do(req)
	case "WorkspacesClient":
		initServer(s, &s.trWorkspacesServer, func() *WorkspacesServerTransport { return NewWorkspacesServerTransport(&s.srv.WorkspacesServer) })
		resp, err = s.trWorkspacesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}

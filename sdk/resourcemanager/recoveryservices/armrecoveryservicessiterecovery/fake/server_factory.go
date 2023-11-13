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

// ServerFactory is a fake server for instances of the armrecoveryservicessiterecovery.ClientFactory type.
type ServerFactory struct {
	MigrationRecoveryPointsServer                  MigrationRecoveryPointsServer
	OperationsServer                               OperationsServer
	RecoveryPointsServer                           RecoveryPointsServer
	ReplicationAlertSettingsServer                 ReplicationAlertSettingsServer
	ReplicationAppliancesServer                    ReplicationAppliancesServer
	ReplicationEligibilityResultsServer            ReplicationEligibilityResultsServer
	ReplicationEventsServer                        ReplicationEventsServer
	ReplicationFabricsServer                       ReplicationFabricsServer
	ReplicationInfrastructureServer                ReplicationInfrastructureServer
	ReplicationJobsServer                          ReplicationJobsServer
	ReplicationLogicalNetworksServer               ReplicationLogicalNetworksServer
	ReplicationMigrationItemsServer                ReplicationMigrationItemsServer
	ReplicationNetworkMappingsServer               ReplicationNetworkMappingsServer
	ReplicationNetworksServer                      ReplicationNetworksServer
	ReplicationPoliciesServer                      ReplicationPoliciesServer
	ReplicationProtectableItemsServer              ReplicationProtectableItemsServer
	ReplicationProtectedItemsServer                ReplicationProtectedItemsServer
	ReplicationProtectionContainerMappingsServer   ReplicationProtectionContainerMappingsServer
	ReplicationProtectionContainersServer          ReplicationProtectionContainersServer
	ReplicationProtectionIntentsServer             ReplicationProtectionIntentsServer
	ReplicationRecoveryPlansServer                 ReplicationRecoveryPlansServer
	ReplicationRecoveryServicesProvidersServer     ReplicationRecoveryServicesProvidersServer
	ReplicationStorageClassificationMappingsServer ReplicationStorageClassificationMappingsServer
	ReplicationStorageClassificationsServer        ReplicationStorageClassificationsServer
	ReplicationVaultHealthServer                   ReplicationVaultHealthServer
	ReplicationVaultSettingServer                  ReplicationVaultSettingServer
	ReplicationvCentersServer                      ReplicationvCentersServer
	SupportedOperatingSystemsServer                SupportedOperatingSystemsServer
	TargetComputeSizesServer                       TargetComputeSizesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armrecoveryservicessiterecovery.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armrecoveryservicessiterecovery.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                              *ServerFactory
	trMu                                             sync.Mutex
	trMigrationRecoveryPointsServer                  *MigrationRecoveryPointsServerTransport
	trOperationsServer                               *OperationsServerTransport
	trRecoveryPointsServer                           *RecoveryPointsServerTransport
	trReplicationAlertSettingsServer                 *ReplicationAlertSettingsServerTransport
	trReplicationAppliancesServer                    *ReplicationAppliancesServerTransport
	trReplicationEligibilityResultsServer            *ReplicationEligibilityResultsServerTransport
	trReplicationEventsServer                        *ReplicationEventsServerTransport
	trReplicationFabricsServer                       *ReplicationFabricsServerTransport
	trReplicationInfrastructureServer                *ReplicationInfrastructureServerTransport
	trReplicationJobsServer                          *ReplicationJobsServerTransport
	trReplicationLogicalNetworksServer               *ReplicationLogicalNetworksServerTransport
	trReplicationMigrationItemsServer                *ReplicationMigrationItemsServerTransport
	trReplicationNetworkMappingsServer               *ReplicationNetworkMappingsServerTransport
	trReplicationNetworksServer                      *ReplicationNetworksServerTransport
	trReplicationPoliciesServer                      *ReplicationPoliciesServerTransport
	trReplicationProtectableItemsServer              *ReplicationProtectableItemsServerTransport
	trReplicationProtectedItemsServer                *ReplicationProtectedItemsServerTransport
	trReplicationProtectionContainerMappingsServer   *ReplicationProtectionContainerMappingsServerTransport
	trReplicationProtectionContainersServer          *ReplicationProtectionContainersServerTransport
	trReplicationProtectionIntentsServer             *ReplicationProtectionIntentsServerTransport
	trReplicationRecoveryPlansServer                 *ReplicationRecoveryPlansServerTransport
	trReplicationRecoveryServicesProvidersServer     *ReplicationRecoveryServicesProvidersServerTransport
	trReplicationStorageClassificationMappingsServer *ReplicationStorageClassificationMappingsServerTransport
	trReplicationStorageClassificationsServer        *ReplicationStorageClassificationsServerTransport
	trReplicationVaultHealthServer                   *ReplicationVaultHealthServerTransport
	trReplicationVaultSettingServer                  *ReplicationVaultSettingServerTransport
	trReplicationvCentersServer                      *ReplicationvCentersServerTransport
	trSupportedOperatingSystemsServer                *SupportedOperatingSystemsServerTransport
	trTargetComputeSizesServer                       *TargetComputeSizesServerTransport
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
	case "MigrationRecoveryPointsClient":
		initServer(s, &s.trMigrationRecoveryPointsServer, func() *MigrationRecoveryPointsServerTransport {
			return NewMigrationRecoveryPointsServerTransport(&s.srv.MigrationRecoveryPointsServer)
		})
		resp, err = s.trMigrationRecoveryPointsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "RecoveryPointsClient":
		initServer(s, &s.trRecoveryPointsServer, func() *RecoveryPointsServerTransport {
			return NewRecoveryPointsServerTransport(&s.srv.RecoveryPointsServer)
		})
		resp, err = s.trRecoveryPointsServer.Do(req)
	case "ReplicationAlertSettingsClient":
		initServer(s, &s.trReplicationAlertSettingsServer, func() *ReplicationAlertSettingsServerTransport {
			return NewReplicationAlertSettingsServerTransport(&s.srv.ReplicationAlertSettingsServer)
		})
		resp, err = s.trReplicationAlertSettingsServer.Do(req)
	case "ReplicationAppliancesClient":
		initServer(s, &s.trReplicationAppliancesServer, func() *ReplicationAppliancesServerTransport {
			return NewReplicationAppliancesServerTransport(&s.srv.ReplicationAppliancesServer)
		})
		resp, err = s.trReplicationAppliancesServer.Do(req)
	case "ReplicationEligibilityResultsClient":
		initServer(s, &s.trReplicationEligibilityResultsServer, func() *ReplicationEligibilityResultsServerTransport {
			return NewReplicationEligibilityResultsServerTransport(&s.srv.ReplicationEligibilityResultsServer)
		})
		resp, err = s.trReplicationEligibilityResultsServer.Do(req)
	case "ReplicationEventsClient":
		initServer(s, &s.trReplicationEventsServer, func() *ReplicationEventsServerTransport {
			return NewReplicationEventsServerTransport(&s.srv.ReplicationEventsServer)
		})
		resp, err = s.trReplicationEventsServer.Do(req)
	case "ReplicationFabricsClient":
		initServer(s, &s.trReplicationFabricsServer, func() *ReplicationFabricsServerTransport {
			return NewReplicationFabricsServerTransport(&s.srv.ReplicationFabricsServer)
		})
		resp, err = s.trReplicationFabricsServer.Do(req)
	case "ReplicationInfrastructureClient":
		initServer(s, &s.trReplicationInfrastructureServer, func() *ReplicationInfrastructureServerTransport {
			return NewReplicationInfrastructureServerTransport(&s.srv.ReplicationInfrastructureServer)
		})
		resp, err = s.trReplicationInfrastructureServer.Do(req)
	case "ReplicationJobsClient":
		initServer(s, &s.trReplicationJobsServer, func() *ReplicationJobsServerTransport {
			return NewReplicationJobsServerTransport(&s.srv.ReplicationJobsServer)
		})
		resp, err = s.trReplicationJobsServer.Do(req)
	case "ReplicationLogicalNetworksClient":
		initServer(s, &s.trReplicationLogicalNetworksServer, func() *ReplicationLogicalNetworksServerTransport {
			return NewReplicationLogicalNetworksServerTransport(&s.srv.ReplicationLogicalNetworksServer)
		})
		resp, err = s.trReplicationLogicalNetworksServer.Do(req)
	case "ReplicationMigrationItemsClient":
		initServer(s, &s.trReplicationMigrationItemsServer, func() *ReplicationMigrationItemsServerTransport {
			return NewReplicationMigrationItemsServerTransport(&s.srv.ReplicationMigrationItemsServer)
		})
		resp, err = s.trReplicationMigrationItemsServer.Do(req)
	case "ReplicationNetworkMappingsClient":
		initServer(s, &s.trReplicationNetworkMappingsServer, func() *ReplicationNetworkMappingsServerTransport {
			return NewReplicationNetworkMappingsServerTransport(&s.srv.ReplicationNetworkMappingsServer)
		})
		resp, err = s.trReplicationNetworkMappingsServer.Do(req)
	case "ReplicationNetworksClient":
		initServer(s, &s.trReplicationNetworksServer, func() *ReplicationNetworksServerTransport {
			return NewReplicationNetworksServerTransport(&s.srv.ReplicationNetworksServer)
		})
		resp, err = s.trReplicationNetworksServer.Do(req)
	case "ReplicationPoliciesClient":
		initServer(s, &s.trReplicationPoliciesServer, func() *ReplicationPoliciesServerTransport {
			return NewReplicationPoliciesServerTransport(&s.srv.ReplicationPoliciesServer)
		})
		resp, err = s.trReplicationPoliciesServer.Do(req)
	case "ReplicationProtectableItemsClient":
		initServer(s, &s.trReplicationProtectableItemsServer, func() *ReplicationProtectableItemsServerTransport {
			return NewReplicationProtectableItemsServerTransport(&s.srv.ReplicationProtectableItemsServer)
		})
		resp, err = s.trReplicationProtectableItemsServer.Do(req)
	case "ReplicationProtectedItemsClient":
		initServer(s, &s.trReplicationProtectedItemsServer, func() *ReplicationProtectedItemsServerTransport {
			return NewReplicationProtectedItemsServerTransport(&s.srv.ReplicationProtectedItemsServer)
		})
		resp, err = s.trReplicationProtectedItemsServer.Do(req)
	case "ReplicationProtectionContainerMappingsClient":
		initServer(s, &s.trReplicationProtectionContainerMappingsServer, func() *ReplicationProtectionContainerMappingsServerTransport {
			return NewReplicationProtectionContainerMappingsServerTransport(&s.srv.ReplicationProtectionContainerMappingsServer)
		})
		resp, err = s.trReplicationProtectionContainerMappingsServer.Do(req)
	case "ReplicationProtectionContainersClient":
		initServer(s, &s.trReplicationProtectionContainersServer, func() *ReplicationProtectionContainersServerTransport {
			return NewReplicationProtectionContainersServerTransport(&s.srv.ReplicationProtectionContainersServer)
		})
		resp, err = s.trReplicationProtectionContainersServer.Do(req)
	case "ReplicationProtectionIntentsClient":
		initServer(s, &s.trReplicationProtectionIntentsServer, func() *ReplicationProtectionIntentsServerTransport {
			return NewReplicationProtectionIntentsServerTransport(&s.srv.ReplicationProtectionIntentsServer)
		})
		resp, err = s.trReplicationProtectionIntentsServer.Do(req)
	case "ReplicationRecoveryPlansClient":
		initServer(s, &s.trReplicationRecoveryPlansServer, func() *ReplicationRecoveryPlansServerTransport {
			return NewReplicationRecoveryPlansServerTransport(&s.srv.ReplicationRecoveryPlansServer)
		})
		resp, err = s.trReplicationRecoveryPlansServer.Do(req)
	case "ReplicationRecoveryServicesProvidersClient":
		initServer(s, &s.trReplicationRecoveryServicesProvidersServer, func() *ReplicationRecoveryServicesProvidersServerTransport {
			return NewReplicationRecoveryServicesProvidersServerTransport(&s.srv.ReplicationRecoveryServicesProvidersServer)
		})
		resp, err = s.trReplicationRecoveryServicesProvidersServer.Do(req)
	case "ReplicationStorageClassificationMappingsClient":
		initServer(s, &s.trReplicationStorageClassificationMappingsServer, func() *ReplicationStorageClassificationMappingsServerTransport {
			return NewReplicationStorageClassificationMappingsServerTransport(&s.srv.ReplicationStorageClassificationMappingsServer)
		})
		resp, err = s.trReplicationStorageClassificationMappingsServer.Do(req)
	case "ReplicationStorageClassificationsClient":
		initServer(s, &s.trReplicationStorageClassificationsServer, func() *ReplicationStorageClassificationsServerTransport {
			return NewReplicationStorageClassificationsServerTransport(&s.srv.ReplicationStorageClassificationsServer)
		})
		resp, err = s.trReplicationStorageClassificationsServer.Do(req)
	case "ReplicationVaultHealthClient":
		initServer(s, &s.trReplicationVaultHealthServer, func() *ReplicationVaultHealthServerTransport {
			return NewReplicationVaultHealthServerTransport(&s.srv.ReplicationVaultHealthServer)
		})
		resp, err = s.trReplicationVaultHealthServer.Do(req)
	case "ReplicationVaultSettingClient":
		initServer(s, &s.trReplicationVaultSettingServer, func() *ReplicationVaultSettingServerTransport {
			return NewReplicationVaultSettingServerTransport(&s.srv.ReplicationVaultSettingServer)
		})
		resp, err = s.trReplicationVaultSettingServer.Do(req)
	case "ReplicationvCentersClient":
		initServer(s, &s.trReplicationvCentersServer, func() *ReplicationvCentersServerTransport {
			return NewReplicationvCentersServerTransport(&s.srv.ReplicationvCentersServer)
		})
		resp, err = s.trReplicationvCentersServer.Do(req)
	case "SupportedOperatingSystemsClient":
		initServer(s, &s.trSupportedOperatingSystemsServer, func() *SupportedOperatingSystemsServerTransport {
			return NewSupportedOperatingSystemsServerTransport(&s.srv.SupportedOperatingSystemsServer)
		})
		resp, err = s.trSupportedOperatingSystemsServer.Do(req)
	case "TargetComputeSizesClient":
		initServer(s, &s.trTargetComputeSizesServer, func() *TargetComputeSizesServerTransport {
			return NewTargetComputeSizesServerTransport(&s.srv.TargetComputeSizesServer)
		})
		resp, err = s.trTargetComputeSizesServer.Do(req)
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

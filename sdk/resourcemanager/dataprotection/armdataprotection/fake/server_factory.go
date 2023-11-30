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

// ServerFactory is a fake server for instances of the armdataprotection.ClientFactory type.
type ServerFactory struct {
	BackupInstancesServer                     BackupInstancesServer
	BackupPoliciesServer                      BackupPoliciesServer
	BackupVaultOperationResultsServer         BackupVaultOperationResultsServer
	BackupVaultsServer                        BackupVaultsServer
	Server                                    Server
	DeletedBackupInstancesServer              DeletedBackupInstancesServer
	DppResourceGuardProxyServer               DppResourceGuardProxyServer
	ExportJobsServer                          ExportJobsServer
	ExportJobsOperationResultServer           ExportJobsOperationResultServer
	FetchCrossRegionRestoreJobServer          FetchCrossRegionRestoreJobServer
	FetchCrossRegionRestoreJobsServer         FetchCrossRegionRestoreJobsServer
	FetchSecondaryRecoveryPointsServer        FetchSecondaryRecoveryPointsServer
	JobsServer                                JobsServer
	OperationResultServer                     OperationResultServer
	OperationStatusBackupVaultContextServer   OperationStatusBackupVaultContextServer
	OperationStatusServer                     OperationStatusServer
	OperationStatusResourceGroupContextServer OperationStatusResourceGroupContextServer
	OperationsServer                          OperationsServer
	RecoveryPointsServer                      RecoveryPointsServer
	ResourceGuardsServer                      ResourceGuardsServer
	RestorableTimeRangesServer                RestorableTimeRangesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdataprotection.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdataprotection.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                         *ServerFactory
	trMu                                        sync.Mutex
	trBackupInstancesServer                     *BackupInstancesServerTransport
	trBackupPoliciesServer                      *BackupPoliciesServerTransport
	trBackupVaultOperationResultsServer         *BackupVaultOperationResultsServerTransport
	trBackupVaultsServer                        *BackupVaultsServerTransport
	trServer                                    *ServerTransport
	trDeletedBackupInstancesServer              *DeletedBackupInstancesServerTransport
	trDppResourceGuardProxyServer               *DppResourceGuardProxyServerTransport
	trExportJobsServer                          *ExportJobsServerTransport
	trExportJobsOperationResultServer           *ExportJobsOperationResultServerTransport
	trFetchCrossRegionRestoreJobServer          *FetchCrossRegionRestoreJobServerTransport
	trFetchCrossRegionRestoreJobsServer         *FetchCrossRegionRestoreJobsServerTransport
	trFetchSecondaryRecoveryPointsServer        *FetchSecondaryRecoveryPointsServerTransport
	trJobsServer                                *JobsServerTransport
	trOperationResultServer                     *OperationResultServerTransport
	trOperationStatusBackupVaultContextServer   *OperationStatusBackupVaultContextServerTransport
	trOperationStatusServer                     *OperationStatusServerTransport
	trOperationStatusResourceGroupContextServer *OperationStatusResourceGroupContextServerTransport
	trOperationsServer                          *OperationsServerTransport
	trRecoveryPointsServer                      *RecoveryPointsServerTransport
	trResourceGuardsServer                      *ResourceGuardsServerTransport
	trRestorableTimeRangesServer                *RestorableTimeRangesServerTransport
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
	case "BackupInstancesClient":
		initServer(s, &s.trBackupInstancesServer, func() *BackupInstancesServerTransport {
			return NewBackupInstancesServerTransport(&s.srv.BackupInstancesServer)
		})
		resp, err = s.trBackupInstancesServer.Do(req)
	case "BackupPoliciesClient":
		initServer(s, &s.trBackupPoliciesServer, func() *BackupPoliciesServerTransport {
			return NewBackupPoliciesServerTransport(&s.srv.BackupPoliciesServer)
		})
		resp, err = s.trBackupPoliciesServer.Do(req)
	case "BackupVaultOperationResultsClient":
		initServer(s, &s.trBackupVaultOperationResultsServer, func() *BackupVaultOperationResultsServerTransport {
			return NewBackupVaultOperationResultsServerTransport(&s.srv.BackupVaultOperationResultsServer)
		})
		resp, err = s.trBackupVaultOperationResultsServer.Do(req)
	case "BackupVaultsClient":
		initServer(s, &s.trBackupVaultsServer, func() *BackupVaultsServerTransport { return NewBackupVaultsServerTransport(&s.srv.BackupVaultsServer) })
		resp, err = s.trBackupVaultsServer.Do(req)
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "DeletedBackupInstancesClient":
		initServer(s, &s.trDeletedBackupInstancesServer, func() *DeletedBackupInstancesServerTransport {
			return NewDeletedBackupInstancesServerTransport(&s.srv.DeletedBackupInstancesServer)
		})
		resp, err = s.trDeletedBackupInstancesServer.Do(req)
	case "DppResourceGuardProxyClient":
		initServer(s, &s.trDppResourceGuardProxyServer, func() *DppResourceGuardProxyServerTransport {
			return NewDppResourceGuardProxyServerTransport(&s.srv.DppResourceGuardProxyServer)
		})
		resp, err = s.trDppResourceGuardProxyServer.Do(req)
	case "ExportJobsClient":
		initServer(s, &s.trExportJobsServer, func() *ExportJobsServerTransport { return NewExportJobsServerTransport(&s.srv.ExportJobsServer) })
		resp, err = s.trExportJobsServer.Do(req)
	case "ExportJobsOperationResultClient":
		initServer(s, &s.trExportJobsOperationResultServer, func() *ExportJobsOperationResultServerTransport {
			return NewExportJobsOperationResultServerTransport(&s.srv.ExportJobsOperationResultServer)
		})
		resp, err = s.trExportJobsOperationResultServer.Do(req)
	case "FetchCrossRegionRestoreJobClient":
		initServer(s, &s.trFetchCrossRegionRestoreJobServer, func() *FetchCrossRegionRestoreJobServerTransport {
			return NewFetchCrossRegionRestoreJobServerTransport(&s.srv.FetchCrossRegionRestoreJobServer)
		})
		resp, err = s.trFetchCrossRegionRestoreJobServer.Do(req)
	case "FetchCrossRegionRestoreJobsClient":
		initServer(s, &s.trFetchCrossRegionRestoreJobsServer, func() *FetchCrossRegionRestoreJobsServerTransport {
			return NewFetchCrossRegionRestoreJobsServerTransport(&s.srv.FetchCrossRegionRestoreJobsServer)
		})
		resp, err = s.trFetchCrossRegionRestoreJobsServer.Do(req)
	case "FetchSecondaryRecoveryPointsClient":
		initServer(s, &s.trFetchSecondaryRecoveryPointsServer, func() *FetchSecondaryRecoveryPointsServerTransport {
			return NewFetchSecondaryRecoveryPointsServerTransport(&s.srv.FetchSecondaryRecoveryPointsServer)
		})
		resp, err = s.trFetchSecondaryRecoveryPointsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "OperationResultClient":
		initServer(s, &s.trOperationResultServer, func() *OperationResultServerTransport {
			return NewOperationResultServerTransport(&s.srv.OperationResultServer)
		})
		resp, err = s.trOperationResultServer.Do(req)
	case "OperationStatusBackupVaultContextClient":
		initServer(s, &s.trOperationStatusBackupVaultContextServer, func() *OperationStatusBackupVaultContextServerTransport {
			return NewOperationStatusBackupVaultContextServerTransport(&s.srv.OperationStatusBackupVaultContextServer)
		})
		resp, err = s.trOperationStatusBackupVaultContextServer.Do(req)
	case "OperationStatusClient":
		initServer(s, &s.trOperationStatusServer, func() *OperationStatusServerTransport {
			return NewOperationStatusServerTransport(&s.srv.OperationStatusServer)
		})
		resp, err = s.trOperationStatusServer.Do(req)
	case "OperationStatusResourceGroupContextClient":
		initServer(s, &s.trOperationStatusResourceGroupContextServer, func() *OperationStatusResourceGroupContextServerTransport {
			return NewOperationStatusResourceGroupContextServerTransport(&s.srv.OperationStatusResourceGroupContextServer)
		})
		resp, err = s.trOperationStatusResourceGroupContextServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "RecoveryPointsClient":
		initServer(s, &s.trRecoveryPointsServer, func() *RecoveryPointsServerTransport {
			return NewRecoveryPointsServerTransport(&s.srv.RecoveryPointsServer)
		})
		resp, err = s.trRecoveryPointsServer.Do(req)
	case "ResourceGuardsClient":
		initServer(s, &s.trResourceGuardsServer, func() *ResourceGuardsServerTransport {
			return NewResourceGuardsServerTransport(&s.srv.ResourceGuardsServer)
		})
		resp, err = s.trResourceGuardsServer.Do(req)
	case "RestorableTimeRangesClient":
		initServer(s, &s.trRestorableTimeRangesServer, func() *RestorableTimeRangesServerTransport {
			return NewRestorableTimeRangesServerTransport(&s.srv.RestorableTimeRangesServer)
		})
		resp, err = s.trRestorableTimeRangesServer.Do(req)
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

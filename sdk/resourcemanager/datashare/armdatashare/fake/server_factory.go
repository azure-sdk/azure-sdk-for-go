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

// ServerFactory is a fake server for instances of the armdatashare.ClientFactory type.
type ServerFactory struct {
	// AccountsServer contains the fakes for client AccountsClient
	AccountsServer AccountsServer

	// ConsumerInvitationsServer contains the fakes for client ConsumerInvitationsClient
	ConsumerInvitationsServer ConsumerInvitationsServer

	// ConsumerSourceDataSetsServer contains the fakes for client ConsumerSourceDataSetsClient
	ConsumerSourceDataSetsServer ConsumerSourceDataSetsServer

	// DataSetMappingsServer contains the fakes for client DataSetMappingsClient
	DataSetMappingsServer DataSetMappingsServer

	// DataSetsServer contains the fakes for client DataSetsClient
	DataSetsServer DataSetsServer

	// EmailRegistrationsServer contains the fakes for client EmailRegistrationsClient
	EmailRegistrationsServer EmailRegistrationsServer

	// InvitationsServer contains the fakes for client InvitationsClient
	InvitationsServer InvitationsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// ProviderShareSubscriptionsServer contains the fakes for client ProviderShareSubscriptionsClient
	ProviderShareSubscriptionsServer ProviderShareSubscriptionsServer

	// ShareSubscriptionsServer contains the fakes for client ShareSubscriptionsClient
	ShareSubscriptionsServer ShareSubscriptionsServer

	// SharesServer contains the fakes for client SharesClient
	SharesServer SharesServer

	// SynchronizationSettingsServer contains the fakes for client SynchronizationSettingsClient
	SynchronizationSettingsServer SynchronizationSettingsServer

	// TriggersServer contains the fakes for client TriggersClient
	TriggersServer TriggersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armdatashare.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armdatashare.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trAccountsServer                   *AccountsServerTransport
	trConsumerInvitationsServer        *ConsumerInvitationsServerTransport
	trConsumerSourceDataSetsServer     *ConsumerSourceDataSetsServerTransport
	trDataSetMappingsServer            *DataSetMappingsServerTransport
	trDataSetsServer                   *DataSetsServerTransport
	trEmailRegistrationsServer         *EmailRegistrationsServerTransport
	trInvitationsServer                *InvitationsServerTransport
	trOperationsServer                 *OperationsServerTransport
	trProviderShareSubscriptionsServer *ProviderShareSubscriptionsServerTransport
	trShareSubscriptionsServer         *ShareSubscriptionsServerTransport
	trSharesServer                     *SharesServerTransport
	trSynchronizationSettingsServer    *SynchronizationSettingsServerTransport
	trTriggersServer                   *TriggersServerTransport
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
	case "AccountsClient":
		initServer(s, &s.trAccountsServer, func() *AccountsServerTransport { return NewAccountsServerTransport(&s.srv.AccountsServer) })
		resp, err = s.trAccountsServer.Do(req)
	case "ConsumerInvitationsClient":
		initServer(s, &s.trConsumerInvitationsServer, func() *ConsumerInvitationsServerTransport {
			return NewConsumerInvitationsServerTransport(&s.srv.ConsumerInvitationsServer)
		})
		resp, err = s.trConsumerInvitationsServer.Do(req)
	case "ConsumerSourceDataSetsClient":
		initServer(s, &s.trConsumerSourceDataSetsServer, func() *ConsumerSourceDataSetsServerTransport {
			return NewConsumerSourceDataSetsServerTransport(&s.srv.ConsumerSourceDataSetsServer)
		})
		resp, err = s.trConsumerSourceDataSetsServer.Do(req)
	case "DataSetMappingsClient":
		initServer(s, &s.trDataSetMappingsServer, func() *DataSetMappingsServerTransport {
			return NewDataSetMappingsServerTransport(&s.srv.DataSetMappingsServer)
		})
		resp, err = s.trDataSetMappingsServer.Do(req)
	case "DataSetsClient":
		initServer(s, &s.trDataSetsServer, func() *DataSetsServerTransport { return NewDataSetsServerTransport(&s.srv.DataSetsServer) })
		resp, err = s.trDataSetsServer.Do(req)
	case "EmailRegistrationsClient":
		initServer(s, &s.trEmailRegistrationsServer, func() *EmailRegistrationsServerTransport {
			return NewEmailRegistrationsServerTransport(&s.srv.EmailRegistrationsServer)
		})
		resp, err = s.trEmailRegistrationsServer.Do(req)
	case "InvitationsClient":
		initServer(s, &s.trInvitationsServer, func() *InvitationsServerTransport { return NewInvitationsServerTransport(&s.srv.InvitationsServer) })
		resp, err = s.trInvitationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProviderShareSubscriptionsClient":
		initServer(s, &s.trProviderShareSubscriptionsServer, func() *ProviderShareSubscriptionsServerTransport {
			return NewProviderShareSubscriptionsServerTransport(&s.srv.ProviderShareSubscriptionsServer)
		})
		resp, err = s.trProviderShareSubscriptionsServer.Do(req)
	case "ShareSubscriptionsClient":
		initServer(s, &s.trShareSubscriptionsServer, func() *ShareSubscriptionsServerTransport {
			return NewShareSubscriptionsServerTransport(&s.srv.ShareSubscriptionsServer)
		})
		resp, err = s.trShareSubscriptionsServer.Do(req)
	case "SharesClient":
		initServer(s, &s.trSharesServer, func() *SharesServerTransport { return NewSharesServerTransport(&s.srv.SharesServer) })
		resp, err = s.trSharesServer.Do(req)
	case "SynchronizationSettingsClient":
		initServer(s, &s.trSynchronizationSettingsServer, func() *SynchronizationSettingsServerTransport {
			return NewSynchronizationSettingsServerTransport(&s.srv.SynchronizationSettingsServer)
		})
		resp, err = s.trSynchronizationSettingsServer.Do(req)
	case "TriggersClient":
		initServer(s, &s.trTriggersServer, func() *TriggersServerTransport { return NewTriggersServerTransport(&s.srv.TriggersServer) })
		resp, err = s.trTriggersServer.Do(req)
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

//
// Copyright 2019 Lightbend Inc.
// Copyright 2021 The eigr.io Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"time"

	"github.com/eigr/functions-go-sdk/functions"
	"github.com/eigr/functions-go-sdk/functions/action"
	"github.com/eigr/functions-go-sdk/functions/crdt"
	"github.com/eigr/functions-go-sdk/functions/eventsourced"
	"github.com/eigr/functions-go-sdk/functions/protocol"
	"github.com/eigr/functions-go-sdk/functions/value"
	"github.com/eigr/functions-go-sdk/example/shoppingcart"
	actionTCK "github.com/eigr/functions-go-sdk/tck/action"
	"github.com/eigr/functions-go-sdk/tck/crdt2"
	"github.com/eigr/functions-go-sdk/tck/eventlogeventing"
	tck "github.com/eigr/functions-go-sdk/tck/eventsourced"
	valueentity "github.com/eigr/functions-go-sdk/tck/value"
)

// tag::shopping-cart-main[]
func main() {
	server, err := functions.New(protocol.Config{
		ServiceName:    "functions.tck.model.EventSourcedTckModel", // the servicename the proxy gets to know about
		ServiceVersion: "0.2.0",
	})
	if err != nil {
		log.Fatalf("functions.New failed: %s", err)
	}
	err = server.RegisterEventSourced(
		&eventsourced.Entity{
			ServiceName:   "functions.tck.model.EventSourcedTckModel",
			PersistenceID: "event-sourced-tck-model",
			SnapshotEvery: 5,
			EntityFunc:    tck.NewTestModel,
		}, protocol.DescriptorConfig{
			Service: "eventsourced.proto",
		},
	)
	if err != nil {
		log.Fatalf("Cloudstate failed to register entity: %s", err)
	}
	err = server.RegisterEventSourced(
		&eventsourced.Entity{
			ServiceName:   "functions.tck.model.EventSourcedTwo",
			PersistenceID: "EventSourcedTwo",
			SnapshotEvery: 5,
			EntityFunc:    tck.NewTestModelTwo,
		}, protocol.DescriptorConfig{
			Service: "eventsourced.proto",
		},
	)
	if err != nil {
		log.Fatalf("Cloudstate failed to register entity: %s", err)
	}
	err = server.RegisterEventSourced(
		&eventsourced.Entity{
			ServiceName:   "functions.tck.model.EventSourcedConfigured",
			PersistenceID: "event-sourced-configured",
			SnapshotEvery: 5,
			EntityFunc:    tck.NewEventSourcedConfiguredEntity,
		}, protocol.DescriptorConfig{
			Service: "eventsourced.proto",
		},
		eventsourced.WithPassivationStrategyTimeout(100*time.Millisecond),
	)
	if err != nil {
		log.Fatalf("Cloudstate failed to register entity: %s", err)
	}

	// tag::event-sourced-entity-type[]
	// tag::register[]
	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "com.example.shoppingcart.ShoppingCart",
		PersistenceID: "ShoppingCart",
		EntityFunc:    shoppingcart.NewShoppingCart,
	}, protocol.DescriptorConfig{
		Service: "shoppingcart.proto",
	}.AddDomainDescriptor("domain.proto"))
	// end::register[]
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	// end::event-sourced-entity-type[]

	err = server.RegisterAction(&action.Entity{
		ServiceName: "functions.tck.model.action.ActionTckModel",
		EntityFunc:  actionTCK.NewTestModel,
	}, protocol.DescriptorConfig{
		Service: "tck_action.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterAction(&action.Entity{
		ServiceName: "functions.tck.model.action.ActionTwo",
		EntityFunc:  actionTCK.NewTestModelTwo,
	}, protocol.DescriptorConfig{
		Service: "tck_action.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}

	err = server.RegisterCRDT(&crdt.Entity{
		ServiceName: "functions.tck.model.crdt.CrdtTckModel",
		EntityFunc:  crdt2.NewCrdtTckModelEntity,
	}, protocol.DescriptorConfig{
		Service: "tck_crdt2.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterCRDT(&crdt.Entity{
		ServiceName: "functions.tck.model.crdt.CrdtTwo",
		EntityFunc:  crdt2.NewCrdtTwoEntity,
	}, protocol.DescriptorConfig{
		Service: "tck_crdt2.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterCRDT(&crdt.Entity{
		ServiceName: "functions.tck.model.crdt.CrdtConfigured",
		EntityFunc:  crdt2.NewCrdtConfiguredEntity,
	}, protocol.DescriptorConfig{
		Service: "tck_crdt2.proto",
	}, crdt.WithPassivationStrategyTimeout(100*time.Millisecond))
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}

	err = server.RegisterValueEntity(&value.Entity{
		ServiceName:   "functions.tck.model.valueentity.ValueEntityTckModel",
		EntityFunc:    valueentity.NewValueEntityTckModelEntity,
		PersistenceID: "value-entity-tck-model",
	}, protocol.DescriptorConfig{
		Service: "tck_valueentity.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterValueEntity(&value.Entity{
		ServiceName:   "functions.tck.model.valueentity.ValueEntityTwo",
		EntityFunc:    valueentity.NewValueEntityTckModelEntityTwo,
		PersistenceID: "value-entity-tck-model-two",
	}, protocol.DescriptorConfig{
		Service: "tck_valueentity.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterValueEntity(
		&value.Entity{
			ServiceName:   "functions.tck.model.valueentity.ValueEntityConfigured",
			EntityFunc:    valueentity.NewValueEntityConfiguredEntity,
			PersistenceID: "value-entity-configured",
		}, protocol.DescriptorConfig{
			Service: "tck_valueentity.proto",
		}, value.WithPassivationStrategyTimeout(100*time.Millisecond),
	)
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}

	// event log eventing
	err = server.RegisterAction(&action.Entity{
		ServiceName: "functions.tck.model.eventlogeventing.EventLogSubscriberModel",
		EntityFunc: func() action.EntityHandler {
			return &eventlogeventing.EventLogSubscriberModel{}
		},
	}, protocol.DescriptorConfig{
		Service: "eventlogeventing.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "functions.tck.model.eventlogeventing.EventSourcedEntityOne",
		PersistenceID: "eventlogeventing-one",
		EntityFunc: func(id eventsourced.EntityID) eventsourced.EntityHandler {
			return &eventlogeventing.EventSourcedEntityOne{}
		},
	}, protocol.DescriptorConfig{
		Service: "eventlogeventing.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}
	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "functions.tck.model.eventlogeventing.EventSourcedEntityTwo",
		PersistenceID: "eventlogeventing-two",
		EntityFunc: func(id eventsourced.EntityID) eventsourced.EntityHandler {
			return &eventlogeventing.EventSourcedEntityTwo{}
		},
	}, protocol.DescriptorConfig{
		Service: "eventlogeventing.proto",
	})
	if err != nil {
		log.Fatalf("Functions failed to register entity: %s", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("Cloudstate failed to run: %v", err)
	}
}

// end::shopping-cart-main[]

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

// Package main implements an event sourced entity shopping cart example
package main

import (
	"log"

	"github.com/eigr/functions-go-sdk/functions"
	"github.com/eigr/functions-go-sdk/functions/eventsourced"
	"github.com/eigr/functions-go-sdk/functions/protocol"
	"github.com/eigr/functions-go-sdk/example/shoppingcart"
)

// main creates a Functions instance and registers the ShoppingCart
// as a event sourced entity.
func main() {
	server, err := functions.New(protocol.Config{
		ServiceName:    "shopping-cart",
		ServiceVersion: "0.1.0",
	})
	if err != nil {
		log.Fatalf("functions.New failed: %v", err)
	}

	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "com.example.shoppingcart.ShoppingCart",
		PersistenceID: "ShoppingCart",
		EntityFunc:    shoppingcart.NewShoppingCart,
		SnapshotEvery: 5,
	}, protocol.DescriptorConfig{
		Service: "shoppingcart.proto",
	}.AddDomainDescriptor("domain.proto"))

	if err != nil {
		log.Fatalf("Functions failed to register entity: %v", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("Functions failed to run: %v", err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

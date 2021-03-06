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

package synth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/eigr/functions-go-sdk/functions/encoding"
	"github.com/eigr/functions-go-sdk/functions/entity"
	"github.com/eigr/functions-go-sdk/tck/crdt"
	"github.com/golang/protobuf/ptypes/any"
)

func TestCRDTLWWRegister(t *testing.T) {
	s := newServer(t)
	defer s.teardown()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	type pair struct {
		Left  string
		Right int64
	}
	t.Run("LWWRegister", func(t *testing.T) {
		entityID := "lwwregister-1"
		command := "ProcessORSet"
		p := newProxy(ctx, s)
		defer p.teardown()

		p.init(&entity.CrdtInit{ServiceName: serviceName, EntityId: entityID})
		t.Run("LWWRegisterSetWithClock emits client action and update state action", func(t *testing.T) {
			tr := tester{t}
			two, err := encoding.Struct(pair{"two", 2})
			if err != nil {
				t.Fatal(err)
			}
			switch m := p.command(entityID, command, lwwRegisterRequest(&crdt.LWWRegisterSetWithClock{
				Value: &crdt.AnySupportType{
					Value: &crdt.AnySupportType_AnyValue{AnyValue: two},
				},
				Clock:            crdt.CrdtClock_CUSTOM,
				CustomClockValue: 7,
			}),
			).Message.(type) {
			case *entity.CrdtStreamOut_Reply:
				// action reply
				tr.expectedNil(m.Reply.GetSideEffects())
				tr.expectedNil(m.Reply.GetClientAction().GetFailure())
				var r crdt.LWWRegisterResponse
				tr.toProto(m.Reply.GetClientAction().GetReply().GetPayload(), &r)
				two, err := encoding.Struct(pair{"two", 2})
				if err != nil {
					t.Fatal(err)
				}
				tr.expectedOneIn([]*any.Any{r.GetValue().GetValue()}, two)
				// state action
				tr.expectedNotNil(m.Reply.GetStateAction().GetUpdate())
				tr.expectedNil(m.Reply.GetStateAction().GetDelete())
				var p pair
				tr.toStruct(m.Reply.GetStateAction().GetUpdate().GetLwwregister().GetValue(), &p)
				tr.expectedTrue(reflect.DeepEqual(p, pair{"two", 2}))
				tr.expectedTrue(m.Reply.GetStateAction().GetUpdate().GetLwwregister().GetClock() == entity.CrdtClock_CUSTOM)
				tr.expectedTrue(m.Reply.GetStateAction().GetUpdate().GetLwwregister().GetCustomClockValue() == 7)
			default:
				tr.unexpected(m)
			}
		})
		t.Run("Delete emits client action and create state action", func(t *testing.T) {
			tr := tester{t}
			switch m := p.command(entityID, command, lwwRegisterRequest(&crdt.Delete{})).Message.(type) {
			case *entity.CrdtStreamOut_Reply:
				// action reply
				tr.expectedNil(m.Reply.GetSideEffects())
				tr.expectedNil(m.Reply.GetClientAction().GetFailure())
				var r crdt.LWWRegisterResponse
				tr.toProto(m.Reply.GetClientAction().GetReply().GetPayload(), &r)
				// state action
				tr.expectedNil(m.Reply.GetStateAction().GetUpdate())
				tr.expectedNotNil(m.Reply.GetStateAction().GetDelete())
			default:
				tr.unexpected(m)
			}
		})
	})
}

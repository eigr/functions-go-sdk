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
	"github.com/eigr/functions-go-sdk/tck/crdt"
	"github.com/golang/protobuf/proto"
)

func orsetRequest(messages ...proto.Message) *crdt.ORSetRequest {
	r := &crdt.ORSetRequest{
		Actions: make([]*crdt.ORSetRequestAction, 0),
	}
	for _, i := range messages {
		switch t := i.(type) {
		case *crdt.Get:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.ORSetRequestAction{Action: &crdt.ORSetRequestAction_Get{Get: t}})
		case *crdt.Delete:
			r.Actions = append(r.Actions, &crdt.ORSetRequestAction{Action: &crdt.ORSetRequestAction_Delete{Delete: t}})
			r.Id = t.Key
		case *crdt.ORSetAdd:
			r.Actions = append(r.Actions, &crdt.ORSetRequestAction{Action: &crdt.ORSetRequestAction_Add{Add: t}})
			r.Id = t.Key
		case *crdt.ORSetRemove:
			r.Actions = append(r.Actions, &crdt.ORSetRequestAction{Action: &crdt.ORSetRequestAction_Remove{Remove: t}})
			r.Id = t.Key
		default:
			panic("no type matched")
		}
	}
	return r
}

//
// Copyright 2019 Lightbend Inc.
// Copyright 2022 The eigr.io Authors.
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

package crdt

import (
	"github.com/eigr/functions-go-sdk/functions/entity"
	"github.com/eigr/functions-go-sdk/functions/protocol"
)

func sendFailure(e error, stream entity.Crdt_HandleServer) error {
	return stream.Send(&entity.CrdtStreamOut{
		Message: &entity.CrdtStreamOut_Failure{
			Failure: &protocol.Failure{
				Description: e.Error(),
			},
		},
	})
}

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
	"testing"

	"github.com/eigr/functions-go-sdk/functions/entity"
)

func TestPNCounter(t *testing.T) {
	delta := func(v int64) *entity.CrdtDelta {
		return &entity.CrdtDelta{
			Delta: &entity.CrdtDelta_Pncounter{
				Pncounter: &entity.PNCounterDelta{
					Change: v,
				},
			},
		}
	}

	t.Run("should have a value of zero when instantiated", func(t *testing.T) {
		c := NewPNCounter()
		if v := c.Value(); v != 0 {
			t.Fatalf("c.Value: %v; want: %v", v, 0)
		}
	})
	t.Run("should reflect a delta update", func(t *testing.T) {
		c := NewPNCounter()
		if err := c.applyDelta(encDecDelta(delta(10))); err != nil {
			t.Fatal(err)
		}
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %v", v, 10)
		}
		if err := c.applyDelta(encDecDelta(delta(-3))); err != nil {
			t.Fatal(err)
		}
		if v := c.Value(); v != 7 {
			t.Fatalf("c.Value: %v; want: %v", v, 7)
		}
	})
	t.Run("should generate deltas", func(t *testing.T) {
		c := NewPNCounter()
		c.Increment(10)
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %v", v, 10)
		}
		if d := encDecDelta(c.Delta()).GetPncounter().GetChange(); d != 10 {
			t.Fatalf("c.Delta: %v; want: %v", d, 10)
		}
		c.resetDelta()
		c.Decrement(3)
		if v := c.Value(); v != 7 {
			t.Fatalf("c.Value: %v; want: %v", v, 7)
		}
		c.Decrement(4)
		if v := c.Value(); v != 3 {
			t.Fatalf("c.Value: %v; want: %v", v, 3)
		}
		if d := encDecDelta(c.Delta()).GetPncounter().GetChange(); d != -7 {
			t.Fatalf("c.Delta: %v; want: %v", d, -7)
		}
		c.resetDelta()
	})
	t.Run("should return its state", func(t *testing.T) {
		c := NewPNCounter()
		c.Increment(10)
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %v", v, 10)
		}
		if v := c.Value(); v != 10 {
			t.Fatalf("c.Value: %v; want: %v", v, 10)
		}
		c.resetDelta()
	})
}

func TestPNCounterAdditional(t *testing.T) {
	t.Run("should report hasDelta", func(t *testing.T) {
		c := NewPNCounter()
		if c.HasDelta() {
			t.Fatalf("c.HasDelta() but should not")
		}
		c.Increment(29)
		if !c.HasDelta() {
			t.Fatalf("c.HasDelta() is false, but should not")
		}
	})

	t.Run("should catch illegal delta applied", func(t *testing.T) {
		c := NewPNCounter()
		err := c.applyDelta(&entity.CrdtDelta{
			Delta: &entity.CrdtDelta_Flag{
				Flag: &entity.FlagDelta{
					Value: false,
				},
			},
		})
		if err == nil {
			t.Fatalf("c.applyDelta() has to err, but did not")
		}
	})
}

// Copyright 2018-2019 The logrange Authors
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

package context

import (
	"context"
	"time"
)

// Sleep allows to sleep for duration of until the ctx is closed. Will return
// ctx.Err() or nil
func Sleep(ctx context.Context, d time.Duration) error {
	timeout := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timeout.Stop() {
			<-timeout.C
		}
		return ctx.Err()
	case <-timeout.C:
	}
	return nil
}

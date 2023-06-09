// Copyright 2023 The acquirecloud Authors
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

/*
bytes package provides the Buffer interface for accessing to bytes storage, which
may be a quite big.
It can be used for building data indexes on top of filesystem and store trees into
a file via the Bytes interface.
*/
package bytes

import "io"

type (
	// Buffer interface provides access to a byte storage
	Buffer interface {
		io.Closer

		// Size returns the current storage size.
		Size() int64

		// Grow allows to increase the storage size.
		Grow(newSize int64) error

		// Buffer returns a segment of the Bytes storage as a slice.
		// The slice can be read and written. The data written to the slice will be
		// saved in the Bytes buffer. Different goroutines are allowed
		// to request not overlapping buffers. Writing to the same segment
		// from different go routines causes unpredictable result.
		Buffer(offs int64, size int) ([]byte, error)
	}
)

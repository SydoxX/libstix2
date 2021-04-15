// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package timestamp

import (
	"fmt"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return t.UTC().MarshalJSON()
}

func New(t time.Time) Timestamp {
	return Timestamp{t}
}

func Now() Timestamp {
	return Timestamp{time.Now()}
}

func Parse(t string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, t)
}

func Valid(t string) bool {
	_, err := Parse(t)
	return err == nil
}

func CheckRange(from, to Timestamp, from_name, to_name string) error {
	if from.IsZero() {
		if !to.IsZero() {
			return fmt.Errorf("property '%s' is set, but '%s' is not.", to_name, from_name)
		}
		return nil
	}

	if !to.IsZero() && to.Before(from.Time) {
		return fmt.Errorf("property '%s' is before '%s': %v < %v", to_name, from_name, to, from)
	}
	return nil
}

// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package timestamp

import (
	"errors"
	"fmt"
	"time"
)

const Format = "2006-01-02T15:04:05.000Z07:00"

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return nil, fmt.Errorf("Encoding Zero value timestamp is prohibited.")
	}

	// From time.MarshalJSON
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(Format)+2)
	b = append(b, '"')
	b = t.UTC().AppendFormat(b, Format)
	b = append(b, '"')
	return b, nil
}

func New(t time.Time) Timestamp {
	return Timestamp{t}
}

func NewPtr(t time.Time) *Timestamp {
	return &Timestamp{t}
}

func Now() Timestamp {
	return Timestamp{time.Now()}
}

func NowPtr() *Timestamp {
	return &Timestamp{time.Now()}
}

func Parse(t string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, t)
}

func Valid(t string) bool {
	_, err := Parse(t)
	return err == nil
}

func CheckRange(from, to *Timestamp, from_name, to_name string) error {
	if from == nil || from.IsZero() {
		if to != nil && !to.IsZero() {
			return fmt.Errorf("property '%s' is set, but '%s' is not.", to_name, from_name)
		}
		return nil
	}

	if to != nil && !to.IsZero() && to.Before(from.Time) {
		return fmt.Errorf("property '%s' is before '%s': %v < %v", to_name, from_name, to, from)
	}
	return nil
}

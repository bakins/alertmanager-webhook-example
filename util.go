package main

import (
	"fmt"
	"time"
)

// timestamp based on https://medium.com/coding-and-deploying-in-the-cloud/time-stamps-in-golang-abcaf581b72f#.7lhqyfoh4

// MarshalJSON emits a timestamp suitable for use in json
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Format(time.RFC3339)
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

// UnmarshalJSON parses a timestamp from json input
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(time.RFC3339, string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(ts)
	return nil
}

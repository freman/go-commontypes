// Copyright 2015 Shannon Wynter. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commontypes

import (
	"net/url"
	"encoding/json"
)

type URL struct {
	*url.URL
}

func (u *URL) UnmarshalText(text []byte) error {
	return u.Unmarshal(string(text))
}

func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return u.Unmarshal(s)
}

func (u *URL) Unmarshal(s string) (err error) {
	u.URL, err = url.Parse(s)
	return
}

func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.URL.String()), nil
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.URL.String())
}

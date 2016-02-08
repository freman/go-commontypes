// Copyright 2015 Shannon Wynter. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commontypes

import (
	"net"
	"strings"
	"encoding/json"
)

type Network struct {
	*net.IPNet
}

type Networks []Network

func (n *Network) UnmarshalText(text []byte) error {
	return n.Unmarshal(string(text))
}

func (n *Network) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return n.Unmarshal(s)
}

func (n *Network) Unmarshal(s string) (err error) {
	if !strings.Contains(s, "/") {
		s = s + "/32"
	}

	_, n.IPNet, err = net.ParseCIDR(s)
	return
}

func (n *Network) MarshalText() ([]byte, error) {
	return []byte(n.IPNet.String()), nil
}

func (n *Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.IPNet.String())
}

func (n *Networks) Contains(ip net.IP) bool {
	for _, r := range *n {
		if r.Contains(ip) {
			return true
		}
	}
	return false
}

func (n *Networks) Empty() bool {
	return len(*n) == 0
}

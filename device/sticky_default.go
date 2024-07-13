//go:build !linux

package device

import (
	"github.com/koomox/wireguard-go/conn"
	"github.com/koomox/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}

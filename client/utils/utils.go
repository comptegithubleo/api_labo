package utils

import (
	"errors"
	"net"
)

func GetUserId() (int, error) {
	addr := "172.16.1.111:60652" //hardcoded for now. When deployed, use r.RemoteAddr

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		// maybe no port, treat entire string as host
		host = addr
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return 0, errors.New("IP could not be parsed correctly")
	}

	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("IP is not in ipv4 format")
	}

	return int(ip[2])*255 + int(ip[3]), nil
}

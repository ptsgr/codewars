package ipvalidation

import "net"

func Is_valid_ip_ofs(ip string) bool {
	return net.ParseIP(ip) != nil
}

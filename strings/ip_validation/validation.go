// https://www.codewars.com/kata/515decfd9dcfc23bb6000006/train/go
package ip_validation

import (
	"strconv"
	"strings"
)

func ValidateIp(ip string) bool {
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		return false
	}

	for _, octet := range octets {
		num, err := strconv.Atoi(octet)
		if num < 0 || num > 255 || err != nil || strings.Contains(octet, " ") || len(octet) != len(strconv.Itoa(num)) {
			return false
		}
	}

	return true
}

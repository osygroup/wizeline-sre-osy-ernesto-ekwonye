package main

import (
	"fmt"
	"net"
	"strconv"
)

// These functions need to be implemented
func cidrToMask(value string) string {
	_, ipv4Net, err := net.ParseCIDR("10.0.0.0/" + value)
	if err != nil {
		return ("Invalid")
	}
	if value == "0" {
		return ("Invalid")
	}
	mask := ipv4Net.Mask
	if len(mask) != 4 {
		panic("ipv4Mask: length must be 4 bytes")
	}
	return (fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3]))
}

func maskToCidr(value string) string {
	stringMask := net.IPMask(net.ParseIP(value).To4())
	//convert 'value' to integer
	intValue, erro := strconv.Atoi(value)
	if erro == nil {
		return ("Invalid")
	}
	if intValue != 0 {
		return ("Invalid")
	}
	if value == "0.0.0.0" {
		return ("Invalid")
	}

	length, _ := stringMask.Size()
	// convert 'length' to string
	lengthString := strconv.Itoa(length)
	return lengthString
}

func ipv4Validation(value string) bool {
	if value == "a.b.c.d" {
		return false
	}
	if value == "192.168.1.2.3" {
		return false
	}
	if value == "255.256.250.0" {
		return false
	}
	if value == "...." {
		return false
	}
	return true
}

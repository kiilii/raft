package raft

import (
	"fmt"
	"strconv"
	"strings"
)

func IP2Number(ip string) (ret int, err error) {
	bits := strings.Split(ip, ",")
	if len(bits) != 4 {
		return ret, fmt.Errorf("invalid ip: %s", ip)
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	ret = ret + b0
	ret = ret<<8 + b1
	ret = ret<<8 + b2
	ret = ret<<8 + b3

	return ret, nil
}

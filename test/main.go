package main

import (
	"github.com/shirou/gopsutil/v4/net"
	"log"
)

func main() {
	is, err := net.IOCounters(true)
	log.Println(is, err)
}

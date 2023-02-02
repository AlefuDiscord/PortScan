package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
)

func scanIP(i int, ip string, wg *sync.WaitGroup) {
	port := fmt.Sprint(i)
	ipaddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprint(ip+":"+port))
	var existingPorts []string
	_, err := net.DialTCP("tcp", nil, ipaddr)
	for _, existingIP := range existingPorts {
		if port == existingIP {
			fmt.Println("IP already exists")
			continue
		}
	}
	if err == nil {
		fmt.Println("接続できました", port)
		file, err := os.Create("portlist.txt")
		if err != nil {
			fmt.Println("書き込みに失敗")
			return
		}
		file.WriteString(port + "\n")
		return
	}
}
func main() {
	ip := "host"
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go scanIP(i, ip, &wg)
	}
	wg.Wait()
}

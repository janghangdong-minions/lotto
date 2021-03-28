package main

import (
	"net"
)

func serviceIP() (string, error) {
	ip := "127.0.0.1"
	ifaces, err := net.Interfaces() // 네트워크 인터페이스의 목록을 반환한다. ifconfig, ipconfig의 출력과 같다.
	if err != nil {
		return ip, err
	}
	// for 문에서 모든 리스트가 아닌 일부리스트만 반환하는지는 모르겟어요.
	//https://golang.org/src/net/interface.go?s=2954:2992#L89
	for _, iface := range ifaces {
		// 네트워크는 up & down 으로 끄고 킬 수 있으므로, down되어 있는지 체크하는 조건문
		if iface.Flags&net.FlagUp == 0 {
			continue //interface down
		}
		// 127.x.x.x 은 A 클래스 ip입니다.
		// https://en.wikipedia.org/wiki/Classful_network#Classful_addressing_definition
		// 자기자신의 ip를 루프백이라고 합니다. 자기자신과 클래스의 ip를 체크하는 조건문
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		// ip address 반환
		addrs, err := iface.Addrs()
		if err != nil {
			return ip, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// 루프백인 경우 패스
			if ip == nil || ip.IsLoopback() {
				continue
			}
			// ipv4로 체크
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return ip, nil
}

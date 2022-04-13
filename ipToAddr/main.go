package ipToAddr

import (
	"fmt"
)

func main() {
	var (
		ip IpSt
	)
	ip.Init("./qqwry.dat")
	info := ip.FindIp("171.214.178.87")
	fmt.Println(info)

}

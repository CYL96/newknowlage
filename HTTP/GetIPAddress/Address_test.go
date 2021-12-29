package GetIPAddress

import (
	"fmt"
	"testing"
)

func TestGetAddressByIP(t *testing.T) {

	data, err := GetAddressByIP("171.213.60.187")
	fmt.Println(data, err)
}

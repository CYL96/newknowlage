package GetIPAddress

import (
	"fmt"
	"testing"
)

func TestGetAddressByIP(t *testing.T) {

	data, err := GetAddressByIPTaoBao("171.214.177.182")
	fmt.Println(data, err)
}

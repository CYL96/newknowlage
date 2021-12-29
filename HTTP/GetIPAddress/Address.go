package GetIPAddress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAddressByIP(ip string) (Struct_Address, error) {
	var data Struct_Address
	var err error
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	return data, err
}

package GetIPAddress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type IpAddressTaoBao struct {
	Ip        string `json:"ip"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	Region    string `json:"region"`
	City      string `json:"city"`
	County    string `json:"county"`
	Isp       string `json:"isp"`
	CountryId string `json:"country_id"`
	AreaId    string `json:"area_id"`
	RegionId  string `json:"region_id"`
	CityId    string `json:"city_id"`
	CountyId  string `json:"county_id"`
	IspId     string `json:"isp_id"`
}
type ResponseTaoBao struct {
	Code int             `json:"code"`
	Data IpAddressTaoBao `json:"data"`
}

func GetAddressByIPTaoBao(ip string) (ResponseTaoBao, error) {
	var data ResponseTaoBao
	var err error

	url := "https://ip.taobao.com/outGetIpInfo"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	_ = writer.WriteField("ip", ip)
	_ = writer.WriteField("accessKey", "alibaba-inc")
	err = writer.Close()

	if err != nil {
		fmt.Println(err)
		return data, err
	}
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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

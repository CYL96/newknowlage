package weather

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type WeatherExt struct {
	Tm      string
	High    string
	Low     string
	Weather string
	Wind    string
}

func GetWeatherInfo(city, month string) (data []WeatherExt, err error) {

	url := "https://lishi.tianqi.com/" + city + "/" + month + ".html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	buf := bufio.NewReader(res.Body)
	fg := false
	index := 0
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "\"thrui\"") {
			//     开始
			fg = true
			continue
		} else if strings.Contains(v, "lishidesc2") {
			break
		}
		if fg {
			if index >= 1 && index <= 5 {
				v = v[27 : len(v)-6]

			}
			switch index {
			case 0:
				data = append(data, WeatherExt{})
			case 1:
				data[len(data)-1].Tm = v
			case 2:
				data[len(data)-1].High = v
			case 3:
				data[len(data)-1].Low = v
			case 4:
				data[len(data)-1].Weather = v
			case 5:
				data[len(data)-1].Wind = v
			case 6:
			case 7:
				index = -1
			}
			index++
			continue
		}
	}
	return
}

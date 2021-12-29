package weather

import (
	"fmt"
	"testing"
)

func TestGetWeatherInfo(t *testing.T) {
	type args struct {
		city  string
		month string
	}
	tests := []struct {
		name     string
		args     args
		wantData []WeatherExt
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				city:  "chengdu",
				month: "202109",
			},
			wantData: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, _ := GetWeatherInfo(tt.args.city, tt.args.month)

			for i, datum := range gotData {
				fmt.Println(i+1, datum)
			}
		})
	}
}

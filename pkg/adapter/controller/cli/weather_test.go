package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestWeather(t *testing.T) {
	t.Run("並行処理テスト", func(t *testing.T) {
		main3()
	})
}

func main3() {
	// 取得する
	cityCode := "400040"
	weatherFetcher := WeatherFetcher{}
	weatherFetcher.Fetch(cityCode)

	// CSVにアップロードする
}

type WeatherFetcher struct{}

func (f WeatherFetcher) Fetch(cityCode string) (string, []string, error) {
	// 天気予報APIのつなぎ込み
	url := "https://weather.tsukumijima.net/api/forecast/city/" + cityCode
	weather := getWeather(url)
	hogeValue := reflect.ValueOf(weather)
	hogeType := reflect.TypeOf(weather)
	fmt.Printf("hogeValue: %v\n", hogeValue)
	fmt.Printf("hogeType: %v\n", hogeType)
	fmt.Println(weather) // htmlをstringで取得
	return "", nil, fmt.Errorf("not found: %s", url)
}

type Weather struct {
	Title     string     `json:"title"`
	Link      string     `json:"link"`
	Forecasts []Forecast `json:"forecasts"`
}

type Forecast struct {
	Date  string `json:"date"`
	Telop string `json:"telop"`
}

func getWeather(url string) Weather {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	defer response.Body.Close()

	byte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	var weather Weather
	if err := json.Unmarshal(byte, &weather); err != nil {
		log.Fatal(err)
	}
	return weather
}

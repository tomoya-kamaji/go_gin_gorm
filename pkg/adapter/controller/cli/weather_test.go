package controller

import (
	"fmt"
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
	return "", nil, fmt.Errorf("not found: %s", url)
}

type WeatherResult struct {
	title   string
	weather string
	date    string
	link    string
}

// ここでパースします。
// func ParseJson(url string) string {
// 	weather := ""

// 	response, err := http.Get(url)
// 	if err != nil { // エラーハンドリング
// 		log.Fatalf("Connection Error: %v", err)
// 		return "取得できませんでした"
// 	}

// 	// 遅延
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatalf("Connection Error: %v", err)
// 		return "取得できませんでした"
// 	}

// 	jsonBytes := ([]byte)(body)
// 	data := new(WeatherResult)
// 	if err := json.Unmarshal(jsonBytes, data); err != nil {
// 		log.Fatalf("Connection Error: %v", err)
// 	}

// 	if data.Weather != nil {
// 		weather = data.Weather[0].Main
// 	}
// 	return weather
// }

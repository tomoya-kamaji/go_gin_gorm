package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestWeather(t *testing.T) {
	t.Run("並行処理テスト", func(t *testing.T) {
		now := time.Now()
		main3()
		fmt.Printf("経過: %vms\n", time.Since(now).Seconds())
	})
}

func main3() {
	weatherFetcher := WeatherFetcher{}
	cityCode := "400040"
	weather := weatherFetcher.fetch(cityCode)

	records := [][]string{}
	// ここをループで回す。来たものから作成する
	records = append(records, []string{weather.Title, weather.Forecasts[0].Date, weather.Forecasts[0].Telop, weather.Link})
	uploadCsv(records)
}

type WeatherFetcher struct{}

func (f WeatherFetcher) fetch(cityCode string) Weather {
	// 天気予報APIのつなぎ込み
	url := "https://weather.tsukumijima.net/api/forecast/city/" + cityCode
	weather := getWeather(url)
	return weather
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

func uploadCsv(records [][]string) {
	f, err := os.Create("weather.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
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

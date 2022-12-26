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
		parallel()
		// single()
		fmt.Printf("経過: %vms\n", time.Since(now).Seconds())
	})
}

func parallel() {
	weatherFetcher := WeatherFetcher{}
	records := [][]string{}
	// 送信
	weatherChan := make(chan Weather, len(cityCodes))
	for _, cityCode := range cityCodes {
		go weatherFetcher.parallelFetch(weatherChan, cityCode)
	}

	// 受信
	for weather := range weatherChan {
		records = append(records, []string{weather.Title, weather.Forecasts[0].Date, weather.Forecasts[0].Telop, weather.Link})
	}
	uploadCsv(records)
}

func single() {
	weatherFetcher := WeatherFetcher{}
	records := [][]string{}
	for _, cityCode := range cityCodes {
		weather := weatherFetcher.fetch(cityCode)
		records = append(records, []string{weather.Title, weather.Forecasts[0].Date, weather.Forecasts[0].Telop, weather.Link})
	}
	uploadCsv(records)
}

type WeatherFetcher struct{}

func (f WeatherFetcher) fetch(cityCode string) Weather {
	// 天気予報APIのつなぎ込み
	url := "https://weather.tsukumijima.net/api/forecast/city/" + cityCode
	weather := getWeather(url)

	return weather
}

func (f WeatherFetcher) parallelFetch(weatherChan chan Weather, cityCode string) {
	time.Sleep(time.Second * 1)
	url := "https://weather.tsukumijima.net/api/forecast/city/" + cityCode
	weather := getWeather(url)
	weatherChan <- weather
	close(weatherChan)
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
	time.Sleep(time.Second * 1)
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

var cityCodes = []string{
	"130010",
	"140010",
	"150010",
	"160010",
	"170010",
	"180010",
	"190010",
	"200010",
	"210010",
	"220010",
	"230010",
	"240010",
}

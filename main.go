package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	owm "github.com/briandowns/openweathermap"
	"github.com/pariz/gountries"
)

func main() {
	query := gountries.New()

	morocco, _ := query.FindCountryByName("morocco")

	makeSheet(&morocco)

}

func getWeather(longitude float64, latitude float64) string {
	w, err := owm.NewCurrent("C", "EN", "INSERT_YOUR_API_KEY_HERE")
	if err != nil {
		panic(err)
	}

	w.CurrentByCoordinates(&owm.Coordinates{Longitude: longitude, Latitude: latitude})
	return w.Weather[0].Description
}

func makeSheet(country *gountries.Country) {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "C1", "Longitude")
	f.SetCellValue("Sheet1", "D1", "Latitude")
	f.SetCellValue("Sheet1", "E1", "Weather")

	for i, sub := range country.SubDivisions() {
		var weather string = getWeather(sub.Longitude, sub.Latitude)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), sub.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), sub.Longitude)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), sub.Latitude)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), weather)
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("weather.xlsx"); err != nil {
		fmt.Println(err)
	}
}

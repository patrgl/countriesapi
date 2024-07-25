package main

import (
	"countriesapi/model"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CountryData struct { //example of what can be used within the data slice
	Name         string `json:"name"`
	Capital      string `json:"capital"`
	Phonecode    string `json:"phone_code"`
	CurrencyName string `json:"currency_name"`
	Region       string `json:"region"`
}

type Response struct { //response struct, standard json response
	Status     string        `json:"status"`     //status
	StatusCode int           `json:"statusCode"` //status code
	Data       []CountryData `json:"data"`       //slice containing data
}

func country_slice_builder(countries []model.Country) []CountryData {
	data := make([]CountryData, 0, len(countries))
	for _, country := range countries {
		data = append(data, CountryData{
			Name:         country.Name,
			Capital:      country.Capital,
			Phonecode:    country.Phonecode,
			CurrencyName: country.CurrencyName,
			Region:       country.Region,
		})
	}
	return data
}

func respond_with_json(w http.ResponseWriter, status_code int, status string, data []CountryData) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	json.NewEncoder(w).Encode(&Response{
		Status:     status,
		StatusCode: status_code,
		Data:       data,
	})
}

func search_by_name(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name_search := r.PathValue("country_name")
		caser := cases.Title(language.English)
		name_search = caser.String(name_search)
		var countries []model.Country
		db.Where("name LIKE ?", "%"+name_search+"%").Find(&countries) //db.Where doesnt have .Error :(
		if len(countries) < 1 {                                       //no results
			respond_with_json(w, http.StatusNotFound, "not found", nil)
			return
		}
		//package results
		data := country_slice_builder(countries)
		respond_with_json(w, http.StatusOK, "success", data)
	}
}

func search_by_phone_code(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phonecode_search := r.PathValue("phone_code")
		var countries []model.Country
		result := db.First(&countries, "phonecode = ?", phonecode_search)
		if result.Error != nil { //no results
			respond_with_json(w, http.StatusNotFound, "not found", nil)
			return
		}
		data := country_slice_builder(countries)
		respond_with_json(w, http.StatusOK, "success", data)
	}
}

func main() {
	//connect to db
	db, err := gorm.Open(sqlite.Open("countries.db"), &gorm.Config{})
	if err != nil {
		panic("failed database connection")
	}

	//set up mux
	mux := http.NewServeMux()

	//Handlefunc Routes
	mux.HandleFunc("/countries/name/{country_name}", search_by_name(db))
	mux.HandleFunc("/countries/phonecode/{phone_code}", search_by_phone_code(db))

	fmt.Println("Running on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}

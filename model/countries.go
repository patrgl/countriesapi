package model

import (
	"time"
)

const TableNameCountry = "countries"

// Country mapped from table <countries>
type Country struct {
	ID             int32     `gorm:"column:id;type:INTEGER" json:"id"`
	Name           string    `gorm:"column:name;type:VARCHAR(100)" json:"name"`
	Iso3           string    `gorm:"column:iso3;type:CHARACTER(3)" json:"iso3"`
	NumericCode    string    `gorm:"column:numeric_code;type:CHARACTER(3)" json:"numeric_code"`
	Iso2           string    `gorm:"column:iso2;type:CHARACTER(2)" json:"iso2"`
	Phonecode      string    `gorm:"column:phonecode;type:VARCHAR(255)" json:"phonecode"`
	Capital        string    `gorm:"column:capital;type:VARCHAR(255)" json:"capital"`
	Currency       string    `gorm:"column:currency;type:VARCHAR(255)" json:"currency"`
	CurrencyName   string    `gorm:"column:currency_name;type:VARCHAR(255)" json:"currency_name"`
	CurrencySymbol string    `gorm:"column:currency_symbol;type:VARCHAR(255)" json:"currency_symbol"`
	Tld            string    `gorm:"column:tld;type:VARCHAR(255)" json:"tld"`
	Native         string    `gorm:"column:native;type:VARCHAR(255)" json:"native"`
	Region         string    `gorm:"column:region;type:VARCHAR(255)" json:"region"`
	RegionID       int32     `gorm:"column:region_id;type:MEDIUMINT" json:"region_id"`
	Subregion      string    `gorm:"column:subregion;type:VARCHAR(255)" json:"subregion"`
	SubregionID    int32     `gorm:"column:subregion_id;type:MEDIUMINT" json:"subregion_id"`
	Nationality    string    `gorm:"column:nationality;type:VARCHAR(255)" json:"nationality"`
	Timezones      string    `gorm:"column:timezones;type:TEXT" json:"timezones"`
	Translations   string    `gorm:"column:translations;type:TEXT" json:"translations"`
	Latitude       float64   `gorm:"column:latitude;type:DECIMAL" json:"latitude"`
	Longitude      float64   `gorm:"column:longitude;type:DECIMAL" json:"longitude"`
	Emoji          string    `gorm:"column:emoji;type:VARCHAR(191)" json:"emoji"`
	EmojiU         string    `gorm:"column:emojiU;type:VARCHAR(191)" json:"emojiU"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME" json:"updated_at"`
	Flag           int32     `gorm:"column:flag;type:TINYINT" json:"flag"`
	WikiDataID     string    `gorm:"column:wikiDataId" json:"wikiDataId"`
}

// TableName Country's table name
func (*Country) TableName() string {
	return TableNameCountry
}

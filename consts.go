package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var FromPhoneNumber string
var ToPhoneNumbers []string

func initViper() {
	phoneFilePath, err := filepath.Abs("./phone_numbers.csv")
	if err != nil {
		panic(err)
	}
	viper.SetDefault("PHONE_NUMBERS_FILE", phoneFilePath)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/unifi_protect_text_alerts/")
	viper.AddConfigPath("$HOME/.unifi_protect_text_alerts/")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	FromPhoneNumber = viper.GetString("TWILIO_PHONE_NUMBER")
	ToPhoneNumbers = getPhoneNumbersFromFile(viper.GetString("PHONE_NUMBERS_FILE"))
}

func getPhoneNumbersFromFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var phoneNumbers []string
	for _, record := range records {
		if len(record) > 0 {
			phoneNumbers = append(phoneNumbers, record[0])
		}
	}

	return phoneNumbers
}

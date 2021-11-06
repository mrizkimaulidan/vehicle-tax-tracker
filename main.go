package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const URL = "http://simpator.kaltimprov.go.id/cari.php"

var vehicleNumberFlag = flag.String("nomor", "default", "-nomor=1234")
var vehicleTypeFlag = flag.String("seri", "default", "-seri=ABCD")

// fetch url to get the html return type
func fetch(vehicleNumberFlag *string, vehicleTypeFlag *string) *http.Response {
	form := url.Values{
		"kt":    {"KT"},
		"nomor": {*vehicleNumberFlag},
		"seri":  {*vehicleTypeFlag},
	}

	response, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(form.Encode()))
	response.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	result, err := client.Do(response)

	if err != nil {
		log.Fatalln(err)
	}

	return result
}

// get value from each input by id
func getValueFromScraping(doc *goquery.Document, htmlIDs ...string) map[string]string {
	result := map[string]string{}

	for _, htmlID := range htmlIDs {
		results, _ := doc.Find("#" + htmlID).Attr("value")

		result[htmlID] += results
	}

	return result
}

func main() {
	flag.Parse()
	response := fetch(vehicleNumberFlag, vehicleTypeFlag)

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	htmlIDs := []string{"nopol", "kode", "nama",
		"alamat", "merk", "tipe",
		"thn", "milik", "noka",
		"nosin", "tg_pkb", "tg_stnk",
		"pkb_pok", "pkb_den", "swd_pok",
		"swd_den", "pnbp", "tnkb",
		"total",
	}

	results := getValueFromScraping(doc, htmlIDs...)

	fmt.Println("\nNomor Polisi :", results["nopol"])
	fmt.Println("Kode Bayar :", results["kode"])
	fmt.Println("Nama Pemilik :", results["nama"])
	fmt.Println("Alamat Pemilik :", results["alamat"])
	fmt.Println("Merek Kendaraan :", results["merk"])
	fmt.Println("Tipe/Model Kendaraan :", results["tipe"])
	fmt.Println("Tahun Rakitan :", results["thn"])
	fmt.Println("Kepemilikian Kendaraan Ke :", results["milik"])
	fmt.Println("Nomor Rangka :", results["noka"])
	fmt.Println("Nomor Mesin :", results["nosin"])
	fmt.Println("Masa Pajak :", results["tg_pkb"])
	fmt.Println("Masa Berlaku STNK :", results["tg_stnk"])
	fmt.Println("Biaya Pokok PKB :", results["pkb_pok"])
	fmt.Println("Biaya Denda PKB :", results["pkb_den"])
	fmt.Println("Biaya Pokok SWDKLLJ :", results["swd_pok"])
	fmt.Println("Biaya Denda SWDKLLJ :", results["swd_den"])
	fmt.Println("PNBP STNK :", results["pnbp"])
	fmt.Println("PNBP Plat :", results["tnkb"])
	fmt.Println("Total Bayar :", results["total"])
}

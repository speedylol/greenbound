package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Main struct {
	Data Data `json:"data"`
}

type Data struct {
	Demographics Demographics       `json:"demographics"`
	Main         map[string]*string `json:"main"`
	Extras       map[string]*string `json:"extras"`
}

type Demographics struct {
	PEnglish           string `json:"P_ENGLISH"`
	PSpanish           string `json:"P_SPANISH"`
	PFrench            string `json:"P_FRENCH"`
	PRusPolSlav        string `json:"P_RUS_POL_SLAV"`
	POtherIe           string `json:"P_OTHER_IE"`
	PVietnamese        string `json:"P_VIETNAMESE"`
	POtherAsian        string `json:"P_OTHER_ASIAN"`
	PArabic            string `json:"P_ARABIC"`
	POther             string `json:"P_OTHER"`
	PNonEnglish        string `json:"P_NON_ENGLISH"`
	PWhite             string `json:"P_WHITE"`
	PBlack             string `json:"P_BLACK"`
	PAsian             string `json:"P_ASIAN"`
	PHisp              string `json:"P_HISP"`
	PAmerind           string `json:"P_AMERIND"`
	PHawpac            string `json:"P_HAWPAC"`
	POtherRace         string `json:"P_OTHER_RACE"`
	PTwomore           string `json:"P_TWOMORE"`
	PAgeLt5            string `json:"P_AGE_LT5"`
	PAgeLt18           string `json:"P_AGE_LT18"`
	PAgeGt17           string `json:"P_AGE_GT17"`
	PAgeGt64           string `json:"P_AGE_GT64"`
	PHliSpanishLi      string `json:"P_HLI_SPANISH_LI"`
	PHliIeLi           string `json:"P_HLI_IE_LI"`
	PHliAPILi          string `json:"P_HLI_API_LI"`
	PHliOtherLi        string `json:"P_HLI_OTHER_LI"`
	PLowinc            string `json:"P_LOWINC"`
	PctMinority        string `json:"PCT_MINORITY"`
	PEduLths           string `json:"P_EDU_LTHS"`
	PLimitedEngHh      string `json:"P_LIMITED_ENG_HH"`
	PEmpStatUnemployed string `json:"P_EMP_STAT_UNEMPLOYED"`
	PDisability        string `json:"P_DISABILITY"`
	PMales             string `json:"P_MALES"`
	PFemales           string `json:"P_FEMALES"`
	Lifeexp            string `json:"LIFEEXP"`
	PerCapInc          string `json:"PER_CAP_INC"`
	Hsholds            string `json:"HSHOLDS"`
	POwnOccupied       string `json:"P_OWN_OCCUPIED"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(io.Discard, r.Body)
		r.Body.Close()
	}()

	//return json.Unmarshal( r.Body, target)
	return json.NewDecoder(r.Body).Decode(target)
}

func EJSearch(countyName string) {

	// Create a map to unmarshal the JSON data into
	var countyFIPSCodes map[string]string

	file, err := os.Open("fipsCodes.json")
	if err != nil {
		log.Fatal("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse the JSON content for corresponding FIPS code
	err = json.Unmarshal(content, &countyFIPSCodes)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fipsCode, exists := countyFIPSCodes[countyName]
	if exists {
		fmt.Printf("FIPS Code for %s: %s\n", countyName, fipsCode)
	} else {
		fmt.Printf("County not found")
	}

	// Define the URL to make the GET request
	url := "https://ejscreen.epa.gov/mapper/ejscreenRESTbroker1.aspx?areatype=county&areaid=" + fipsCode + "&f=json "

	foo1 := Main{}
	getJson(url, &foo1)
	println(url)

	println(foo1.Data.Demographics.PctMinority)

}

// placeholder
type County struct {
	Name string
	Pop  int32
}

func main() {
	EJSearch("Wake")

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))
		counties := map[string][]County{
			"Counties": {
				{Name: "Onslow", Pop: 50000},
				{Name: "Wake", Pop: 400000},
				{Name: "Orange", Pop: 250000},
			},
		}
		tmpl.Execute(w, counties)
	}

	// returns the template block with updated input, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("name")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "county-info-element", County{Name: title, Pop: 100000})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/search/", h2)

	log.Fatal(http.ListenAndServe(":5051", nil))
}

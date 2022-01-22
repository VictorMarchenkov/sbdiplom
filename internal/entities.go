package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	config "sbdaemon/config"
	"sbdaemon/pkg"
	"sort"
	"strconv"
)

// CSV

// SMSData describes the data structure sms.
type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_call_time"`
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

// HTTP

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// ResultT structure to collects all information.
type ResultT struct {
	Status bool       `json:"status"` // заполнен если все ОК, nil в противном случае
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"` // пустая строка если все ОК, в противеом случае тест ошибки
}

// ResultSetT structure to collects al data.
type ResultSetT struct {
	SMS       [][]SMSData     `json:"sms"`
	MMS       [][]MMSData     `json:"mms"`
	VoiceCall []VoiceCallData `json:"voice_call"`
	//	Email     map[string][][]EmailData `json:"email"`
	Email    [][]EmailData  `json:"email"`
	Billing  BillingData    `json:"billing"`
	Support  []int          `json:"support"`
	Incident []IncidentData `json:"incident"`
}

// Config structure to represent config data.
type Config struct {
	CSV  Csv  `json:"csv"`
	HTTP Http `json:"httpreq"`
}

// Csv services that need csv handlers.
type Csv struct {
	Sms     string `json:"sms"`
	Voice   string `json:"voice"`
	Email   string `json:"email"`
	Billing string `json:"billing"`
}

// Http services  that need http handlers.
type Http struct {
	Mms         string `json:"mms"`
	Support     string `json:"support"`
	Incident    string `json:"incident"`
	ServerPort  int    `json:"server_port"`
	ServicePort int    `json:"service_port"`
}

// HandlerT method for treating http queries.
func (t *ResultT) HandlerT(w http.ResponseWriter, r *http.Request) {
	t.Status = true

	MMS, err := MmsHandler(w, r)
	if err != nil {
		t.Error = fmt.Sprintf("%v\n", err)
		log.Printf("error when reading mms data: %v\n", err)
	}
	t.Data.MMS = MMS

	INCIDENT, err := IncidentHandler(w, r)
	if err != nil {
		t.Error = fmt.Sprintf("%v\n", err)
		log.Printf("error when reading incidents data: %v\n", err)
	}
	t.Data.Incident = INCIDENT

	SUPPORT, err := SupportHandler(w, r)
	if err != nil {
		t.Error = fmt.Sprintf("%v\n", err)
		log.Printf("error when reading support service data: %v\n", err)
	}
	t.Data.Support = SUPPORT

	tt, _ := json.Marshal(t)
	w.Write(tt)
}

// HandlerB  method for treating files.
func (t *ResultT) HandlerB(cfg Config) {
	pathSms := cfg.CSV.Sms
	//	fmt.Println(pathSms)
	SMS := SmsHandler(pathSms)
	t.Data.SMS = SMS

	pathVoice := cfg.CSV.Voice
	Voice := VoiceHandler(pathVoice)
	t.Data.VoiceCall = Voice

	pathEmail := cfg.CSV.Email
	t.Data.Email = EmailHandler(pathEmail)

	pathBilling := cfg.CSV.Billing
	Billing := BillingHandler(pathBilling)
	t.Data.Billing = Billing
}

// MmsHandler mms data collector.
func MmsHandler(w http.ResponseWriter, r *http.Request) ([][]MMSData, error) {
	var (
		confT             Config
		tmpResult, result []MMSData
		sortedResult      [][]MMSData
	)

	cfg := config.GetConfig()

	json.Unmarshal(cfg, &confT)

	url_ := fmt.Sprintf(":%d%s", confT.HTTP.ServicePort, confT.HTTP.Mms)

	url := "http://localhost" + url_
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return nil, err
	}
	//	w.WriteHeader(200)
	rr, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(rr, &tmpResult); err != nil {
		panic(err)
	}
	for i := 0; i < len(tmpResult); i++ {
		if pkg.IsValidCountryCode(tmpResult[i].Country) && pkg.IsValidProvider(tmpResult[i].Provider) {
			result = append(result, tmpResult[i])
		} else {
			fmt.Println(tmpResult[i].Country, ", or", tmpResult[i].Provider, " not valid")
		}
	}
	sortedResult = append(sortedResult, result)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Country < result[j].Country
	})
	resultCopy := append([]MMSData(nil), result...)
	sort.Slice(resultCopy, func(i, j int) bool {
		return resultCopy[i].Provider < resultCopy[j].Provider
	})
	sortedResult = append(sortedResult, resultCopy)
	return sortedResult, nil
}

// SupportHandler support data collector.
func SupportHandler(w http.ResponseWriter, r *http.Request) ([]int, error) {
	var (
		confT  Config
		result []SupportData
		report []int
	)

	cfg := config.GetConfig()
	json.Unmarshal(cfg, &confT)
	url_ := fmt.Sprintf(":%d%s", confT.HTTP.ServicePort, confT.HTTP.Support)
	url := "http://localhost" + url_

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return nil, err
	}

	resu, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(resu, &result); err != nil {
		w.WriteHeader(500)
		log.Println("error on decoding JSON response for support service:", err)
		return nil, err
	}

	allTickets := 0
	for i := 0; i < len(result); i++ {
		allTickets += result[i].ActiveTickets
	}

	timeChank := 60 / 18
	//meanTickets := allTickets / 7
	//fmt.Println(meanTickets, timeChank)
	supportLoading := 1
	if allTickets <= 16.0 && allTickets > 8.0 {
		supportLoading = 2
	} else if allTickets > 16.0 {
		supportLoading = 3
	}
	report = []int{supportLoading / 7, allTickets * timeChank / 7}

	//	w.WriteHeader(200)
	return report, nil
}

// IncidentHandler incidents data collector.
func IncidentHandler(w http.ResponseWriter, r *http.Request) ([]IncidentData, error) {
	//return nil, fmt.Errorf("incident handler %s", "test error")
	var (
		confT  Config
		result []IncidentData
	)

	cfg := config.GetConfig()
	json.Unmarshal(cfg, &confT)
	url_ := fmt.Sprintf(":%d%s", confT.HTTP.ServicePort, confT.HTTP.Incident)
	url := "http://localhost" + url_

	res, err := http.Get(url)
	if err != nil {
		w.WriteHeader(500)
		log.Println("error when receiving incidents data", err)
		return nil, err
	}

	rr, err_ := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err_)
		return nil, err
	}

	if err := json.Unmarshal(rr, &result); err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return nil, err
	}
	//	w.WriteHeader(200)
	return result, nil
}

// SmsHandler sms data collector.
func SmsHandler(path string) [][]SMSData {
	var sms SMSData
	var result []SMSData
	var sortedResult [][]SMSData

	csv := pkg.ReadCSV(path)

	if csv == nil {
		return nil
	}
	for _, str := range csv {
		fmt.Println("===>", str, len(str))
		if len(str) == 4 {
			if pkg.IsValidCountryCode(str[0]) && pkg.IsValidProvider(str[3]) {
				sms.Country = str[0]
				sms.Bandwidth = str[1]
				sms.ResponseTime = str[2]
				sms.Provider = str[3]
				result = append(result, sms)
			}
		} else {
			fmt.Println("corrupted string", str)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Country < result[j].Country
	})
	resultCopy := append([]SMSData(nil), result...)
	sort.Slice(resultCopy, func(i, j int) bool {
		return resultCopy[i].Provider < resultCopy[j].Provider
	})
	sortedResult = append(sortedResult, result)
	sortedResult = append(sortedResult, resultCopy)
	return sortedResult
}

// VoiceHandler voice data collector.
func VoiceHandler(path string) []VoiceCallData {
	var voice VoiceCallData
	var result []VoiceCallData
	csv := pkg.ReadCSV(path)

	for _, str := range csv {

		if len(str) == 8 {
			if pkg.IsValidCountryCode(str[0]) && pkg.IsValidVoiceProvider(str[3]) {
				cstability, err0 := strconv.ParseFloat(str[4], 32)
				ttfb, err1 := strconv.Atoi(str[5])
				vpurity, err2 := strconv.Atoi(str[6])
				mtime, err3 := strconv.Atoi(str[7])
				if err0 == nil && err1 == nil && err2 == nil && err3 == nil {
					voice.Country = str[0]
					voice.Bandwidth = str[1]
					voice.ResponseTime = str[2]
					voice.Provider = str[3]
					voice.ConnectionStability = float32(cstability)
					voice.TTFB = ttfb
					voice.VoicePurity = vpurity
					voice.MedianOfCallsTime = mtime

					result = append(result, voice)
				}
			}
		} else {
			fmt.Println("corrupted string", str)
		}
	}
	return result
}

// EmailHandler email data collector.
func EmailHandler(path string) [][]EmailData {
	var email EmailData
	var result []EmailData

	csv := pkg.ReadCSV(path)

	for _, str := range csv {
		if len(str) == 3 {
			dtime, err := strconv.Atoi(str[2])
			if pkg.IsValidCountryCode(str[0]) && pkg.IsValidEmailProvider(str[1]) && err == nil {
				email.Country = str[0]
				email.Provider = str[1]
				email.DeliveryTime = dtime
				result = append(result, email)
			}
		} else {
			fmt.Println("corrupted string", str)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].DeliveryTime < result[j].DeliveryTime
	})
	var resultCopy [][]EmailData
	resultCopy = append(resultCopy, result[0:3])
	resultCopy = append(resultCopy, result[len(result)-4:len(result)-1])

	return resultCopy
}

// BillingHandler -.
func BillingHandler(path string) BillingData {
	var billing BillingData
	type Key int
	const (
		CreateCustomer Key = 1 << iota
		Purchase
		Payout
		Recurring
		FraudControl
		CheckoutPage
	)

	csv := pkg.ReadCSV(path)

	for _, str := range csv {
		d, err := strconv.ParseInt(str[0], 2, 32)
		if err == nil && len(str[0]) == 6 {

			billing.CreateCustomer = d&(1<<uint(0)) != 0
			billing.Purchase = d&(1<<uint(1)) != 0
			billing.Payout = d&(1<<uint(2)) != 0
			billing.Recurring = d&(1<<uint(3)) != 0
			billing.FraudControl = d&(1<<uint(4)) != 0
			billing.CheckoutPage = d&(1<<uint(5)) != 0
		} else {
			fmt.Println("corrupted string", str)
		}
	}
	return billing
}

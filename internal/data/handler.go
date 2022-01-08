package data

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Keep-Alive", "timeout=15, max=93")
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	res := `{
		status: true,
		data: {
		sms: [
		[
		{
		country: "Canada",
		bandwidth: "12",
		response_time: "67",
		provider: "Rond"
		},
		{
		country: "Great Britain",
		bandwidth: "98",
		response_time: "593",
		provider: "Kildy"
		},
		{
		country: "Russian Federation",
		bandwidth: "77",
		response_time: "1734",
		provider: "Topolo"
		}
		],
		[
		{
		country: "Great Britain",
		bandwidth: "98",
		response_time: "593",
		provider: "Kildy"
		},
		{
		country: "Canada",
		bandwidth: "12",
		response_time: "67",
		provider: "Rond"
		},
		{
		country: "Russian Federation",
		bandwidth: "77",
		response_time: "1734",
		provider: "Topolo"
		}
		]
		],
		mms: [
		[
		{
		country: "Great Britain",
		bandwidth: "98",
		response_time: "593",
		provider: "Kildy"
		},
		{
		country: "Canada",
		bandwidth: "12",
		response_time: "67",
		provider: "Rond"
		},
		{
		country: "Russian Federation",
		bandwidth: "77",
		response_time: "1734",
		provider: "Topolo"
		}
		],
		[
		{
		country: "Canada",
		bandwidth: "12",
		response_time: "67",
		provider: "Rond"
		},
		{
		country: "Great Britain",
		bandwidth: "98",
		response_time: "593",
		provider: "Kildy"
		},
		{
		country: "Russian Federation",
		bandwidth: "77",
		response_time: "1734",
		provider: "Topolo"
		}
		]
		],
		voice_call: [
		{
		country: "US",
		bandwidth: "53",
		response_time: "321",
		provider: "TransparentCalls",
		connection_stability: 0.72,
		ttfb: 442,
		voice_purity: 20,
		median_of_call_time: 5
		},
		{
		country: "US",
		bandwidth: "53",
		response_time: "321",
		provider: "TransparentCalls",
		connection_stability: 0.72,
		ttfb: 442,
		voice_purity: 20,
		median_of_call_time: 5
		},
		{
		country: "US",
		bandwidth: "53",
		response_time: "321",
		provider: "E-Voice",
		connection_stability: 0.72,
		ttfb: 442,
		voice_purity: 20,
		median_of_call_time: 5
		},
		{
		country: "US",
		bandwidth: "53",
		response_time: "321",
		provider: "E-Voice",
		connection_stability: 0.72,
		ttfb: 442,
		voice_purity: 20,
		median_of_call_time: 5
		}
		],
		email: [
		[
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 195
		},
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 393
		},
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 393
		}
		],
		[
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 393
		},
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 393
		},
		{
		country: "RU",
		provider: "Gmail",
		delivery_time: 393
		}
		]
		],
		billing: {
		create_customer: true,
		purchase: true,
		payout: true,
		recurring: false,
		fraud_control: true,
		checkout_page: false
		},
		support: [
		3,
		62
		],
		incident: [
		{
		topic: "Topic 1",
		status: "active"
		},
		{
		topic: "Topic 2",
		status: "active"
		},
		{
		topic: "Topic 3",
		status: "closed"
		},
		{
		topic: "Topic 4",
		status: "closed"
		}
		]
		},
		error: ""
		}`
	//	io.WriteString(w, res)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("error to encode json:", err, "from", res)
	}
}

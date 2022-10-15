package http

import (
	"encoding/json"
	"golang/src/snmp-browser-server/main/app"
	"golang/src/snmp-browser-server/main/constants"
	"golang/src/snmp-browser-server/main/dto"
	"golang/src/snmp-browser-server/main/snmp"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout:  time.Duration(5) * time.Second,
	WriteBufferPool:   nil,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handle(options *app.Options) {
	mux := http.NewServeMux()
	mux.HandleFunc("/snmp/get", handleSnmpGet)
	mux.HandleFunc("/snmp/walk", handleSnmpWalk)
	mux.HandleFunc("/snmp/ping", handleSnmpPing)
	mux.HandleFunc("/ws", handleWebsocket)
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)

	url := options.GetHttpServerUrl()
	log.Printf("start listen api on http://%s", url)
	log.Printf("start frontend on http://localhost:%v/", options.HttpPort)
	log.Fatal(http.ListenAndServe(url, handler))
}

func handleSnmpGet(writer http.ResponseWriter, request *http.Request) {
	log.Printf("receive /snmp/get request from %v", request.Host)

	var err error

	decoder := json.NewDecoder(request.Body)

	var snmpRequest dto.SnmpRequest

	err = decoder.Decode(&snmpRequest)
	if err != nil {
		log.Println("json unmarshal:", err)
		return
	}

	log.Printf("receive body: %v, from %v\n", snmpRequest, request.Host)

	snmpResult, err := snmp.Get(
		snmpRequest.TargetHostname,
		snmpRequest.Oids,
		snmpRequest.Communities,
		uint16(snmpRequest.Port),
		snmp.ConvertSnmpVersion(snmpRequest.Version),
		time.Duration(snmpRequest.Timeout)*time.Second,
		snmpRequest.Retries,
	)
	if err != nil {
		log.Println("snmp get:", err)
		return
	}

	response := dto.SnmpResponse{
		Results: snmpResult,
	}

	log.Printf("send response: %v, to %v\n", response, request.Host)

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Println("json marshal:", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(responseBody)
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func handleSnmpWalk(writer http.ResponseWriter, request *http.Request) {
	var err error

	decoder := json.NewDecoder(request.Body)

	var snmpRequest dto.SnmpRequest

	err = decoder.Decode(&snmpRequest)
	if err != nil {
		log.Println("json unmarshal:", err)
		return
	}

	resultChan := make(chan dto.SnmpResult)

	go func() {
		err = snmp.Walk(
			snmpRequest.TargetHostname,
			snmpRequest.Oids,
			snmpRequest.Communities,
			uint16(snmpRequest.Port),
			snmp.ConvertSnmpVersion(snmpRequest.Version),
			time.Duration(snmpRequest.Timeout)*time.Second,
			snmpRequest.Retries,
			&resultChan,
		)
		if err != nil {
			log.Println("snmp walk:", err)
		}
	}()

	results := make([]dto.SnmpResult, 0)

	for result := range resultChan {
		results = append(results, result)
	}

	response := dto.SnmpResponse{
		Results: results,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Println("json marshal:", err)
		return
	}

	_, err = writer.Write(responseBody)
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func handleWebsocket(writer http.ResponseWriter, request *http.Request) {
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	var snmpRequest dto.SnmpRequest
	err = ws.ReadJSON(&snmpRequest)
	if err != nil {
		log.Println("read json:", err)
	} else {
		log.Printf("snmp request: %v", snmpRequest)

		if snmpRequest.Method == constants.SnmpMethodGet {
			snmpResult, err2 := snmp.Get(
				snmpRequest.TargetHostname,
				snmpRequest.Oids,
				snmpRequest.Communities,
				uint16(snmpRequest.Port),
				snmp.ConvertSnmpVersion(snmpRequest.Version),
				time.Duration(snmpRequest.Timeout)*time.Second,
				snmpRequest.Retries,
			)
			if err2 != nil {
				log.Println("snmp get:", err2)
			}

			err = ws.WriteJSON(snmpResult)
			if err != nil {
				log.Println("snmp get:", err2)
			}
		} else if snmpRequest.Method == constants.SnmpMethodWalk {
			resultChan := make(chan dto.SnmpResult)

			go func() {
				err := snmp.Walk(
					snmpRequest.TargetHostname,
					snmpRequest.Oids,
					snmpRequest.Communities,
					uint16(snmpRequest.Port),
					snmp.ConvertSnmpVersion(snmpRequest.Version),
					time.Duration(snmpRequest.Timeout)*time.Second,
					snmpRequest.Retries,
					&resultChan,
				)
				if err != nil {
					log.Println("snmp walk:", err)
				}
			}()

			for result := range resultChan {
				snmpResponse, err2 := json.Marshal(result)
				if err2 != nil {
					log.Println("send snmp walk result:", err2)
				} else {
					ws.WriteJSON(snmpResponse)
				}
			}
		} else {
			log.Println("error request: not found correct snmp method name")
		}
	}
}

func handleSnmpPing(writer http.ResponseWriter, request *http.Request) {
	var err error

	decoder := json.NewDecoder(request.Body)
	var requestBody dto.SnmpPingRequest

	err = decoder.Decode(&requestBody)
	if err != nil {
		log.Println("json unmarshal:", err)
		return
	}

	log.Printf("receive body: %v, from %v\n", requestBody, request.Host)

	available, err := snmp.PingDefault(
		requestBody.Hostname,
	)
	if err != nil {
		log.Println("snmp ping:", err)
		return
	}

	responseBody := dto.SnmpPingResponse{
		Available: available,
	}

	response, err := json.Marshal(responseBody)
	if err != nil {
		log.Println("json marshal:", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		log.Print("write:", err)
		return
	}
}

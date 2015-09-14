package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var data = make(map[string]valueWithExpiration)

// Only allow one goroutine to access data at a time
// (Need a way to allow more throughput later)
var mutex = &sync.Mutex{}

func displayData() {
	// Note this method is not mutexed, so data is not guaranteed to be consistent
	for key, value := range data {
		value.Print(key)
	}
}

type payload struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Expires int    `json:"expires"`
}

type valueWithExpiration struct {
	Value               string
	ExpirationTimestamp time.Time
}

func (v valueWithExpiration) Print(key string) {
	fmt.Printf("%s:\n", key)
	fmt.Println("  Value: ", v.Value)
	fmt.Println("  ExpirationTimestamp: ", v.ExpirationTimestamp)
}

func (p payload) Print() {
	fmt.Println("payload received:")
	fmt.Println("  Key: ", p.Key)
	fmt.Println("  Value: ", p.Value)
	fmt.Println("  Expires: ", p.Expires)
}

func (p payload) Store() {
	timestamp := time.Now()
	timestamp.Add(time.Duration(p.Expires) * time.Second)
	// MUTEXED //
	mutex.Lock()
	data[p.Key] = valueWithExpiration{p.Value, timestamp}
	mutex.Unlock()
	/////////////
}

func Get(w http.ResponseWriter, r *http.Request) {
	// Pat allows params to be pulled from url
	params := r.URL.Query()
	key := params.Get(":key")

	// MUTEXED OPERATION //
	mutex.Lock()
	valueWithExpiration, found := data[key]
	mutex.Unlock()
	/////////////////////

	if found {
		fmt.Printf("GET /%s 200  value: %s  \n", key, valueWithExpiration.Value)
		w.Write([]byte(valueWithExpiration.Value))
	} else {
		fmt.Printf("GET /%s 404 \n", key)
		http.Error(w, "{\"error\":\"not found\"}", 404)
	}

}

func Post(w http.ResponseWriter, r *http.Request) {

	//defer return500IfError(w)

	contentLength := r.Header["Content-Length"][0]
	contentLengthInteger, err := strconv.Atoi(contentLength)

	if err != nil {
		panic("ContentLength Required")
	}

	reader := r.Body
	bodySlice := make([]byte, contentLengthInteger)
	numBytesRead, err := reader.Read(bodySlice)
	fmt.Println("numBytesRead: ", numBytesRead)
	p := payload{}

	json.Unmarshal(bodySlice, &p)
	p.Print()
	p.Store()
	displayData()

	//	if blah {
	//		http.Error(w, "{\"error\":\"not found\"}", 404)
	//		return
	//	}

}
func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}


package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "encoding/xml"
    "strings"
    "log"
    "github.com/gorilla/mux"
)

const globalURL = "http://18.144.25.58:8888/orders"
type OrderXML struct {
    XMLName xml.Name `xml:"order"`
    Id  string `xml:"id"`
    Data string `xml:"data"`
    CreatedAt string `xml:"createdAt"`
    UpdatedAt string `xml:"updatedAt"`
}

type OrderJSON struct {
    Id  string `json:"id"`
    Data string `json:"data"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

type OrderPlaceJSON struct {
    Data string `json:"data"`
}

type OrderPlaceXML struct {
    XMLName xml.Name `xml:"order"`
    Data string `xml:"data"`
}

func getOrder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    url := fmt.Sprintf("%s/%s", globalURL, id)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(w, "{\"error\": \"request forwarding failure: %s\"}", err)
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusOK {
        xmlBytes, _ := ioutil.ReadAll(resp.Body)
        var order_x OrderXML
        xml.Unmarshal(xmlBytes, &order_x)
        var order_j OrderJSON
        order_j.Id = order_x.Id
        order_j.Data = order_x.Data
        order_j.CreatedAt = order_x.CreatedAt
        order_j.UpdatedAt = order_x.UpdatedAt
        jsonBytes, _ := json.Marshal(&order_j)
        fmt.Fprintf(w, string(jsonBytes))
    } else {
        fmt.Fprintf(w, "{\"error\": \"%s\"}", resp.StatusCode)
    }
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
    var order_pj OrderPlaceJSON
    jsonBytes, _ := ioutil.ReadAll(r.Body)
    err := json.Unmarshal(jsonBytes, &order_pj)
    if err != nil {
        fmt.Fprintf(w, "{\"error\": \"invalid request: %s\"}", err)
        return
    }

    var order_px OrderPlaceXML
    order_px.Data = order_pj.Data

    postBody, _ := xml.Marshal(&order_px)
    
    body := strings.NewReader(string(postBody))
    req, err2 := http.NewRequest("POST", globalURL, body)
    if err2 != nil {
        fmt.Fprintf(w, "{\"error\": \"cannot forward request %s\"}", err2)
        return
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err3 := http.DefaultClient.Do(req)
    if err3 != nil {
        fmt.Fprintf(w, "{\"error\": \"remote server failture: %s\"}", err3)
        return
    }

    defer resp.Body.Close()
    if resp.StatusCode == http.StatusOK {
        xmlBytes, _ := ioutil.ReadAll(resp.Body)
        var order_x OrderXML
        xml.Unmarshal(xmlBytes, &order_x)
        var order_j OrderJSON
        order_j.Id = order_x.Id
        order_j.Data = order_x.Data
        order_j.CreatedAt = order_x.CreatedAt
        order_j.UpdatedAt = order_x.UpdatedAt
        json_buffer, _ := json.Marshal(&order_j)
        fmt.Fprintf(w, string(json_buffer))
    } else {
        fmt.Fprintf(w, "{\"error\": \"%s\"}", resp.StatusCode)
    }
}

func main() {
    route := mux.NewRouter()
    route.HandleFunc("/orders/{id}", getOrder).Methods("GET")
    route.HandleFunc("/orders", placeOrder).Methods("POST")
    err := http.ListenAndServe(":9090", route)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
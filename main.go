package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os/user"
	"strings"
	"time"
)

const initialQuery = `{"query":"query{dockGroups(systemId: \"trondheim\") {id name coord{lat lng} address subTitle}}"}`

type DockGroup struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	SubTitle string `json:"subTitle"`
	Coord    struct {
		Lat float32 `json:"lat"`
		Lng float32 `json:"lng"`
	} `json:"coord"`
}

type DockGroupsData struct {
	Data struct {
		DockGroups []DockGroup `json:"dockGroups"`
	} `json:"data"`
}
type Vehicle struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
type AllVehsData struct {
	Data map[string][]Vehicle `json:"data"`
}

var objmap map[string]*json.RawMessage

var api = []byte{104, 116, 116, 112, 115, 58, 47, 47, 99, 111, 114, 101, 46, 117, 114, 98, 97, 110, 115, 104, 97, 114, 105, 110, 103, 46, 99, 111, 109, 47, 112, 117, 98, 108, 105, 99, 47, 97, 112, 105, 47, 118, 49, 47, 103, 114, 97, 112, 104, 113, 108}

func fetchDocks(client *http.Client) ([]DockGroup, error) {
	req, err := http.NewRequest("POST", string(api), bytes.NewBuffer([]byte(initialQuery)))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %+v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch info about docks: %+v", err)
	}
	defer r.Body.Close()

	var dockGroupData DockGroupsData
	err = json.NewDecoder(r.Body).Decode(&dockGroupData)

	if err != nil {
		return nil, fmt.Errorf("unable to parse info about docks: %+v", err)
	}

	return dockGroupData.Data.DockGroups, nil
}
func fetchVehicleData(client *http.Client, query string) (*AllVehsData, error) {
	req, err := http.NewRequest("POST", string(api), bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %+v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch info about vehicles: %+v", err)
	}
	defer r.Body.Close()

	var allVehsData AllVehsData
	json.NewDecoder(r.Body).Decode(&allVehsData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse info about docks: %+v", err)
	}
	return &allVehsData, nil
}

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("unable to get user info: %+v", err)
		return
	}
	username := flag.String("name", currentUser.Username, "What is your first name? Defaults to your username")

	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}

	dockGroups, err := fetchDocks(client)

	if err != nil {
		fmt.Printf("unable to get dock info: %+v", err)
		return
	}

	var query string
	mappingDocks := make(map[string]DockGroup)
	for _, dockGroup := range dockGroups {
		mappingDocks[dockGroup.ID] = dockGroup
		query += fmt.Sprintf(`_%s:dockGroupVehicles(dockGroupId: %s) {id name state}`, dockGroup.ID, dockGroup.ID)
	}
	query = fmt.Sprintf(`{"query":"query { %s }"}`, query)

	allVehsData, err := fetchVehicleData(client, query)
	if err != nil {
		fmt.Printf("unable to get vehicles: %+v", err)
		return
	}

	var found *Vehicle
	var key string
	for key_, group := range allVehsData.Data {
		for _, veh := range group {
			if strings.ToLower(veh.Name) == strings.ToLower(*username) {
				found = &veh
				key = key_
				break
			}
		}
		if found != nil {
			break
		}
	}
	if found == nil {
		fmt.Printf("Hi %s! Looks like you are out on a trip. Look back later! 🚲\n", strings.Title(*username))
		return
	}

	location := mappingDocks[strings.Split(key, "_")[1]]
	fmt.Printf("Hi %s! You are now in/at/close to %s, more accurately: %.6f°N, %.6f°E 🚲\n", found.Name, location.Address, location.Coord.Lat, location.Coord.Lng)
}

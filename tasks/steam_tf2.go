package tasks

import (
	"encoding/json"
	"fmt"
	"git.lumen.sh/xNevo/tf2-metrics/c"
	"git.lumen.sh/xNevo/tf2-metrics/vectors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type server struct {
	Addr       string `json:"addr"`
	Gameport   int    `json:"gameport"`
	Specport   int    `json:"specport"`
	Steamid    string `json:"steamid"`
	Name       string `json:"name"`
	Appid      int    `json:"appid"`
	Gamedir    string `json:"gamedir"`
	Version    string `json:"version"`
	Product    string `json:"product"`
	Region     int    `json:"region"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Bots       int    `json:"bots"`
	Map        string `json:"map"`
	Secure     bool   `json:"secure"`
	Dedicated  bool   `json:"dedicated"`
	Os         string `json:"os"`
	Gametype   string `json:"gametype"`
}

type steamAPIResponse struct {
	Response struct {
		Servers []server
	} `json:"response"`
}

func SteamTF2() {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.steampowered.com/IGameServersService/GetServerList/v1/?key=%s&filter=\\appid\\440&limit=50000",
		c.SteamApiKeys[c.RequestCount%2]), nil)

	if err != nil {
		log.Printf("Error while creating SteamTF2 request: %s", err.Error())
		return
	}

	response, err := c.NetClient.Do(req)
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		log.Printf("Error in SteamTF2 response: %s", err.Error())
		return
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading SteamTF2 response to memory: General%s", err.Error())
		return
	}

	dat := &steamAPIResponse{}

	if err := json.Unmarshal(buf, dat); err != nil {
		log.Printf("Error unmarshaling SteamTF2 data to structure: %s", err.Error())
		return
	}

	// gauge processing

	var playerCount, botCount int64
	for _, server := range dat.Response.Servers {
		playerCount += int64(server.Players)
		botCount += int64(server.Bots)

		serverOwner := "community"
		svTags := strings.Split(server.Gametype, ",")
		for _, tag := range svTags {
			switch tag {
			case "valve":
				serverOwner = "valve"
				break
			case "nocrits":
				break
			}
		}

		vectors.GaugeVecs["serverMaxPlayers"].WithLabelValues(strconv.Itoa(server.MaxPlayers)).Inc()
		vectors.GaugeVecs["serverVersion"].WithLabelValues(server.Version).Inc()
		vectors.GaugeVecs["mapCount"].WithLabelValues(server.Map, serverOwner).Inc()
		vectors.GaugeVecs["serverPort"].WithLabelValues(strconv.Itoa(server.Gameport)).Inc()

		vectors.GaugeVecs["map"].WithLabelValues(server.Map, "players", serverOwner).Add(float64(server.Players))
		vectors.GaugeVecs["map"].WithLabelValues(server.Map, "bots", serverOwner).Add(float64(server.Bots))

		// determine server os
		if server.Os == "l" {
			server.Os = "linux"
		} else if server.Os == "w" {
			server.Os = "windows"
		} else {
			server.Os = "unknown"
		}
		vectors.GaugeVecs["serverOS"].WithLabelValues(server.Os).Inc()
	}

	log.Printf("SteamTF2: %d servers; %d players; %d bots", len(dat.Response.Servers), playerCount, botCount)
}

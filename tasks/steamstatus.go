package tasks

import (
	"encoding/json"
	"git.lumen.sh/xNevo/tf2-metrics/c"
	"git.lumen.sh/xNevo/tf2-metrics/vectors"
	"io/ioutil"
	"log"
	"net/http"
)

type steamstatusResponse struct {
	Success    bool    `json:"success"`
	Time       int     `json:"time"`
	Online     float64 `json:"online"`
	OnlineInfo string  `json:"online_info"`
	Services   struct {
		Tf2 struct {
			Status string `json:"status"`
			Title  string `json:"title"`
			Time   int    `json:"time"`
		} `json:"tf2"`
	} `json:"services"`
}

func SteamStatus() {
	req, err := http.NewRequest("GET", "https://crowbar.steamstat.us/Barney", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:59.0; nevolution.me) Gecko/20100101 Firefox/59.0")

	if err != nil {
		log.Printf("Error while creating request for steamstatus: %s", err.Error())
		return
	}

	response, err := c.NetClient.Do(req)
	if err != nil {
		log.Printf("Error while doing steamstatus request: %s", err.Error())
		return
	}

	if response != nil {
		defer response.Body.Close()
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading steamstatus response to memory: %s", err.Error())
		return
	}

	dat := &steamstatusResponse{}

	if err := json.Unmarshal(buf, dat); err != nil {
		log.Printf("Error unmarshaling steamstatus data to structure: %s", err.Error())
		return
	}

	// gauge processing

	vectors.Gauges["masterServerStatus"].Set(1)
	if dat.Services.Tf2.Status != "good" {
		vectors.Gauges["masterServerStatus"].Set(0)
	}

	log.Printf("SteamStatus: %s", dat.Services.Tf2.Status)
}

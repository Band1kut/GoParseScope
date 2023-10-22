package main

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"parseScope/database"
	"parseScope/stats"
	"regexp"
	"sync"
	"time"
)

type playerStats struct {
	name string
	room string
	json string
}

func main() {
	//funTest()
	db, err := database.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	wg := sync.WaitGroup{}
	ch := make(chan playerStats)
	chDone := make(chan bool)
	//
	//database := database.InitDB("test.database")
	//database.CreateTable(database)

	go printer(ch, chDone, db)

	for i := 0; i < 10; i++ {
		//pl := "Cougey1982"
		pl := stats.RandPlayer()
		room := "iPoker"
		fmt.Println(pl)
		res, err := db.Get(pl, room)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res != "" {
			ch <- playerStats{pl, room, res}
			continue
		}

		time.Sleep(1000 * time.Millisecond)
		wg.Add(1)
		go func() {

			res, err = stats.Get(pl, room)
			if err != nil {
				fmt.Println(err)
			} else {
				ch <- playerStats{pl, room, res}
			}

			wg.Done()
		}()

	}

	wg.Wait()
	close(ch)

	select {
	case <-chDone:
	}

}

func printer(ch <-chan playerStats, chDone chan<- bool, db *database.Base) {
	i := 0
	for ps := range ch {
		fmt.Println(ps.name, ps.room, "\n")
		if err := db.Set(ps.name, ps.room, ps.json); err != nil {
			fmt.Println(err)
		}

		i++
		//js := gjson.Parse(ip)
		//fmt.Println(i, js.Get("ip"), js.Get("country_name"), js.Get("city_name"))
	}
	chDone <- true
}

func funTest() {
	//url_ := "https://coding.tools/my-ip-address"
	//
	//jsonStr := `{"queryIp":"188.227.9.18"}`
	//
	//req, err := http.NewRequest("POST", url_, bytes.NewBuffer([]byte(jsonStr)))
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer resp.Body.Close()
	//
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//data := []byte(`{"queryIp":"188.227.9.18"}`)
	data := []byte(`queryIp=188.227.9.18`)
	r := bytes.NewReader(data)
	resp, err := http.Post("https://coding.tools/my-ip-address", "application/json", r)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func getIP(prx string) string {
	tmpUrl := "https://api.ipify.org?format=json"

	proxyUrl, _ := url.Parse(prx)

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	req, err := http.NewRequest("GET", tmpUrl, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("User-Agent", )

	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	resp := gjson.ParseBytes(body)

	return resp.Get("ip").String()

}

func parseInfo(resp string) { //map[string][]string {
	var result []string
	re := regexp.MustCompile(`/geoip/"/>(.*?)</a>`)
	m := re.FindStringSubmatch(resp)
	if len(m) > 0 {
		result = append(result, m[1])
	}

	re = regexp.MustCompile(`<i class="ip-icon-small ip-icon-device-desktop"></i>\s*<div class="(.*?)">`)
	m = re.FindStringSubmatch(resp)
	if len(m) > 0 {
		re = regexp.MustCompile(fmt.Sprintf(`class="%s">\s*(.*?)\s*<`, m[1]))
		for _, m = range re.FindAllStringSubmatch(resp, -1) {
			result = append(result, m[1])
		}
	}

	fmt.Println(result)
}

//func getInfo(prx string) []byte {
//	tmpUrl := "https://2ip.ru/"
//
//	proxyUrl, _ := url.Parse(prx)
//
//	httpClient := &http.Client{
//		Transport: &http.Transport{
//			Proxy: http.ProxyURL(proxyUrl),
//		},
//	}
//
//	req, err := http.NewRequest("GET", tmpUrl, nil)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
//	req.Header.Add("User-Agent", ua.Get())
//
//	response, err := httpClient.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	defer response.Body.Close()
//
//	body, _ := io.ReadAll(response.Body)
//	return body
//
//}

//func parseTest() {
//	//var httpClient *http.Client
//	var p string
//	names := []string{"Alcyoneus"} //, "MrRobot69", "KORKUT34", "Inkin88", "Gasman1981", "Engapenga1", "DakataKalinov", "Cougey1982", "Borislavhs", "Arruolato"}
//	//prx := []string{"46acd86e65:33000997_country-am_session-2870nefz_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-m92b8q3j_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-bzt5hc7q_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-gmzshjm6_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-jzuzjvla_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-le63z5hz_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-h8uvgkgp_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-d0e8qbjn_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-xkub60vy_lifetime-1s@185.156.177.59:8377", "46acd86e65:33000997_country-am_session-nrsnjmfs_lifetime-1s@185.156.177.59:8377"}
//	//p = "http://46acd86e65:33000997_country-ru@185.156.177.59:8377" //http://46acd86e65:33000997@185.156.177.59:8377"
//	for _, player := range names {
//
//		st, err := stats.Get(player, "iPoker")
//		//resp := gjson.ParseBytes(body)
//		fmt.Println(player, string(body))
//
//		if err := os.WriteFile(player+".json", body, 0666); err != nil {
//			log.Fatal(err)
//		}
//
//		//response.Body.Close()
//	}
//}

//func getStats(name, prx string) []byte {
//	//tmpUrl := fmt.Sprintf("https://sharkscope.com/poker-statistics/networks/iPoker/players/%s?&Currency=USD", name)
//	tmpUrl := "https://api.ipify.org?format=json"
//	//tmpUrl := "https://ru.sharkscope.com/"
//	//_ = prx
//
//	proxyUrl, _ := url.Parse(prx)
//
//	httpClient := &http.Client{
//		Transport: &http.Transport{
//			Proxy: http.ProxyURL(proxyUrl),
//		},
//	}
//
//	req, err := http.NewRequest("GET", tmpUrl, nil)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
//	req.Header.Add("User-Agent", ua.Get())
//	req.Header.Add("Accept-Language", "en")
//	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
//	req.Header.Add("Username", "")
//	req.Header.Add("Password", "")
//	req.Header.Add("X-Requested-With", "XMLHttpRequest")
//	req.Header.Add("Connection", "keep-alive")
//	req.Header.Add("Referer", "https://sharkscope.com/")
//
//	response, err := httpClient.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	defer response.Body.Close()
//
//	body, _ := io.ReadAll(response.Body)
//	return body
//
//}

//if err := os.WriteFile("file.txt", []byte("Hello GOSAMPLES!"), 0666); err != nil {
//log.Fatal(err)
//}

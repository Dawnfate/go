package main

/*func main() {

	lng := "111.046906000000000"
	lat := "21.799383000000000"
	testUrl := "http://latlngto.market.alicloudapi.com/lundroid/queryLocation?lat=" + lat + "&lng=" + lng

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, testUrl,nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Authorization", "APPCODE")
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	rsp := new(Result)
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Data.Province)
	fmt.Println(rsp.Data.City)
	fmt.Println(rsp.Data.District)
}*/

type Result struct {
	Data struct {
		Province     string `json:"province"`
		StreetNumber string `json:"street_number"`
		District     string `json:"district"`
		Street       string `json:"street"`
		Lng          string `json:"lng"`
		Address      string `json:"address"`
		City         string `json:"city"`
		Reference    struct {
			Town struct {
				Distance  float64 `json:"distance"`
				Direction string  `json:"direction"`
				Name      string  `json:"name"`
				Location  struct {
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"location"`
			} `json:"town"`
			StreetNumber struct {
				Distance  float64 `json:"distance"`
				Direction string  `json:"direction"`
				Name      string  `json:"name"`
				Location  struct {
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"location"`
			} `json:"street_number"`
			Street struct {
				Distance float64 `json:"distance"`
				Name     string  `json:"name"`
				Location struct {
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"location"`
			} `json:"street"`
			LandmarkL2 struct {
				Distance  float64 `json:"distance"`
				Direction string  `json:"direction"`
				Name      string  `json:"name"`
				Location  struct {
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"location"`
			} `json:"landmark_l2"`
		} `json:"reference"`
		Nation string `json:"nation"`
		Lat    string `json:"lat"`
	} `json:"data"`
	Resp struct {
		RespMsg  string `json:"RespMsg"`
		RespCode int    `json:"RespCode"`
	} `json:"resp"`
}

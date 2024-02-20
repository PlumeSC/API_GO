package utils

import (
	"fmt"
	"io"
	"net/http"
)

func Standing() {

	url := "https://api-football-v1.p.rapidapi.com/v3/standings?season=2023&league=39"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "7ebe202ab2msh4a5e6520f04bbbap12fbf7jsn3d5fef1a70a2")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

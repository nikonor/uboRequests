package main

import (
	"fmt"
	"net/http"
	"time"

	"uboRequests"
)

func main() {

	cli := http.Client{
		Timeout: 100 * time.Millisecond,
	}

	u := uboRequests.NewRule(cli, time.Second, 10)
	req, err := http.NewRequest("GET", "http://localhost:8095/ok", nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 15; i++ {
		fmt.Printf("Try#%d", i)
		if i < 10 {
			fmt.Printf("::ждем 200\n")
		} else {
			fmt.Printf("::ждем ошибку\n")
		}
		resp, err := u.DoRequest(req)
		if err != nil {
			fmt.Printf("\terr=%s\n", err.Error())
			continue
		}
		fmt.Printf("\tresp: status=%s\n", resp.Status)
	}

}

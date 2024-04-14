package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error while sending request to %s: %v", url, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %v", err)
	}

	// NOTE: Print response body in string format
	// fmt.Println(string(body))

	// NOTE: Parse []byte to the go struct pointer
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	fmt.Println("Can not unmarshal JSON")
	// }

	return body, nil
}

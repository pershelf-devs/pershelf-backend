package helperContact

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

// DbApiMainPath is the main path for the helper API
const (
	DbApiMainPath = "/restapi/helper/v1.0"
)

// HelperRequest makes an API call by sending JSON data to the specified endpoint.
//
// Parameters:
//   - endpointPath: API endpoint path (appended to DbApiMainPath)
//   - payload: JSON data to be sent in the request body
//
// Return values:
//   - []byte: Response data received from the API
//   - error: Error that occurred during the process, nil if successful
//
// Usage:
//
//	payload, _ := json.Marshal(map[string]interface{}{"id": 123})
//	response, err := HelperRequest("/users", payload)
func HelperRequest(endpointPath string, payload []byte) ([]byte, error) {
	url := fmt.Sprintf("http://127.0.0.1:55000%s%s", DbApiMainPath, endpointPath)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.SetBody(payload)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	c := &fasthttp.Client{}
	if err := c.Do(req, resp); err != nil {
		log.Printf("(Error): error sending the POST request to database (%s): %v", url, err)
		return nil, err
	}

	bodyCopy := append([]byte(nil), resp.Body()...)

	return bodyCopy, nil
}

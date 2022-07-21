package authentication

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ClientRequest HTTP METHOD WRAPPER
type ClientRequest struct {
	Path    string
	Query   map[string][]string
	Headers map[string]string
	Body    []byte

	method             string
	Username, Password string
}

type ClientResponse struct {
	Code int   // http status code and default err code(0)
	Err  error // when program err occurred
	Body []byte
	Raw  *http.Response // be careful, response.Body can be read exactly once
}

// generate http request
func (c *Client) httpRequest(req *ClientRequest) *ClientResponse {
	var err error
	resp := &ClientResponse{}

	// generate req
	var newReq *http.Request
	reqUrl := c.Endpoint + req.Path
	if req.Body != nil {
		newReq, err = http.NewRequest(req.method, reqUrl, bytes.NewBuffer(req.Body))
	} else {
		newReq, err = http.NewRequest(req.method, reqUrl, nil)
	}
	if err != nil {
		resp.Err = err
		return resp
	}
	newReq.Header.Set("User-Agent", c.userAgent)

	// process url query string
	if req.Query != nil {
		urlV := url.Values{}
		for k, vs := range req.Query {
			if len(vs) == 1 {
				urlV.Set(k, vs[0])
			} else if len(vs) > 1 {
				for _, v := range vs {
					urlV.Add(k, v)
				}
			}
		}
		newReq.URL.RawQuery = urlV.Encode()
	}

	// process headers
	for k, v := range req.Headers {
		newReq.Header.Add(k, v)
	}
	if req.Username != "" && req.Password != "" {
		newReq.SetBasicAuth(req.Username, req.Password)
	}

	doResp, err := c.client.Do(newReq)
	if err != nil {
		resp.Err = err
		return resp
	}
	resp.Raw = doResp
	resp.Code = doResp.StatusCode

	respBody, err := ioutil.ReadAll(doResp.Body)
	if err != nil {
		resp.Err = err
		return resp
	}
	defer doResp.Body.Close()
	resp.Body = respBody

	// check if need unmarshal response body
	return resp
}

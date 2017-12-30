package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/linchongky/btcsoon/proto"
	"golang.org/x/net/context/ctxhttp"
)

func ParseQuery(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(x)
}

func HttpRequest(method string, url string, bodyType string, body io.Reader, ctx context.Context) ([]byte, error) {
	var resp *http.Response
	var err error
	Umethod := strings.ToUpper(method)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	//defer cancel()

	c := &http.Client{}
	req, err := http.NewRequest(Umethod, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err = ctxhttp.Do(ctx, c, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("req data:", string(data))
	if len(data) == 0 {
		return nil, errors.New("request timeout")
	}
	return data, nil
}

func HttpResponse(w http.ResponseWriter, code int, msg string, data interface{}) {
	resp := &proto.Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	b, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(b))
}

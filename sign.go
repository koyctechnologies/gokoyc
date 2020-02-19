package koyc

import (
	"bytes"
	"github.com/bmizerany/aws4"
	"net/http"
	"time"	
)

const (
	VERSION = "1"
	PROTOCOL = "https"
	DOMAIN = "api.koyc.in"
)

var Client = &http.Client{Timeout: 30 * time.Second}

type Keys struct {
	CustomerKey string
	ApiKey string
	AccessKey string
	SecretKey string
}

func KoycAwsSignHttpRequest(r *http.Request, keys Keys) (*http.Request, error) {
	awsKeys := &aws4.Keys{
		AccessKey: keys.AccessKey,
		SecretKey: keys.SecretKey,
	}
	newService := aws4.Service{
		Name:   "execute-api",
		Region: "ap-south-1",
	}
	r.Header.Set("x-api-key", keys.ApiKey)
	r.Header.Set("x-api-agent", "gokyc."+VERSION)
	err := newService.Sign(awsKeys, r)
	return r, err
}

func NewRequest(keys Keys, method string, path string, jsonBytes []byte) (*http.Response, error) {
	url := PROTOCOL + "://" + DOMAIN + "/"+ keys.CustomerKey+"/api/v" + VERSION + path
	r, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))

	if err == nil {
		return RequestWith(r, keys)
	}else{
		return nil, err
	}
}

func RequestWith(r *http.Request, keys Keys) (*http.Response, error) {
	r, err := KoycAwsSignHttpRequest(r, keys)
	if err != nil {
		return nil, err 
	}else{
		return Client.Do(r)
	}
}

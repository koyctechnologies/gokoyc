
## Getting Started
	
	import (
		koyc "github.com/koyctechnologies/gokoyc"
	)
	
	keys := koyc.Keys {
		CustomerKey: "REPLACE_HERE",
		ApiKey: "REPLACE_HERE",
		AccessKey: "REPLACE_HERE",
		SecretKey: "REPLACE_HERE",
	}
		
	var jsonBytes = []byte(`{"id_number":"` + SOME_PAN + `"}`)

	resp, err := koyc.NewRequest(keys, "POST", "/pan/pan-only", jsonBytes)
	if err != nil {
		t.Error("Error in making the request, error: ", err)
	}
	
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("resp.Body:", string(body)) // {"success":true|false, ...

You can also unmarshal the response to `koyc.DefaultPanVerifyResponse`
	
## Complete Test	

Here is a complete test, where you need to find and replace `REPLACE_HERE`


	package koyc_test
	
	import (
		koyc "github.com/koyctechnologies/gokoyc"
		"io/ioutil"
		"testing"
	)

	const (
		PAN_NUMBER = "REPLACE_HERE"
	)
	
	func TestKoycPan(t *testing.T) {
	
		keys := koyc.Keys {
			CustomerKey: "REPLACE_HERE",
			ApiKey: "REPLACE_HERE",
			AccessKey: "REPLACE_HERE",
			SecretKey: "REPLACE_HERE",
		}
		
		var jsonBytes = []byte(`{"id_number":"` + PAN_NUMBER+ `","dob":"REPLACE_HERE(YYYY-MM-DD)"}`)

		resp, err := koyc.NewRequest(keys, "POST", "/pan/pan-only", jsonBytes)
		if err != nil {
			t.Error("Error in making the request, error: ", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Log("resp.Body:", string(body))   // {"success":true|false, ...
	}


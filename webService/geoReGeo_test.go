package webService

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGeo(t *testing.T) {
	client := NewClient(http.DefaultClient, "")
	get, _, err := client.Geo.BatchGet([]string{"北京市朝阳区阜通东大街6号"}, nil)
	if err != nil {
		t.Error(err)
		return
	} else {
		for i, geo := range get {
			t.Log(fmt.Sprintf("%d: %+v", i, geo))
		}
	}

	code, _, err := client.Geo.Get("北京市朝阳区阜通东大街6号", nil)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(code)
	}
}

func TestReGeo(t *testing.T) {
	client := NewClient(http.DefaultClient, "")
	get, _, err := client.ReGeo.BatchGet([]string{"116.481488,39.990464"}, nil)
	if err != nil {
		t.Error(err)
		return
	} else {
		for i, reGeo := range get {
			t.Log(fmt.Sprintf("%d: %+v", i, reGeo))
		}
	}

	code, _, err := client.ReGeo.Get("116.481488,39.990464", nil)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(code)
	}
}

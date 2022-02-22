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
		return
	}
	for i, geo := range get {
		fmt.Println(i, geo)
	}
}

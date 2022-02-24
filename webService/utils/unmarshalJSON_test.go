package utils

import (
	"fmt"
	"github.com/asjdf/amap-SDK/webService"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	rList := &[]string{}
	geoResp := new(webService.GeoResp)

	genJSONSliceToStrRegexp(geoResp, []string{}, rList)
	fmt.Println(rList)
}

package webService

import (
	"github.com/dghubble/sling"
	"net/http"
)

const amapApi = "https://restapi.amap.com/v3/"

type Client struct {
	sling *sling.Sling
	Geo   *geoService   //地理编码
	Regeo *reGeoService //逆地理编码
}

func NewClient(httpClient *http.Client, key string) *Client {
	base := sling.New().Client(httpClient).Base(amapApi)
	return &Client{
		sling: base,
		Geo:   newGeoService(base.New(), key),
		Regeo: newReGeoService(base.New(), key),
	}
}

package webService

import (
	"github.com/dghubble/sling"
	"net/http"
)

type geoService struct {
	sling *sling.Sling
	Key   string
}

func newGeoService(sling *sling.Sling, key string) *geoService {
	return &geoService{
		sling: sling.Path("geocode/geo"),
		Key:   key,
	}
}

type GeoParam struct {
	Key      string `url:"key,omitempty"`
	Address  string `url:"address"`
	City     string `url:"city,omitempty"`
	Batch    bool   `url:"batch,omitempty"`
	Sig      string `url:"sig,omitempty"`
	Output   string `url:"output,omitempty"`
	Callback string `url:"callback,omitempty"`
}

type GeoResp struct {
	Status   string     `json:"status"`
	Info     string     `json:"info"`
	InfoCode string     `json:"infocode"`
	Count    string     `json:"count"`
	Geocodes []*GeoCode `json:"geocodes"`
}

type GeoCode struct {
	FormattedAddress string `json:"formatted_address"`
	Country          string `json:"country"`
	Province         string `json:"province"`
	CityCode         string `json:"citycode"`
	City             string `json:"city"`
	District         string `json:"district"`
	AdCode           string `json:"adcode"`
	Street           string `json:"street"`
	Number           string `json:"number"`
	Location         string `json:"location"`
	Level            string `json:"level"`
}

func (s *geoService) Get(addr string, param *GeoParam) (*GeoCode, *http.Response, error) {
	geoResp, resp, err := s.BatchGet([]string{addr}, param)
	if err != nil || len(geoResp) == 0 {
		return nil, resp, err
	}
	return geoResp[0], resp, nil
}

func (s *geoService) BatchGet(addr []string, param *GeoParam) ([]*GeoCode, *http.Response, error) {
	if param == nil {
		param = &GeoParam{Key: s.Key}
	} else {
		param.Key = s.Key
	}
	if len(addr) > 1 {
		param.Batch = true
	}
	if len(addr) > 0 {
		param.Address = addr[0]
		for i := 1; i < len(addr); i++ {
			param.Address += "|" + addr[i]
		}
	}
	geoResp := new(GeoResp)
	geoResp.Geocodes = make([]*GeoCode, 0)
	resp, err := s.sling.New().Get("").QueryStruct(param).ReceiveSuccess(geoResp)
	if err != nil {
		return nil, resp, err
	}
	return geoResp.Geocodes, resp, nil
}

type reGeoService struct {
	sling *sling.Sling
	Key   string
}

func newReGeoService(sling *sling.Sling, key string) *reGeoService {
	return &reGeoService{
		sling: sling.Path("geocode/regeo"),
		Key:   key,
	}
}

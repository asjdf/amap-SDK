package webService

import (
	"errors"
	"github.com/dghubble/sling"
	"net/http"
)

type geoService struct {
	sling *sling.Sling
	key   string
}

func newGeoService(sling *sling.Sling, key string) *geoService {
	return &geoService{
		sling: sling.Path("geocode/geo"),
		key:   key,
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
		param = &GeoParam{Key: s.key}
	} else {
		param.Key = s.key
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
	} else if geoResp.Status != "1" {
		return nil, resp, errors.New(geoResp.Info)
	}
	return geoResp.Geocodes, resp, nil
}

type reGeoService struct {
	sling *sling.Sling
	key   string
}

func newReGeoService(sling *sling.Sling, key string) *reGeoService {
	return &reGeoService{
		sling: sling.Path("geocode/regeo"),
		key:   key,
	}
}

type ReGeoParam struct {
	Key        string `url:"key"`
	Location   string `url:"location"`
	PoiType    string `url:"poitype,omitempty"`
	Radius     string `url:"radius,omitempty"`
	Extensions string `url:"extensions,omitempty"`
	Batch      bool   `url:"batch,omitempty"`
	RoadLevel  int    `url:"roadlevel,omitempty"`
	Sig        string `url:"sig,omitempty"`
	Output     string `url:"output,omitempty"`
	Callback   string `url:"callback,omitempty"`
	HomeOrCorp int    `url:"homeorcorp,omitempty"`
}

type ReGeoResp struct {
	Status    string       `json:"status"`
	Info      string       `json:"info"`
	InfoCode  string       `json:"infocode"`
	ReGeoCode []*ReGeoCode `json:"regeocodes"`
}

type ReGeoCode struct {
	FormattedAddress string `json:"formatted_address"`
	AddressComponent struct {
		Country      string   `json:"country"`
		Province     string   `json:"province"`
		City         []string `json:"city"`
		CityCode     string   `json:"citycode"`
		District     string   `json:"district"`
		AdCode       string   `json:"adcode"`
		Township     string   `json:"township"`
		TownCode     string   `json:"towncode"`
		Neighborhood struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"neighborhood"`
		Building struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"building"`
		StreetNumber struct {
			Street    string `json:"street"`
			Number    string `json:"number"`
			Location  string `json:"location"`
			Direction string `json:"direction"`
			Distance  string `json:"distance"`
		} `json:"streetNumber"`
		BusinessAreas []struct {
			Location string `json:"location"`
			Name     string `json:"name"`
			Id       string `json:"id"`
		} `json:"businessAreas"`
	} `json:"addressComponent"`
	Pois []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Tel          string `json:"tel"`
		Direction    string `json:"direction"`
		Distance     string `json:"distance"`
		Location     string `json:"location"`
		Address      string `json:"address"`
		PoiWeight    string `json:"poiweight"`
		BusinessArea string `json:"businessarea"`
	} `json:"pois"`
	Roads []struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Direction string `json:"direction"`
		Distance  string `json:"distance"`
		Location  string `json:"location"`
	} `json:"roads"`
	RoadInters []struct {
		Direction  string `json:"direction"`
		Distance   string `json:"distance"`
		Location   string `json:"location"`
		FirstId    string `json:"first_id"`
		FirstName  string `json:"first_name"`
		SecondId   string `json:"second_id"`
		SecondName string `json:"second_name"`
	} `json:"roadinters"`
	Aois []struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		AdCode   string `json:"adcode"`
		Location string `json:"location"`
		Area     string `json:"area"`
		Distance string `json:"distance"`
		Type     string `json:"type"`
	} `json:"aois"`
}

func (s *reGeoService) Get(location string, param *ReGeoParam) (*ReGeoCode, *http.Response, error) {
	ReGeoResp, resp, err := s.BatchGet([]string{location}, param)
	if err != nil || len(ReGeoResp) == 0 {
		return nil, resp, err
	}
	return ReGeoResp[0], resp, nil
}

func (s *reGeoService) BatchGet(location []string, param *ReGeoParam) ([]*ReGeoCode, *http.Response, error) {
	if param == nil {
		param = &ReGeoParam{Key: s.key}
	} else {
		param.Key = s.key
	}
	param.Batch = true
	if len(location) > 0 {
		param.Location = location[0]
		for i := 1; i < len(location); i++ {
			param.Location += "|" + location[i]
		}
	}
	reGeoResp := new(ReGeoResp)
	reGeoResp.ReGeoCode = make([]*ReGeoCode, 0)
	resp, err := s.sling.New().Get("").QueryStruct(param).ReceiveSuccess(reGeoResp)
	if err != nil {
		return nil, resp, err
	} else if reGeoResp.Status != "1" {
		return nil, resp, errors.New(reGeoResp.Info)
	}
	return reGeoResp.ReGeoCode, resp, nil
}

package service

import (
	"time"

	gjson "github.com/paulmach/go.geojson"
)

type (
	// GEOJSON Struct
	Polygon struct {
		Type        string        `json:"type,omitempty"`
		Coordinates [][][]float64 `json:"coordinates"`
	}
	Point struct {
		Type        string    `json:"type,omitempty"`
		Coordinates []float64 `json:"coordinates,coordinates"`
	}

	// LineString
	LineString struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	}

	// location
	AddLocationRequest struct {
		LocationName   string  `json:"location_name,omitempty"`
		IsActive       int     `json:"is_active,omitempty"`
		CityId         int64   `json:"city_id,omitempty"`
		ProvinceId     int64   `json:"province_id,omitempty"`
		Shape          Polygon `json:"shape,omitempty"`
		LocationTypeId int64   `json:"location_type"`
	}
	AddLocationResponse struct {
		Id           int64  `json:"id,omitempty"`
		LocationName string `json:"location_name,omitempty"`
		IsActive     int    `json:"is_active,omitempty"`
		CityId       int64  `json:"city_id,omitempty"`
	}

	GetLocationsRequest struct {
		Page        int `json:"page"`
		ItemPerPage int `json:"item_per_page"`
	}
	AdministratifItem struct {
		ProvinceCode int64  `json:"province_code,omitempty"`
		CitiesId     int64  `json:"cities_id,omitempty"`
		ProvinceName string `json:"province_name,omitempty"`
		CityName     string `json:"city_name,omitempty"`
	}
	LocationItem struct {
		Id                int64             `json:"id,omitempty"`
		LocationName      string            `json:"location_name,omitempty"`
		Area              float64           `json:"area,omitempty"`
		AdministratifArea AdministratifItem `json:"administratif_area,omitempty"`
		Shape             *gjson.Feature    `json:"shape,omitempty"`
	}
	GetLocationsResponse struct {
		Count     int64          `json:"count"`
		Locations []LocationItem `json:"locations"`
	}

	// nearby location
	Position struct {
		Lat    float64 `json:"lat"`
		Long   float64 `json:"long"`
		Radius int64   `json:"radius,omitempty"`
	}

	NearbyFilter struct {
		WithType bool   `json:"with_type"`
		TypeName string `json:"type_name"`
	}

	GetNearbyLocationRequest struct {
		Count           int64        `json:"count"`
		ShowDistance    bool         `json:"show_distance"`
		Filter          NearbyFilter `json:"nearby_filter"`
		CurrentPosition Position     `json:"current_position"`
	}

	EstimatedTime struct {
		Walking    float64 `json:"walking"`
		MotorCycle float64 `json:"motorcycle"`
		DrivingCar float64 `json:"driving_car"`
	}

	NearbyLocationItem struct {
		Id            int64         `json:"id"`
		LocationName  string        `json:"location_name"`
		Point         Position      `json:"point,omitempty"`
		Distance      *float64      `json:"distance,omitempty"`
		Detail        string        `json:"detail,omitempty"`
		EstimatedTime EstimatedTime `json:"estimated_time"`
		LocationType  string        `json:"location_type"`
	}
	GetNearbyLocationResponse struct {
		Count          int                  `json:"count"`
		NearbyLocation []NearbyLocationItem `json:"nearby_locations,omitempty"`
	}
	// sub location
	AddSubLocationRequest struct {
		SubLocationName string  `json:"sub_location_name,omitempty"`
		LocationId      int64   `json:"location_id,omitempty"`
		IsGeofence      int64   `json:"is_geofence"`
		Shape           Polygon `json:"shape,omitempty"`
		Point           Point   `json:"point,omitempty"`
	}

	// add location type
	AddLocationTypeRequest struct {
		LocationType string `json:"location_type"`
	}
	AddLocationTypeResponse struct {
		Id        int64     `json:"id"`
		CreatedAt time.Time `json:"created_at"`
	}
	// get location type
	LocationTypeItem struct {
		Id        int64     `json:"id"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
	}
	GetLocationTypeResponse struct {
		Count        int64              `json:"count,omitempty"`
		LocationType []LocationTypeItem `json:"location_type"`
	}

	// EditLocation
	EditLocationRequest struct {
		Locationid   int64   `json:"location_id"`
		LocationType int64   `json:"location_type"`
		CityId       int64   `json:"city_id,omitempty"`
		ProvinceId   int64   `json:"province_id,omitempty"`
		LocationName string  `json:"location_name,omitempty"`
		IsActive     int     `json:"is_active,omitempty"`
		Shape        Polygon `json:"shape,omitempty"`
	}
	EditLocationResponse struct {
		Message    string    `json:"message"`
		ModifiedAt time.Time `json:"modified_at"`
	}
	// detail location
	DetailLocationRequest struct {
		LocationId int64 `json:"location_id"`
	}

	SublocationItems struct {
		Id              int64          `json:"id,omitempty"`
		SubLocationName string         `json:"sub_location_name"`
		IsGeofence      bool           `json:"is_geofence"`
		Shape           *gjson.Feature `json:"polygon"`
		Point           *gjson.Feature `json:"point"`
	}

	DetailLocationResponse struct {
		Id                int64              `json:"location_id"`
		LocationName      string             `json:"location_name"`
		LocationType      string             `json:"location_type"`
		AdministrafifItem AdministratifItem  `json:"administratif_info"`
		Shape             *gjson.Feature     `json:"shape"`
		CountSubLocation  int64              `json:"count_sublocation"`
		SublocationItems  []SublocationItems `json:"sub_locations"`
	}
	// direction/route location
	GetDirectionLocationRequest struct {
		FromLocation         string `json:"from"`
		DestionationLocation string `json:"destination"`
		// type of direction Ex: Walking,driving car etc...
		Profile string  `json:"profile"`
		SrcLat  float64 `json:"-"`
		SrcLong float64 `json:"-"`
		DstLat  float64 `json:"-"`
		DstLong float64 `json:"-"`
	}
	GetDirectionLocationResponse struct {
		From        string     `json:"from"`
		Destination string     `json:"destinantion"`
		Distance    float64    `json:"distance"`
		Duration    float64    `json:"duration"`
		Profile     string     `json:"profile"`
		Route       LineString `json:"route"`
	}

	// properties payload
	PropertiesPayload struct {
		LocationName  string
		SubLocationId int64
		LocationId    int64
		LocationType  string
		Area          float64
		ProvinceName  string
		CityName      string
	}
)

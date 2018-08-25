package model

import (
	"time"
)

const (
	SORT_PARAM_LICENSETYPE = "license_type"
	SORT_PARAM_TYPE        = "type"
	SORT_PARAM_STATE       = "state"
	SORT_PARAM_CITY        = "city"
	SORT_PARAM_NAME        = "name"
	SORT_PARAM_DISTANCE    = "distance"

	SORT_ORDER_ASC  = "asc"
	SORT_ORDER_DESC = "desc"
)

type RequestGeo struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latiude"`
	SortParam string  `json:"sort"`
	SortOrder string  `json:"order"`
	Limit     int     `json:"limit"`
	Radius    int     `json:"radius"`
}

type ResponseGeo struct {
	Locations []MITLocations  `json:"locations,omitempty"`
	Summary   ResponseSummary `json:"summary"`
}

type ResponseSummary struct {
	Count    int        `json:"count"`
	Duration string     `json:"duration"`
	Date     time.Time  `json:"date"`
	Params   RequestGeo `json:"parameters"`
}

type MITLocations struct {
	Id           int     `bson:"id" json:"id"`
	Name         string  `bson:"name" json:"name"`
	Slug         string  `bson:"slug" json:"slug"`
	State        string  `bson:"state" json:"state"`
	City         string  `bson:"city" json:"city"`
	Type         string  `bson:"type" json:"type"`
	Wmid         int     `bson:"wmid" json:"wmid"`
	Latitude     float64 `bson:"latitude" json:"latitude"`
	Longitude    float64 `bson:"longitude" json:"longitude"`
	WebURL       string  `bson:"web_url" json:"web_url"`
	PackageLevel string  `bson:"package_level" json:"package_level"`
	FeatureOrder int     `bson:"feature_order" json:"feature_order"`
	Ranking      float64 `bson:"ranking" json:"ranking"`
	Rating       float64 `bson:"rating" json:"rating"`
	ReviewsCount int     `bson:"reviews_count" json:"reviews_count"`
	AvatarImage  struct {
		SmallURL string `bson:"small_url" json:"small_url"`
	} `bson:"avatar_image" json:"avatar_image"`
	LicenseType            string      `bson:"license_type" json:"license_type"`
	Address                string      `bson:"address" json:"address"`
	Distance               interface{} `bson:"distance" json:"distance"`
	ZipCode                string      `bson:"zip_code" json:"zip_code"`
	Timezone               string      `bson:"timezone" json:"timezone"`
	IntroBody              string      `bson:"intro_body" json:"intro_body"`
	StaticMapURL           string      `bson:"static_map_url" json:"static_map_url"`
	OpenNow                bool        `bson:"open_now" json:"open_now"`
	ClosesIn               interface{} `bson:"closes_in" json:"closes_in"`
	TodaysHoursStr         string      `bson:"todays_hours_str" json:"todays_hours_str"`
	MinAge                 int         `bson:"min_age" json:"min_age"`
	MenuItemsCount         int         `bson:"menu_items_count" json:"menu_items_count"`
	VerifiedMenuItemsCount int         `bson:"verified_menu_items_count" json:"verified_menu_items_count"`
	EndorsementBadgeCount  int         `bson:"endorsement_badge_count" json:"endorsement_badge_count"`
	IsPublished            bool        `bson:"is_published" json:"is_published"`
	OnlineOrdering         struct {
		EnabledForPickup   bool `bson:"enabled_for_pickup" json:"enabled_for_pickup"`
		EnabledForDelivery bool `bson:"enabled_for_delivery" json:"enabled_for_delivery"`
	} `bson:"online_ordering" json:"online_ordering"`
	Location GeoJson `bson:"location" json:"location"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

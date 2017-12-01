package revgeocode

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	 geo "github.com/martinlindhe/google-geolocate"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-revgeocode")

const (
	ivAPIkey = "apiKey"
	ivLat    = "lat"
	ivLang   = "lang"
	ovLocation = "location"
)

// GeoCodeActivity is a Geocode Activity implementation
type RevGeoCodeActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new GeoCodeActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &RevGeoCodeActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *RevGeoCodeActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *RevGeoCodeActivity) Eval(context activity.Context) (done bool, err error) {

	apiKey := context.GetInput(ivAPIkey).(string)
	lat := context.GetInput(ivLat).(float64)
	lang := context.GetInput(ivLang).(float64)
	location:= "location"

	gclient := geo.NewGoogleGeo(apiKey)
	gpoint := geo.Point{Lat: lat, Lng: lang}
	resp, _ := gclient.ReverseGeocode(&gpoint)

  log.Debug("Response:", resp)

  context.SetOutput(location, resp)
	return true, nil
}

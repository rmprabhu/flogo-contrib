package revgeocode

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/martinlindhe/google-geolocate"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-revgeocode")

const (
	ivAPIkey = "apiKey"
	ivLat    = "lat"
	ivLang   = "lang"
)

// GeoCodeActivity is a Geocode Activity implementation
type RevGeoCodeActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new TwilioActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &RevGeoCodeActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *RevGeoCodeActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *GeoCodeActivity) Eval(context activity.Context) (done bool, err error) {

	apiKey := context.GetInput(ivAPIkey).(string)
	lat := context.GetInput(ivLat).(string)
	lang := context.GetInput(ivLang).(string)

	gclient := geo.NewGoogleGeo(apiKey)
	gpoint := geo.Point{Lat: lat, Lng: lang}
	resp, _, err := gclient.ReverseGeocode(&gpoint)

	if err != nil {
		log.Error("Error translating location:", err)
	}

	log.Debug("Response:", resp)

	return resp
}

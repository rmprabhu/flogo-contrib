package twilliocall

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/sfreiberg/gotwilio"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-twilio")

const (
	ivAcctSID   = "accountSID"
	ivAuthToken = "authToken"
	ivFrom      = "from"
	ivTo        = "to"
	ivURL       = "url"
)

// TwilioActivity is a Twilio Activity implementation
type TwilioActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new TwilioActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TwilioActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *TwilioActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *TwilioActivity) Eval(context activity.Context) (done bool, err error) {

	accountSID := context.GetInput(ivAcctSID).(string)
	authToken := context.GetInput(ivAuthToken).(string)
	from := context.GetInput(ivFrom).(string)
	to := context.GetInput(ivTo).(string)
	url := context.GetInput(ivURL).(string)

	//Added callbackparams ivURL

	twilio := gotwilio.NewTwilioClient(accountSID, authToken)

	callbackParams := gotwilio.NewCallbackParameters(url)

	resp, _, err := twilio.CallWithUrlCallbacks(from, to, callbackParams)

	if err != nil {
		log.Error("Error making voice call:", err)
	}

	log.Debug("Response:", resp)

	return true, nil
}

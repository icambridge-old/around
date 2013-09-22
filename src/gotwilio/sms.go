package gotwilio

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// SmsResponse is returned after a text/sms message is posted to Twilio
type SmsResponse struct {
	XMLName     xml.Name `xml:"TwilioResponse"`
	Sid         string   `xml:"SMSMessage>Sid"`
	DateCreated string   `xml:"SMSMessage>DateCreated"`
	DateUpdate  string   `xml:"SMSMessage>DateUpdated"`
	DateSent    string   `xml:"SMSMessage>DateSent"`
	AccountSid  string   `xml:"SMSMessage>AccountSid"`
	To          string   `xml:"SMSMessage>To"`
	From        string   `xml:"SMSMessage>From"`
	Body        string   `xml:"SMSMessage>Body"`
	Status      string   `xml:"SMSMessage>Status"`
	Direction   string   `xml:"SMSMessage>Direction"`
	ApiVersion  string   `xml:"SMSMessage>ApiVersion"`
	Price       float32  `xml:"SMSMessage>Price"`
	Url         string   `xml:"SMSMessage>Uri"`
}

// Returns SmsResponse.DateCreated as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateCreatedAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateCreated)
}

// Returns SmsResponse.DateUpdate as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateUpdateAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateUpdate)
}

// Returns SmsResponse.DateSent as a time.Time object
// instead of a string.
func (sms *SmsResponse) DateSentAsTime() (time.Time, error) {
	return time.Parse(time.RFC1123Z, sms.DateSent)
}

// SendTextMessage uses Twilio to send a text message.
// See http://www.twilio.com/docs/api/rest/sending-sms for more information.
func (twilio *Twilio) SendSMS(from, to, body, statusCallback, applicationSid string) (*SmsResponse, *Exception, error) {
	var smsResponse *SmsResponse
	var exception *Exception
	twilioUrl := twilio.BaseUrl + "/Accounts/" + twilio.AccountSid + "/SMS/Messages"

	formValues := url.Values{}
	formValues.Set("From", from)
	formValues.Set("To", to)
	formValues.Set("Body", body)
	if statusCallback != "" {
		formValues.Set("StatusCallback", statusCallback)
	}
	if applicationSid != "" {
		formValues.Set("ApplicationSid", applicationSid)
	}

	res, err := twilio.post(formValues, twilioUrl)
	if err != nil {
		return smsResponse, exception, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return smsResponse, exception, err
	}

	if res.StatusCode != http.StatusCreated {
		exception = new(Exception)
		err = xml.Unmarshal(responseBody, exception)

		// We aren't checking the error because we don't actually care.
		// It's going to be passed to the client either way.
		return smsResponse, exception, err
	}

	smsResponse = new(SmsResponse)
	err = xml.Unmarshal(responseBody, smsResponse)
	return smsResponse, exception, err
}

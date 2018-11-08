package onesignal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/janmir/gorequest"
)

var (
	f = fmt.Sprintf
)

const (
	createNotificationURL  = "https://onesignal.com/api/v1/notifications"
	viewAppsURL            = "https://onesignal.com/api/v1/apps"
	createApp              = "https://onesignal.com/api/v1/apps"
	addADevice             = "https://onesignal.com/api/v1/players"
	editDevice             = "https://onesignal.com/api/v1/players/%s"                                 //PlayerID
	viewAnAppURL           = "https://onesignal.com/api/v1/apps/%s"                                    //AppID
	cancelNotificationURL  = "https://onesignal.com/api/v1/notifications/%s?app_id=%s"                 //NotificationID, AppID
	updateAnApp            = "https://onesignal.com/api/v1/apps/%s"                                    //AppID
	viewDevices            = "https://onesignal.com/api/v1/players?app_id=%s&limit=%d&offset=%d"       //AppID, limit, offset
	viewDevice             = "https://onesignal.com/api/v1/players/%s?app_id=%s"                       //PlayerID, AppID
	newSession             = "https://onesignal.com/api/v1/players/%s/on_session"                      //PlayerID
	newPurchase            = "https://onesignal.com/api/v1/players/:id/on_purchase"                    //PlayerID
	incrementSessionLength = "https://onesignal.com/api/v1/players/:id/on_focus"                       //PlayerID
	csvExport              = "https://onesignal.com/api/v1/players/csv_export?app_id=:app_id"          //AppID
	viewNotification       = "https://onesignal.com/api/v1/notifications/%s?app_id=%s"                 //NotificationID, AppID
	viewNotifications      = "https://onesignal.com/api/v1/notifications?app_id=%s&limit=%s&offset=%s" //AppID, lmit, offset
	trackOpen              = "https://onesignal.com/api/v1/notifications/%s"                           //NotificationID
)

//OpenSignal Object
type OpenSignal struct {
	Client  *gorequest.SuperAgent
	AuthKey string
	APIKey  string
}

//New creates a new OpenSignal object
func New(auth, api string) OpenSignal {
	o := OpenSignal{}
	o.AuthKey = auth
	o.APIKey = api
	o.Client = gorequest.New()
	return o
}

func catch(res gorequest.Response, body []byte) error {
	//check statusCode
	if res.StatusCode != http.StatusOK {
		b := Response{}
		//parseBody
		err := json.Unmarshal(body, &b)
		if err != nil {
			return err
		}

		if len(b.Errors) > 0 {
			msg := ""
			for _, er := range b.Errors {
				msg += er + ". "
			}
			return errors.New(msg)
		}
	}

	return nil
}

//CreateNotification Sends notifications to your users
//Post -> https://onesignal.com/api/v1/notifications
func (o OpenSignal) CreateNotification(req NotificationRequest) (Response, error) {
	strResponse := Response{}

	res, body, errs := o.Client.Post(createNotificationURL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(req).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//CancelNotification Stop a scheduled or currently outgoing notification
//Delete -> https://onesignal.com/api/v1/notifications/:id?app_id=:app_id
func (o OpenSignal) CancelNotification(notificationID, appID string) (Response, error) {
	strResponse := Response{}

	URL := f(cancelNotificationURL, notificationID, appID)
	res, body, errs := o.Client.Delete(URL).
		Set("Authorization", "Basic "+o.APIKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewApps View the details of all of your current OneSignal apps
//Get -> https://onesignal.com/api/v1/apps
func (o OpenSignal) ViewApps() ([]App, error) {
	strResponse := []App{}

	res, body, errs := o.Client.Get(viewAppsURL).
		Set("Authorization", "Basic "+o.AuthKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewAnApp View the details of a single OneSignal app
//Get -> https://onesignal.com/api/v1/apps/:id
func (o OpenSignal) ViewAnApp(appID string) (App, error) {
	strResponse := []App{}

	URL := f(viewAnAppURL, appID)
	res, body, errs := o.Client.Get(URL).
		Set("Authorization", "Basic "+o.AuthKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	out := App{}
	if len(strResponse) > 0 {
		out = strResponse[0]
	}
	return out, err
}

//CreateAnApp Creates a new OneSignal app
//Post -> https://onesignal.com/api/v1/apps
func (o OpenSignal) CreateAnApp(app App) (App, error) {
	strResponse := App{}

	res, body, errs := o.Client.Post(createApp).
		Send(app).
		Set("Authorization", "Basic "+o.AuthKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//UpdateAnApp Updates the name or configuration settings of an existing OneSignal app
//Put -> https://onesignal.com/api/v1/apps/:id
func (o OpenSignal) UpdateAnApp(app App) (App, error) {
	strResponse := App{}

	URL := f(updateAnApp, app.ID)
	res, body, errs := o.Client.Put(URL).
		Send(app).
		Set("Authorization", "Basic "+o.AuthKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewDevices View the details of multiple devices in one of your OneSignal apps
//Get -> https://onesignal.com/api/v1/players?app_id=:app_id&limit=:limit&offset=:offset
func (o OpenSignal) ViewDevices(appID string, limit, offset int) (Devices, error) {
	strResponse := Devices{}

	URL := f(viewDevices, appID, limit, offset)
	res, body, errs := o.Client.Get(URL).
		Set("Authorization", "Basic "+o.APIKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewDevice View the details of an existing device in one of your OneSignal apps
//Get -> https://onesignal.com/api/v1/players/:id
func (o OpenSignal) ViewDevice(appID, playerID string) (Player, error) {
	strResponse := Player{}

	URL := f(viewDevice, playerID, appID)
	res, body, errs := o.Client.Get(URL).
		Set("Authorization", "Basic "+o.APIKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//AddADevice Register a new device to one of your OneSignal apps
//Post -> https://onesignal.com/api/v1/players
//Warning Don't use, this API endpoint is designed to be used from our open source Mobile and
//Web Push SDKs. It is not designed for developers to use it directly, unless instructed to
//do so by OneSignal support.
func (o OpenSignal) AddADevice(device Device) (Response, error) {
	strResponse := Response{}

	res, body, errs := o.Client.Post(addADevice).
		Set("Authorization", "Basic "+o.APIKey).
		Send(device).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//EditDevice Update an existing device in one of your OneSignal apps
//Put -> https://onesignal.com/api/v1/players/:id
//Warning Instead of using this REST API call we recommend using our Mobile
//Web SDK methods. Changes values with this REST API call may create
//synchronization issues with the SDK.
func (o OpenSignal) EditDevice(playerID string, device Device) (Response, error) {
	strResponse := Response{}

	URL := f(editDevice, playerID)
	res, body, errs := o.Client.Put(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(device).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//NewSession Update a device's session information
//Post -> https://onesignal.com/api/v1/players/:id/on_session
func (o OpenSignal) NewSession(playerID string, session Session) (Response, error) {
	strResponse := Response{}

	URL := f(newSession, playerID)
	res, body, errs := o.Client.Post(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(session).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//NewPurchase Track a new purchase in your app
//Post -> https://onesignal.com/api/v1/players/:id/on_purchase
func (o OpenSignal) NewPurchase(playerID string, purchase Purchase) (Response, error) {
	strResponse := Response{}

	URL := f(newPurchase, playerID)
	res, body, errs := o.Client.Post(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(purchase).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//IncrementSessionLength Update a device's session length upon app resuming
//Post -> https://onesignal.com/api/v1/players/:id/on_focus
func (o OpenSignal) IncrementSessionLength(playerID string, sessionLength SessionLength) (Response, error) {
	strResponse := Response{}

	URL := f(newPurchase, playerID)
	res, body, errs := o.Client.Post(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(sessionLength).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//CSVExport Generate a compressed CSV export of all of your current user data
//Post -> https://onesignal.com/api/v1/players/csv_export?app_id=:app_id
func (o OpenSignal) CSVExport(appID string, csv CSV) (Response, error) {
	strResponse := Response{}

	URL := f(csvExport, appID)
	res, body, errs := o.Client.Post(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(csv).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewNotification View the details of a single notification
//Get -> https://onesignal.com/api/v1/notifications/:id?app_id=:app_id
func (o OpenSignal) ViewNotification(notificationID, appID string) (Notification, error) {
	strResponse := Notification{}

	URL := f(viewNotification, notificationID, appID)
	res, body, errs := o.Client.Get(URL).
		Set("Authorization", "Basic "+o.APIKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//ViewNotifications View the details of multiple notifications
//Get -> https://onesignal.com/api/v1/notifications?app_id=:app_id&limit=:limit&offset=:offset
func (o OpenSignal) ViewNotifications(appID string, limit, offset int) (Notifications, error) {
	strResponse := Notifications{}

	URL := f(viewNotifications, appID, limit, offset)
	res, body, errs := o.Client.Get(URL).
		Set("Authorization", "Basic "+o.APIKey).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

//TrackOpen Track when users open a notification
//Put -> https://onesignal.com/api/v1/notifications/:id
func (o OpenSignal) TrackOpen(notificationID string, track Track) (Response, error) {

	strResponse := Response{}

	URL := f(trackOpen, notificationID)
	res, body, errs := o.Client.Put(URL).
		Set("Authorization", "Basic "+o.APIKey).
		Send(track).
		EndStruct(&strResponse)
	err := catch(res, body)
	if err == nil {
		for _, e := range errs {
			if e != nil {
				err = e
				break
			}
		}
	}
	return strResponse, err
}

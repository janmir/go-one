package one

import (
	"fmt"
	"testing"
	"time"
)

var (
	playerID  = "20562684-494d-4b06-af41-868a40a62e1b"
	authKey   = "OTI5NmU1ZmQtM2UzMS00ZTg5LWI0OTctYmE2NDdhMzJhNGRh"
	apiKey    = "MTdiZWNjMTEtYzA0Yy00MmQ5LWFhYzMtZjE4NDQyYzJhZDgw"
	appID     = "87a38029-bda5-4a10-a4c9-697492b36d8f"
	notifID   = "e5d74486-ce0c-4bb6-9660-2fea19eb2920"
	oneSignal = New(authKey, apiKey)
)

func TestCreateNotification(t *testing.T) {
	return
	//make a time 2 mins from now
	timeNow := time.Now().Local()
	timeMinuteFromNow := timeNow.Add(time.Minute * 2)
	out, err := oneSignal.CreateNotification(NotificationRequest{
		AppID: appID,
		Filters: []Filter{
			{
				Field:    "tag",
				Key:      "user",
				Relation: "=",
				Value:    "jp.miranda",
			},
			{
				Operator: "=",
			},
			{
				Field:    "tag",
				Key:      "realm",
				Relation: "=",
				Value:    "odtr",
			},
		},
		TemplateID:        "5ad9bae6-7930-43ea-9c17-22efd98bcc08",
		URL:               "https://odtr.awsys-i.com/jp-odtr/DTRMainLoginv2.aspx",
		DelayedOption:     "timezone",
		DeliveryTimeOfDay: timeMinuteFromNow.Format("03:04PM"),
		WebIcon:           "https://notif.janmir.me/static/icon5.png",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestCancelNotification(t *testing.T) {
	return
	out, err := oneSignal.CancelNotification("something", appID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewApps(t *testing.T) {
	return
	out, err := oneSignal.ViewApps()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewAnApp(t *testing.T) {
	return
	out, err := oneSignal.ViewAnApp("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewDevices(t *testing.T) {
	return
	out, err := oneSignal.ViewDevices(appID, 100, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewDevice(t *testing.T) {
	return
	out, err := oneSignal.ViewDevice(appID, playerID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestAddDevice(t *testing.T) {
	return
	out, err := oneSignal.AddADevice(Device{
		AppID:       "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		Identifier:  "ce777617da7f548fe7a9ab6febb56cf39fba6d382000c0395666288d961ee566",
		Language:    "en",
		Timezone:    -28800,
		GameVersion: "1.0",
		DeviceOs:    "7.0.4",
		DeviceType:  0,
		DeviceModel: "iPhone 8,2",
		Tags:        map[string]string{"a": "1", "foo": "bar"},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewNotif(t *testing.T) {
	return
	out, err := oneSignal.ViewNotification(notifID, appID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestViewNotifs(t *testing.T) {
	return
	out, err := oneSignal.ViewNotifications(appID, 100, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

func TestTrackOpen(t *testing.T) {
	return
	out, err := oneSignal.TrackOpen(notifID, Track{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", out)
}

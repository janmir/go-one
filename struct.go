package one

//App structure
type App struct {
	ID                               string `json:"id"`
	Name                             string `json:"name,omitempty"`
	Players                          int    `json:"players,omitempty"`
	MessagablePlayers                int    `json:"messagable_players,omitempty"`
	UpdatedAt                        string `json:"updated_at,omitempty"`
	CreatedAt                        string `json:"created_at,omitempty"`
	GcmKey                           string `json:"gcm_key,omitempty"`
	ChromeKey                        string `json:"chrome_key,omitempty"`
	ChromeWebOrigin                  string `json:"chrome_web_origin,omitempty"`
	ChromeWebGcmSenderID             string `json:"chrome_web_gcm_sender_id,omitempty"`
	ChromeWebDefaultNotificationIcon string `json:"chrome_web_default_notification_icon,omitempty"`
	ChromeWebSubDomain               string `json:"chrome_web_sub_domain,omitempty"`
	ApnsEnv                          string `json:"apns_env,omitempty"`
	ApnsCertificates                 string `json:"apns_certificates,omitempty"`
	SafariApnsCertificate            string `json:"safari_apns_certificate,omitempty"`
	SafariSiteOrigin                 string `json:"safari_site_origin,omitempty"`
	SafariPushID                     string `json:"safari_push_id,omitempty"`
	SafariIcon1616                   string `json:"safari_icon_16_16,omitempty"`
	SafariIcon3232                   string `json:"safari_icon_32_32,omitempty"`
	SafariIcon6464                   string `json:"safari_icon_64_64,omitempty"`
	SafariIcon128128                 string `json:"safari_icon_128_128,omitempty"`
	SafariIcon256256                 string `json:"safari_icon_256_256,omitempty"`
	SiteName                         string `json:"site_name,omitempty"`
	BasicAuthKey                     string `json:"basic_auth_key,omitempty"`
}

//Response response json
type Response struct {
	ID         string   `json:"id,omitempty"`
	Success    string   `json:"success,omitempty"`
	Recipients int      `json:"recipients,omitempty"`
	Errors     []string `json:"errors"`
	CSVFileURL string   `json:"csv_file_url,omitempty"`
}

//NotificationRequest used in creating notifications
type NotificationRequest struct {
	AppID             string            `json:"app_id"`
	Filters           []Filter          `json:"filters,omitempty"`
	TemplateID        string            `json:"template_id,omitempty"`
	URL               string            `json:"url,omitempty"`
	DelayedOption     string            `json:"delayed_option,omitempty"`
	DeliveryTimeOfDay string            `json:"delivery_time_of_day,omitempty"`
	WebIcon           string            `json:"chrome_web_icon,omitempty"`
	Data              map[string]string `json:"data,omitempty"`
	Contents          map[string]string `json:"contents,omitempty"`
	IncludedSegments  []string          `json:"included_segments,omitempty"`
}

//Filter struct
type Filter struct {
	Field    string `json:"field,omitempty"`
	Key      string `json:"key,omitempty"`
	Relation string `json:"relation,omitempty"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator,omitempty"`
}

//Devices device structure
type Devices struct {
	TotalCount int      `json:"total_count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
	Players    []Player `json:"players"`
}

//Player struct for device players
type Player struct {
	Identifier        string            `json:"identifier"`
	SessionCount      int               `json:"session_count"`
	Language          string            `json:"language"`
	Timezone          int               `json:"timezone"`
	GameVersion       string            `json:"game_version"`
	DeviceOs          string            `json:"device_os"`
	DeviceType        int               `json:"device_type"`
	DeviceModel       string            `json:"device_model"`
	AdID              interface{}       `json:"ad_id"`
	Tags              map[string]string `json:"tags"`
	LastActive        int               `json:"last_active"`
	AmountSpent       float64           `json:"amount_spent"`
	CreatedAt         int               `json:"created_at"`
	InvalidIdentifier bool              `json:"invalid_identifier"`
	BadgeCount        int               `json:"badge_count"`
}

//Device struct device
type Device struct {
	AppID       string            `json:"app_id"`
	Identifier  string            `json:"identifier"`
	Language    string            `json:"language"`
	Timezone    int               `json:"timezone"`
	GameVersion string            `json:"game_version"`
	DeviceOs    string            `json:"device_os"`
	DeviceType  int               `json:"device_type"`
	DeviceModel string            `json:"device_model"`
	Tags        map[string]string `json:"tags"`
}

//Session structure
type Session struct {
	Language    string `json:"language"`
	Timezone    int    `json:"timezone"`
	GameVersion string `json:"game_version"`
	DeviceOs    string `json:"device_os"`
}

//SessionLength data
type SessionLength struct {
	State      string `json:"state"`
	ActiveTime int    `json:"active_time"`
}

//Purchase structure
type Purchase struct {
	Purchases []struct {
		Sku    string `json:"sku"`
		Iso    string `json:"iso"`
		Amount string `json:"amount"`
	} `json:"purchases"`
}

//CSV request data
type CSV struct {
	ExtraFields     []string `json:"extra_fields"`
	LastActiveSince string   `json:"last_active_since"`
}

//Notification used in notification get/view
type Notification struct {
	ID         string            `json:"id"`
	Successful int               `json:"successful"`
	Failed     int               `json:"failed"`
	Converted  int               `json:"converted"`
	Remaining  int               `json:"remaining"`
	QueuedAt   int               `json:"queued_at"`
	SendAfter  int               `json:"send_after"`
	URL        string            `json:"url"`
	Data       map[string]string `json:"data"` //key: value
	Canceled   bool              `json:"canceled"`
	Headings   map[string]string `json:"headings"` //en, es
	Contents   map[string]string `json:"contents"` //en, es
}

//Notifications for viewNotifications
type Notifications struct {
	TotalCount    int `json:"total_count"`
	Offset        int `json:"offset"`
	Limit         int `json:"limit"`
	Notifications []struct {
		ID         string `json:"id"`
		Successful int    `json:"successful"`
		Failed     int    `json:"failed"`
		Converted  int    `json:"converted"`
		Remaining  int    `json:"remaining"`
		QueuedAt   int    `json:"queued_at"`
		SendAfter  int    `json:"send_after"`
		Canceled   bool   `json:"canceled"`
		URL        string `json:"url,omitempty"`
		Headings   struct {
			En string `json:"en"`
			Es string `json:"es"`
		} `json:"headings"`
		Contents struct {
			En string `json:"en"`
			Es string `json:"es"`
		} `json:"contents"`
		Data struct {
			Foo  string `json:"foo"`
			Your string `json:"your"`
		} `json:"data,omitempty"`
	} `json:"notifications"`
}

//Track for trackOpen
type Track struct {
	Opened bool   `json:"opened"`
	AppID  string `json:"app_id"`
}

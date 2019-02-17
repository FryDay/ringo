package ringo

import (
	"strconv"
	"strings"
	"time"
)

// Devices contains lists of different devices.
type Devices struct {
	Doorbots    []*device `json:"doorbots,omitempty"`
	StickupCams []*device `json:"stickup_cams,omitempty"`
}

// DeviceHealth contains the health of a device.
type DeviceHealth struct {
	ID                        int64       `json:"id,omitempty"`
	WifiName                  string      `json:"wifi_name,omitempty"`
	BatteryPercentage         batteryLife `json:"battery_percentage,omitempty"`
	BatteryPercentageCategory string      `json:"battery_percentage_category,omitempty"`
	BatteryVoltage            batteryLife `json:"battery_voltage,omitempty"`
	BatteryVoltageCategory    string      `json:"battery_voltage_category,omitempty"`
	LatestSignalStrength      int64       `json:"latest_signal_strength,omitempty"`
	LatestSignalCategory      string      `json:"latest_signal_category,omitempty"`
	AverageSignalStrength     int64       `json:"average_signal_strength,omitempty"`
	AverageSignalCategory     string      `json:"average_signal_category,omitempty"`
	Firmware                  string      `json:"firmware,omitempty"`
	UpdatedAt                 time.Time   `json:"updated_at,omitempty"`
	WifiIsRingNetwork         bool        `json:"wifi_is_ring_network,omitempty"`
	PacketLossCategory        string      `json:"packet_loss_category,omitempty"`
	PacketLossStrength        int64       `json:"packet_loss_strength,omitempty"`
}

type device struct {
	ID                    int64        `json:"id,omitempty"`
	Description           string       `json:"description,omitempty"`
	DeviceID              string       `json:"device_id,omitempty"`
	TimeZone              string       `json:"time_zone,omitempty"`
	Subscribed            bool         `json:"subscribed,omitempty"`
	SubscribedMotions     bool         `json:"subscribed_motions,omitempty"`
	BatteryLife           batteryLife  `json:"battery_life,omitempty"`
	ExternalConnection    bool         `json:"external_connection,omitempty"`
	FirmwareVersion       string       `json:"firmware_version,omitempty"`
	Kind                  string       `json:"kind,omitempty"`
	Latitude              float64      `json:"latitude,omitempty"`
	Longitude             float64      `json:"longitude,omitempty"`
	Address               string       `json:"address,omitempty"`
	Settings              *settings    `json:"settings,omitempty"`
	Features              *features    `json:"features,omitempty"`
	Owned                 bool         `json:"owned,omitempty"`
	Alerts                *alerts      `json:"alerts,omitempty"`
	Stolen                bool         `json:"stolen,omitempty"`
	LocationID            string       `json:"location_id,omitempty"`
	Owner                 *owner       `json:"owner,omitempty"`
	NightModeStatus       string       `json:"night_mode_status,omitempty"`
	LEDStatus             string       `json:"led_status,omitempty"`
	RingCamLightInstalled string       `json:"ring_cam_light_installed,omitempty"`
	RingCamSetupFlow      string       `json:"ring_cam_setup_flow,omitempty"`
	SirenStatus           *sirenStatus `json:"siren_status,omitempty"`
	// MotionSnooze ?
	// RingID ?
}

type sirenStatus struct {
	SecondsRemaining int64 `json:"seconds_remaining,omitempty"`
}

type owner struct {
	ID        int64  `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type alerts struct {
	Connection string `json:"connection,omitempty"`
}

type features struct {
	MotionsEnabled          bool `json:"motions_enabled,omitempty"`
	ShowRecordings          bool `json:"show_recordings,omitempty"`
	AdvancedMotionEnabled   bool `json:"advanced_motion_enabled,omitempty"`
	PeopleOnlyEnabled       bool `json:"people_only_enabled,omitempty"`
	ShadowCorrectionEnabled bool `json:"shadow_correction_enabled,omitempty"`
	MotionMessageEnabled    bool `json:"motion_message_enabled,omitempty"`
	NightVisionEnabled      bool `json:"night_vision_enabled,omitempty"`
}

type settings struct {
	EnableVod                      bool                   `json:"enable_vod,omitempty"`
	MotionZones                    *motionZones           `json:"motion_zones,omitempty"`
	MotionSnoozePresetProfile      string                 `json:"motion_snooze_preset_profile,omitempty"`
	LiveViewPresetProfile          string                 `json:"live_view_preset_profile,omitempty"`
	LiveViewPresets                []string               `json:"live_view_presets,omitempty"`
	MotionSnoozePresets            []string               `json:"motion_snooze_presets,omitempty"`
	DoorbellVolume                 int64                  `json:"doorbell_volume,omitempty"`
	ChimeSettings                  *chimeSettings         `json:"chime_settings,omitempty"`
	VideoSettings                  *videoSettings         `json:"video_settings,omitempty"`
	MotionAnnouncement             bool                   `json:"motion_announcement,omitempty"`
	StreamSetting                  int64                  `json:"stream_setting,omitempty"`
	AdvancedMotionDetectionEnabled bool                   `json:"advanced_motion_detection_enabled,omitempty"`
	PIRSettings                    *pirSettings           `json:"pir_settings,omitempty"`
	PIRMotionZones                 []int64                `json:"pir_motion_zones,omitempty"`
	FloodlightSettings             *floodlightSettings    `json:"floodlight_settings,omitempty"`
	LightScheduleSettings          *lightScheduleSettings `json:"light_schedule_settings,omitempty"`
}

type floodlightSettings struct {
	Priority int64 `json:"priority,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	AlwaysOn bool  `json:"always_on,omitempty"`
}

type lightScheduleSettings struct {
	StartHour   int64 `json:"start_hour,omitempty"`
	StartMinute int64 `json:"start_minute,omitempty"`
	EndHour     int64 `json:"end_hour,omitempty"`
	EndMinute   int64 `json:"end_minute,omitempty"`
}

type motionZones struct {
	EnableAudio            bool                    `json:"enable_audio,omitempty"`
	ActiveMotionFilter     int64                   `json:"active_motion_filter,omitempty"`
	Sensitivity            int64                   `json:"sensitivity,omitempty"`
	AdvancedObjectSettings *advancedObjectSettings `json:"advanced_object_settings,omitempty"`
	Zone1                  *zone                   `json:"zone1,omitempty"`
	Zone2                  *zone                   `json:"zone2,omitempty"`
	Zone3                  *zone                   `json:"zone3,omitempty"`
	PIRSettings            *pirSettings            `json:"pir_settings,omitempty"`
}

type pirSettings struct {
	Sensitivity1 int64 `json:"sensitivity_1,omitempty"`
	Sensitivity2 int64 `json:"sensitivity_2,omitempty"`
	Sensitivity3 int64 `json:"sensitivity_3,omitempty"`
	ZoneMask     int64 `json:"zone_mask,omitempty"`
}

type advancedObjectSettings struct {
	HumanDetectionConfidence *dayNight `json:"human_detection_confidence,omitempty"`
	MotionZoneOverlap        *dayNight `json:"motion_zone_overlap,omitempty"`
	ObjectTimeOverlap        *dayNight `json:"object_time_overlap,omitempty"`
	ObjectSizeMinimum        *dayNight `json:"object_size_minimum,omitempty"`
	ObjectSizeMaximum        *dayNight `json:"object_size_maximum,omitempty"`
}

type dayNight struct {
	Day   float64 `json:"day,omitempty"`
	Night float64 `json:"night,omitempty"`
}

type zone struct {
	Name    string  `json:"name,omitempty"`
	State   int64   `json:"state,omitempty"`
	Vertex1 *vertex `json:"vertex1,omitempty"`
	Vertex2 *vertex `json:"vertex2,omitempty"`
	Vertex3 *vertex `json:"vertex3,omitempty"`
	Vertex4 *vertex `json:"vertex4,omitempty"`
	Vertex5 *vertex `json:"vertex5,omitempty"`
	Vertex6 *vertex `json:"vertex6,omitempty"`
	Vertex7 *vertex `json:"vertex7,omitempty"`
	Vertex8 *vertex `json:"vertex8,omitempty"`
}

type vertex struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type chimeSettings struct {
	Type     int64 `json:"type,omitempty"`
	Enabled  bool  `json:"enabled,omitempty"`
	Duration int64 `json:"duration,omitempty"`
}

type videoSettings struct {
	Birton     string `json:"birton,omitempty"`
	Brightness int64  `json:"brightness,omitempty"`
	Contrast   int64  `json:"contrast,omitempty"`
	Saturation int64  `json:"saturation,omitempty"`
	AELevel    int64  `json:"ae_level,omitempty"`
}

type batteryLife string

func (bl *batteryLife) UnmarshalJSON(b []byte) error {
	asString, err := strconv.Unquote(string(b))
	if err != nil {
		asString = string(b)
	}
	*bl = batteryLife(strings.ToLower(asString))

	return nil
}

package models

import "time"

type RTCWAKE struct {
	Type string
	Date time.Time
}

var RTCWAKE_TYPE_STANDBY = "standby"
var RTCWAKE_TYPE_MEMORY = "mem"
var RTCWAKE_TYPE_DISK = "disk"
var RTCWAKE_TYPE_OFF = "off"
var RTCWAKE_TYPE_ON = "on"
var RTCWAKE_TYPE_NO = "no"

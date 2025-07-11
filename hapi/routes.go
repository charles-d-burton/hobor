// Package hapi package to hold the home assistnt api spec
package hapi

const (
	API = `/api/`
)

const (
	// GET routes
	GetConfig   = API + `config`
	GetEvents   = API + `events`
	GetServices = API + `services`
	GetHistory  = API + `history/period/%s` // String should be timstamp

	GetLogbook     = API + `logbook/%s` // String should be timstamp
	GetStates      = API + `states`
	GetEntity      = GetStates + `/%s` // String is the entity_id
	GetErrorLog    = API + `error_log`
	GetCameraProxy = API + `camera_proxy/%s` // String is the camera entity_id
	GetCalendars   = API + `calendars`
	// Strings are <calendar_entity_id>, <start_timestampe>, <end_timestamp>
	GetCalendarEntity = API + `calendars/%s?start=%s&end=%s`
)

const (
	// POST Routes
	PostStates      = API + `states/%s`      // String is the entity_id
	PostEvents      = API + `events/%s`      // String is the event_type
	PostServices    = API + `services/%s/%s` // Strings are <domain>/<service>
	PostTemplate    = API + `template`
	PostCheckConfig = API + `config/core/check_config`
	PostHandle      = API + `intent/handle`
)

const (
	// Delete Routes
	DeleteStates = API + `states/%s` // String is the entity_id
)

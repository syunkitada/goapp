package spec

import "time"

type ResourceEvent struct {
	Name            string
	Time            string
	Level           string
	Handler         string
	Msg             string
	ReissueDuration int
	Tag             map[string]string
}

type GetEvents struct {
	Cluster string `validate:"required"`
}

type GetEventsData struct {
	Events []Event
}

type Event struct {
	Check           string
	Handler         string
	Level           string
	Project         string
	Node            string
	Msg             string
	ReissueDuration int
	Time            time.Time
}

type IssueEvent struct {
	Event Event
}

type IssueEventData struct{}

type GetIssuedEvents struct {
}

type GetIssuedEventsData struct {
	Events []Event
}

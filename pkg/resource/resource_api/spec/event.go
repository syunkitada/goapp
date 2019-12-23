package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

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
	Level           string
	Project         string
	Node            string
	Msg             string
	ReissueDuration int
	Ignored         int
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

type CreateEventRules struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateEventRulesData struct{}

type UpdateEventRules struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateEventRulesData struct{}

type DeleteEventRules struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteEventRulesData struct{}

type GetEventRules struct {
	Cluster string `validate:"required"`
}

type EventRule struct {
	Node  string
	Name  string
	Kind  string // Ignore, Aggregation
	Until time.Time
}

type GetEventRulesData struct {
	EventRules []EventRule
}

var EventsTable = index_model.Table{
	Name:    "Events",
	Route:   "/Events",
	Kind:    "Table",
	DataKey: "Events",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Check", IsSearch: true,
		},
		index_model.TableColumn{
			Name: "Node", IsSearch: true,
		},
		index_model.TableColumn{
			Name:           "Level",
			RowColoringMap: map[string]string{"Warning": "Warning", "Critical": "Critical"},
			FilterValues: []map[string]string{
				map[string]string{
					"Icon":  "Warning",
					"Value": "Warning",
				},
				map[string]string{
					"Icon":  "Critical",
					"Value": "Critical",
				},
			},
		},
		index_model.TableColumn{Name: "Msg"},
		index_model.TableColumn{Name: "Time", Kind: "Time"},
	},
}

var EventRulesTable = index_model.Table{
	Name:    "EventRules",
	Route:   "/EventRules",
	Kind:    "Table",
	DataKey: "EventRules",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
		},
		index_model.TableColumn{Name: "Host"},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "Until", Kind: "Time"},
	},
}

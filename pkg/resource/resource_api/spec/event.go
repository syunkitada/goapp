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
	Silenced        int
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
	Cluster string `validate:"required"`
	Specs   string `validate:"required" flagKind:"file"`
}

type CreateEventRulesData struct{}

type UpdateEventRules struct {
	Cluster string `validate:"required"`
	Specs   string `validate:"required" flagKind:"file"`
}

type UpdateEventRulesData struct{}

type DeleteEventRules struct {
	Cluster string `validate:"required"`
	Specs   string `validate:"required" flagKind:"file"`
}

type DeleteEventRulesData struct{}

type GetEventRule struct {
	Cluster string `validate:"required"`
	Name    string
}

type GetEventRuleData struct {
	EventRule EventRule
}

type GetEventRules struct {
	Cluster string `validate:"required"`
}

type EventRule struct {
	Project string
	Node    string
	Name    string
	Msg     string
	Check   string
	Level   string
	Kind    string // Filter, Silence, Aggregate, Action
	Until   *time.Time
	Spec    interface{}
}

type GetEventRulesData struct {
	EventRules []EventRule
}

type EventRuleFilterSpec struct {
}

type EventRuleSilenceSpec struct {
}

type EventRuleAggregateSpec struct {
	AggregateNode  bool
	AggregateCheck bool
	Priority       int
}

type EventRuleActionSpec struct {
	Actions      []string
	Priority     int
	ContinueNext bool
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
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "Node"},
		index_model.TableColumn{Name: "Check"},
		index_model.TableColumn{Name: "Msg"},
		index_model.TableColumn{Name: "Until", Kind: "Time"},
	},
}

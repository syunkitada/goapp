package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type ResourceEvent struct {
	Name            string
	Time            time.Time
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
	Node string
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

var EventsTable = base_index_model.Table{
	Name:        "Events",
	Kind:        "Table",
	DataQueries: []string{"GetEvents"},
	DataKey:     "Events",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Check", IsSearch: true,
		},
		base_index_model.TableColumn{
			Name: "Node", IsSearch: true,
		},
		base_index_model.TableColumn{
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
		base_index_model.TableColumn{Name: "Msg"},
		base_index_model.TableColumn{Name: "Time", Kind: "Time"},
	},
}

var EventRulesTable = base_index_model.Table{
	Name:        "EventRules",
	Kind:        "Table",
	DataQueries: []string{"GetEventRules"},
	DataKey:     "EventRules",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "Node"},
		base_index_model.TableColumn{Name: "Check"},
		base_index_model.TableColumn{Name: "Msg"},
		base_index_model.TableColumn{Name: "Until", Kind: "Time"},
	},
}

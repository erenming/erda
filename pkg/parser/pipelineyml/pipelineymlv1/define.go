// All struct fields are required, unless "Optional" explicitly declared.
package pipelineymlv1

type Pipeline struct {
	Version string   `json:"version"`
	Stages  []*Stage `json:"stages,omitempty"`

	Envs map[string]string `json:"envs,omitempty"`

	Resources     []Resource     `json:"resources,omitempty"`
	ResourceTypes []ResourceType `json:"resource_types,omitempty"`

	Triggers []Trigger `json:"triggers,omitempty"`
}

type Stage struct {
	Name        string       `json:"name,omitempty"`
	TaskConfigs []TaskConfig `json:"tasks,omitempty"`
	Tasks       []StepTask   `json:"-"`
}

type TaskConfig map[string]interface{}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Optional
	Source Source `json:"source"`
}

type ResourceType struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Source Source `json:"source"`
}

type Trigger struct {
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	Cron    string  `json:"cron"`
	Filters Filters `json:"filters,omitempty"`
}

type Version map[string]string

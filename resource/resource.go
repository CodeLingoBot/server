package resource

import (
	"github.com/awethome/server/action"
)

type Resource struct {
	Name string
	Actions map[string]action.Action
}

func (resource *Resource) AddAction(resourceAction action.Action) {
	if len(resource.Actions) == 0 {
		resource.Actions = make(map[string]action.Action)
	}
	resource.Actions[resourceAction.Name] =  resourceAction
}
package job

import (
	"github.com/equinor/radix-common/utils"
	"github.com/equinor/radix-common/utils/slice"
)

type ExpectedSteps interface {
	Add(stepName string, components ...string) ExpectedSteps
	Count() int
	HasStepWithComponent(stepName string, components []string) bool
}

type expectedStep struct {
	components []string
	name       string
}

type expectedSteps struct {
	steps []expectedStep
}

func NewExpectedSteps() ExpectedSteps {
	return &expectedSteps{
		steps: make([]expectedStep, 0),
	}
}

func (es *expectedSteps) Add(stepName string, components ...string) ExpectedSteps {
	es.steps = append(es.steps, expectedStep{name: stepName, components: components})
	return es
}

func (es *expectedSteps) Count() int {
	return len(es.steps)
}

func (es *expectedSteps) HasStepWithComponent(stepName string, components []string) bool {
	return slice.Any(es.steps, func(step expectedStep) bool {
		return step.name == stepName && utils.ArrayEqualElements(step.components, components)
	})
}

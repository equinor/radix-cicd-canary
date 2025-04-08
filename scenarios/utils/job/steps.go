package job

import (
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-common/utils"
	"github.com/equinor/radix-common/utils/slice"
)

type ExpectedSteps interface {
	Add(stepName string, components ...string) ExpectedSteps
	AddForSubPipeline(stepName string, subPipelineTaskStep *models.SubPipelineTaskStep) ExpectedSteps
	Count() int
	HasStepWithComponent(stepName string, components []string) bool
	HasStepWithSubPipelineTaskStep(stepName string, subPipelineTaskStep *models.SubPipelineTaskStep) bool
}

type expectedStep struct {
	components          []string
	name                string
	subPipelineTaskStep *models.SubPipelineTaskStep
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

func (es *expectedSteps) AddForSubPipeline(stepName string, subPipelineTaskStep *models.SubPipelineTaskStep) ExpectedSteps {
	es.steps = append(es.steps, expectedStep{name: stepName, subPipelineTaskStep: subPipelineTaskStep})
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

func (es *expectedSteps) HasStepWithSubPipelineTaskStep(stepName string, subPipelineTaskStep *models.SubPipelineTaskStep) bool {
	if subPipelineTaskStep == nil {
		panic("subPipelineTaskStep is nil")
	}
	return slice.Any(es.steps, func(step expectedStep) bool {
		return step.name == stepName && step.subPipelineTaskStep != nil &&
			step.subPipelineTaskStep.Environment == subPipelineTaskStep.Environment &&
			*step.subPipelineTaskStep.PipelineName == *subPipelineTaskStep.PipelineName
	})
}

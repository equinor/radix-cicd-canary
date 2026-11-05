package job

import (
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-common/utils/slice"
)

type StepType string

const (
	StepTypeCloneConfig     StepType = "clone-config"
	StepTypePipelineRunner  StepType = "pipeline-runner"
	StepTypeClone           StepType = "clone"
	StepTypeBuild           StepType = "build"
	StepTypeSubPipelineStep StepType = "sub-pipeline-step"
	StepTypeUnknown         StepType = "unknown"
)

type StepStatus string

const (
	StepStatusQueued           StepStatus = "Queued"
	StepStatusWaiting          StepStatus = "Waiting"
	StepStatusRunning          StepStatus = "Running"
	StepStatusSucceeded        StepStatus = "Succeeded"
	StepStatusCompleted        StepStatus = "Completed" // Hack: the status for sub-pipelines are taken raw from taskrun.status.steps.terminated
	StepStatusFailed           StepStatus = "Failed"
	StepStatusStopped          StepStatus = "Stopped"
	StepStatusStoppedNoChanges StepStatus = "StoppedNoChanges"
)

type expectedStep struct {
	StepType            StepType
	Status              StepStatus
	Components          []string
	Environments        []string
	SubPipelineTaskStep *models.SubPipelineTaskStep
}

type ExpectedSteps struct {
	steps []expectedStep
}

func NewExpectedSteps() *ExpectedSteps {
	return &ExpectedSteps{
		steps: make([]expectedStep, 0),
	}
}

type Option func(*expectedStep)

func WithStatus(status StepStatus) Option {
	return func(step *expectedStep) {
		step.Status = status
	}
}

func WithComponents(components ...string) Option {
	return func(step *expectedStep) {
		step.Components = components
	}
}

func WithEnvironments(environments ...string) Option {
	return func(step *expectedStep) {
		step.Environments = environments
	}
}

func WithSubPipelineTaskStep(subPipelineTaskStep *models.SubPipelineTaskStep) Option {
	return func(step *expectedStep) {
		step.SubPipelineTaskStep = subPipelineTaskStep
	}
}

func (es *ExpectedSteps) Add(stepType StepType, options ...Option) *ExpectedSteps {
	step := expectedStep{StepType: stepType}

	for _, option := range options {
		option(&step)
	}

	if step.Status == "" {
		step.Status = StepStatusSucceeded
	}

	es.steps = append(es.steps, step)
	return es
}

func (es *ExpectedSteps) Count() int {
	return len(es.steps)
}

func (es *ExpectedSteps) HasStep(step models.Step) bool {

	return slice.Any(es.steps, func(expectedStep expectedStep) bool {
		if step.Type != string(expectedStep.StepType) {
			return false
		}

		if step.Status != string(expectedStep.Status) {
			return false
		}

		if !slice.ElementsMatch(step.Components, expectedStep.Components) {
			return false
		}

		if !slice.ElementsMatch(step.Environments, expectedStep.Environments) {
			return false
		}

		if !compareSubPipelineTaskStep(&step, expectedStep) {
			return false
		}

		return true
	})
}

func compareSubPipelineTaskStep(step *models.Step, expectedStep expectedStep) bool {
	if expectedStep.SubPipelineTaskStep == nil && step.SubPipelineTaskStep == nil {
		return true
	}
	if step.SubPipelineTaskStep == nil || expectedStep.SubPipelineTaskStep == nil {
		return false
	}
	return equalStringPointers(step.SubPipelineTaskStep.Environment, expectedStep.SubPipelineTaskStep.Environment) &&
		equalStringPointers(step.SubPipelineTaskStep.PipelineName, expectedStep.SubPipelineTaskStep.PipelineName)
}

func equalStringPointers(first, second *string) bool {
	if first == nil || second == nil {
		return first == second
	}
	return *first == *second
}

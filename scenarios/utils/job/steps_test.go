package job

import (
	"reflect"
	"testing"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
)

func TestExpectedSteps(t *testing.T) {
	tests := map[string]struct {
		add  func(*ExpectedSteps)
		want []expectedStep
	}{
		"empty": {
			add:  func(*ExpectedSteps) {},
			want: []expectedStep{},
		},
		"defaults status to succeeded": {
			add: func(steps *ExpectedSteps) {
				steps.Add(StepTypeBuild)
			},
			want: []expectedStep{{
				StepType: StepTypeBuild,
				Status:   StepStatusSucceeded,
			}},
		},
		"applies options and adds multiple steps": {
			add: func(steps *ExpectedSteps) {
				steps.Add(StepTypeBuild,
					WithStatus(StepStatusRunning),
					WithComponents("api", "web"),
					WithEnvironments("dev"),
				)
				steps.Add(StepTypeSubPipelineStep, WithSubPipelineTaskStep(&models.SubPipelineTaskStep{}))
			},
			want: []expectedStep{
				{
					StepType:     StepTypeBuild,
					Status:       StepStatusRunning,
					Components:   []string{"api", "web"},
					Environments: []string{"dev"},
				},
				{
					StepType:            StepTypeSubPipelineStep,
					Status:              StepStatusSucceeded,
					SubPipelineTaskStep: &models.SubPipelineTaskStep{},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			steps := NewExpectedSteps()
			if steps == nil {
				t.Fatal("NewExpectedSteps() returned nil")
			}

			test.add(steps)
			if steps.Count() != len(test.want) {
				t.Fatalf("Count() = %d, want %d", steps.Count(), len(test.want))
			}
			if !reflect.DeepEqual(steps.steps, test.want) {
				t.Errorf("steps = %#v, want %#v", steps.steps, test.want)
			}
		})
	}
}

func TestExpectedStepsHasStep(t *testing.T) {
	tests := map[string]struct {
		expected expectedStep
		actual   models.Step
		want     bool
	}{
		"matches type and default status": {
			expected: expectedStep{
				StepType: StepTypeBuild,
				Status:   StepStatusSucceeded,
			},
			actual: models.Step{Type: string(StepTypeBuild), Status: string(StepStatusSucceeded)},
			want:   true,
		},
		"matches component and environment lists independent of order": {
			expected: expectedStep{
				StepType:     StepTypeBuild,
				Status:       StepStatusRunning,
				Components:   []string{"api", "web"},
				Environments: []string{"dev", "prod"},
			},
			actual: models.Step{
				Type:         string(StepTypeBuild),
				Status:       string(StepStatusRunning),
				Components:   []string{"web", "api"},
				Environments: []string{"prod", "dev"},
			},
			want: true,
		},
		"does not match type": {
			expected: expectedStep{
				StepType: StepTypeBuild,
				Status:   StepStatusSucceeded,
			},
			actual: models.Step{Type: string(StepTypeClone), Status: string(StepStatusSucceeded)},
		},
		"does not match status": {
			expected: expectedStep{
				StepType: StepTypeBuild,
				Status:   StepStatusSucceeded,
			},
			actual: models.Step{Type: string(StepTypeBuild), Status: string(StepStatusFailed)},
		},
		"does not match components": {
			expected: expectedStep{
				StepType:   StepTypeBuild,
				Status:     StepStatusSucceeded,
				Components: []string{"api"},
			},
			actual: models.Step{
				Type:       string(StepTypeBuild),
				Status:     string(StepStatusSucceeded),
				Components: []string{"web"},
			},
		},
		"does not match environments": {
			expected: expectedStep{
				StepType:     StepTypeBuild,
				Status:       StepStatusSucceeded,
				Environments: []string{"dev"},
			},
			actual: models.Step{
				Type:         string(StepTypeBuild),
				Status:       string(StepStatusSucceeded),
				Environments: []string{"prod"},
			},
		},
		"matches when both sub-pipeline task steps are nil": {
			expected: expectedStep{
				StepType: StepTypeSubPipelineStep,
				Status:   StepStatusSucceeded,
			},
			actual: models.Step{Type: string(StepTypeSubPipelineStep), Status: string(StepStatusSucceeded)},
			want:   true,
		},
		"does not match when only one sub-pipeline task step is nil": {
			expected: expectedStep{
				StepType:            StepTypeSubPipelineStep,
				Status:              StepStatusSucceeded,
				SubPipelineTaskStep: &models.SubPipelineTaskStep{},
			},
			actual: models.Step{Type: string(StepTypeSubPipelineStep), Status: string(StepStatusSucceeded)},
		},
		"matches equal sub-pipeline fields": {
			expected: expectedStep{
				StepType: StepTypeSubPipelineStep,
				Status:   StepStatusSucceeded,
				SubPipelineTaskStep: &models.SubPipelineTaskStep{
					Environment:  new("dev"),
					PipelineName: new("build"),
				},
			},
			actual: models.Step{
				Type:   string(StepTypeSubPipelineStep),
				Status: string(StepStatusSucceeded),
				SubPipelineTaskStep: &models.SubPipelineTaskStep{
					Environment:  new("dev"),
					PipelineName: new("build"),
				},
			},
			want: true,
		},
		"does not match different sub-pipeline fields": {
			expected: expectedStep{
				StepType: StepTypeSubPipelineStep,
				Status:   StepStatusSucceeded,
				SubPipelineTaskStep: &models.SubPipelineTaskStep{
					Environment:  new("dev"),
					PipelineName: new("build"),
				},
			},
			actual: models.Step{
				Type:   string(StepTypeSubPipelineStep),
				Status: string(StepStatusSucceeded),
				SubPipelineTaskStep: &models.SubPipelineTaskStep{
					Environment:  new("dev"),
					PipelineName: new("test"),
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			steps := NewExpectedSteps().Add(test.expected.StepType,
				WithStatus(test.expected.Status),
				WithComponents(test.expected.Components...),
				WithEnvironments(test.expected.Environments...),
				WithSubPipelineTaskStep(test.expected.SubPipelineTaskStep),
			)
			if got := steps.HasStep(test.actual); got != test.want {
				t.Errorf("HasStep() = %t, want %t", got, test.want)
			}
		})
	}
}

package service

import (
	"go.temporal.io/sdk/workflow"
)

type WorkflowCompensation struct {
	compensations []any
	arguments     [][]any
}

func NewWorkflowCompensation() *WorkflowCompensation {
	return &WorkflowCompensation{
		compensations: []any{},
		arguments:     [][]any{},
	}
}

func (s *WorkflowCompensation) AddCompensation(activity any, parameters ...any) {
	s.compensations = append(s.compensations, activity)
	s.arguments = append(s.arguments, parameters)
}

func (s *WorkflowCompensation) Compensate(ctx workflow.Context, inParallel bool) {
	if !inParallel {
		// Compensate in Last-In-First-Out order, to undo in the reverse order that activies were applied.
		for i := len(s.compensations) - 1; i >= 0; i-- {
			errCompensation := workflow.ExecuteActivity(ctx, s.compensations[i], s.arguments[i]...).Get(ctx, nil)
			if errCompensation != nil {
				workflow.GetLogger(ctx).Error("Executing compensation failed", "Error", errCompensation)
			}
		}
	} else {
		selector := workflow.NewSelector(ctx)
		for i := 0; i < len(s.compensations); i++ {
			execution := workflow.ExecuteActivity(ctx, s.compensations[i], s.arguments[i]...)
			selector.AddFuture(execution, func(f workflow.Future) {
				if errCompensation := f.Get(ctx, nil); errCompensation != nil {
					workflow.GetLogger(ctx).Error("Executing compensation failed", "Error", errCompensation)
				}
			})
		}
		for range s.compensations {
			selector.Select(ctx)
		}
	}
}

func (s *WorkflowCompensation) GetCompensationList() []any {
	return s.compensations
}

func (s *WorkflowCompensation) ClearCompensationList() {
	s.compensations = []any{}
	s.arguments = [][]any{}
}

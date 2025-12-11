package provider

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

// ScheduleInterval defines predefined recurring schedule intervals.
type ScheduleInterval string

const (
	// ScheduleIntervalDaily runs the schedule every 24 hours.
	ScheduleIntervalDaily ScheduleInterval = "daily"
	// ScheduleIntervalWeekly runs the schedule every 7 days.
	ScheduleIntervalWeekly ScheduleInterval = "weekly"
	// ScheduleIntervalMonthly runs the schedule every 30 days.
	ScheduleIntervalMonthly ScheduleInterval = "monthly"
	// ScheduleIntervalYearly runs the schedule every 365 days.
	ScheduleIntervalYearly ScheduleInterval = "yearly"
	// ScheduleIntervalCustom allows defining a custom interval duration.
	ScheduleIntervalCustom ScheduleInterval = "custom"
)

// ScheduleIntervalUnit defines the time unit for custom intervals.
type ScheduleIntervalUnit string

const (
	// ScheduleIntervalUnitSecond represents intervals in seconds.
	ScheduleIntervalUnitSecond ScheduleIntervalUnit = "second"
	// ScheduleIntervalUnitMinute represents intervals in minutes.
	ScheduleIntervalUnitMinute ScheduleIntervalUnit = "minute"
	// ScheduleIntervalUnitHour represents intervals in hours.
	ScheduleIntervalUnitHour ScheduleIntervalUnit = "hour"
	// ScheduleIntervalUnitDay represents intervals in days.
	ScheduleIntervalUnitDay ScheduleIntervalUnit = "day"
	// ScheduleIntervalUnitMonth represents intervals in months.
	ScheduleIntervalUnitMonth ScheduleIntervalUnit = "month"
	// ScheduleIntervalUnitYear represents intervals in years.
	ScheduleIntervalUnitYear ScheduleIntervalUnit = "year"
)

// ScheduleCustomInterval defines a custom recurring interval.
type ScheduleCustomInterval struct {
	// Value is the numeric value of the interval.
	Value int
	// Unit is the time unit for the interval.
	Unit ScheduleIntervalUnit
}

// WorkflowActivity defines the workflow to be executed by a schedule.
type WorkflowActivity struct {
	// WorkflowName is the name of the Temporal workflow to execute.
	WorkflowName string
	// TaskQueue is the optional task queue for the workflow.
	// If nil, uses the current service's task queue.
	TaskQueue *string
	// Args are the arguments to pass to the workflow.
	Args []any
}

// ScheduleOptions defines configuration for creating or updating a schedule.
type ScheduleOptions struct {
	// NextSchedule is the delay before first execution (for one-time schedules).
	// If zero, the schedule is assumed to be recurring only.
	NextSchedule time.Duration
	// Interval defines a recurring schedule interval.
	// If nil, creates a one-time schedule.
	Interval *ScheduleInterval
	// CustomInterval defines a custom recurring interval.
	// Only used when Interval is ScheduleIntervalCustom.
	CustomInterval *ScheduleCustomInterval
	// WorkflowActivity defines the workflow to execute.
	WorkflowActivity *WorkflowActivity
}

// ScheduleClient is the concrete implementation of Scheduler interface.
type ScheduleClient struct {
	ts *service.TemporalService
}

// Scheduler defines the interface for managing Temporal schedules.
// It supports creating one-time and recurring schedules for workflow execution.
type ScheduleProvider interface {
	// Init initializes the scheduler (required for ServiceManager compatibility).
	Init()
	// Close cleans up scheduler resources (required for ServiceManager compatibility).
	Close()
	// CreateSchedule creates a new Temporal schedule.
	CreateSchedule(ctx context.Context, options ScheduleOptions) error
	// UpdateSchedule modifies an existing schedule.
	UpdateSchedule(ctx context.Context, scheduleID string, options ScheduleOptions) error
	// DeleteSchedule removes a schedule.
	DeleteSchedule(ctx context.Context, scheduleID string) error
}

// NewScheduleClient creates a new Scheduler instance.
func NewScheduleClient(ts *service.TemporalService) ScheduleProvider {
	return &ScheduleClient{
		ts: ts,
	}
}

// Init implements Scheduler interface for ServiceManager compatibility.
// ScheduleClient has no initialization logic.
func (s *ScheduleClient) Init() {
	// No initialization needed
}

// Close implements Scheduler interface for ServiceManager compatibility.
// ScheduleClient has no cleanup logic.
func (s *ScheduleClient) Close() {
	// No cleanup needed
}

// CreateSchedule creates a new Temporal schedule.
//
// For one-time schedules:
//   - Set NextSchedule to the delay before execution
//   - Leave Interval as nil
//
// For recurring schedules:
//   - Set Interval to a predefined interval (daily, weekly, etc.)
//   - Or use ScheduleIntervalCustom with CustomInterval
//
// TaskQueue defaults to the current service's queue unless overridden in WorkflowActivity.
func (s *ScheduleClient) CreateSchedule(
	ctx context.Context,
	options ScheduleOptions,
) error {
	if options.WorkflowActivity == nil {
		return errors.New("workflow activity is nil")
	}

	if options.WorkflowActivity.Args == nil {
		return errors.New("workflow activity args is nil")
	}

	if options.NextSchedule == 0 {
		return errors.New("next schedule is zero")
	}

	taskQueue := s.ts.TaskQueue
	if options.WorkflowActivity.TaskQueue != nil {
		taskQueue = *options.WorkflowActivity.TaskQueue
	}

	scheduleSpec, err := s.generateScheduleSpec(options)
	if err != nil {
		return err
	}

	_, err = s.ts.Client.ScheduleClient().Create(ctx, client.ScheduleOptions{
		Spec: *scheduleSpec,
		Action: &client.ScheduleWorkflowAction{
			Workflow:  options.WorkflowActivity.WorkflowName,
			TaskQueue: taskQueue,
			Args:      options.WorkflowActivity.Args,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 2,
			},
		},
	})
	return err
}

// UpdateSchedule modifies an existing Temporal schedule.
//
// Supports partial updates:
//   - If NextSchedule is 0, schedule spec is not updated
//   - If WorkflowActivity is nil, workflow action is not updated
//
// This allows updating only the schedule time or only the workflow parameters.
func (s *ScheduleClient) UpdateSchedule(ctx context.Context, scheduleID string, options ScheduleOptions) error {
	// Generate new schedule spec
	newScheduleSpec, err := s.generateScheduleSpec(options)
	if err != nil {
		return err
	}

	// Get schedule handle
	handle := s.ts.Client.ScheduleClient().GetHandle(ctx, scheduleID)
	_, err = handle.Describe(ctx) // Temporal will throw error if schedule not found
	if err != nil {
		return errors.New("schedule not found")
	}

	// Update the schedule
	err = handle.Update(ctx, client.ScheduleUpdateOptions{
		DoUpdate: func(input client.ScheduleUpdateInput) (*client.ScheduleUpdate, error) {
			// Modify the existing schedule
			schedule := input.Description.Schedule

			// Update the schedule spec (when to run), only if it is not nil
			if newScheduleSpec != nil {
				if len(schedule.Spec.Intervals) == 0 && len(newScheduleSpec.Intervals) > 0 {
					slog.Info("Schedule spec updated from one time to interval", "scheduleID", scheduleID)
				} else if len(schedule.Spec.Intervals) > 0 && len(newScheduleSpec.Intervals) == 0 {
					slog.Info("Schedule spec updated from interval to one time", "scheduleID", scheduleID)
				}

				schedule.Spec = newScheduleSpec
			}

			// Update the action (what to run), only if it is not nil
			if options.WorkflowActivity != nil {
				if options.WorkflowActivity.Args == nil {
					return nil, errors.New("workflow activity args is nil")
				}
				taskQueue := s.ts.TaskQueue
				if options.WorkflowActivity.TaskQueue != nil {
					taskQueue = *options.WorkflowActivity.TaskQueue
				}
				schedule.Action = &client.ScheduleWorkflowAction{
					Workflow:  options.WorkflowActivity.WorkflowName,
					TaskQueue: taskQueue,
					Args:      options.WorkflowActivity.Args,
					RetryPolicy: &temporal.RetryPolicy{
						MaximumAttempts: 2,
					},
				}
			}

			return &client.ScheduleUpdate{
				Schedule: &schedule,
			}, nil
		},
	})

	return err
}

// DeleteSchedule removes a Temporal schedule.
// The schedule will stop executing and be removed from the Temporal server.
func (s *ScheduleClient) DeleteSchedule(ctx context.Context, scheduleID string) error {
	handle := s.ts.Client.ScheduleClient().GetHandle(ctx, scheduleID)
	_, err := handle.Describe(ctx) // Temporal will throw error if schedule not found
	if err != nil {
		return errors.New("schedule not found")
	}
	return handle.Delete(ctx)
}

// generateScheduleSpec creates a Temporal schedule specification from options.
// Returns nil if NextSchedule is 0 (indicating no schedule time update needed).
func (s *ScheduleClient) generateScheduleSpec(options ScheduleOptions) (*client.ScheduleSpec, error) {
	if options.NextSchedule == 0 {
		return nil, nil
	}

	scheduledAt := time.Now().Add(options.NextSchedule)

	scheduleSpec := client.ScheduleSpec{}

	if options.Interval != nil {
		intervalSpec := client.ScheduleIntervalSpec{}
		switch *options.Interval {
		case ScheduleIntervalDaily:
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Hour * 24,
			}
		case ScheduleIntervalWeekly:
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Hour * 24 * 7,
			}
		case ScheduleIntervalMonthly:
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Hour * 24 * 30,
			}
		case ScheduleIntervalYearly:
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Hour * 24 * 365,
			}
		case ScheduleIntervalCustom:
			if options.CustomInterval == nil {
				return &client.ScheduleSpec{}, errors.New("custom interval is nil")
			}
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Duration(options.CustomInterval.Value) * time.Minute,
			}
		}

		scheduleSpec.Intervals = append(scheduleSpec.Intervals, intervalSpec)
	} else {
		scheduleSpec.Calendars = []client.ScheduleCalendarSpec{
			{
				Second:     []client.ScheduleRange{{Start: scheduledAt.Second()}},
				Minute:     []client.ScheduleRange{{Start: scheduledAt.Minute()}},
				Hour:       []client.ScheduleRange{{Start: scheduledAt.Hour()}},
				DayOfMonth: []client.ScheduleRange{{Start: scheduledAt.Day()}},
				Month:      []client.ScheduleRange{{Start: int(scheduledAt.Month())}},
				Year:       []client.ScheduleRange{{Start: scheduledAt.Year()}},
			},
		}
	}

	return &scheduleSpec, nil
}

package provider

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
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

const (
	// ScheduleIntervalUnitSecond represents intervals in seconds.
	ScheduleIntervalUnitSecond time.Duration = time.Second
	// ScheduleIntervalUnitMinute represents intervals in minutes.
	ScheduleIntervalUnitMinute time.Duration = time.Minute
	// ScheduleIntervalUnitHour represents intervals in hours.
	ScheduleIntervalUnitHour time.Duration = time.Hour
	// ScheduleIntervalUnitDay represents intervals in days.
	ScheduleIntervalUnitDay time.Duration = time.Hour * 24
	// ScheduleIntervalUnitMonth represents intervals in months.
	ScheduleIntervalUnitMonth time.Duration = time.Hour * 24 * 30
	// ScheduleIntervalUnitYear represents intervals in years.
	ScheduleIntervalUnitYear time.Duration = time.Hour * 24 * 365
)

// ScheduleCustomInterval defines a custom recurring interval.
type ScheduleCustomInterval struct {
	// Value is the numeric value of the interval.
	Value int
	// Unit is the time unit for the interval.
	Unit time.Duration
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
	//  - For one-time schedules: this must be > 0.
	//  - For recurring schedules: this can be used to delay the first run, if not set or zero, the schedule starts immediately.
	NextSchedule time.Time
	// Interval defines a recurring schedule interval.
	// If empty, creates a one-time schedule.
	Interval ScheduleInterval
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
	// CreateSchedule creates a new Temporal schedule. Fails if schedule already exists.
	CreateSchedule(ctx context.Context, scheduleId string, options ScheduleOptions) (string, error)
	// CreateOrUpdateSchedule creates a new schedule or updates an existing one (upsert).
	CreateOrUpdateSchedule(ctx context.Context, scheduleId string, options ScheduleOptions) (string, error)
	// UpdateSchedule modifies an existing schedule. Fails if schedule doesn't exist.
	UpdateSchedule(ctx context.Context, scheduleId string, options ScheduleOptions) error
	// DeleteSchedule removes a schedule.
	DeleteSchedule(ctx context.Context, scheduleId string) error
}

// NewScheduleClient creates a new Scheduler instance.
func NewScheduleClient(ts *service.TemporalService) ScheduleProvider {
	db.InitSnowflake()
	return &ScheduleClient{
		ts: ts,
	}
}

func (s *ScheduleClient) Init() {
	// No initialization needed
}

func (s *ScheduleClient) Close() {
	// No cleanup needed
}

// CreateSchedule creates a new Temporal schedule.
// Use CreateOrUpdateSchedule if you want upsert behavior.
//
// For one-time schedules:
//   - Set NextSchedule to the delay before execution
//   - Leave Interval empty
//
// For recurring schedules:
//   - Set Interval to a predefined interval (daily, weekly, etc.)
//   - Or use ScheduleIntervalCustom with CustomInterval defined
//
// TaskQueue defaults to the current service's queue unless overridden in WorkflowActivity.
func (s *ScheduleClient) CreateSchedule(
	ctx context.Context,
	scheduleId string,
	options ScheduleOptions,
) (string, error) {
	if scheduleId == "" {
		return "", errors.New("schedule identifier is required")
	}

	if options.WorkflowActivity == nil {
		return "", errors.New("workflow activity is nil")
	}

	if options.WorkflowActivity.WorkflowName == "" {
		return "", errors.New("workflow name is required")
	}

	if options.Interval == "" && options.NextSchedule.IsZero() {
		return "", errors.New("next schedule should not be zero for one-time schedule")
	}

	taskQueue := s.ts.TaskQueue
	if options.WorkflowActivity.TaskQueue != nil {
		taskQueue = *options.WorkflowActivity.TaskQueue
	}

	scheduleSpec, err := s.generateScheduleSpec(options)
	if err != nil {
		return "", err
	}

	h, err := s.ts.Client.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:   scheduleId,
		Spec: *scheduleSpec,
		Action: &client.ScheduleWorkflowAction{
			Workflow:  options.WorkflowActivity.WorkflowName,
			TaskQueue: taskQueue,
			Args:      options.WorkflowActivity.Args,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts:        2,
				NonRetryableErrorTypes: []string{"WORKFLOW_PANIC"},
			},
		},
	})
	if err != nil {
		return "", err
	}

	return h.GetID(), nil
}

// CreateOrUpdateSchedule creates a new schedule or updates an existing one.
// This is the recommended method when you want upsert behavior (e.g., updating user expiry dates).
func (s *ScheduleClient) CreateOrUpdateSchedule(
	ctx context.Context,
	scheduleId string,
	options ScheduleOptions,
) (string, error) {
	// Try to update first
	err := s.UpdateSchedule(ctx, scheduleId, options)
	if err == nil {
		// Successfully updated, return the schedule ID
		return scheduleId, nil
	}

	// Schedule doesn't exist, create new one
	return s.CreateSchedule(ctx, scheduleId, options)
}

// UpdateSchedule modifies an existing Temporal schedule.
//
// Supports partial updates:
//   - If NextSchedule is 0, schedule spec is not updated
//   - If WorkflowActivity is nil, workflow action is not updated
//
// This allows updating only the schedule time or only the workflow parameters.
func (s *ScheduleClient) UpdateSchedule(ctx context.Context, scheduleId string, options ScheduleOptions) error {
	// Generate new schedule spec (may be nil to indicate no schedule change)
	newScheduleSpec, err := s.generateScheduleSpec(options)
	if err != nil {
		return err
	}

	scheduleID := scheduleId

	// Get schedule handle
	handle := s.ts.Client.ScheduleClient().GetHandle(ctx, scheduleID)
	desc, err := handle.Describe(ctx) // Temporal will throw error if schedule not found
	if err != nil {
		return errors.New("schedule not found")
	}

	// Guard: if caller appears to convert an interval schedule to one-time (Interval == empty)
	// but provides no NextSchedule, reject to avoid creating an undeliverable one-time schedule.
	if options.Interval == "" && options.NextSchedule.IsZero() {
		if len(desc.Schedule.Spec.Intervals) > 0 {
			return errors.New("next schedule is required when updating to one-time schedule")
		}
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
				taskQueue := s.ts.TaskQueue
				if options.WorkflowActivity.TaskQueue != nil {
					taskQueue = *options.WorkflowActivity.TaskQueue
				}
				if options.WorkflowActivity.WorkflowName == "" {
					return nil, errors.New("workflow name is required for updating workflow activity")
				}
				schedule.Action = &client.ScheduleWorkflowAction{
					Workflow:  options.WorkflowActivity.WorkflowName,
					TaskQueue: taskQueue,
					Args:      options.WorkflowActivity.Args,
					RetryPolicy: &temporal.RetryPolicy{
						MaximumAttempts:        2,
						NonRetryableErrorTypes: []string{"WORKFLOW_PANIC"},
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
// This operation is idempotent - if the schedule doesn't exist, it returns nil.
func (s *ScheduleClient) DeleteSchedule(ctx context.Context, scheduleId string) error {
	handle := s.ts.Client.ScheduleClient().GetHandle(ctx, scheduleId)
	_, err := handle.Describe(ctx)
	if err != nil {
		// Schedule doesn't exist - idempotent delete, return success
		return nil
	}
	return handle.Delete(ctx)
}

// generateScheduleSpec creates a Temporal schedule specification from options.
// Returns nil if NextSchedule is 0 and schedule is one-time schedule (indicating no schedule time update needed).
func (s *ScheduleClient) generateScheduleSpec(options ScheduleOptions) (*client.ScheduleSpec, error) {
	if options.NextSchedule.IsZero() && options.Interval == "" {
		return nil, nil
	}

	// Temporal uses UTC time for schedules
	scheduledAt := options.NextSchedule.UTC()
	slog.Info("Scheduled at", "scheduledAt", scheduledAt)

	scheduleSpec := client.ScheduleSpec{}

	if options.Interval != "" { // Recurring schedule
		intervalSpec := client.ScheduleIntervalSpec{}
		switch options.Interval {
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
			if options.CustomInterval.Value <= 0 {
				return &client.ScheduleSpec{}, errors.New("custom interval value must be greater than zero")
			}
			intervalSpec = client.ScheduleIntervalSpec{
				Every: time.Duration(options.CustomInterval.Value) * options.CustomInterval.Unit,
			}
		}

		scheduleSpec.Intervals = append(scheduleSpec.Intervals, intervalSpec)
		// For recurring schedules, when NextSchedule is provided, delay the first run.
		// If NextSchedule is zero, the schedule starts immediately.
		if !options.NextSchedule.IsZero() {
			scheduleSpec.StartAt = scheduledAt
		}
	} else { // One-time schedule
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

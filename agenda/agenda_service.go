package agenda

import (
	"context"
	"fmt"
	"time"

	"github.com/betsv2/tsv2-go-utils/service"
	"github.com/go-co-op/gocron/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StatusScheduleWaiting  = "waiting"
	StatusScheduleFinished = "finished"
	StatusScheduleCanceled = "canceled"
)

type Task = map[string]func(data any)

type TaskMetadata struct {
	Metadata any
}

type Schedule struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	RunAt  time.Time          `json:"runAt" bson:"runAt"`
	Data   any                `json:"data" bson:"data"`
	Status string             `json:"status" bson:"status"`
}

type Agenda struct {
	Service   service.Service //Base Service. Use MakeGlobalService function
	Scheduler gocron.Scheduler
	Tasks     Task
	Schedules []Schedule
}

func (u Agenda) Listen() {

	u.Scheduler.Start()
	fmt.Println("Agenda is running")
}

// Get all incompleted schedules from database and run all the schedules.
// This will be called every time the agenda set to ready
func (u Agenda) retrieveAndRunSchedules(ctx context.Context) {
	err := u.retrieveSchedules(ctx)

	if err != nil {
		panic("Something went wrong")
	}

	for _, schedule := range u.Schedules {
		u.createJob(ctx, schedule)
	}
}

// Run your schedule and save it in database
func (u Agenda) RunSchedule(ctx context.Context, taskName string, runAt time.Time, data TaskMetadata) error {
	newSchedule := Schedule{
		Name:   taskName,
		RunAt:  runAt,
		Status: StatusScheduleWaiting,
		Data:   data.Metadata,
	}

	result, err := u.Service.InsertOne(ctx, newSchedule)

	if err != nil {
		return err
	}

	newSchedule.ID = result.(primitive.ObjectID)

	u.createJob(ctx, newSchedule)

	return nil
}

// Define your task that will be called at the scheduled time
func (u Agenda) DefineTask(name string, action func(data any)) {
	u.Tasks[name] = action
}

// Get all waiting schedules from database
func (u *Agenda) retrieveSchedules(ctx context.Context) error {

	var result []Schedule
	err := u.Service.Find(&result, ctx, bson.M{"status": StatusScheduleWaiting})

	if err != nil {
		return err
	}

	u.Schedules = result

	return err
}

// Set the Scheduled Task into finished
func (u Agenda) setScheduleToFinished(ctx context.Context, scheduleId primitive.ObjectID) {
	u.Service.UpdateOne(ctx, bson.M{"_id": scheduleId}, bson.M{"$set": bson.M{"status": StatusScheduleFinished}})
}

// Change the agenda state into ready.
// Call this method when you finished defining all of your task
func (u Agenda) Ready(ctx context.Context) {
	u.retrieveAndRunSchedules(ctx)
}

func (u Agenda) createJob(ctx context.Context, schedule Schedule) {
	job := gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(schedule.RunAt))

	//check if the scheduled time is in the past
	if schedule.RunAt.Compare(time.Now()) == -1 {
		job = gocron.OneTimeJob(gocron.OneTimeJobStartImmediately())
	}

	//run a new job
	_, err := u.Scheduler.NewJob(
		job,
		gocron.NewTask(
			func(metadata TaskMetadata) {
				u.Tasks[schedule.Name](metadata.Metadata)
				u.setScheduleToFinished(ctx, schedule.ID)
			}, TaskMetadata{Metadata: schedule.Data}),
		gocron.WithTags(schedule.ID.String()),
	)

	if err != nil {
		panic(err.Error())
	}
}

// Cancel running schedule.
// Use provided metadata as the identifier
func (u Agenda) CancelSchedule(ctx context.Context, metadata TaskMetadata) error {
	var schedule Schedule

	//find one of the schedules with provided metadata that are still have status waiting
	err := u.Service.FindOne(&schedule, ctx, bson.M{"data": metadata.Metadata, "status": StatusScheduleWaiting})

	if err != nil {
		return err
	}

	//Remove the running schedule by the tag its given, which is the mongoId
	u.Scheduler.RemoveByTags(schedule.ID.String())

	//Update the schedule status to canceled
	_, err = u.Service.UpdateOne(ctx, bson.M{"_id": schedule.ID}, bson.M{"$set": bson.M{"status": StatusScheduleCanceled}})

	return err
}

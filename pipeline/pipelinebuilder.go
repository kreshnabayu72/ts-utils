package pipeline

import (
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PipelineBuilder is a class for building pipelines
type PipelineBuilder struct {
	pipelines mongo.Pipeline
}

// NewPipelineBuilder is a constructor to initialize PipelineBuilder
func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		pipelines: mongo.Pipeline{},
	}
}

// Group adds a $group stage to the pipeline
func (pb *PipelineBuilder) Group(groupData bson.D) *PipelineBuilder {
	if groupData != nil {
		pb.pipelines = append(pb.pipelines, bson.D{{Key: "$group", Value: groupData}})
	}
	return pb
}

// Match adds a $match stage to the pipeline
func (pb *PipelineBuilder) Match(matchData bson.M) *PipelineBuilder {
	if matchData != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$match", Value: matchData},
		})
	}
	return pb
}

// OptionalMatch optionally adds a $match stage to the pipeline
func (pb *PipelineBuilder) OptionalMatch(matchData bson.M, isMatch bool) *PipelineBuilder {
	if matchData != nil && isMatch {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$match", Value: matchData},
		})
	}
	return pb
}

// Set adds a $set stage to the pipeline
func (pb *PipelineBuilder) Set(setData bson.M) *PipelineBuilder {
	if setData != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$set", Value: setData},
		})
	}
	return pb
}

// Addfields adds a $addFields stage to the pipeline
func (pb *PipelineBuilder) Addfields(setData bson.M) *PipelineBuilder {
	if setData != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$addFields", Value: setData},
		})
	}
	return pb
}

// Project adds a $project stage to the pipeline
func (pb *PipelineBuilder) Project(projectData bson.M) *PipelineBuilder {
	if projectData != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$project", Value: projectData},
		})
	}
	return pb
}

// Adds a $unwind stage to the pipeline
func (pb *PipelineBuilder) Unwind(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$unwind", Value: data},
		})
	}
	return pb
}

// limit adds a $limit stage to the pipeline
func (pb *PipelineBuilder) Limit(limit int) *PipelineBuilder {

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$limit", Value: limit},
	})
	return pb
}

// Sort adds a $sort stage to the pipeline
func (pb *PipelineBuilder) Sort(field bson.M) *PipelineBuilder {
	if field == nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$sort", Value: field},
		})
	}
	return pb
}

// Facet adds a $facet stage to the pipeline
func (pb *PipelineBuilder) Facet(facets bson.M) *PipelineBuilder {

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$facet", Value: facets}, // $facet should be a single object with stages as its values
	})
	return pb
}

type PaginationQuery struct {
	Page      int32  `bson:"page" json:"page"`
	Limit     int32  `bson:"limit" json:"limit"`
	SortBy    string `bson:"sortBy" json:"sortBy"`
	SortOrder int32  `bson:"sortOrder" json:"sortOrder"`
}

// Pagination adds a $facet stage to the pipeline, with additional limit & skip stage
func (pb *PipelineBuilder) Pagination(query PaginationQuery) *PipelineBuilder {

	var skip int32 = 0
	var sortBy = "name"
	var sortOrder int32 = 1
	var limit int32 = math.MaxInt32

	if query.SortBy != "" {
		sortBy = query.SortBy
	}

	if query.SortOrder != 0 {
		sortOrder = query.SortOrder
	}

	if query.Page > 0 && query.Limit > 0 {
		skip = (query.Page - 1) * query.Limit
		limit = query.Limit
	}

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$facet", Value: bson.M{
			"data": bson.A{
				bson.D{{Key: "$sort", Value: bson.D{{Key: sortBy, Value: sortOrder}}}}, bson.D{{Key: "$skip", Value: skip}}, bson.D{{Key: "$limit", Value: limit}},
			},
			"totalRecords": bson.A{
				bson.D{{Key: "$count", Value: "total"}},
			},
		}},
	})
	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$set", Value: bson.M{
			"totalRecords": bson.D{{Key: "$first", Value: "$totalRecords.total"}},
		}},
	})
	return pb
}

// Build returns the constructed pipeline
func (pb *PipelineBuilder) Build() mongo.Pipeline {
	return pb.pipelines
}

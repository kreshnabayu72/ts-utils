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

// GraphLookup adds a $graphLookup stage to the pipeline
func (pb *PipelineBuilder) GraphLookup(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$graphLookup", Value: data},
		})
	}
	return pb
}

// Adds a $lookup stage to the pipeline
func (pb *PipelineBuilder) Lookup(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$lookup", Value: data},
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
	if field != nil {
		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$sort", Value: field},
		})
	}
	return pb
}

// Facet adds a $facet stage to the pipeline
func (pb *PipelineBuilder) Facet(facets bson.M) *PipelineBuilder {

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$facet", Value: facets},
	})
	return pb
}

// ReplaceRoot adds a $replaceRoo stage to the pipeline
func (pb *PipelineBuilder) ReplaceRoot(data bson.M) *PipelineBuilder {

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$replaceRoot", Value: data},
	})
	return pb
}

type PaginationQuery struct {
	Page      int32  `bson:"page"      json:"page"      transform:"int"`
	Limit     int32  `bson:"limit"     json:"limit"     transform:"int"`
	SortBy    string `bson:"sortBy"    json:"sortBy"    transform:"string"`
	SortOrder int32  `bson:"sortOrder" json:"sortOrder" transform:"int"`
}
type Sort struct {
	SortBy    string `bson:"sortBy"    json:"sortBy"`
	SortOrder int32  `bson:"sortOrder" json:"sortOrder"`
}

// Pagination adds a $facet stage to the pipeline, with additional limit & skip stage
// Custom sort can be used as a default sort
//
// Priority order:
// 1. query.SortBy
// 2. customSort
// 3. internal SortBy & SortOrder
//
// For some complex sorting logic, set query.SortBy to "" and use customSort
func (pb *PipelineBuilder) Pagination(query PaginationQuery, customSort ...Sort) *PipelineBuilder {
	var skip int32 = 0
	var sortBy = "name"
	var sortOrder int32 = 1
	var limit int32 = math.MaxInt32

	// Override the default values if provided by query
	if query.SortOrder != 0 {
		sortOrder = query.SortOrder
	}
	if query.SortBy != "" {
		sortBy = query.SortBy
	}

	// Use default sort or by query
	formattedSort := bson.D{
		{Key: sortBy, Value: sortOrder},
	}

	// Only use customSort if provided AND query.SortBy is empty
	// Meaning the default (sortby and sortOrder) will get ignored
	if len(customSort) > 0 && query.SortBy == "" {
		formattedSort = bson.D{}

		for _, sort := range customSort {
			formattedSort = append(
				formattedSort,
				bson.E{Key: sort.SortBy, Value: sort.SortOrder},
			)
		}
	}

	if query.Page > 0 && query.Limit > 0 {
		skip = (query.Page - 1) * query.Limit
		limit = query.Limit
	}

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$facet", Value: bson.M{
			"data": bson.A{
				bson.D{{Key: "$sort", Value: formattedSort}},
				bson.D{{Key: "$skip", Value: skip}},
				bson.D{{Key: "$limit", Value: limit}},
			},
			"totalRecords": bson.A{
				bson.D{{Key: "$count", Value: "total"}},
			},
		}},
	})

	pb.pipelines = append(pb.pipelines, bson.D{{Key: "$unwind", Value: "$totalRecords"}})

	pb.pipelines = append(pb.pipelines, bson.D{
		{Key: "$addFields", Value: bson.M{
			"totalRecords": "$totalRecords.total",
		}},
	})
	return pb
}

func (pb *PipelineBuilder) Search(keyword string, searchFields []string) *PipelineBuilder {
	if keyword != "" && len(searchFields) > 0 {
		query := bson.A{}
		for _, field := range searchFields {
			query = append(query, bson.M{field: GenerateSearchCondition(keyword)})
		}

		pb.pipelines = append(pb.pipelines, bson.D{
			{Key: "$match", Value: bson.M{"$or": query}},
		})
	}

	return pb
}

// Build returns the constructed pipeline
func (pb *PipelineBuilder) Build() mongo.Pipeline {
	return pb.pipelines
}

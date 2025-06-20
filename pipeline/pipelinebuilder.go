package pipeline

import (
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PipelineBuilder is a class for building Pipelines
type PipelineBuilder struct {
	Pipelines mongo.Pipeline
}

// NewPipelineBuilder is a constructor to initialize PipelineBuilder
func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{
		Pipelines: mongo.Pipeline{},
	}
}

// Group adds a $group stage to the pipeline
func (pb *PipelineBuilder) Group(groupData bson.D) *PipelineBuilder {
	if groupData != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{{Key: "$group", Value: groupData}})
	}
	return pb
}

// Match adds a $match stage to the pipeline
func (pb *PipelineBuilder) Match(matchData bson.M) *PipelineBuilder {
	if matchData != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$match", Value: matchData},
		})
	}
	return pb
}

// OptionalMatch optionally adds a $match stage to the pipeline
func (pb *PipelineBuilder) OptionalMatch(matchData bson.M, isMatch bool) *PipelineBuilder {
	if matchData != nil && isMatch {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$match", Value: matchData},
		})
	}
	return pb
}

// Set adds a $set stage to the pipeline
func (pb *PipelineBuilder) Set(setData bson.M) *PipelineBuilder {
	if setData != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$set", Value: setData},
		})
	}
	return pb
}

// Addfields adds a $addFields stage to the pipeline
func (pb *PipelineBuilder) Addfields(setData bson.M) *PipelineBuilder {
	if setData != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$addFields", Value: setData},
		})
	}
	return pb
}

// Project adds a $project stage to the pipeline
func (pb *PipelineBuilder) Project(projectData bson.M) *PipelineBuilder {
	if projectData != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$project", Value: projectData},
		})
	}
	return pb
}

// GraphLookup adds a $graphLookup stage to the pipeline
func (pb *PipelineBuilder) GraphLookup(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$graphLookup", Value: data},
		})
	}
	return pb
}

// Adds a $lookup stage to the pipeline
func (pb *PipelineBuilder) Lookup(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$lookup", Value: data},
		})
	}
	return pb
}

// OptionalLookup Adds a $lookup stage to the pipeline if the condition is met
func (pb *PipelineBuilder) OptionalLookup(data bson.M, isMatch bool) *PipelineBuilder {
	if data != nil && isMatch {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$lookup", Value: data},
		})
	}
	return pb
}

type LookupData struct {
	From         string
	LocalField   string
	ForeignField string
	As           string
	Let          bson.M
	Pipeline     mongo.Pipeline
}

// Adds a $lookup stage to the pipeline
func (pb *PipelineBuilder) LookupStr(rawData LookupData) *PipelineBuilder {
	data := bson.M{}

	if rawData.From != "" {
		data["from"] = rawData.From
	}
	if rawData.LocalField != "" {
		data["localField"] = rawData.LocalField
	}
	if rawData.ForeignField != "" {
		data["foreignField"] = rawData.ForeignField
	}
	if rawData.As != "" {
		data["as"] = rawData.As
	}
	if rawData.Let != nil {
		data["let"] = rawData.Let
	}
	if rawData.Pipeline != nil {
		data["pipeline"] = rawData.Pipeline
	}

	pb.Pipelines = append(pb.Pipelines, bson.D{
		{Key: "$lookup", Value: data},
	})

	return pb
}

// Adds a $unwind stage to the pipeline
func (pb *PipelineBuilder) Unwind(data bson.M) *PipelineBuilder {
	if data != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$unwind", Value: data},
		})
	}
	return pb
}

// limit adds a $limit stage to the pipeline
func (pb *PipelineBuilder) Limit(limit int) *PipelineBuilder {

	pb.Pipelines = append(pb.Pipelines, bson.D{
		{Key: "$limit", Value: limit},
	})
	return pb
}

// Skip adds a $skip stage to the pipeline
func (pb *PipelineBuilder) Skip(skip int) *PipelineBuilder {
	pb.Pipelines = append(pb.Pipelines, bson.D{
		{Key: "$skip", Value: skip},
	})

	return pb
}

// Sort adds a $sort stage to the pipeline
func (pb *PipelineBuilder) Sort(field bson.M) *PipelineBuilder {
	if field != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$sort", Value: field},
		})
	}
	return pb
}

// Sort multiple attributes (has different input than sort (bson.D) since bson.M cant handle sort)
func (pb *PipelineBuilder) SortMultiAttribute(field bson.D) *PipelineBuilder {
	if field != nil {
		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$sort", Value: field},
		})
	}
	return pb
}

// Facet adds a $facet stage to the pipeline
func (pb *PipelineBuilder) Facet(facets bson.M) *PipelineBuilder {

	pb.Pipelines = append(pb.Pipelines, bson.D{
		{Key: "$facet", Value: facets},
	})
	return pb
}

// ReplaceRoot adds a $replaceRoo stage to the pipeline
func (pb *PipelineBuilder) ReplaceRoot(data bson.M) *PipelineBuilder {

	pb.Pipelines = append(pb.Pipelines, bson.D{
		{Key: "$replaceRoot", Value: data},
	})
	return pb
}

type PaginationQuery struct {
	Page      int    `bson:"page"      json:"page"      transform:"int"`
	Limit     int    `bson:"limit"     json:"limit"     transform:"int"`
	SortBy    string `bson:"sortBy"    json:"sortBy"    transform:"string"`
	SortOrder int    `bson:"sortOrder" json:"sortOrder" transform:"int"`
	MultiSort []Sort `bson:"multiSort" json:"multiSort"`
}
type Sort struct {
	SortBy    string `bson:"sortBy"    json:"sortBy"`
	SortOrder int    `bson:"sortOrder" json:"sortOrder"`
}

// Pagination adds a $facet stage to the pipeline, with additional limit & skip stage
// Custom sort can be used as a default sort
//
// Priority order:
// 1. query.MultiSort
// 2. query.SortBy
// 3. customSort
// 4. internal SortBy & SortOrder
//
// For some complex sorting logic, set query.SortBy to "" and use customSort
func (pb *PipelineBuilder) Pagination(query PaginationQuery, customSort ...Sort) *PipelineBuilder {
	var skip int = 0
	var sortBy = "name"
	var sortOrder int = 1
	var limit int = math.MaxInt

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

	//MULTISORT WILL BE PRIORITIZED
	if len(query.MultiSort) > 0 {
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

	pb.Pipelines = append(pb.Pipelines, bson.D{
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

	pb.Pipelines = append(pb.Pipelines, bson.D{{Key: "$unwind", Value: "$totalRecords"}})

	pb.Pipelines = append(pb.Pipelines, bson.D{
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

		pb.Pipelines = append(pb.Pipelines, bson.D{
			{Key: "$match", Value: bson.M{"$or": query}},
		})
	}

	return pb
}

// Build returns the constructed pipeline
func (pb *PipelineBuilder) Build() mongo.Pipeline {
	return pb.Pipelines
}

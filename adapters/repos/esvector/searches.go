package esvector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v5/esapi"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

// ClassSearch searches for classes with optional filters without vector scoring
func (r *Repo) ClassSearch(ctx context.Context, kind kind.Kind,
	className string, limit int, filters *filters.LocalFilter) ([]traverser.VectorSearchResult, error) {
	index := classIndexFromClassName(kind, className)
	return r.search(ctx, index, nil, limit, filters)
}

// VectorClassSearch limits the vector search to a specific class (and kind)
func (r *Repo) VectorClassSearch(ctx context.Context, kind kind.Kind,
	className string, vector []float32, limit int,
	filters *filters.LocalFilter) ([]traverser.VectorSearchResult, error) {
	index := classIndexFromClassName(kind, className)
	return r.search(ctx, index, vector, limit, filters)
}

// VectorSearch retrives the closest concepts by vector distance
func (r *Repo) VectorSearch(ctx context.Context, vector []float32,
	limit int, filters *filters.LocalFilter) ([]traverser.VectorSearchResult, error) {
	return r.search(ctx, "*", vector, limit, filters)
}

func (r *Repo) search(ctx context.Context, index string,
	vector []float32, limit int,
	filters *filters.LocalFilter) ([]traverser.VectorSearchResult, error) {
	var buf bytes.Buffer

	query, err := queryFromFilter(filters)
	if err != nil {
		return nil, err
	}

	body := r.buildSearchBody(query, vector, limit)

	err = json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return nil, fmt.Errorf("vector search: encode json: %v", err)
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(index),
		r.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("vector search: %v", err)
	}

	return r.searchResponse(res)
}

func (r *Repo) buildSearchBody(filterQuery map[string]interface{}, vector []float32, limit int) map[string]interface{} {
	var query map[string]interface{}

	if vector == nil {
		query = filterQuery
	} else {
		query = map[string]interface{}{
			"function_score": map[string]interface{}{
				"query":      filterQuery,
				"boost_mode": "replace",
				"functions": []interface{}{
					map[string]interface{}{
						"script_score": map[string]interface{}{
							"script": map[string]interface{}{
								"inline": "binary_vector_score",
								"lang":   "knn",
								"params": map[string]interface{}{
									"cosine": false,
									"field":  keyVector,
									"vector": vector,
								},
							},
						},
					},
				},
			},
		}
	}

	return map[string]interface{}{
		"query": query,
		"size":  limit,
	}
}

type searchResponse struct {
	Hits struct {
		Hits []hit `json:"hits"`
	} `json:"hits"`
}

type hit struct {
	ID     string                 `json:"_id"`
	Source map[string]interface{} `json:"_source"`
	Score  float32                `json:"_score"`
}

func (r *Repo) searchResponse(res *esapi.Response) ([]traverser.VectorSearchResult,
	error) {
	if err := errorResToErr(res, r.logger); err != nil {
		return nil, fmt.Errorf("vector search: %v", err)
	}

	var sr searchResponse

	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&sr)
	if err != nil {
		return nil, fmt.Errorf("vector search: decode json: %v", err)
	}

	return sr.toVectorSearchResult()
}

func (sr searchResponse) toVectorSearchResult() ([]traverser.VectorSearchResult, error) {
	hits := sr.Hits.Hits
	output := make([]traverser.VectorSearchResult, len(hits), len(hits))
	for i, hit := range hits {
		k, err := kind.Parse(hit.Source[keyKind.String()].(string))
		if err != nil {
			return nil, fmt.Errorf("vector search: result %d: %v", i, err)
		}

		vector, err := base64ToVector(hit.Source[keyVector.String()].(string))
		if err != nil {
			return nil, fmt.Errorf("vector search: result %d: %v", i, err)
		}

		schema, err := parseSchema(hit.Source)
		if err != nil {
			return nil, fmt.Errorf("vector search: result %d: %v", i, err)
		}

		output[i] = traverser.VectorSearchResult{
			ClassName: hit.Source[keyClassName.String()].(string),
			ID:        strfmt.UUID(hit.ID),
			Kind:      k,
			Score:     hit.Score,
			Vector:    vector,
			Schema:    schema,
		}
	}

	return output, nil
}

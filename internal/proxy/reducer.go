package proxy

import (
	"context"

	"github.com/milvus-io/milvus-proto/go-api/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/schemapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/planpb"
)

type milvusReducer interface {
	Reduce([]*internalpb.RetrieveResults) (*milvuspb.QueryResults, error)
}

func createMilvusReducer(ctx context.Context, params *queryParams, req *internalpb.RetrieveRequest, schema *schemapb.CollectionSchema, plan *planpb.PlanNode, collectionName string) milvusReducer {
	if plan.GetQuery().GetIsCount() {
		return &cntReducer{}
	} else if req.GetIterationExtensionReduceRate() > 0 {
		params.limit = params.limit * req.GetIterationExtensionReduceRate()
		if params.limit > Params.CommonCfg.TopKLimit.GetAsInt64() {
			params.limit = Params.CommonCfg.TopKLimit.GetAsInt64()
		}
	}
	return newDefaultLimitReducer(ctx, params, req, schema, collectionName)
}

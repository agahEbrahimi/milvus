/*
 * Licensed to the LF AI & Data foundation under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hookutil

var (
	// WARN: Please DO NOT modify all constants.

	OpTypeKey     = "op_type"
	DatabaseKey   = "database"
	UsernameKey   = "username"
	DataSizeKey   = "data_size"
	SuccessCntKey = "success_cnt"
	FailCntKey    = "fail_cnt"
	RelatedCntKey = "related_cnt"
	NodeIDKey     = "id"
	DimensionKey  = "dim"

	OpTypeInsert       = "insert"
	OpTypeDelete       = "delete"
	OpTypeUpsert       = "upsert"
	OpTypeQuery        = "query"
	OpTypeSearch       = "search"
	OpTypeHybridSearch = "hybrid_search"
	OpTypeNodeID       = "node_id"
)

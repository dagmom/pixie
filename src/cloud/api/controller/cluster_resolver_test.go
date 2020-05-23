package controller_test

import (
	"context"
	"testing"
	"time"

	types "github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"

	"pixielabs.ai/pixielabs/src/cloud/api/controller"
	"pixielabs.ai/pixielabs/src/cloud/api/controller/schema"
	"pixielabs.ai/pixielabs/src/cloud/api/controller/testutils"
	"pixielabs.ai/pixielabs/src/cloud/cloudapipb"
	"pixielabs.ai/pixielabs/src/shared/services/authcontext"
	svcutils "pixielabs.ai/pixielabs/src/shared/services/utils"
	"pixielabs.ai/pixielabs/src/utils"
)

func CreateTestContext() context.Context {
	sCtx := authcontext.New()
	sCtx.Claims = svcutils.GenerateJWTForUser("abcdef", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "test@test.com", time.Now())
	return authcontext.NewContext(context.Background(), sCtx)
}

func LoadSchema(gqlEnv controller.GraphQLEnv) *graphql.Schema {
	schemaData := schema.MustLoadSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	qr := &controller.QueryResolver{gqlEnv}
	gqlSchema := graphql.MustParseSchema(schemaData, qr, opts...)
	return gqlSchema
}

func TestCreateCluster(t *testing.T) {
	clusterID := "7ba7b810-9dad-11d1-80b4-00c04fd430c8"

	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	mockVzSvc.EXPECT().CreateCluster(gomock.Any(), &cloudapipb.CreateClusterRequest{}).
		Return(&cloudapipb.CreateClusterResponse{
			ClusterID: utils.ProtoFromUUIDStrOrNil(clusterID),
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				mutation {
					CreateCluster {
						id
					}
				}
			`,
			ExpectedResult: `
				{
					"CreateCluster": {
						"id":"7ba7b810-9dad-11d1-80b4-00c04fd430c8"
					}
				}
			`,
		},
	})
}

func TestClusterInfoWithoutID(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	clusterInfo := &cloudapipb.ClusterInfo{
		ID:              utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Status:          cloudapipb.CS_HEALTHY,
		LastHeartbeatNs: 4 * 1000 * 1000,
		Config: &cloudapipb.VizierConfig{
			PassthroughEnabled: false,
		},
		VizierVersion:  "vzVersion",
		ClusterVersion: "clusterVersion",
		ClusterName:    "clusterName",
		ClusterUID:     "clusterUID",
	}

	mockVzSvc.EXPECT().GetClusterInfo(gomock.Any(), &cloudapipb.GetClusterInfoRequest{}).
		Return(&cloudapipb.GetClusterInfoResponse{
			Clusters: []*cloudapipb.ClusterInfo{clusterInfo},
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					cluster {
						id
						status
						lastHeartbeatMs
						vizierConfig {
							passthroughEnabled
						}
						vizierVersion
						clusterVersion
						clusterName
						clusterUID
					}
				}
			`,
			ExpectedResult: `
				{
					"cluster": {
						"id":"7ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"status": "CS_HEALTHY",
						"lastHeartbeatMs": 4,
						"vizierConfig": {
							"passthroughEnabled": false
						},
						"vizierVersion": "vzVersion",
						"clusterVersion": "clusterVersion",
						"clusterName": "clusterName",
						"clusterUID": "clusterUID"
					}
				}
			`,
		},
	})
}

func TestClusterInfo(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	clusterInfo := &cloudapipb.ClusterInfo{
		ID:              utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Status:          cloudapipb.CS_HEALTHY,
		LastHeartbeatNs: 4 * 1000 * 1000,
		Config: &cloudapipb.VizierConfig{
			PassthroughEnabled: false,
		},
		VizierVersion:  "vzVersion",
		ClusterVersion: "clusterVersion",
		ClusterName:    "clusterName",
		ClusterUID:     "clusterUID",
	}

	mockVzSvc.EXPECT().GetClusterInfo(gomock.Any(), &cloudapipb.GetClusterInfoRequest{
		ID: clusterInfo.ID,
	}).
		Return(&cloudapipb.GetClusterInfoResponse{
			Clusters: []*cloudapipb.ClusterInfo{clusterInfo},
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					cluster(id: "7ba7b810-9dad-11d1-80b4-00c04fd430c8") {
						id
						status
						lastHeartbeatMs
						vizierConfig {
							passthroughEnabled
						}
						vizierVersion
						clusterVersion
						clusterName
						clusterUID
					}
				}
			`,
			ExpectedResult: `
				{
					"cluster": {
						"id":"7ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"status": "CS_HEALTHY",
						"lastHeartbeatMs": 4,
						"vizierConfig": {
							"passthroughEnabled": false
						},
						"vizierVersion": "vzVersion",
						"clusterVersion": "clusterVersion",
						"clusterName": "clusterName",
						"clusterUID": "clusterUID"
					}
				}
			`,
		},
	})
}

func TestClusterConnectionInfoWithoutID(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	clusterID := utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8")

	clusterInfo := &cloudapipb.ClusterInfo{
		ID:              clusterID,
		Status:          cloudapipb.CS_HEALTHY,
		LastHeartbeatNs: 4 * 1000 * 1000,
		Config: &cloudapipb.VizierConfig{
			PassthroughEnabled: false,
		},
	}

	mockVzSvc.EXPECT().GetClusterInfo(gomock.Any(), &cloudapipb.GetClusterInfoRequest{}).
		Return(&cloudapipb.GetClusterInfoResponse{
			Clusters: []*cloudapipb.ClusterInfo{clusterInfo},
		}, nil)

	mockVzSvc.EXPECT().GetClusterConnectionInfo(gomock.Any(), &cloudapipb.GetClusterConnectionInfoRequest{
		ID: clusterID,
	}).
		Return(&cloudapipb.GetClusterConnectionInfoResponse{
			IPAddress: "127.0.0.1",
			Token:     "this-is-a-token",
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					clusterConnection {
						ipAddress
						token
					}
				}
			`,
			ExpectedResult: `
				{
					"clusterConnection": {
						"ipAddress": "127.0.0.1",
						"token": "this-is-a-token"
					}
				}
			`,
		},
	})
}

func TestClusterConnectionInfo(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	clusterID := utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8")

	mockVzSvc.EXPECT().GetClusterConnectionInfo(gomock.Any(), &cloudapipb.GetClusterConnectionInfoRequest{
		ID: clusterID,
	}).
		Return(&cloudapipb.GetClusterConnectionInfoResponse{
			IPAddress: "127.0.0.1",
			Token:     "this-is-a-token",
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					clusterConnection(id: "7ba7b810-9dad-11d1-80b4-00c04fd430c8") {
						ipAddress
						token
					}
				}
			`,
			ExpectedResult: `
				{
					"clusterConnection": {
						"ipAddress": "127.0.0.1",
						"token": "this-is-a-token"
					}
				}
			`,
		},
	})
}

func TestUpdateClusterVizierConfig(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	mockVzSvc.EXPECT().UpdateClusterVizierConfig(gomock.Any(), &cloudapipb.UpdateClusterVizierConfigRequest{
		ID: utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		ConfigUpdate: &cloudapipb.VizierConfigUpdate{
			PassthroughEnabled: &types.BoolValue{Value: true},
		},
	}).
		Return(&cloudapipb.UpdateClusterVizierConfigResponse{}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				mutation {
					UpdateVizierConfig(clusterID: "7ba7b810-9dad-11d1-80b4-00c04fd430c8", passthroughEnabled: true)
				}
			`,
			ExpectedResult: `
				{
					"UpdateVizierConfig": true
				}
			`,
		},
	})
}

func TestUpdateClusterVizierConfigNoUpdates(t *testing.T) {
	gqlEnv, _, mockVzSvc, _, _, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	mockVzSvc.EXPECT().UpdateClusterVizierConfig(gomock.Any(), &cloudapipb.UpdateClusterVizierConfigRequest{
		ID:           utils.ProtoFromUUIDStrOrNil("7ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		ConfigUpdate: &cloudapipb.VizierConfigUpdate{},
	}).
		Return(&cloudapipb.UpdateClusterVizierConfigResponse{}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				mutation {
					UpdateVizierConfig(clusterID: "7ba7b810-9dad-11d1-80b4-00c04fd430c8")
				}
			`,
			ExpectedResult: `
				{
					"UpdateVizierConfig": true
				}
			`,
		},
	})
}

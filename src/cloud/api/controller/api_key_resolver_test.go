package controller_test

import (
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go/gqltesting"

	"px.dev/pixie/src/cloud/api/controller/testutils"
	"px.dev/pixie/src/cloud/cloudapipb"
	"px.dev/pixie/src/utils"
)

func TestAPIKey(t *testing.T) {
	keyID := "7ba7b810-9dad-11d1-80b4-00c04fd430c8"

	gqlEnv, mockClients, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	createTime := time.Date(2020, 03, 9, 17, 46, 100, 1232409, time.UTC)
	createTimePb, err := types.TimestampProto(createTime)
	if err != nil {
		t.Fatalf("could not write time %+v as protobuf", createTime)
	}

	mockClients.MockAPIKey.EXPECT().
		Get(gomock.Any(), &cloudapipb.GetAPIKeyRequest{
			ID: utils.ProtoFromUUIDStrOrNil(keyID),
		}).
		Return(&cloudapipb.GetAPIKeyResponse{
			Key: &cloudapipb.APIKey{
				ID:        utils.ProtoFromUUIDStrOrNil(keyID),
				Key:       "foobar",
				CreatedAt: createTimePb,
				Desc:      "key description",
			},
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					apiKey(id: "7ba7b810-9dad-11d1-80b4-00c04fd430c8") {
						id
						key
						createdAtMs
						desc
					}
				}
			`,
			ExpectedResult: `
				{
					"apiKey": {
						"id": "7ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"key": "foobar",
						"createdAtMs": 1583776060001.2324,
						"desc": "key description"
					}
				}
			`,
		},
	})
}

func TestAPIKeys(t *testing.T) {
	key1ID := "7ba7b810-9dad-11d1-80b4-00c04fd430c8"
	key2ID := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	key3ID := "8cb848c6-9dad-11d1-80b4-00c04fd430c8"

	gqlEnv, mockClients, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	createTime1 := time.Date(2020, 03, 9, 17, 46, 100, 1232409, time.UTC)
	createTime1Pb, err := types.TimestampProto(createTime1)
	if err != nil {
		t.Fatalf("could not write time %+v as protobuf", createTime1)
	}
	createTime2 := time.Date(2019, 11, 3, 17, 46, 100, 412401, time.UTC)
	createTime2Pb, err := types.TimestampProto(createTime2)
	if err != nil {
		t.Fatalf("could not write time %+v as protobuf", createTime2)
	}
	createTime3 := time.Date(2020, 10, 3, 17, 46, 100, 412401, time.UTC)
	createTime3Pb, err := types.TimestampProto(createTime3)
	if err != nil {
		t.Fatalf("could not write time %+v as protobuf", createTime3)
	}

	// Inserted keys are not sorted by creation time.
	mockClients.MockAPIKey.EXPECT().
		List(gomock.Any(), &cloudapipb.ListAPIKeyRequest{}).
		Return(&cloudapipb.ListAPIKeyResponse{
			Keys: []*cloudapipb.APIKey{
				{
					ID:        utils.ProtoFromUUIDStrOrNil(key1ID),
					Key:       "abc",
					CreatedAt: createTime1Pb,
					Desc:      "key description 1",
				},
				{
					ID:        utils.ProtoFromUUIDStrOrNil(key2ID),
					Key:       "def",
					CreatedAt: createTime2Pb,
					Desc:      "key description 2",
				},
				{
					ID:        utils.ProtoFromUUIDStrOrNil(key3ID),
					Key:       "ghi",
					CreatedAt: createTime3Pb,
					Desc:      "key description 3",
				},
			},
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	// Expect returned keys to be sorted.
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				query {
					apiKeys {
						id
						key
						createdAtMs
						desc
					}
				}
			`,
			ExpectedResult: `
				{
					"apiKeys": [{
						"id": "8cb848c6-9dad-11d1-80b4-00c04fd430c8",
						"key": "ghi",
						"createdAtMs": 1601747260000.4124,
						"desc": "key description 3"
					}, {
						"id": "7ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"key": "abc",
						"createdAtMs": 1583776060001.2324,
						"desc": "key description 1"
					},
					{
						"id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"key": "def",
						"createdAtMs": 1572803260000.4124,
						"desc": "key description 2"
					}]
				}
			`,
		},
	})
}

func TestCreateAPIKey(t *testing.T) {
	keyID := "7ba7b810-9dad-11d1-80b4-00c04fd430c8"

	gqlEnv, mockClients, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	createTime := time.Date(2020, 03, 9, 17, 46, 100, 1232409, time.UTC)
	createTimePb, err := types.TimestampProto(createTime)
	if err != nil {
		t.Fatalf("could not write time %+v as protobuf", createTime)
	}

	mockClients.MockAPIKey.EXPECT().
		Create(gomock.Any(), &cloudapipb.CreateAPIKeyRequest{}).
		Return(&cloudapipb.APIKey{
			ID:        utils.ProtoFromUUIDStrOrNil(keyID),
			Key:       "foobar",
			CreatedAt: createTimePb,
			Desc:      "key description",
		}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				mutation {
					CreateAPIKey {
						id
						key
						createdAtMs
						desc
					}
				}
			`,
			ExpectedResult: `
				{
					"CreateAPIKey": {
						"id": "7ba7b810-9dad-11d1-80b4-00c04fd430c8",
						"key": "foobar",
						"createdAtMs": 1583776060001.2324,
						"desc": "key description"				
					}
				}
			`,
		},
	})
}

func TestDeleteAPIKey(t *testing.T) {
	keyID := "7ba7b810-9dad-11d1-80b4-00c04fd430c8"

	gqlEnv, mockClients, cleanup := testutils.CreateTestGraphQLEnv(t)
	defer cleanup()
	ctx := CreateTestContext()

	mockClients.MockAPIKey.EXPECT().
		Delete(gomock.Any(), utils.ProtoFromUUIDStrOrNil(keyID)).
		Return(&types.Empty{}, nil)

	gqlSchema := LoadSchema(gqlEnv)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema:  gqlSchema,
			Context: ctx,
			Query: `
				mutation {
					DeleteAPIKey(id: "7ba7b810-9dad-11d1-80b4-00c04fd430c8")
				}
			`,
			ExpectedResult: `
				{
					"DeleteAPImentKey": true
				}
			`,
		},
	})
}

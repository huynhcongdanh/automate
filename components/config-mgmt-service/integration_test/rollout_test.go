package integration_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/chef/automate/api/external/cfgmgmt/request"
	"github.com/chef/automate/api/external/cfgmgmt/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRolloutCreate(t *testing.T) {
	ctx := context.Background()

	cleanup := func(t *testing.T) {
		err := cfgmgmt.ClearPg()
		require.NoError(t, err)
	}

	t.Run("create fails when the request is blank", func(t *testing.T) {
		cleanup(t)
		req := request.CreateRollout{}
		_, err := cfgmgmt.CreateRollout(ctx, &req)
		require.Error(t, err)
	})

	reqWithAllRequiredFields := request.CreateRollout{
		PolicyName:       "example-policy-name",
		PolicyNodeGroup:  "example-policy-node-group",
		PolicyRevisionId: "abc123",
		PolicyDomainUrl:  "https://chef-server.example/organizations/example_org",
	}

	t.Run("create fails when the request has a blank PolicyName", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		req.PolicyName = ""
		_, err := cfgmgmt.CreateRollout(ctx, &req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "required field(s) \"policy_name\" were blank")
	})

	t.Run("create fails when the request has a blank PolicyNodeGroup", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		req.PolicyNodeGroup = ""
		_, err := cfgmgmt.CreateRollout(ctx, &req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "required field(s) \"policy_node_group\" were blank")
	})

	t.Run("create fails when the request has a blank PolicyRevisionId", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		req.PolicyRevisionId = ""
		_, err := cfgmgmt.CreateRollout(ctx, &req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "required field(s) \"policy_revision_id\" were blank")
	})

	t.Run("create fails when the request has a blank PolicyDomainUrl", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		req.PolicyDomainUrl = ""
		_, err := cfgmgmt.CreateRollout(ctx, &req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "required field(s) \"policy_domain_url\" were blank")
	})

	t.Run("create succeeds with a minimum valid request", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		created, err := cfgmgmt.CreateRollout(ctx, &req)
		require.NoError(t, err)
		assert.Equal(t, "example-policy-name", created.PolicyName)
		assert.Equal(t, "example-policy-node-group", created.PolicyNodeGroup)
		assert.Equal(t, "abc123", created.PolicyRevisionId)
		assert.Equal(t, "https://chef-server.example/organizations/example_org", created.PolicyDomainUrl)
		s, err := time.Parse(time.RFC3339, created.StartTime)
		require.NoError(t, err)
		assert.WithinDuration(t, time.Now().UTC(), s, 1*time.Minute)
		assert.Empty(t, created.EndTime)
	})

	t.Run("after creating a rollout, fetch by id succeeds", func(t *testing.T) {
		cleanup(t)
		req := reqWithAllRequiredFields
		created, err := cfgmgmt.CreateRollout(ctx, &req)
		require.NoError(t, err)
		fetched, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: created.Id})
		assert.Equal(t, "example-policy-name", fetched.PolicyName)
		assert.Equal(t, "example-policy-node-group", fetched.PolicyNodeGroup)
		assert.Equal(t, "abc123", fetched.PolicyRevisionId)
		assert.Equal(t, "https://chef-server.example/organizations/example_org", fetched.PolicyDomainUrl)
		s, err := time.Parse(time.RFC3339, fetched.StartTime)
		require.NoError(t, err)
		assert.WithinDuration(t, time.Now().UTC(), s, 1*time.Minute)
		assert.Empty(t, fetched.EndTime)
	})

	t.Run("create succeeds with all available fields filled", func(t *testing.T) {
		cleanup(t)
		req := request.CreateRollout{
			PolicyName:           "example-policy-name-full",
			PolicyNodeGroup:      "example-policy-node-group",
			PolicyRevisionId:     "abc123",
			PolicyDomainUrl:      "https://chef-server.example/organizations/example_org",
			PolicyDomainUsername: "bobo",
			ScmType:              request.SCMType_GIT,
			ScmWebType:           request.SCMWebType_GITHUB,
			ScmAuthorName:        "Bobo Tiberius Clown",
			ScmAuthorEmail:       "bobo@example.com",
			PolicyScmUrl:         "git@github.com:chef/automate.git",
			PolicyScmWebUrl:      "https://github.com/chef/automate",
			PolicyScmCommit:      "a2a344e6804629de85ffa50e84caad18ac42cf50",
			Description:          "install winamp",
			CiJobId:              "buildkite/chef-automate-master-verify#11875",
			CiJobUrl:             "https://buildkite.com/chef-oss/chef-automate-master-verify/builds/11875",
		}
		createResponse, err := cfgmgmt.CreateRollout(ctx, &req)
		require.NoError(t, err)
		assert.NotEmpty(t, createResponse.Id)
		_, err = strconv.Atoi(createResponse.Id)
		assert.NoError(t, err)
		checkFields := func(res *response.Rollout) {
			assert.Equal(t, "example-policy-name-full", res.PolicyName)
			assert.Equal(t, "example-policy-node-group", res.PolicyNodeGroup)
			assert.Equal(t, "abc123", res.PolicyRevisionId)
			assert.Equal(t, "https://chef-server.example/organizations/example_org", res.PolicyDomainUrl)
			assert.Equal(t, "bobo", res.PolicyDomainUsername)
			assert.Equal(t, response.SCMType_GIT, res.ScmType)
			assert.Equal(t, response.SCMWebType_GITHUB, res.ScmWebType)
			assert.Equal(t, "git@github.com:chef/automate.git", res.PolicyScmUrl)
			assert.Equal(t, "https://github.com/chef/automate", res.PolicyScmWebUrl)
			assert.Equal(t, "a2a344e6804629de85ffa50e84caad18ac42cf50", res.PolicyScmCommit)
			assert.Equal(t, "Bobo Tiberius Clown", res.ScmAuthorName)
			assert.Equal(t, "bobo@example.com", res.ScmAuthorEmail)
			assert.Equal(t, "install winamp", res.Description)
			assert.Equal(t, "buildkite/chef-automate-master-verify#11875", res.CiJobId)
			assert.Equal(t, "https://buildkite.com/chef-oss/chef-automate-master-verify/builds/11875", res.CiJobUrl)
			assert.NotEmpty(t, res.StartTime)
			s, err := time.Parse(time.RFC3339, res.StartTime)
			require.NoError(t, err)
			assert.WithinDuration(t, time.Now().UTC(), s, 1*time.Minute)
			assert.Empty(t, res.EndTime)
		}
		checkFields(createResponse)
		fetchedById, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: createResponse.Id})
		require.NoError(t, err)
		checkFields(fetchedById)
	})

	t.Run("creating a second rollout with a different target node segment doesn't set end_time", func(t *testing.T) {
		cleanup(t)
		req1 := reqWithAllRequiredFields
		req2 := reqWithAllRequiredFields
		req2.PolicyName = "example-policy-two"
		created1, err := cfgmgmt.CreateRollout(ctx, &req1)
		require.NoError(t, err)
		created2, err := cfgmgmt.CreateRollout(ctx, &req2)
		require.NoError(t, err)

		fetched1, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: created1.Id})
		require.NoError(t, err)
		fetched2, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: created2.Id})
		require.NoError(t, err)

		assert.Empty(t, fetched1.EndTime)
		assert.Empty(t, fetched2.EndTime)

		list, err := cfgmgmt.GetRollouts(ctx, &request.Rollouts{})
		require.NoError(t, err)
		assert.Len(t, list.Rollouts, 2)
	})
	t.Run("creating a second rollout with the same target node segment sets end_time on the older one", func(t *testing.T) {
		cleanup(t)
		req1 := reqWithAllRequiredFields
		req2 := reqWithAllRequiredFields
		req2.PolicyRevisionId = "def456"
		created1, err := cfgmgmt.CreateRollout(ctx, &req1)
		require.NoError(t, err)
		created2, err := cfgmgmt.CreateRollout(ctx, &req2)
		require.NoError(t, err)

		fetched1, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: created1.Id})
		require.NoError(t, err)
		fetched2, err := cfgmgmt.GetRolloutById(ctx, &request.RolloutById{RolloutId: created2.Id})
		require.NoError(t, err)

		assert.NotEmpty(t, fetched1.EndTime, "EndTime was not set on the rollout when it was replaced with a newer one")
		assert.Empty(t, fetched2.EndTime)

		list, err := cfgmgmt.GetRollouts(ctx, &request.Rollouts{})
		require.NoError(t, err)
		assert.Len(t, list.Rollouts, 2)
	})
	t.Run("creating a rollout with the same target node segment and policy revision fails", func(t *testing.T) {
		cleanup(t)
		req1 := reqWithAllRequiredFields
		req2 := reqWithAllRequiredFields
		_, err := cfgmgmt.CreateRollout(ctx, &req1)
		require.NoError(t, err)
		_, err = cfgmgmt.CreateRollout(ctx, &req2)
		require.Error(t, err)

		list, err := cfgmgmt.GetRollouts(ctx, &request.Rollouts{})
		require.NoError(t, err)
		assert.Len(t, list.Rollouts, 1)
	})
	t.Run("rolling back to an old policy revision works", func(t *testing.T) {
		cleanup(t)
		req1 := reqWithAllRequiredFields
		req2 := reqWithAllRequiredFields
		req2.PolicyRevisionId = "def456"
		req3 := reqWithAllRequiredFields

		_, err := cfgmgmt.CreateRollout(ctx, &req1)
		require.NoError(t, err)
		_, err = cfgmgmt.CreateRollout(ctx, &req2)
		require.NoError(t, err)
		_, err = cfgmgmt.CreateRollout(ctx, &req3)
		require.NoError(t, err)

		list, err := cfgmgmt.GetRollouts(ctx, &request.Rollouts{})
		require.NoError(t, err)
		assert.Len(t, list.Rollouts, 3)
	})
}

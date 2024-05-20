// +acceptance

package test

import (
	"context"
	"testing"

	rocket "github.com/karokojnr/rocketbin-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("add a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "afdba19d-9153-4246-9c22-9aaf05b736ea",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "afdba19d-9153-4246-9c22-9aaf05b736ea", resp.Rocket.Id)
	})

	s.T().Run("validates the uuid in the new rocket is a uuid", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "not-a-valid-uuid",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.Error(s.T(), err)
		st := status.Convert(err)
		assert.Equal(s.T(), codes.InvalidArgument, st.Code())
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

// +acceptance

package test

import (
	"context"
	"testing"

	rocket "github.com/karokojnr/rocketbin-protos/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
					Id:   "9",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "9", resp.Rocket.Id)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"testing"
)

type RootCmdSuite struct {
	suite.Suite
}

func (s *RootCmdSuite) SetupTest() {
	// todo: add setup
}

func (s *RootCmdSuite) TestCreateCmd() {
	cmd := newRootCmd()
	s.Require().NoError(cmd.Execute())

	s.Run("config-file-created", func() {
		home, err := homedir.Dir()
		s.Require().NoError(err)
		configFile := path.Join(home, ".wnab.yaml")
		_, err = os.Stat(configFile)
		s.Require().NoError(err, "config file does not exists")
	})
}

func (s *RootCmdSuite) TearDownTest() {
	home, err := homedir.Dir()
	s.Require().NoError(err)
	configFile := path.Join(home, ".wnab.yaml")
	s.Require().NoError(os.Remove(configFile), "failed to delete test generated config file")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RootCmdSuite))
}

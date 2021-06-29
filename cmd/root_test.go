package cmd

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/suite"
	"go.bmvs.io/ynab"
)

type RootCmdSuite struct {
	suite.Suite
}

func (s *RootCmdSuite) SetupTest() {
	// todo: add setup
}

func (s *RootCmdSuite) TestInitCreateCmd() {
	var client ynab.ClientServicer
	cmd := newRootCmd(client)
	s.Require().NoError(cmd.Execute())

	home, err := homedir.Dir()
	s.Require().NoError(err)
	configFile := path.Join(home, ".wnab.yaml")
	_, err = os.Stat(configFile)
	s.Require().NoError(err, "config file does not exists")
	s.Require().NoError(os.Remove(configFile), "failed to delete test generated config file")
}

func (s *RootCmdSuite) TestCreateCmd() {
	fmt.Println("")
}

func (s *RootCmdSuite) TearDownTest() {
	//
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(RootCmdSuite))
}

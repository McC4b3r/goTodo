package test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/testing/protocmp"
)

type Go_appSuite struct {
	suite.Suite
}

func TestGo_appSuite(t *testing.T) {
	suite.Run(t, new(Go_appSuite))
}

func (s *Go_appSuite) SetupSuite() {
	initializeApiClient(5 * time.Second)
	loadTestTodoData(s.T())
}

func assertProtoEqualitySortById(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	theOpts := []cmp.Option{
		cmpopts.SortSlices(func(x, y *go_appv1.Todo) bool {
			return *x.Id < *y.Id
		}),
		protocmp.Transform(),
		protocmp.SortRepeated(func(x, y *go_appv1.Todo) bool {
			return *x.Id < *y.Id
		}),
	}
	theOpts = append(theOpts, opts...)
	diff := cmp.Diff(expected, actual, theOpts...)
	require.Empty(t, diff)
}

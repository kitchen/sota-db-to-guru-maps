package sotadb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SummitDateTestSuite struct {
	suite.Suite
}

func (suite *SummitDateTestSuite) TestMarshalUnmarshal() {
	date := SummitDate{}
	err := date.UnmarshalCSV("31/01/2020")
	suite.NoError(err)
	suite.Equal(time.January, date.Month())
	suite.Equal(31, date.Day())
	suite.Equal(2020, date.Year())

	marshalled, err := date.MarshalCSV()
	suite.NoError(err)
	suite.Equal("31/01/2020", marshalled)
}

func (suite *SummitDateTestSuite) TestEmptyDate() {
	date := SummitDate{}
	err := date.UnmarshalCSV("")
	suite.NoError(err)
	suite.True(date.IsZero())

	csv, err := date.MarshalCSV()
	suite.NoError(err)
	suite.Equal("", csv)
}

func TestSummitDateTestSuite(t *testing.T) {
	suite.Run(t, new(SummitDateTestSuite))
}

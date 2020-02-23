package sotadb

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CSVTestSuite struct {
	suite.Suite
}

func (suite *CSVTestSuite) TestParseCSVFile() {
	db, err := ParseCSVFile("../examples/first-100.csv")

	suite.NoError(err)
	suite.Equal(100, len(db.Summits))

	summit := db.Summits[0]
	suite.Equal("3Y/BV-001", summit.SummitCode)
	suite.Equal("Bouvet Island", summit.AssociationName)
	suite.Equal("Bouvetøya (Bouvet Island)", summit.RegionName)
	suite.Equal("Olavtoppen", summit.SummitName)
	suite.Equal(780, summit.AltitudeMeters)
	suite.Equal(2559, summit.AltitudeFeet)
	suite.Equal(3.3565, summit.GridRef1)
	suite.Equal(-54.4104, summit.GridRef2)
	suite.Equal(3.3565, summit.Longitude)
	suite.Equal(-54.4104, summit.Latitude)
	suite.Equal(10, summit.Points)
	suite.Equal(3, summit.BonusPoints)
	suite.Equal("2018-03-01", summit.ValidFrom.Format(ISO8601DateFormat))
	suite.Equal("2099-12-31", summit.ValidTo.Format(ISO8601DateFormat))
	suite.Equal(0, summit.ActivationCount)
	suite.True(summit.LastActivatedOn.IsZero())

	// this one has been activated
	summit = db.Summits[89]
	suite.Equal("4O/JC-009", summit.SummitCode)
	suite.Equal("4O7DST/P", summit.LastActivatedBy)
	suite.Equal("2019-06-27", summit.LastActivatedOn.Format(ISO8601DateFormat))
}

func TestCSVTestSuite(t *testing.T) {
	suite.Run(t, new(CSVTestSuite))
}

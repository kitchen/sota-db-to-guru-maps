package sotadb

type Summit struct {
	SummitCode      string     `csv:"SummitCode"`
	AssociationName string     `csv:"AssociationName"`
	RegionName      string     `csv:"RegionName"`
	SummitName      string     `csv:"SummitName"`
	AltitudeMeters  int        `csv:"AltM"`
	AltitudeFeet    int        `csv:"AltFt"`
	GridRef1        float64    `csv:"GridRef1"`
	GridRef2        float64    `csv:"GridRef2"`
	Longitude       float64    `csv:"Longitude"`
	Latitude        float64    `csv:"Latitude"`
	Points          int        `csv:"Points"`
	BonusPoints     int        `csv:"BonusPoints"`
	ValidFrom       SummitDate `csv:"ValidFrom"`
	ValidTo         SummitDate `csv:"ValidTo"`
	ActivationCount int        `csv:"ActivationCount"`
	LastActivatedOn SummitDate `csv:"ActivationDate"`
	LastActivatedBy string     `csv:"ActivationCall"`
}

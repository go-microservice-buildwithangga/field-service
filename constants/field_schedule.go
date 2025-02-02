package constants

type FieldStatusStatusName string
type FieldScheduleStatus int

const (
	Available FieldScheduleStatus = 100
	Booked    FieldScheduleStatus = 200

	AvailableString FieldStatusStatusName = "Available"
	BookedString    FieldStatusStatusName = "Booked"
)

var mapFieldScheduleStatusIntToString = map[FieldScheduleStatus]FieldStatusStatusName{
	Available: AvailableString,
	Booked:    BookedString,
}

var mapFieldScheduleStatusStringToInt = map[FieldStatusStatusName]FieldScheduleStatus{
	AvailableString: Available,
	BookedString:    Booked,
}

func (f FieldScheduleStatus) GetStatusString() string {
	return string(mapFieldScheduleStatusIntToString[f])
}

func (f FieldStatusStatusName) GetStatusInt() FieldScheduleStatus {
	return mapFieldScheduleStatusStringToInt[f]
}


Available.GetStatusString()
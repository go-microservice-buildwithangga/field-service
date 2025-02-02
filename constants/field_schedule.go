package constants

type FieldSScheduleStatusName string
type FieldScheduleStatus int

const (
	Available FieldScheduleStatus = 100
	Booked    FieldScheduleStatus = 200

	AvailableString FieldSScheduleStatusName = "Available"
	BookedString    FieldSScheduleStatusName = "Booked"
)

var mapFieldScheduleStatusIntToString = map[FieldScheduleStatus]FieldSScheduleStatusName{
	Available: AvailableString,
	Booked:    BookedString,
}

var mapFieldScheduleStatusStringToInt = map[FieldSScheduleStatusName]FieldScheduleStatus{
	AvailableString: Available,
	BookedString:    Booked,
}

func (f FieldScheduleStatus) GetStatusString() string {
	return string(mapFieldScheduleStatusIntToString[f])
}

func (f FieldSScheduleStatusName) GetStatusInt() FieldScheduleStatus {
	return mapFieldScheduleStatusStringToInt[f]
}

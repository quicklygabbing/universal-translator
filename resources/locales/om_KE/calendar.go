package om_KE

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Hag", 1: "Ama", 6: "Wax", 3: "Bit", 7: "Ado", 9: "Ful", 10: "Onk", 4: "Elb", 11: "Sad", 12: "Mud", 2: "Gur", 5: "Cam"}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 1: "J", 9: "S", 2: "F", 3: "M", 6: "J", 10: "O", 5: "M", 7: "J", 8: "A", 11: "N", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "Hagayya", 9: "Fuulbana", 11: "Sadaasa", 10: "Onkololeessa", 1: "Amajjii", 5: "Caamsa", 6: "Waxabajjii", 7: "Adooleessa", 3: "Bitooteessa", 12: "Muddee", 2: "Guraandhala", 4: "Elba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Qib", 3: "Rob", 4: "Kam", 5: "Jim", 6: "San", 0: "Dil", 1: "Wix"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{5: "Jimaata", 6: "Sanbata", 0: "Dilbata", 1: "Wiixata", 2: "Qibxata", 3: "Roobii", 4: "Kamiisa"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "WD", "pm": "WB"}}}}

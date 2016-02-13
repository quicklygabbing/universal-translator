package en_AG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Aug", 12: "Dec", 5: "May", 4: "Apr", 9: "Sep", 2: "Feb", 7: "Jul", 10: "Oct", 1: "Jan", 6: "Jun", 11: "Nov", 3: "Mar"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 2: "F", 5: "M", 9: "S", 11: "N", 12: "D", 4: "A", 8: "A", 6: "J", 1: "J", 3: "M", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "November", 12: "December", 1: "January", 6: "June", 7: "July", 9: "September", 2: "February", 3: "March", 10: "October", 4: "April", 5: "May", 8: "August"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue"}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S"}, Short: ut.CalendarDayFormatNameValue{2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo"}, Wide: ut.CalendarDayFormatNameValue{5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p", "morning1": "in the morning"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight"}}}}

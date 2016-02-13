package en_MO

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 8: "Aug", 9: "Sep", 11: "Nov", 3: "Mar", 6: "Jun", 10: "Oct", 12: "Dec", 5: "May", 7: "Jul", 2: "Feb", 4: "Apr"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 3: "M", 4: "A", 9: "S", 11: "N", 12: "D", 2: "F", 7: "J", 8: "A", 1: "J", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "February", 3: "March", 12: "December", 10: "October", 6: "June", 7: "July", 9: "September", 11: "November", 1: "January", 4: "April", 5: "May", 8: "August"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun"}, Narrow: ut.CalendarDayFormatNameValue{4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W"}, Short: ut.CalendarDayFormatNameValue{2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo"}, Wide: ut.CalendarDayFormatNameValue{5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p", "morning1": "in the morning"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight"}}}}

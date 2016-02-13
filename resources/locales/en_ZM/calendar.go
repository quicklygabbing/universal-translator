package en_ZM

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Nov", 8: "Aug", 12: "Dec", 1: "Jan", 5: "May", 6: "Jun", 3: "Mar", 7: "Jul", 10: "Oct", 4: "Apr", 9: "Sep", 2: "Feb"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 9: "S", 1: "J", 4: "A", 10: "O", 6: "J", 2: "F", 5: "M", 11: "N", 12: "D", 3: "M", 8: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "February", 1: "January", 4: "April", 6: "June", 8: "August", 11: "November", 3: "March", 12: "December", 10: "October", 5: "May", 7: "July", 9: "September"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We"}, Wide: ut.CalendarDayFormatNameValue{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight"}}}}

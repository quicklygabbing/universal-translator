package en_NR

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "May", 10: "Oct", 2: "Feb", 3: "Mar", 11: "Nov", 6: "Jun", 9: "Sep", 1: "Jan", 7: "Jul", 8: "Aug", 12: "Dec", 4: "Apr"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 4: "A", 9: "S", 6: "J", 11: "N", 12: "D", 2: "F", 8: "A", 10: "O", 1: "J", 3: "M", 5: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "January", 6: "June", 8: "August", 9: "September", 10: "October", 11: "November", 3: "March", 12: "December", 4: "April", 5: "May", 7: "July", 2: "February"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri"}, Narrow: ut.CalendarDayFormatNameValue{4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W"}, Short: ut.CalendarDayFormatNameValue{4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We"}, Wide: ut.CalendarDayFormatNameValue{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM"}}}}

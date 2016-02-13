package af

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Des.", 4: "Apr.", 7: "Jul.", 10: "Okt.", 11: "Nov.", 6: "Jun.", 8: "Aug.", 9: "Sep.", 1: "Jan.", 2: "Feb.", 3: "Mrt.", 5: "Mei"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 4: "A", 5: "M", 6: "J", 7: "J", 8: "A", 11: "N", 1: "J", 9: "S", 10: "O", 12: "D", 3: "M"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "November", 12: "Desember", 4: "April", 6: "Junie", 7: "Julie", 8: "Augustus", 9: "September", 10: "Oktober", 1: "Januarie", 2: "Februarie", 3: "Maart", 5: "Mei"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "So.", 1: "Ma.", 2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa."}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "D", 3: "W", 4: "D", 5: "V", 6: "S"}, Short: ut.CalendarDayFormatNameValue{0: "So.", 1: "Ma.", 2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa."}, Wide: ut.CalendarDayFormatNameValue{6: "Saterdag", 0: "Sondag", 1: "Maandag", 2: "Dinsdag", 3: "Woensdag", 4: "Donderdag", 5: "Vrydag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag", "midnight": "middernag", "am": "vm.", "pm": "nm."}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "n", "midnight": "mn", "am": "v", "pm": "n", "morning1": "o", "afternoon1": "m", "evening1": "a"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "middag", "evening1": "aand", "night1": "nag", "midnight": "middernag", "am": "vm.", "pm": "nm.", "morning1": "oggend"}}}}

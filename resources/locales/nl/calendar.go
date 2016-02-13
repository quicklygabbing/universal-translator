package nl

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan.", 4: "apr.", 7: "jul.", 9: "sep.", 10: "okt.", 11: "nov.", 12: "dec.", 2: "feb.", 3: "mrt.", 5: "mei", 6: "jun.", 8: "aug."}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 2: "F", 3: "M", 4: "A", 5: "M", 6: "J", 7: "J", 8: "A", 1: "J", 9: "S", 10: "O", 11: "N"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "januari", 4: "april", 6: "juni", 7: "juli", 10: "oktober", 11: "november", 2: "februari", 3: "maart", 5: "mei", 8: "augustus", 9: "september", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma", 2: "di", 3: "wo"}, Narrow: ut.CalendarDayFormatNameValue{0: "Z", 1: "M", 2: "D", 3: "W", 4: "D", 5: "V", 6: "Z"}, Short: ut.CalendarDayFormatNameValue{0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za"}, Wide: ut.CalendarDayFormatNameValue{6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag", 4: "donderdag", 5: "vrijdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht"}}}}

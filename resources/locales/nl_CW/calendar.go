package nl_CW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "jul.", 6: "jun.", 2: "feb.", 12: "dec.", 1: "jan.", 4: "apr.", 11: "nov.", 8: "aug.", 3: "mrt.", 5: "mei", 9: "sep.", 10: "okt."}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 8: "A", 1: "J", 9: "S", 10: "O", 12: "D", 2: "F", 11: "N", 3: "M", 4: "A", 5: "M", 6: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "december", 1: "januari", 11: "november", 8: "augustus", 10: "oktober", 2: "februari", 3: "maart", 5: "mei", 9: "september", 4: "april", 6: "juni", 7: "juli"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma", 2: "di", 3: "wo"}, Narrow: ut.CalendarDayFormatNameValue{6: "Z", 0: "Z", 1: "M", 2: "D", 3: "W", 4: "D", 5: "V"}, Short: ut.CalendarDayFormatNameValue{4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma", 2: "di", 3: "wo"}, Wide: ut.CalendarDayFormatNameValue{6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag", 4: "donderdag", 5: "vrijdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag"}}}}

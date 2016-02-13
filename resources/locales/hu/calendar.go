package hu

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y. MMMM d., EEEE", Long: "y. MMMM d.", Medium: "y. MMM d.", Short: "y. MM. dd."}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "febr.", 6: "jún.", 7: "júl.", 12: "dec.", 1: "jan.", 4: "ápr.", 5: "máj.", 8: "aug.", 9: "szept.", 10: "okt.", 11: "nov.", 3: "márc."}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 12: "D", 2: "F", 6: "J", 7: "J", 5: "M", 8: "A", 9: "Sz", 10: "O", 1: "J", 3: "M", 4: "Á"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "február", 6: "június", 7: "július", 10: "október", 1: "január", 4: "április", 5: "május", 8: "augusztus", 9: "szeptember", 11: "november", 12: "december", 3: "március"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Szo", 0: "V", 1: "H", 2: "K", 3: "Sze", 4: "Cs", 5: "P"}, Narrow: ut.CalendarDayFormatNameValue{4: "Cs", 5: "P", 6: "Sz", 0: "V", 1: "H", 2: "K", 3: "Sz"}, Short: ut.CalendarDayFormatNameValue{0: "V", 1: "H", 2: "K", 3: "Sze", 4: "Cs", 5: "P", 6: "Szo"}, Wide: ut.CalendarDayFormatNameValue{0: "vasárnap", 1: "hétfő", 2: "kedd", 3: "szerda", 4: "csütörtök", 5: "péntek", 6: "szombat"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "de.", "morning1": "reggel", "morning2": "délelőtt", "afternoon1": "délután", "evening1": "este", "night2": "hajnal", "midnight": "éjfél", "noon": "dél", "pm": "du.", "night1": "éjjel"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "dél", "pm": "du.", "morning2": "délelőtt", "afternoon1": "délután", "evening1": "este", "night2": "hajnal", "midnight": "éjfél", "am": "de.", "morning1": "reggel", "night1": "éjjel"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "du.", "evening1": "este", "night2": "hajnal", "am": "de.", "noon": "dél", "morning2": "délelőtt", "afternoon1": "délután", "night1": "éjjel", "midnight": "éjfél", "morning1": "reggel"}}}}

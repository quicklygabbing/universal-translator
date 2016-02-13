package bs

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd. MMM. y.", Short: "dd.MM.yy."}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "apr", 5: "maj", 6: "jun", 7: "jul", 11: "nov", 3: "mar", 2: "feb", 8: "aug", 9: "sep", 10: "okt", 12: "dec", 1: "jan"}, Narrow: ut.CalendarMonthFormatNameValue{8: "a", 10: "o", 2: "f", 4: "a", 5: "m", 7: "j", 11: "n", 12: "d", 1: "j", 3: "m", 6: "j", 9: "s"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "april", 5: "maj", 9: "septembar", 12: "decembar", 1: "januar", 3: "mart", 6: "juni", 7: "juli", 8: "august", 10: "oktobar", 11: "novembar", 2: "februar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "sri", 4: "čet", 5: "pet", 6: "sub", 0: "ned", 1: "pon", 2: "uto"}, Narrow: ut.CalendarDayFormatNameValue{2: "u", 3: "s", 4: "č", 5: "p", 6: "s", 0: "n", 1: "p"}, Short: ut.CalendarDayFormatNameValue{3: "sri", 4: "čet", 5: "pet", 6: "sub", 0: "ned", 1: "pon", 2: "uto"}, Wide: ut.CalendarDayFormatNameValue{6: "subota", 0: "nedjelja", 1: "ponedjeljak", 2: "utorak", 3: "srijeda", 4: "četvrtak", 5: "petak"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "večer", "night1": "noć", "midnight": "ponoć", "noon": "podne", "morning1": "jutro", "afternoon1": "poslijepodne"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "prijepodne", "noon": "podne", "pm": "popodne"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"evening1": "veče", "night1": "noć", "midnight": "u ponoć", "am": "prije podne", "noon": "podne", "pm": "popodne", "morning1": "jutro", "afternoon1": "poslijepodne"}}}}

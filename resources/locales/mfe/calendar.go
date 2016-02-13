package mfe

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "avr", 5: "me", 8: "out", 12: "des", 10: "okt", 11: "nov", 1: "zan", 2: "fev", 3: "mar", 6: "zin", 7: "zil", 9: "sep"}, Narrow: ut.CalendarMonthFormatNameValue{3: "m", 4: "a", 7: "z", 9: "s", 11: "n", 1: "z", 2: "f", 8: "o", 10: "o", 12: "d", 5: "m", 6: "z"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "zin", 7: "zilye", 10: "oktob", 11: "novam", 12: "desam", 2: "fevriye", 4: "avril", 5: "me", 8: "out", 9: "septam", 1: "zanvie", 3: "mars"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "mar", 3: "mer", 4: "ze", 5: "van", 6: "sam", 0: "dim", 1: "lin"}, Narrow: ut.CalendarDayFormatNameValue{3: "m", 4: "z", 5: "v", 6: "s", 0: "d", 1: "l", 2: "m"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "lindi", 2: "mardi", 3: "merkredi", 4: "zedi", 5: "vandredi", 6: "samdi", 0: "dimans"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}

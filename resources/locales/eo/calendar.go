package eo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d-'a' 'de' MMMM y", Long: "y-MMMM-dd", Medium: "y-MMM-dd", Short: "yy-MM-dd"}, Time: ut.CalendarDateFormat{Full: "H-'a' 'horo' 'kaj' m:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan", 2: "feb", 3: "mar", 5: "maj", 8: "aŭg", 12: "dec", 4: "apr", 6: "jun", 7: "jul", 9: "sep", 10: "okt", 11: "nov"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 7: "J", 10: "O", 12: "D", 11: "N", 1: "J", 2: "F", 3: "M", 4: "A", 6: "J", 8: "A", 9: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{12: "decembro", 1: "januaro", 3: "marto", 6: "junio", 7: "julio", 8: "aŭgusto", 9: "septembro", 10: "oktobro", 2: "februaro", 4: "aprilo", 5: "majo", 11: "novembro"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "me", 4: "ĵa", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma"}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "M", 4: "Ĵ", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "mardo", 3: "merkredo", 4: "ĵaŭdo", 5: "vendredo", 6: "sabato", 0: "dimanĉo", 1: "lundo"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "atm", "pm": "ptm"}}}}

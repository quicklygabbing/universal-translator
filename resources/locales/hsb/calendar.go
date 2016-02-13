package hsb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm 'hodź'."}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "now", 12: "dec", 7: "jul", 10: "okt", 3: "měr", 4: "apr", 5: "mej", 6: "jun", 8: "awg", 9: "sep", 1: "jan", 2: "feb"}, Narrow: ut.CalendarMonthFormatNameValue{6: "j", 7: "j", 8: "a", 11: "n", 1: "j", 3: "m", 4: "a", 10: "o", 12: "d", 2: "f", 5: "m", 9: "s"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "januar", 3: "měrc", 11: "nowember", 10: "oktober", 2: "februar", 4: "apryl", 5: "meja", 6: "junij", 7: "julij", 8: "awgust", 9: "september", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "štw", 5: "pja", 6: "sob", 0: "nje", 1: "pón", 2: "wut", 3: "srj"}, Narrow: ut.CalendarDayFormatNameValue{0: "n", 1: "p", 2: "w", 3: "s", 4: "š", 5: "p", 6: "s"}, Short: ut.CalendarDayFormatNameValue{0: "nj", 1: "pó", 2: "wu", 3: "sr", 4: "št", 5: "pj", 6: "so"}, Wide: ut.CalendarDayFormatNameValue{6: "sobota", 0: "njedźela", 1: "póndźela", 2: "wutora", 3: "srjeda", 4: "štwórtk", 5: "pjatk"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"am": "dop.", "pm": "pop."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "dopołdnja", "pm": "popołdnju"}}}}

package gsw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Jul", 8: "Aug", 9: "Sep", 12: "Dez", 1: "Jan", 4: "Apr", 5: "Mai", 10: "Okt", 11: "Nov", 2: "Feb", 3: "Mär", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 1: "J", 2: "F", 4: "A", 5: "M", 6: "J", 9: "S", 10: "O", 3: "M", 7: "J", 8: "A", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "Mai", 10: "Oktoober", 12: "Dezämber", 1: "Januar", 3: "März", 4: "April", 6: "Juni", 7: "Juli", 8: "Auguscht", 9: "Septämber", 11: "Novämber", 2: "Februar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mä.", 2: "Zi.", 3: "Mi.", 4: "Du.", 5: "Fr.", 6: "Sa.", 0: "Su."}, Narrow: ut.CalendarDayFormatNameValue{2: "D", 3: "M", 4: "D", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "Dunschtig", 5: "Friitig", 6: "Samschtig", 0: "Sunntig", 1: "Määntig", 2: "Ziischtig", 3: "Mittwuch"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "v.m.", "pm": "n.m."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Vormittag", "pm": "Namittag"}}}}

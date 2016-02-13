package twq

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Okt", 11: "Noo", 12: "Dee", 5: "Me", 6: "Žuw", 7: "Žuy", 9: "Sek", 8: "Ut", 1: "Žan", 2: "Fee", 3: "Mar", 4: "Awi"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 4: "A", 7: "Ž", 11: "N", 12: "D", 1: "Ž", 5: "M", 6: "Ž", 8: "U", 9: "S", 10: "O", 3: "M"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Sektanbur", 10: "Oktoobur", 5: "Me", 6: "Žuweŋ", 3: "Marsi", 4: "Awiril", 7: "Žuyye", 8: "Ut", 11: "Noowanbur", 12: "Deesanbur", 1: "Žanwiye", 2: "Feewiriye"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Alh", 1: "Ati", 2: "Ata", 3: "Ala", 4: "Alm", 5: "Alz", 6: "Asi"}, Narrow: ut.CalendarDayFormatNameValue{1: "T", 2: "T", 3: "L", 4: "L", 5: "L", 6: "S", 0: "H"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{5: "Alzuma", 6: "Asibti", 0: "Alhadi", 1: "Atinni", 2: "Atalaata", 3: "Alarba", 4: "Alhamiisa"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Subbaahi", "pm": "Zaarikay b"}}}}

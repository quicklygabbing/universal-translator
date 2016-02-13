package tzm

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Duj", 3: "Mar", 4: "Ibr", 6: "Yun", 7: "Yul", 10: "Kṭu", 11: "Nwa", 1: "Yen", 2: "Yeb", 5: "May", 8: "Ɣuc", 9: "Cut"}, Narrow: ut.CalendarMonthFormatNameValue{6: "Y", 11: "N", 3: "M", 4: "I", 5: "M", 8: "Ɣ", 9: "C", 10: "K", 12: "D", 1: "Y", 2: "Y", 7: "Y"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "Ibrir", 5: "Mayyu", 7: "Yulyuz", 11: "Nwanbir", 12: "Dujanbir", 3: "Mars", 2: "Yebrayer", 6: "Yunyu", 8: "Ɣuct", 9: "Cutanbir", 10: "Kṭuber", 1: "Yennayer"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Ayn", 2: "Asn", 3: "Akr", 4: "Akw", 5: "Asm", 6: "Asḍ", 0: "Asa"}, Narrow: ut.CalendarDayFormatNameValue{5: "A", 6: "A", 0: "A", 1: "A", 2: "A", 3: "A", 4: "A"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{5: "Asimwas", 6: "Asiḍyas", 0: "Asamas", 1: "Aynas", 2: "Asinas", 3: "Akras", 4: "Akwas"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Zdat azal", "pm": "Ḍeffir aza"}}}}

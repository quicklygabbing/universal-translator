package cs

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "bře", 5: "kvě", 8: "srp", 9: "zář", 11: "lis", 12: "pro", 1: "led", 4: "dub", 6: "čvn", 7: "čvc", 10: "říj", 2: "úno"}, Narrow: ut.CalendarMonthFormatNameValue{10: "10", 12: "12", 2: "2", 4: "4", 6: "6", 7: "7", 9: "9", 11: "11", 1: "1", 3: "3", 5: "5", 8: "8"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "březen", 4: "duben", 7: "červenec", 8: "srpen", 9: "září", 11: "listopad", 1: "leden", 2: "únor", 5: "květen", 6: "červen", 10: "říjen", 12: "prosinec"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "pá", 6: "so", 0: "ne", 1: "po", 2: "út", 3: "st", 4: "čt"}, Narrow: ut.CalendarDayFormatNameValue{5: "P", 6: "S", 0: "N", 1: "P", 2: "Ú", 3: "S", 4: "Č"}, Short: ut.CalendarDayFormatNameValue{5: "pá", 6: "so", 0: "ne", 1: "po", 2: "út", 3: "st", 4: "čt"}, Wide: ut.CalendarDayFormatNameValue{2: "úterý", 3: "středa", 4: "čtvrtek", 5: "pátek", 6: "sobota", 0: "neděle", 1: "pondělí"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "dopoledne", "afternoon1": "odpoledne", "evening1": "večer", "midnight": "půlnoc", "am": "dop.", "noon": "poledne", "pm": "odp.", "morning1": "ráno", "night1": "noc"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "půl.", "noon": "pol.", "night1": "noc", "morning2": "dop.", "afternoon1": "odp.", "evening1": "več.", "am": "dop.", "pm": "odp.", "morning1": "ráno"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "odp.", "afternoon1": "odpoledne", "evening1": "večer", "night1": "noc", "noon": "poledne", "am": "dop.", "morning1": "ráno", "morning2": "dopoledne", "midnight": "půlnoc"}}}}

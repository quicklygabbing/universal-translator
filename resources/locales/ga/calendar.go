package ga

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "Feabh", 3: "Márta", 4: "Aib", 8: "Lún", 9: "MFómh", 10: "DFómh", 12: "Noll", 1: "Ean", 5: "Beal", 6: "Meith", 7: "Iúil", 11: "Samh"}, Narrow: ut.CalendarMonthFormatNameValue{1: "E", 5: "B", 6: "M", 10: "D", 11: "S", 12: "N", 2: "F", 3: "M", 4: "A", 7: "I", 8: "L", 9: "M"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "Samhain", 1: "Eanáir", 2: "Feabhra", 7: "Iúil", 9: "Meán Fómhair", 10: "Deireadh Fómhair", 12: "Nollaig", 3: "Márta", 4: "Aibreán", 5: "Bealtaine", 6: "Meitheamh", 8: "Lúnasa"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Céad", 4: "Déar", 5: "Aoine", 6: "Sath", 0: "Domh", 1: "Luan", 2: "Máirt"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "C", 4: "D", 5: "A"}, Short: ut.CalendarDayFormatNameValue{4: "Dé", 5: "Ao", 6: "Sa", 0: "Do", 1: "Lu", 2: "Má", 3: "Cé"}, Wide: ut.CalendarDayFormatNameValue{0: "Dé Domhnaigh", 1: "Dé Luain", 2: "Dé Máirt", 3: "Dé Céadaoin", 4: "Déardaoin", 5: "Dé hAoine", 6: "Dé Sathairn"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "am": "a"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m."}}}}

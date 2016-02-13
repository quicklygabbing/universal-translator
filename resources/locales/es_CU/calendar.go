package es_CU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "jun.", 10: "oct.", 2: "feb.", 11: "nov.", 1: "ene.", 9: "sept.", 4: "abr.", 7: "jul.", 12: "dic.", 3: "mar.", 5: "may.", 8: "ago."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 12: "D", 1: "E", 6: "J", 3: "M", 11: "N", 2: "F", 8: "A", 5: "M", 7: "J", 9: "S", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "abril", 12: "diciembre", 6: "junio", 7: "julio", 1: "enero", 2: "febrero", 3: "marzo", 8: "agosto", 9: "septiembre", 10: "octubre", 11: "noviembre", 5: "mayo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar."}, Narrow: ut.CalendarDayFormatNameValue{2: "M", 3: "X", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L"}, Short: ut.CalendarDayFormatNameValue{2: "MA", 3: "MI", 4: "JU", 5: "VI", 6: "SA", 0: "DO", 1: "LU"}, Wide: ut.CalendarDayFormatNameValue{4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde"}}}}

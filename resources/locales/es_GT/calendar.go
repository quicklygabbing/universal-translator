package es_GT

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d/MM/y", Short: "d/MM/yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "feb.", 3: "mar.", 6: "jun.", 9: "sept.", 11: "nov.", 5: "may.", 8: "ago.", 7: "jul.", 10: "oct.", 12: "dic.", 1: "ene.", 4: "abr."}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 1: "E", 4: "A", 8: "A", 3: "M", 2: "F", 10: "O", 6: "J", 5: "M", 7: "J", 9: "S", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "febrero", 8: "agosto", 9: "septiembre", 1: "enero", 3: "marzo", 4: "abril", 5: "mayo", 10: "octubre", 6: "junio", 7: "julio", 11: "noviembre", 12: "diciembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar.", 3: "mié."}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "X", 4: "J", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI", 6: "SA", 0: "DO"}, Wide: ut.CalendarDayFormatNameValue{0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}}}}

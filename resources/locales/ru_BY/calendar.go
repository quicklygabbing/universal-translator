package ru_BY

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "нояб.", 9: "сент.", 10: "окт.", 2: "февр.", 3: "март", 4: "апр.", 6: "июнь", 8: "авг.", 12: "дек.", 5: "май", 7: "июль", 1: "янв."}, Narrow: ut.CalendarMonthFormatNameValue{11: "Н", 12: "Д", 1: "Я", 6: "И", 8: "А", 10: "О", 2: "Ф", 7: "И", 3: "М", 4: "А", 9: "С", 5: "М"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "февраль", 3: "март", 10: "октябрь", 12: "декабрь", 1: "январь", 8: "август", 9: "сентябрь", 11: "ноябрь", 5: "май", 6: "июнь", 4: "апрель", 7: "июль"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "чт", 5: "пт", 6: "сб", 0: "вс", 1: "пн", 2: "вт", 3: "ср"}, Narrow: ut.CalendarDayFormatNameValue{6: "С", 0: "В", 1: "П", 2: "В", 3: "С", 4: "Ч", 5: "П"}, Short: ut.CalendarDayFormatNameValue{2: "вт", 3: "ср", 4: "чт", 5: "пт", 6: "сб", 0: "вс", 1: "пн"}, Wide: ut.CalendarDayFormatNameValue{0: "воскресенье", 1: "понедельник", 2: "вторник", 3: "среда", 4: "четверг", 5: "пятница", 6: "суббота"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "полн.", "am": "ДП", "noon": "полд.", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "веч.", "night1": "ночь"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "полн.", "am": "ДП", "noon": "полд.", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "веч.", "night1": "ночь"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "ночь", "midnight": "полночь", "am": "ДП", "noon": "полдень", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "вечер"}}}}

package yo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Owewe", 12: "Ọ̀pẹ̀", 2: "Èrèlè", 3: "Ẹrẹ̀nà", 4: "Ìgbé", 5: "Ẹ̀bibi", 7: "Agẹmọ", 8: "Ògún", 1: "Ṣẹ́rẹ́", 6: "Òkúdu", 10: "Ọ̀wàrà", 11: "Bélú"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "Oṣù Ọ̀wàrà", 11: "Oṣù Bélú", 2: "Oṣù Èrèlè", 3: "Oṣù Ẹrẹ̀nà", 4: "Oṣù Ìgbé", 8: "Oṣù Ògún", 9: "Oṣù Owewe", 12: "Oṣù Ọ̀pẹ̀", 1: "Oṣù Ṣẹ́rẹ́", 5: "Oṣù Ẹ̀bibi", 6: "Oṣù Òkúdu", 7: "Oṣù Agẹmọ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Àìkú", 1: "Ajé", 2: "Ìsẹ́gun", 3: "Ọjọ́rú", 4: "Ọjọ́bọ", 5: "Ẹtì", 6: "Àbámẹ́ta"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Ọjọ́ Àìkú", 1: "Ọjọ́ Ajé", 2: "Ọjọ́ Ìsẹ́gun", 3: "Ọjọ́rú", 4: "Ọjọ́bọ", 5: "Ọjọ́ Ẹtì", 6: "Ọjọ́ Àbámẹ́ta"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Àárọ̀", "pm": "Ọ̀sán"}}}}

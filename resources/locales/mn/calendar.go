package mn

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, y 'оны' MM 'сарын' d", Long: "y 'оны' MM 'сарын' d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "7-р сар", 11: "11-р сар", 12: "12-р сар", 1: "1-р сар", 2: "2-р сар", 3: "3-р сар", 5: "5-р сар", 6: "6-р сар", 4: "4-р сар", 8: "8-р сар", 9: "9-р сар", 10: "10-р сар"}, Narrow: ut.CalendarMonthFormatNameValue{9: "9", 10: "10", 12: "12", 1: "1", 7: "7", 4: "4", 5: "5", 6: "6", 8: "8", 11: "11", 2: "2", 3: "3"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "Наймдугаар сар", 2: "Хоёрдугаар сар", 4: "Дөрөвдүгээр сар", 5: "Тавдугаар сар", 6: "Зургадугаар сар", 10: "Аравдугаар сар", 11: "Арван нэгдүгээр сар", 12: "Арван хоёрдугаар сар", 1: "Нэгдүгээр сар", 3: "Гуравдугаар сар", 7: "Долдугаар сар", 9: "Есдүгээр сар"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Ня", 1: "Да", 2: "Мя", 3: "Лх", 4: "Пү", 5: "Ба", 6: "Бя"}, Narrow: ut.CalendarDayFormatNameValue{0: "1", 1: "2", 2: "3", 3: "4", 4: "5", 5: "6", 6: "7"}, Short: ut.CalendarDayFormatNameValue{0: "Ня", 1: "Да", 2: "Мя", 3: "Лх", 4: "Пү", 5: "Ба", 6: "Бя"}, Wide: ut.CalendarDayFormatNameValue{1: "даваа", 2: "мягмар", 3: "лхагва", 4: "пүрэв", 5: "баасан", 6: "бямба", 0: "ням"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"noon": "үд", "pm": "үх", "morning1": "өглөө", "afternoon1": "өдөр", "evening1": "орой", "night1": "шөнө", "midnight": "шөнө дунд", "am": "үө"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "ҮХ", "morning1": "өглөө", "afternoon1": "өдөр", "evening1": "орой", "night1": "шөнө", "midnight": "шөнө дунд", "am": "ҮӨ", "noon": "үд дунд"}}}}

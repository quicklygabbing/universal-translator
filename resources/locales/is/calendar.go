package is

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "d.M.y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "okt.", 12: "des.", 2: "feb.", 3: "mar.", 4: "apr.", 6: "jún.", 8: "ágú.", 1: "jan.", 5: "maí", 7: "júl.", 9: "sep.", 11: "nóv."}, Narrow: ut.CalendarMonthFormatNameValue{8: "Á", 9: "S", 10: "O", 11: "N", 1: "J", 3: "M", 4: "A", 5: "M", 2: "F", 6: "J", 7: "J", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "apríl", 5: "maí", 7: "júlí", 9: "september", 11: "nóvember", 12: "desember", 2: "febrúar", 3: "mars", 6: "júní", 8: "ágúst", 10: "október", 1: "janúar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "mið.", 4: "fim.", 5: "fös.", 6: "lau.", 0: "sun.", 1: "mán.", 2: "þri."}, Narrow: ut.CalendarDayFormatNameValue{6: "L", 0: "S", 1: "M", 2: "Þ", 3: "M", 4: "F", 5: "F"}, Short: ut.CalendarDayFormatNameValue{4: "fi.", 5: "fö.", 6: "la.", 0: "su.", 1: "má.", 2: "þr.", 3: "mi."}, Wide: ut.CalendarDayFormatNameValue{0: "sunnudagur", 1: "mánudagur", 2: "þriðjudagur", 3: "miðvikudagur", 4: "fimmtudagur", 5: "föstudagur", 6: "laugardagur"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "miðn.", "noon": "hád.", "morning1": "morg.", "afternoon1": "síðd.", "evening1": "kv.", "night1": "nótt"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "kv.", "night1": "n.", "midnight": "mn.", "noon": "hd.", "morning1": "mrg.", "afternoon1": "sd."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "nótt", "midnight": "miðnætti", "am": "f.h.", "noon": "hádegi", "pm": "e.h.", "morning1": "morgunn", "afternoon1": "eftir hádegi", "evening1": "kvöld"}}}}

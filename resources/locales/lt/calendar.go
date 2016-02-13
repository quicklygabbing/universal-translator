package lt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "rugs.", 12: "gruod.", 3: "kov.", 4: "bal.", 6: "birž.", 7: "liep.", 8: "rugp.", 10: "spal.", 11: "lapkr.", 1: "saus.", 2: "vas.", 5: "geg."}, Narrow: ut.CalendarMonthFormatNameValue{1: "S", 2: "V", 3: "K", 4: "B", 7: "L", 5: "G", 6: "B", 8: "R", 9: "R", 10: "S", 11: "L", 12: "G"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "kovas", 8: "rugpjūtis", 9: "rugsėjis", 11: "lapkritis", 12: "gruodis", 10: "spalis", 1: "sausis", 2: "vasaris", 4: "balandis", 5: "gegužė", 6: "birželis", 7: "liepa"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "an", 3: "tr", 4: "kt", 5: "pn", 6: "št", 0: "sk", 1: "pr"}, Narrow: ut.CalendarDayFormatNameValue{4: "K", 5: "P", 6: "Š", 0: "S", 1: "P", 2: "A", 3: "T"}, Short: ut.CalendarDayFormatNameValue{0: "Sk", 1: "Pr", 2: "An", 3: "Tr", 4: "Kt", 5: "Pn", 6: "Št"}, Wide: ut.CalendarDayFormatNameValue{1: "pirmadienis", 2: "antradienis", 3: "trečiadienis", 4: "ketvirtadienis", 5: "penktadienis", 6: "šeštadienis", 0: "sekmadienis"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis", "pm": "popiet", "morning1": "rytas"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "vidurdienis", "pm": "pop.", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "pr. p."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis", "pm": "popiet"}}}}

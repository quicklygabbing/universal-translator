package nmg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "ng2", 3: "ng3", 5: "ng5", 7: "ng7", 9: "ng9", 10: "ng10", 1: "ng1", 4: "ng4", 6: "ng6", 8: "ng8", 11: "ng11", 12: "kris"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "ngwɛn matáhra", 2: "ngwɛn ńmba", 3: "ngwɛn ńlal", 4: "ngwɛn ńna", 7: "ngwɛn hɛmbuɛrí", 9: "ngwɛn rɛbvuâ", 5: "ngwɛn ńtan", 6: "ngwɛn ńtuó", 8: "ngwɛn lɔmbi", 10: "ngwɛn wum", 11: "ngwɛn wum navǔr", 12: "krísimin"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "sɔ́n", 1: "mɔ́n", 2: "smb", 3: "sml", 4: "smn", 5: "mbs", 6: "sas"}, Narrow: ut.CalendarDayFormatNameValue{0: "s", 1: "m", 2: "s", 3: "s", 4: "s", 5: "m", 6: "s"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "sɔ́ndɔ", 1: "mɔ́ndɔ", 2: "sɔ́ndɔ mafú mába", 3: "sɔ́ndɔ mafú málal", 4: "sɔ́ndɔ mafú mána", 5: "mabágá má sukul", 6: "sásadi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "maná", "pm": "kugú"}}}}

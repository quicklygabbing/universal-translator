package ti_ER

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE፡ dd MMMM መዓልቲ y G", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "ታሕሳ", 1: "ጥሪ", 4: "ሚያዝ", 8: "ነሓሰ", 10: "ጥቅም", 11: "ሕዳር", 9: "መስከ", 2: "ለካቲ", 3: "መጋቢ", 5: "ግንቦ", 6: "ሰነ", 7: "ሓምለ"}, Narrow: ut.CalendarMonthFormatNameValue{10: "ኦ", 9: "ሴ", 1: "ጃ", 5: "ሜ", 12: "ዲ", 2: "ፌ", 3: "ማ", 4: "ኤ", 6: "ጁ", 11: "ኖ", 7: "ጁ", 8: "ኦ"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "ለካቲት", 4: "ሚያዝያ", 6: "ሰነ", 9: "መስከረም", 10: "ጥቅምቲ", 11: "ሕዳር", 1: "ጥሪ", 3: "መጋቢት", 5: "ግንቦት", 7: "ሓምለ", 8: "ነሓሰ", 12: "ታሕሳስ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{}, Narrow: ut.CalendarDayFormatNameValue{2: "ሠ", 3: "ረ", 4: "ኃ", 5: "ዓ", 6: "ቀ", 0: "ሰ", 1: "ሰ"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{2: "ሰሉስ", 3: "ረቡዕ", 4: "ሓሙስ", 5: "ዓርቢ", 6: "ቀዳም", 0: "ሰንበት", 1: "ሰኑይ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "ንጉሆ ሰዓተ", "pm": "ድሕር ሰዓት"}}}}

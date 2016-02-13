package mas

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "Rok", 7: "Sás", 2: "Ará", 3: "Ɔɛn", 5: "Lép", 8: "Bɔ́r", 9: "Kús", 10: "Gís", 11: "Shʉ́", 12: "Ntʉ́", 1: "Dal", 4: "Doy"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "Olodoyíóríê inkókúâ", 5: "Oloilépūnyīē inkókúâ", 6: "Kújúɔrɔk", 10: "Olgísan", 11: "Pʉshʉ́ka", 1: "Oladalʉ́", 2: "Arát", 8: "Ɔlɔ́ɨ́bɔ́rárɛ", 9: "Kúshîn", 12: "Ntʉ́ŋʉ́s", 3: "Ɔɛnɨ́ɔɨŋɔk", 7: "Mórusásin"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Jnn", 3: "Jtn", 4: "Alh", 5: "Iju", 6: "Jmo", 0: "Jpi", 1: "Jtt"}, Narrow: ut.CalendarDayFormatNameValue{3: "5", 4: "6", 5: "7", 6: "1", 0: "2", 1: "3", 2: "4"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Jumatátu", 2: "Jumane", 3: "Jumatánɔ", 4: "Alaámisi", 5: "Jumáa", 6: "Jumamósi", 0: "Jumapílí"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Ɛnkakɛnyá", "pm": "Ɛndámâ"}}}}

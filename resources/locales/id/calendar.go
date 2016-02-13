package id

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 2: "Feb", 5: "Mei", 7: "Jul", 8: "Agt", 10: "Okt", 12: "Des", 3: "Mar", 4: "Apr", 6: "Jun", 9: "Sep", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 2: "F", 4: "A", 6: "J", 7: "J", 11: "N", 3: "M", 5: "M", 8: "A", 9: "S", 10: "O", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Januari", 2: "Februari", 5: "Mei", 6: "Juni", 8: "Agustus", 11: "November", 3: "Maret", 4: "April", 7: "Juli", 9: "September", 10: "Oktober", 12: "Desember"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Min", 1: "Sen", 2: "Sel", 3: "Rab", 4: "Kam", 5: "Jum", 6: "Sab"}, Narrow: ut.CalendarDayFormatNameValue{1: "S", 2: "S", 3: "R", 4: "K", 5: "J", 6: "S", 0: "M"}, Short: ut.CalendarDayFormatNameValue{5: "Jum", 6: "Sab", 0: "Min", 1: "Sen", 2: "Sel", 3: "Rab", 4: "Kam"}, Wide: ut.CalendarDayFormatNameValue{2: "Selasa", 3: "Rabu", 4: "Kamis", 5: "Jumat", 6: "Sabtu", 0: "Minggu", 1: "Senin"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "pagi", "afternoon1": "siang", "evening1": "sore", "night1": "malam", "midnight": "tengah malam", "am": "AM", "noon": "tengah hari", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "malam", "midnight": "tengah malam", "am": "AM", "noon": "tengah hari", "pm": "PM", "morning1": "pagi", "afternoon1": "siang", "evening1": "sore"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "tengah hari", "pm": "PM", "morning1": "pagi", "afternoon1": "siang", "evening1": "sore", "night1": "malam", "midnight": "tengah malam"}}}}

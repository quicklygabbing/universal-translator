package rw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "ugu.", 12: "uku.", 1: "mut.", 3: "wer.", 4: "mat.", 6: "kam.", 7: "nya.", 8: "kan.", 2: "gas.", 5: "gic.", 9: "nze.", 10: "ukw."}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "Werurwe", 10: "Ukwakira", 12: "Ukuboza", 6: "Kamena", 7: "Nyakanga", 8: "Kanama", 9: "Nzeli", 1: "Mutarama", 2: "Gashyantare", 4: "Mata", 5: "Gicuransi", 11: "Ugushyingo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "cyu.", 1: "mbe.", 2: "kab.", 3: "gtu.", 4: "kan.", 5: "gnu.", 6: "gnd."}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Kuwa gatandatu", 0: "Ku cyumweru", 1: "Kuwa mbere", 2: "Kuwa kabiri", 3: "Kuwa gatatu", 4: "Kuwa kane", 5: "Kuwa gatanu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}

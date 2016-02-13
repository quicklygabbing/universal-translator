package haw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Kek.", 1: "Ian.", 3: "Mal.", 6: "Iun.", 8: "ʻAu.", 10: "ʻOk.", 11: "Now.", 2: "Pep.", 4: "ʻAp.", 5: "Mei", 7: "Iul.", 9: "Kep."}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "ʻApelila", 5: "Mei", 6: "Iune", 8: "ʻAukake", 9: "Kepakemapa", 10: "ʻOkakopa", 1: "Ianuali", 2: "Pepeluali", 11: "Nowemapa", 12: "Kekemapa", 3: "Malaki", 7: "Iulai"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "P5", 6: "P6", 0: "LP", 1: "P1", 2: "P2", 3: "P3", 4: "P4"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "Poʻalua", 3: "Poʻakolu", 4: "Poʻahā", 5: "Poʻalima", 6: "Poʻaono", 0: "Lāpule", 1: "Poʻakahi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}

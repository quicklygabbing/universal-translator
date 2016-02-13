package ar

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "مايو", 8: "أغسطس", 9: "سبتمبر", 10: "أكتوبر", 11: "نوفمبر", 1: "يناير", 2: "فبراير", 3: "مارس", 12: "ديسمبر", 4: "أبريل", 6: "يونيو", 7: "يوليو"}, Narrow: ut.CalendarMonthFormatNameValue{3: "م", 6: "ن", 8: "غ", 9: "س", 11: "ب", 12: "د", 1: "ي", 2: "ف", 4: "أ", 5: "و", 7: "ل", 10: "ك"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "فبراير", 3: "مارس", 4: "أبريل", 8: "أغسطس", 9: "سبتمبر", 12: "ديسمبر", 1: "يناير", 5: "مايو", 6: "يونيو", 7: "يوليو", 10: "أكتوبر", 11: "نوفمبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}, Narrow: ut.CalendarDayFormatNameValue{4: "خ", 5: "ج", 6: "س", 0: "ح", 1: "ن", 2: "ث", 3: "ر"}, Short: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}, Wide: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "ص", "pm": "م", "afternoon2": "بعد الظهر", "evening1": "مساءً", "morning1": "فجرا", "morning2": "صباحًا", "afternoon1": "ظهرًا", "night1": "منتصف الليل", "night2": "ليلاً"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "ظهرًا", "night2": "ليلاً", "pm": "م", "morning1": "فجرا", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "evening1": "مساءً", "night1": "منتصف الليل", "am": "ص"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "منتصف الليل", "am": "صباحًا", "morning1": "فجرا", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "pm": "مساءً", "afternoon1": "ظهرًا", "evening1": "مساءً", "night2": "ليلاً"}}}}

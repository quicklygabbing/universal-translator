package ar_DZ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "مارس", 4: "أفريل", 5: "ماي", 6: "جوان", 9: "سبتمبر", 11: "نوفمبر", 2: "فيفري", 7: "جويلية", 8: "أوت", 10: "أكتوبر", 12: "ديسمبر", 1: "جانفي"}, Narrow: ut.CalendarMonthFormatNameValue{7: "ج", 9: "س", 10: "أ", 12: "د", 1: "ج", 4: "أ", 6: "ج", 8: "أ", 11: "ن", 2: "ف", 3: "م", 5: "م"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "جانفي", 2: "فيفري", 6: "جوان", 8: "أوت", 9: "سبتمبر", 12: "ديسمبر", 3: "مارس", 4: "أفريل", 5: "ماي", 7: "جويلية", 10: "أكتوبر", 11: "نوفمبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}, Narrow: ut.CalendarDayFormatNameValue{1: "ن", 2: "ث", 3: "ر", 4: "خ", 5: "ج", 6: "س", 0: "ح"}, Short: ut.CalendarDayFormatNameValue{1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد"}, Wide: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "م", "night1": "منتصف الليل", "morning2": "صباحًا", "afternoon1": "ظهرًا", "night2": "ليلاً", "am": "ص", "afternoon2": "بعد الظهر", "evening1": "مساءً", "morning1": "فجرا"}, Narrow: ut.CalendarPeriodFormatNameValue{"night2": "ليلاً", "morning1": "فجرا", "morning2": "صباحًا", "evening1": "مساءً", "night1": "منتصف الليل", "afternoon1": "ظهرًا", "afternoon2": "بعد الظهر", "am": "ص", "pm": "م"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "منتصف الليل", "morning2": "صباحًا", "pm": "مساءً", "afternoon1": "ظهرًا", "night2": "ليلاً", "am": "صباحًا", "morning1": "فجرا", "afternoon2": "بعد الظهر", "evening1": "مساءً"}}}}

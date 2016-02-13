package fa_AF

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "جول", 12: "دسم", 6: "ژوئن", 11: "نوامبر", 2: "فوریه", 1: "جنو", 5: "می", 8: "اوت", 10: "اکتبر", 3: "مارس", 4: "آوریل", 9: "سپتامبر"}, Narrow: ut.CalendarMonthFormatNameValue{10: "ا", 1: "ج", 2: "ف", 5: "م", 6: "ج", 7: "ج", 8: "ا", 3: "م", 4: "ا", 9: "س", 11: "ن", 12: "د"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "جولای", 8: "اگست", 10: "اکتوبر", 12: "دسمبر", 3: "مارچ", 4: "اپریل", 5: "می", 6: "جون", 9: "سپتمبر", 11: "نومبر", 1: "جنوری", 2: "فبروری"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه"}, Narrow: ut.CalendarDayFormatNameValue{2: "س", 3: "چ", 4: "پ", 5: "ج", 6: "ش", 0: "ی", 1: "د"}, Short: ut.CalendarDayFormatNameValue{2: "۳ش", 3: "۴ش", 4: "۵ش", 5: "ج", 6: "ش", 0: "۱ش", 1: "۲ش"}, Wide: ut.CalendarDayFormatNameValue{3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه", 0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "ظ", "pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ."}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظهر"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "قبل\u200cازظهر", "noon": "ظهر", "pm": "بعدازظهر", "morning1": "صبح", "afternoon1": "عصر"}}}}

package si

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "ජූනි", 8: "අගෝ", 9: "සැප්", 2: "පෙබ", 3: "මාර්", 5: "මැයි", 7: "ජූලි", 10: "ඔක්", 11: "නොවැ", 12: "දෙසැ", 1: "ජන", 4: "අප්\u200dරේල්"}, Narrow: ut.CalendarMonthFormatNameValue{5: "මැ", 7: "ජූ", 9: "සැ", 11: "නෙ", 12: "දෙ", 1: "ජ", 4: "අ", 6: "ජූ", 8: "අ", 10: "ඔ", 2: "පෙ", 3: "මා"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "ඔක්තෝබර්", 11: "නොවැම්බර්", 4: "අප්\u200dරේල්", 9: "සැප්තැම්බර්", 3: "මාර්තු", 5: "මැයි", 6: "ජූනි", 7: "ජූලි", 8: "අගෝස්තු", 12: "දෙසැම්බර්", 1: "ජනවාරි", 2: "පෙබරවාරි"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "අඟහ", 3: "බදාදා", 4: "බ්\u200dරහස්", 5: "සිකු", 6: "සෙන", 0: "ඉරිදා", 1: "සඳුදා"}, Narrow: ut.CalendarDayFormatNameValue{4: "බ්\u200dර", 5: "සි", 6: "සෙ", 0: "ඉ", 1: "ස", 2: "අ", 3: "බ"}, Short: ut.CalendarDayFormatNameValue{0: "ඉරි", 1: "සඳු", 2: "අඟ", 3: "බදා", 4: "බ්\u200dරහ", 5: "සිකු", 6: "සෙන"}, Wide: ut.CalendarDayFormatNameValue{6: "සෙනසුරාදා", 0: "ඉරිදා", 1: "සඳුදා", 2: "අඟහරුවාදා", 3: "බදාදා", 4: "බ්\u200dරහස්පතින්දා", 5: "සිකුරාදා"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "ප.ව.", "morning2": "උදේ", "night1": "රෑ", "night2": "මැදියමට පසු", "midnight": "මැදියම", "am": "පෙ.ව.", "noon": "මධ්\u200dයාහ්නය", "morning1": "පාන්දර", "afternoon1": "දවල්", "evening1": "හවස"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "හ", "midnight": "මැ", "noon": "ද", "pm": "ප", "morning2": "උ", "night2": "මැ", "am": "පෙ", "morning1": "පා", "afternoon1": "ද", "night1": "රෑ"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night2": "මැදියමට පසු", "am": "පෙ.ව.", "pm": "ප.ව.", "morning1": "පාන්දර", "morning2": "උදේ", "afternoon1": "දවල්", "evening1": "හවස", "night1": "රෑ", "midnight": "මැදියම", "noon": "මධ්\u200dයාහ්නය"}}}}

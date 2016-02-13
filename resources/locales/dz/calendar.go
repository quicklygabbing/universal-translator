package dz

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, སྤྱི་ལོ་y MMMM ཚེས་dd", Long: "སྤྱི་ལོ་y MMMM ཚེས་ dd", Medium: "སྤྱི་ལོ་y ཟླ་MMM ཚེས་dd", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a zzzz", Long: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a z", Medium: "ཆུ་ཚོད་h:mm:ss a", Short: "ཆུ་ཚོད་ h སྐར་མ་ mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "ཟླ་༡༡", 1: "ཟླ་༡", 3: "ཟླ་༣", 4: "ཟླ་༤", 5: "ཟླ་༥", 8: "ཟླ་༨", 10: "ཟླ་༡༠", 2: "ཟླ་༢", 6: "ཟླ་༦", 7: "ཟླ་༧", 9: "ཟླ་༩", 12: "ཟླ་༡༢"}, Narrow: ut.CalendarMonthFormatNameValue{8: "༨", 10: "༡༠", 4: "༤", 5: "༥", 3: "༣", 6: "༦", 7: "༧", 9: "༩", 11: "༡༡", 12: "༡༢", 1: "༡", 2: "༢"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "སྤྱི་ཟླ་བརྒྱད་པ་", 11: "སྤྱི་ཟླ་བཅུ་གཅིག་པ་", 1: "སྤྱི་ཟླ་དངཔ་", 3: "སྤྱི་ཟླ་གསུམ་པ་", 4: "སྤྱི་ཟླ་བཞི་པ", 5: "སྤྱི་ཟླ་ལྔ་པ་", 6: "སྤྱི་ཟླ་དྲུག་པ", 2: "སྤྱི་ཟླ་གཉིས་པ་", 7: "སྤྱི་ཟླ་བདུན་པ་", 9: "སྤྱི་ཟླ་དགུ་པ་", 10: "སྤྱི་ཟླ་བཅུ་པ་", 12: "སྤྱི་ཟླ་བཅུ་གཉིས་པ་"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "ཕུར་", 4: "སངས་", 5: "སྤེན་", 6: "ཉི་", 0: "ཟླ་", 1: "མིར་", 2: "ལྷག་"}, Narrow: ut.CalendarDayFormatNameValue{6: "ཉི", 0: "ཟླ", 1: "མིར", 2: "ལྷག", 3: "ཕུར", 4: "སངྶ", 5: "སྤེན"}, Short: ut.CalendarDayFormatNameValue{4: "སངས་", 5: "སྤེན་", 6: "ཉི་", 0: "ཟླ་", 1: "མིར་", 2: "ལྷག་", 3: "ཕུར་"}, Wide: ut.CalendarDayFormatNameValue{0: "གཟའ་ཟླ་བ་", 1: "གཟའ་མིག་དམར་", 2: "གཟའ་ལྷག་པ་", 3: "གཟའ་ཕུར་བུ་", 4: "གཟའ་པ་སངས་", 5: "གཟའ་སྤེན་པ་", 6: "གཟའ་ཉི་མ་"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"pm": "ཕྱི་ཆ་", "am": "སྔ་ཆ་"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "སྔ་ཆ་", "pm": "ཕྱི་ཆ་"}}}}

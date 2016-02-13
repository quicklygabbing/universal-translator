package fr_MA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "avr.", 7: "juil.", 10: "oct.", 11: "nov.", 1: "jan.", 2: "fév.", 3: "mar.", 5: "mai", 6: "jui.", 8: "août", 9: "sept.", 12: "déc."}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 1: "J", 8: "A", 11: "N", 4: "A", 6: "J", 7: "J", 2: "F", 5: "M", 10: "O", 12: "D", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 9: "septembre", 4: "avril", 1: "janvier", 6: "juin", 10: "octobre", 11: "novembre", 2: "février", 3: "mars", 7: "juillet", 8: "août", 12: "décembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me"}, Wide: ut.CalendarDayFormatNameValue{6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "noon": "midi", "am": "a.m.", "pm": "p.m."}}}}

package fr_MR

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "juin", 1: "janv.", 4: "avr.", 5: "mai", 9: "sept.", 10: "oct.", 3: "mars", 12: "déc.", 8: "août", 11: "nov.", 2: "févr.", 7: "juil."}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 3: "M", 4: "A", 9: "S", 5: "M", 7: "J", 11: "N", 6: "J", 8: "A", 10: "O", 1: "J", 2: "F"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "mars", 5: "mai", 7: "juillet", 8: "août", 12: "décembre", 6: "juin", 9: "septembre", 2: "février", 10: "octobre", 11: "novembre", 1: "janvier", 4: "avril"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{3: "me", 4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma"}, Wide: ut.CalendarDayFormatNameValue{6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir"}}}}

package he

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"DOP": ut.Currency{Currency: "DOP", DisplayName: "פזו דומיניקני", Symbol: "DOP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "וון דרום-קוריאני", Symbol: "₩"}, "WST": ut.Currency{Currency: "WST", DisplayName: "טאלה סמואי", Symbol: "WST"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "מרק גרמני", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "סדי גאני", Symbol: "GHS"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "לירה ישראלית", Symbol: "ל״י"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "לוטי לסותי", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "דינר לובי", Symbol: "LYD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "דינר טוניסאי", Symbol: "TND"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "קורונה צ׳כית", Symbol: "CZK"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "מרק מזרח גרמני", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "פאונד גיברלטר", Symbol: "GIP"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "פזו ארגנטינאי (1983–1985)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "שילינג אוסטרי", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "דינר אלג׳ירי", Symbol: "DZD"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "וון צפון-קוריאני", Symbol: "KPW"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "רופי פקיסטני", Symbol: "PKR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "סומוני טג׳קיסטני", Symbol: "TJS"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "פורינט הונגרי", Symbol: "HUF"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "ליאו מולדובני", Symbol: "MDL"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ריאל עומאני", Symbol: "OMR"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "פזטה אנדורית", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "דינר בחרייני", Symbol: "BHD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "מרק פיני", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "פרנק צרפתי", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "דולר הונג קונגי", Symbol: "HK$"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "לירה קפריסאית", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "פרנק מרוקאי", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "לירה מלטית", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "פרנק רואנדי", Symbol: "RWF"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "דולר סורינאמי", Symbol: "SRD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "פרנק בלגי", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "קרוזדו ברזילאי", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "דלאסי גמבי", Symbol: "GMD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "פאנגה טונגי", Symbol: "TOP"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "פרנק שוויצרי", Symbol: "CHF"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "פזו מקסיקני", Symbol: "MX$"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "סוקר אקואדורי", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "אירו", Symbol: "€"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "שילינג אוגנדי (1966 – 1987)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "זהב", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "דולר ברבדיאני", Symbol: "BBD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "בוליביאנו", Symbol: "BOB"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "פלטינה", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "דראם ארמני", Symbol: "AMD"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "קואנזה רג׳וסטדו אנגולי (1995–1999)", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "דינר כוויתי", Symbol: "KWD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "סימון למטרות בדיקה", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "פסו צ׳ילאני", Symbol: "CLP"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "זאיר חדש", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "גילדר של האנטילים ההולנדיים", Symbol: "ANG"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "לירה סודנית", Symbol: "SDG"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "לירה דרום-סודנית", Symbol: "SSP"}, "USS": ut.Currency{Currency: "USS", DisplayName: "דולר אמריקאי (היום הזה)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "פרנק CFA BEAC", Symbol: "FCFA"}, "COP": ut.Currency{Currency: "COP", DisplayName: "פסו קולומביאני", Symbol: "COP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "לירה של איי פוקלנד", Symbol: "FKP"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "ליטא ליטאי", Symbol: "LTL"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "דולר נמיבי", Symbol: "NAD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "דולר מזרח קריבי", Symbol: "EC$"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "פרנק בלגי (בר המרה)", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "אסקודו כף ורדה", Symbol: "CVE"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "דינר סודני", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "דונג וייטנאמי", Symbol: "₫"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "דולר איי שלמה", Symbol: "SBD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "לירה לבנונית", Symbol: "LBP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "בהט תאילנדי", Symbol: "฿"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "מנאט אזרביג׳אני (1993–2006)", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "קצל גואטמלי", Symbol: "GTQ"}, "YER": ut.Currency{Currency: "YER", DisplayName: "ריאל תימני", Symbol: "YER"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "דולר בהאמי", Symbol: "BSD"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "דולר קנדי", Symbol: "CA$"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "דולר אוסטרלי", Symbol: "A$"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "רופיה מלדיבית", Symbol: "MVR"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "נאירה ניגרי", Symbol: "NGN"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "טוגריק מונגולי", Symbol: "MNT"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "קוואצ׳ה זמבית", Symbol: "ZMW"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "דולר ברמודה", Symbol: "BMD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "פזו קובני להמרה", Symbol: "CUC"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "לירה אירית", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "סום קירגיזי", Symbol: "KGS"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "סום אוזבקי", Symbol: "UZS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "טולאר סלובני", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ריאל סעודי", Symbol: "SAR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "לירה סודנית (1957–1998)", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "רובל סובייטי", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "קיפ לאי", Symbol: "LAK"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "קואצ׳ה מלאוי", Symbol: "MWK"}, "AON": ut.Currency{Currency: "AON", DisplayName: "קואנזה חדש אנגולי (1990–2000)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "מאנאט אזרביג׳ני", Symbol: "AZN"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "קרון אסטוני", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "לירה מצרית", Symbol: "EGP"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "למפירה הונדורי", Symbol: "HNL"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "פרנק לוקסמבורגי", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "פרנק מדגסקארי", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "פרנק ג׳יבוטי", Symbol: "DJF"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "לירה שטרלינג", Symbol: "£"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "פזו פיליפיני", Symbol: "PHP"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "רובל רוסי", Symbol: "RUB"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "דינר יגוסלבי", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "פלורין של ארובה", Symbol: "AWG"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "דינר עיראקי", Symbol: "IQD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "גוורני פראגוואי", Symbol: "PYG"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "שילינג טנזני", Symbol: "TZS"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "פרנק CFA BCEAO", Symbol: "CFA"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "ראנד דרום אפריקאי (כספי)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "יואן סיני", Symbol: "CN¥"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "כתר דני", Symbol: "DKK"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "דולר סינגפורי", Symbol: "SGD"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "קורונה סלובקי", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "טאקה בנגלדשי", Symbol: "BDT"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "לב בולגרי ישן", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "פרנק בורונדי", Symbol: "BIF"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "דרכמה", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "רופיה אינדונזית", Symbol: "IDR"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ריאל איראני", Symbol: "IRR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "דירהם מרוקאי", Symbol: "MAD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "דולר טייוואני חדש", Symbol: "NT$"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "טנגה קזחסטני", Symbol: "KZT"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "אוגוויה מאוריטני", Symbol: "MRO"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "בלבואה פנמי", Symbol: "PAB"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "מנאט טורקמאני", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "לירה טורקית", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "כסף", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "אפגני אפגני", Symbol: "AFN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "פולה בוצוואני", Symbol: "BWP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "שילינג קנייאתי", Symbol: "KES"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "רופי סרי לנקי", Symbol: "LKR"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "כתר נורבגי", Symbol: "NOK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "לאו רומני", Symbol: "RON"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "נגולטרום בהוטני", Symbol: "BTN"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "פסטה ספרדי", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "קרונה איסלנדית", Symbol: "ISK"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "דולר ליברי", Symbol: "LRD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "פטקה של מקאו", Symbol: "MOP"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "אסקודו מוזמביקי", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "מטיקל", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "שילינג סומאלי", Symbol: "SOS"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "שילינג אוגנדי", Symbol: "UGX"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "דינר של בוסניה־הרצגובינה", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "לב בולגרי", Symbol: "BGN"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "רובל רוסי (1991 – 1998)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "זכויות משיכה מיוחדות", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "מטבע שאינו ידוע", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "קרוזיארו חדש ברזילאי (1967–1986)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "לרי גאורגי", Symbol: "GEL"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ין יפני", Symbol: "JP¥"}, "USD": ut.Currency{Currency: "USD", DisplayName: "דולר אמריקאי", Symbol: "$"}, "AED": ut.Currency{Currency: "AED", DisplayName: "דירהם של איחוד הנסיכויות הערביות", Symbol: "AED"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "לק אלבני", Symbol: "ALL"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "פזו בוליבי", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "דינר מקדוני", Symbol: "MKD"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "פזו מקסיקני (1861 – 1992)", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "דולר ניו זילנדי", Symbol: "NZ$"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "דולר זימבבואי", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "רובל בלרוסי", Symbol: "BYR"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "ריל קמבודי", Symbol: "KHR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "פסו אורוגוואי", Symbol: "UYU"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "פזטה [ESB]", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "דולר בליזי", Symbol: "BZD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "פרנק קונגולזי", Symbol: "CDF"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "דינר סרבי", Symbol: "RSD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "כתר שוודי", Symbol: "SEK"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "לירה טורקית חדשה", Symbol: "TRY"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "דינר תימני", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "ראנד דרום אפריקאי", Symbol: "ZAR"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "קואנזה אנגולי", Symbol: "AOA"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "דינר סרבי ישן", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "רינגיט מלזי", Symbol: "MYR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "זלוטי פולני", Symbol: "PLN"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "אסקודו טימוראי", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "ואטו של ונואטו", Symbol: "VUV"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "דולר טרינידדי", Symbol: "TTD"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "פזטה [ESA]", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "רופי נפאלי", Symbol: "NPR"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "בוליבר ונצואלי", Symbol: "VEF"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "אסקודו פורטוגלי", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "גילדר סורינאמי", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "לירה איטלקית", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "דולר ג׳מייקני", Symbol: "JMD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "דולר קיימאני", Symbol: "KYD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ליאון סיירה לאוני", Symbol: "SLL"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "גריבנה אוקראיני", Symbol: "UAH"}, "BND": ut.Currency{Currency: "BND", DisplayName: "דולר ברוניי", Symbol: "BND"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "פזו גינאי", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "דינר ירדני", Symbol: "JOD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "קורדובה ניקראגי", Symbol: "NIO"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "מארק בר המרה של בוסניה־הרצגובינה", Symbol: "BAM"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "גורד האיטי", Symbol: "HTG"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "אריארי מלגשי", Symbol: "MGA"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "רופי מאוריציני", Symbol: "MUR"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "קולון סלבדורי", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ביר אתיופי", Symbol: "ETB"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "קינה של פפואה גינאה החדשה", Symbol: "PGK"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "זלוטי (1950 – 1995)", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "רופי סיישלי", Symbol: "SCR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "דוברה של סן טומה ופרינסיפה", Symbol: "STD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "מאנאט טורקמני", Symbol: "TMT"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "פרנק זהב", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "דינר יגוסלבי חדש", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "קוואצ׳ה זמבית (1968–2012)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "פלדיום", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "פסו ארגנטינאי", Symbol: "ARS"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "קיאט מיאנמרי", Symbol: "MMK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "סול פרואני חדש", Symbol: "PEN"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ריאל קטארי", Symbol: "QAR"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ריאל ברזילאי", Symbol: "R$"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "קונה קרואטי", Symbol: "HRK"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "לירה סורית", Symbol: "SYP"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "דולר גיאני", Symbol: "GYD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ש״ח", Symbol: "₪"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "מטיקל מוזמביני", Symbol: "MZN"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "קולון קוסטה־ריקני", Symbol: "CRC"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "נאקפה אריתראי", Symbol: "ERN"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "בוליבר ונצואלי (1871–2008)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "פרנק פולינזיה הצרפתית", Symbol: "CFPF"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "לילנגני סווזי", Symbol: "SZL"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "פרנק גינאי", Symbol: "GNF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "פאונד סנט הלני", Symbol: "SHP"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "פזו קובני", Symbol: "CUP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "לט לטבי", Symbol: "LVL"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "גילדן הולנדי", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "לאו רומני ישן", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "דולר פיג׳י", Symbol: "FJD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "רופי הודי", Symbol: "₹"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "פרנק קומורואי", Symbol: "KMF"}, "USN": ut.Currency{Currency: "USN", DisplayName: "דולר אמריקאי (היום הבא)", Symbol: ""}}

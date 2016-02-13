package it

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltese", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguaiano (1975–1993)", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum uzbeco", Symbol: "UZS"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso della Guinea-Bissau", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduregna", Symbol: "HNL"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollaro liberiano", Symbol: "LRD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "scellino keniota", Symbol: "KES"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia di Sri Lanka", Symbol: "LKR"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinaro sudanese", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinaro forte yugoslavo", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasiliano", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreo", Symbol: "ERN"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo della Guinea portoghese", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danese", Symbol: "DKK"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "złoty polacco", Symbol: "PLN"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona slovacca", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinaro dello Yemen", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marco convertibile della Bosnia-Herzegovina", Symbol: "BAM"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombiano", Symbol: "COP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicano", Symbol: "MZN"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinaro serbo", Symbol: "RSD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollaro zimbabwiano (2009)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca di Macao", Symbol: "MOP"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "złoty Polacco (1950–1995)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "diritti speciali di incasso", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladio", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolano (1977–1990)", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidades de fomento chilene", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "sterlina israeliana", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rublo della CSI", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona svedese", Symbol: "SEK"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollaro della Guyana", Symbol: "GYD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo del Mozambico", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollaro della Rhodesia", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinaro noviy yugoslavo", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franco congolese", Symbol: "CDF"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversion (UDI) messicana", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "leu della Romania", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocchino", Symbol: "MAD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanita", Symbol: "OMR"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele della Guinea Equatoriale", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz ucraino", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platino", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franco belga", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bhutanese", Symbol: "BTN"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franco svizzero", Symbol: "CHF"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thailandese", Symbol: "฿"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguaiano in unità indicizzate", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentino", Symbol: "ARS"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) dell’Ecuador", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghiso", Symbol: "KGS"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "fiorino delle Antille olandesi", Symbol: "ANG"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinaro algerino", Symbol: "DZD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona dell’Estonia", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argento", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat di Myanmar", Symbol: "MMK"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani", Symbol: "AFN"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "sterlina delle Falkland", Symbol: "FKP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambiano", Symbol: "GMD"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni dello Swaziland", Symbol: "SZL"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgaro", Symbol: "BGN"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliviano", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "lira siriana", Symbol: "SYP"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruviano", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tagiko", Symbol: "TJS"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unità composita europea", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza reajustado angolano (1995–1999)", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capoverdiano", Symbol: "CVE"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panamense", Symbol: "PAB"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha dello Zambia", Symbol: "ZMW"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba nicaraguense", Symbol: "NIO"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar venezuelano (1871–2008)", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sudafricano", Symbol: "ZAR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franco CFP", Symbol: "CFPF"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha dello Zambia (1968–2012)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "scellino austriaco", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nuovo siclo israeliano", Symbol: "₪"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franco comoriano", Symbol: "KMF"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "sterlina egiziana", Symbol: "EGP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgiano", Symbol: "GEL"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sudcoreano", Symbol: "KRW"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazako", Symbol: "KZT"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "fiorino di Aruba", Symbol: "AWG"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmano", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ostmark della Germania Orientale", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricano", Symbol: "CRC"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "markka finlandese", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublo lettone", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "riyal yemenita", Symbol: "YER"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol boliviano", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brasiliano", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollaro canadese", Symbol: "CA$"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriziana", Symbol: "MUR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayano", Symbol: "PYG"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinaro tunisino", Symbol: "TND"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azero (1993–2006)", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franco di Mali", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "sterlina maltese", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta spagnola", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "fiorino ungherese", Symbol: "HUF"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso messicano", Symbol: "MXN"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen giapponese", Symbol: "JPY"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambogiano", Symbol: "KHR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filippino", Symbol: "PHP"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollaro di Trinidad e Tobago", Symbol: "TTD"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentino (vecchio Cod.)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marco tedesco", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franco della Guinea", Symbol: "GNF"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franco belga (finanziario)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitiano", Symbol: "HTG"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona norvegese", Symbol: "NOK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franco ruandese", Symbol: "RWF"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicano", Symbol: "DOP"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iraniano", Symbol: "IRR"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laotiano", Symbol: "LAK"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "sterlina irlandese", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti del Lesotho", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistana", Symbol: "PKR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollaro del Suriname", Symbol: "SRD"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "renminbi cinese", Symbol: "CN¥"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma greca", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "sterlina cipriota", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spagnola account convertibile", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgascio", Symbol: "MGA"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tallero sloveno", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nuovo zaire dello Zaire", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franco del Burundi", Symbol: "BIF"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "cruzeiro novo brasiliano (1967–1986)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollaro delle Bahamas", Symbol: "BSD"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli della Guinea", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublo russo", Symbol: "RUB"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unità di acconto europea (XBD)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinaro convertibile yugoslavo", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona forte cecoslovacca", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra di Sao Tomé e Principe", Symbol: "STD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franco CFA BCEAO", Symbol: "CFA"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franco belga (convertibile)", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollaro delle Bermuda", Symbol: "BMD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollaro delle Figi", Symbol: "FJD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmeno (1993–2009)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmeno", Symbol: "TMT"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franco finanziario del Lussemburgo", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayano", Symbol: "UYU"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso cileno", Symbol: "CLP"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubano", Symbol: "CUP"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldavo", Symbol: "MDL"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre dell’Ecuador", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollaro statunitense (next day)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinaro croato", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollaro dei Caraibi orientali", Symbol: "EC$"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bosnia-Herzegovina", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladese", Symbol: "BDT"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "scellino ugandese (1966–1987)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezuelano", Symbol: "VEF"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoano", Symbol: "WST"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "valuta sconosciuta", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiope", Symbol: "ETB"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "scellino della Tanzania", Symbol: "TZS"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongolo", Symbol: "MNT"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya della Mauritania", Symbol: "MRO"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollaro neozelandese", Symbol: "NZ$"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollaro delle Isole Salomone", Symbol: "SBD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham degli Emirati Arabi Uniti", Symbol: "AED"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nuovo kwanza angolano (1990–2000)", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franco marocchino", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "oro", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "cruzado novo brasiliano", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollaro del Belize", Symbol: "BZD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "sterlina britannica", Symbol: "£"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nuovo rublo bielorusso (1994–1999)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nuovo sol peruviano", Symbol: "PEN"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paʻanga tongano", Symbol: "TOP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "scellino ugandese", Symbol: "UGX"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollaro statunitense", Symbol: "US$"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "antico dinaro serbo", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi del Ghana", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublo del Tajikistan", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriana", Symbol: "NGN"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo di Timor", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unità di acconto europea (XBC)", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sudafricano (finanziario)", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublo bielorusso", Symbol: "BYR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia indiana", Symbol: "₹"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "scellino somalo", Symbol: "SOS"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "fiorino del Suriname", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollaro del Brunei", Symbol: "BND"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nordcoreano", Symbol: "KPW"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franco del Lussemburgo", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franco oro francese", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonesiana", Symbol: "IDR"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguense", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papuana", Symbol: "PGK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "sterlina sud-sudanese", Symbol: "SSP"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublo sovietico", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamita", Symbol: "₫"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollaro delle Isole Cayman", Symbol: "KYD"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franco malgascio", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollaro namibiano", Symbol: "NAD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinaro libico", Symbol: "LYD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia delle Seychelles", Symbol: "SCR"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinaro del Bahrein", Symbol: "BHD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "sterlina di Gibilterra", Symbol: "GIP"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinaro giordano", Symbol: "JOD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macedone", Symbol: "MKD"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portoghese", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "sterlina di Sant’Elena", Symbol: "SHP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croata", Symbol: "HRK"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinaro kuwaitiano", Symbol: "KWD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa delle Maldive", Symbol: "MVR"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire dello Zaire", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona ceca", Symbol: "CZK"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franco di Gibuti", Symbol: "DJF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats lettone", Symbol: "LVL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatariano", Symbol: "QAR"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "codice di verifica della valuta", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollaro statunitense (same day)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanese", Symbol: "ALL"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubano convertibile", Symbol: "CUC"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgiano", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollaro di Hong Kong", Symbol: "HKD"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sterlina sudanese", Symbol: "SDG"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgaro (1962–1999)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasiliano", Symbol: "BRL"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula del Botswana", Symbol: "BWP"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malese", Symbol: "MYR"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unità monetaria europea", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franco UIC francese", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "fiorino olandese", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "riyal saudita", Symbol: "SAR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollaro di Singapore", Symbol: "SGD"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franco francese", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinaro iracheno", Symbol: "IQD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawiano", Symbol: "MWK"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalteco", Symbol: "GTQ"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituani", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadoregno", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruviano", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nuovo dollaro taiwanese", Symbol: "TWD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolano", Symbol: "AOA"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brasiliano (1990–1993)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandese", Symbol: "ISK"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso messicano d’argento (1861–1992)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalese", Symbol: "NPR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "grivnia ucraina", Symbol: "UAH"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fondi RINET", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollaro australiano", Symbol: "A$"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franco convertibile del Lussemburgo", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franco CFA BEAC", Symbol: "FCFA"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "lira libanese", Symbol: "LBP"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituano", Symbol: "LTL"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumeno", Symbol: "RON"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone della Sierra Leone", Symbol: "SLL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu di Vanuatu", Symbol: "VUV"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollaro dello Zimbabwe", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentino", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azero", Symbol: "AZN"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollaro giamaicano", Symbol: "JMD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armeno", Symbol: "AMD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanese", Symbol: "GHS"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollaro di Barbados", Symbol: "BBD"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spagnola account", Symbol: ""}}

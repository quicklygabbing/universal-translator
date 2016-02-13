package ro

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AUD": ut.Currency{Currency: "AUD", DisplayName: "dolar australian", Symbol: "AUD"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "liră malteză", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "gulden Surinam", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublă belarusă", Symbol: "BYR"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertibil luxemburghez", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC francez", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guineean", Symbol: "GNF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazahă", Symbol: "KZT"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc Madagascar", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc francez de aur", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "leva bulgărească", Symbol: "BGN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula Botswana", Symbol: "BWP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo din Capul Verde", Symbol: "CVE"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre Ecuador", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-african (financiar)", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-african", Symbol: "ZAR"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "aur", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azer (1993–2006)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar Serbia și Muntenegru (2002–2006)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "liră italiană", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thailandez", Symbol: "THB"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundez", Symbol: "BIF"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "coroană islandeză", Symbol: "ISK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexican", Symbol: "MXN"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zlot polonez", Symbol: "PLN"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo Mozambic", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambican", Symbol: "MZN"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unitate de cont europeană (XBD)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platină", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "marcă est-germană", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanez", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruvian (1863–1965)", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivian", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolez", Symbol: "CDF"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubanez convertibil", Symbol: "CUC"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coreean", Symbol: "KRW"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "carboavă ucraineană", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unitate monetară europeană", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu din Vanuatu", Symbol: "VUV"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc elvețian", Symbol: "CHF"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "liră sterlină", Symbol: "GBP"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupie mauritiană", Symbol: "MUR"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "coroană norvegiană", Symbol: "NOK"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "leu românesc (1952–2006)", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmen (1993–2009)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dolar din Belize", Symbol: "BZD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "liră din Gibraltar", Symbol: "GIP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupie indoneziană", Symbol: "IDR"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorian", Symbol: "KMF"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dolar din Brunei", Symbol: "BND"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "coroană estoniană", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litu lituanian", Symbol: "LTL"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupie nepaleză", Symbol: "NPR"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldiviană", Symbol: "MVR"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominican", Symbol: "DOP"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "pesetă spaniolă", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croată", Symbol: "HRK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonez", Symbol: "JPY"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde din Haiti", Symbol: "HTG"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "coroană slovacă", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint maghiar", Symbol: "HUF"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argint", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham din Emiratele Arabe Unite", Symbol: "AED"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "liră sudaneză", Symbol: "SDG"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "șiling tanzanian", Symbol: "TZS"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "șechel israelian nou", Symbol: "ILS"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuweitian", Symbol: "KWD"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso mexican de argint (1861–1992)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "metical Mozambic vechi", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupie indiană", Symbol: "INR"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dolar rhodesian", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar sloven", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA BEAC", Symbol: "FCFA"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unitate compusă europeană", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angoleză", Symbol: "AOA"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduriană", Symbol: "HNL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "sol nou peruvian", Symbol: "PEN"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum Uzbekistan", Symbol: "UZS"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algerian", Symbol: "DZD"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financiar luxemburghez", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar sârbesc", Symbol: "RSD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisian", Symbol: "TND"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dolar american", Symbol: "USD"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bosnia-Herțegovina (1992–1994)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dolar canadian", Symbol: "CAD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "liră din Insulele Falkland", Symbol: "FKP"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothian", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spaniolă (cont convertibil)", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjic", Symbol: "TJS"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "liră egipteană", Symbol: "EGP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocan", Symbol: "MAD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakian", Symbol: "IQD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc Mali", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar venezuelean", Symbol: "VEF"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "gulden din Antilele Olandeze", Symbol: "ANG"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "liră cipriotă", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "șiling somalez", Symbol: "SOS"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni din Swaziland", Symbol: "SZL"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar Yemen", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanez", Symbol: "GHS"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dolar guyanez", Symbol: "GYD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupie din Sri Lanka", Symbol: "LKR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudit", Symbol: "SAR"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "liră turcească (1922–2005)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dolar american (ziua următoare)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "coroană cehă", Symbol: "CZK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgian", Symbol: "GEL"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanez", Symbol: "OMR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dolar surinamez", Symbol: "SRD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "monedă necunoscută", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dolar liberian", Symbol: "LRD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina din Papua-Noua Guinee", Symbol: "PGK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "liră Insula Sf. Elena", Symbol: "SHP"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladiu", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi din Gambia", Symbol: "GMD"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar Venezuela (1871–2008)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivian", Symbol: "BOB"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum din Bhutan", Symbol: "BTN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan chinezesc", Symbol: "CNY"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marcă germană", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin aruban", Symbol: "AWG"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamez", Symbol: "VND"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zair nou", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat din Myanmar", Symbol: "MMK"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azer", Symbol: "AZN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar din Bahrain", Symbol: "BHD"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "liră irlandeză", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublă Letonia", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croat", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonean", Symbol: "MKD"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unitate de monedă europeană", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "liră libaneză", Symbol: "LBP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "liră sudaneză (1957–1998)", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublă Tadjikistan", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belgian (convertibil)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brazilian (1993–1994)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dolar jamaican", Symbol: "JMD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brazilian (1990–1993)", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libian", Symbol: "LYD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "dolar nou din Taiwan", Symbol: "TWD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dolar din Caraibele de Est", Symbol: "XCD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "leka albaneză", Symbol: "ALL"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dolar din Insulele Cayman", Symbol: "KYD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "CFPF"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso columbian", Symbol: "COP"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriană", Symbol: "NGN"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayan", Symbol: "UYU"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kârgâz", Symbol: "KGS"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublă sovietică", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "EUR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "coroană suedeză", Symbol: "SEK"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldovenesc", Symbol: "MDL"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar iugoslav convertibil", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambian (1968–2012)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "cod monetar de test", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambian", Symbol: "ZMW"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "pesetă andorrană", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilian", Symbol: "CLP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiopian", Symbol: "ETB"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dolar din Trinidad-Tobago", Symbol: "TTD"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "șiling austriac", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coreean", Symbol: "KPW"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unitate de cont europeană (XBC)", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dolar Zimbabwe (2008)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dolar din Bahamas", Symbol: "BSD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca din Macao", Symbol: "MOP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "liră turcească", Symbol: "TRY"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dolar din Bermuda", Symbol: "BMD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "coroană daneză", Symbol: "DKK"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc francez", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso Guineea-Bissau", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar iordanian", Symbol: "JOD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letonian", Symbol: "LVL"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brazilian", Symbol: "BRL"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalez", Symbol: "GTQ"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar iugoslav nou", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupie din Seychelles", Symbol: "SCR"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso Uruguay (1975–1993)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marcă convertibilă din Bosnia și Herțegovina", Symbol: "BAM"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxemburghez", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruvian", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublă rusească", Symbol: "RUB"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "șiling ugandez", Symbol: "UGX"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani afgan", Symbol: "AFN"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dolar Barbados", Symbol: "BBD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dolar fijian", Symbol: "FJD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone din Sierra Leone", Symbol: "SLL"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentinian", Symbol: "ARS"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar iugoslav greu", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandez", Symbol: "RWF"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "liră sud-sudaneză", Symbol: "SSP"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colon El Salvador", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "liră siriană", Symbol: "SYP"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubanez", Symbol: "CUP"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutian", Symbol: "DJF"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dolar din Hong Kong", Symbol: "HKD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgian", Symbol: "KHR"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "drepturi speciale de tragere", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "șiling ugandez (1966–1987)", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoană", Symbol: "WST"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemenit", Symbol: "YER"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "cordoba nicaraguană", Symbol: "NIO"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipinez", Symbol: "PHP"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupie pakistaneză", Symbol: "PKR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dolar Singapore", Symbol: "SGD"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentinian (1983–1985)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka din Bangladesh", Symbol: "BDT"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spaniolă (cont A)", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drahmă grecească", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armenesc", Symbol: "AMD"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi Ghana (1979–2007)", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dolar namibian", Symbol: "NAD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belgian", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreeană", Symbol: "ERN"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra Sao Tome și Principe", Symbol: "STD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guarani paraguayan", Symbol: "PYG"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatarian", Symbol: "QAR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongană", Symbol: "TOP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgaș", Symbol: "MGA"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawiană", Symbol: "MWK"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malaiezian", Symbol: "MYR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zlot polonez (1950–1995)", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranian", Symbol: "IRR"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmen", Symbol: "TMT"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvna ucraineană", Symbol: "UAH"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dolar neozeelandez", Symbol: "NZD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameză", Symbol: "PAB"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu românesc", Symbol: "RON"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dolar din Insulele Solomon", Symbol: "SBD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belgian (financiar)", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colon costarican", Symbol: "CRC"}, "KES": ut.Currency{Currency: "KES", DisplayName: "șiling kenyan", Symbol: "KES"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritană", Symbol: "MRO"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dolar Zimbabwe (2009)", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "gulden olandez", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivian", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marcă finlandeză", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "liră israeliană", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laoțian", Symbol: "LAK"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocan", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguană (1988–1991)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dolar american (aceeași zi)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA BCEAO", Symbol: "CFA"}}

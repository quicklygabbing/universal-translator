package fr_TN

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicain", Symbol: "MZN"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florin antillais", Symbol: "ANG"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "livre turque (1844–2005)", Symbol: "TRL"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argent", Symbol: "XAG"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituanien", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "nouveau dinar yougoslave", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "couronne danoise", Symbol: "DKK"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "livre libanaise", Symbol: "£LB"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat myanmarais", Symbol: "MMK"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombien", Symbol: "$CO"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricain", Symbol: "CRC"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxembourgeois", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "mark allemand", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro", Symbol: "BRR"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgare (1962–1999)", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rouble soviétique", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "roupie indienne", Symbol: "₹"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar bahaméen", Symbol: "$BS"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar barbadien", Symbol: "BBD"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundais", Symbol: "BIF"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ancien leu roumain", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unité d’investissement chilienne", Symbol: "CLF"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saoudien", Symbol: "SAR"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rouble biélorusse", Symbol: "BYR"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigérian", Symbol: "NGN"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatari", Symbol: "QAR"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "livre sterling", Symbol: "£GB"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswanais", Symbol: "BWP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nouveau shekel israélien", Symbol: "₪"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: "BUK"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicain", Symbol: "$MX"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croate", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierra-léonais", Symbol: "SLL"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "métical", Symbol: "MZM"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar soudanais", Symbol: "SDD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranien", Symbol: "IRR"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lire maltaise", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "livre égyptienne", Symbol: "EGP"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta espagnole", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr éthiopien", Symbol: "ETB"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubain", Symbol: "CUP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "livre syrienne", Symbol: "SYP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadien", Symbol: "$CA"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coréen", Symbol: "₩"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldave", Symbol: "MDL"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonais", Symbol: "PLN"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bouthanais", Symbol: "BTN"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilien", Symbol: "CLE"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanien", Symbol: "MRO"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurien", Symbol: "HNL"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothan", Symbol: "lLS"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmène", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "livre de Gibraltar", Symbol: "£GI"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar de Singapour", Symbol: "$SG"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "mark finlandais", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: "ARA"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: "LVL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malien", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belge (convertible)", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar brunéien", Symbol: "$BN"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani afghan", Symbol: "AFN"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgien", Symbol: "KHR"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unité monétaire européenne", Symbol: "XBB"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar de Hong Kong", Symbol: "HKD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "roupie srilankaise", Symbol: "LKR"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocain", Symbol: "fMA"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thaïlandais", Symbol: "THB"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonais", Symbol: "JPY"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar slovène", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar libérien", Symbol: "LRD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malais", Symbol: "MYR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suisse", Symbol: "CHF"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "FCFP"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivien", Symbol: "BOV"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgare", Symbol: "BGN"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belge", Symbol: "FB"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanais", Symbol: "OMR"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haïtienne", Symbol: "HTG"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar vénézuélien", Symbol: "VEF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar jamaïcain", Symbol: "JMD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belge (financier)", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar trinidadien", Symbol: "$TT"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrane", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platine", Symbol: "XPT"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "livre chypriote", Symbol: "£CY"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or", Symbol: "XFO"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croate", Symbol: "HRK"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libyen", Symbol: "LYD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambien", Symbol: "GMD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari géorgien", Symbol: "GEL"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilling tanzanien", Symbol: "TZS"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso bissau-guinéen", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panaméen", Symbol: "PAB"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawite", Symbol: "MWK"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba", Symbol: "NIC"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisien", Symbol: "DT"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc français", Symbol: "F"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram arménien", Symbol: "AMD"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorais", Symbol: "TPE"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rouble russe (1991–1998)", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar yougoslave Noviy", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayen", Symbol: "PYG"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinée portugaise", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldivien", Symbol: "MVR"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar du Guyana", Symbol: "GYD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "livre des îles Malouines", Symbol: "£FK"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "livre soudanaise", Symbol: "SDG"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africain", Symbol: "ZAR"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nouveau cruzado", Symbol: "BRN"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "couronne norvégienne", Symbol: "NOK"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol péruvien", Symbol: "PES"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip loatien", Symbol: "LAK"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilling somalien", Symbol: "SOS"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar néo-zélandais", Symbol: "$NZ"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra santoméen", Symbol: "STD"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rhodésien", Symbol: "$RH"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbo-monténégrin", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financier luxembourgeois", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: "AFA"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorien", Symbol: "KMF"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "type de fonds RINET", Symbol: "XRE"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "nouveau manat turkmène", Symbol: "TMT"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azéri", Symbol: "AZN"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guinéen", Symbol: "GNF"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macédonien", Symbol: "MKD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florin néerlandais", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC", Symbol: "XFU"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unité européenne composée", Symbol: "XBA"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "livre turque", Symbol: "TRY"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladeshi", Symbol: "BDT"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: "CDF"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoan", Symbol: "WS$"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutien", Symbol: "DJF"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuan", Symbol: "VUV"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar koweïtien", Symbol: "KWD"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu roumain", Symbol: "RON"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivien", Symbol: "BOB"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "roupie indonésienne", Symbol: "IDR"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australien", Symbol: "$AU"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicain", Symbol: "DOP"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "couronne tchèque", Symbol: "CZK"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "roupie des Seychelles", Symbol: "SCR"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint hongrois", Symbol: "HUF"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanien", Symbol: "JOD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "roupie pakistanaise", Symbol: "PKR"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinois", Symbol: "CNY"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "couronne suédoise", Symbol: "SEK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso philippin", Symbol: "PHP"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rouble russe", Symbol: "RUB"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nafka érythréen", Symbol: "ERN"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coréen", Symbol: "KPW"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "roupie népalaise", Symbol: "NPR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar des îles Caïmans", Symbol: "KYD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar des États-Unis", Symbol: "$US"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "livre maltaise", Symbol: "£MT"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar zimbabwéen", Symbol: "ZWD"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "livre irlandaise", Symbol: "£IE"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "mark est-allemand", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "roupie mauricienne", Symbol: "MUR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilling kényan", Symbol: "KES"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "droit de tirage spécial", Symbol: "DTS"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lire italienne", Symbol: "₤IT"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "livre sud-soudanaise", Symbol: "SSP"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosniaque", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar des îles Salomon", Symbol: "$SB"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "roupie maldivienne", Symbol: "MVP"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghize", Symbol: "KGS"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgache", Symbol: "MGA"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambien", Symbol: "ZMW"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxembourgeois", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rouble tadjik", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unité de compte européenne (ECU)", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yéménite", Symbol: "YER"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty (1950–1995)", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambicain", Symbol: "MZE"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "couronne estonienne", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca macanaise", Symbol: "MOP"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorien", Symbol: "SVC"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "couronne slovaque", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "livre israélienne", Symbol: "£IL"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling autrichien", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgache", Symbol: "Fmg"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar bélizéen", Symbol: "$BZ"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dông vietnamien", Symbol: "₫"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilling ougandais", Symbol: "UGX"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "couronne islandaise", Symbol: "ISK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinamais", Symbol: "$SR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti péruvien", Symbol: "PEI"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayen", Symbol: "$UY"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubain convertible", Symbol: "CUC"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cédi", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unité de compte ADB", Symbol: "XUA"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbe", Symbol: "RSD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cédi ghanéen", Symbol: "GHS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "réal brésilien", Symbol: "R$"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta espagnole (compte convertible)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar yougoslave convertible", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capverdien", Symbol: "CVE"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachme grecque", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolais", Symbol: "AOA"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum ouzbek", Symbol: "UZS"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar bermudien", Symbol: "$BM"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilien", Symbol: "$CL"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ukrainienne", Symbol: "UAH"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guinéen", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakien", Symbol: "IQD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongan", Symbol: "TOP"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivien", Symbol: "BOP"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar fidjien", Symbol: "$FJ"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandais", Symbol: "RWF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algérien", Symbol: "DZD"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: "$AR"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "couronne forte tchécoslovaque", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nouveau sol péruvien", Symbol: "PEN"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocain", Symbol: "MAD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "devise inconnue ou non valide", Symbol: "XXX"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "coupon de lari géorgien", Symbol: "GEK"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar du Yémen", Symbol: "YDD"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rouble letton", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "mark convertible bosniaque", Symbol: "BAM"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibien", Symbol: "$NA"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre équatorien", Symbol: "ECS"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florin surinamais", Symbol: "SRG"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreïni", Symbol: "BHD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubais", Symbol: "AWG"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanais", Symbol: "ALL"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaïre zaïrois", Symbol: "ZRZ"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta espagnole (compte A)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituanien", Symbol: "LTL"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "(devise de test)", Symbol: "XTS"}}

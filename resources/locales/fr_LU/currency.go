package fr_LU

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GEK": ut.Currency{Currency: "GEK", DisplayName: "coupon de lari géorgien", Symbol: "GEK"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbo-monténégrin", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: "$AR"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu roumain", Symbol: "RON"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar des îles Caïmans", Symbol: "KYD"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar jamaïcain", Symbol: "JMD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "roupie mauricienne", Symbol: "MUR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandais", Symbol: "RWF"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "mark allemand", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar barbadien", Symbol: "BBD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldivien", Symbol: "MVR"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubain convertible", Symbol: "CUC"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "FCFP"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lire italienne", Symbol: "₤IT"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar bélizéen", Symbol: "$BZ"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreïni", Symbol: "BHD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or", Symbol: "XFO"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinamais", Symbol: "$SR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ukrainienne", Symbol: "UAH"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayen", Symbol: "PYG"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "couronne forte tchécoslovaque", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coréen", Symbol: "KPW"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argent", Symbol: "XAG"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rouble tadjik", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro", Symbol: "BRR"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florin néerlandais", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra santoméen", Symbol: "STD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "couronne tchèque", Symbol: "CZK"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldave", Symbol: "MDL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutien", Symbol: "DJF"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC", Symbol: "XFU"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unité européenne composée", Symbol: "XBA"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar brunéien", Symbol: "$BN"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorais", Symbol: "TPE"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "livre libanaise", Symbol: "£LB"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rhodésien", Symbol: "$RH"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar soudanais", Symbol: "SDD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dông vietnamien", Symbol: "₫"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africain", Symbol: "ZAR"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unité de compte ADB", Symbol: "XUA"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bouthanais", Symbol: "BTN"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "mark convertible bosniaque", Symbol: "BAM"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricain", Symbol: "CRC"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar bermudien", Symbol: "$BM"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilling tanzanien", Symbol: "TZS"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanien", Symbol: "JOD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "livre syrienne", Symbol: "SYP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unité de compte européenne (ECU)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "livre israélienne", Symbol: "£IL"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothan", Symbol: "lLS"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram arménien", Symbol: "AMD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "roupie des Seychelles", Symbol: "SCR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilling ougandais", Symbol: "UGX"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rouble russe (1991–1998)", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta espagnole", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilien", Symbol: "CLE"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre équatorien", Symbol: "ECS"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanien", Symbol: "MRO"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghize", Symbol: "KGS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar slovène", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar vénézuélien", Symbol: "VEF"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol péruvien", Symbol: "PES"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundais", Symbol: "BIF"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar zimbabwéen", Symbol: "ZWD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "devise inconnue ou non valide", Symbol: "XXX"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "roupie népalaise", Symbol: "NPR"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolais", Symbol: "AOA"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanais", Symbol: "ALL"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivien", Symbol: "BOV"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorien", Symbol: "SVC"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "nouveau manat turkmène", Symbol: "TMT"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar bahaméen", Symbol: "$BS"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panaméen", Symbol: "PAB"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso bissau-guinéen", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar libérien", Symbol: "LRD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambicain", Symbol: "MZE"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "(devise de test)", Symbol: "XTS"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belge (financier)", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgache", Symbol: "MGA"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coréen", Symbol: "₩"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgien", Symbol: "KHR"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nouveau sol péruvien", Symbol: "PEN"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso philippin", Symbol: "PHP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: "LVL"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florin surinamais", Symbol: "SRG"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macédonien", Symbol: "MKD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayen", Symbol: "$UY"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar du Guyana", Symbol: "GYD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar fidjien", Symbol: "$FJ"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonais", Symbol: "JPY"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croate", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libyen", Symbol: "LYD"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling autrichien", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani afghan", Symbol: "AFN"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituanien", Symbol: "LTL"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suisse", Symbol: "CHF"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "roupie srilankaise", Symbol: "LKR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambien", Symbol: "GMD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswanais", Symbol: "BWP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cédi ghanéen", Symbol: "GHS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malais", Symbol: "MYR"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belge", Symbol: "FB"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azéri", Symbol: "AZN"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachme grecque", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbe", Symbol: "RSD"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ancien leu roumain", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisien", Symbol: "TND"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilling somalien", Symbol: "SOS"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "livre maltaise", Symbol: "£MT"}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilling kényan", Symbol: "KES"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonais", Symbol: "PLN"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaïre zaïrois", Symbol: "ZRZ"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubais", Symbol: "AWG"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unité monétaire européenne", Symbol: "XBB"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "roupie maldivienne", Symbol: "MVP"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar trinidadien", Symbol: "$TT"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haïtienne", Symbol: "HTG"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financier luxembourgeois", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurien", Symbol: "HNL"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocain", Symbol: "MAD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guinéen", Symbol: "GNF"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nouveau cruzado", Symbol: "BRN"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar des États-Unis", Symbol: "$US"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "livre chypriote", Symbol: "£CY"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algérien", Symbol: "DZD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "livre de Gibraltar", Symbol: "£GI"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "couronne islandaise", Symbol: "ISK"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosniaque", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australien", Symbol: "$AU"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivien", Symbol: "BOP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinois", Symbol: "CNY"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar des îles Salomon", Symbol: "$SB"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lire maltaise", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuan", Symbol: "VUV"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: "ARA"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "couronne norvégienne", Symbol: "NOK"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilien", Symbol: "$CL"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari géorgien", Symbol: "GEL"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrane", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmène", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgare", Symbol: "BGN"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoan", Symbol: "WS$"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unité d’investissement chilienne", Symbol: "CLF"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba", Symbol: "NIC"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "mark est-allemand", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicain", Symbol: "DOP"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malien", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar yougoslave convertible", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "livre soudanaise", Symbol: "SDG"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croate", Symbol: "HRK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicain", Symbol: "$MX"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "métical", Symbol: "MZM"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty (1950–1995)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cédi", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint hongrois", Symbol: "HUF"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadien", Symbol: "$CA"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgare (1962–1999)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta espagnole (compte convertible)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakien", Symbol: "IQD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rouble biélorusse", Symbol: "BYR"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar néo-zélandais", Symbol: "$NZ"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorien", Symbol: "KMF"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum ouzbek", Symbol: "UZS"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "couronne estonienne", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "livre égyptienne", Symbol: "EGP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "mark finlandais", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar du Yémen", Symbol: "YDD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr éthiopien", Symbol: "ETB"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "type de fonds RINET", Symbol: "XRE"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "livre turque (1844–2005)", Symbol: "TRL"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: "AFA"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rouble soviétique", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saoudien", Symbol: "SAR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "couronne suédoise", Symbol: "SEK"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "livre turque", Symbol: "TRY"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "roupie indonésienne", Symbol: "IDR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatari", Symbol: "QAR"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxembourgeois", Symbol: "F"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocain", Symbol: "fMA"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti péruvien", Symbol: "PEI"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar de Singapour", Symbol: "$SG"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "couronne danoise", Symbol: "DKK"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "livre irlandaise", Symbol: "£IE"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc français", Symbol: "FRF"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: "BUK"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubain", Symbol: "CUP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip loatien", Symbol: "LAK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "livre sud-soudanaise", Symbol: "SSP"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "roupie pakistanaise", Symbol: "PKR"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanais", Symbol: "OMR"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambien", Symbol: "ZMW"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guinéen", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar de Hong Kong", Symbol: "HKD"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranien", Symbol: "IRR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar yougoslave Noviy", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierra-léonais", Symbol: "SLL"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxembourgeois", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belge (convertible)", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florin antillais", Symbol: "ANG"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgache", Symbol: "Fmg"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinée portugaise", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar koweïtien", Symbol: "KWD"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rouble letton", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombien", Symbol: "$CO"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "nouveau dinar yougoslave", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yéménite", Symbol: "YER"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta espagnole (compte A)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "couronne slovaque", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "roupie indienne", Symbol: "₹"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "livre des îles Malouines", Symbol: "£FK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca macanaise", Symbol: "MOP"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "droit de tirage spécial", Symbol: "DTS"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituanien", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawite", Symbol: "MWK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibien", Symbol: "$NA"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigérian", Symbol: "NGN"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "réal brésilien", Symbol: "R$"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capverdien", Symbol: "CVE"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat myanmarais", Symbol: "MMK"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "livre sterling", Symbol: "£GB"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nouveau shekel israélien", Symbol: "₪"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nafka érythréen", Symbol: "ERN"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thaïlandais", Symbol: "THB"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: "CDF"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platine", Symbol: "XPT"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rouble russe", Symbol: "RUB"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivien", Symbol: "BOB"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladeshi", Symbol: "BDT"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicain", Symbol: "MZN"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongan", Symbol: "TOP"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"}}

package fr_GN

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: "BUK"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubain", Symbol: "CUP"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "couronne slovaque", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar du Guyana", Symbol: "GYD"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocain", Symbol: "fMA"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coréen", Symbol: "KPW"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar barbadien", Symbol: "BBD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "nouveau manat turkmène", Symbol: "TMT"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "livre israélienne", Symbol: "£IL"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawite", Symbol: "MWK"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rhodésien", Symbol: "$RH"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financier luxembourgeois", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "livre des îles Malouines", Symbol: "£FK"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba", Symbol: "NIC"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar jamaïcain", Symbol: "JMD"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argent", Symbol: "XAG"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinois", Symbol: "CNY"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australien", Symbol: "$AU"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar brunéien", Symbol: "$BN"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanien", Symbol: "JOD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre équatorien", Symbol: "ECS"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unité monétaire européenne", Symbol: "XBB"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bouthanais", Symbol: "BTN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswanais", Symbol: "BWP"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florin néerlandais", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanais", Symbol: "ALL"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum ouzbek", Symbol: "UZS"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saoudien", Symbol: "SAR"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgache", Symbol: "MGA"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolais", Symbol: "AOA"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadien", Symbol: "$CA"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar de Hong Kong", Symbol: "HKD"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cédi ghanéen", Symbol: "GHS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unité de compte européenne (ECU)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croate", Symbol: "HRK"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libyen", Symbol: "LYD"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guinéen", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belge (financier)", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar trinidadien", Symbol: "$TT"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmène", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatari", Symbol: "QAR"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgien", Symbol: "KHR"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundais", Symbol: "BIF"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "réal brésilien", Symbol: "R$"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "roupie maldivienne", Symbol: "MVP"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nouveau cruzado", Symbol: "BRN"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta espagnole (compte A)", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombien", Symbol: "$CO"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgache", Symbol: "Fmg"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar bélizéen", Symbol: "$BZ"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivien", Symbol: "BOV"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilling somalien", Symbol: "SOS"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricain", Symbol: "CRC"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malais", Symbol: "MYR"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar fidjien", Symbol: "$FJ"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suisse", Symbol: "CHF"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "mark allemand", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbe", Symbol: "RSD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "roupie népalaise", Symbol: "NPR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "roupie des Seychelles", Symbol: "SCR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayen", Symbol: "$UY"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "FCFP"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro", Symbol: "BRR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilling ougandais", Symbol: "UGX"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituanien", Symbol: "LTL"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thaïlandais", Symbol: "THB"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonais", Symbol: "JPY"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr éthiopien", Symbol: "ETB"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dông vietnamien", Symbol: "₫"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "droit de tirage spécial", Symbol: "DTS"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorais", Symbol: "TPE"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanien", Symbol: "MRO"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "couronne forte tchécoslovaque", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "nouveau dinar yougoslave", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: "ARA"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurien", Symbol: "HNL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malien", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "couronne suédoise", Symbol: "SEK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu roumain", Symbol: "RON"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxembourgeois", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierra-léonais", Symbol: "SLL"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cédi", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "livre soudanaise", Symbol: "SDG"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram arménien", Symbol: "AMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "livre libanaise", Symbol: "£LB"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "roupie srilankaise", Symbol: "LKR"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rouble tadjik", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar néo-zélandais", Symbol: "$NZ"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgare (1962–1999)", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algérien", Symbol: "DZD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unité d’investissement chilienne", Symbol: "CLF"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonais", Symbol: "PLN"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "devise inconnue ou non valide", Symbol: "XXX"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranien", Symbol: "IRR"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: "AFA"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc français", Symbol: "F"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxembourgeois", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisien", Symbol: "TND"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "type de fonds RINET", Symbol: "XRE"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC", Symbol: "XFU"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling autrichien", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "roupie pakistanaise", Symbol: "PKR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capverdien", Symbol: "CVE"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "livre chypriote", Symbol: "£CY"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florin surinamais", Symbol: "SRG"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosniaque", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "couronne islandaise", Symbol: "ISK"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "livre de Gibraltar", Symbol: "£GI"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nouveau shekel israélien", Symbol: "₪"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgare", Symbol: "BGN"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilien", Symbol: "CLE"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayen", Symbol: "PYG"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "mark finlandais", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lire maltaise", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: "CDF"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar yougoslave Noviy", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakien", Symbol: "IQD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "couronne danoise", Symbol: "DKK"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat myanmarais", Symbol: "MMK"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nafka érythréen", Symbol: "ERN"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocain", Symbol: "MAD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani afghan", Symbol: "AFN"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinamais", Symbol: "$SR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "livre syrienne", Symbol: "SYP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso philippin", Symbol: "PHP"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "livre sud-soudanaise", Symbol: "SSP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guinéen", Symbol: "FG"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrane", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ancien leu roumain", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panaméen", Symbol: "PAB"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint hongrois", Symbol: "HUF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituanien", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "roupie mauricienne", Symbol: "MUR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilling kényan", Symbol: "KES"}, "INR": ut.Currency{Currency: "INR", DisplayName: "roupie indienne", Symbol: "₹"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari géorgien", Symbol: "GEL"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "couronne estonienne", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "couronne norvégienne", Symbol: "NOK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rouble russe", Symbol: "RUB"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorien", Symbol: "SVC"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar yougoslave convertible", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azéri", Symbol: "AZN"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip loatien", Symbol: "LAK"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belge (convertible)", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol péruvien", Symbol: "PES"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rouble biélorusse", Symbol: "BYR"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldivien", Symbol: "MVR"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambien", Symbol: "GMD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty (1950–1995)", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar des îles Caïmans", Symbol: "KYD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilien", Symbol: "$CL"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macédonien", Symbol: "MKD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambicain", Symbol: "MZE"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belge", Symbol: "FB"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or", Symbol: "XFO"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaïre zaïrois", Symbol: "ZRZ"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unité européenne composée", Symbol: "XBA"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar du Yémen", Symbol: "YDD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubain convertible", Symbol: "CUC"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: "$AR"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar vénézuélien", Symbol: "VEF"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambien", Symbol: "ZMW"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicain", Symbol: "MZN"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothan", Symbol: "lLS"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar libérien", Symbol: "LRD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar soudanais", Symbol: "SDD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanais", Symbol: "OMR"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "livre maltaise", Symbol: "£MT"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lire italienne", Symbol: "₤IT"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "livre égyptienne", Symbol: "EGP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: "LVL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutien", Symbol: "DJF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorien", Symbol: "KMF"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuan", Symbol: "VUV"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar de Singapour", Symbol: "$SG"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "métical", Symbol: "MZM"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "(devise de test)", Symbol: "XTS"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongan", Symbol: "TOP"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreïni", Symbol: "BHD"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "livre turque (1844–2005)", Symbol: "TRL"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "livre irlandaise", Symbol: "£IE"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghize", Symbol: "KGS"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivien", Symbol: "BOP"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca macanaise", Symbol: "MOP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coréen", Symbol: "₩"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicain", Symbol: "$MX"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar des États-Unis", Symbol: "$US"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unité de compte ADB", Symbol: "XUA"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar koweïtien", Symbol: "KWD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti péruvien", Symbol: "PEI"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar bermudien", Symbol: "$BM"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivien", Symbol: "BOB"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haïtienne", Symbol: "HTG"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "mark est-allemand", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "mark convertible bosniaque", Symbol: "BAM"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso bissau-guinéen", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubais", Symbol: "AWG"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandais", Symbol: "RWF"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoan", Symbol: "WS$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nouveau sol péruvien", Symbol: "PEN"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachme grecque", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florin antillais", Symbol: "ANG"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigérian", Symbol: "NGN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibien", Symbol: "$NA"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar zimbabwéen", Symbol: "ZWD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar bahaméen", Symbol: "$BS"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croate", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilling tanzanien", Symbol: "TZS"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "roupie indonésienne", Symbol: "IDR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rouble soviétique", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta espagnole (compte convertible)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rouble russe (1991–1998)", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra santoméen", Symbol: "STD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "livre turque", Symbol: "TRY"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rouble letton", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ukrainienne", Symbol: "UAH"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldave", Symbol: "MDL"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "coupon de lari géorgien", Symbol: "GEK"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar slovène", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platine", Symbol: "XPT"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar des îles Salomon", Symbol: "$SB"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yéménite", Symbol: "YER"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta espagnole", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbo-monténégrin", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinée portugaise", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africain", Symbol: "ZAR"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "livre sterling", Symbol: "£GB"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladeshi", Symbol: "BDT"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicain", Symbol: "DOP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "couronne tchèque", Symbol: "CZK"}}

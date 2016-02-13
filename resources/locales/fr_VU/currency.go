package fr_VU

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GHC": ut.Currency{Currency: "GHC", DisplayName: "cédi", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saoudien", Symbol: "SAR"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinée portugaise", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suisse", Symbol: "CHF"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreïni", Symbol: "BHD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "livre libanaise", Symbol: "£LB"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorais", Symbol: "TPE"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "mark est-allemand", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatari", Symbol: "QAR"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "livre de Gibraltar", Symbol: "£GI"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: "ARA"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar yougoslave convertible", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso bissau-guinéen", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbo-monténégrin", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haïtienne", Symbol: "HTG"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belge", Symbol: "FB"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandais", Symbol: "RWF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libyen", Symbol: "LYD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgien", Symbol: "KHR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutien", Symbol: "DJF"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "livre soudanaise", Symbol: "SDG"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswanais", Symbol: "BWP"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platine", Symbol: "XPT"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilien", Symbol: "CLE"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "roupie srilankaise", Symbol: "LKR"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorien", Symbol: "KMF"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "couronne estonienne", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbe", Symbol: "RSD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr éthiopien", Symbol: "ETB"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar vénézuélien", Symbol: "VEF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "livre chypriote", Symbol: "£CY"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilling ougandais", Symbol: "UGX"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "couronne islandaise", Symbol: "ISK"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta espagnole (compte convertible)", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "couronne norvégienne", Symbol: "NOK"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar trinidadien", Symbol: "$TT"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "nouveau manat turkmène", Symbol: "TMT"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivien", Symbol: "BOP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coréen", Symbol: "KPW"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigérian", Symbol: "NGN"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florin surinamais", Symbol: "SRG"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: "LVL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra santoméen", Symbol: "STD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari géorgien", Symbol: "GEL"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldave", Symbol: "MDL"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: "CDF"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belge (financier)", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panaméen", Symbol: "PAB"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malais", Symbol: "MYR"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubain convertible", Symbol: "CUC"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladeshi", Symbol: "BDT"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unité de compte ADB", Symbol: "XUA"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ukrainienne", Symbol: "UAH"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxembourgeois", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar slovène", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capverdien", Symbol: "CVE"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "roupie indonésienne", Symbol: "IDR"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "livre irlandaise", Symbol: "£IE"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "couronne forte tchécoslovaque", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nouveau sol péruvien", Symbol: "PEN"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty (1950–1995)", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti péruvien", Symbol: "PEI"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar des îles Caïmans", Symbol: "KYD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "roupie mauricienne", Symbol: "MUR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rouble soviétique", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar yougoslave Noviy", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambien", Symbol: "ZMW"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivien", Symbol: "BOV"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrane", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africain", Symbol: "ZAR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanien", Symbol: "JOD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar bahaméen", Symbol: "$BS"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bouthanais", Symbol: "BTN"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachme grecque", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoan", Symbol: "WS$"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nouveau cruzado", Symbol: "BRN"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre équatorien", Symbol: "ECS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "roupie indienne", Symbol: "₹"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "couronne danoise", Symbol: "DKK"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmène", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldivien", Symbol: "MVR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "FCFP"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "livre sud-soudanaise", Symbol: "SSP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu roumain", Symbol: "RON"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat myanmarais", Symbol: "MMK"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituanien", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unité de compte européenne (ECU)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "coupon de lari géorgien", Symbol: "GEK"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar zimbabwéen", Symbol: "ZWD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar du Yémen", Symbol: "YDD"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivien", Symbol: "BOB"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricain", Symbol: "CRC"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint hongrois", Symbol: "HUF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florin néerlandais", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgare", Symbol: "BGN"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro", Symbol: "BRR"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubain", Symbol: "CUP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unité monétaire européenne", Symbol: "XBB"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso philippin", Symbol: "PHP"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financier luxembourgeois", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorien", Symbol: "SVC"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar du Guyana", Symbol: "GYD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "nouveau dinar yougoslave", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "mark allemand", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "livre israélienne", Symbol: "£IL"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rouble letton", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani afghan", Symbol: "AFN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibien", Symbol: "$NA"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar jamaïcain", Symbol: "JMD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar brunéien", Symbol: "$BN"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: "AFA"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar de Hong Kong", Symbol: "HKD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca macanaise", Symbol: "MOP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ancien leu roumain", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar bélizéen", Symbol: "$BZ"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "roupie pakistanaise", Symbol: "PKR"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algérien", Symbol: "DZD"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar néo-zélandais", Symbol: "$NZ"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicain", Symbol: "$MX"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "(devise de test)", Symbol: "XTS"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol péruvien", Symbol: "PES"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belge (convertible)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or", Symbol: "XFO"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanais", Symbol: "OMR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituanien", Symbol: "LTL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonais", Symbol: "PLN"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurien", Symbol: "HNL"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghize", Symbol: "KGS"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilien", Symbol: "$CL"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rhodésien", Symbol: "$RH"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakien", Symbol: "IQD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "livre maltaise", Symbol: "£MT"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuan", Symbol: "VT"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicain", Symbol: "MZN"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croate", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macédonien", Symbol: "MKD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "mark finlandais", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thaïlandais", Symbol: "THB"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guinéen", Symbol: "GNF"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombien", Symbol: "$CO"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "livre turque", Symbol: "TRY"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "roupie des Seychelles", Symbol: "SCR"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar soudanais", Symbol: "SDD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar bermudien", Symbol: "$BM"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongan", Symbol: "TOP"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "devise inconnue ou non valide", Symbol: "XXX"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rouble tadjik", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum ouzbek", Symbol: "UZS"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonais", Symbol: "JPY"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: "$AR"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawite", Symbol: "MWK"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lire italienne", Symbol: "₤IT"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croate", Symbol: "HRK"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgache", Symbol: "MGA"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azéri", Symbol: "AZN"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nouveau shekel israélien", Symbol: "₪"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "réal brésilien", Symbol: "R$"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rouble russe (1991–1998)", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unité européenne composée", Symbol: "XBA"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayen", Symbol: "PYG"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: "BUK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinamais", Symbol: "$SR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba", Symbol: "NIC"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadien", Symbol: "$CA"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dông vietnamien", Symbol: "₫"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgare (1962–1999)", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC", Symbol: "XFU"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "livre égyptienne", Symbol: "EGP"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubais", Symbol: "AWG"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "livre turque (1844–2005)", Symbol: "TRL"}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilling kényan", Symbol: "KES"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "mark convertible bosniaque", Symbol: "BAM"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rouble russe", Symbol: "RUB"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinois", Symbol: "CNY"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta espagnole", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundais", Symbol: "BIF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australien", Symbol: "$AU"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florin antillais", Symbol: "ANG"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar barbadien", Symbol: "BBD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanais", Symbol: "ALL"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "roupie népalaise", Symbol: "NPR"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranien", Symbol: "IRR"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "droit de tirage spécial", Symbol: "DTS"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxembourgeois", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaïre zaïrois", Symbol: "ZRZ"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambien", Symbol: "GMD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram arménien", Symbol: "AMD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "livre sterling", Symbol: "£GB"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rouble biélorusse", Symbol: "BYR"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lire maltaise", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling autrichien", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosniaque", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "type de fonds RINET", Symbol: "XRE"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar libérien", Symbol: "LRD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicain", Symbol: "DOP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guinéen", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "couronne tchèque", Symbol: "CZK"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilling somalien", Symbol: "SOS"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "roupie maldivienne", Symbol: "MVP"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argent", Symbol: "XAG"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolais", Symbol: "AOA"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgache", Symbol: "Fmg"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip loatien", Symbol: "LAK"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malien", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothan", Symbol: "lLS"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocain", Symbol: "MAD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "couronne suédoise", Symbol: "SEK"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambicain", Symbol: "MZE"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yéménite", Symbol: "YER"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nafka érythréen", Symbol: "ERN"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar koweïtien", Symbol: "KWD"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar de Singapour", Symbol: "$SG"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "couronne slovaque", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocain", Symbol: "fMA"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc français", Symbol: "F"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar des États-Unis", Symbol: "$US"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta espagnole (compte A)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "livre syrienne", Symbol: "SYP"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar des îles Salomon", Symbol: "$SB"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "métical", Symbol: "MZM"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilling tanzanien", Symbol: "TZS"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coréen", Symbol: "₩"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayen", Symbol: "$UY"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar fidjien", Symbol: "$FJ"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisien", Symbol: "TND"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unité d’investissement chilienne", Symbol: "CLF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cédi ghanéen", Symbol: "GHS"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanien", Symbol: "MRO"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "livre des îles Malouines", Symbol: "£FK"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierra-léonais", Symbol: "SLL"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"}}

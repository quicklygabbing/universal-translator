package sl

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BRE": ut.Currency{Currency: "BRE", DisplayName: "stari brazilski kruzeiro (1990–1993)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "gvajanski dolar", Symbol: "GYD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haitski gurd", Symbol: "HTG"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norveška krona", Symbol: "NOK"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "frank UIC", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "stari afganistanski afgani (1927–2002)", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "bolivijski peso", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongoški frank", Symbol: "CDF"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "čilski unidades de fomento", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "libijski dinar", Symbol: "LYD"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "novozelandski dolar", Symbol: "NZ$"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "švedska krona", Symbol: "SEK"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "nizozemsko-antilski gulden", Symbol: "ANG"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "argentinski avstral", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "ciprski funt", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "stari urugvajski peso (1975–1993)", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "brazilski novi kruzeiro (1967–1986)", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "gvinejski frank", Symbol: "GNF"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "grška drahma", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "irski funt", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "mehiški srebrni peso (1861–1992)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepalska rupija", Symbol: "NPR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paragvajski gvarani", Symbol: "PYG"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "surinamski gulden", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "saotomejska dobra", Symbol: "STD"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelski bolivar", Symbol: "VEF"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "zambijska kvača (1968–2012)", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ekvadorski sukre", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordanski dinar", Symbol: "JOD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "poljski novi zlot", Symbol: "PLN"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "urugvajski peso", Symbol: "UYU"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundski frank", Symbol: "BIF"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "finska marka", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "islandska krona", Symbol: "ISK"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "malteška lira", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "mozambiški eskudo", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "stari mozambiški metikal", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laoški kip", Symbol: "LAK"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tongovska paanga", Symbol: "TOP"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "venezuelski bolivar (1871–2008)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "novi brazilski kruzado", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "malgaški ariarij", Symbol: "MGA"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrajinska grivna", Symbol: "UAH"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frank", Symbol: "CFPF"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahranski dinar", Symbol: "BHD"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "španska pezeta (račun B)", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "svazijski lilangeni", Symbol: "SZL"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "evropska obračunska enota (XBC)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "argentinski peso (1983–1985)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamski dolar", Symbol: "BSD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "beloruski rubelj", Symbol: "BYR"}, "COP": ut.Currency{Currency: "COP", DisplayName: "kolumbijski peso", Symbol: "COP"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigerijska naira", Symbol: "NGN"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "tadžikistanski rubelj", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandski šiling", Symbol: "UGX"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani", Symbol: "AFN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadski dolar", Symbol: "CAD"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "severnokorejski von", Symbol: "KPW"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "latvijski rubelj", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "omanski rial", Symbol: "OMR"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "stari jugoslovanski dinar", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "stara angolska kvanza (1977–1990)", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "mehiška inverzna enota (UDI)", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "evropska monetarna enota", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zairski novi zaire", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritrejska nakfa", Symbol: "ERN"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "jemenski dinar", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "brazilski kruzado", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "čilski peso", Symbol: "CLP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambijski dalasi", Symbol: "GMD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japonski jen", Symbol: "¥"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namibijski dolar", Symbol: "NAD"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "novi jugoslovanski dinar", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zairski zaire", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "angolska nova kvanza (1990–2000)", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "danska krona", Symbol: "DKK"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiopski bir", Symbol: "ETB"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladij", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "burmanski kjat", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "zelenortski eskudo", Symbol: "CVE"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "nemška marka", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "evro", Symbol: "€"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "stari ganski cedi (1979–2007)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "malgaški frank", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "sejšelska rupija", Symbol: "SCR"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "andorska peseta", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "estonska krona", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ganski cedi", Symbol: "GHS"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "neznana valuta", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "brunejski dolar", Symbol: "BND"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "italijanska lira", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "brazilski kruzeiro", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "perujski novi sol", Symbol: "PEN"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "stari sudanski dinar", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "armenski dram", Symbol: "AMD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolivijski boliviano", Symbol: "BOB"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "evro WIR", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "ameriški dolar", Symbol: "$"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "stari azerbajdžanski manat (1993–2006)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltarski funt", Symbol: "GIP"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "honduraška lempira", Symbol: "HNL"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamajški dolar", Symbol: "JMD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "makedonski denar", Symbol: "MKD"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "portugalski eskudo", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "somalski šiling", Symbol: "SOS"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "španska pezeta", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "gvinejski sili", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "stara turška lira", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dolar Trinidada in Tobaga", Symbol: "TTD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "iraški dinar", Symbol: "IQD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litovski litas", Symbol: "LTL"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mavricijska rupija", Symbol: "MUR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatujski vatu", Symbol: "VUV"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "jugoslovanski konvertibilni dinar", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentinski peso", Symbol: "ARS"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "stari bolgarski lev", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "stari srbski dinar", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "češkoslovaška krona", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "stari romunski leu", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "novi tajvanski dolar", Symbol: "NT$"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "belgijski konvertibilni frank", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubanski peso", Symbol: "CUP"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "luksemburški konvertibilni frank", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "romunski leu", Symbol: "RON"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fidžijski dolar", Symbol: "FJD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "hrvaški dinar", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "madžarski forint", Symbol: "HUF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "liberijski dolar", Symbol: "LRD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "tajski baht", Symbol: "฿"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tadžikistanski somoni", Symbol: "TJS"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadoški dolar", Symbol: "BBD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenijski šiling", Symbol: "KES"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nikaraška zlata kordova", Symbol: "NIO"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ruandski frank", Symbol: "RWF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "funt Sv. Helene", Symbol: "SHP"}, "VND": ut.Currency{Currency: "VND", DisplayName: "vientnamski dong", Symbol: "₫"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "koda za potrebe testiranja", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "gruzijski bon lari", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "lesoški loti", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "maroški frank", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malezijski ringit", Symbol: "MYR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "solomonski dolar", Symbol: "SBD"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sudanski funt", Symbol: "SDG"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "stari sudanski funt", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "belgijski finančni frank", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "indijska rupija", Symbol: "₹"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "filipinski peso", Symbol: "PHP"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "stari poljski zlot (1950–1995)", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rodezijski dolar", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "srebro", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "belgijski frank", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "nizozemski gulden", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "kajmanski dolar", Symbol: "KYD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "južnosudanski funt", Symbol: "SSP"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "bocvanska pula", Symbol: "BWP"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgiški som", Symbol: "KGS"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "luksemburški finančni frank", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mozambiški metikal", Symbol: "MZN"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "salvadorski kolon", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "uzbeški sum", Symbol: "UZS"}, "WST": ut.Currency{Currency: "WST", DisplayName: "samoanska tala", Symbol: "WST"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kitajski juan", Symbol: "CN¥"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "kostariški kolon", Symbol: "CRC"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mavretanska uguija", Symbol: "MRO"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "srbski dinar", Symbol: "RSD"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "evropska obračunska enota (XBD)", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele Ekvatorialne Gvineje", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonezijska rupija", Symbol: "IDR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "perujski sol", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falklandski funt", Symbol: "FKP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "gruzijski lari", Symbol: "GEL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranski rial", Symbol: "IRR"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kamboški riel", Symbol: "KHR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "saudski rial", Symbol: "SAR"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA frank BEAC", Symbol: "FCFA"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "zahodnoafriški frank CFA", Symbol: "CFA"}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemenski rial", Symbol: "YER"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "avstralski dolar", Symbol: "A$"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egiptovski funt", Symbol: "EGP"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "španska pezeta (račun A)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brazilski real", Symbol: "R$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "izraelski šekel", Symbol: "₪"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "surinamski dolar", Symbol: "SRD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "ameriški dolar, isti dan", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham Združenih arabskih emiratov", Symbol: "AED"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "izraelski funt", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanonski funt", Symbol: "LBP"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ruski rubelj", Symbol: "RUB"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bolgarski lev", Symbol: "BGN"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "luksemburški frank", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "zimbabvejski dolar", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudski dolar", Symbol: "BMD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "švicarski frank", Symbol: "CHF"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "džibutski frank", Symbol: "DJF"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "eskudo Portugalske Gvineje", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malavijska kvača", Symbol: "MWK"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "sieraleonski leone", Symbol: "SLL"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "južnoafriški rand", Symbol: "ZAR"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "bosansko-hercegovski dinar", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "hrvaška kuna", Symbol: "HRK"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "komorski frank", Symbol: "KMF"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "šrilanška rupija", Symbol: "LKR"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambijska kvača", Symbol: "ZMW"}, "COU": ut.Currency{Currency: "COU", DisplayName: "kolumbijska enota realne vrednosti", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikanski peso", Symbol: "DOP"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "litvanski litas", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "maroški dirham", Symbol: "MAD"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "vzhodnonemška marka", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "malteški funt", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ukrajinski karbovanci", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "ameriški dolar, naslednji dan", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "evropska sestavljena enota", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "azerbajdžanski manat", Symbol: "AZN"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "češka krona", Symbol: "CZK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "katarski rial", Symbol: "QAR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzanijski šiling", Symbol: "TZS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "evropska denarna enota", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "beloruski novi rubelj (1994–1999)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso Gvineje Bissau", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "ruski rubelj (1991–1998)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "turkmenistanski novi manat", Symbol: "TMT"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "zimbabvejski dolar (2009)", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongolski tugrik", Symbol: "MNT"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makavska pataka", Symbol: "MOP"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "butanski ngultrum", Symbol: "BTN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belizejski dolar", Symbol: "BZD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubanski konvertibilni peso", Symbol: "CUC"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "turkmenski manat", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "posebne pravice črpanja", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ekvadorska enota realne vrednosti (UVC)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "mjanmarski kjat", Symbol: "MMK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "mehiški peso", Symbol: "MX$"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "sovjetski rubelj", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albanski lek", Symbol: "ALL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosansko-hercegovska konvertibilna marka", Symbol: "BAM"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "nova turška lira", Symbol: "TRY"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "britanski funt", Symbol: "£"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "latvijski lats", Symbol: "LVL"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "perujski inti", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "slovenski tolar", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "gvatemalski kecal", Symbol: "GTQ"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kazahstanski tenge", Symbol: "KZT"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "timorski eskudo", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "vzhodnokaribski dolar", Symbol: "EC$"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "zlati frank", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "južnoafriški finančni rand", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "konvertibilna angolska kvanza (1995–1999)", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "francoski frank", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "hongkonški dolar", Symbol: "HK$"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakistanska rupija", Symbol: "PKR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "singapurski dolar", Symbol: "SGD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "bolivijski mvdol", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "alžirski dinar", Symbol: "DZD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "južnokorejski von", Symbol: "₩"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "malijski frank", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldivska rufija", Symbol: "MVR"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "frank WIR", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunizijski dinar", Symbol: "TND"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolska kvanza", Symbol: "AOA"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "sirijski funt", Symbol: "SYP"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "avstrijski šiling", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldavijski leu", Symbol: "MDL"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubski florin", Symbol: "AWG"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina Papue Nove Gvineje", Symbol: "PGK"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuvajtski dinar", Symbol: "KWD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "nikaraška kordova", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladeška taka", Symbol: "BDT"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamska balboa", Symbol: "PAB"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "slovaška krona", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "stari ugandski šiling (1966–1987)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "zlato", Symbol: ""}}

package es_NI

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta española (cuenta convertible)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platino", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "antiguo dinar serbio", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turcomano (1993–2009)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franco CFA BEAC", Symbol: "XAF"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "derechos especiales de giro", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "chelín ugandés", Symbol: "UGX"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasileño", Symbol: "BRL"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugués", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreiní", Symbol: "BHD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florín arubeño", Symbol: "AWG"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "chelín tanzano", Symbol: "TZS"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen", Symbol: "JPY"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "libra siria", Symbol: "SYP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo de Cabo Verde", Symbol: "CVE"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unidad de cuenta europea (XBD)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonesia", Symbol: "IDR"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dólar guyanés", Symbol: "GYD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong", Symbol: "₫"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumano", Symbol: "RON"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azerí (1993–2006)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruceiro brasileño", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruano", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marco convertible de Bosnia-Herzegovina", Symbol: "BAM"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "chelín somalí", Symbol: "SOS"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza reajustado angoleño (1995–1999)", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma griego", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayo", Symbol: "UYU"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rublo ruso (1991–1998)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turcomano", Symbol: "TMT"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chileno", Symbol: "CLP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial catarí", Symbol: "QAR"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "MVDOL boliviano", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidad de fomento chilena", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franco marroquí", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalí", Symbol: "NPR"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unidad monetaria europea", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franco malgache", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondureño", Symbol: "HNL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu", Symbol: "VUV"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadoreño", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franco suizo", Symbol: "CHF"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dólar bahameño", Symbol: "BSD"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unidad de moneda europea", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta española", Symbol: "₧"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dólar estadounidense", Symbol: "$"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi", Symbol: "GHS"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dírham de los Emiratos Árabes Unidos", Symbol: "AED"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dólar rodesiano", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgiano", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliviano", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dólar hongkonés", Symbol: "HKD"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florín surinamés", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libio", Symbol: "LYD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florín neerlandés", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fondos RINET", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca de Macao", Symbol: "MOP"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florín de las Antillas Neerlandesas", Symbol: "ANG"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasileño", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dírham marroquí", Symbol: "MAD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariari", Symbol: "MGA"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia seychellense", Symbol: "SCR"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franco UIC francés", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franco CFP", Symbol: "CFPF"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real colombiana", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringit", Symbol: "MYR"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franco luxemburgués", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kuanza", Symbol: "AOA"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayo en unidades indexadas", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina", Symbol: "PGK"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azerí", Symbol: "AZN"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marco alemán", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iraní", Symbol: "IRR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franco yibutiano", Symbol: "DJF"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosnio", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambiano (1968–2012)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dólar canadiense", Symbol: "CA$"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazako", Symbol: "KZT"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franco malí", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "plata", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladio", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dólar bruneano", Symbol: "BND"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia esrilanquesa", Symbol: "LKR"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franco francés", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tayiko", Symbol: "TJS"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanés", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemení", Symbol: "YER"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "super dinar yugoslavo", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franco congoleño", Symbol: "CDF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unidad compuesta europea", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunecino", Symbol: "TND"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublo ruso", Symbol: "RUB"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba nicaragüense", Symbol: "C$"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dólar estadounidense (día siguiente)", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbio", Symbol: "RSD"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "antiguo metical mozambiqueño", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna", Symbol: "HRK"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltesa", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar yemení", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ostmark de Alemania del Este", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitiano", Symbol: "HTG"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "libra irlandesa", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som", Symbol: "KGS"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dólar zimbabuense", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka", Symbol: "BDT"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona estonia", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar argelino", Symbol: "DZD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar fuerte yugoslavo", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineano", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franco comorense", Symbol: "KMF"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nuevo dólar taiwanés", Symbol: "TWD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip", Symbol: "LAK"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dólar de Trinidad y Tobago", Symbol: "TTD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayo", Symbol: "PYG"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira", Symbol: "NGN"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leona", Symbol: "SLL"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won surcoreano", Symbol: "KRW"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentino", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nuevo cruceiro brasileño (1967–1986)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dólar beliceño", Symbol: "BZD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dólar de las Islas Caimán", Symbol: "KYD"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghanés (1979–2007)", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dólar neozelandés", Symbol: "NZD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "uguiya", Symbol: "MRO"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nuevo zaire zaireño", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekuele de Guinea Ecuatorial", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula", Symbol: "BWP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum", Symbol: "UZS"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franco belga (financiero)", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordano", Symbol: "JOD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costarricense", Symbol: "CRC"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "libra gibraltareña", Symbol: "GIP"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) ecuatoriana", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire zaireño", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni", Symbol: "SZL"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angoleño (1977–1990)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iraquí", Symbol: "IQD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek", Symbol: "ALL"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forinto húngaro", Symbol: "HUF"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nuevo sol peruano", Symbol: "PEN"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruano", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra", Symbol: "STD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dólar namibio", Symbol: "NAD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldavo", Symbol: "MDL"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "libra maltesa", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudí", Symbol: "SAR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical", Symbol: "MZN"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruceiro brasileño (1990–1993)", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franco ruandés", Symbol: "RWF"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dólar salomonense", Symbol: "SBD"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublo bielorruso", Symbol: "BYR"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik", Symbol: "MNT"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nuevo cruzado brasileño", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipino", Symbol: "PHP"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "oro", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram", Symbol: "AMD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistaní", Symbol: "PKR"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta española (cuenta A)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "código reservado para pruebas", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "gultrum", Symbol: "BTN"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kiat", Symbol: "MMK"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franco belga", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre ecuatoriano", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dólar estadounidense (mismo día)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezolano", Symbol: "VEF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dólar jamaicano", Symbol: "JMD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand", Symbol: "ZAR"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "libra británica", Symbol: "GBP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "libra malvinense", Symbol: "FKP"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franco financiero luxemburgués", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "grivna", Symbol: "UAH"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari", Symbol: "GEL"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev búlgaro", Symbol: "BGN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan", Symbol: "CNY"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversión (UDI) mexicana", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiya", Symbol: "MVR"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublo tayiko", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuwaití", Symbol: "KWD"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorense", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "moneda desconocida", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombiano", Symbol: "COP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letón", Symbol: "LVL"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "libra egipcia", Symbol: "EGP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dólar del Caribe Oriental", Symbol: "XCD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "libra sudanesa antigua", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi", Symbol: "GMD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel", Symbol: "KHR"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanet ucraniano", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franco CFA BCEAO", Symbol: "XOF"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolívar venezolano (1871–2008)", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalteco", Symbol: "GTQ"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubano", Symbol: "CUP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona fuerte checoslovaca", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublo soviético", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tólar esloveno", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "chelín ugandés (1966–1987)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentino", Symbol: "ARS"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "libra chipriota", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "chelín austriaco", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won norcoreano", Symbol: "KPW"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dólar de Bermudas", Symbol: "BMD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambiqueño", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituano", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "chelín keniano", Symbol: "KES"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriciana", Symbol: "MUR"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unidad de cuenta europea (XBC)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franco convertible luxemburgués", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franco WIR", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona checa", Symbol: "CZK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "córdoba nicaragüense (1988–1991)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "libra israelí", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dólar fiyiano", Symbol: "FJD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "bat", Symbol: "฿"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dólar de Zimbabue", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malauí", Symbol: "MWK"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dólar barbadense", Symbol: "BBD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franco guineano", Symbol: "GNF"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayo (1975–1993)", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "antiguo leu rumano", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kuacha zambiano", Symbol: "ZMW"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dólar liberiano", Symbol: "LRD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dólar singapurense", Symbol: "SGD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "libra de Santa Elena", Symbol: "SHP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "bir", Symbol: "ETB"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonio", Symbol: "MKD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani", Symbol: "AFN"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franco burundés", Symbol: "BIF"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omaní", Symbol: "OMR"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franco belga (convertible)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty polaco (1950–1995)", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa", Symbol: "ERN"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmano", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "esloti", Symbol: "PLN"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "libra sursudanesa", Symbol: "SSP"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso de Guinea-Bissáu", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franco oro francés", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicano", Symbol: "MXN"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublo letón", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paanga", Symbol: "TOP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dólar australiano", Symbol: "AUD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia india", Symbol: "INR"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marco finlandés", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinea Portuguesa", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubano convertible", Symbol: "CUC"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso de plata mexicano (1861–1992)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar convertible yugoslavo", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentino (1983–1985)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libra libanesa", Symbol: "LBP"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala", Symbol: "WST"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameño", Symbol: "PAB"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona eslovaca", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituano", Symbol: "LTL"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nuevo séquel israelí", Symbol: "ILS"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicano", Symbol: "DOP"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croata", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev fuerte búlgaro", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothense", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dólar surinamés", Symbol: "SRD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sudafricano (financiero)", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "libra sudanesa", Symbol: "SDG"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nuevo rublo bielorruso (1994–1999)", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "nuevo kwanza angoleño (1990–2000)", Symbol: ""}}

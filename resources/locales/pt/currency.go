package pt

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"FRF": ut.Currency{Currency: "FRF", DisplayName: "Franco francês", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar kuwaitiano", Symbol: "KWD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Franco luxemburguês", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra de São Tomé e Príncipe", Symbol: "STD"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Direitos Especiais de Giro", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev búlgaro (1879–1952)", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano (1863–1963)", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Libra sul-sudanesa", Symbol: "SSP"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar bahamense", Symbol: "BSD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Dracma grego", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary malgaxe", Symbol: "MGA"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublo russo", Symbol: "RUB"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolívar venezuelano", Symbol: "VEF"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dólar do Zimbábue (2009)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dólar fijiano", Symbol: "FJD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Cupom Lari georgiano", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dólar guianense", Symbol: "GYD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iraniano", Symbol: "IRR"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franco UIC francês", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar reformado iugoslavo (1992–1993)", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso cubano", Symbol: "CUP"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Coroa estoniana", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar namibiano", Symbol: "NAD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira nigeriana", Symbol: "NGN"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire Novo zairense (1993–1998)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso argentino", Symbol: "ARS"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Unidades de Fomento chilenas", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel cambojano", Symbol: "KHR"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats letão", Symbol: "LVL"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rublo letão", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical de Moçambique (1980–2006)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dólar norte-americano (Dia seguinte)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Ouro", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unidade Composta Europeia", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Peso argentino (1983–1985)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: "ETB"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti peruano", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Xelim tanzaniano", Symbol: "TZS"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Libra egípcia", Symbol: "EGP"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira hondurenha", Symbol: "HNL"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Franco financeiro de Luxemburgo", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso filipino", Symbol: "PHP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dólar do Caribe Oriental", Symbol: "EC$"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: "BOB"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta espanhola (conta A)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won norte-coreano", Symbol: "KPW"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip laosiano", Symbol: "LAK"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi de Gana (1979–2007)", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani paraguaio", Symbol: "PYG"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia ucraniano", Symbol: "UAH"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dólar barbadense", Symbol: "BBD"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro brasileiro (1993–1994)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi ganês", Symbol: "GHS"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA de BCEAO", Symbol: "CFA"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Peso uruguaio en unidades indexadas", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloti polonês (1950–1995)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek albanês", Symbol: "ALL"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Franco belga (financeiro)", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Franco monegasco", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dólar bermudense", Symbol: "BMD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari georgiano", Symbol: "GEL"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguense", Symbol: "NIO"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Dinar sudanês (1992–2007)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA de BEAC", Symbol: "FCFA"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral argentino", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Libra israelita", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia ceilandesa", Symbol: "LKR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Novo dinar da Bósnia-Herzegovina (1994–1997)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real brasileiro", Symbol: "R$"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco djibutiense", Symbol: "DJF"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dólar do Zimbábue (2008)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Florim das Antilhas Holandesas", Symbol: "ANG"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Novo cuanza angolano (1990–2000)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev forte búlgaro", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat mianmarense", Symbol: "MMK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso mexicano", Symbol: "MX$"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu romeno", Symbol: "RON"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar sérvio", Symbol: "RSD"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Florim do Suriname", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala samoano", Symbol: "WST"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Cuacha zambiano (1968–2012)", Symbol: "ZMK"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franco belga (conversível)", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dólar do Banco Popular da China", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Marca finlandesa", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Coroa sueca", Symbol: "SEK"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rublo soviético", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso colombiano", Symbol: "COP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Coroa dinamarquesa", Symbol: "DKK"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Unidade de Conta Europeia (XBC)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Código de Moeda de Teste", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Sheqel antigo israelita", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dólar neozelandês", Symbol: "NZ$"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Libra de Santa Helena", Symbol: "SHP"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiro brasileiro (1942–1967)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso cubano conversível", Symbol: "CUC"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dólar jamaicano", Symbol: "JMD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dólar singapuriano", Symbol: "SGD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong vietnamita", Symbol: "₫"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta espanhola", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa panamenha", Symbol: "PAB"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar Bons esloveno", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar venezuelano (1871–2008)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar iemenita", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbábue (1980–2008)", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Libra cipriota", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint húngaro", Symbol: "HUF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dólar liberiano", Symbol: "LRD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal saudita", Symbol: "SAR"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rublo do Tadjiquistão", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudo chileno", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Novo sol peruano", Symbol: "PEN"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht tailandês", Symbol: "฿"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira turca", Symbol: "TRY"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar noviy iugoslavo (1994–2002)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev socialista búlgaro", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Peso boliviano", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha malawiana", Symbol: "MWK"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Marco alemão", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal guatemalense", Symbol: "GTQ"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco comorense", Symbol: "KMF"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas lituano", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Xelim ugandense", Symbol: "UGX"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Prata", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afegane (1927–2002)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat birmanês", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Iene japonês", Symbol: "JP¥"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit malaio", Symbol: "MYR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical moçambicano", Symbol: "MZN"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanetz ucraniano", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidade de Moeda Europeia", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar jordaniano", Symbol: "JOD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti do Lesoto", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudo de Moçambique", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga tonganesa", Symbol: "TOP"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Cuanza angolano (1977–1990)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde haitiano", Symbol: "HTG"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni tadjique", Symbol: "TJS"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev búlgaro", Symbol: "BGN"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Mvdol boliviano", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franco suíço", Symbol: "CHF"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Libra malvinense", Symbol: "FKP"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Libra de Gibraltar", Symbol: "GIP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirrã marroquino", Symbol: "MAD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franco de Mali", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial catariano", Symbol: "QAR"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong vietnamita (1978–1985)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paládio", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand sul-africano (financeiro)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso chileno", Symbol: "CLP"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar croata", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupon moldávio", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu vanuatuense", Symbol: "VUV"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidade de Valor Constante (UVC) do Equador", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won sul-coreano", Symbol: "₩"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rublo russo (1991–1998)", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni suazi", Symbol: "SZL"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Cuanza angolano reajustado (1995–1999)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado novo brasileiro (1989–1990)", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Coroa norueguesa", Symbol: "NOK"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo português", Symbol: "Esc."}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somaliano", Symbol: "SOS"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudo timorense", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dólar norte-americano (Mesmo dia)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franco-ouro francês", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum butanês", Symbol: "BTN"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa da Eritreia", Symbol: "ERN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franco ruandês", Symbol: "RWF"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadense", Symbol: "CA$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel novo israelita", Symbol: "₪"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Lira maltesa", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franco burundiano", Symbol: "BIF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge cazaque", Symbol: "KZT"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya mauritana", Symbol: "MRO"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol peruano (1863–1965)", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire zairense (1971–1993)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirrã dos Emirados Árabes Unidos", Symbol: "AED"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rublo bielorrusso", Symbol: "BYR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo cabo-verdiano", Symbol: "CVE"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso dominicano", Symbol: "DOP"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Lira italiana", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libra libanesa", Symbol: "LBP"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unidade Monetária Europeia", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial iemenita", Symbol: "YER"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Marco bósnio-herzegovino conversível", Symbol: "BAM"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublo novo bielo-russo (1994–1999)", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Franco marroquino", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia mauriciana", Symbol: "MUR"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rupia maldiva", Symbol: "MVR"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat turcomeno", Symbol: "TMT"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Lira turca (1922–2005)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan chinês", Symbol: "CN¥"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas lituano", Symbol: "LTL"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dólar rodesiano", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Peso argentino (1881–1970)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Coroa tcheca", Symbol: "CZK"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna croata", Symbol: "HRK"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Florim holandês", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Libra sudanesa", Symbol: "SDG"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dólar surinamês", Symbol: "SRD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Libra síria", Symbol: "SYP"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar conversível iugoslavo (1990–1992)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek Albanês (1946–1965)", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar argelino", Symbol: "DZD"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Coroa antiga islandesa", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Libra sudanesa (1957–1998)", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Libra maltesa", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia paquistanesa", Symbol: "PKR"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram armênio", Symbol: "AMD"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta espanhola (conta conversível)", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar líbio", Symbol: "LYD"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Coroa eslovaca", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat azeri", Symbol: "AZN"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre equatoriano", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan da Coreia do Sul (1953–1962)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Franco conversível de Luxemburgo", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Franco de Madagascar", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca macaense", Symbol: "MOP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha zambiano", Symbol: "ZMW"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dólar de Hong Kong", Symbol: "HK$"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Xelim queniano", Symbol: "KES"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloti polonês", Symbol: "PLN"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dólar das Ilhas Salomão", Symbol: "SBD"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dólar de Trinidad e Tobago", Symbol: "TTD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta de Andorra", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Peso lei argentino (1970–1983)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Prata mexicano (1861–1992)", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina papuásia", Symbol: "PGK"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar da Bósnia-Herzegovina (1992–1994)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euro WIR", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele da Guiné Equatorial", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia indiana", Symbol: "₹"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dólar das Ilhas Caiman", Symbol: "KYD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik mongol", Symbol: "MNT"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi gambiano", Symbol: "GMD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial omanense", Symbol: "OMR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia seichelense", Symbol: "SCR"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar sérvio (2002–2006)", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Dinar macedônio (1992–1993)", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso uruguaio", Symbol: "UYU"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidade Mexicana de Investimento (UDI)", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisiano", Symbol: "TND"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som uzbeque", Symbol: "UZS"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Xelim austríaco", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som quirguiz", Symbol: "KGS"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguense (1988–1991)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula botsuanesa", Symbol: "BWP"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franco congolês", Symbol: "CDF"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Coroa Forte checoslovaca", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Ostmark da Alemanha Oriental", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franco guineano", Symbol: "GNF"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli da Guiné", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo da Guiné Portuguesa", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sul-africano", Symbol: "ZAR"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza angolano", Symbol: "AOA"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florim arubano", Symbol: "AWG"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado brasileiro (1986–1989)", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Dinar macedônio", Symbol: "MKD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat do Turcomenistão (1993–2009)", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Fundos RINET", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Moeda desconhecida", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat azerbaijano (1993–2006)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colom salvadorenho", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afegane afegão", Symbol: "AFN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka bengalesa", Symbol: "BDT"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dólar bruneano", Symbol: "BND"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Libra irlandesa", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar iraquiano", Symbol: "IQD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone de Serra Leoa", Symbol: "SLL"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Coroa islandesa", Symbol: "ISK"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar forte iugoslavo (1966–1990)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australiano", Symbol: "AU$"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizenho", Symbol: "BZD"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidade de Valor Real", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón costarriquenho", Symbol: "CRC"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupia indonésia", Symbol: "IDR"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldávio", Symbol: "MDL"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia nepalesa", Symbol: "NPR"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Novo dólar taiwanês", Symbol: "NT$"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Unidade de Conta Europeia (XBD)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Franco CFP", Symbol: "CFPF"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu romeno (1952–2006)", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar americano", Symbol: "US$"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Franco belga", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar bareinita", Symbol: "BHD"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Franco WIR", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra britânica", Symbol: "£"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Peso da Guiné-Bissau", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won da Coreia do Sul (1945–1953)", Symbol: ""}}

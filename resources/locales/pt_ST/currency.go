package pt_ST

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal guatemalense", Symbol: "GTQ"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franco de Mali", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa da Eritreia", Symbol: "ERN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franco ruandês", Symbol: "RWF"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Cuanza angolano reajustado (1995–1999)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Franco conversível de Luxemburgo", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar kuwaitiano", Symbol: "KWD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real brasileiro", Symbol: "R$"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso colombiano", Symbol: "COP"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dólar neozelandês", Symbol: "NZ$"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar venezuelano (1871–2008)", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti do Lesoto", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso chileno", Symbol: "CLP"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Lira turca (1922–2005)", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar americano", Symbol: "US$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Libra malvinense", Symbol: "FKP"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rublo bielorrusso", Symbol: "BYR"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Cupom Lari georgiano", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel cambojano", Symbol: "KHR"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dólar do Caribe Oriental", Symbol: "EC$"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta espanhola", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Peso boliviano", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Lira italiana", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franco congolês", Symbol: "CDF"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Fundos RINET", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi ganês", Symbol: "GHS"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral argentino", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Unidade de Conta Europeia (XBC)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna croata", Symbol: "HRK"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franco guineano", Symbol: "GNF"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta espanhola (conta A)", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia ucraniano", Symbol: "UAH"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Peso uruguaio en unidades indexadas", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Coroa dinamarquesa", Symbol: "DKK"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial iemenita", Symbol: "YER"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo cabo-verdiano", Symbol: "CVE"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublo novo bielo-russo (1994–1999)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Unidade de Conta Europeia (XBD)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical de Moçambique (1980–2006)", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rublo do Tadjiquistão", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Mvdol boliviano", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florim arubano", Symbol: "AWG"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone de Serra Leoa", Symbol: "SLL"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dólar bruneano", Symbol: "BND"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dólar guianense", Symbol: "GYD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar iemenita", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco comorense", Symbol: "KMF"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Coroa sueca", Symbol: "SEK"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afegane (1927–2002)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat birmanês", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas lituano", Symbol: "LTL"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi gambiano", Symbol: "GMD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Coroa islandesa", Symbol: "ISK"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Prata", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudo timorense", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Peso argentino (1881–1970)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Franco de Madagascar", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre equatoriano", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli da Guiné", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar iraquiano", Symbol: "IQD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unidade Composta Europeia", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: "BOB"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Coroa norueguesa", Symbol: "NOK"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Coroa eslovaca", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial omanense", Symbol: "OMR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso uruguaio", Symbol: "UYU"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat azerbaijano (1993–2006)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat mianmarense", Symbol: "MMK"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha zambiano", Symbol: "ZMW"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afegane afegão", Symbol: "AFN"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira hondurenha", Symbol: "HNL"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dólar barbadense", Symbol: "BBD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dólar bermudense", Symbol: "BMD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Libra israelita", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar argelino", Symbol: "DZD"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Coroa antiga islandesa", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Libra irlandesa", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dólar norte-americano (Mesmo dia)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franco-ouro francês", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Dinar macedônio (1992–1993)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbábue (1980–2008)", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanetz ucraniano", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol peruano (1863–1965)", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirrã marroquino", Symbol: "MAD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa panamenha", Symbol: "PAB"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudo de Moçambique", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni tadjique", Symbol: "TJS"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rublo russo (1991–1998)", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Prata mexicano (1861–1992)", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia seichelense", Symbol: "SCR"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado brasileiro (1986–1989)", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Libra sul-sudanesa", Symbol: "SSP"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Cuacha zambiano (1968–2012)", Symbol: "ZMK"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudo chileno", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Xelim ugandense", Symbol: "UGX"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya mauritana", Symbol: "MRO"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Franco marroquino", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar sérvio (2002–2006)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev búlgaro (1879–1952)", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani paraguaio", Symbol: "PYG"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek albanês", Symbol: "ALL"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat turcomeno", Symbol: "TMT"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra de São Tomé e Príncipe", Symbol: "Db"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euro WIR", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira nigeriana", Symbol: "NGN"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Unidades de Fomento chilenas", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Florim das Antilhas Holandesas", Symbol: "ANG"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni suazi", Symbol: "SZL"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won sul-coreano", Symbol: "₩"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso cubano", Symbol: "CUP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip laosiano", Symbol: "LAK"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro brasileiro (1993–1994)", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Iene japonês", Symbol: "JP¥"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Xelim queniano", Symbol: "KES"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo da Guiné Portuguesa", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franco belga (conversível)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Marca finlandesa", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iraniano", Symbol: "IRR"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libra libanesa", Symbol: "LBP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele da Guiné Equatorial", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidade Mexicana de Investimento (UDI)", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Dinar macedônio", Symbol: "MKD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka bengalesa", Symbol: "BDT"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA de BEAC", Symbol: "FCFA"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge cazaque", Symbol: "KZT"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar da Bósnia-Herzegovina (1992–1994)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso argentino", Symbol: "ARS"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Florim do Suriname", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Coroa tcheca", Symbol: "CZK"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta espanhola (conta conversível)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat azeri", Symbol: "AZN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula botsuanesa", Symbol: "BWP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australiano", Symbol: "AU$"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Franco belga", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franco francês", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som uzbeque", Symbol: "UZS"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga tonganesa", Symbol: "TOP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Coroa Forte checoslovaca", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra britânica", Symbol: "£"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Franco belga (financeiro)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso cubano conversível", Symbol: "CUC"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somaliano", Symbol: "SOS"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Peso argentino (1983–1985)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloti polonês (1950–1995)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia paquistanesa", Symbol: "PKR"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón costarriquenho", Symbol: "CRC"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum butanês", Symbol: "BTN"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia mauriciana", Symbol: "MUR"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar conversível iugoslavo (1990–1992)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloti polonês", Symbol: "PLN"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dólar das Ilhas Caiman", Symbol: "KYD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldávio", Symbol: "MDL"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Cuanza angolano (1977–1990)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia nepalesa", Symbol: "NPR"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dólar liberiano", Symbol: "LRD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirrã dos Emirados Árabes Unidos", Symbol: "AED"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Marco bósnio-herzegovino conversível", Symbol: "BAM"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar líbio", Symbol: "LYD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina papuásia", Symbol: "PGK"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira turca", Symbol: "TRY"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidade de Valor Constante (UVC) do Equador", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan da Coreia do Sul (1953–1962)", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano (1863–1963)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dólar fijiano", Symbol: "FJD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: "ETB"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia ceilandesa", Symbol: "LKR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franco suíço", Symbol: "CHF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Florim holandês", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dólar de Hong Kong", Symbol: "HK$"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupia indonésia", Symbol: "IDR"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Direitos Especiais de Giro", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dólar norte-americano (Dia seguinte)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev socialista búlgaro", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Libra sudanesa (1957–1998)", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat do Turcomenistão (1993–2009)", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas lituano", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical moçambicano", Symbol: "MZN"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolívar venezuelano", Symbol: "VEF"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso filipino", Symbol: "PHP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dólar singapuriano", Symbol: "SGD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sul-africano", Symbol: "ZAR"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Moeda desconhecida", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo português", Symbol: "Esc."}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan chinês", Symbol: "CN¥"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublo russo", Symbol: "RUB"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Novo dinar da Bósnia-Herzegovina (1994–1997)", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiro brasileiro (1942–1967)", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dólar jamaicano", Symbol: "JMD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal saudita", Symbol: "SAR"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram armênio", Symbol: "AMD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisiano", Symbol: "TND"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza angolano", Symbol: "AOA"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizenho", Symbol: "BZD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Franco luxemburguês", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi de Gana (1979–2007)", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Lira maltesa", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som quirguiz", Symbol: "KGS"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidade de Valor Real", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Libra egípcia", Symbol: "EGP"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Xelim austríaco", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Dracma grego", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Ouro", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong vietnamita (1978–1985)", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rupia maldiva", Symbol: "MVR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Libra sudanesa", Symbol: "SDG"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Ostmark da Alemanha Oriental", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand sul-africano (financeiro)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar croata", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupon moldávio", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franco burundiano", Symbol: "BIF"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu romeno (1952–2006)", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire Novo zairense (1993–1998)", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Dinar sudanês (1992–2007)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial catariano", Symbol: "QAR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti peruano", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won norte-coreano", Symbol: "KPW"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dólar surinamês", Symbol: "SRD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Código de Moeda de Teste", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Sheqel antigo israelita", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Coroa estoniana", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Novo cuanza angolano (1990–2000)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev forte búlgaro", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong vietnamita", Symbol: "₫"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha malawiana", Symbol: "MWK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Novo sol peruano", Symbol: "PEN"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta de Andorra", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Peso lei argentino (1970–1983)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia indiana", Symbol: "₹"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Franco WIR", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dólar do Zimbábue (2008)", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso mexicano", Symbol: "MX$"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint húngaro", Symbol: "HUF"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colom salvadorenho", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar bahamense", Symbol: "BSD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Marco alemão", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unidade Monetária Europeia", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rublo letão", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Xelim tanzaniano", Symbol: "TZS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar Bons esloveno", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar jordaniano", Symbol: "JOD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguense (1988–1991)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Franco CFP", Symbol: "CFPF"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar reformado iugoslavo (1992–1993)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Franco financeiro de Luxemburgo", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Franco monegasco", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguense", Symbol: "NIO"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Libra de Santa Helena", Symbol: "SHP"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dólar do Zimbábue (2009)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari georgiano", Symbol: "GEL"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dólar do Banco Popular da China", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar noviy iugoslavo (1994–2002)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde haitiano", Symbol: "HTG"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado novo brasileiro (1989–1990)", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Libra síria", Symbol: "SYP"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca macaense", Symbol: "MOP"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won da Coreia do Sul (1945–1953)", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar namibiano", Symbol: "NAD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rublo soviético", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev búlgaro", Symbol: "BGN"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu vanuatuense", Symbol: "VUV"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Libra de Gibraltar", Symbol: "GIP"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek Albanês (1946–1965)", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik mongol", Symbol: "MNT"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar forte iugoslavo (1966–1990)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA de BCEAO", Symbol: "CFA"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco djibutiense", Symbol: "DJF"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu romeno", Symbol: "RON"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidade de Moeda Europeia", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Libra maltesa", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dólar das Ilhas Salomão", Symbol: "SBD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar bareinita", Symbol: "BHD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats letão", Symbol: "LVL"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht tailandês", Symbol: "฿"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel novo israelita", Symbol: "₪"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dólar rodesiano", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dólar de Trinidad e Tobago", Symbol: "TTD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar sérvio", Symbol: "RSD"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Libra cipriota", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso dominicano", Symbol: "DOP"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Peso da Guiné-Bissau", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary malgaxe", Symbol: "MGA"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franco UIC francês", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit malaio", Symbol: "MYR"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paládio", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadense", Symbol: "CA$"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Novo dólar taiwanês", Symbol: "NT$"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala samoano", Symbol: "WST"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire zairense (1971–1993)", Symbol: ""}}

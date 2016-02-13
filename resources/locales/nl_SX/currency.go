package nl_SX

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"NLG": ut.Currency{Currency: "NLG", DisplayName: "Nederlandse gulden", Symbol: "NLG"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinaamse gulden", Symbol: "SRG"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Boliviaanse mvdol", Symbol: "BOV"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatoriaal-Guinese ekwele guineana", Symbol: "GQE"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutaanse ngultrum", Symbol: "BTN"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgische lari", Symbol: "GEL"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Soedanese dinar", Symbol: "SDD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Braziliaanse cruzeiro (1990–1993)", Symbol: "BRE"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivische rufiyaa", Symbol: "MVR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbeidzjaanse manat", Symbol: "AZN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinse dinar", Symbol: "BHD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayaanse peso", Symbol: "UYU"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistaanse roepie", Symbol: "PKR"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagassische franc", Symbol: "MGF"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentijnse austral", Symbol: "ARA"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Wit-Russische roebel", Symbol: "BYR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Sint-Heleens pond", Symbol: "SHP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Verenigde Arabische Emiraten-dirham", Symbol: "AED"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australische dollar", Symbol: "AU$"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaanse paʻanga", Symbol: "TOP"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutiaanse frank", Symbol: "DJF"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franse gouden franc", Symbol: "XFO"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruaanse inti", Symbol: "PEI"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Cubaanse convertibele peso", Symbol: "CUC"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Sloveense tolar", Symbol: "SIT"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kaapverdische escudo", Symbol: "CVE"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poolse zloty", Symbol: "PLN"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA-franc BCEAO", Symbol: "CFA"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese frank", Symbol: "CDF"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litouwse talonas", Symbol: "LTT"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadoraanse sucre", Symbol: "ECS"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji-dollar", Symbol: "FJ$"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Braziliaanse cruzeiro (1942–1967)", Symbol: "BRZ"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Deense kroon", Symbol: "DKK"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuaanse vatu", Symbol: "VUV"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Oegandese shilling", Symbol: "UGX"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaïrese zaïre", Symbol: "ZRZ"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundese frank", Symbol: "BIF"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezolaanse bolivar", Symbol: "VEF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmarese kyat", Symbol: "MMK"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptisch pond", Symbol: "EGP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinese syli", Symbol: "GNS"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotiaanse kip", Symbol: "LAK"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Oude Servische dinar", Symbol: "CSD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominicaanse peso", Symbol: "DOP"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkaanse franc", Symbol: "MAF"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadzjiekse somoni", Symbol: "TJS"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luxemburgse financiële franc", Symbol: "LUL"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Koeweitse dinar", Symbol: "KWD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemenitische dinar", Symbol: "YDD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Valutacode voor testdoeleinden", Symbol: "XTS"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Oude Zuid-Koreaanse won (1945–1953)", Symbol: "KRO"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papoea-Nieuw-Guinese kina", Symbol: "PGK"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Oezbeekse sum", Symbol: "UZS"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bengalese taka", Symbol: "BDT"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estlandse kroon", Symbol: "EEK"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaanse ouguiya", Symbol: "MRO"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Europese monetaire eenheid", Symbol: "XBB"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Poolse zloty (1950–1995)", Symbol: "PLZ"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgizische som", Symbol: "KGS"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litouwse litas", Symbol: "LTL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatarese rial", Symbol: "QAR"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Oostenrijkse schilling", Symbol: "ATS"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkse lira", Symbol: "TRY"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israëlische nieuwe shekel", Symbol: "₪"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Zuid-Koreaanse hwan (1953–1962)", Symbol: "KRH"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Goud", Symbol: "XAU"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belizaanse dollar", Symbol: "BZD"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Boliviaanse peso", Symbol: "BOP"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnische dinar", Symbol: "BAD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Soedanees pond (1957–1998)", Symbol: "SDP"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinese frank", Symbol: "GNF"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadoraanse colón", Symbol: "SVC"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambikaanse metical", Symbol: "MZN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopische birr", Symbol: "ETB"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraakse dinar", Symbol: "IQD"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentijnse peso", Symbol: "ARS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiase roepie", Symbol: "₹"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaanse dirham", Symbol: "MAD"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldivische roepie", Symbol: "MVP"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Duitse mark", Symbol: "DEM"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwaanse dollar (2008)", Symbol: "ZWR"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviaanse boliviano", Symbol: "BOB"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luxemburgse convertibele franc", Symbol: "LUC"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyaanse dollar", Symbol: "GYD"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Braziliaanse cruzado novo", Symbol: "BRN"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymaneilandse dollar", Symbol: "KYD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwaanse dollar (2009)", Symbol: "ZWL"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiaanse roepie", Symbol: "MUR"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwaanse dollar", Symbol: "ZWD"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghani (1927–2002)", Symbol: "AFA"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugees-Guinese escudo", Symbol: "GWE"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazachse tenge", Symbol: "KZT"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrisch pond", Symbol: "SYP"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychelse roepie", Symbol: "SCR"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Braziliaanse real", Symbol: "R$"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Zuid-Koreaanse won", Symbol: "₩"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipijnse peso", Symbol: "PHP"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaïrese nieuwe zaïre", Symbol: "ZRN"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorrese peseta", Symbol: "ADP"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldavische leu", Symbol: "MDL"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spaanse peseta", Symbol: "ESP"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-dollar", Symbol: "BMD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Maleisische ringgit", Symbol: "MYR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-frank", Symbol: "XPF"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinee-Bissause peso", Symbol: "GWP"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgische frank", Symbol: "BEF"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hongaarse forint", Symbol: "HUF"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgische kupon larit", Symbol: "GEK"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saoedi-Arabische riyal", Symbol: "SAR"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentijnse peso (1983–1985)", Symbol: "ARP"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolese kwanza", Symbol: "AOA"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatische kuna", Symbol: "HRK"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamese balboa", Symbol: "PAB"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombiaanse peso", Symbol: "COP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Brits pond", Symbol: "£"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "IJslandse kroon", Symbol: "ISK"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugese escudo", Symbol: "PTE"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Santomese dobra", Symbol: "STD"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Joegoslavische noviy-dinar", Symbol: "YUM"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanese lek", Symbol: "ALL"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Roemeense leu", Symbol: "RON"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albanese lek (1946–1965)", Symbol: "ALK"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iraanse rial", Symbol: "IRR"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegaskische frank", Symbol: "MCF"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Joegoslavische hervormde dinar (1992–1993)", Symbol: "YUR"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Oost-Duitse ostmark", Symbol: "DDM"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruaanse sol", Symbol: "PES"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgaarse lev (1879–1952)", Symbol: "BGO"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexicaanse zilveren peso (1861–1992)", Symbol: "MXP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Noord-Koreaanse won", Symbol: "KPW"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spaanse peseta (account A)", Symbol: "ESA"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghaanse afghani", Symbol: "AFN"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadzjikistaanse roebel", Symbol: "TJR"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgische frank (financieel)", Symbol: "BEL"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Cubaanse peso", Symbol: "CUP"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: "XPD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayaanse guarani", Symbol: "PYG"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spaanse peseta (convertibele account)", Symbol: "ESB"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgische frank (convertibel)", Symbol: "BEC"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Zuid-Afrikaanse rand (financieel)", Symbol: "ZAL"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Zwitserse frank", Symbol: "CHF"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaanse rial", Symbol: "OMR"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Joegoslavische convertibele dinar", Symbol: "YUN"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macause pataca", Symbol: "MOP"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinaamse dollar", Symbol: "SRD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothaanse loti", Symbol: "LSL"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagassische ariary", Symbol: "MGA"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: "COU"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanees pond", Symbol: "LBP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadese dollar", Symbol: "C$"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeense dram", Symbol: "AMD"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Tsjechoslowaakse harde koruna", Symbol: "CSK"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkongse dollar", Symbol: "HK$"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoolse tugrik", Symbol: "MNT"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Braziliaanse cruzado", Symbol: "BRC"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Wit-Russische nieuwe roebel (1994–1999)", Symbol: "BYB"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Europese rekeneenheid (XBD)", Symbol: "XBD"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnamese dong (1978–1985)", Symbol: "VNN"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexicaanse peso", Symbol: "MX$"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese roepie", Symbol: "NPR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keniaanse shilling", Symbol: "KES"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaanse nieuwe sol", Symbol: "PEN"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltees pond", Symbol: "MTP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haïtiaanse gourde", Symbol: "HTG"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkse lire", Symbol: "TRL"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Amerikaanse dollar", Symbol: "US$"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanese cedi (1979–2007)", Symbol: "GHC"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Oekraïense hryvnia", Symbol: "UAH"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Oegandese shilling (1966–1987)", Symbol: "UGS"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurese lempira", Symbol: "HNL"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnamese dong", Symbol: "₫"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Europese rekeneenheid (XBC)", Symbol: "XBC"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bulgaarse socialistische lev", Symbol: "BGM"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriaanse naira", Symbol: "NGN"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thaise baht", Symbol: "฿"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolese nieuwe kwanza (1990–2000)", Symbol: "AON"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Zilver", Symbol: "XAG"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaarse lev", Symbol: "BGN"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-fondsen", Symbol: "XRE"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonische denar", Symbol: "MKD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tsjechische kroon", Symbol: "CZK"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgaarse harde lev", Symbol: "BGL"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Zuid-Afrikaanse rand", Symbol: "ZAR"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Oude Roemeense leu", Symbol: "ROL"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sovjet-roebel", Symbol: "SUR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiaanse kwacha (1968–2012)", Symbol: "ZMK"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad en Tobago-dollar", Symbol: "TTD"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadoraanse unidad de valor constante (UVC)", Symbol: "ECV"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malinese franc", Symbol: "MLF"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA-frank", Symbol: "FCFA"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franc", Symbol: "CHW"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentijnse peso (1881–1970)", Symbol: "ARM"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Oude Mozambikaanse metical", Symbol: "MZM"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubaanse gulden", Symbol: "AWG"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Argentijnse peso ley (1970–1983)", Symbol: "ARL"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Nederlands-Antilliaanse gulden", Symbol: "NAf."}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Zweedse kroon", Symbol: "SEK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russische roebel", Symbol: "RUB"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Macedonische denar (1992–1993)", Symbol: "MKN"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomon-dollar", Symbol: "SI$"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviaanse boliviano (1863–1963)", Symbol: "BOL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnische convertibele mark", Symbol: "BAM"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letse lats", Symbol: "LVL"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrese nakfa", Symbol: "ERN"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolese kwanza (1977–1990)", Symbol: "AOK"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarees pond", Symbol: "GIP"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Griekse drachme", Symbol: "GRD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayaanse peso en geïndexeerde eenheden", Symbol: "UYI"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Cyprisch pond", Symbol: "CYP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franse franc", Symbol: "FRF"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezolaanse bolivar (1871–2008)", Symbol: "VEB"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiaanse dollar", Symbol: "LRD"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiaanse lire", Symbol: "ITL"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Braziliaanse cruzeiro", Symbol: "BRR"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: "XDR"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Bruneise dollar", Symbol: "BND"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbeidzjaanse manat (1993–2006)", Symbol: "AZM"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaanse dollar", Symbol: "BSD"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerijnse dinar", Symbol: "DZD"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nieuw-Zeelandse dollar", Symbol: "NZ$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Oost-Caribische dollar", Symbol: "EC$"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesische dinar", Symbol: "TND"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmeense manat", Symbol: "TMT"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: "XPT"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Braziliaanse cruzeiro novo (1967–1986)", Symbol: "BRB"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadaanse dollar", Symbol: "BBD"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB-rekeneenheid", Symbol: "XUA"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandese frank", Symbol: "RWF"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawische kwacha", Symbol: "MWK"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "onbekende munteenheid", Symbol: "XXX"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalteekse quetzal", Symbol: "GTQ"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaicaanse dollar", Symbol: "JMD"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaanse tala", Symbol: "WST"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Soedanees pond", Symbol: "SDG"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "IJslandse kroon (1918–1981)", Symbol: "ISJ"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comorese frank", Symbol: "KMF"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Amerikaanse dollar (zelfde dag)", Symbol: "USS"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambikaanse escudo", Symbol: "MZE"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenitische rial", Symbol: "YER"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalische shilling", Symbol: "SOS"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franse UIC-franc", Symbol: "XFU"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Iers pond", Symbol: "IEP"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguaanse córdoba (1988–1991)", Symbol: "NIC"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesische dollar", Symbol: "RHD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiaanse dalasi", Symbol: "GMD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Servische dinar", Symbol: "RSD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Joegoslavische harde dinar", Symbol: "YUD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Kroatische dinar", Symbol: "HRD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswaanse pula", Symbol: "BWP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Noorse kroon", Symbol: "NOK"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltese lire", Symbol: "MTL"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israëlisch pond", Symbol: "ILP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinese renminbi", Symbol: "CN¥"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesische roepia", Symbol: "IDR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexicaanse unidad de inversion (UDI)", Symbol: "MXV"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierraleoonse leone", Symbol: "SLL"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Birmese kyat", Symbol: "BUK"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Letse roebel", Symbol: "LVR"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar van de Chinese Volksbank", Symbol: "CNX"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finse markka", Symbol: "FIM"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxemburgse frank", Symbol: "LUF"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniaanse shilling", Symbol: "TZS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "European Currency Unit", Symbol: "XEU"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: "CHE"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiaanse kwacha", Symbol: "ZMW"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libische dinar", Symbol: "LYD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singaporese dollar", Symbol: "SGD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Europese samengestelde eenheid", Symbol: "XBA"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Ricaanse colon", Symbol: "CRC"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldavische cupon", Symbol: "MDC"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorese escudo", Symbol: "TPE"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Israëlische sjekel (1980–1985)", Symbol: "ILR"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Oekraïense karbovanetz", Symbol: "UAK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Zuid-Soedanees pond", Symbol: "SSP"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayaanse peso (1975–1993)", Symbol: "UYP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazische lilangeni", Symbol: "SZL"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanese cedi", Symbol: "GHS"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibische dollar", Symbol: "NAD"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Amerikaanse dollar (volgende dag)", Symbol: "USN"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Cambodjaanse riel", Symbol: "KHR"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chileense peso", Symbol: "CLP"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nieuwe Taiwanese dollar", Symbol: "NT$"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Chileense escudo", Symbol: "CLE"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankaanse roepie", Symbol: "LKR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaanse dinar", Symbol: "JOD"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolese kwanza reajustado (1995–1999)", Symbol: "AOR"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russische roebel (1991–1998)", Symbol: "RUR"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slowaakse koruna", Symbol: "SKK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandeilands pond", Symbol: "FKP"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmeense manat (1993–2009)", Symbol: "TMM"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chileense unidades de fomento", Symbol: "CLF"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanse yen", Symbol: "JP¥"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Nieuwe Bosnische dinar (1994–1997)", Symbol: "BAN"}}

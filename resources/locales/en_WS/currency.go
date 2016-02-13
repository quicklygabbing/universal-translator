package en_WS

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldovan Leu", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolan New Kwanza (1990–2000)", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentine Peso (1881–1970)", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad & Tobago Dollar", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP Franc", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanaian Cedi", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "European Composite Unit", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgian Lari", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Belarusian New Ruble (1994–1999)", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "South Sudanese Pound", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruban Florin", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolan Kwanza", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerian Dinar", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellois Rupee", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanaian Cedi (1979–2007)", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruvian Sol (1863–1965)", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrainian Karbovanets", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgian Kupon Larit", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Cypriot Pound", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé & Príncipe Dobra", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Chinese People’s Bank Dollar", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawian Kwacha", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezuelan Bolívar", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambican Metical (1980–2006)", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rican Colón", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordanian Dinar", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambian Dalasi", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgarian Lev (1879–1952)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Croatian Dinar", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spanish Peseta (A account)", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankan Rupee", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolan Kwanza (1977–1991)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenyan Shilling", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portuguese Guinea Escudo", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadoran Colón", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwean Dollar (2008)", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapore Dollar", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brazilian New Cruzado (1989–1990)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagasy Franc", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruvian Nuevo Sol", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Romanian Leu (1952–2006)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Icelandic Króna", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Cambodian Riel", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghan Afghani (1927–2002)", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hungarian Forint", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danish Krone", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Israeli Sheqel (1980–1985)", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldovan Cupon", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayan Peso", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "German Mark", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentine Peso (1983–1985)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopian Birr", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrian Pound", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbaijani Manat (1993–2006)", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamanian Balboa", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaican Dollar", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Yugoslavian Hard Dinar (1966–1990)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotian Kip", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Uzbekistani Som", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luxembourg Financial Franc", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polish Zloty", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Irish Pound", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamian Dollar", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnamese Dong", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burmese Kyat", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russian Ruble", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptian Pound", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkish Lira (1922–2005)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australian Dollar", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comorian Franc", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Moroccan Dirham", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese Rupee", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Dutch Guilder", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwean Dollar (1980–2008)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Austrian Schilling", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakhstani Tenge", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leonean Leone", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentine Austral", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberian Dollar", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswanan Pula", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltese Pound", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "South Korean Won", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagasy Ariary", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalan Quetzal", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwaiti Dinar", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Philippine Peso", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Unknown Currency", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Slovenian Tolar", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Romanian Leu", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolan Readjusted Kwanza (1995–1999)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "US Dollar (Next day)", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermudan Dollar", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei Dollar", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Yugoslavian New Dinar (1994–2002)", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnamese Dong (1978–1985)", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "South Korean Won (1945–1953)", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegasque Franc", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandan Franc", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonian Denar", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Yemeni Dinar", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorese Escudo", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruvian Inti", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoan Tala", Symbol: "WS$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swiss Franc", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadorian Unit of Constant Value", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norwegian Krone", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi Riyal", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Testing Currency Code", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland Islands Pound", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brazilian Real", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar Kyat", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Silver", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbaijani Manat", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Latvian Lats", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET Funds", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau Peso", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayan Peso (1975–1993)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar Pound", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrean Nakfa", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho Loti", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Cayman Islands Dollar", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudanese Dinar (1992–2007)", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Bolivian Boliviano (1863–1963)", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Macedonian Denar (1992–1993)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Herzegovina Convertible Mark", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Gold", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese Franc", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Ugandan Shilling (1966–1987)", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbian Dinar", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombian Peso", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Soviet Rouble", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadian Dollar", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigerian Naira", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadian Dollar", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua New Guinean Kina", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verdean Escudo", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutian Franc", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Czechoslovak Hard Koruna", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraqi Dinar", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnia-Herzegovina Dinar (1992–1994)", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkish Lira", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italian Lira", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "European Unit of Account (XBC)", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolian Tugrik", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russian Ruble (1991–1998)", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltese Lira", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brazilian Cruzeiro (1993–1994)", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "French UIC-Franc", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indian Rupee", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexican Silver Peso (1861–1992)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chilean Peso", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeshi Taka", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Swedish Krona", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranian Rial", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesian Dollar", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandan Shilling", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "French Gold Franc", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "South Korean Hwan (1953–1962)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeli New Sheqel", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zairean Zaire (1971–1993)", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibian Dollar", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Polish Zloty (1950–1995)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgian Franc", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmenistani Manat (1993–2009)", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgian Franc (convertible)", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Greek Drachma", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Croatian Kuna", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "European Currency Unit", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thai Baht", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "British Pound", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuelan Bolívar (1871–2008)", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaysian Ringgit", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadorian Sucre", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Czech Republic Koruna", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambian Kwacha", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgian Franc (financial)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR Euro", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brazilian Cruzeiro (1942–1967)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haitian Gourde", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudanese Pound", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Latvian Ruble", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguan Córdoba", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malian Franc", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunisian Dinar", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayan Guarani", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "US Dollar (Same day)", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambican Escudo", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Serbian Dinar (2002–2006)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albanian Lek (1946–1965)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Central African CFA Franc", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenian Dram", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Yemeni Rial", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tajikistani Ruble", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Netherlands Antillean Guilder", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzanian Shilling", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spanish Peseta", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyanaese Dollar", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Cuban Peso", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinese Yuan", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omani Rial", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanese Yen", Symbol: "¥"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrainian Hryvnia", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fijian Dollar", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luxembourgian Convertible Franc", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivian Boliviano", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Lithuanian Talonas", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena Pound", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "East German Mark", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "New Taiwan Dollar", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macanese Pataca", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinean Franc", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "South African Rand", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chilean Unit of Account (UF)", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "European Monetary Unit", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Yugoslavian Convertible Dinar (1990–1992)", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxembourgian Franc", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongan Paʻanga", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "South African Rand (financial)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platinum", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivian Rufiyaa", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivian Mvdol", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Cuban Convertible Peso", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinamese Dollar", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "US Dollar", Symbol: "$"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Yugoslavian Reformed Dinar (1992–1993)", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudanese Pound (1957–1998)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "West African CFA Franc", Symbol: ""}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zairean New Zaire (1993–1998)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israeli Pound", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somali Shilling", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritanian Ouguiya", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estonian Kroon", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwean Dollar (2009)", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Chilean Escudo", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Moroccan Franc", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belarusian Ruble", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "Colombian Real Value Unit", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "North Korean Won", Symbol: ""}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldivian Rupee (1947–1981)", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR Franc", Symbol: ""}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB Unit of Account", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinean Syli", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivian Peso", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tajikistani Somoni", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistani Manat", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazi Lilangeni", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatari Rial", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayan Peso (Indexed Units)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Lithuanian Litas", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu Vatu", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanian Lek", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatorial Guinean Ekwele", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesian Rupiah", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgarian Hard Lev", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambican Metical", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portuguese Escudo", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritian Rupee", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "East Caribbean Dollar", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovak Koruna", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Bosnia-Herzegovina New Dinar (1994–1997)", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libyan Dinar", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghan Afghani", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexican Peso", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "European Unit of Account (XBD)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finnish Markka", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentine Peso", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahraini Dinar", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistani Rupee", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong Dollar", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brazilian Cruzado (1986–1989)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brazilian Cruzeiro (1990–1993)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgarian Lev", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominican Peso", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Argentine Peso Ley (1970–1983)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize Dollar", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "French Franc", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spanish Peseta (convertible account)", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinamese Guilder", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexican Investment Unit", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutanese Ngultrum", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguan Córdoba (1988–1991)", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduran Lempira", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundian Franc", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Icelandic Króna (1918–1981)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bulgarian Socialist Lev", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Lebanese Pound", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "United Arab Emirates Dirham", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Solomon Islands Dollar", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "New Zealand Dollar", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kyrgystani Som", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorran Peseta", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brazilian New Cruzeiro (1967–1986)", Symbol: ""}}

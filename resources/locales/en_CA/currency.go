package en_CA

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GBP": ut.Currency{Currency: "GBP", DisplayName: "British Pound", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finnish Markka", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Central African CFA Franc", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP Franc", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzanian Shilling", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivian Boliviano", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polish Zloty", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Yemeni Rial", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanaian Cedi", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxembourgian Franc", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Soviet Rouble", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "New Taiwan Dollar", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolan Readjusted Kwanza (1995–1999)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgarian Lev (1879–1952)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanian Lek", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Bosnia-Herzegovina New Dinar (1994–1997)", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Argentine Peso Ley (1970–1983)", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenian Dram", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigerian Naira", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Croatian Dinar", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonian Denar", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei Dollar", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguan Córdoba", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "German Mark", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivian Peso", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Unknown Currency", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistani Rupee", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolian Tugrik", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnamese Dong (1978–1985)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolan Kwanza", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libyan Dinar", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hungarian Forint", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau Peso", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuelan Bolívar (1871–2008)", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR Franc", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luxembourgian Convertible Franc", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Herzegovina Convertible Mark", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Romanian Leu", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakhstani Tenge", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Czechoslovak Hard Koruna", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistani Manat", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruvian Inti", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutian Franc", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brazilian Cruzeiro (1993–1994)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgian Kupon Larit", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesian Rupiah", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar Pound", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platinum", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadoran Colón", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kyrgystani Som", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagasy Franc", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Dutch Guilder", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chilean Peso", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambican Escudo", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagasy Ariary", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Polish Zloty (1950–1995)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Lebanese Pound", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comorian Franc", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Gold", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentine Peso (1881–1970)", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Moroccan Dirham", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Cambodian Riel", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Testing Currency Code", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaican Dollar", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malian Franc", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Cuban Convertible Peso", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estonian Kroon", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruvian Sol (1863–1965)", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayan Peso", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Serbian Dinar (2002–2006)", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong Dollar", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnamese Dong", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadorian Unit of Constant Value", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "US Dollar (Next day)", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena Pound", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesian Dollar", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraqi Dinar", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Yugoslavian Reformed Dinar (1992–1993)", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leonean Leone", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinamese Guilder", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrian Pound", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somali Shilling", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgarian Hard Lev", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudanese Dinar (1992–2007)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bulgarian Socialist Lev", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russian Ruble", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland Islands Pound", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldovan Cupon", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapore Dollar", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Cuban Peso", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Bolivian Boliviano (1863–1963)", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thai Baht", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Macedonian Denar (1992–1993)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezuelan Bolívar", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "South African Rand", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldovan Leu", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fijian Dollar", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Chilean Escudo", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Philippine Peso", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brazilian New Cruzeiro (1967–1986)", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexican Investment Unit", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brazilian Real", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatorial Guinean Ekwele", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadian Dollar", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé & Príncipe Dobra", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduran Lempira", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Lithuanian Litas", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "French Franc", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "US Dollar", Symbol: "$"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "New Zealand Dollar", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambian Dalasi", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Latvian Ruble", Symbol: ""}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omani Rial", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombian Peso", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgian Franc (convertible)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italian Lira", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese Rupee", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexican Silver Peso (1861–1992)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswanan Pula", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanese Yen", Symbol: "¥"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldivian Rupee (1947–1981)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brazilian New Cruzado (1989–1990)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkish Lira (1922–2005)", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandan Shilling", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritanian Ouguiya", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "US Dollar (Same day)", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnia-Herzegovina Dinar (1992–1994)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrean Nakfa", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghan Afghani (1927–2002)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeshi Taka", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellois Rupee", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwean Dollar (2008)", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwaiti Dinar", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israeli Pound", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Czech Republic Koruna", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "French Gold Franc", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tajikistani Somoni", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verdean Escudo", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotian Kip", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwean Dollar (2009)", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinean Franc", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerian Dinar", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Cypriot Pound", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgian Lari", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentine Austral", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrainian Hryvnia", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "French UIC-Franc", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranian Rial", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutanese Ngultrum", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "South African Rand (financial)", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltese Lira", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR Euro", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibian Dollar", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudanese Pound", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgian Franc", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgarian Lev", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swiss Franc", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Netherlands Antillean Guilder", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanaian Cedi (1979–2007)", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorran Peseta", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brazilian Cruzeiro (1942–1967)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolan Kwanza (1977–1991)", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmenistani Manat (1993–2009)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macanese Pataca", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazi Lilangeni", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovak Koruna", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptian Pound", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambian Kwacha", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Greek Drachma", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalan Quetzal", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "European Composite Unit", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorese Escudo", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudanese Pound (1957–1998)", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Ugandan Shilling (1966–1987)", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Latvian Lats", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spanish Peseta", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Yemeni Dinar", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Swedish Krona", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgian Franc (financial)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albanian Lek (1946–1965)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahraini Dinar", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "United Arab Emirates Dirham", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambican Metical", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamanian Balboa", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "European Unit of Account (XBD)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spanish Peseta (convertible account)", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkish Lira", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongan Paʻanga", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadorian Sucre", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luxembourg Financial Franc", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbian Dinar", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Belarusian New Ruble (1994–1999)", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunisian Dinar", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "South Korean Won (1945–1953)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatari Rial", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyanaese Dollar", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Yugoslavian Convertible Dinar (1990–1992)", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua New Guinean Kina", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "European Unit of Account (XBC)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeli New Sheqel", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Lithuanian Talonas", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Irish Pound", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad & Tobago Dollar", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belarusian Ruble", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australian Dollar", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawian Kwacha", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize Dollar", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET Funds", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayan Guarani", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentine Peso (1983–1985)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Yugoslavian Hard Dinar (1966–1990)", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "Colombian Real Value Unit", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norwegian Krone", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danish Krone", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russian Ruble (1991–1998)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinese Yuan", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tajikistani Ruble", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivian Mvdol", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritian Rupee", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "South Korean Won", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankan Rupee", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "West African CFA Franc", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayan Peso (1975–1993)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brazilian Cruzeiro (1990–1993)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "North Korean Won", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho Loti", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Croatian Kuna", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rican Colón", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordanian Dinar", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Icelandic Króna", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chilean Unit of Account (UF)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar Kyat", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Icelandic Króna (1918–1981)", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "East Caribbean Dollar", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Romanian Leu (1952–2006)", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandan Franc", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Chinese People’s Bank Dollar", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi Riyal", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "European Currency Unit", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegasque Franc", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zairean New Zaire (1993–1998)", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Cayman Islands Dollar", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexican Peso", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruban Florin", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayan Peso (Indexed Units)", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinamese Dollar", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zairean Zaire (1971–1993)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadian Dollar", Symbol: "$"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguan Córdoba (1988–1991)", Symbol: ""}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB Unit of Account", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Silver", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indian Rupee", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermudan Dollar", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruvian Nuevo Sol", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenyan Shilling", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamian Dollar", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Solomon Islands Dollar", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haitian Gourde", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Israeli Sheqel (1980–1985)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burmese Kyat", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "South Sudanese Pound", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portuguese Escudo", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopian Birr", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberian Dollar", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Uzbekistani Som", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentine Peso", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltese Pound", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolan New Kwanza (1990–2000)", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominican Peso", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoan Tala", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Austrian Schilling", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghan Afghani", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbaijani Manat", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaysian Ringgit", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwean Dollar (1980–2008)", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Yugoslavian New Dinar (1994–2002)", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu Vatu", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Moroccan Franc", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundian Franc", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spanish Peseta (A account)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambican Metical (1980–2006)", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivian Rufiyaa", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "South Korean Hwan (1953–1962)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese Franc", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinean Syli", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbaijani Manat (1993–2006)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "East German Mark", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portuguese Guinea Escudo", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "European Monetary Unit", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brazilian Cruzado (1986–1989)", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Slovenian Tolar", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrainian Karbovanets", Symbol: ""}}

package gsw

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"FRF": ut.Currency{Currency: "FRF", DisplayName: "Französische Franc", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugiisische Guinea Escudo", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda-Franc", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brasilianische Cruzeiro", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israelischs Pfund", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Cordoba", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russische Rubel (alt)", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayische Nöie Peso (1975–1993)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Gold", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA-Franc (Wescht)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Auschtralische Dollar", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belarus Rubel (nöi)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgarische Lew", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea-Franc", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Kaiman-Dollar", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Uganda-Schilling (1966–1987)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR-Euro", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Schpanischi Peseeta (A–Kontene)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irak-Dinar", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutanische Ngultrum", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Renminbi Yuan", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Zypere-Pfund", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkanische Dirham", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoslawische Dinar (konvertibel)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize-Dollar", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland-Pfund", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanaische Cedi (GHC)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luxemburgische Franc (konvertibel)", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Malteesischi Lira", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litauische Litas", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexikanische Unidad de Inversion (UDI)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nöii Taiwan-Dollar", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha (1968–2012)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianische Real", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Belarus-Rubel (alt)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Verrächnigsäiheit für EC", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaanische Dinar", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Eestnischi Chroone", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgische Kupon Larit", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwait-Dinar", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados-Dollar", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgische Finanz-Franc", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Birmanische Kyat", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luxemburgischer Finanz-Franc", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiänischi Lira", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi-Riyal", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Tschechoslowakischi Chroone", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomone-Dollar", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayische Peso", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentinische Peso", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Konvertierbari Mark vo Bosnie und Herzegowina", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadorianische Sucre", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudaneesischs Pfund", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri-Lanka-Rupie", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbische Dinar", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Oschtkaribische Dollar", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentinische Peso (1983–1985)", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnie-und-Herzegowina-Dinar", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukraiinische Karbovanetz", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "US-Dollar", Symbol: "$"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswanische Pula", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Tüütschi Mark", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guineische Syli", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Türkischi Liire", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norweegischi Chroone", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Europääischi Währigseinheit (XEU)", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Holländische Gulde", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Dschibuti-Franc", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambische Dalasi", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawi-Kwacha", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Simbabwe-Dollar", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rica Colon", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar-Pfund", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seyschelle-Rupie", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Nöii Türkischi Liire", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgische Lari", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russische Rubel", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadschikischtan-Somoni", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brasilianische Cruzado", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "DDR-Mark", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Löi", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jeme-Dinar", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kap Verde Escudo", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malische Franc", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadschikischtan-Rubel", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litauische Talonas", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkong-Dollar", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghani", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Aserbeidschanische Manat (1993–2006)", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Schpanischi Peseeta", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Süürischs Pfund", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolanische Kwanza (1977–1990)", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaysische Ringgit", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudaneesischs Pfund (alt)", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinamische Gulde", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Französische Gold-Franc", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Jugoslawische Dinar (1966–1990)", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudaneesische Dinar", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brasilianische Cruzeiro Novo (1967–1986)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahama-Dollar", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Kroazische Dinar", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Iirischs Pfund", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakischtanischi Rupie", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei-Dollar", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brasilianische Cruzado Novo", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Süüdkoreanische Won", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanaische Cedi (GHS)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Iisländischi Chroone", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Nordkoreanische Won", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mosambikanische Escudo", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahrain-Dinar", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brasilianische Cruzeiro (1990–1993)", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR-Franke", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugiisische Escudo", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmeenischtan-Manat", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Silber", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkanischer Franc", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaschkar-Franc", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesische Dollar", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Kwanza Reajustado", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentinische Auschtral", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivianische Mvdol", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Lüübische Dinar", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumäänische Löi", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Neuseeland-Dollar", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorranischi Peseete", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolumbianische Peso", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lettische Lats", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Lettische Rubel", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platin", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Äquatorialguinea-Ekwele", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesischi Rupie", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxemburgische Franc", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial Omani", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar-Riyal", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timor-Escudo", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexikanische Peso", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepaleesischi Rupie", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Malteesischs Pfund", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinamische Dollar", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Unbekannti Währig", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Nöie Zaire", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivianische Peso", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Tschileenische Unidad de Fomento", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Tschileenische Peso", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiaanische Dollar", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libaneesischs Pfund", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sowjetische Rubel", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau-Peso", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Schweedischi Chroone", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba Florin", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruanische Inti", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "El-Salvador-Colon", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Europääischi Rächnigseinheit (XBD)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "UAE Dirham", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Schekel", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slowakischi Chroone", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda-Schilling", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Schpanischi Peseeta (konvertibel)", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komore-Franc", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pfund Schtörling", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbeidschanische Manat", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidschi Dollar", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen", Symbol: "¥"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Maurizius-Rupie", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Öschtriichische Schilling", Symbol: "öS"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi-Franc", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indischi Rupie", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-Funds", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Nöii Dinar", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "US Dollar (Gliiche Taag)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgische Franc", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldau-Löi", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaschkar-Ariary", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "US Dollar (Nöchschte Taag)", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritreische Nakfa", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibia-Dollar", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Sunderziäigsrächt", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenia-Schilling", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tuneesische Dinar", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-Franc", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapur-Dollar", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Tänischi Chroone", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Ägüptischs Pfund", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finnischi Mark", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaragua-Córdoba", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgische Franc (konvertibel)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Schwiizer Franke", Symbol: "CHF"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jeme-Rial", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-Dollar", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika-Dollar", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St.-Helena-Pfund", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalia-Schilling", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Nöie Sol", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolivar (1871–2008)", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Europääischi Währigseinheit (XBB)", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Europääischi Rächnigseinheit (XBC)", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghani (1927–2002)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongolesische Franc", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Alte Serbische Dinar", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Alte Metical", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyana-Dollar", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Europääischi Rächnigseinheit", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Französische UIC-Franc", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbeekischtan-Sum", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolivar", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexikanische Silber-Peso (1861–1992)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloty (1950–1995)", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Süüdsudaneesischs Pfund", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeerischi Dinar", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Niderländischi-Antille-Gulde", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lew (1962–1999)", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Griechische Trachme", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubanische Peso", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tschechischi Chroone", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Äthiopische Birr", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Nöie Kwanza", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Philippiinische Peso", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansania-Schilling", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad-und-Tobago-Dollar", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA-Franc (Äquatoriaal)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Teschtwährig", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadische Dollar", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Tominikanische Peso", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia", Symbol: ""}}

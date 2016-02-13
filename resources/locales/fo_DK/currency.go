package fo_DK

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbekistan som", Symbol: "UZS"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesia rupiah", Symbol: "IDR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordan dinar", Symbol: "JOD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansania skillingur", Symbol: "TZS"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "unse palladium", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar pund", Symbol: "GIP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irak dinar", Symbol: "IQD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaoyggjar dollari", Symbol: "BSD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Suðursudan pund", Symbol: "SSP"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepal rupi", Symbol: "NPR"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkaland liri", Symbol: "TRY"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukraina hryvnia", Symbol: "UAH"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "bretsk pund", Symbol: "£"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar (Burma) kyat", Symbol: "MMK"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Vesturafrika CFA frankur", Symbol: "CFA"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panama balboa", Symbol: "PAB"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japanskur yen", Symbol: "JP¥"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominika peso", Symbol: "DOP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksiko peso", Symbol: "MX$"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanon pund", Symbol: "LBP"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda dollari", Symbol: "BMD"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanada dollari", Symbol: "CA$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "sveisiskur frankur", Symbol: "CHF"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumenia leu", Symbol: "RON"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Ísrael new sheqel", Symbol: "₪"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena pund", Symbol: "SHP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokko dirham", Symbol: "MAD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptaland pund", Symbol: "EGP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemala quetzal", Symbol: "GTQ"}, "USD": ut.Currency{Currency: "USD", DisplayName: "US dollari", Symbol: "US$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Kosta Rika colón", Symbol: "CRC"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laos kip", Symbol: "LAK"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopia birr", Symbol: "ETB"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haiti gourde", Symbol: "HTG"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Hvítarussland ruble", Symbol: "BYR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambik metical", Symbol: "MZN"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbadjan manat", Symbol: "AZN"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatia kuna", Symbol: "HRK"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolombia peso", Symbol: "COP"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Butan ngultrum", Symbol: "BTN"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kuba peso", Symbol: "CUP"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymanoyggjar dollari", Symbol: "KYD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistan manat", Symbol: "TMT"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Ungarn forint", Symbol: "HUF"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Evra", Symbol: "€"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Gujana dollari", Symbol: "GYD"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaria lev", Symbol: "BGN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Barein dinar", Symbol: "BHD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taivan new dollari", Symbol: "NT$"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiarabia riyal", Symbol: "SAR"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ókent gjaldoyra", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peru nuevo sol", Symbol: "PEN"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seyskelloyggjar rupi", Symbol: "SCR"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Makao pataca", Symbol: "MOP"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeria naira", Symbol: "NGN"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tailand baht", Symbol: "THB"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenjanskur skillingur", Symbol: "KES"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguai guarani", Symbol: "PYG"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norsk króna", Symbol: "NOK"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi frankur", Symbol: "BIF"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldova leu", Symbol: "MDL"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistan afghani", Symbol: "AFN"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuvait dinar", Symbol: "KWD"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemen rial", Symbol: "YER"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Eystur Karibia dollari", Symbol: "EC$"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kinesiskur yuan", Symbol: "CN¥"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda skillingur", Symbol: "UGX"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolia tugrik", Symbol: "MNT"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentina peso", Symbol: "ARS"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadsjikistan somoni", Symbol: "TJS"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaragua córdoba", Symbol: "NIO"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Sameindu Emirríkini dirham", Symbol: "AED"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivia boliviano", Symbol: "BOB"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranskir rials", Symbol: "IRR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipsoyggjar peso", Symbol: "PHP"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russland ruble", Symbol: "RUB"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Niðurlonds Karibia gyllin", Symbol: "ANG"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua Nýguinea kina", Symbol: "PGK"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Kekkia koruna", Symbol: "CZK"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djibuti frankur", Symbol: "DJF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoroyggjar frankur", Symbol: "KMF"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "íslendsk króna", Symbol: "ISK"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji dollari", Symbol: "FJD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibia dollari", Symbol: "NAD"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svasiland lilangeni", Symbol: "SZL"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudan pund", Symbol: "SDG"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lanka rupi", Symbol: "LKR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistan rupi", Symbol: "PKR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Oman rial", Symbol: "OMR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonga paʻanga", Symbol: "TOP"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenia dram", Symbol: "AMD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leona leone", Symbol: "SLL"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguai peso", Symbol: "UYU"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar rial", Symbol: "QAR"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba florin", Symbol: "AWG"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaisia ringgit", Symbol: "MYR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Sýria pund", Symbol: "SYP"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angola kwanza", Symbol: "AOA"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalia skillingur", Symbol: "SOS"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivoyggjar rufiyaa", Symbol: "MVR"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea frankur", Symbol: "GNF"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo frankur", Symbol: "CDF"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Sambia kwacha", Symbol: "ZMW"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskar ariary", Symbol: "MGA"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Norðurkorea won", Symbol: "KPW"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vjetnam dong", Symbol: "₫"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedónia denar", Symbol: "MKD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kasakstan tenge", Symbol: "KZT"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malavi kwacha", Symbol: "MWK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgia lari", Symbol: "GEL"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svensk króna", Symbol: "SEK"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libya dinar", Symbol: "LYD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kuba peso (sum kann vekslast)", Symbol: "CUC"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome & Prinsipi dobra", Symbol: "STD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Suðurafrika rand", Symbol: "ZAR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinam dollari", Symbol: "SRD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Hersegovina mark (kann vekslast)", Symbol: "BAM"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodja riel", Symbol: "KHR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Pólland zloty", Symbol: "PLN"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Miðafrika CFA frankur", Symbol: "FCFA"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda frankur", Symbol: "RWF"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Móritius rupi", Symbol: "MUR"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albania lek", Symbol: "ALL"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Móritania ouguiya", Symbol: "MRO"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeria dinar", Symbol: "DZD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Gana cedi", Symbol: "GHS"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrea nakfa", Symbol: "ERN"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoa tala", Symbol: "WST"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu vatu", Symbol: "VUV"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Grønhøvdaoyggjar escudo", Symbol: "CVE"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianskur real", Symbol: "R$"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapor dollari", Symbol: "SGD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "donsk króna", Symbol: "kr."}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberia dollari", Symbol: "LRD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Avstralskur dollari", Symbol: "A$"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "unse platin", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frankur", Symbol: "CFPF"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad & Tobago dollari", Symbol: "TTD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "unse guld", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika dollari", Symbol: "JMD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "unse sølv", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Kili peso", Symbol: "CLP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Suðurkorea won", Symbol: "₩"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisia som", Symbol: "KGS"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados dollari", Symbol: "BBD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduras lempira", Symbol: "HNL"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botsvana pula", Symbol: "BWP"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesj taka", Symbol: "BDT"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venesuela bolívar", Symbol: "VEF"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei dollari", Symbol: "BND"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong dollari", Symbol: "HK$"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nýsæland dollari", Symbol: "NZ$"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonoyggjar dollari", Symbol: "SBD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesia dinar", Symbol: "TND"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandsoyggjar pund", Symbol: "FKP"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbia dinar", Symbol: "RSD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambia dalasi", Symbol: "GMD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belis dollari", Symbol: "BZD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiskir rupis", Symbol: "₹"}}

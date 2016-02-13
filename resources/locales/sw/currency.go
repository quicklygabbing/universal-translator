package sw

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat ya Azebaijani", Symbol: "AZN"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faranga ya Burundi", Symbol: "BIF"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso ya Cuba", Symbol: "CUP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: "SZL"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Faranga ya CFP", Symbol: "CFPF"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr ya Uhabeshi", Symbol: "ETB"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Pauni ya Visiwa vya Falkland", Symbol: "FKP"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira ya Hondurasi", Symbol: "HNL"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won ya Korea Kaskazini", Symbol: "KPW"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dola ya Singapore", Symbol: "SGD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso ya Cuba Inayoweza Kubadilishwa", Symbol: "CUC"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Pauni ya Misri", Symbol: "EGP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi ya Ghana", Symbol: "GHS"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Pauni ya Sudani Kusini", Symbol: "SSP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "STD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som ya Uzibekistani", Symbol: "UZS"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghani ya Afghanistan", Symbol: "AFN"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ya Gambia", Symbol: "GMD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar ya Serbia", Symbol: "RSD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinari ya Tunisia", Symbol: "TND"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dola ya Fiji", Symbol: "FJD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dola ya Hong Kong", Symbol: "HK$"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Krona ya Aisilandi", Symbol: "ISK"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Pauni ya Lebanon", Symbol: "LBP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial ya Omani", Symbol: "OMR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty ya Polandi", Symbol: "PLN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dola ya Kanada", Symbol: "CA$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faranga ya Uswisi", Symbol: "CHF"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pauni ya Uingereza", Symbol: "£"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala ya Samoa", Symbol: "WST"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Koruna ya Jamhuri ya Cheki", Symbol: "CZK"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats ya Lativia", Symbol: "LVL"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha ya Malawi", Symbol: "MWK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dola ya Visiwa vya Solomon", Symbol: "SBD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dola ya Marekani", Symbol: "US$"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek ya Albania", Symbol: "ALL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial ya Iran", Symbol: "IRR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia ya India", Symbol: "₹"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip ya Laosi", Symbol: "LAK"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham ya Falme za Kiarabu", Symbol: "AED"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram ya Armenia", Symbol: "AMD"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev ya Bulgaria", Symbol: "BGN"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Ruble ya Belarusi", Symbol: "BYR"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon ya Kostarika", Symbol: "CRC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Eskudo ya Cape Verde", Symbol: "CVE"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya ya Moritania", Symbol: "MRO"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga ya Tonga", Symbol: "TOP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha ya Zambia", Symbol: "ZMW"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka ya Bangladeshi", Symbol: "BDT"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia ya Ushelisheli", Symbol: "SCR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Krona ya Uswidi", Symbol: "SEK"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa ya Maldivi", Symbol: "MVR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metikali ya Msumbiji", Symbol: "MZN"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Faranga ya Jibuti", Symbol: "DJF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinari ya Aljeria", Symbol: "DZD"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Sedi ya Ghana", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dola ya Guyana", Symbol: "GYD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som ya Kyrgystani", Symbol: "KGS"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ya Lesoto", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pauni ya Sudani (1957–1998)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel ya Kambodia", Symbol: "KHR"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu ya Moldova", Symbol: "MDL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu ya Vanuatu", Symbol: "VUV"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar ya Bahareni", Symbol: "BHD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Faranga ya Guinea", Symbol: "GNF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faranga ya Komoro", Symbol: "KMF"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilingi ya Uganda", Symbol: "UGX"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna ya Kroeshia", Symbol: "HRK"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar ya Masedonia", Symbol: "MKD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca ya Macau", Symbol: "MOP"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metikali ya Msumbiji (1980–2006)", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht ya Tailandi", Symbol: "฿"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dola ya Taiwan", Symbol: "NT$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dola ya Karibea ya Mashariki", Symbol: "EC$"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari ya Georgia", Symbol: "GEL"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal ya Guatemala", Symbol: "GTQ"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar ya Yordani", Symbol: "JOD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia ya Nepali", Symbol: "NPR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani ya Paragwai", Symbol: "PYG"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Pauni ya Sudani", Symbol: "SDG"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso ya Urugwai", Symbol: "UYU"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Sarafu isiyojulikana", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dola ya Barbados", Symbol: "BBD"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dola ya Liberia", Symbol: "LRD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik ya Mongolia", Symbol: "MNT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ya Nijeria", Symbol: "NGN"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Cordoba ya Nikaragua", Symbol: "NIO"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina ya Papua New Guinea", Symbol: "PGK"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial ya Yemeni", Symbol: "YER"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan ya Uchina", Symbol: "CN¥"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Pauni ya Gibraltar", Symbol: "GIP"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin ya Aruba", Symbol: "AWG"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso ya Chile", Symbol: "CLP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dola ya Jamaica", Symbol: "JMD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit ya Malaysia", Symbol: "MYR"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Krone ya Norwe", Symbol: "NOK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen ya Japani", Symbol: "JP¥"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dola ya Nyuzilandi", Symbol: "NZ$"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni ya Tajikistani", Symbol: "TJS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real ya Brazil", Symbol: "R$"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faranga ya Rwanda", Symbol: "RWF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ya Afrika Kusini", Symbol: "ZAR"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Krone ya Denmaki", Symbol: "DKK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pauni ya Santahelena", Symbol: "SHP"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dola ya Trinidad na Tobago", Symbol: "TTD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong ya Vietnam", Symbol: "₫"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dola ya Visiwa vya Cayman", Symbol: "KYD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari ya Madagaska", Symbol: "MGA"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dola ya Namibia", Symbol: "NAD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso ya Ufilipino", Symbol: "PHP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faranga ya Gine", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham ya Moroko", Symbol: "MAD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia ya Morisi", Symbol: "MUR"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu ya Romania", Symbol: "RON"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dola ya Brunei", Symbol: "BND"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum ya Bhutan", Symbol: "BTN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dola ya Belize", Symbol: "BZD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde ya Haiti", Symbol: "HTG"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia ya Sri Lanka", Symbol: "LKR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas ya Lithuania", Symbol: "LTL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial ya Qatari", Symbol: "QAR"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira ya Uturuki", Symbol: "TRY"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dola ya Bahamas", Symbol: "BSD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ya Botswana", Symbol: "BWP"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso ya Kolombia", Symbol: "COP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupiah ya Indonesia", Symbol: "IDR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won ya Korea Kusini", Symbol: "₩"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar ya Kuwaiti", Symbol: "KWD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilingi ya Somalia", Symbol: "SOS"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Guilder ya Antili za Kiholanzi", Symbol: "ANG"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dola ya Bermuda", Symbol: "BMD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano ya Bolivia", Symbol: "BOB"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso ya Dominika", Symbol: "DOP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilingi ya Kenya", Symbol: "Ksh"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat ya Myama", Symbol: "MMK"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso ya Ajentina", Symbol: "ARS"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinari ya Libya", Symbol: "LYD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal ya Saudia", Symbol: "SAR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dola ya Suriname", Symbol: "SRD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Pauni ya Syria", Symbol: "SYP"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dola ya Zimbabwe", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dola ya Australia", Symbol: "A$"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Yuro", Symbol: "€"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint ya Hungaria", Symbol: "HUF"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat ya Turukimenistani", Symbol: "TMT"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faranga ya Kongo", Symbol: "CDF"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa ya Panama", Symbol: "PAB"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia ya Pakistani", Symbol: "PKR"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolivar ya Venezuela", Symbol: "VEF"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA faranga za BCEAO", Symbol: "CFA"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza ya Angola", Symbol: "AOA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Convertible Mark ya Bosnia na Hezegovina", Symbol: "BAM"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge ya Kazakistani", Symbol: "KZT"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso ya Meksiko", Symbol: "MX$"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone", Symbol: "SLL"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA faranga ya BEAC", Symbol: "FCFA"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa ya Eritrea", Symbol: "ERN"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel Mpya ya Israeli", Symbol: "₪"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar ya Iraki", Symbol: "IQD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Nuevo Sol ya Peru", Symbol: "PEN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Ruble ya Urusi", Symbol: "RUB"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia ya Ukrania", Symbol: "UAH"}}

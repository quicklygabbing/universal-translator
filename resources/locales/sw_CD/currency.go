package sw_CD

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ANG": ut.Currency{Currency: "ANG", DisplayName: "Guilder ya Antili za Kiholanzi", Symbol: "ANG"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dola ya Brunei", Symbol: "BND"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dola ya Karibea ya Mashariki", Symbol: "EC$"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinari ya Aljeria", Symbol: "DZD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Sarafu ya Kijapani", Symbol: "JP¥"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Krone ya Norwe", Symbol: "NOK"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia ya Sri Lanka", Symbol: "LKR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal ya Saudia", Symbol: "SAR"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dola ya Australia", Symbol: "A$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira ya Hondurasi", Symbol: "HNL"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Pauni ya Lebanon", Symbol: "LBP"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dola ya Visiwa vya Solomon", Symbol: "SBD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilingi ya Uganda", Symbol: "UGX"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham ya Falme za Kiarabu", Symbol: "AED"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial ya Iran", Symbol: "IRR"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia ya Nepali", Symbol: "NPR"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde ya Haiti", Symbol: "HTG"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilingi ya Somalia", Symbol: "SOS"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Pauni ya Misri", Symbol: "EGP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Krona ya Uswidi", Symbol: "SEK"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Faranga CFA BCEAO", Symbol: "CFA"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Pauni ya Syria", Symbol: "SYP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metikali ya Msumbiji", Symbol: "MZN"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso ya Chile", Symbol: "CLP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso ya Ufilipino", Symbol: "PHP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial ya Qatari", Symbol: "QAR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilingi ya Kenya", Symbol: "Ksh"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong ya Vietnam", Symbol: "₫"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dola ya Zimbabwe", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ya Gambia", Symbol: "GMD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek ya Albania", Symbol: "ALL"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ya Nijeria", Symbol: "NGN"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty ya Polandi", Symbol: "PLN"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ya Lesoto", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faranga ya Komoro", Symbol: "KMF"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Sarafu isiyojulikana", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa ya Maldivi", Symbol: "MVR"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary ya Bukini", Symbol: "MGA"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht ya Tailandi", Symbol: "฿"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ya Botswana", Symbol: "BWP"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolivar ya Venezuela", Symbol: "VEF"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar ya Serbia", Symbol: "RSD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinari ya Tunisia", Symbol: "TND"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala ya Samoa", Symbol: "WST"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Faranga ya Jibuti", Symbol: "DJF"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna ya Kroeshia", Symbol: "HRK"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal ya Guatemala", Symbol: "GTQ"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat ya Turukimenistani", Symbol: "TMT"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza ya Angola", Symbol: "AOA"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faranga ya Burundi", Symbol: "BIF"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won ya Korea Kaskazini", Symbol: "KPW"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Pauni ya Sudani Kusini", Symbol: "SSP"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faranga ya Kongo", Symbol: "FC"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pauni ya Sudani (1957–1998)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dola ya Kanada", Symbol: "CA$"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dola ya Nyuzilandi", Symbol: "NZ$"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum ya Bhutan", Symbol: "BTN"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dola ya Bermuda", Symbol: "BMD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dola ya Taiwan", Symbol: "NT$"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pauni ya Santahelena", Symbol: "SHP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu ya Romania", Symbol: "RON"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge ya Kazakistani", Symbol: "KZT"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghani ya Afghanistan", Symbol: "AFN"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial ya Omani", Symbol: "OMR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso ya Urugwai", Symbol: "UYU"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dola ya Trinidad na Tobago", Symbol: "TTD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dola ya Fiji", Symbol: "FJD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga ya Tonga", Symbol: "TOP"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dola ya Guyana", Symbol: "GYD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Pauni ya Gibraltar", Symbol: "GIP"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Cordoba ya Nikaragua", Symbol: "NIO"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won ya Korea Kusini", Symbol: "₩"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat ya Myama", Symbol: "MMK"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso ya Dominika", Symbol: "DOP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Yuro", Symbol: "€"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso ya Cuba", Symbol: "CUP"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia ya India", Symbol: "₹"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani ya Paragwai", Symbol: "PYG"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Sedi ya Ghana", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon ya Kostarika", Symbol: "CRC"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia ya Pakistani", Symbol: "PKR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Faranga ya CFP", Symbol: "CFPF"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dola ya Hong Kong", Symbol: "HK$"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka ya Bangladeshi", Symbol: "BDT"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som ya Kyrgystani", Symbol: "KGS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real ya Brazil", Symbol: "R$"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Faranga ya Guinea", Symbol: "GNF"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca ya Macau", Symbol: "MOP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dola ya Jamaica", Symbol: "JMD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Koruna ya Jamhuri ya Cheki", Symbol: "CZK"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha ya Zambia", Symbol: "ZMW"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel ya Kambodia", Symbol: "KHR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni ya Tajikistani", Symbol: "TJS"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dola ya Bahamas", Symbol: "BSD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar ya Iraki", Symbol: "IQD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi ya Ghana", Symbol: "GHS"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu ya Vanuatu", Symbol: "VUV"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia ya Morisi", Symbol: "MUR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som ya Uzibekistani", Symbol: "UZS"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev ya Bulgaria", Symbol: "BGN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ugwiya ya Moritania", Symbol: "MRO"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metikali ya Msumbiji (1980–2006)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dola ya Barbados", Symbol: "BBD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Krone ya Denmaki", Symbol: "DKK"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia ya Ukrania", Symbol: "UAH"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Krona ya Aisilandi", Symbol: "ISK"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar ya Masedonia", Symbol: "MKD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa ya Panama", Symbol: "PAB"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dola ya Suriname", Symbol: "SRD"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso ya Ajentina", Symbol: "ARS"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint ya Hungaria", Symbol: "HUF"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial ya Yemeni", Symbol: "YER"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar ya Bahareni", Symbol: "BHD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano ya Bolivia", Symbol: "BOB"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Ruble ya Urusi", Symbol: "RUB"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: "SZL"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats ya Lativia", Symbol: "LVL"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram ya Armenia", Symbol: "AMD"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas ya Lithuania", Symbol: "LTL"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dola ya Belize", Symbol: "BZD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira ya Uturuki", Symbol: "TRY"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso ya Kolombia", Symbol: "COP"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel Mpya ya Israeli", Symbol: "₪"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Eskudo ya Cape Verde", Symbol: "CVE"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari ya Georgia", Symbol: "GEL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faranga ya Rwanda", Symbol: "RWF"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faranga ya Gine", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pauni ya Uingereza", Symbol: "£"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Pauni ya Sudani", Symbol: "SDG"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham ya Moroko", Symbol: "MAD"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar ya Kuwaiti", Symbol: "KWD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinari ya Libya", Symbol: "LYD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Convertible Mark ya Bosnia na Hezegovina", Symbol: "BAM"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Pauni ya Visiwa vya Falkland", Symbol: "FKP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "STD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina ya Papua New Guinea", Symbol: "PGK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Nuevo Sol ya Peru", Symbol: "PEN"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faranga ya Uswisi", Symbol: "CHF"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupiah ya Indonesia", Symbol: "IDR"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso ya Cuba Inayoweza Kubadilishwa", Symbol: "CUC"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik ya Mongolia", Symbol: "MNT"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar ya Yordani", Symbol: "JOD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin ya Aruba", Symbol: "AWG"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu ya Moldova", Symbol: "MDL"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit ya Malaysia", Symbol: "MYR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dola ya Visiwa vya Cayman", Symbol: "KYD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone", Symbol: "SLL"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip ya Laosi", Symbol: "LAK"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Faranga CFA BEAC", Symbol: "FCFA"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia ya Shelisheli", Symbol: "SCR"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dola ya Liberia", Symbol: "LRD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ya Afrika Kusini", Symbol: "ZAR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dola ya Namibia", Symbol: "NAD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso ya Meksiko", Symbol: "MX$"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa ya Eritrea", Symbol: "ERN"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat ya Azebaijani", Symbol: "AZN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Bir ya Uhabeshi", Symbol: "ETB"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha ya Malawi", Symbol: "MWK"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dola ya Marekani", Symbol: "US$"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Ruble ya Belarusi", Symbol: "BYR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dola ya Singapore", Symbol: "SGD"}}

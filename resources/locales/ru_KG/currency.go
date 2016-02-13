package ru_KG

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BRC": ut.Currency{Currency: "BRC", DisplayName: "Бразильское крузадо", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Южносуданский фунт", Symbol: "SSP"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Неизвестная или недействительная валюта", Symbol: "XXXX"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гайанский доллар", Symbol: "GYD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Белорусский рубль", Symbol: "BYR"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Турецкая лира (1922–2005)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катарский риал", Symbol: "QAR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонганская паанга", Symbol: "TOP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Таиландский бат", Symbol: "฿"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузинский лари", Symbol: "GEL"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Рубль СССР", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонезийская рупия", Symbol: "IDR"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Эквеле экваториальной Гвинеи", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Доллар Каймановых островов", Symbol: "KYD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Кина Папуа — Новой Гвинеи", Symbol: "PGK"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Джа", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Гаитянский гурд", Symbol: "HTG"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Европейская составная единица", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермудский доллар", Symbol: "BMD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Австралийский доллар", Symbol: "A$"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоанская тала", Symbol: "WST"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR франк", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Бразильское новое крузадо", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монгольский тугрик", Symbol: "MNT"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ганский седи (1979–2007)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кенийский шиллинг", Symbol: "KES"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминиканское песо", Symbol: "DOP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Оманский риал", Symbol: "OMR"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Афгани (1927–2002)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ганский седи", Symbol: "GHS"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Бразильский новый крузейро (1967–1986)", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Словенский толар", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Иранский риал", Symbol: "IRR"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Доллар Зимбабве", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малайзийский ринггит", Symbol: "MYR"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Иракский динар", Symbol: "IQD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Восточно-карибский доллар", Symbol: "EC$"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбийская квача", Symbol: "ZMW"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Польский злотый", Symbol: "PLN"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокканский дирхам", Symbol: "MAD"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Единица реальной стоимости Колумбии", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Испанская песета (конвертируемая)", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Франк Коморских островов", Symbol: "KMF"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомалийский шиллинг", Symbol: "SOS"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "расчетная единица европейского валютного соглашения (XBC)", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Российский рубль", Symbol: "₽"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR евро", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадосский доллар", Symbol: "BBD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапурский доллар", Symbol: "SGD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Европейская денежная единица", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Ботсванская пула", Symbol: "BWP"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Боливийское песо", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигерийская найра", Symbol: "NGN"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Доллар США следующего дня", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Испанская песета (А)", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинейский франк", Symbol: "GNF"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзанийский шиллинг", Symbol: "TZS"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Брунейский доллар", Symbol: "BND"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Ямайский доллар", Symbol: "JMD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Багамский доллар", Symbol: "BSD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Кубинское конвертируемое песо", Symbol: "CUC"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Греческая драхма", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сейшельская рупия", Symbol: "SCR"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Бельгийский франк", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Эскудо Португальской Гвинеи", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдавский лей", Symbol: "MDL"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутанский нгултрум", Symbol: "BTN"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Платина", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казахский тенге", Symbol: "KZT"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Французский тихоокеанский франк", Symbol: "CFPF"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Французский UIC-франк", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Ливийский динар", Symbol: "LYD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Эстонская крона", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Армянский драм", Symbol: "AMD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мавританская угия", Symbol: "MRO"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Таджикский сомони", Symbol: "TJS"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Чехословацкая твердая крона", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чешская крона", Symbol: "CZK"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Дирхам ОАЭ", Symbol: "AED"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Малагасийский ариари", Symbol: "MGA"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Югославский твердый динар", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Эфиопский быр", Symbol: "ETB"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Доллар Соломоновых островов", Symbol: "SBD"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Малагасийский франк", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексиканское песо", Symbol: "MX$"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непальская рупия", Symbol: "NPR"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Старый угандийский шиллинг", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Афганский афгани", Symbol: "AFN"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвайское песо", Symbol: "UYU"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Саудовский риал", Symbol: "SAR"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Сальвадорский колон", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Китайский юань", Symbol: "CN¥"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фунт Фолклендских островов", Symbol: "FKP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Камбоджийский риель", Symbol: "KHR"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Эквадорский сукре", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Мозамбикское эскудо", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербайджанский манат", Symbol: "AZN"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Южнокорейская вона", Symbol: "₩"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Динар Боснии и Герцеговины", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Люксембургский франк", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Австрийский шиллинг", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Южноафриканский рэнд", Symbol: "ZAR"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Карбованец (украинский)", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Иорданский динар", Symbol: "JOD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Датская крона", Symbol: "DKK"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Латвийский рубль", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Ангольская кванза", Symbol: "AOA"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Марокканский франк", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филиппинское песо", Symbol: "PHP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Доллар США", Symbol: "$"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Аргентинское песо", Symbol: "ARS"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Доллар Зимбабве (2009)", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Старый суданский фунт", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свазилендский лилангени", Symbol: "SZL"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбийский даласи", Symbol: "GMD"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжирский динар", Symbol: "DZD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Туркменский новый манат", Symbol: "ТМТ"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Мальтийский фунт", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Тиморское эскудо", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Доллар Фиджи", Symbol: "FJD"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либерийский доллар", Symbol: "LRD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Суринамский доллар", Symbol: "SRD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Перуанское инти", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "ЭКЮ (единица европейской валюты)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Российский рубль (1991–1998)", Symbol: "р."}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Костариканский колон", Symbol: "CRC"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Новый тайваньский доллар", Symbol: "NT$"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Немецкая марка", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Накфа", Symbol: "ERN"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Мальдивская руфия", Symbol: "MVR"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вьетнамский донг", Symbol: "₫"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македонский динар", Symbol: "MKD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "тестовый валютный код", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Аргентинский аустрал", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Арубанский флорин", Symbol: "AWG"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Условная расчетная единица Чили", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Ангольская кванза (1977–1990)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Английский фунт стерлингов", Symbol: "£"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Гондурасская лемпира", Symbol: "HNL"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувейтский динар", Symbol: "KWD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Бразильский крузейро (1990–1993)", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сирийский фунт", Symbol: "SYP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Леоне", Symbol: "SLL"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Йеменский динар", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панамское бальбоа", Symbol: "PAB"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Мексиканское серебряное песо (1861–1992)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Бразильский крузейро", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Мальтийская лира", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Финская марка", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландская крона", Symbol: "ISK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Японская иена", Symbol: "¥"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ирландский фунт", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбикский метикал", Symbol: "MZN"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Старый Сербский динар", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Франк Руанды", Symbol: "RWF"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Патака Макао", Symbol: "MOP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Эскудо Кабо-Верде", Symbol: "CVE"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Малийский франк", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Франк КФА ВЕАС", Symbol: "FCFA"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Новозеландский доллар", Symbol: "NZ$"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Южноафриканский рэнд (финансовый)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливийский боливиано", Symbol: "BOB"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Уругвайское старое песо (1975–1993)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Злотый", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Венгерский форинт", Symbol: "HUF"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Квача (замбийская) (1968–2012)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Болгарский лев", Symbol: "BGN"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Фунт острова Святой Елены", Symbol: "SHP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурундийский франк", Symbol: "BIF"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Киргизский сом", Symbol: "сом"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малавийская квача", Symbol: "MWK"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Итальянская лира", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Аргентинское песо (1983–1985)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Израильский фунт", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Андоррская песета", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Кипрский фунт", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Суринамский гульден", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Португальское эскудо", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Добра Сант-Томе и Принсипи", Symbol: "STD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Доллар США текущего дня", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Гонконгский доллар", Symbol: "HK$"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румынский лей", Symbol: "RON"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Ливанский фунт", Symbol: "LBP"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Нидерландский антильский гульден", Symbol: "ANG"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Шведская крона", Symbol: "SEK"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Старый мозамбикский метикал", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Бельгийский франк (конвертируемый)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Новый израильский шекель", Symbol: "₪"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Йеменский риал", Symbol: "YER"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразильский реал", Symbol: "R$"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вату Вануату", Symbol: "VUV"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Таджикский рубль", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Золото", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралтарский фунт", Symbol: "GIP"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чилийское песо", Symbol: "CLP"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Новый заир", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Восточногерманская марка", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венесуэльский боливар", Symbol: "VEF"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конголезский франк", Symbol: "CDF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвежская крона", Symbol: "NOK"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагуанская кордоба", Symbol: "NIO"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Лоти", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литовский лит", Symbol: "LTL"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвайский гуарани", Symbol: "PYG"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Суданский фунт", Symbol: "SDG"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Кубинское песо", Symbol: "CUP"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Никарагуанская кордоба (1988–1991)", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шри-Ланкийская рупия", Symbol: "LKR"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Испанская песета", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Словацкая крона", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албанский лек", Symbol: "ALL"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Французский золотой франк", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Мексиканская пересчетная единица (UDI)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "расчетная единица европейского валютного соглашения (XBD)", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Бельгийский франк (финансовый)", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Ангольская новая кванза (1990–2000)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Постоянная единица стоимости Эквадора", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Ангольская кванза реюстадо (1995–1999)", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Перуанский соль", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Белорусский рубль (1994–1999)", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Старый Румынский лей", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Югославский динар", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Серебро", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Суданский динар", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Гвинейская сили", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Боливийский мвдол", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбийское песо", Symbol: "COP"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Тунисский динар", Symbol: "TND"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Франк Джибути", Symbol: "DJF"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Египетский фунт", Symbol: "EGP"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Уругвайский песо (индекс инфляции)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Грузинский купон", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Конвертируемый франк Люксембурга", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швейцарский франк", Symbol: "CHF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мьянманский кьят", Symbol: "MMK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Доллар Намибии", Symbol: "NAD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахрейнский динар", Symbol: "BHD"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Лев", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хорватская куна", Symbol: "HRK"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Заир", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Турецкая лира", Symbol: "TRY"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Северокорейская вона", Symbol: "KPW"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистанская рупия", Symbol: "PKR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Угандийский шиллинг", Symbol: "UGX"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Югославский новый динар", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Сербский динар", Symbol: "RSD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перуанский новый соль", Symbol: "PEN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канадский доллар", Symbol: "CA$"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Литовский талон", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладешская така", Symbol: "BDT"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Песо Гвинеи-Бисау", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемальский кетсаль", Symbol: "GTQ"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Французский франк", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Старый азербайджанский манат", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвийский лат", Symbol: "LVL"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Франк КФА ВСЕАО", Symbol: "CFA"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Маврикийская рупия", Symbol: "MUR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Узбекский сум", Symbol: "UZS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украинская гривна", Symbol: "₴"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индийская рупия", Symbol: "₹"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Нидерландский гульден", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "СДР (специальные права заимствования)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Финансовый франк Люксембурга", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Доллар Тринидада и Тобаго", Symbol: "TTD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белизский доллар", Symbol: "BZD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Хорватский динар", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Родезийский доллар", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Конвертируемая марка Боснии и Герцеговины", Symbol: "BAM"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Туркменский манат", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаосский кип", Symbol: "LAK"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Палладий", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Венесуэльский боливар (1871–2008)", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "единица RINET-фондов", Symbol: ""}}

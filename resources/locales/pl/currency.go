package pl

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CRC": ut.Currency{Currency: "CRC", DisplayName: "colon kostarykański", Symbol: "CRC"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "korona czechosłowacka", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso kubańskie", Symbol: "CUP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "korona norweska", Symbol: "NOK"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dolar nowozelandzki", Symbol: "NZD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "złoty polski", Symbol: "zł"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni Suazi", Symbol: "SZL"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolańska (1977–1990)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "stary dinar serbski", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "jen japoński", Symbol: "JPY"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "funt maltański", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "frank belgijski", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "frank belgijski (finansowy)", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "frank komoryjski", Symbol: "KMF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dolar liberyjski", Symbol: "LRD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauretańska", Symbol: "MRO"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zair zairski", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistańska", Symbol: "PKR"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rubel rosyjski (1991–1998)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahrański", Symbol: "BHD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina Papua Nowa Gwinea", Symbol: "PGK"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso Guinea-Bissau", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazachskie", Symbol: "KZT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeryjska", Symbol: "NGN"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruwiański", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dolar wschodniokaraibski", Symbol: "EC$"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albański", Symbol: "ALL"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolańska Reajustado (1995–1999)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentyński", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubański", Symbol: "AWG"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol boliwijski", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambijska", Symbol: "ZMW"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bengalska", Symbol: "BDT"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "funt libański", Symbol: "LBP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicki", Symbol: "MZN"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "gulden holenderski", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dolar Wysp Salomona", Symbol: "SBD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "funt Wyspy Świętej Heleny", Symbol: "SHP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turecka", Symbol: "TRY"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentyńskie (1983–1985)", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliwijskie", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dolar belizeński", Symbol: "BZD"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "frank dżibutyjski", Symbol: "DJF"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambijskie", Symbol: "GMD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malezyjski", Symbol: "MYR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "gulden surinamski", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "nieznana waluta", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nowy zair zairski", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "funt egipski", Symbol: "EGP"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta hiszpańska (Konto A)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitańskie", Symbol: "HTG"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "szyling somalijski", Symbol: "SOS"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoańska", Symbol: "WST"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "frank złoty francuski", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armeński", Symbol: "AMD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bhutański", Symbol: "BTN"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon gruziński larit", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "lej mołdawski", Symbol: "MDL"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial jemeński", Symbol: "YER"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "nowy dinar jugosławiański", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham arabski", Symbol: "AED"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "funt szterling", Symbol: "GBP"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar chorwacki", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "frank kongijski", Symbol: "CDF"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso meksykańskie", Symbol: "MXN"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "korona szwedzka", Symbol: "SEK"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbowaniec ukraiński", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo Gwinea Portugalska", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel kambodżański", Symbol: "KHR"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rupia malediwska", Symbol: "MVR"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brazylijski", Symbol: "R$"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "rubel białoruski (1994–1999)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marka fińska", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nowy sol peruwiański", Symbol: "PEN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "funt sudański", Symbol: "SDG"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar jugosławiański wymienny", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lir włoski", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "frank belgijski (zamienny)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghański", Symbol: "GHS"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli gwinejskie", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talon litewski", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca Makau", Symbol: "MOP"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dolar rodezyjski", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rubel tadżycki", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marka zamienna Bośni i Hercegowiny", Symbol: "BAM"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuwejcki", Symbol: "KWD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawska", Symbol: "MWK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rubel rosyjski", Symbol: "RUB"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "srebro", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "frank CFA", Symbol: "CFA"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentyńskie", Symbol: "ARS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "ECU", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "juan chiński", Symbol: "CNY"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dolar namibijski", Symbol: "NAD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "szyling ugandyjski", Symbol: "UGX"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libijski", Symbol: "LYD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "frank malijski", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmeński (1993–2009)", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunezyjski", Symbol: "TND"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia indyjska", Symbol: "INR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial katarski", Symbol: "QAR"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dolar Trynidadu i Tobago", Symbol: "TTD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "korona estońska", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kiat birmański", Symbol: "MMK"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nowe cruzado brazylijskie", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele gwinejskie Gwinei Równikowej", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "szyling kenijski", Symbol: "KES"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruwiański", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorskie", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "nowa kwanza angolańska (1990–2000)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "wschodnia marka wschodnioniemiecka", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "som uzbecki", Symbol: "UZS"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "frank CFA BEAC", Symbol: "FCFA"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "testowy kod waluty", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dolar brunejski", Symbol: "BND"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "frank szwajcarski", Symbol: "CHF"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta hiszpańska", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marokański", Symbol: "MAD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudyjski", Symbol: "SAR"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "europejska jednostka rozrachunkowa (XBD)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lew bułgarski wymienny", Symbol: "BGL"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonezyjska", Symbol: "IDR"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "frank malgaski", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambijska (1968–2012)", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brazylijskie", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brazylijskie (1990–1993)", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint węgierski", Symbol: "HUF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dolar jamajski", Symbol: "JMD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "łat łotewski", Symbol: "LVL"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierraleoński", Symbol: "SLL"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand południowoafrykański", Symbol: "ZAR"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "szyling austriacki", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudański", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "frank CFP", Symbol: "CFPF"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiopski", Symbol: "ETB"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirgiski", Symbol: "KGS"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "cordoba nikaraguańska", Symbol: "NIO"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "jednostka emisji euroobligacji", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "frank gwinejski", Symbol: "GNF"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia lankijska", Symbol: "LKR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dolar singapurski", Symbol: "SGD"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "szyling ugandyjski (1966–1987)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "pallad", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "gulden antylski", Symbol: "ANG"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rubel łotewski", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia maurytyjska", Symbol: "MUR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "funt syryjski", Symbol: "SYP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo zielonoprzylądkowe", Symbol: "CVE"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "frank luksemburski", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macedoński", Symbol: "MKD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nikaraguańska (1988–1991)", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panamski", Symbol: "PAB"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "lej rumuński (1952–2006)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lew bułgarski", Symbol: "BGN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algierski", Symbol: "DZD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nowy szekel izraelski", Symbol: "ILS"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "europejska jednostka rozrachunkowa (XBC)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azerbejdżański", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilijskie", Symbol: "CLP"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dolar hongkoński", Symbol: "HKD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "funt izraelski", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "lew bułgarski (1879–1952)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marka niemiecka", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghańskie (1979–2007)", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omański", Symbol: "OMR"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "specjalne prawa ciągnienia", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iracki", Symbol: "IQD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari gruzińskie", Symbol: "GEL"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipińskie", Symbol: "PHP"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "złoto", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "lew bułgarski socjalistyczny", Symbol: "BGM"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "funt falklandzki", Symbol: "FKP"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltańska", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dolar Zimbabwe (2008)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azerski", Symbol: "AZN"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre ekwadorski", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dolar gujański", Symbol: "GYD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "lit litewski", Symbol: "LTL"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colon salwadorski", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rubel białoruski", Symbol: "BYR"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "frank francuski", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna chorwacka", Symbol: "HRK"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "europejska jednostka monetarna", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmański", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordański", Symbol: "JOD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambickie", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongolski", Symbol: "MNT"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hrywna ukraińska", Symbol: "UAH"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dolar Barbadosu", Symbol: "BBD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal gwatemalski", Symbol: "GTQ"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu Vanuatu", Symbol: "VUV"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "cruzeiro novo brazylijskie (1967–1986)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "boliwar wenezuelski", Symbol: "VEF"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platyna", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dolar Zimbabwe (2009)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "frank burundyjski", Symbol: "BIF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "funt cypryjski", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "funt gibraltarski", Symbol: "GIP"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "korona islandzka", Symbol: "ISK"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "frank marokański", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbski", Symbol: "RSD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht tajski", Symbol: "THB"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadżyckie", Symbol: "TJS"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turecka (1922–2005)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dolar fidżi", Symbol: "FJD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti Lesoto", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "metical Mozambik", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "frank ruandyjski", Symbol: "RWF"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia seszelska", Symbol: "SCR"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani", Symbol: "AFN"}, "RON": ut.Currency{Currency: "RON", DisplayName: "lej rumuński", Symbol: "RON"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "funt sudański (1957–1998)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso urugwajskie (1975–1993)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong wietnamski", Symbol: "VND"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa erytrejska", Symbol: "ERN"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugalskie", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmeński", Symbol: "TMT"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dolar bahamski", Symbol: "BSD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso kolumbijskie", Symbol: "COP"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta hiszpańska (konto wymienne)", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar słoweński", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laotański", Symbol: "LAK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalska", Symbol: "NPR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra Wysp Świętego Tomasza i Książęcej", Symbol: "STD"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "UIC-frank francuski", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bośni i Hercegowiny", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dolar bermudzki", Symbol: "BMD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso urugwajskie", Symbol: "UYU"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brazylijskie", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won północnokoreański", Symbol: "KPW"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dolar kajmański", Symbol: "KYD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgaski", Symbol: "MGA"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso srebrne meksykańskie (1861–1992)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolańska", Symbol: "AOA"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso kubańskie wymienialne", Symbol: "CUC"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachma grecka", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial irański", Symbol: "IRR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dolar surinamski", Symbol: "SRD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nowy dolar tajwański", Symbol: "TWD"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "funt irlandzki", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "dolar amerykański", Symbol: "USD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar jemeński", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorska", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswańska", Symbol: "BWP"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "boliwar wenezuelski (1871–2008)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dolar kanadyjski", Symbol: "CAD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "korona duńska", Symbol: "DKK"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominikańskie", Symbol: "DOP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dolar australijski", Symbol: "AUD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "złoty polski (1950–1995)", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongijska", Symbol: "TOP"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "szyling tanzański", Symbol: "TZS"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "korona czeska", Symbol: "CZK"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduraska", Symbol: "HNL"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won południowokoreański", Symbol: "KRW"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guarani paragwajskie", Symbol: "PYG"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rubel radziecki", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "korona słowacka", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "funt południowosudański", Symbol: "SSP"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand południowoafrykański (finansowy)", Symbol: ""}}

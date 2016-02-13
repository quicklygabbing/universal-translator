package et

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbekistani somm", Symbol: "UZS"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Prantsuse UIC-frank", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombia peeso", Symbol: "COP"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iraani riaal", Symbol: "IRR"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mosambiigi metikal (1980–2006)", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Jaapani jeen", Symbol: "¥"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldiivi ruupia (1947–1981)", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi Araabia riaal", Symbol: "SAR"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Iisraeli nael", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruu soll (1863–1965)", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipiini peeso", Symbol: "PHP"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovaki kroon", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidadi ja Tobago dollar", Symbol: "TTD"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Bosnia ja Hertsegoviina uus dinaar (1994–1997)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kuuba konverteeritav peeso", Symbol: "CUC"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmari kjatt", Symbol: "MMK"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Aserbaidžaani manat (1993–2006)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreini dinaar", Symbol: "BHD"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Iiri nael", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Malta liir", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Euroopa rahaline arvestusühik (XBD)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia ja Hertsegoviina konverteeritav mark", Symbol: "BAM"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Soome mark", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Gruusia lari", Symbol: "GEL"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svaasimaa lilangeni", Symbol: "SZL"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Saint Helena nael", Symbol: "SHP"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jeemeni riaal", Symbol: "YER"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brasiilia krusado", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Leedu litt", Symbol: "LTL"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaragua kordoba (1988–1991)", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistani afgaani", Symbol: "AFN"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Rumeenia leu (1952–2006)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indoneesia ruupia", Symbol: "IDR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaani riaal", Symbol: "OMR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katari riaal", Symbol: "QAR"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Kesk-Aafrika CFA frank", Symbol: "FCFA"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Euroopa rahaline arvestusühik (XBC)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brasiilia kruseiro (1993–1994)", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapuri dollar", Symbol: "SGD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leone leoone", Symbol: "SLL"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadori sukre", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau peeso", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyana dollar", Symbol: "GYD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Hollandi kulden", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Araabia Ühendemiraatide dirhem", Symbol: "AED"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi frank", Symbol: "BIF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Alžeeria dinaar", Symbol: "DZD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "EURCO", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Ida-Kariibi dollar", Symbol: "EC$"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "plaatina", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwe dollar (2009)", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorra peseeta", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghana sedi (1979–2007)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poola zlott", Symbol: "PLN"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Jugoslaavia uus dinaar (1994–2002)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Põhja-Korea vonn", Symbol: "KPW"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Mali frank", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Türgi liir (1922–2005)", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djibouti frank", Symbol: "DJF"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Itaalia liir", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monaco frank", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "vääringute testkood", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldiivi ruupia", Symbol: "MVR"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Sambia kvatša", Symbol: "ZMW"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angola kvanza (1977–1990)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Eesti kroon", Symbol: "kr"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Iisraeli uus seekel", Symbol: "₪"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tai baat", Symbol: "฿"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afganistani afgaani (1927–2002)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Saksa mark", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghana sedi", Symbol: "GHS"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanada dollar", Symbol: "CA$"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuveidi dinaar", Symbol: "KWD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruu inti", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "NSVL-i rubla", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuela boliivar (1871–2008)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentina peeso", Symbol: "ARS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasiilia reaal", Symbol: "R$"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Birma kjatt", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Islandi kroon", Symbol: "ISK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepali ruupia", Symbol: "NPR"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Lõuna-Aafrika rand", Symbol: "ZAR"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo frank", Symbol: "CDF"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Hispaania peseeta", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Iisraeli seekel (1980–1985)", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaragua kordoba", Symbol: "NIO"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Sloveenia tolar", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnami dong", Symbol: "₫"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidži dollar", Symbol: "FJD"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Suriname kulden", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Türkmenistani manat", Symbol: "TMT"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinea syli", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Libeeria dollar", Symbol: "LRD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panama balboa", Symbol: "PAB"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Lõuna-Korea vonn (1945–1953)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laose kiip", Symbol: "LAK"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldova leu", Symbol: "MDL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruu uus soll", Symbol: "PEN"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikaani peeso", Symbol: "DOP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambia dalasi", Symbol: "GMD"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoori frank", Symbol: "KMF"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba kulden", Symbol: "AWG"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodža riaal", Symbol: "KHR"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawi kvatša", Symbol: "MWK"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbia dinaar", Symbol: "RSD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Saalomoni Saarte dollar", Symbol: "SBD"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Rahvusvahelise Valuutafondi arvestusühik", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "pallaadium", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliivia boliviaano", Symbol: "BOB"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahama dollar", Symbol: "BSD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Kaimanisaarte dollar", Symbol: "KYD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Lääne-Aafrika CFA frank", Symbol: "CFA"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoslaavia konverteeritav dinaar (1990–1992)", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Sambia kvatša (1968–2012)", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliivia boliviaano (1863–1963)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Suurbritannia naelsterling", Symbol: "£"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnami dong (1978–1985)", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kõrgõzstani somm", Symbol: "KGS"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kasahstani tenge", Symbol: "KZT"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Venemaa rubla", Symbol: "RUB"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé ja Príncipe dobra", Symbol: "STD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "kuld", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeenia dramm", Symbol: "AMD"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgaaria püsiv leev", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Valgevene uus rubla (1994–1999)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumeenia leu", Symbol: "RON"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgaaria leev (1879–1952)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tšehhi kroon", Symbol: "CZK"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Islandi kroon (1918–1981)", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu vatu", Symbol: "VUV"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "hõbe", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwe dollar (1980–2008)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Valgevene rubla", Symbol: "BYR"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaisia ringgit", Symbol: "MYR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seišelli ruupia", Symbol: "SCR"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angola reformitud kvanza, 1995–1999", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentina peeso (1983–1985)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraagi dinaar", Symbol: "IQD"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Malta nael", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeeria naira", Symbol: "NGN"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angola kvanza", Symbol: "AOA"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Kreeka drahm", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Horvaatia dinaar", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoa taala", Symbol: "WST"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Serbia dinaar (2002–2006)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudaani nael", Symbol: "SDG"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgia arveldusfrank", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei dollar", Symbol: "BND"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiuse ruupia", Symbol: "MUR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaalia šilling", Symbol: "SOS"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taiwani dollar", Symbol: "NT$"}, "USD": ut.Currency{Currency: "USD", DisplayName: "USA dollar", Symbol: "$"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeshi taka", Symbol: "BDT"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize’i dollar", Symbol: "BZD"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rodeesia dollar", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Euroopa rahaühik", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandi saarte nael", Symbol: "FKP"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Paapua Uus-Guinea kina", Symbol: "PGK"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadžikistani rubla", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luksemburgi konverteeritav frank", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaania ugia", Symbol: "MRO"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timori eskuudo", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguay guaranii", Symbol: "PYG"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoolia tugrik", Symbol: "MNT"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luksemburgi frank", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Maroko dirhem", Symbol: "MAD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambiigi metikal", Symbol: "MZN"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukraina grivna", Symbol: "UAH"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentina austral", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kuuba peeso", Symbol: "CUP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Prantsuse frank", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Ida-Saksa mark", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norra kroon", Symbol: "NOK"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudaani dinaar (1992–2007)", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Süüria nael", Symbol: "SYP"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jeemeni dinaar", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Hollandi Antillide kulden", Symbol: "ANG"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgia konverteeritav frank", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Hiina jüaan", Symbol: "CN¥"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Uganda šilling (1966–1987)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezuela boliivar", Symbol: "VEF"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cabo Verde eskuudo", Symbol: "CVE"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaskar frank", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tuneesia dinaar", Symbol: "TND"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugali eskuudo", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguay peeso", Symbol: "UYU"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Lõuna-Sudaani nael", Symbol: "SSP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskari ariari", Symbol: "MGA"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistani ruupia", Symbol: "PKR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudaani nael (1957–1998)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Rootsi kroon", Symbol: "SEK"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albaania lekk", Symbol: "ALL"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswana pula", Symbol: "BWP"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Ungari forint", Symbol: "HUF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Liibüa dinaar", Symbol: "LYD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frank", Symbol: "CFPF"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentina peeso (1881–1970)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaaria leev", Symbol: "BGN"}, "INR": ut.Currency{Currency: "INR", DisplayName: "India ruupia", Symbol: "₹"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "määramata rahaühik", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemala ketsaal", Symbol: "GTQ"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Lõuna-Korea vonn", Symbol: "₩"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "eküü", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macau pataaka", Symbol: "MOP"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Prantsuse kuldfrank", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Austria šilling", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadose dollar", Symbol: "BBD"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugali Guinea eskuudo", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaania dinaar", Symbol: "JOD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lanka ruupia", Symbol: "LKR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namiibia dollar", Symbol: "NAD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgia frank", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rica koloon", Symbol: "CRC"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkongi dollar", Symbol: "HK$"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnia ja Hertsegoviina dinaar (1992–1994)", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Boliivia peeso", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Suriname dollar", Symbol: "SRD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurase lempiira", Symbol: "HNL"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Liibanoni nael", Symbol: "LBP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Austraalia dollar", Symbol: "AU$"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda dollar", Symbol: "BMD"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrea nakfa", Symbol: "ERN"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brasiilia uus kruseiro (1967–1986)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etioopia birr", Symbol: "ETB"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea frank", Symbol: "GNF"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angola kvanza (1990–2000)", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Küprose nael", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egiptuse nael", Symbol: "EGP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mehhiko peeso", Symbol: "MX$"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "El Salvadori koloon", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brasiilia kruseiro (1942–1967)", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukraina karbovanets", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Tšiili eskuudo", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Taani kroon", Symbol: "DKK"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedoonia dinaar", Symbol: "MKD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Türkmenistani manat (1993–2009)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansaania šilling", Symbol: "TZS"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutani ngultrum", Symbol: "BTN"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Šveitsi frank", Symbol: "CHF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Läti latt", Symbol: "LVL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwanda frank", Symbol: "RWF"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Türgi liir", Symbol: "TRY"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Läti rubla", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Maroko frank", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Venemaa rubla (1991–1998)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Horvaatia kuna", Symbol: "HRK"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Poola zlott (1950–1995)", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadžikistani somoni", Symbol: "TJS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbaidžaani manat", Symbol: "AZN"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brasiilia kruseiro (1990–1993)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Tšiili peeso", Symbol: "CLP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keenia šilling", Symbol: "KES"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho loti", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mehhiko peeso (1861–1992)", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mosambiigi eskuudo", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albaania lekk (1946–1965)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltari nael", Symbol: "GIP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaica dollar", Symbol: "JMD"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Makedoonia dinaar (1992–1993)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "USA järgmise päeva dollar", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Sairi zaire", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonga pa’anga", Symbol: "TOP"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguay peeso (1975–1993)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "USA sama päeva dollar", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwe dollar (2008)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haiti gurd", Symbol: "HTG"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Uus-Meremaa dollar", Symbol: "NZ$"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda šilling", Symbol: "UGX"}}

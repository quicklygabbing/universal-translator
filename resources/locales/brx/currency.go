package brx

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"KPW": ut.Currency{Currency: "KPW", DisplayName: "ऊत्तर कोरियाई वोन", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "युगाँडाई शीलींग (1996–1987)", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "कोलम्बियाई पेसो", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "एक्वादोर सुक्रे", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "कंबोडिया का रिएल", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "नाईजीरीयाई नाईरा", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "स्वीडन क्रोना", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ताईवानी नया डॉलर", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "अमरिकी डॉलर (इसी दिन का)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "बुरुंदी फ्राँ", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "गाँबिया का दलासी", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "न्यूज़ीलैंड डॉलर", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "फ्रानसीसी सुवर्ण फ्राँ", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "सर्बियाई दिनार", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "एम्यु", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "चिली पेसो", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "पेरूवाई ईंटी", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "ऑस्ट्रियन शीलींग", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "बोसनिया हेर्ज़ेगोविना का दीनार", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "रां", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "नेदरलैण्ड गीलडर", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ऊज़बेक सुम", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "आज़रबैजानी मनात (1993–2006)", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "बेलारुसी नया रूबल (194–1999)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "यूरोपी एकाऊंट का युनीट (एक्स बी डी)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "युगोस्लावी कनवर्टीबल दीनार", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "एंगोला क्वानज़ा", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "नेपाली रुपी", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "संयुक्त अरब अमीरात का दिर्हाम", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "युक्रेनी ह्रीवनिया", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "इस्राइली पौंड", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "कोमोरो का फ्राँ", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "मोरिशियस का रूपी", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "नीकारागुआई कोर्दोबा", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "ताजीक़ीस्तानी रूबल", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "एंगोला क्वानज़ा (1977–1990)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "डॉईच मार्क", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "सोलोमन द्वीप का डॉलर", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "अमरिकी डॉलर", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "ऊरुगुए का पेसो आन ऊनीदादोस ईंदेक्सादास", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "नेदरलैण्ड एन्टीलीज़ का गील्डर", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "बेल्जियन फ्राँ कनवर्टीबल", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "तान्ज़ेनियाई शीलींग", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "बहामा डॉलर", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "लुक्ज़मबुर्गी फ्राँ", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "जॉर्जिया का लारी", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ऐक्वाटरी गीनी एक्वेले गीनीआना", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "रोमानियाई पुरानी ल", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "(सीएफ़ए) फ्रानसीसी फेदरेशनी फ्राँ", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "युगोस्लावी हार्ड दीनार", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "चीनी युआन रेनमीनबी", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "चेक गनतंत्र का कोरुना", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "मोज़ांबीक एस्कुएदो", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "पैलेडियम", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "सुदानी पुराना पौंड", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "स्लोवेनियाई तोलार", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "श्री लंका रूपी", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "लिथुआनियाई लिता", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "डेनमार्क का क्रोन", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "जीब्रालटर का पौण्ड", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "बारबादोस डॉलर", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "बरमुडी डॉलर", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ओमानी रियाल", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "पाकिस्तानी रुपया", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "सुरीनाम डॉलर", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "एल सालवादर कोलोन", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "क़ुवैती दीनार", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "नॉर्वे का क्रोन", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "फीनीश मार्क्का", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "माल्टी लीरा", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "मालवी क्वाचा", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "थाई बाह्ट", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "अमरिकी डॉलर (अगले दिन का)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "स्विस फ़्रैंक", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "डॉमीनीकन पेसो", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "मादागास्करी फ्राँ", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "सेशेल रूपी", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "तुर्कमेनीस्तानी मानाट", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "यूरोपी एकाऊंट का युनीट (एक्स बी सी)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "बोलिवियानो", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "बेलीज़ डॉलर", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "रीनैट फंड्स", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "लाटवियाई लाट्स", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "सोमाली शीलींग", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "लुक्ज़मबुर्गी कनवर्टीबल फ्राँ", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "माकाव पाताचा", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "जमाईका का डॉलर", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "लिथुआनियाई टालोनास", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "दक्षिण कोरियाई वोन", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "जीबुती फ्राँ", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ईथिओपिया का बीर्र", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "फाल्कलैण्ड द्वीप पौण्ड", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "ताजीक़ीस्तानी सोमोनी", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "भुतान का नगुलत्रुम", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "युरो", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "फ्रानसीसी फेदेरेशनी फ्राँ", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "बल्गेरियाई हार्ड लेव", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ग्वातेमाला क़्वेत्ज़ाल", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "लिबियाई दीनार", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "सोना", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "आरमिनियाई दिर्हाम", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "युनानी द्राखमा", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "सर्बिया का डीनार", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "काप वेर्दे का एस्कुदो", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ब्रज़ीली रेयाल", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "बोट्सवाना का पुलाट", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "हंगेरियाई फ़ोरिण्ट", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "मौंगोलीयाई तुग्रीक", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ऊरुगुए का ऊरुगुआयो पेसो", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "ब्रज़ीली क्रुज़ाडो", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "बेलारुसी रूबल", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "इतली का लीरा", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "कैमान द्वीप का डॉलर", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "यमनी रीयाल", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "इस्राइली शेकेल", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "पोलिश ज़्लॉटी", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ईस्ट जर्मन ओस्टमार्क", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "लाओस का कीप", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "नमीबिया डॉलर", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "वेनेज़ुएलाई बोलिवार", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "आईरलैण्ड का युरो", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "चैकोस्लोवाकिय हार्ड कोरुना", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "सीयेरा लीयोनेई लीयोने", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "तीमोरी एस्कुदो", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "ज़ाम्बियाई क्वाचा (1968–2012)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "ज़ीम्बाबवेई डॉलर", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ऑस्ट्रेलियन डॉलर", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "रूसी रूबल", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "आल्बेनिया का लेक", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ईराक़ी दीनार", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "स्वाज़ीलैण्ड लीलांगेनी", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "सिंगापुर डॉलर", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "लीबेरियाई डॉलर", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "मौरिटानी ऊगुया", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "फ्राँसीसी फ्राँ", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "यमनी दीनार", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "बोलिवियाई डॉलर", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "हीस्पानी पेसेता (ए अकाऊँट)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "रोमानियाई ल", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "स्लोवाकी कोरुना", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "क्युबा का पेसो", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "घाना चेदी (1979–2007)", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "युगाँडाई शीलींग", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "ऊरुगुए का पेसो (1975–1993)", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "वेनेज़ुएलाई बोलिवार (1871–2008)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "स्पेशियल ड्राईंग राईट्स", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "ज़ाईरी ज़ाईर", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "अफ़ग़ानी", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "मालदीव द्वीप का रूफिया", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "त्युनीसी दीनर", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "ऐन्डोरा का पेसेता", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "सोवियत रूबल", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "पापुआ न्यु गीनी का कीना", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "बोलिवियाई पेसो", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "कॉंगोलीज़ फ्राँ", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "आईरलैण्ड का फ़्रैंक", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "लाटवियाई रूबल", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "मादागास्करी आरिआरी", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "अफ़ग़ानी 1927–2002", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "ब्रुनई डॉलर", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "क्रोएशियाई कुना", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "आईरीश पौंड", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ईरानी रीयाल", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "त्रीनीदाद एवं टोबागो डॉलर", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "वनुआटु वटु", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "जॉर्जिया का कुपोन लारीत", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "हाँग काँग डॉलर", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "क़ाज़ाख़स्तान तेंगे", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "तुर्की लीरा", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "घाना चेदी", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ईस्ट करिबियन डॉलर", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "बल्गेरियाई लेव", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "हीस्पानी पेसेता", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "बाहरैनी दीनार", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "मेक्सिकन रजती पेसो (1861–1992)", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "एंगोला नया क्वानज़ा (1990–20000)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "बेल्जियन फ्राँ", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "वीयतनामी डॉंग", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "एंगोला क्वानज़ा सुधारीत (1995–1999)", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "सुदानी पौंड", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "ज़ाईरी नये ज़ाईर", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ज़ाम्बियाई क्वाचा", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "कनेडियन डॉलर", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "पेरुवाई सोल", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "म्यानमारी क्याट", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "ब्रज़ीली नया क्रुज़ाडो", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "इण्डोनेशियाई रुपिया", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "हॉंडुरास लेंपीरा", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "बेल्जियन फ्राँ फिनानसीयल (वित्तीय)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "ब्रज़ीली क्रुज़ेरो", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "गीनी फ्राँ", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "पुरतुगी गीनी का एस्कुएदो", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "मोज़ांबीक मेतीकाल", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "पेरुवाई नया सोल", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "साईप्रस का पाऊंड", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "हीस्पानी पेसेता (कनवर्टीबल अकाऊँट)", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "जोर्डनी दीनार", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "लसोथो का लोटी", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "मेक्सिकन पेसो", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "मोज़ांबीक पुराना मेतीकाल", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "युक्रेनी कार्बोवानेत्ज़", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "टेस्टींग करनसी कोड", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "एक्वादोर युनीदाद दे वालोर कॉनस्तांते", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "आईसलैण्ड क्रोना", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "मिस्री पाउण्ड", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "फ़िजी का डॉलर", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "किनियाई शीलींग", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "मोल्डोवियाई ल", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "आज़रबैजानी मनात", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "कॉस्टा रीका का कोलोन", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "सुदानी पुराना डॉलर", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "सेंट हेलीना पौंड", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "अज्ञात या अवैध मुद्रा", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "दक्षिण अफ़्रीकी रॅण्ड", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "गीनी सीली", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "लीबानी पौंड", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "माली का फ्राँ", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "साँव तोमे एवं प्रीन्सीपे का डोब्रा", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "टॉंगा पाईंगा", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "वेस्टर्न समोआ ताला", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "यूरोपी मुद्रा (एक्यु)", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "चीली का ऊनीदादेस द फोमेंटो", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "क्रोएशियाई दीनार", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "फ्रानसीसी फेदेरेशनी बीसीएआओ फ्राँ", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "पनामा का बालबोआ", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "फ्रानसीसी युआईसी फ्राँ", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "पुरतुगी एस्कुदो", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "रुआँदा फ्राँ", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "बोसनिया हेर्ज़ेगोविना कनवर्टीबल मार्क", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "फ़िलिपीन का पेसो", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "प्लैटीनम", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "आल्जीरी दीनार", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "रोडेशियाई डॉलर", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "जापानी येन", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "लुक्ज़मबुर्गी वीत्ती फ्राँ", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "मेक्सिकन युनीदाद द ईनवेरसिओन (युडीआई)", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "सीरियाई पौंड", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "बर्मी (म्यानमारी) क्याट", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "हाईती गुर्द", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "गुयाना डॉलर", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "नीकारागुआई सुवर्ण कोर्दोबा", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "रजत", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "अर्जेण्टीनी पेसो", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "ब्रज़ीली क्रुज़ेरो (190–1993)", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "माल्टी पौंड", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "मलेशियन रिंगित", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "युनीदाद द वालोर रेआल", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ब्रितन का पौण्ड स्टर्लिग", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ऐरित्रीया का नाफ़का", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "गीनी बिस्साऊ का पेसो", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "तुर्की नया लीरा", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "एक्यु", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "युगोस्लावी नोवीय (नये) दीनार", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "ब्रज़ीली नया क्रुज़ेरो (1967–1986)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "एस्टोनियाई क्रून", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "क़ीर्ग़ीज़स्तानी सोम", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "मोरक्किय फ्राँ", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "क़तारी रीयाल", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "अर्जेण्टीनी ओस्ट्राल", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "अर्जेण्टीनी पेसो (1983–1985)", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "पारागुऐई गुआरानी", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "मोरक्किय दिर्हाम", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "मसेदोनियाई दीनार", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "साउदी रियाल", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "सुरीनाम गील्डर", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "अरुबा गील्डर", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "बांगलादेश टका", Symbol: ""}}

package th

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AZM": ut.Currency{Currency: "AZM", DisplayName: "มานัตอาเซอร์ไบจาน (1993–2006)", Symbol: "AZM"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "มาร์กบอสเนีย-เฮอร์เซโกวีนา", Symbol: "BAM"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "มฟดอลโบลิเวีย", Symbol: "BOV"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "ฮาร์ดโครูนาเช็กโกสโลวัก", Symbol: "CSK"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "แนกฟาเอริเทรีย", Symbol: "ERN"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ดอลลาร์ไต้หวันใหม่", Symbol: "NT$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "ดอลลาร์แคนาดา", Symbol: "CA$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "โกลองคอสตาริกา", Symbol: "CRC"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "ดีนาร์ลิเบีย", Symbol: "LYD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "เอสคูโดโมซัมบิก", Symbol: "MZE"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "กิลเดอร์ซูรินาเม", Symbol: "SRG"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "เปโซอาร์เจนตินา (1881–1970)", Symbol: "ARM"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "โบลิเวียโนโบลิเวีย", Symbol: "BOB"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "ฟรังก์ดับเบิลยูไออาร์", Symbol: "CHW"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "ไซลีกินี", Symbol: "GNS"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "ลีตัสลิทัวเนีย", Symbol: "LTL"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "เมติคัลโมซัมบิกเก่า", Symbol: "MZM"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "ฟรังก์ซีเอฟพี", Symbol: "CFPF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ดอลลาร์ออสเตรเลีย", Symbol: "AU$"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "เอสคูโดชิลี", Symbol: "CLE"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "หยวนจีน", Symbol: "CN¥"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "โคลอนเอลซัลวาดอร์", Symbol: "SVC"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "ปอนด์ซีเรีย", Symbol: "SYP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "เดอร์แฮมสหรัฐอาหรับเอมิเรตส์", Symbol: "AED"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "ฟรังก์คองโก", Symbol: "CDF"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "โครูนาสาธารณรัฐเช็ก", Symbol: "CZK"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "ฟรังก์จิบูตี", Symbol: "DJF"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "รูเบิลทาจิกิสถาน", Symbol: "TJR"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "อัฟกานีอัฟกานิสถาน (1927–2002)", Symbol: "AFA"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "เลกแอลเบเนีย", Symbol: "ALL"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "ชิลลิงออสเตรีย", Symbol: "ATS"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ปอนด์สเตอร์ลิง (สหราชอาณาจักร)", Symbol: "£"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "ฟรังก์คอโมโรส", Symbol: "KMF"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "พาแองกาตองกา", Symbol: "TOP"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ฟรังก์เซฟาธนาคารรัฐแอฟริกากลาง", Symbol: "FCFA"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "วอนเกาหลีเหนือ", Symbol: "KPW"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "ดีแรห์มโมร็อกโก", Symbol: "MAD"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "คอนเวอร์ทิเบิลฟรังก์ลักเซมเบิร์ก", Symbol: "LUC"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "รูฟิยามัลดีฟส์", Symbol: "MVR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "เรียลโอมาน", Symbol: "OMR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "บาทไทย", Symbol: "THB"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ทอง", Symbol: "XAU"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "ตากาบังกลาเทศ", Symbol: "BDT"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "ดอลลาร์เบลีซ", Symbol: "BZD"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "เอสคูโดกินีโปรตุเกส", Symbol: "GWE"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "โครนาสวีเดน", Symbol: "SEK"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "ดีนาร์ปฏิรูปยูโกสลาเวีย (1992–1993)", Symbol: "YUR"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ปอนด์ไอริช", Symbol: "IEP"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "ลีราอิตาลี", Symbol: "ITL"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "ดีนาร์เซอร์เบีย", Symbol: "RSD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "ปอนด์เซนต์เฮเลนา", Symbol: "SHP"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "โบลิวาร์เวเนซุเอลา (1871–2008)", Symbol: "VEB"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "พัลเลเดียม", Symbol: "XPD"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "แซร์คองโก", Symbol: "ZRZ"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ดอลลาร์ไลบีเรีย", Symbol: "LRD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "คอร์โดบานิการากัว", Symbol: "NIC"}, "USN": ut.Currency{Currency: "USN", DisplayName: "ดอลลาร์สหรัฐ (วันถัดไป)", Symbol: "USN"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "หน่วยบัญชียุโรป [XBC]", Symbol: "XBC"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "แพลตินัม", Symbol: "XPT"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "ดอลลาร์ฟิจิ", Symbol: "FJD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "ดีนาร์มาซิโดเนีย", Symbol: "MKD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "เมติคัลโมซัมบิก", Symbol: "MZN"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "สิทธิถอนเงินพิเศษ", Symbol: "XDR"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "ดอลลาร์บาร์เบโดส", Symbol: "BBD"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "จ๊าดพม่า", Symbol: "BUK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "รูเบิลรัสเซีย", Symbol: "RUB"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ซอมอุซเบกิสถาน", Symbol: "UZS"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "หน่วยบัญชียุโรป [XBD]", Symbol: "XBD"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "หน่วยสกุลเงินยุโรป", Symbol: "XEU"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "เปโซอาร์เจนตินา (1983–1985)", Symbol: "ARP"}, "BND": ut.Currency{Currency: "BND", DisplayName: "ดอลลาร์บรูไน", Symbol: "BND"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "ปอนด์หมู่เกาะฟอล์กแลนด์", Symbol: "FKP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "ฟรังก์ฝรั่งเศส", Symbol: "FRF"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "ดรัชมากรีก", Symbol: "GRD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "เควตซัลกัวเตมาลา", Symbol: "GTQ"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "บัตรปันส่วนมอลโดวา", Symbol: "MDC"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "รูปีมอริเชียส", Symbol: "MUR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "รูเบิลโซเวียต", Symbol: "SUR"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "เรียลบราซิล", Symbol: "R$"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ยูโร", Symbol: "€"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "ทอลาร์สโลวีเนีย", Symbol: "SIT"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ควาชาแซมเบีย", Symbol: "ZMW"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "ฟรังก์กินี", Symbol: "GNF"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "ฮวานเกาหลีใต้ (1953–1962)", Symbol: "KRH"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "ดีนาร์ซูดานเก่า", Symbol: "SDD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ฟรังก์เซฟาธนาคารกลางรัฐแอฟริกาตะวันตก", Symbol: "CFA"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "อัฟกานิอัฟกานิสถาน", Symbol: "AFN"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "ครูนเอสโตเนีย", Symbol: "EEK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "ลิวโรมาเนีย", Symbol: "RON"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "ดีนาร์บาห์เรน", Symbol: "BHD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "ครูเซโรบราซิล (1990–1993)", Symbol: "BRE"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "เลมปิราฮอนดูรัส", Symbol: "HNL"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "อินตีเปรู", Symbol: "PEI"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "เอสคูโดโปรตุเกส", Symbol: "PTE"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ฟรังก์รวันดา", Symbol: "RWF"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "ชิลลิงแทนซาเนีย", Symbol: "TZS"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "หน่วยโมเนทารียุโรป", Symbol: "XBB"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "ดอลลาร์ซิมบับเว", Symbol: "ZWD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "กวานซาแองโกลา (1977–1990)", Symbol: "AOK"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "ทาโลนัสลิทัวเนีย", Symbol: "LTT"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "ซลอตีโปแลนด์", Symbol: "PLN"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "เอสคูโดเคปเวิร์ด", Symbol: "CVE"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "โครูนาสโลวัก", Symbol: "SKK"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ชิลลิงยูกันดา (1966–1987)", Symbol: "UGS"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "เปโซอุรุกวัย (1975–1993)", Symbol: "UYP"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "นิวรูเบิลเบลารุส (1994–1999)", Symbol: "BYB"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "ปอนด์ไซปรัส", Symbol: "CYP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "ดีนาร์คูเวต", Symbol: "KWD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "นูโวซอลเปรู", Symbol: "PEN"}, "USD": ut.Currency{Currency: "USD", DisplayName: "ดอลลาร์สหรัฐ", Symbol: "US$"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "โบลิวาร์เวเนซุเอลา", Symbol: "VEF"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "ฟรังก์ยูไอซีฝรั่งเศส", Symbol: "XFU"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "ครูเซโรบราซิล", Symbol: "BRR"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "มาร์กเยอรมัน", Symbol: "DEM"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "เปโซโดมินิกัน", Symbol: "DOP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "อาเรียรีมาลากาซี", Symbol: "MGA"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ฮรีฟเนียยูเครน", Symbol: "UAH"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "เปโซเอนยูนิแดดเซสอินเด็กซาแดสอุรุกวัย", Symbol: "UYI"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "ดอลลาร์เบอร์มิวดา", Symbol: "BMD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "คูปอนลาริตจอร์เจีย", Symbol: "GEK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "ลารีจอร์เจีย", Symbol: "GEL"}, "INR": ut.Currency{Currency: "INR", DisplayName: "รูปีอินเดีย", Symbol: "₹"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "วอนเกาหลีใต้ (1945–1953)", Symbol: "KRO"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "ดอลลาร์นามิเบีย", Symbol: "NAD"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "นิวแซร์คองโก", Symbol: "ZRN"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "กวานซารีจัสทาโดแองโกลา (1995–1999)", Symbol: "AOR"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "ครูเซโรโนโวบราซิล (1967–1986)", Symbol: "BRB"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "วาเลอร์คอนสแตนต์เอกวาดอร์", Symbol: "ECV"}, "YER": ut.Currency{Currency: "YER", DisplayName: "เรียลเยเมน", Symbol: "YER"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "เอ็งกุลตรัมภูฏาน", Symbol: "BTN"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ลิวโรมาเนียเก่า", Symbol: "ROL"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "คอนเวอร์ทิเบิลดีนาร์ยูโกสลาเวีย", Symbol: "YUN"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "แรนด์แอฟริกาใต้", Symbol: "ZAR"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "ยูโรดับเบิลยูไออาร์", Symbol: "CHE"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "นิวเชเกลอิสราเอล", Symbol: "₪"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ดีนาร์อิรัก", Symbol: "IQD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "จ๊าตพม่า", Symbol: "MMK"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "กองทุนไรเน็ต", Symbol: "XRE"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "ฟรังก์เบลเยียม", Symbol: "BEF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "เซดีกานา", Symbol: "GHS"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "ดอลลาร์หมู่เกาะเคย์แมน", Symbol: "KYD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "ดองเวียดนาม", Symbol: "₫"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "หน่วยบัญชี เอดีบี", Symbol: "XUA"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "เอ็กเวเลอิเควทอเรียลกินี", Symbol: "GQE"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "เงินเปโซเม็กซิโก (1861–1992)", Symbol: "MXP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ชิลลิงยูกันดา", Symbol: "UGX"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "หน่วยคอมโพสิตยุโรป", Symbol: "XBA"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "ฮาร์ดดีนาร์ยูโกสลาเวีย", Symbol: "YUD"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "เลฟบัลแกเรีย", Symbol: "BGN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "ดีนาร์แอลจีเรีย", Symbol: "DZD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "บัลบัวปานามา", Symbol: "PAB"}, "PES": ut.Currency{Currency: "PES", DisplayName: "ซอลเปรู", Symbol: "PES"}, "AON": ut.Currency{Currency: "AON", DisplayName: "นิวกวานซาแองโกลา (1990–2000)", Symbol: "AON"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "ดีนาร์ใหม่บอสเนีย-เฮอร์เซโกวีนา (1994–1997)", Symbol: "BAN"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "มานัตเติร์กเมนิสถาน", Symbol: "TMT"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "ดอลลาร์ซิมบับเว (2009)", Symbol: "ZWL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "เรียลอิหร่าน", Symbol: "IRR"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "ลีรามอลตา", Symbol: "MTL"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "มานัตเติร์กเมนิสถาน (1993–2009)", Symbol: "TMM"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "ดอลลาร์ตรินิแดดและโตเบโก", Symbol: "TTD"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "รูเบิลลัตเวีย", Symbol: "LVR"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "ฟูเมนโตชิลี", Symbol: "CLF"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "เปเซตาสเปน", Symbol: "ESP"}, "USS": ut.Currency{Currency: "USS", DisplayName: "ดอลลาร์สหรัฐ (วันเดียวกัน)", Symbol: "USS"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "มาร์กเยอรมันตะวันออก", Symbol: "DDM"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ดอลลาร์จาเมกา", Symbol: "JMD"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "ครูซาโดบราซิล", Symbol: "BRC"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "ครูซาโดโนโวบราซิล", Symbol: "BRN"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "ฟรังก์มาดากัสการ์", Symbol: "MGF"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "เปโซฟิลิปปินส์", Symbol: "PHP"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "ดอลลาร์บาฮามาส", Symbol: "BSD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "ฟรังก์สวิส", Symbol: "CHF"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "ฟรังก์ลักเซมเบิร์ก", Symbol: "LUF"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "ฟรังก์มาลี", Symbol: "MLF"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "อูกียามอริเตเนีย", Symbol: "MRO"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "รูปีปากีสถาน", Symbol: "PKR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "วาตูวานูอาตู", Symbol: "VUV"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "ฟรังก์ทองฝรั่งเศส", Symbol: "XFO"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "ฮาร์ดเลฟบัลแกเรีย", Symbol: "BGL"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ดาลาซีแกมเบีย", Symbol: "GMD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "ดอลลาร์กายอานา", Symbol: "GYD"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "ยูนิแดด ดี อินเวอร์ชั่น เม็กซิโก", Symbol: "MXV"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "ลีราตุรกีเก่า", Symbol: "TRL"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "คูนาโครเอเชีย", Symbol: "HRK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "ดอบราเซาตูเมและปรินซิปี", Symbol: "STD"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "ออสตรัลอาร์เจนตินา", Symbol: "ARA"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "เซดีกานา (1979–2007)", Symbol: "GHC"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "ลิวมอลโดวา", Symbol: "MDL"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "ไนราไนจีเรีย", Symbol: "NGN"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "เปโซโบลิเวีย", Symbol: "BOP"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "เปโซคิวบา (แปลงสภาพ)", Symbol: "CUC"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "ดอลลาร์นิวซีแลนด์", Symbol: "NZ$"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "รูเบิลรัสเซีย (1991–1998)", Symbol: "RUR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "ปอนด์ซูดานเก่า", Symbol: "SDP"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "รหัสทดสอบสกุลเงิน", Symbol: "XTS"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "รูเบิลเบลารุส", Symbol: "BYR"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "เปโซคิวบา", Symbol: "CUP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "เยนญี่ปุ่น", Symbol: "¥"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "ดีนาร์มาซิโดเนีย (1992–1993)", Symbol: "MKN"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "รูปีเซเชลส์", Symbol: "SCR"}, "COP": ut.Currency{Currency: "COP", DisplayName: "เปโซโคลอมเบีย", Symbol: "COP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "ชิลลิ่งเคนยา", Symbol: "KES"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "กวารานีปารากวัย", Symbol: "PYG"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "ดอลลาร์สิงคโปร์", Symbol: "SGD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "ดอลลาร์ซูรินาเม", Symbol: "SRD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "แรนด์แอฟริกาใต้ (การเงิน)", Symbol: "ZAL"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "กิลเดอร์เนเธอร์แลนด์แอนทิลลิส", Symbol: "ANG"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "ครูเซโรบราซิล (1942–1967)", Symbol: "BRZ"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ปอนด์อียิปต์", Symbol: "EGP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ปอนด์เลบานอน", Symbol: "LBP"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "โลตีเลโซโท", Symbol: "LSL"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "คาร์โบวาเนตซ์ยูเครน", Symbol: "UAK"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "เปโซอาร์เจนตินา", Symbol: "ARS"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "เปเซตาสเปน (บัญชีเอ)", Symbol: "ESA"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "ดีนาร์โครเอเชีย", Symbol: "HRD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "เปโซเม็กซิโก", Symbol: "MX$"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "โซโมนิทาจิกิสถาน", Symbol: "TJS"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "ฟรังก์เบลเยียม (การเงิน)", Symbol: "BEL"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "โบลิเวียโนโบลิเวีย (1863–1963)", Symbol: "BOL"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "ดีนาร์จอร์แดน", Symbol: "JOD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "เรียลกัมพูชา", Symbol: "KHR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "รูปีศรีลังกา", Symbol: "LKR"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "รูปีเนปาล", Symbol: "NPR"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "กีนาปาปัวนิวกินี", Symbol: "PGK"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "ดอลลาร์ซิมบับเว (2008)", Symbol: "ZWR"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "ฟรังก์เบลเยียม (เปลี่ยนแปลงได้)", Symbol: "BEC"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "โซเชียลลิสต์เลฟบัลแกเรีย", Symbol: "BGM"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "กูร์ดเฮติ", Symbol: "HTG"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "ควาชามาลาวี", Symbol: "MWK"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ริงกิตมาเลเซีย", Symbol: "MYR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "ดอลลาร์หมู่เกาะโซโลมอน", Symbol: "SBD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "ดีนาร์ตูนิเซีย", Symbol: "TND"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "ดอลลาร์โรดีเซีย", Symbol: "RHD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ริยัลซาอุดีอาระเบีย", Symbol: "SAR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "วอนเกาหลีใต้", Symbol: "₩"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "กีบลาว", Symbol: "LAK"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "ลัตส์ลัตเวีย", Symbol: "LVL"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "กอร์โดบานิการากัว", Symbol: "NIO"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "ปอนด์ซูดาน", Symbol: "SDG"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "เปโซอุรุกวัย", Symbol: "UYU"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ดอลลาร์แคริบเบียนตะวันออก", Symbol: "EC$"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ปูลาบอตสวานา", Symbol: "BWP"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "เปโซกินี-บิสเซา", Symbol: "GWP"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "ซอมคีร์กีซสถาน", Symbol: "KGS"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "ฟรังก์โมนาโก", Symbol: "MCF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "โครนนอร์เวย์", Symbol: "NOK"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "เปเซตาอันดอร์รา", Symbol: "ADP"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "ดีนาร์บอสเนีย-เฮอร์เซโกวีนา", Symbol: "BAD"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "รูเปียห์อินโดนีเซีย", Symbol: "IDR"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "ปอนด์มอลตา", Symbol: "MTP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "ลีราตุรกี", Symbol: "TRY"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "เปเซตาสเปน (บัญชีที่เปลี่ยนแปลงได้)", Symbol: "ESB"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "ปอนด์ยิบรอลตาร์", Symbol: "GIP"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "ปอนด์อิสราเอล", Symbol: "ILP"}, "WST": ut.Currency{Currency: "WST", DisplayName: "ทาลาซามัว", Symbol: "WST"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "เงิน", Symbol: "XAG"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "ไฟแนลเชียลฟรังก์ลักเซมเบิร์ก", Symbol: "LUL"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "ปาตากามาเก๊า", Symbol: "MOP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "เรียลกาตาร์", Symbol: "QAR"}, "COU": ut.Currency{Currency: "COU", DisplayName: "วาเลอร์เรียลโคลอมเบีย", Symbol: "COU"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "สกุลเงินที่ไม่รู้จัก", Symbol: "XXX"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "แดรมอาร์เมเนีย", Symbol: "AMD"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "เปโซเลย์อาร์เจนตินา (1970–1983)", Symbol: "ARL"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "ดีนาร์เซอร์เบียเก่า", Symbol: "CSD"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "ฟรังก์โมร็อกโก", Symbol: "MAF"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ลีโอนเซียร์ราลีโอน", Symbol: "SLL"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "กวานซาแองโกลา", Symbol: "AOA"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "ฟลอรินอารูบา", Symbol: "AWG"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "เลฟบัลเกเรีย (1879–1952)", Symbol: "BGO"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "เปโซชิลี", Symbol: "CLP"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "โนวิย์ดีนาร์ยูโกสลาเวีย", Symbol: "YUM"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "โครนเดนมาร์ก", Symbol: "DKK"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ดอลลาร์ฮ่องกง", Symbol: "HK$"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "ทูกริกมองโกเลีย", Symbol: "MNT"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "ปอนด์ซูดานใต้", Symbol: "SSP"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "เอสคูโดติมอร์", Symbol: "TPE"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "ซูเกร", Symbol: "XSU"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "โครนาไอซ์แลนด์", Symbol: "ISK"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "ซลอตีโปแลนด์ (1950–1995)", Symbol: "PLZ"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "ดีนาร์เยเมน", Symbol: "YDD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "ฟรังก์บุรุนดี", Symbol: "BIF"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ซูเกรเอกวาดอร์", Symbol: "ECS"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "เทงเจคาซัคสถาน", Symbol: "KZT"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "ลิลันเจนีสวาซิ", Symbol: "SZL"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "ดองเวียดนาม (1978–1985)", Symbol: "VNN"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "ควาชาแซมเบีย (1968–2012)", Symbol: "ZMK"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "มานัตอาเซอร์ไบจาน", Symbol: "AZN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "เบอรร์เอธิโอเปีย", Symbol: "ETB"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "มาร์กกาฟินแลนด์", Symbol: "FIM"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ฟอรินต์ฮังการี", Symbol: "HUF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "กิลเดอร์เนเธอร์แลนด์", Symbol: "NLG"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "ชิลลิงโซมาเลีย", Symbol: "SOS"}}

package yo

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ti Orílẹ́ède Lesoto", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kaṣa ti Orílẹ́ède Malawi", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ti Orílẹ́ède Ariwa Afirika", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupi ti Orílẹ́ède Maritiusi", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupi ti Orílẹ́ède Sayiselesi", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakifa ti Orílẹ́ède Eriteriani", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ti Orílẹ́ède Gamibia", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yeni ti Orílẹ́ède Japani", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dọla ti Orílẹ́ède Liberia", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dọla ti Orílẹ́ède Amerika", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Biri ti Orílẹ́ède Eutopia", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dọla ti Orílẹ́ède Ástràlìá", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ti Orílẹ́ède Bọ̀tìsúwánà", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metika ti Orílẹ́ède Mosamibiki", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faransi ti Orílẹ́ède Ruwanda", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kawaṣa ti Orílẹ́ède Saabia (1968–2012)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "pọọn ti Orílẹ́ède Egipiti", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupi ti Orílẹ́ède Indina", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Faransi ti Orílẹ́ède BEKA", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dọla ti Orílẹ́ède Kánádà", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Sile ti Orílẹ́ède Somali", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Faransi ti Orílẹ́ède Dibouti", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Diami ti Awon Orílẹ́ède Arabu", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faransi ti Orílẹ́ède Bùùrúndì", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "ṣiili ti Orílẹ́ède Kenya", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Siile ti Orílẹ́ède Uganda", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dọla ti Orílẹ́ède Namibia", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faransi ti Orílẹ́ède Kóngò", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dina ti Orílẹ́ède Tunisia", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Uro", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pọọun ti Orílẹ́ède Sudani", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kawaṣa ti Orílẹ́ède Saabia", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ti Orílẹ́ède Nàìjíríà", Symbol: "₦"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Lioni", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Faransi ti Orílẹ́ède BIKEAO", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Wansa ti Orílẹ́ède Àngólà", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faransi ti Orílẹ́ède Siwisi", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faransi ti Orílẹ́ède ṣomoriani", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dina ti Orílẹ́ède Libiya", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya ti Orílẹ́ède Maritania", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Reminibi ti Orílẹ́ède ṣáínà", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirami ti Orílẹ́ède Moroko", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Sile ti Orílẹ́ède Tansania", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kabofediano ti Orílẹ́ède Esuodo", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pọọn ti Orílẹ́ède Bírítísì", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riya ti Orílẹ́ède Saudi", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ṣidi ti Orílẹ́ède Gana", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faransi ti Orílẹ́ède Gini", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Dina ti Orílẹ́ède Sudani", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobira ti Orílẹ́ède Sao tome Ati Pirisipe", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dọla ti Orílẹ́ède Siibabuwe", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dina ti Orílẹ́ède Báránì", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dina ti Orílẹ́ède Àlùgèríánì", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Faransi ti Orílẹ́ède Malagasi", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pọọun ti Orílẹ́ède ̣Elena", Symbol: ""}}

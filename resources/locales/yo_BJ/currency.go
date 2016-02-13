package yo_BJ

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"DJF": ut.Currency{Currency: "DJF", DisplayName: "Faransi ti Orílɛ́ède Dibouti", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupi ti Orílɛ́ède Sayiselesi", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Biri ti Orílɛ́ède Eutopia", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pɔɔn ti Orílɛ́ède Bírítísì", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Sile ti Orílɛ́ède Tansania", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Wansa ti Orílɛ́ède Àngólà", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faransi ti Orílɛ́ède Bùùrúndì", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faransi ti Orílɛ́ède Siwisi", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faransi ti Orílɛ́ède Gini", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ti Orílɛ́ède Nàìjíríà", Symbol: "₦"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faransi ti Orílɛ́ède Ruwanda", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dina ti Orílɛ́ède Báránì", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faransi ti Orílɛ́ède Kóngò", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kabofediano ti Orílɛ́ède Esuodo", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Uro", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kasha ti Orílɛ́ède Malawi", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kawasha ti Orílɛ́ède Saabia", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dɔla ti Orílɛ́ède Ástràlìá", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ti Orílɛ́ède Bɔ̀tìsúwánà", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirami ti Orílɛ́ède Moroko", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dɔla ti Orílɛ́ède Namibia", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Faransi ti Orílɛ́ède BIKEAO", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Faransi ti Orílɛ́ède Malagasi", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pɔɔun ti Orílɛ́ède Sudani", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Sile ti Orílɛ́ède Somali", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dina ti Orílɛ́ède Tunisia", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ti Orílɛ́ède Ariwa Afirika", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dɔla ti Orílɛ́ède Siibabuwe", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Dina ti Orílɛ́ède Sudani", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kawasha ti Orílɛ́ède Saabia (1968–2012)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dɔla ti Orílɛ́ède Liberia", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupi ti Orílɛ́ède Maritiusi", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riya ti Orílɛ́ède Saudi", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "shidi ti Orílɛ́ède Gana", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Lioni", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dɔla ti Orílɛ́ède Amerika", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Faransi ti Orílɛ́ède BEKA", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dɔla ti Orílɛ́ède Kánádà", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "pɔɔn ti Orílɛ́ède Egipiti", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ti Orílɛ́ède Gamibia", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupi ti Orílɛ́ède Indina", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dina ti Orílɛ́ède Àlùgèríánì", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yeni ti Orílɛ́ède Japani", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faransi ti Orílɛ́ède shomoriani", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dina ti Orílɛ́ède Libiya", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Siile ti Orílɛ́ède Uganda", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakifa ti Orílɛ́ède Eriteriani", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Diami ti Awon Orílɛ́ède Arabu", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "shiili ti Orílɛ́ède Kenya", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ti Orílɛ́ède Lesoto", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metika ti Orílɛ́ède Mosamibiki", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Reminibi ti Orílɛ́ède sháínà", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya ti Orílɛ́ède Maritania", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pɔɔun ti Orílɛ́ède ̣Elena", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobira ti Orílɛ́ède Sao tome Ati Pirisipe", Symbol: ""}}

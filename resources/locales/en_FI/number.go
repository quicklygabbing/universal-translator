package en_FI

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = ut.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0\u00a0%"}
)

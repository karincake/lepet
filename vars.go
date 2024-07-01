package lepet

// The built-in language item list for error messages
var defaultList = LangItem{
	"required":         "value is required",
	"data-notFound":    "data can not be found",
	"integer":          "value must be an integer",
	"integer-positive": "value must be a positive integer",
	"uuid":             "value must be a valid UUID",
	"parse-fail":       "parsing failed",
}

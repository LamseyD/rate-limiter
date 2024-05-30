package server

type IPExtractorOption int

// See documentation of extractors at https://echo.labstack.com/docs/ip-address
// TODO Perhaps discuss with Michael what different options mean?
const (
	IPExtractorOptionNone IPExtractorOption = iota
	IPExtractorNoProxy
	IPExtractorXFF // X-Forwarded-For
	IPExtractorXRI // X-Real-IP
)

type Config struct {
	IPExtractorOption int `envconfig:"IP_EXTRACTOR_OPTION" default:"0"`
}

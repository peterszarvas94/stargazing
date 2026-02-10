package middleware

import (
	"github.com/CAFxX/httpcompression"
	"github.com/CAFxX/httpcompression/contrib/andybalholm/brotli"
	"github.com/labstack/echo/v4"
)

// Compression returns an Echo middleware that compresses responses using Brotli.
func Compression() echo.MiddlewareFunc {
	brEnc, err := brotli.New(brotli.Options{})
	if err != nil {
		panic("failed to create brotli encoder: " + err.Error())
	}

	compress, err := httpcompression.Adapter(
		httpcompression.Compressor(brotli.Encoding, 0, brEnc),
		httpcompression.MinSize(256),
	)
	if err != nil {
		panic("failed to create compression adapter: " + err.Error())
	}

	return echo.WrapMiddleware(compress)
}

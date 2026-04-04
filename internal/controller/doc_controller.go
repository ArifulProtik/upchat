package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

// OpenAPISpec serves the generated OpenAPI JSON spec.
//
//	@Summary	OpenAPI spec
//	@Tags		docs
//	@Produce	json
//	@Success	200	{object}	object
//	@Router		/docs/openapi.json [get]
func (c *Controller) OpenAPISpec(ctx *gin.Context) {
	doc, err := swag.ReadDoc()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to read API spec"},
		)
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, doc)
}

// ScalarDocs serves the Scalar API reference UI.
func (c *Controller) ScalarDocs(ctx *gin.Context) {
	html := `<!doctype html>
<html>
<head>
	<title>UpChat API Docs</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<style>body { margin: 0; }</style>
</head>
<body>
	<script
		id="api-reference"
		data-url="/api/docs/openapi.json"
		src="https://cdn.jsdelivr.net/npm/@scalar/api-reference">
	</script>
</body>
</html>`
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, html)
}

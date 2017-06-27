package posts

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get post.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	post, err := arn.GetPost(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Post not found", err)
	}

	return ctx.HTML(components.Post(post))
}

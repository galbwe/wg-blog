package actions

import (
	"net/http"
	"wg_blog/models"

	"github.com/gobuffalo/buffalo"
)

func BlogHandler(c buffalo.Context) error {
	posts := models.GetAllPublishedPosts(models.DB)
	c.Set("posts", posts)
	return c.Render(http.StatusOK, r.HTML("blog/blog.html"))
}

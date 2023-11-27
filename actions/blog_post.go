package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"

	"wg_blog/models"
)

// BlogPostBlogPost default implementation.
func BlogPostHandler(c buffalo.Context) error {
	post := models.GetPostById(models.DB, c.Param("id"))
	// TODO: handle nil return value by redirecting to not found page
	c.Set("post", post)
	return c.Render(http.StatusOK, r.HTML("blog_post/blog_post.html"))
}

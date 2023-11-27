package grifts

import (
	"time"

	"wg_blog/models"

	"github.com/gobuffalo/grift/grift"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		conn, _ := pop.Connect("development")
		posts := []*models.Post{
			{
				// CreatedAt:   time.Now(),
				// UpdatedAt:   time.Now(),
				Title:       "My Third Post",
				Content:     "This is my third post",
				PublishedAt: nulls.NewTime(time.Now()),
			},
			{
				// CreatedAt:   time.Now(),
				// UpdatedAt:   time.Now(),
				Title:       "My Second Post",
				Content:     "This is my second post",
				PublishedAt: nulls.NewTime(time.Now()),
			},
			{
				// CreatedAt:   time.Now(),
				// UpdatedAt:   time.Now(),
				Title:       "My First Post",
				Content:     "This is my first post",
				PublishedAt: nulls.NewTime(time.Now()),
			},
			{
				// CreatedAt:   time.Now(),
				// UpdatedAt:   time.Now(),
				Title:   "Work in Progress",
				Content: "This is a work in progress",
			},
		}
		for _, p := range posts {
			conn.Save(p)
		}

		return nil
	})

})

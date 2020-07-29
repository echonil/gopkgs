package middleware

import (
	"context"
	"net/http"

	"clevergo.tech/clevergo"
	"github.com/echonil/gopkgs/internal/models"
	"github.com/echonil/gopkgs/internal/web"
	"github.com/jmoiron/sqlx"
)

func GoGet(db *sqlx.DB) clevergo.MiddlewareFunc {
	return func(next clevergo.Handle) clevergo.Handle {
		return func(c *clevergo.Context) error {
			if !c.IsGet() || c.QueryParam("go-get") != "1" {
				return next(c)
			}

			var pkg models.Package
			ctx := c.Context()
			if err := models.FindPackageByDomainAndPath(ctx, db, &pkg, c.Host(), c.Request.URL.Path[1:]); err != nil {
				return err
			}

			go func() {
				action := models.NewAction(models.ActionGoGet, pkg.ID)
				if err := action.Save(context.Background(), db); err != nil {
					c.Logger().Errorf("failed to insert an action record: %s", err.Error())
				}
			}()

			return c.Render(http.StatusOK, "package/view.tmpl", clevergo.Map{
				"page": web.NewPage(pkg.Prefix()),
				"pkg":  pkg,
			})
		}
	}
}

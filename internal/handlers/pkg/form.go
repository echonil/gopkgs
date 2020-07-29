package pkg

import (
	"context"
	"errors"

	"github.com/echonil/gopkgs/internal/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jmoiron/sqlx"
)

type Form struct {
	db       *sqlx.DB
	pkg      *models.Package
	userID   string
	DomainID int64  `json:"domain_id" schema:"domain_id"`
	Path     string `json:"path" schema:"path"`
	VCS      string `json:"vcs" schema:"vcs"`
	Root     string `json:"root" schema:"root"`
	Docs     string `json:"docs" schema:"docs"`
}

func newForm(db *sqlx.DB, userID string) *Form {
	return &Form{
		db:     db,
		userID: userID,
	}
}

func newFormPkg(db *sqlx.DB, userID string, pkg *models.Package) *Form {
	f := newForm(db, userID)
	f.pkg = pkg
	f.DomainID = f.pkg.DomainID
	f.Path = f.pkg.Path
	f.Root = f.pkg.Root
	f.VCS = f.pkg.VCS
	f.Docs = f.pkg.Docs
	return f
}

// Validate validates package.
func (f *Form) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.DomainID, validation.Required, validation.WithContext(f.validateDomain)),
		validation.Field(&f.Path, validation.Required, validation.WithContext(f.validatePath)),
		validation.Field(&f.VCS, validation.Required, validation.In(models.VCSGit, models.VCSSvn, models.VCSBzr, models.VCSHg, models.VCSFossil)),
		validation.Field(&f.Root, validation.Required),
		validation.Field(&f.Docs, is.URL),
	)
}

func (f *Form) validateDomain(ctx context.Context, value interface{}) error {
	query := "SELECT * FROM domains WHERE id = ? AND user_id = ?"
	var domain models.Domain
	if err := f.db.GetContext(ctx, &domain, query, value.(int64), f.userID); err != nil {
		return err
	}
	if !domain.Verified {
		return errors.New("domain has not been verified")
	}
	return nil
}

func (f *Form) validatePath(ctx context.Context, value interface{}) error {
	query := "SELECT COUNT(1) FROM packages WHERE domain_id = ? AND path = ?"
	args := []interface{}{f.DomainID, value.(string)}
	if f.pkg != nil {
		query += " AND id != ?"
		args = append(args, f.pkg.ID)
	}
	var count int64
	if err := f.db.GetContext(ctx, &count, query, args...); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("package already exists")
	}
	return nil
}

func (f *Form) Create(ctx context.Context) (pkg *models.Package, err error) {
	pkg = models.NewPackage(f.DomainID, f.Path, f.VCS, f.Root)
	pkg.Docs = f.Docs
	err = pkg.Insert(ctx, f.db)
	return
}

func (f *Form) Update(ctx context.Context) (err error) {
	f.pkg.DomainID = f.DomainID
	f.pkg.Path = f.Path
	f.pkg.Root = f.Root
	f.pkg.VCS = f.VCS
	f.pkg.Docs = f.Docs
	return f.pkg.Update(ctx, f.db)
}

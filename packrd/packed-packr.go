// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "119ad2426fc2e48763bf6f42f22872f4"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"32c5425df5a87715e327f43ea547be19": "1f8b08000000000000ff6c90bf6eb34010c47b9e62bfabbe14702923715c63a788943f28728a941b3c70c8c09163b16c59bc7b141312474ab55aedcc6f76d7fc5b3fad36aff92d39691b1b99a580b736322d84c989f431dec77a9fa995ef049dc49b630f45c5dc654a7010fd694da9701c0648364a19df28bd403a6e91a9cac775dbfb2017ded389927e572577e7c10384ff5fd134a9bfe203ca80c15db8af531a43932d90b52f8697e7fb6f82d4d2c0e65cecb8022daa3ca0ac0f344d46cf82c8e8af8bdffcf668a3472faeee2a124f03400e012919261750fe6cfc3bacf57b1037beab8c669b4446cf2c737e8cfd080000ffffde412fef6a010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("views", "../views")
		b.SetResolver("package.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "32c5425df5a87715e327f43ea547be19"})
	}()
	return nil
}()

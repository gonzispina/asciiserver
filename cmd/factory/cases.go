package factory

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
)

func NewCasesFactory(
	repositories *Repositories,
) *Cases {
	if repositories == nil {
		panic("repostories must be initialized")
	}
	return &Cases{
		repositories: repositories,
	}
}

// Cases factory
type Cases struct {
	repositories *Repositories

	asciiDrawer asciidrawer.Drawer
}

/*
 *
 * Canvas
 *
 */

// ASCIIDrawer use case
func (f *Cases) ASCIIDrawer() asciidrawer.Drawer {
	if f.asciiDrawer == nil {
		f.asciiDrawer = asciidrawer.NewDrawer(f.repositories.CanvasStorage())
	}
	return f.asciiDrawer
}

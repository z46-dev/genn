// Package fonts embeds the archived client's required Ubuntu typeface.
package fonts

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/z46-dev/gctx2d"
)

//go:embed Ubuntu-Bold.ttf
var ubuntuBold []byte

// LoadUbuntu loads the embedded Ubuntu Bold face through Gctx2D.
func LoadUbuntu() (font *gctx2d.FontHandle, err error) {
	var temporary *os.File
	if temporary, err = os.CreateTemp("", "genn-ubuntu-*.ttf"); err != nil {
		err = fmt.Errorf("create embedded Ubuntu font file: %w", err)
		return
	}

	var name string = temporary.Name()
	defer func() { _ = os.Remove(name) }()

	if _, err = temporary.Write(ubuntuBold); err != nil {
		_ = temporary.Close()
		err = fmt.Errorf("write embedded Ubuntu font: %w", err)
		return
	}

	if err = temporary.Close(); err != nil {
		err = fmt.Errorf("close embedded Ubuntu font: %w", err)
		return
	}

	if font, err = gctx2d.LoadFont(name); err != nil {
		err = fmt.Errorf("load embedded Ubuntu font: %w", err)
	}

	return
}

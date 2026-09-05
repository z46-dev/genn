package render

import (
	"strconv"
	"strings"

	"github.com/z46-dev/gctx2d"
)

func MixColors(primary, secondary gctx2d.Color, factor float32) (mixed gctx2d.Color) {
	mixed = gctx2d.Color{
		R: primary.R*(1-factor) + secondary.R*factor,
		G: primary.G*(1-factor) + secondary.G*factor,
		B: primary.B*(1-factor) + secondary.B*factor,
		A: primary.A*(1-factor) + secondary.A*factor,
	}

	return
}

func Text(ctx *gctx2d.Context, text string, x, y float32, fontSize float64, fontFamily *gctx2d.FontHandle, color gctx2d.Color) {
	ctx.SetFillStyle(color)
	ctx.SetFont(fontSize, nil)
	ctx.FillText(text, x, y)
}

func HexToColor(hex string) (color gctx2d.Color) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 3 && len(hex) != 6 && len(hex) != 8 {
		return
	}

	parse := func(value string) (float32, bool) {
		parsed, err := strconv.ParseUint(value, 16, 8)
		if err != nil {
			return 0, false
		}
		return float32(parsed) / 255, true
	}

	if len(hex) == 3 {
		values := [3]float32{}
		for i := range values {
			value, ok := parse(hex[i:i+1] + hex[i:i+1])
			if !ok {
				return gctx2d.Color{}
			}
			values[i] = value
		}
		color.R, color.G, color.B, color.A = values[0], values[1], values[2], 1
		return
	}

	values := [4]float32{0, 0, 0, 1}
	for i := 0; i < len(hex)/2; i++ {
		value, ok := parse(hex[i*2 : i*2+2])
		if !ok {
			return gctx2d.Color{}
		}
		values[i] = value
	}
	color.R, color.G, color.B, color.A = values[0], values[1], values[2], values[3]
	return
}

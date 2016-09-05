package utils

import (
	"math"

	"github.com/lsegal/gorb/test/crosspkg/data"
)

func ToHSV(rgb data.RGB) data.HSV {
	rf := float64(rgb.R) / 255.0
	gf := float64(rgb.G) / 255.0
	bf := float64(rgb.B) / 255.0
	minRGB := math.Min(rf, math.Min(gf, bf))
	maxRGB := math.Max(rf, math.Max(gf, bf))

	if minRGB == maxRGB {
		return data.HSV{V: minRGB}
	}

	var d float64
	var h float64
	if rf == minRGB {
		d = gf - bf
		h = 3.0
	} else if bf == minRGB {
		d = rf - gf
		h = 1.0
	} else {
		d = bf - rf
		h = 5.0
	}

	return data.HSV{
		H: 60 * (h - d/(maxRGB-minRGB)),
		S: (maxRGB - minRGB) / maxRGB,
		V: maxRGB,
	}
}

package isogrids

import (
	"crypto/md5"
	"fmt"

	"io"
	"log"
	"net/http"

	tgColors "github.com/taironas/tinygraphs/colors"
	"github.com/taironas/tinygraphs/draw"
	"github.com/taironas/tinygraphs/extract"

	"github.com/taironas/tinygraphs/misc"
	"github.com/taironas/tinygraphs/write"
)

// Isogrids is the handler for /isogrids/[a-zA-Z0-9]+/?.
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Isogrids(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 2); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.Isogrids(w, key, bg, fg, size)
	}
}

// Isogrids is the handler for /isogrids/[a-zA-Z0-9]+/?.
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Color(w http.ResponseWriter, r *http.Request) {
	if colorId, err := misc.PermalinkID(r, 2); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		if id, err1 := misc.PermalinkString(r, 3); err1 == nil {
			h := md5.New()
			io.WriteString(h, id)
			key := fmt.Sprintf("%x", h.Sum(nil)[:])

			colorMap := tgColors.MapOfColorPatterns()
			bg, err1 := extract.Background(r)
			if err1 != nil {
				bg = colorMap[int(colorId)][0]
			}
			fg, err2 := extract.Foreground(r)
			if err2 != nil {
				fg = colorMap[int(colorId)][1]
			}
			size := extract.Size(r)
			write.ImageSVG(w)
			draw.Isogrids(w, key, bg, fg, size)
		} else {
			log.Printf("error when extracting permalink string: %v", err)
		}
	}
}

func Skeleton(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsSkeleton(w, key, bg, fg, size)
	}
}

func Diagonals(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.Diagonals(w, key, bg, fg, size)
	}
}

func HalfDiagonals(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.HalfDiagonals(w, key, bg, fg, size)
	}
}

func GridBW(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsBW(w, key, bg, fg, size)
	}
}

func Grid2Colors(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.Isogrids2Colors(w, key, bg, fg, size)
	}
}

func Random(w http.ResponseWriter, r *http.Request) {

	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsRandom(w, key, bg, fg, size)
	}
}

func RandomColor(w http.ResponseWriter, r *http.Request) {
	if id, err := misc.PermalinkID(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[int(id)][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[int(id)][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsRandom(w, "", bg, fg, size)
	}
}

func RandomMirror(w http.ResponseWriter, r *http.Request) {
	if id, err := misc.PermalinkString(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		h := md5.New()
		io.WriteString(h, id)
		key := fmt.Sprintf("%x", h.Sum(nil)[:])

		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[0][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[0][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsRandomMirror(w, key, bg, fg, size)
	}
}

func RandomMirrorColor(w http.ResponseWriter, r *http.Request) {
	if id, err := misc.PermalinkID(r, 3); err != nil {
		log.Printf("error when extracting permalink id: %v", err)
	} else {
		colorMap := tgColors.MapOfColorPatterns()
		bg, err1 := extract.Background(r)
		if err1 != nil {
			bg = colorMap[int(id)][0]
		}
		fg, err2 := extract.Foreground(r)
		if err2 != nil {
			fg = colorMap[int(id)][1]
		}
		size := extract.Size(r)
		write.ImageSVG(w)
		draw.IsogridsRandomMirror(w, "", bg, fg, size)
	}
}

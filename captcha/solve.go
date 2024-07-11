package captcha

import (
	"bytes"
	"github.com/otiai10/gosseract/v2"
	"image/jpeg"
	"io"
)

// width of each character in pixels
var W = 30

func findMostFrequent(chars map[byte]int) byte {
	var m int
	var r byte

	for c, f := range chars {
		if f > m {
			m = f
			r = c
		}
	}
	return r
}

type Solver struct {
  difficulty int
	gosseract.Client
}

func NewSolver(difficulty int) *Solver {
	client := gosseract.NewClient()
	client.SetWhitelist("QWERTYUIOPASDFGHJKLZXCVBNM")
	client.SetPageSegMode(gosseract.PSM_SINGLE_CHAR)
	return &Solver{
		Client: *client,
    difficulty: difficulty, 
	}
}

/*
higher the n, takes longer to solve and is more accurate
*/
func (solver *Solver) Solve(getCaptcha func() io.Reader) string {
	img, err := jpeg.Decode(getCaptcha())
	if err != nil {
		panic(err)
	}

	nParts := img.Bounds().Dx() / W

	parts := make([]map[byte]int, nParts)

	for i := range parts {
		parts[i] = make(map[byte]int)
	}

	for k := 0; k < solver.difficulty; k++ {
		img, err := jpeg.Decode(getCaptcha())

		if err != nil {
			panic(err)
		}

		for i, p := range parts {
			bytesBuf := new(bytes.Buffer)

			Part(img, i, W, bytesBuf)

			imgBytes, _ := io.ReadAll(bytesBuf)
			solver.SetImageFromBytes(imgBytes)

			text, _ := solver.Text()

			if len(text) == 1 {
				p[text[0]]++
			}
		}
	}

	r := new(bytes.Buffer)

	for _, p := range parts {
		r.WriteByte(findMostFrequent(p))
	}

	return r.String()
}

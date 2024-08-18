package avatar

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"log"
	"sort"
	"strconv"
	"unicode"
)

func GenerateAvatar(username string, xMax, yMax, blockSize int, steps int) []byte {
	gradient := generateGradient(username, steps)
	img := image.NewRGBA(image.Rect(0, 0, xMax, yMax))
	stepsCount := len(gradient)

	for x := 0; x < xMax/2; x += blockSize {
		for y := 0; y < yMax; y += blockSize {

			index := (x + y) / blockSize % stepsCount
			col := gradient[index]

			if render(username, x, y) {
				for i := 0; i < blockSize; i++ {
					for j := 0; j < blockSize; j++ {
						img.Set(x+i, y+j, col)
						img.Set(xMax-x-i-1, y+j, col)
					}
				}
			}
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}

	return buf.Bytes()
}

func render(username string, x, y int) bool {
	hash := generateHash(username + strconv.Itoa(x) + strconv.Itoa(y))
	val := hexToInt(hash[:8]) // Используем первые 8 символов хеша
	return val%4 > 1
}

func generateHash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func hexToInt(hexStr string) int {
	val, _ := strconv.ParseInt(hexStr, 16, 0)
	return int(val)
}

func generateGradient(username string, steps int) []color.RGBA {
	runes := []rune(username)
	for i, r := range runes {
		runes[i] = unicode.ToUpper(r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	hash := generateHash(username)
	baseColor := color.RGBA{
		uint8(hexToInt(hash[0:2])),
		uint8(hexToInt(hash[2:4])),
		uint8(hexToInt(hash[4:6])),
		255,
	}

	gradient := make([]color.RGBA, steps)
	for i := 0; i < steps; i++ {
		gradient[i] = color.RGBA{
			uint8((int(baseColor.R) + i*20) % 255),
			uint8((int(baseColor.G) + i*20) % 255),
			uint8((int(baseColor.B) + i*20) % 255),
			255,
		}
	}
	return gradient
}

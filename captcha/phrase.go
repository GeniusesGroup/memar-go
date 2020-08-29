/* For license and copyright information please see LEGAL file in repository */

package captcha

import (
	"bytes"
	"container/list"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math/rand"
	"strconv"
	"time"
	"unsafe"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobolditalic"

	etime "../earth-time"
	"../log"
	"../uuid"
)

/*
Usage:
var phraseCaptchas = captcha.NewDefaultPhraseCaptchas()
var pc *captcha.PhraseCaptcha = phraseCaptchas.NewImage(req.Language, req.ImageFormat)

*/

// PhraseCaptchas store
type PhraseCaptchas struct {
	Len        uint8       // Number of captcha solution. can't be more than 16!
	Difficulty uint8       // 0:very-easy, 1:easy, 2:medium, 3:hard, 4:very-hard, 5:extreme-hard
	Type       uint8       // 0:Number(625896), 1:Word(A19Cat), 2:Math(+ - * /),
	Duration   int64       // The number of seconds indicate expiration time of captchas
	ImageSize  image.Point // Standard width & height of a captcha image.
	Pool       map[[16]byte]*PhraseCaptcha
	idByTime   *list.List
}

// PhraseCaptcha store
type PhraseCaptcha struct {
	ID       [16]byte
	Answer   string
	ExpireIn int64
	State    state
	Image    []byte // In requested lang & format
	Audio    []byte // In requested lang & format
}

// NewDefaultPhraseCaptchas use to make new captchas with defaults values!
func NewDefaultPhraseCaptchas() (pcs *PhraseCaptchas) {
	pcs = &PhraseCaptchas{
		Len:        6,
		Difficulty: 2,
		Type:       0,
		Duration:   2 * 60, // 2 Minute
		ImageSize:  image.Point{128, 64},
		Pool:       make(map[[16]byte]*PhraseCaptcha, 1024),
		idByTime:   list.New(),
	}
	// cleaner for expired captchas!
	go pcs.expirationProcessing()
	return
}

// NewImage make, store and return new captcha!
func (pcs *PhraseCaptchas) NewImage(lang Language, imageformat ImageFormat) (pc *PhraseCaptcha) {
	pc = &PhraseCaptcha{
		ID:       uuid.NewV4(),
		ExpireIn: etime.Now() + pcs.Duration,
		State:    StateCreated,
	}
	switch pcs.Type {
	case 0:
		pc.Answer = pcs.randomDigits()
	case 1:
		pc.Answer = pcs.randomWord(lang)
	case 2:
		pc.Answer = pcs.randomMath()
	}
	pc.Image = pcs.createImage(pc.Answer, imageformat)

	pcs.Pool[pc.ID] = pc
	return
}

// GetAudio return exiting captcha with audio generated if exits otherwise returns nil!
func (pcs *PhraseCaptchas) GetAudio(captchaID [16]byte, lang Language, audioFormat AudioFormat) (pc *PhraseCaptcha) {
	pc = pcs.Pool[captchaID]
	if pc == nil {
		return nil
	}
	pc.Audio = pcs.createAudio(pc.Answer, audioFormat)
	return
}

// Get return exiting captcha if exits otherwise returns nil!
func (pcs *PhraseCaptchas) Get(captchaID [16]byte) (pc *PhraseCaptcha) {
	pc = pcs.Pool[captchaID]
	if pc != nil && pc.ExpireIn < etime.Now() {
		delete(pcs.Pool, captchaID)
		return nil
	}
	return
}

// Solve check answer and return captcha state!
func (pcs *PhraseCaptchas) Solve(captchaID [16]byte, answer string) error {
	var pc *PhraseCaptcha
	pc = pcs.Pool[captchaID]
	if pc == nil {
		return ErrCaptchaNotFound
	}
	if pc.ExpireIn < etime.Now() {
		delete(pcs.Pool, captchaID)
		return ErrCaptchaExpired
	}
	if pc.Answer != answer {
		pc.State = StateLastAnswerNotValid
		return ErrCaptchaAnswerNotValid
	}
	// Give more time to user to complete any proccess need captcha!
	pc.ExpireIn += pcs.Duration
	pc.State = StateSolved
	return nil
}

// Check return true if captcha exits and solved otherwise returns false!
func (pcs *PhraseCaptchas) Check(captchaID [16]byte) error {
	var pc *PhraseCaptcha
	pc = pcs.Pool[captchaID]
	if pc == nil {
		return ErrCaptchaNotFound
	}
	if pc.ExpireIn < etime.Now() {
		delete(pcs.Pool, captchaID)
		return ErrCaptchaExpired
	}
	if pc.State != StateSolved {
		return ErrCaptchaNotSolved
	}
	return nil
}

func (pcs *PhraseCaptchas) randomDigits() string {
	var low, hi int64
	switch pcs.Len {
	case 6:
		low, hi = 100000, 999999
	case 7:
		low, hi = 1000000, 9999999
	case 8:
		low, hi = 10000000, 99999999
	case 9:
		low, hi = 100000000, 999999999
	case 10:
		low, hi = 1000000000, 9999999999
	default:
		low, hi = 1000000000, 9999999999
	}
	var rand = low + rand.Int63n(hi-low)
	return strconv.FormatInt(rand, 10)
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)
const englishLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789-"

func (pcs *PhraseCaptchas) randomWord(lang Language) string {
	var b = make([]byte, pcs.Len)
	var i uint8
	for i = 0; i < pcs.Len; i++ {
		var random = src.Int63()
		if idx := int(random & letterIdxMask); idx < len(englishLetters) {
			b[i] = englishLetters[idx]
		}
		random >>= letterIdxBits
	}
	return *(*string)(unsafe.Pointer(&b))
}

func (pcs *PhraseCaptchas) randomMath() string {
	// TODO:::
	return ""
}

var goBoldItalic *truetype.Font

func init() {
	var err error
	goBoldItalic, err = freetype.ParseFont(gobolditalic.TTF)
	if err != nil {
		// Almost never occur!
		log.Fatal(err)
	}
}

func (pcs *PhraseCaptchas) createImage(answer string, imageFormat ImageFormat) []byte {
	var img = image.NewRGBA(image.Rect(0, 0, pcs.ImageSize.X, pcs.ImageSize.Y))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	// TODO::: Difficulty??!!
	var c = freetype.NewContext()
	var point = freetype.Pt(10, 10+int(c.PointToFixed(24)>>6))
	c.SetDst(img)
	c.SetSrc(image.Black) //(image.NewUniform(color.RGBA{200, 100, 0, 255}))
	c.SetFont(goBoldItalic)
	c.SetFontSize(24)
	c.SetDPI(72)
	c.SetClip(img.Bounds())
	c.DrawString(answer, point)

	var buf bytes.Buffer
	switch imageFormat {
	case ImageFormatPNG:
		png.Encode(&buf, img)
	case ImageFormatJPEG:
		jpeg.Encode(&buf, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
	return buf.Bytes()
}

func (pcs *PhraseCaptchas) createAudio(answer string, audioFormat AudioFormat) []byte {
	return []byte{}
}

func (pcs *PhraseCaptchas) expirationProcessing() {
	var timer = time.NewTimer(time.Duration(pcs.Duration) * time.Second)
	for {
		select {
		// case shutdownFeedback := <-pcs.shutdownSignal:
		// 	timer.Stop()
		// 	shutdownFeedback <- struct{}{}
		// 	return
		case <-timer.C:
			timer.Reset(time.Duration(pcs.Duration) * time.Second)

			if len(pcs.Pool) == 0 {
				continue
			}

			// Usually this proccess is less than one second, so get time once for compare!
			var timeNow = etime.Now()
			for _, captcha := range pcs.Pool {
				if captcha.ExpireIn < timeNow {
					delete(pcs.Pool, captcha.ID)
				}
			}
		}
	}
}

// https://github.com/search?l=Go&q=captcha&type=Repositories
// https://github.com/dchest/captcha
// https://github.com/afocus/captcha
// https://github.com/steambap/captcha
// https://github.com/lifei6671/gocaptcha
// https://www.hcaptcha.com/

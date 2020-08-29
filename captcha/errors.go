/* For license and copyright information please see LEGAL file in repository */

package captcha

import "../errors"

// package errors
var (
	ErrCaptchaNotFound       = errors.New("CaptchaNotFound", "Given CaptchaID point to the captcha that not exist")
	ErrCaptchaExpired        = errors.New("CaptchaExpired", "Given CaptchaID point to the captcha that expired")
	ErrCaptchaAnswerNotValid = errors.New("CaptchaAnswerNotValid", "Given answer for captcha not valid")
	ErrCaptchaNotSolved      = errors.New("CaptchaNotSolved", "Given CaptchaID point to the captcha not solved yet, Solve it before use it to prove human being")
)

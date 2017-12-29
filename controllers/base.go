package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type BaseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (this *BaseController) Prepare() {

	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}

}

// AppController handles the welcome screen that allows user to pick a technology and username.
func (this *BaseController) setLangVer() bool {
	// Reset language option.
	lang := "" // This field is from i18n.Locale.
	isNeedRedir := false
	hasCookie := false

	//1.Get language information from input.
	lang = this.Input().Get("lang")
	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		fmt.Println("从Accept-Language中获取")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 2. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		fmt.Println("从默认中获取")
	}

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", lang, 1<<31-1, "/")
	}

	this.Lang = lang
	// Set template level language option.
	this.Data["Lang"] = this.Lang
	return isNeedRedir
}

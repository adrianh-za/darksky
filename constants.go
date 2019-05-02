// MIT License
// 
// Copyright (c) 2019 Adrian Houghton
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.


package darksky

// The base URL for the DarkSKy API
const BaseURL = "https://api.darksky.net/forecast"

// Language - The language to request from the API
type ForecastLanguage string

const (
	LangArabic             		ForecastLanguage = "ar"
	LangAzerbaijani        		ForecastLanguage = "az"
	LangBelarusian         		ForecastLanguage = "be"
	LangBulgarian				ForecastLanguage = "bg"
	LangBengali					ForecastLanguage = "bn"
	LangBosnian            		ForecastLanguage = "bs"
	LangCatalan            		ForecastLanguage = "ca"
	LangCzech              		ForecastLanguage = "cs"
	LangDanish					ForecastLanguage = "da"
	LangGerman             		ForecastLanguage = "de"
	LangGreek              		ForecastLanguage = "el"
	LangEnglish            		ForecastLanguage = "en"		//Default language
	LangEsperanto            	ForecastLanguage = "eo"
	LangSpanish            		ForecastLanguage = "es"
	LangEstonian           		ForecastLanguage = "et"
	LangFinnish					ForecastLanguage = "fi"
	LangFrench             		ForecastLanguage = "fr"
	LangHebrew             		ForecastLanguage = "he"
	LangHindi             		ForecastLanguage = "hi"
	LangCroatian           		ForecastLanguage = "hr"
	LangHungarian          		ForecastLanguage = "hu"
	LangIndonesian         		ForecastLanguage = "id"
	LangIcelandic          		ForecastLanguage = "is"
	LangItalian            		ForecastLanguage = "it"
	LangJapanese				ForecastLanguage = "ja"
	LangGeorgian				ForecastLanguage = "ka"
	LangKannada            		ForecastLanguage = "kn"
	LangKorean            		ForecastLanguage = "ko"
	LangCornish            		ForecastLanguage = "kw"
	LangLatvian					ForecastLanguage = "lv"
	LangMalayam					ForecastLanguage = "ml"
	LangMarathi					ForecastLanguage = "mr"
	LangNorwegianBokmål			ForecastLanguage = "nb"
	LangDutch              		ForecastLanguage = "nl"
	LangNorwegianBokmålAlt  	ForecastLanguage = "no"
	LangPunjabi             	ForecastLanguage = "pa"
	LangPolish             		ForecastLanguage = "pl"
	LangPortuguese         		ForecastLanguage = "pt"
	LangRomanian            	ForecastLanguage = "ro"
	LangRussian            		ForecastLanguage = "ru"
	LangSlovak             		ForecastLanguage = "sk"
	LangSlovenian          		ForecastLanguage = "sl"
	LangSerbian            		ForecastLanguage = "sr"
	LangSwedish            		ForecastLanguage = "sv"
	LangTamil            		ForecastLanguage = "ta"
	LangTelugu            		ForecastLanguage = "te"
	LangTetum              		ForecastLanguage = "tet"
	LangTurkish            		ForecastLanguage = "tr"
	LangUkrainian          		ForecastLanguage = "uk"
	LangUrdu          			ForecastLanguage = "ur"
	LangIgpayAtinlay       		ForecastLanguage = "x-pig-latin"
	LangSimplifiedChinese  		ForecastLanguage = "zh"
	LangTraditionalChinese 		ForecastLanguage = "zh-tw"
)

// ForecastExclude are the various info blocks that can be excluded when forecast is returned
type ForecastExclude string

const (
	ExcludeCurrently	ForecastExclude = "currently"
	ExcludeMinutely		ForecastExclude = "minutely"
	ExcludeHourly		ForecastExclude = "hourly"
	ExcludeDaily		ForecastExclude = "daily"
	ExcludeAlerts		ForecastExclude = "alerts"
	ExcludeFlags		ForecastExclude = "flags"
)

// ForecastUnits represent the different kind of measurements unit kinds
type ForecastUnits string

const (
	UnitAuto		ForecastUnits = "auto"
	UnitMetricCA	ForecastUnits = "ca"
	UnitMetricUK	ForecastUnits = "uk2"
	UnitImperial	ForecastUnits = "us"
	UnitMetric		ForecastUnits = "si"	//Default units
)
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


package main

import (
	darksky "github.com/adrianh-za/go-darksky"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Example URL for Johannesburg, South Africa in English, Metric units, no extended hourly forecast, no exclusions
	// https://api.darksky.net/forecast/<YOUR_API_KEY>/-26.1357,28.1392?units=si&lang=en&extend=none&exclude=

	// Build argument
	var argument darksky.ForecastArgument
	argument.Key = "<YOUR_API_KEY>"	//Put your API key here
	argument.Latitude = -26.135746
	argument.Longitude = 28.139208
	argument.Language = darksky.LangEnglish
	argument.Units = darksky.UnitMetric
	argument.ExtendHourly = false
	argument.ExcludeInfo = []darksky.ForecastExclude{darksky.ExcludeAlerts, darksky.ExcludeFlags}
	log.Infoln(argument)

	// Build URL from argument
	var url = darksky.GetForecastURL(argument)
	log.Infoln(url)

	// Do call to API
	var resp, err = darksky.DoForecastCall(url, true)
	if err != nil {
		log.Error(err)
		return
	}

	// Parse the response and build ForecastResponse
	defer resp.Body.Close()
	forecast, err := darksky.ParseForecastJSON(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}

	// Call successful and reponse parsed, but error returned
	if forecast.StatusCode != 0 {
		log.Error(forecast.StatusError)
		return
	}

	// Dump the forecast
	log.Infoln(forecast)
}

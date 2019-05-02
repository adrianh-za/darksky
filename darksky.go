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

import (
	"fmt"
	"strings"
	"net/http"
	"io"
	"encoding/json"
	"io/ioutil"
)

// GetForecastURL returns the API URL to be called built up from the inputted ForecastArgument.
// Defaults will be set if certain argument values are left blank.
func GetForecastURL(arg ForecastArgument) string {

	//If set to hourly, then the total hours will increase from the next 48 hours to the next 168 hours
	var extendHourlyString = "none"
	if (arg.ExtendHourly) {
		extendHourlyString = "hourly"
	}

	//Each listed exclude block will be excluded from return data of the API call (Reference: https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string)
	var excludeString = strings.Trim(strings.Replace(fmt.Sprint(arg.ExcludeInfo), " ", ",", -1), "[]")
	
	//The language that must be returned in the forecast
	var langString = "en"
	if (arg.Language != "") {
		langString = string(arg.Language)
	}

	//String representations of the latitude and longitude
	var latString = fmt.Sprintf("%.4f", arg.Latitude)		//Max precision of 4
	var longString = fmt.Sprintf("%.4f", arg.Longitude)		//Max precision of 4
	
	//Build up the URL to be called
	var forecastURL = BaseURL + "/" + arg.Key + "/" + latString + "," + longString + "?units=" + string(arg.Units) + "&lang=" + langString + "&extend=" + extendHourlyString + "&exclude=" + excludeString
	return forecastURL
}

// DoForecastCall returns a ForecastResponse depending on the API URL passed in.
// Set "compress" to true if the call to the URL must have http compression enabled. 
func DoForecastCall(url string, compress bool) (*http.Response,  error) {
	
	// Setup the transport to handle gzip compression
	var transport = &http.Transport{
		DisableCompression: !compress,
	}
	client := &http.Client{Transport: transport}
	
	// Do the call
	resp, err := client.Get(url)
	if (err != nil) {
		return nil, err
	}

	return resp, nil
}

// ParseForecastJSON returns a ForecastResponse depending parsed from the inputted json
func ParseForecastJSON(reader io.Reader) (*ForecastResponse, error) {
	var forecast ForecastResponse

	err := json.NewDecoder(reader).Decode(&forecast)
	if (err != nil ) {
		return nil, err
	}

	return &forecast, nil
}

// GetForecastText return the text content of the call response
func GetForecastText(reader io.Reader) (string, error) {
	responseData, err := ioutil.ReadAll(reader)
    if err != nil {
        return "", err
    }
 
    return string(responseData), nil
}
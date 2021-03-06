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

// INFO: These JSON structs were mostly generated by using https://mholt.github.io/json-to-go/

// ForecastArgument - the argument object used to generated the call URL
type ForecastArgument struct {
	Key string
	Latitude  float64
	Longitude float64
	ExcludeInfo []ForecastExclude
	ExtendHourly bool
	Language ForecastLanguage
	Units ForecastUnits
	DailyAPICalls int
}
 
// ForecastResponse - the main forecast response containing all the details for the location requested
type ForecastResponse struct {
	Latitude  	float64 		`json:"latitude,omitempty"`
	Longitude 	float64 		`json:"longitude,omitempty"`
	Timezone  	string      	`json:"timezone,omitempty"`
	Currently 	DataPoint  		`json:"currently,omitempty"`
	Minutely  	DataBlock  		`json:"minutely,omitempty"`
	Hourly    	DataBlock  		`json:"hourly,omitempty"`
	Daily     	DataBlock  		`json:"daily,omitempty"`
	Alerts    	[]AlertBlock 	`json:"alerts,omitempty"`
	Flags     	FlagsBlock    	`json:"flags,omitempty"`
	StatusCode	int64			`json:"code,omitempty"`
	StatusError	string      	`json:"error,omitempty"`
}

// DataPoint - A data point object contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon 
// occurring during a period of time: an instant in the case of currently, a minute for minutely, an hour for hourly, and a day for daily.
// See https://darksky.net/dev/docs#data-point-object
type DataPoint struct {
	Time                        int64   `json:"time,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	SunriseTime                 int64   `json:"sunriseTime,omitempty"`
	SunsetTime                  int64   `json:"sunsetTime,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      int64   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         int64   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          int64   `json:"temperatureLowTime,omitempty"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime int64   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  int64   `json:"apparentTemperatureLowTime,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                int64   `json:"windGustTime,omitempty"`
	WindBearing                 int64   `json:"windBearing,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	UvIndex                     int64   `json:"uvIndex,omitempty"`
	UvIndexTime                 int64   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          int64   `json:"temperatureMinTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          int64   `json:"temperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  int64   `json:"apparentTemperatureMinTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  int64   `json:"apparentTemperatureMaxTime,omitempty"`
}

// DataBlock - A data block object represents the various weather phenomena occurring over a period of time. 
// See https://darksky.net/dev/docs#data-block-object
type DataBlock struct {
	Data    []DataPoint `json:"data,omitempty"`
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
}

// Alert -  The alert obkect representing the severe weather warnings issued for the requested location by a governmental authority
// See https://darksky.net/dev/docs#response-alerts
type AlertBlock struct {
	Description string    	`json:"description,omitempty"`
	Expires     int64		`json:"expires,omitempty"`
	Regions     []string  	`json:"regions,omitempty"`
	Severity    string    	`json:"severity,omitempty"`
	Time        int64 		`json:"time,omitempty"`
	Title       string    	`json:"title,omitempty"`
	Uri         string    	`json:"uri,omitempty"`
}

// Flags - The flags object contains various metadata information related to the request.
// See https://darksky.net/dev/docs#response-flags
type FlagsBlock struct {
	DarkSkyUnavailable string      	`json:"darksky-unavailable,omitempty"`
	NearestStation     float64 		`json:"nearest-station,omitempty"`
	Sources            []string    	`json:"sources,omitempty"`
	Units              string      	`json:"units,omitempty"`
}
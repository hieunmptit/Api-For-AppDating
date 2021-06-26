package geo_location

import geo "github.com/martinlindhe/google-geolocate"

func Geolocation() string {
	client := geo.NewGoogleGeo("api-key")
	res, _ := client.Geolocate()
	p := geo.Point{Lat: res.Lat, Lng: res.Lng}
	return p.Address
}

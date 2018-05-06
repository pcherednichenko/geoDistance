[![Torii Build Status](https://circleci.com/gh/pcherednichenko/geoDistance.png?circle-token=b18fdca0e536d760d3faef3906e2cbbba7500ef1 "Geo Distance Build Status")](https://circleci.com/gh/pcherednichenko/geoDistance)

## GeoDistance 

![geo logo](/github/logo.png)

Geo Distance allows you to find the nearest point by coordinate to the original using the [haversine formula](https://en.wikipedia.org/wiki/Haversine_formula) ([great-circle distance](https://en.wikipedia.org/wiki/Great-circle_distance))

### How to run

```
go run main.go -filename=geoData.csv
```
Result:
```
Top 5 closest:
City ID: 442406, distance: 334.210193 m, coordinates: (51.927167, 4.482217), url: https://www.google.com/maps/?q=51.92716710,4.48221710
City ID: 285782, distance: 528.620510 m, coordinates: (51.925356, 4.486310), url: https://www.google.com/maps/?q=51.92535590,4.48630980
City ID: 429151, distance: 648.732565 m, coordinates: (51.925630, 4.488034), url: https://www.google.com/maps/?q=51.92562970,4.48803440
City ID: 512818, distance: 741.377832 m, coordinates: (51.926815, 4.489072), url: https://www.google.com/maps/?q=51.92681520,4.48907200
City ID: 25182, distance: 822.557625 m, coordinates: (51.924912, 4.490593), url: https://www.google.com/maps/?q=51.92491220,4.49059300
Top 5 furthest:
City ID: 382693, distance: 1443325.841185 m, lat: 40.970240, lon: -5.661052, url: https://www.google.com/maps/?q=40.97023990,-5.66105200
City ID: 382582, distance: 1760039.861645 m, lat: 37.176867, lon: -3.608897, url: https://www.google.com/maps/?q=37.17686720,-3.60889700
City ID: 381823, distance: 1760808.374509 m, lat: 37.168004, lon: -3.602987, url: https://www.google.com/maps/?q=37.16800400,-3.60298700
City ID: 382013, distance: 1812135.215856 m, lat: 37.399450, lon: -5.971514, url: https://www.google.com/maps/?q=37.39945000,-5.97151400
City ID: 7818, distance: 8786427.189942 m, lat: 37.866754, lon: -122.259099, url: https://www.google.com/maps/?q=37.86675380,-122.25909900
Running time: 903.196µs
```

If you want to use goroutines for search, run with flag `-useGoroutines`:
```
go run main.go -filename=geoData.csv -useGoroutines
```

### Data source

In my example a csv file is used, but if you want to use a different source, simply use the [data interface](/geo/data/data_interface.go)
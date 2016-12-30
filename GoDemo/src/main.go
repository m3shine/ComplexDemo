package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	/*http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9090", nil)*/
	a :=  [...]int{1,2,1,1}
	var s []int
	fmt.Printf("%T\n%T\n",a,s)
}
func handler(w http.ResponseWriter, r *http.Request) {

	str := "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width=" + strconv.Itoa(width) + " height=" + strconv.Itoa(height) + ">"

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			a := strconv.FormatFloat(ax, 'f', 2, 64)
			b := strconv.FormatFloat(ay, 'f', 2, 64)
			c := strconv.FormatFloat(bx, 'f', 2, 64)
			d := strconv.FormatFloat(by, 'f', 2, 64)
			e := strconv.FormatFloat(cx, 'f', 2, 64)
			f := strconv.FormatFloat(cy, 'f', 2, 64)
			g := strconv.FormatFloat(dx, 'f', 2, 64)
			h := strconv.FormatFloat(dy, 'f', 2, 64)
			str += "<polygon points='" + a + "," + b + ","  + c + "," + d + "," + e + "," + f + "," + g + "," + h + "'/>\n"
		}
	}
	str += "</svg>"
	fmt.Fprintf(w, "%s", str)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

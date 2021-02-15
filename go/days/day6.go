package main

import (
	"fmt"
	"strings"
	"strconv"
	"math/rand"
)

type colorT struct {
	r int
	g int
	b int
}

type coordT struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	} else {
		return x
	}
}

func main() {
	raw_str := `69, 102
118, 274
150, 269
331, 284
128, 302
307, 192
238, 52
240, 339
111, 127
180, 156
248, 265
160, 69
58, 136
43, 235
154, 202
262, 189
309, 53
292, 67
335, 198
99, 199
224, 120
206, 313
359, 352
101, 147
301, 47
255, 347
121, 153
264, 343
252, 225
48, 90
312, 139
90, 277
203, 227
315, 328
330, 81
190, 191
89, 296
312, 255
218, 181
299, 149
151, 254
209, 212
42, 76
348, 183
333, 227
44, 210
293, 356
44, 132
175, 77
215, 109`

	lines := strings.Split(raw_str, "\n")

	var coords []coordT

	for _, line := range lines {
		words := strings.Split(line, ", ")
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		coord := coordT{x: x, y: y}
		coords = append(coords, coord)
	}

	min_x := coords[0].x
	max_x := coords[0].x
	min_y := coords[0].y
	max_y := coords[0].y
	for _, c := range coords {
		if c.x < min_x {
			min_x = c.x
		}
		if c.x > max_x {
			max_x = c.x
		}
		if c.y < min_y {
			min_y = c.y
		}
		if c.y > max_y {
			max_y = c.y
		}
	}

	var min_xs []int
	var min_ys []int
	var max_xs []int
	var max_ys []int
	for k, c := range coords {
		if c.x == min_x {
			min_xs = append(min_xs, k)
		}
		if c.x == max_x {
			max_xs = append(max_xs, k)
		}
		if c.y == min_y {
			min_ys = append(min_ys, k)
		}
		if c.y == max_y {
			max_ys = append(max_ys, k)
		}
	}

	dx := max_x - min_x
	dy := max_y - min_y

	min_i := -30
	max_i := dx+90
	min_j := -30
	max_j := dy+90

	dx = max_i - min_i
	dy = max_j - min_j

	grids := [][][]int{}

	for _, c := range coords {
		grid := [][]int{}
		for i := min_i; i < max_i; i++ {
			var ar []int
			for j := min_j; j < max_j; j++ {
				md := abs(c.x-i) + abs(c.y-j)
				ar = append(ar, md)
			}
			grid = append(grid, ar)
		}
		grids = append(grids, grid)
	}

	min_grid := [][]int{}

	for i := min_i; i < max_i; i++ {
		var ar []int
		for j := min_j; j < max_j; j++ {
			min_md := grids[0][i-min_i][j-min_j]+1
			min_md_k := -1
			for k := 0; k < len(grids); k++ {
				v := grids[k][i-min_i][j-min_j]
				if v < min_md {
					min_md = v
					min_md_k = k
				} else if (v == min_md) {
					min_md_k = -1
				}
			}
			ar = append(ar, min_md_k)
		}
		min_grid = append(min_grid, ar)
	}

	area_counts := make([]int, len(coords))

	for i := min_i; i < max_i; i++ {
		for j := min_j; j < max_j; j++ {
			v := min_grid[i-min_i][j-min_j]
			if v > 0 {
				area_counts[v]++
			}
			//st := string("A"[0] + byte(v % 26))
			//fmt.Print(st)
		}
		//fmt.Println()
	}

	exclude_list := make(map[int]bool)

	for i := min_i; i < max_i; i++ {
		v1 := min_grid[i-min_i][0]
		v2 := min_grid[i-min_i][max_j-min_j-1]
		exclude_list[v1] = true
		exclude_list[v2] = true
	}

	for j := min_j; j < max_j; j++ {
		v1 := min_grid[0][j-min_j]
		v2 := min_grid[max_i-min_i-1][j-min_j]
		exclude_list[v1] = true
		exclude_list[v2] = true
	}

	max_area := 0
	var max_area_k int
	for i, v := range area_counts {
		_, ok := exclude_list[i]
		if ok {
			fmt.Println(i,"is in exclude_list; skipping...")
			continue
		}
		if v > max_area {
			fmt.Println("found new max:", i, v)
			max_area = v
			max_area_k = i
		}
	}

	fmt.Println(len(grids))
	fmt.Println(max_area_k, max_area)
	fmt.Println(area_counts)


	fmt.Print(`<html>
	<head>
	</head>
	<body>
		<canvas id="myCanvas" width="`, dx ,`" height="`, dy ,`"></canvas>
		<script>
			var canvas = document.getElementById('myCanvas');
			var myContext = canvas.getContext('2d');
			var id = myContext.createImageData(` + strconv.Itoa(dx) + ", " + strconv.Itoa(dy) + `);
			var d  = id.data;
			function pp(x, y, r, g, b, a=256) {
				var w = id.width;
				var h = id.height;
				var offset = y*w*4+x*4;
				d[offset+0]   = r;
				d[offset+1]   = g;
				d[offset+2]   = b;
				d[offset+3]   = a;
			}
`)

	var colors []colorT
	for i := 0; i < len(coords); i++ {
		rand.Intn(256)
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		colors = append(colors, colorT{r: r, g: g, b: b})
	}

	for i := min_i; i < max_i; i++ {
		for j := min_j; j < max_j; j++ {
			gk := min_grid[i-min_i][j-min_j]
			var color colorT
			if gk < 0 {
				color.r = 0
				color.g = 0
				color.b = 0
			} else if gk == max_area_k {
				color.r = 50
				color.g = 50
				color.b = 50
			} else {
				color = colorT{r: colors[gk].r, g: colors[gk].g, b: colors[gk].b}
			}
			fmt.Print(`			pp(` + strconv.Itoa(i-min_i) + ", " + strconv.Itoa(j-min_j) + ", " + strconv.Itoa(color.r) + ", " + strconv.Itoa(color.g) + ", " + strconv.Itoa(color.b) + `);
`)
		}
	}

	fmt.Print(`			myContext.putImageData( id, 0, 0);
		</script>
	</body>
</html>`)

}

package main

// import "golang.org/x/tour/pic"

// you can view it by going to https://go.dev/tour/moretypes/18 and pasting code there

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
	}

	for i := range pic {
		for j := range pic[i] {

			// interesting images:
			// pic[i][j] = uint8((i + j) / 2)
			// pic[i][j] = uint8(i * j)
			pic[i][j] = uint8(i ^ j)
		}
	}

	return pic
}

// func main() {
// 	pic.Show(Pic)
// }

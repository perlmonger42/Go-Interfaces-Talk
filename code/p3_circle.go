package main
import ("image"; "image/png"; "image/color"; "log"; "os")

func main() {
  outFilename, m := makeSampleImage()
  outFile, err := os.Create(outFilename)
  if err != nil {
    log.Fatal(err)
  }
  defer outFile.Close()
  log.Print("Saving image to: ", outFilename)
  png.Encode(outFile, m)
}

func makeSampleImage() (string, image.Image) {
  hotpink := color.RGBA{255, 25, 155, 255}
  return "circle.png", Circle{Radius: 200, Color: hotpink}
}

type Circle struct {
  Radius int
  Color  color.Color
}

func (shape Circle) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape Circle) Bounds() image.Rectangle {
  r := shape.Radius
  return image.Rect(-r, -r, r, r)
}

func (shape Circle) At(x, y int) color.Color {
  r := shape.Radius
  if x*x + y*y < r*r {
    return shape.Color
  }
  return color.RGBA{0, 0, 0, 0}
}

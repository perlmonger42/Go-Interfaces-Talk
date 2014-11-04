package main
import ("image"; "image/png"; "image/color"; "log"; "math"; "os")

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
  return "gingham.png", Gingham{image.Point{850, 400}}
}

type Gingham struct {
  Size image.Point
}

func (shape Gingham) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape Gingham) Bounds() image.Rectangle {
  return image.Rectangle{image.ZP, shape.Size}
}

func (shape Gingham) At(x, y int) color.Color {
  sinx := math.Sin(float64(x)/25.0)
  cosy := math.Cos(float64(y)/25.0)
  return color.RGBA{
    uint8(math.Abs(sinx) * 255),
    uint8(math.Abs(cosy) * 255),
    128,
    255,
  }
}

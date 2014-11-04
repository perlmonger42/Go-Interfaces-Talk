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
  return "rectangle.png", Rectangle{}
}

type Rectangle struct { }

func (shape Rectangle) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape Rectangle) Bounds() image.Rectangle {
  return image.Rect(0, 0, 850, 400)
}

func (shape Rectangle) At(x, y int) color.Color {
  return color.RGBA{0, 0, 255, 255} // blue
}

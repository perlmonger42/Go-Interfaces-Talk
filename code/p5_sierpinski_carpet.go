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
  return "sierpinski.png", SierpinskiCarpet{}
}

type SierpinskiCarpet struct {
}

func (shape SierpinskiCarpet) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape SierpinskiCarpet) Bounds() image.Rectangle {
  return image.Rect(0, 0, 850, 400)
}

func (shape SierpinskiCarpet) At(x, y int) color.Color {
  for ; x > 0 || y > 0; x, y = x/3, y/3 {
    if x%3 == 1 && y%3 == 1 {
      return color.RGBA{0, 0, 0, 0}
    }
  }
  return color.RGBA{0,55,255,255};
}

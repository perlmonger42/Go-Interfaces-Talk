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
  return "colorings.png", ColorRings{Radius: 200}
}

type ColorRings struct {
  Radius int
}

func (shape ColorRings) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape ColorRings) Bounds() image.Rectangle {
  r := shape.Radius
  return image.Rect(-r, -r, r, r)
}

func (shape ColorRings) At(x, y int) color.Color {
  distance := int(math.Hypot(float64(x), float64(y)))
  return RainbowColor(distance * 6 / shape.Radius)
}

func RainbowColor(n int) color.Color {
  switch n {
  case 0: return color.RGBA{0xFF, 0x00, 0x00, 255} // Red
  case 1: return color.RGBA{0xFF, 0xA5, 0x00, 255} // Orange
  case 2: return color.RGBA{0xFF, 0xFF, 0x00, 255} // Yellow
  case 3: return color.RGBA{0x3C, 0xB3, 0x71, 255} // Medium Sea Green
  case 4: return color.RGBA{0x1E, 0x90, 0xFF, 255} // Dodger Blue
  case 5: return color.RGBA{0x93, 0x70, 0xDB, 255} // Medium Purple
  }
  return color.RGBA{0, 0, 0, 0}
}

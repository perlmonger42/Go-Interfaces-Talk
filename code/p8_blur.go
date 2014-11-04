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
  return "blur.png", Blur{
    Image: ColorRings{Radius: 200},
    BlurRadius: 8,
  }
}

type Blur struct {
  Image  image.Image
  BlurRadius int
}

func (Blur) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape Blur) Bounds() image.Rectangle {
  return shape.Image.Bounds()
}

func (shape Blur) At(x, y int) color.Color {
  var sumWeight, sumR, sumG, sumB float64
  radius := float64(shape.BlurRadius)
  for i := x - shape.BlurRadius; i < x+shape.BlurRadius+1; i++ {
    for j := y - shape.BlurRadius; j < y+shape.BlurRadius+1; j++ {
      dist := math.Hypot(float64(i-x), float64(j-y))
      if dist <= radius {
        weight := 1 / (1 + dist)
        r, g, b, _ := shape.Image.At(i, j).RGBA()
        sumR += float64(r) * weight
        sumG += float64(g) * weight
        sumB += float64(b) * weight
        sumWeight += weight
      }
    }
  }
  _, _, _, a := shape.Image.At(x, y).RGBA()
  return color.RGBA{
    uint8(sumR / sumWeight * (255.0 / 65535)),
    uint8(sumG / sumWeight * (255.0 / 65535)),
    uint8(sumB / sumWeight * (255.0 / 65535)),
    uint8(a * 255 / 65535),
  }
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

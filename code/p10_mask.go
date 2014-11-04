package main
import (
  "image"
  "image/draw"
  "image/jpeg"
  "image/png"
  "image/color"
  "log"
  "os"
)

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
  gopher := loadGopher()
  rect := gopher.Bounds()
  carpet := SierpinskiCarpet{rect.Size(), image.Point{0, 200}}

  // Copy the gopher into an RGBA image, using Sierpinski mask.
  result := image.NewRGBA(rect)
  draw.DrawMask(result, rect,
                gopher, rect.Min,
                carpet, image.ZP,
                draw.Over)
  return "mask.png", result
}

func loadGopher() image.Image {
  inFilename := "gophers1.jpg"
  inFile, err := os.Open(inFilename)
  if err != nil { log.Fatal(err) }
  defer inFile.Close()
  log.Print("Reading image from: ", inFilename)
  gopher, err := jpeg.Decode(inFile)
  if err != nil { log.Fatal(err) }
  return gopher
}


type SierpinskiCarpet struct {
  Size image.Point
  Offset image.Point
}

func (shape SierpinskiCarpet) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape SierpinskiCarpet) Bounds() image.Rectangle {
  return image.Rectangle{image.ZP, shape.Size}
}

func (shape SierpinskiCarpet) At(x, y int) color.Color {
  x += shape.Offset.X
  y += shape.Offset.Y
  for ; x > 0 || y > 0; x, y = x/3, y/3 {
    if x%3 == 1 && y%3 == 1 {
      return color.RGBA{255, 255, 255, 255}
    }
  }
  return color.RGBA{0,0,0,0};
}

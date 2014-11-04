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
  circle := Circle{Radius: 100, Color: RainbowColor(3)}
  mandelbrot := Mandelbrot{
    Size: image.Point{350, 200},
    Cx:   -0.75,
    Cy:   -0.0,
    Zoom: 1,
  }
  gingham := Gingham{Size: image.Point{175, 175}}
  composedImage := hcAppend(circle, mandelbrot, gingham)
  return "composite.png", composedImage
}

type Composite struct {
  images []image.Image
  bounds []image.Rectangle
}

func (Composite) ColorModel() color.Model {
  return color.RGBAModel
}

func (shape Composite) Bounds() image.Rectangle {
  if len(shape.images) == 0 {
    return image.Rect(0, 0, 0, 0)
  }
  bounds := shape.bounds[0]
  for _, b := range shape.bounds {
    bounds = bounds.Union(b)
  }
  return bounds
}

func (shape Composite) At(x, y int) color.Color {
  p := image.Point{x, y}
  for i, b := range shape.bounds {
    if p.In(b) {
      img := shape.images[i]
      ib := img.Bounds()
      p = p.Sub(b.Min).Add(ib.Min)
      return img.At(p.X, p.Y)
    }
  }
  return color.RGBA{0, 0, 0, 0}
}

func NewComposite() Composite {
  return Composite{make([]image.Image, 0), make([]image.Rectangle, 0)}
}

func maxheight(images []image.Image) int {
  tallest := 0
  for _, img := range images {
    if h := img.Bounds().Dy(); h > tallest {
      tallest = h
    }
  }
  return tallest
}

func maxwidth(images []image.Image) int {
  widest := 0
  for _, img := range images {
    if w := img.Bounds().Dx(); w > widest {
      widest = w
    }
  }
  return widest
}

// hcAppend builds a composite image from a list of images, arranging them
// horizontally from left to right, with adjacent images touching, and with
// their vertical centers aligned. The 'h' in 'hc' means horizontally-arranged,
// and the 'c' in 'hc' means center-aligned.
func hcAppend(images ...image.Image) image.Image {
  shape := NewComposite()
  tallest := maxheight(images)
  left := 0
  for _, img := range images {
    sz := img.Bounds().Size()
    w, h := sz.X, sz.Y
    top := (tallest - h) / 2
    shape.images = append(shape.images, img)
    shape.bounds = append(shape.bounds, image.Rect(left, top, left+w, top+h))
    left += w
  }
  return shape
}

// vcAppend builds a composite image from a list of images, arranging them
// vertically from top to bottom, with adjacent images touching, and with their
// horizontal centers aligned. The 'v' in 'vc' means vertically-arranged, and
// the 'c' in 'vc' means center-aligned.
func vcAppend(images ...image.Image) image.Image {
  shape := NewComposite()
  widest := maxwidth(images)
  top := 0
  for _, img := range images {
    sz := img.Bounds().Size()
    w, h := sz.X, sz.Y
    left := (widest - w) / 2
    shape.images = append(shape.images, img)
    shape.bounds = append(shape.bounds, image.Rect(left, top, left+w, top+h))
    top += h
  }
  return shape
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

type Mandelbrot struct {
  Size  image.Point
  Cx    float64
  Cy    float64
  Zoom  float64
}

func (shape Mandelbrot) ColorModel() color.Model {
  return color.RGBAModel
}

// The Mandelbrot set is roughly in the area -2.5 < x < 1 and -1 < y < 1.
func (shape Mandelbrot) Bounds() image.Rectangle {
  return image.Rectangle{image.ZP, shape.Size}
}

func (shape Mandelbrot) At(xi, yi int) color.Color {
  x0 := float64(xi-shape.Size.X/2)/(shape.Zoom*100) + shape.Cx
  y0 := float64(yi-shape.Size.Y/2)/(shape.Zoom*100) + shape.Cy
  x, y := x0, y0
  xx, yy := x*x, y*y

  iteration := 0
  max_iterations := 1000
  for xx+yy < 4 && iteration < max_iterations {
    x, y = xx-yy+x0, 2*x*y+y0
    xx, yy = x*x, y*y
    iteration += 1
  }
  return ThousandColors(iteration)
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

func ThousandColors(n int) color.Color {
  // Interpolate a thousand colors between yellow and purple.
  if n < 0 || n >= 1000 {
    return color.RGBA{0, 0, 0, 0}
  }
  const W = 250 // band width: number of values to map between 2 adjacent colors
  i := n / W    // convert 0..999 to 0..3
  c1 := RainbowColor(i + 2)
  c2 := RainbowColor(i + 3)
  // interpolate between c1 and c2
  d := uint32(n % W) // d is 0..W-1, the relative distance between c1 & c2
  d2 := (W - 1) - d  // d2 is W-1..0, the relative distance between c2 & c1
  r1, g1, b1, a1 := c1.RGBA()
  r2, g2, b2, a2 := c2.RGBA()
  r := ((r1*d2 + r2*d) / W) * 255 / 65535
  g := ((g1*d2 + g2*d) / W) * 255 / 65535
  b := ((b1*d2 + b2*d) / W) * 255 / 65535
  a := ((a1*d2 + a2*d) / W) * 255 / 65535
  return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

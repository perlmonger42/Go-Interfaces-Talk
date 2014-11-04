package main
import ("image"; "image/png"; "log"; "os")

func main() {
  size := image.Point{850, 400}
  m := image.NewRGBA(image.Rectangle{image.ZP, size})
  outFilename := "transparent.png"
  outFile, err := os.Create(outFilename)
  if err != nil {
    log.Fatal(err)
  }
  defer outFile.Close()
  log.Print("Saving image to: ", outFilename)
  png.Encode(outFile, m)
}

package main
//
//import (
//	_ "image/png"
//	"log"
//
//	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
//)
//
//var img *ebiten.Image
//
//func init() {
//	var err error
//	img, _, err = ebitenutil.NewImageFromFile("drphil.png")
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//type Game struct{}
//
//func (g *Game) Update() error {
//	return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	op.GeoM.Translate(435/2, 300/2)
//	op.GeoM.Scale(1, 1)
//	screen.DrawImage(img, op)
//}
//
//func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return 435, 600
//}
//
//func main() {
//	ebiten.SetWindowSize(435, 600)
//	ebiten.SetWindowTitle("Geometry Matrix")
//	if err := ebiten.RunGame(&Game{}); err != nil {
//		log.Fatal(err)
//	}
//}

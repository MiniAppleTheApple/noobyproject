package sys
import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"
	"os"
	"strconv"
	//"fmt"
	//"time"
)
var (
	player GameObject
	movepermission map[string]bool
	//test GameEnemy
	movetemplate  map[string]bool
	font rl.Font
	background1 GameObject
	background2 GameObject
	enemylist []GameObject
	bulletcollection [][]GameObject
	enemycollection [][]GameObject
	imagetable = make(map[string]rl.Texture2D)
	moveleft = true
	times = 0
)
type GameMove struct {
	left bool
	right bool
	up bool
	down bool
}
type GameObject struct {
	x int32
	y int32
	texture rl.Texture2D
	size [2]int32
	vel int32
	movepermission GameMove
	atk int32
	hp int32
}
const (
	WinWidth int32 = 800
	WinHeight int32 = 600
)
type Map struct{
	list []int32
}
func (object GameObject)Draw(){
	rl.DrawTexture(object.texture,object.x,object.y,rl.White)
}
func (object *GameObject)Move(x int32,y int32){
	object.x += x
	object.y += y
}
func Draw(){
	background1.Draw()
	background2.Draw()
	text := strconv.Itoa(int(player.hp))
	rl.DrawTextEx(font, text, rl.Vector2{ 20.0, 100.0 }, float32(32), 2, rl.Lime)
	player.Draw()
	for i := 0;i < len(bulletcollection);i++{
		AllDraw(bulletcollection[i])
	}
	AllDraw(enemylist)
}
func DrawIt(object GameObject){
	rl.DrawTexture(object.texture,object.x,object.y,rl.White)
}
func Onload(){
	imagetable["player"] = rl.LoadTexture("resource/texture/character.png")
	imagetable["enemy"] = rl.LoadTexture("resource/texture/enemy.png")
	imagetable["enemy2"] = rl.LoadTexture("resource/texture/enemy2.png")
	imagetable["enemy3"] = rl.LoadTexture("resource/texture/enemy3.png")
	for i := 0;i <= 9;i++{
		imagetable["bullet" + strconv.Itoa(i)] = rl.LoadTexture("resource/texture/bullet" + strconv.Itoa(i) + ".png")
	}
	imagetable["null"] = rl.LoadTexture("resource/texture/null.png")
	imagetable["bulletex"] = rl.LoadTexture("resource/texture/bulletex.png")
	imagetable["bulletblack"] = rl.LoadTexture("resource/texture/bulletblack.png")
	player = GameObject{WinWidth / 2,WinHeight / 2,imagetable["player"],[2]int32{32,32},6,GameMove{false,false,false,false},10,3}
	enemylist = []GameObject{GameObject{0,0,imagetable["enemy3"],[2]int32{32,32},5,GameMove{false,false,false,false},10,10}}
	font = rl.LoadFont("resource/UbuntuMono-R.ttf")
	background2 = GameObject{
	0,
	-600,
	rl.LoadTexture("resource/texture/background1.png"),
	[2]int32{0,0},
	3,
	GameMove{
	false,
	false,
	false,
	false},10,10}
	background1 = GameObject{
	0,
	0,
	rl.LoadTexture("resource/texture/background1.png"),
	[2]int32{0,0},
	3,
	GameMove{
	false,
	false,
	false,
	false},10,10}
	enemylist = append(enemylist,GameObject{0,0,imagetable["enemy3"],[2]int32{32,32},2,GameMove{false,false,false,false},10,10})
	bulletlist := []GameObject{GameObject{400,400,imagetable["bulletblack"],[2]int32{32,32},2,GameMove{false,false,false,false},10,10}}
	bulletlist2 := []GameObject{GameObject{400,400,imagetable["bullet5"],[2]int32{32,32},2,GameMove{false,false,false,false},10,10}}
	bulletlist3 := []GameObject{GameObject{400,400,imagetable["bulletblack"],[2]int32{32,32},2,GameMove{false,false,false,false},10,10}}
	bulletlist4 := []GameObject{GameObject{400,400,imagetable["bullet5"],[2]int32{32,32},2,GameMove{false,false,false,false},10,10}}
	for i := 0;i < 10;i++{
		bulletlist = append(bulletlist,GameObject{Random(WinWidth),Random(WinHeight),imagetable["bulletblack"],[2]int32{32,32},3,GameMove{false,false,false,false},10,10})
		bulletlist2 = append(bulletlist2,GameObject{Random(WinWidth),Random(WinHeight),imagetable["bullet5"],[2]int32{32,32},3,GameMove{false,false,false,false},10,10})
		bulletlist3 = append(bulletlist3,GameObject{Random(WinWidth),Random(WinHeight),imagetable["bulletblack"],[2]int32{32,32},3,GameMove{false,false,false,false},10,10})
		bulletlist4 = append(bulletlist4,GameObject{Random(WinWidth),Random(WinHeight),imagetable["bullet5"],[2]int32{32,32},3,GameMove{false,false,false,false},10,10})
	}
	bulletcollection = append(bulletcollection,bulletlist)
	bulletcollection = append(bulletcollection,bulletlist2)
	bulletcollection = append(bulletcollection,bulletlist3)
	bulletcollection = append(bulletcollection,bulletlist4)
	for i := 0;i < 10;i++{
		enemylist = append(enemylist,GameObject{Random(WinWidth),Random(WinHeight),imagetable["enemy3"],[2]int32{32,32},5,GameMove{false,false,false,false},10,10})
	}
	Shoot(bulletcollection[0],enemylist[0])
	Shoot(bulletcollection[1],enemylist[1])
}
func Update(){
	Movecheck()
	WindowCollision()
	BackgroundMovement()
	Move()
	PlayerCollision()

	
}
func PlayerCollision(){
	for i := 0;i < len(bulletcollection);i++{
		if AllCollision(player,bulletcollection[i]){
			*&player.hp -= 1
		}
		if player.hp < 0{
			os.Exit(0)
		}
	}
}
func Movecheck(){
	if (rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)){
		player.movepermission.left = true
	}else{
		player.movepermission.left = false
	}
	if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)){
		player.movepermission.right = true
	}else{
		player.movepermission.right = false
	}
	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)){
		player.movepermission.up = true
	}else{
		player.movepermission.up = false
	}
	if (rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)){
		player.movepermission.down = true
	}else{
		player.movepermission.down = false
	}
}
func Move(){
	if (player.movepermission.left){
		player.x -= player.vel
	}
	if (player.movepermission.right){
		player.x += player.vel
	}
	if (player.movepermission.down){
		player.y += player.vel
	}
	if (player.movepermission.up){
		player.y -= player.vel
	}
}
func WindowCollision(){
	if player.x > 600 - player.size[0]{
		player.movepermission.right = false
	}
	if player.x < 0{
		player.movepermission.left = false
	}
	if player.y > 600 - player.size[1]{
		player.movepermission.down = false
	}
	if player.y < 0{
		player.movepermission.up = false
	}
}
func DistanceCal(x1 int32,y1 int32,x2 int32,y2 int32) float64{
	d := math.Sqrt(math.Pow(float64(x1) - float64(x2),2) + math.Pow(float64(y1) - float64(y2),2))
	return d
}
func ObjectDistanceCal(obj1 GameObject,obj2 GameObject) float64{
	d := math.Sqrt(math.Pow(float64(obj1.x) - float64(obj2.x),2) + math.Pow(float64(obj1.y) - float64(obj2.y),2))
	return d
}
func BackgroundMovement(){
	if background1.y >= 600{
		background1.y = -600
	}
	if background2.y >= 600{
		background2.y = -600
	}
	background1.y += background1.vel
	background2.y += background2.vel
}
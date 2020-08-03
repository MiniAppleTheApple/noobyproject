
package sys
import (
	//"github.com/gen2brain/raylib-go/raylib"
	//"math/rand"
)
func AllCollision(obj GameObject,list []GameObject) bool{
	pre_list := &list[0]
	pre_obj := &obj
	for i := 0;i <= len(list) - 1;i++{
		pre_list = &list[i]
		if DistanceCal(pre_list.x,pre_list.y,pre_obj.x,pre_obj.y) <= 8{
			return true
			break;
		}
	}
	return false
}
func AllDraw(list []GameObject){
	pre_list := &list[0]
	for i := 0;i <= len(list) - 1;i++{
		pre_list = &list[i]
		pre_list.Draw()
	}
}
func AllMove(x int32,y int32,list []GameObject){
	pre_list := &list[0]
	for i := 0;i <= len(list) - 1;i++{
		pre_list = &list[i]
		pre_list.Move(x,y)
	}
}
func AllRandomMove(x int32,y int32,list []GameObject){
	pre_list := &list[0]
	for i := 0;i <= len(list) - 1;i++{
		r := Random(1)
		b := Random(1)
		if r == 0{
			r = Random(x)
		}else if r == 1{
			r = -Random(x)
		}
		if b == 0{
			b = Random(y)
		}else if b == 1{
			b = -Random(y)
		}
		pre_list = &list[i]
		pre_list.Move(r,b)
	}
}
func AllMoveTo(x int32,y int32,list []GameObject){
	pre_list := &list[0]
	for i := 0;i <= len(list) - 1;i++{
		pre_list = &list[i]
		pre_list.x = x
		pre_list.y = y
	}
}
func CheckOutScreen(list []GameObject) (bool,GameMove){
	leftboolean := len(list) - 1
	rightboolean := len(list) - 1
	downboolean := len(list) - 1
	upboolean := len(list) - 1
	boolean := len(list) - 1
	for i := 0;i <= len(list) - 1;i++{
		if list[i].x > WinWidth - list[i].size[0]{
			leftboolean -= 1
			boolean -= 1
		}
		if list[i].x < 0{
			rightboolean -= 1
			boolean -= 1
		}
		if list[i].y > WinHeight - list[i].size[1]{
			downboolean -= 1
			boolean -= 1
		}
		if list[i].y < 0{
			upboolean -= 1
			boolean -= 1
		}
	}
	if boolean == 0{
		return true,GameMove{leftboolean == 0,rightboolean == 0,upboolean == 0,downboolean == 0}
	}
	return false,GameMove{leftboolean == 0,rightboolean == 0,upboolean == 0,downboolean == 0}	
}
func CheckOutScreenSimple(list []GameObject) bool{
	leftboolean := len(list) - 1
	rightboolean := len(list) - 1
	downboolean := len(list) - 1
	upboolean := len(list) - 1
	boolean := len(list) - 1
	for i := 0;i <= len(list) - 1;i++{
		if list[i].x > 600 - list[i].size[0]{
			leftboolean -= 1
			boolean -= 1
		}
		if list[i].x < 0{
			rightboolean -= 1
			boolean -= 1
		}
		if list[i].y > 600 - list[i].size[1]{
			downboolean -= 1
			boolean -= 1
		}
		if list[i].y < 0{
			upboolean -= 1
			boolean -= 1
		}
	}
	if boolean == 0{
		return true
	}
	return false
}
func CheckOutScreenX(list []GameObject) GameMove{
	leftboolean := len(list) - 1
	rightboolean := len(list) - 1
	for i := 0;i <= len(list) - 1;i++{
		if list[i].x > 600 - list[i].size[0]{
			leftboolean -= 1
		}
		if list[i].x < 0{
			rightboolean -= 1
		}
	}
	return GameMove{leftboolean - leftboolean / 10 < 0 ,rightboolean - rightboolean / 10 < 0,false,false}	
}
func CheckOutScreenNumber(list []GameObject) []int32{
	leftboolean := len(list) - 1
	rightboolean := len(list) - 1
	downboolean := len(list) - 1
	upboolean := len(list) - 1
	for i := 0;i <= len(list) - 1;i++{
		if list[i].x > 600 - list[i].size[0]{
			leftboolean -= 1
		}
		if list[i].x < 0{
			rightboolean -= 1
		}
		if list[i].y > 600 - list[i].size[1]{
			downboolean -= 1
		}
		if list[i].y < 0{
			upboolean -= 1
		}
	}
	return []int32{int32(leftboolean),int32(rightboolean),int32(upboolean),int32(downboolean)}
}
func MoveLogicFollow(list []GameObject,obj GameObject){
	for i := 0;i < len(list) - 1;i++{
		if obj.x > list[i].x{
			*&list[i].x += list[i].vel
		}
		if obj.x < list[i].x{
			*&list[i].x -= list[i].vel
		}
		if obj.y > list[i].y{
			*&list[i].y += list[i].vel
		}
		if obj.y < list[i].y{
			*&list[i].y -= list[i].vel
		}
	}
}
func MoveLogic1(list []GameObject,oper string){
	for i := 0;i < len(list);i++{
		if oper == "+"{
			*&list[i].x += list[i].vel 
		}
		if oper == "-"{
			*&list[i].x -= list[i].vel 
		}
		if oper == "*"{
			*&list[i].x *= list[i].vel 
		}
		if oper == "/"{
			*&list[i].x /= list[i].vel 
		}
	}
}
func (obj GameObject)Shoot(list []GameObject){
	for i := 0;i <= len(list) - 1;i++{
		*&list[i].x = obj.x
	}
}
func Shoot(list []GameObject,obj GameObject){
	for i := 0;i < len(list);i++{
		*&list[i].x = obj.x
		*&list[i].y = obj.y
	}
}
func MoveLogicFollow2(list []GameObject,obj GameObject,dis int32){
	for i := 0;i < len(list);i++{
		if int32(DistanceCal(list[i].x,list[i].y,obj.x,obj.y)) >= dis + obj.size[0]{
			if obj.x > list[i].x{
				*&list[i].x += list[i].vel
			}
			if obj.x < list[i].x{
				*&list[i].x -= list[i].vel
			}
			if obj.y > list[i].y{
				*&list[i].y += list[i].vel
			}
			if obj.y < list[i].y{
				*&list[i].y -= list[i].vel
			}
		}
		if int32(DistanceCal(list[i].x,list[i].y,obj.x,obj.y)) <= dis + obj.size[0]{
			break
		}
	}
}
func MoveLogicFollow3(y int32,x int32,list []GameObject,obj GameObject,dis int32){
	for i := 0;i < len(list);i++{
		if int32(DistanceCal(list[i].x,list[i].y,obj.x,obj.y)) >= dis + obj.size[0]{
			if obj.x > list[i].x{
				*&list[i].x += list[i].vel
			}
			if obj.x < list[i].x{
				*&list[i].x -= list[i].vel
			}
			if obj.y > list[i].y{
				*&list[i].y += list[i].vel
			}
			if obj.y < list[i].y{
				*&list[i].y -= list[i].vel
			}
		}else{
			r := Random(1)
			a := Random(1)
			if r == 1{
				list[i].x += Random(list[i].vel * x)
			}
			if r == 0{
				list[i].x -= Random(list[i].vel * x)
			}
			if a == 1{
				list[i].y += Random(list[i].vel * y)
			}
			if a == 0{
				list[i].y -= Random(list[i].vel * y)
			}
		}
	}
}
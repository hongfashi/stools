package stools

import (
	"io"
	"image"
	"fmt"
	"os"
	"math"
	"image/png"
	"crypto/md5"
)

func WriteWithIoutil(name string,content string,append int) {
	
	var method int 
	
	if append == 1 {
	
		method = os.O_APPEND|os.O_CREATE|os.O_WRONLY
	
	}else{
	
		method = os.O_CREATE|os.O_WRONLY
		
	}

	f, err := os.OpenFile(name, method, 0644)
	
	if err != nil {
		// 打开文件失败处理

	} else {
      
		// 从末尾的偏移量开始写入内容
		_, err = f.Write([]byte(content))
	}

	defer f.Close()

}

/**
  *根据圆的半径，坐标x,y中的一个值，推出另一个值
 *return(左侧，右侧)
 */
func CircleYtoX(y float64,r float64)(int,int){
	x1 := r - math.Sqrt(math.Pow(r,2)-math.Pow(y-r,2))
	x2 := r +  math.Sqrt(math.Pow(r,2)-math.Pow(y-r,2))
	return int(x1+0.5),int(x2+0.5)
}


func StringMd5(str string) string {
    w := md5.New()
    io.WriteString(w, str)
    md5str := fmt.Sprintf("%x", w.Sum(nil))
    return md5str
}

/**
*@desc 剪切图片
**/
func ImageCut(sourceName string,targetName string,x1 int,y1 int,x2 int,y2 int){

	source, _ := os.Open(sourceName)
	
	defer source.Close()
	  
	dist, _ := os.Create(targetName)
	
	defer dist.Close()
	// 图片文件解码
	m, _:= png.Decode(source)     
             
	rgbImg := m.(*image.NRGBA)
	subImg := rgbImg.SubImage(image.Rect(x1, y1, x2, y2)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	png.Encode(dist, subImg)       //写入文件
	
}
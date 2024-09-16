package main

/*
#include <stdio.h>

void PutChar(char ch){
	putchar(ch);
}
*/
import "C"
import (
	"fmt"
	"math"
	"time"
)

var A, B, c float64
var cubeWidth float64
var width, height int = 160, 44
var zBuffer [160 * 44 * 4]float64
var buffer [160 * 44]rune
var backgrountASCIICode rune = ' '
var distanceFromCam float64 = 100
var horizontalOffset float64

var incrementSpeed float64 = 0.8

var x, y, z, ooz float64
var K1 float64 = 40
var xp, yp, idx int

func main() {
	fmt.Print("\033[2J")
	fmt.Print("\033[?25l")
	for {
		memsetBuffer(&buffer, backgrountASCIICode)
		memsetZBuffer(&zBuffer, '\000')
		// First Cube
		cubeWidth = 20
		horizontalOffset = -2 * cubeWidth
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateSurface(cubeWidth, cubeY, cubeX, '$')
				calculateSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		// Second
		cubeWidth = 10
		horizontalOffset = 1 * cubeWidth
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateSurface(cubeWidth, cubeY, cubeX, '$')
				calculateSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}
		// Third
		cubeWidth = 5
		horizontalOffset = 8 * cubeWidth
		for cubeX := -cubeWidth; cubeX < cubeWidth; cubeX += incrementSpeed {
			for cubeY := -cubeWidth; cubeY < cubeWidth; cubeY += incrementSpeed {
				calculateSurface(cubeX, cubeY, -cubeWidth, '@')
				calculateSurface(cubeWidth, cubeY, cubeX, '$')
				calculateSurface(-cubeWidth, cubeY, -cubeX, '~')
				calculateSurface(-cubeX, cubeY, cubeWidth, '#')
				calculateSurface(cubeX, -cubeWidth, -cubeY, ';')
				calculateSurface(cubeX, cubeWidth, cubeY, '+')
			}
		}

		fmt.Print("\033[H")
		for k := 0; k < width*height; k++ {
			if k%width > 0 {

				C.PutChar(C.char(buffer[k]))

			} else if k%width == 0 {

				C.PutChar(10)

			}
		}
		A += 0.005
		B += 0.005
		c += 0.001
		time.Sleep(time.Microsecond * 16)
	}
}

func calculateSurface(cubeX, cubeY, cubeZ float64, ch rune) {
	x = calculateX(cubeX, cubeY, cubeZ)
	y = calculateY(cubeX, cubeY, cubeZ)
	z = calculateZ(cubeX, cubeY, cubeZ) + distanceFromCam

	ooz = 1 / z
	xp = int(float64(width)/2 + horizontalOffset + K1*ooz*x*2)
	yp = int(float64(height)/2 + K1*ooz*y)

	idx = xp + yp*width
	if idx >= 0 && idx < width*height {
		if ooz > zBuffer[idx] {
			zBuffer[idx] = ooz
			buffer[idx] = ch
		}
	}
}

func calculateX(i float64, j float64, k float64) float64 {
	return j*math.Sin(A)*math.Sin(B)*math.Cos(c) - k*math.Cos(A)*math.Sin(B)*math.Cos(c) + j*math.Cos(A)*math.Sin(c) + k*math.Sin(A)*math.Sin(c) + i*math.Cos(B)*math.Cos(c)
}

func calculateY(i float64, j float64, k float64) float64 {
	return j*math.Cos(A)*math.Cos(c) + k*math.Sin(A)*math.Cos(c) - j*math.Sin(A)*math.Sin(B)*math.Sin(c) + k*math.Cos(A)*math.Sin(B)*math.Sin(c) - i*math.Cos(B)*math.Sin(c)
}

func calculateZ(i float64, j float64, k float64) float64 {
	return k*math.Cos(A)*math.Cos(B) - j*math.Sin(A)*math.Cos(B) + i*math.Sin(B)
}

func memsetBuffer(str *[160 * 44]rune, c rune) {
	lenth := len(str)
	for i := 0; i < lenth; i++ {
		str[i] = c
	}
}

func memsetZBuffer(str *[160 * 44 * 4]float64, c float64) {
	lenth := len(str)
	for i := 0; i < lenth; i++ {
		str[i] = c
	}
}

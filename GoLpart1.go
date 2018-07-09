package main

import (
	"image"
	"image/color"
	"image/jpeg"
	//"image/draw"
	"os"
	"fmt"
	"./Rules"
	"strconv"
	"math/rand"
	"time"
	//"os/exec"

)


//
// Define the cell data struct. This technically isn't an 'object' cause of GO but I'm gonna refer to it as an object later.
//
type Cell struct {
	X int // X coordinate in the cell array
	Y int // Y coordinate in the cell array
	alive bool // Boolean to represent wether the cell is CURRENTLY alive or dead
	index int // index number in the 1D cell array... this comes in handy when iterating. 
	neighbors int // Number of neighbors that are alive
}


//
// Function to update the image no matter what our cell size is as long as we understand it is a square.
//
func imgUpdate(img *image.RGBA, cells []Cell, cellSize int) *image.RGBA{
	for i := 0; i < len(cells); i ++ {
		for y := cells[i].Y; y < cells[i].Y + cellSize; y ++ {
			for x := cells[i].X; x < cells[i].X + cellSize; x ++ {
				if cells[i].alive {
					img.Set(x, y, color.White)
				} else {
					img.Set(x, y, color.Black)
				}
			}
		}
	}
	return img
}


func checkUp(cells []Cell, cell Cell, numCol int) int{
	if cell.index > numCol {
		if cells[cell.index - numCol].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkDown(cells []Cell, cell Cell, numCol int) int{
	if cell.index < numCol*numCol - numCol {
		if cells[cell.index + numCol].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkRight(cells []Cell, cell Cell, numCol int) int{
	if cell.index + 1 < numCol*numCol && cells[cell.index + 1].X > cell.X {
		if cells[cell.index + 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkLeft(cells []Cell, cell Cell, numCol int) int{
	if cell.index > 0 && cells[cell.index - 1].X < cell.X {
		if cells[cell.index - 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkUpRight(cells []Cell, cell Cell, numCols int) int{
	if cell.index - numCols + 1 > 0 && cells[cell.index - numCols + 1].X > cell.X{
		if cells[cell.index - numCols + 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkUpLeft(cells []Cell, cell Cell, numCols int) int{
	if cell.index - numCols - 1 > 0 && cells[cell.index - numCols - 1].X < cell.X {
		if cells[cell.index - numCols - 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkDownRight(cells []Cell, cell Cell, numCells int) int{
	if cell.index + numCells + 1 < numCells*numCells && cells[cell.index + numCells + 1].X > cell.X {
		if cells[cell.index + numCells + 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func checkDownLeft(cells []Cell, cell Cell, numCells int) int{
	if cell.index + numCells - 1 < numCells*numCells && cells[cell.index + numCells - 1].X < cell.X{
		if cells[cell.index + numCells - 1].alive {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func countNeighbors(cells []Cell, numCells int) []Cell{
	for i := 0; i < len(cells); i++ {
		cells[i].neighbors = checkUp(cells, cells[i], numCells) + checkDown(cells, cells[i], numCells) + checkRight(cells, cells[i], numCells) + checkLeft(cells, cells[i], numCells) + checkUpRight(cells, cells[i], numCells) + checkUpLeft(cells, cells[i], numCells) + checkDownRight(cells, cells[i], numCells) + checkDownLeft(cells, cells[i], numCells)
	}
	return cells
}

func numAlive(cells []Cell) int{
	numCellsAlive := 0
	for i := 0; i < len(cells); i++ {
		if cells[i].alive {
			numCellsAlive += 1
		}
	}
	return numCellsAlive
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	//
	// The Image is always square, so the cellSize variable indicates the pixel density of each cell. 
	// cellSize = 1 would mean each pixel is a square. 
	// cellSize = 2 would mean each square is 2 pixels high and 2 pixels wide.
	//
	cellSize := 2

	// Resolution of the output image frames i.e. 1080x1080 in this case (instagram optimization)
	imgResolution := 1080

	//
	// Bad naming convention, this actually isn't the number of cells, it actually the number of cells in a row or column
	// The true number of cells is numCells * numCells
	//
	numCells := imgResolution / cellSize

	initialArray := []int{}

	// This is initialized to 2 because the first frame is hardcoded (see the next to lines of code).
	frameNum := 2


	// Create the primary image object to hold the pixel data
	img := image.NewRGBA(image.Rect(0,0,imgResolution,imgResolution))
	// See how the first frame is hardcoded to 1? Yeah I could've used the variable but eh...
	file1, err := os.Create("frames/frame1.jpg")
	if err != nil {
		
	}

	//
	// Create the 1D array of cells
	// Note that the rows and columns both iterate to numCells so it will be a square.
	//
	cells := []Cell{}
	index := 0
	for i := 0; i < numCells; i ++ {
		for j := 0; j < numCells; j++ {
			cell := Cell{j * cellSize, i * cellSize, false, index, 0}
			cells = append(cells, cell)
			index += 1
		}
	}


	// Choose a random number of cells to initialize to true between a certain range
	randNum := 0
	for {
		randNum = rand.Intn(numCells*numCells - 1)
		if (randNum > 15000 && randNum < 180000){
			break
		}
	}

	// Now set that random number of cells to true and make sure it always sets a unique cell aka doesn't repeat the same cell
	for i := 0; i < randNum; i++ {
		for {
			n := rand.Intn(numCells*numCells - 1)
			unique := true
			for j := 0; j < len(initialArray); j++ {
				if initialArray[j] == n {
					unique = false
				}
			}
			if unique {
				initialArray = append(initialArray, n)
				cells[n].alive = true
				break
			}
		}
	}


	// Update the image object with the initialized version
	img = imgUpdate(img, cells, cellSize)

	// Write the image to disk
	jpeg.Encode(file1, img, &jpeg.Options{80})

	// Now count how many neighbors each cell has for the first iteration of the primary simulation loop
	cells = countNeighbors(cells, numCells);


	// 
	// Hey Hey we finally made it to the main loop.
	// The i < blah number indicates the number of frames to capture as part of the simulation.
	// Keep in mind though that the loop will break if all the cells are dead. No point in capturing more frames at that point.
	//
	for i := 0; i < 400; i++ {
		// Print the current frame number
		fmt.Println("Frame: ", i+2)

		// Check and see that we still have at least 1 live cell
		if numAlive(cells) == 0 {
			break
		}

		// Update the filename using the frameNum variable
		fileName := "frames/frame" + strconv.Itoa(frameNum) + ".jpg"

		// Create the file for this frame
		file, err := os.Create(fileName)
		if err != nil {
			// Haha you really expect me to handle errors with a simple script like this?
		}
		//defer file.Close()

		// Alright lets loop through all the cells
		for i := 0; i < len(cells); i++ {
			// We have 2 rules so lets figure out if we need to call the alive rule or dead rule
			if cells[i].alive {
				cells[i].alive = Rules.CheckStillAlive(cells[i].neighbors)
			} else {
				cells[i].alive = Rules.CheckStillDead(cells[i].neighbors)
			}
		}
		// Create the tremp object. This will be destroyed at each loop iteration
		tempImg := image.NewRGBA(image.Rect(0,0,imgResolution,imgResolution))

		// Update the image with the cell data
		tempImg = imgUpdate(img, cells, cellSize)

		// Write the frame file
		jpeg.Encode(file, img, &jpeg.Options{80})

		// Don't forget to close the file since we commented out the defer line above for some dumb reason. 
		file.Close()

		// I honestly don't remember why I did this but eh...
		img = tempImg

		// Iterate that frame number please
		frameNum += 1

		// Count the number of neighbors before the next iteration
		cells = countNeighbors(cells, numCells);
	}
	
	//cmd := exec.Command("ffmpeg", "-r", "5/1", "-i", "frames/frame%1d.jpg", "-vcodec","mpeg4","vid.mp4")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	defer file1.Close()
	//defer file2.Close()
	//jpeg.Encode(file2, img, &jpeg.Options{80})
}


/*

knight tour puzzle  17.2.2018 George Loo




*/


package main
 
import "fmt"

const (
	N = 8
	
)

var (
	sol[N][N] int
	calls int
)

/* A utility function to check if i,j are valid indexes
   for N*N chessboard */
func isSafe(x int,  y int, sol[N][N] int) bool {
	ret := ( x >= 0 && x < N && y >= 0 &&
          y < N && sol[x][y] == -1)
    return ret
}

/* A utility function to print solution matrix sol[N][N] */
func printSolution( sol[N][N] int) {
    for x := 0; x < N; x++ { 
         for y := 0; y < N; y++ {
            fmt.Printf(" %2d ", sol[x][y])
        }
        fmt.Printf("\n")
    }
}


/* This function solves the Knight Tour problem using
   Backtracking.  This function mainly uses solveKTUtil()
   to solve the problem. It returns false if no complete
   tour is possible, otherwise return true and prints the
   tour.
   Please note that there may be more than one solutions,
   this function prints one of the feasible solutions.  */
func solveKT() bool {

	//var xMove[8] int
	//var yMove[8] int

    /* Initialization of solution matrix */
    for x := 0; x < N; x++ {
        for y := 0; y < N; y++ {
            sol[x][y] = -1
        }
    }

    /* xMove[] and yMove[] define next move of Knight.
       xMove[] is for next value of x coordinate
       yMove[] is for next value of y coordinate */
       //s := []int{5, 2, 6, 3, 1, 4}
    xMove := []int{  2, 1, -1, -2, -2, -1,  1,  2 }
    yMove := []int{  1, 2,  2,  1, -1, -2, -2, -1 }

 
    // Since the Knight is initially at the first block
    sol[0][0] = 0;
   	calls = 0	
   /* Start from 0,0 and explore all tours using
       solveKTUtil() */
    if (solveKTUtil(0, 0, 1, xMove, yMove) == false)   {
        fmt.Printf("Solution does not exist\n");
        return false
    }    else {
    	fmt.Printf("solveKTUtil calls %d\n", calls)
        printSolution(sol)
    }


    return true
}

/* A recursive utility function to solve Knight Tour
   problem */

func solveKTUtil(x int, y int, movei int ,xMove[] int, yMove[] int) bool {
   var  k, next_x, next_y int

   //fmt.Printf("solveKTUtil move %d calls %d\n",movei, calls)
   calls += 1
   if movei == N*N {
   		fmt.Printf("when 64\n")
       	return true
   }


 
   /* Try all next moves from the current coordinate x, y */
   for k = 0; k < 8; k++   {
       next_x = x + xMove[k]
       next_y = y + yMove[k]
       if isSafe(next_x, next_y, sol) {
         	sol[next_x][next_y] = movei
         	if solveKTUtil(next_x, next_y, movei+1, xMove, yMove) == true {
             	return true
       		} else {
       			sol[next_x][next_y] = -1// backtracking
       		}
       }
   }
 
   return false
}

func mutate(j[]int) {
	j[2] = 333
}

func mutate2(j[][]int) {
	j[2][2] = 333
}

func xperiment() {

  	var p[] int
  	var s2d[][]int

  	p = []int{1,2,3,4,5}
  	mutate(p)
  	for i:=0; i < 5; i++ {
  		fmt.Printf("p %d\n",p[i])
	}

	s2d = [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	mutate2(s2d)
    for x := 0; x < 3; x++ { 
         for y := 0; y < 3; y++ {
            fmt.Printf(" %2d ", s2d[x][y])
        }
        fmt.Printf("\n")
    }


}
 
func main() {
  	fmt.Printf("Knight Tour Puzzle!\n")
  	xperiment()

	//solveKT()
}

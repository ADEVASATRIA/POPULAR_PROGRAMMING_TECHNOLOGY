//ADEVA SATRIA ARIF WIBAWA
//LA20
//2502012464
//BINUS UNIVERSITY MALANG
package main

import "fmt"

// fungsi untuk mengisi nilai array
func inputArray(size int) []string {
    array := make([]string, size)
    for i := 0; i < size; i++ {
        fmt.Printf("Masukkan nilai array[%d]: ", i)
        fmt.Scanln(&array[i])
    }
    return array
}

// fungsi untuk membandingkan dua array
func compareArray(array1 []string, array2 []string) bool {
    if len(array1) != len(array2) {
        return false
    }
    for i := 0; i < len(array1); i++ {
        if array1[i] != array2[i] {
            return false
        }
    }
    return true
}

// fungsi untuk mencetak hasil perbandingan
func printComparison(array1 []string, array2 []string) {
    if compareArray(array1, array2) {
        fmt.Println("Kedua array memiliki isi yang sama!")
    } else {
        fmt.Println("Kedua array memiliki isi yang berbeda:")
        for i := 0; i < len(array1); i++ {
            if array1[i] != array2[i] {
                fmt.Printf("- Index ke-%d berbeda: %s != %s\n", i, array1[i], array2[i])
            }
        }
    }
}

func main() {
    var size int
    fmt.Print("Masukkan ukuran array: ")
    fmt.Scanln(&size)

    fmt.Println("Isi array 1:")
    array1 := inputArray(size)

    fmt.Println("Isi array 2:")
    array2 := inputArray(size)

    printComparison(array1, array2)
}


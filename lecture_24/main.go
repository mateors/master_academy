package main

import (
	"fmt"
	"os"
)

func main() {

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(dir)

	// isErr := createFile("master_academy2.txt", "hello bangladesh")
	// fmt.Println(isErr)

	// fi, err := os.Stat("master_academy2.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())

	//how to make a folder
	// err := os.Mkdir("master_academy", 0777)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//base := filepath.Base(dir)
	//relativePath := filepath.Join("master_academy")
	//absolutePath, _ := filepath.Abs("master_academy")

	//E:\GOLANG\src\master_academy\lecture_24\master_academy
	// newPath := filepath.Join(absolutePath, "..", "..", "lecture_25")
	// fmt.Println(base)
	// //fmt.Println(relativePath)
	// fmt.Println(absolutePath)
	// fmt.Println(newPath)
	// os.Mkdir(newPath, 777)

	//os.Mkdir(`D:\TEST`, 777)
	os.Rename(`D:\TEST`, `D:\TEST_01`)

}

func createFile(fileName, content string) bool {

	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer posf.Close()
	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)
	if err != nil {
		return false
	}
	return true
}

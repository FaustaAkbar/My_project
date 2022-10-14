package main
import "fmt"
func soal(a,b int)int{
  var hasil = a+b
  return hasil
}
func main(){
  var a, b, hasil int
  fmt.Println("Program untuk penjumlahan")
  fmt.Scanln(&a,&b)
  hasil = soal(a,b)
  fmt.Println(hasil)
  
}

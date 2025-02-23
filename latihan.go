package main
import "fmt"

func main(){
    var namaProduk, Status string
    var harga, stok, diskon, hargaAsli float64

    fmt.Scan(&namaProduk, &harga, &stok)
    for namaProduk != "#"{
        hargaAsli = harga
        if harga >= 1000000{
            diskon = 0.9
        }else if harga < 1000000 && harga >= 500000{
            diskon = 0.95
        }else{
            diskon = 1
        }
        harga = harga*diskon
        if stok > 0{
            Status = "Produk tersedia"
        }else{
            Status = "Produk habis"
        }
        fmt.Println("Detail Produk")
        fmt.Println("Nama :", namaProduk)
        fmt.Printf("Harga asli : %.0f\n", hargaAsli)
        fmt.Printf("Diskon : %.0f%%\n", (1-diskon)*100)
        fmt.Printf("Harga setelah diskon : %.0f\n", harga)
        fmt.Printf("Stok : %.0f\n", stok)
        fmt.Println("Status :", Status)
        fmt.Println()

        fmt.Scan(&namaProduk, &harga, &stok)
    }
}
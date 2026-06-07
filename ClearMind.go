package main
import "fmt"

func main(){
	//Sebagai main menu. Terdapat dua pilihan, Porduktivitas dan Kesehatan Mental
	fmt.Printf("===== ClearMind =====\n")
	fmt.Println("[1] Menu Produktivitas")
	fmt.Println("[2] Menu Kesehatan Mental")
	fmt.Println("[0] Keluar Progam")
	fmt.Println("Pilih 1 atau 2: ")
	
	var pilihan string
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case "1":
			fmt.Print("On Progress")
	case "2":
			MenuKesehatanMental()
	case "0":
			fmt.Print("Terima Kasih sudah menggunakan ClearMind. Semoga hari Anda menyenangkan.")
	}
}

func MenuKesehatanMental(){
	fmt.Printf("=== CEK KESEHATAN MENTAL ===\n")
	fmt.Println("[1] Cek Koin Mental dan Stress Meter")
	fmt.Println("[2] Isi Jurnal Suasana Hati")
	fmt.Println("[3] Lihat Stastistik Mingguan")
	fmt.Println("[0] Kembali ke Menu Utama")
	
	var pilihan string
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case "1":
			fmt.Print("On Progress")
	case "2":
			fmt.Print("On Progress")
	case "3":
			fmt.Print("On Progress")
	case "0":
			main()
	}
}

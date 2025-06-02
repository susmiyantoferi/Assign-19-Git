Penjelasan: 
- func reverse(str string) string{} : membuat fungsi yang menerima parameter bertipe data string dan akan mengembalikan data dengan tipe data string.
- var hasil string : membuat variable hasil dengan tipe data string digunakan untuk menampung hasil reverse kata nantinya.
- for i := len(str)-1 : mendefined variable i dengan panjang karakter parameter yang dikurangi 1.
- i >= 0 : memberikan kondisi bahwa i harus lebih besar sama dengan 0.
- i-- : akan mengurangi 1 nilai variable i setiap perulangan dijalankan.
- hasil += string(str[i]) : disini setiap perulangan dijalankan karakter yang ada di parameter dengan index ke i akan ditambahkan ke variable hasil. disini juga dilakukan konversi dari tipe data bawaan int ke string.
- return hasil : kembalikan semua data yang ada dengan tipe data string

Contoh Penggunaan fungsi: 
- kata := "feri susmiyanto" : variable kata dengan isi digunakan sebagai parameter nantinya.
- fmt.Println(reverse(kata)) : menampilkan kata ke console dengan memanggil fungsi reverse beserta variable kata sebgai parameter.

Cara menjalankan program : 
- ketik go run di terminal diikuti dengan nama file 
- contoh go run reverse.go
 

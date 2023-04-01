// BEFORE
// package main

// type kendaraan struct {
// 	totalroda       int
// 	kecepatanperjam int
// }

// type mobil struct {
// 	kendaraan
// }

// func (m *mobil) berjalan() {
// 	m.tambahkecepatan(10)
// }

// func (m *mobil) tambahkecepatan(kecepatanbaru int) {
// 	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
// }

// func main() {
// 	mobilcepat := mobil{}
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()

// 	mobillamban := mobil{}
// 	mobillamban.berjalan()
// }

// ----------------------------------------------------------------------------------

// AFTER: BEFORE + COMMENT
// package main

// // 1: nama struct sebaiknya diawali huruf kapital. karena struct sebagai sebuah tanda pengenal yang mewakili object didalamnya. contoh: Kendaraan
// type Kendaraan struct {
// 	// 2: nama variabel lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: totalRoda
// 	totalRoda int
// 	// 3: nama variabel lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: kecepatanPerJam
// 	kecepatanPerJam int
// }

// // 4: nama struct sebaiknya diawali huruf kapital. karena struct sebagai sebuah tanda pengenal yang mewakili object didalamnya. contoh: Mobil
// type Mobil struct {
// 	Kendaraan
// }

// func (m *Mobil) berjalan() {
// 	// 5: memasukan nilai sebagai parameter sebaiknya dibuatkan variabel tersendiri agar memudahkan untuk memahami
// 	kilometerPerJam := 10
// 	m.tambahKecepatan(kilometerPerJam)
// }

// // 6: nama function lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: tambahKecepatan
// // 7: nama parameter lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: kecepatanBaru
// func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
// 	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
// }

// func main() {
// 	// 8: nama variabel lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: mobilCepat
// 	mobilCepat := Mobil{}
// 	// 9: Dont Repeat Your Self: artinya jangan melakukan hal yang sama berkali kali (copy and paste). lebih baik menggunakan looping
// 	for i := 0; i < 3; i++ {
// 		mobilCepat.berjalan()
// 	}

// 	// 10: nama variabel lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: mobilLamban
// 	mobilLamban := Mobil{}
// 	mobilLamban.berjalan()
// }

// // TOTAL KESALAHAN PENULISAN DALAM CODE INI SEJUMLAH: 10

// ---------------------------------------------------------------------------------------
// AFTER: FINAL RESULT

package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan() {
	// jika dijalankan maka kecepatan mobil akan bertambah sebesar nilai dari variabel kilometerPerJam
	kilometerPerJam := 10
	m.tambahKecepatan(kilometerPerJam)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	// jika dijalankan kecepatam mobil saat ini akan ditambahkan dengan kecepatan baru
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	for i := 0; i < 3; i++ {
		mobilCepat.berjalan()
	}

	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}

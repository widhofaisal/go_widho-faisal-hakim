package main

// 1: nama struct sebaiknya diawali huruf kapital. karena struct sebagai sebuah tanda pengenal yang mewakili object didalamnya. contoh: User
type user struct {
	id       int
	username int	// 2: tipe data untuk username seharusnya adalah string
	password int	// 3: tipe data untuk password seharusnya adalah string
}

// 4: nama function lebih dari satu kata sebaiknya menggunakan camel case. karena yang seperti ini sulit untuk dikenali. contoh: userService
type userservice struct {
	// 5: nama array harus jelas mendefinisikan kegunaannya, bila seperti ini menyebabkan miss understanding dan sulit dipahami
	t []user
}

// 6: nama function lebih dari satu kata sebaiknya menggunakan camel case. karena lebih sulit untuk dikenali. contoh: getAllUser
func (u userservice) getallusers() []user {
	return u.t
}


// 7: nama function lebih dari satu kata sebaiknya menggunakan camel case. karena lebih sulit untuk dikenali. contoh: getUserById
func (u userservice) getuserbyid(id int) user {
	// 8: value yang didapatkan dari looping harus diberi nama yang jelas agar mudah di pahami. contoh: eachUser
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

// TOTAL KESALAHAN PENULISAN DALAM CODE INI SEJUMLAH: 8
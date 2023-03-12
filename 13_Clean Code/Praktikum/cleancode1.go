package main

// Menurut saya ada beberapa kekurangan pada kode tersebut: 9
// Tidak ada komentar untuk menjelaskan masing-masing baris kode.
// tipe struct user harus diawali dengan huruf kapital di huruf pertama.
type user struct {
	// Penamaan variable/property harus menggunakan huruf kapital di awal.
	id       int
	username int
	password int
}

// Penamaan struct harus diawali dengan huruf kapital di huruf pertama.
type userservice struct {
	// Penamaan variable harus jelas dan deskriptif.
	t []user
}

// Tidak menggunakan pointer pada parameter dan awalan method harus huruf kapital.
func (u userservice) getallusers() []user {
	return u.t
}

// Tidak menggunakan pointer pada parameter dan awalan method harus huruf kapital.
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

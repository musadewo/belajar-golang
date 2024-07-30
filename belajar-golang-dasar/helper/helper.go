package helper

// version adalah variabel yang menyimpan versi aplikasi
var version = "1.0.0" // bisa dipanggil asal masih 1 package

// Application adalah variabel yang menyimpan nama aplikasi
var Application = "golang" // bisa dipanggil dari luar atau dalam package

// SayHello adalah fungsi yang menerima sebuah string 'name' dan mengembalikan
// sebuah string yang merupakan sapaan "Hello" diikuti dengan nama yang diberikan.
func SayHello(name string) string {
    // Mengembalikan string sapaan "Hello" diikuti dengan nama yang diberikan
    return "Hello " + name
}
// Package helper berisi fungsi-fungsi pembantu dan pengujiannya
// Package ini dinamakan "helper" untuk mengelompokkan fungsi-fungsi utilitas
package helper

// Import package yang diperlukan untuk pengujian
import (
	// fmt digunakan untuk mencetak pesan ke konsol
	// Dalam konteks pengujian, ini digunakan untuk mencetak pesan keberhasilan test
	"fmt"

	// runtime digunakan untuk mendapatkan informasi tentang lingkungan runtime Go
	// Dalam kode ini, digunakan untuk memeriksa sistem operasi dalam TestSkip
	"runtime"

	// testing adalah package standar Go untuk menulis dan menjalankan unit test
	// Menyediakan fungsi dan metode untuk melakukan assertions dan melaporkan hasil test
	"testing"

	// github.com/stretchr/testify/assert digunakan untuk membuat assertions yang lebih ekspresif
	// Assert memungkinkan test untuk berlanjut meskipun assertion gagal
	"github.com/stretchr/testify/assert"

	// github.com/stretchr/testify/require juga digunakan untuk assertions
	// Require akan menghentikan eksekusi test jika assertion gagal
	"github.com/stretchr/testify/require"
)

// TestTableHelloWorld adalah fungsi pengujian yang menggunakan teknik table-driven tests
// Ini memungkinkan pengujian beberapa kasus uji dalam satu fungsi
func TestTableHelloWorld(t *testing.T) {
    // Mendefinisikan slice dari struct yang berisi kasus uji
    tests := []struct {
        name     string  // Nama kasus uji
        request  string  // Input untuk fungsi HelloWorld
        expected string  // Output yang diharapkan
    }{
        // Mendefinisikan beberapa kasus uji
        {
            name:     "Dewo",
            request:  "Dewo",
            expected: "Hello, Dewo",
        },
        {
            name:     "Wo",
            request:  "Wo",
            expected: "Hello, Wo",
        },
        {
            name:     "De",
            request:  "De",
            expected: "Hello, De",
        },
    }

    // Melakukan iterasi melalui setiap kasus uji
    for _, test := range tests {
        // Menjalankan subtest untuk setiap kasus
        t.Run(test.name, func(t *testing.T) {
            result := HelloWorld(test.request)
            // Menggunakan require.Equal untuk memastikan hasil sesuai dengan yang diharapkan
            require.Equal(t, test.expected, result, test.expected)
        })
    }
}

// TestSubTest mendemonstrasikan penggunaan subtests dalam Go
func TestSubTest(t *testing.T) {
    // Menjalankan beberapa subtest
    t.Run("Test 1", func(t *testing.T) {
        result := HelloWorld("Dewo")
        require.Equal(t, "Hello, Dewo", result, "Result must be Hello, Dewo")
    }) 
    t.Run("Test 2", func(t *testing.T) {
        result := HelloWorld("Wo")
        require.Equal(t, "Hello, Wo", result, "Result must be Hello, Wo")
    }) 
    t.Run("Test 3", func(t *testing.T) {
        result := HelloWorld("De")
        require.Equal(t, "Hello, De", result, "Result must be Hello, De")
    })
}

// TestMain adalah fungsi khusus yang dijalankan sebelum dan sesudah semua test dalam package
func TestMain(m *testing.M) {
    fmt.Println("Sebelum unit test")
    m.Run() // Menjalankan semua test
    fmt.Println("Sesudah unit test")
}

// TestSkip mendemonstrasikan cara melewati test berdasarkan kondisi tertentu
func TestSkip(t *testing.T) {
    if runtime.GOOS == "darwin" {
        t.Skip("Can't run on Mac OS") // Melewati test jika berjalan di Mac OS
    }
    fmt.Println("TestSkip passed")
}

// TestHelloWorldRequire mendemonstrasikan penggunaan require untuk assertion
func TestHelloWorldRequire(t *testing.T) {
    result := HelloWorld("Dewo")
    require.Equal(t, "Hello, Dewo", result, "Result must be Hello, Dewo")
    fmt.Println("TestHelloWorldRequire passed")
}

// TestHelloWorldAssert mendemonstrasikan penggunaan assert untuk assertion
func TestHelloWorldAssert(t *testing.T) {
    result := HelloWorld("Dewo")
    assert.Equal(t, "Hello, Dewo", result, "Result must be Hello, Dewo")
    fmt.Println("TestHelloWorldAssert passed")
}

// TestHelloWorld adalah fungsi pengujian dasar untuk HelloWorld
func TestHelloWorld(t *testing.T) {
    result := HelloWorld("Dewo")
    if result != "Hello, Dewo" {
        panic("Harusnya Hello, Dewo") // Tidak disarankan menggunakan panic dalam test
        t.Fail()
        t.Error("Result must be Hello, Dewo")
    }
    fmt.Println("TestHelloWorld passed")
}

// TestHelloWorldWo adalah fungsi pengujian lain untuk HelloWorld
func TestHelloWorldWo(t *testing.T) {
    result := HelloWorld("Wo")
    if result != "Hello, Wo" {
        // panic("Harusnya Hello, Wo") // Tidak disarankan menggunakan panic dalam test
        // t.FailNow()
        t.Fatal("Result must be Hello, Wo")
    }
    fmt.Println("TestHelloWorldWo passed")
}

// TestHelloWorldDe adalah fungsi pengujian tambahan untuk HelloWorld
func TestHelloWorldDe(t *testing.T) {
    result := HelloWorld("De")
    if result != "Hello, De" {
        t.Fatal("Result must be Hello, De")
    }
    fmt.Println("TestHelloWorldDe passed")
}
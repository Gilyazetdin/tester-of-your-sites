package main
import 
(
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func Gettest(site_url string) {
	for {
		_, err := http.Get(site_url) 
		if err != nil { 
			fmt.Println("Problems with sending: ", err) 
		} else {
			fmt.Println("Success.")				
		}
	}
}
func Posttest(site_url string, keys url.Values) {
	for {
		_, err := http.PostForm(site_url, keys) 
		if err != nil { 
			fmt.Println("Problems with sending: ", err) 
		} else {
			fmt.Println("Success.")				
		}
	}
}

func main() {
	fmt.Print("Enter the name of the site: ")
	var site_url string
	fmt.Scanln(&site_url)

	fmt.Print("Enter number of threads: ")
	var times int
	fmt.Scanln(&times)

	fmt.Print("How long will the program work? (seconds): ")
	var seconds time.Duration
	fmt.Scanln(&seconds)

	var answer int
	fmt.Print("GET or POST (0/1): ")
	fmt.Scanln(&answer)

	if answer == 0 {
		for i := 0; i < 100; i++ {
			go Gettest(site_url)
		}
	} else if answer == 1 {
		fmt.Print("Enter number of keys for POST request: ")
		var numb_of_keys int
		fmt.Scanln(&numb_of_keys)

		keys := url.Values{}

		for i:=0; i < numb_of_keys; i++ {
			fmt.Print("Enter name of key: ")
			var name string
			fmt.Scanln(&name)

			fmt.Print("Enter value of key: ")
			var value string
			fmt.Scanln(&value)

			keys.Add(name, value)
		}

		for i := 0; i < 100; i++ {
			go Posttest(site_url, keys)
		}
	}
	
	time.Sleep(seconds * time.Second)
}
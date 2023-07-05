//HTTP request

package main
import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"

type Post struct{
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func main(){
	//URL to send the GET request 
	url := "https://jsonplaceholder.typicode.com/posts/1"

	//Send the GET request
	response, err := http.Get(url)
	if err != nil{
		fmt.Printf("HTTP GET request failed with error %s\n",err)
		return 
	}

//Ensure the response body is closed after we are done with it
defer response.Body.Close()

//Read the response body
body, err := ioutil.ReadAll(response.Body)
if err != nil{
	fmt.Printf("Failed to read the response body %s\n",err)
	return
}

//Parse the JSON response into a post struct
var post Post
err = json.Unmarshal(body, &post)
if err != nil{
	fmt.Printf("Failed to parse JSON: %s\n",err)
	return
}

//Print the parsed post
fmt.Printf("UserID : %d\n",post.UserID)
fmt.Printf("ID : %d\n",post.ID)
fmt.Printf("Title : %s\n",post.Title)
fmt.Printf("Body : %s\n",post.Body)
}

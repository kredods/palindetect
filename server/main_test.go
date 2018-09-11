package main
import(
	"bytes"
	"testing"
	"net/http/httptest"
	"net/http"
	"log"
	"github.com/go-chi/chi"
	"github.com/kredods/palindetect/server/models"
	"encoding/json"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
	  return
	  }
	  if len(message) == 0 {
		  message = fmt.Sprintf("%v != %v", a, b)
	  }
	  t.Fatal(message)
  }

func TestRoutes(t *testing.T){

	router := Routes()

	server:= httptest.NewServer(router)
	defer server.Close()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) 
		return nil
	}
	chi.Walk(router, walkFunc); 

	
	t.Run("Test Create Message", func(t *testing.T) { 
		newMessage := models.Message{Body: "testCreate"}
		buffer,_:= json.Marshal(newMessage)

		res, err := http.Post( server.URL + "/v1/messages","application/json", bytes.NewBuffer(buffer))
		if err != nil {
			log.Fatal(err)
		}

		assertEqual(t ,res.StatusCode, http.StatusCreated, "Invalid status code")
		var responseMessage models.Message
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&responseMessage)
		assertEqual(t ,responseMessage.Body, "testCreate", "Invalid message")
	 })

	 t.Run("Test GetMessage by Id", func (t *testing.T){
		newMessage := models.Message{Body: "testGetById"}
		buffer,_:= json.Marshal(newMessage)

		res, err := http.Post( server.URL + "/v1/messages","application/json", bytes.NewBuffer(buffer))
		if err != nil {
			log.Fatal(err)
		}

		var responseMessage models.Message
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&responseMessage)

		res, err = http.Post( server.URL + "/v1/messages","application/json", bytes.NewBuffer(buffer))
		if err != nil {
			log.Fatal(err)
		}

		res, err = http.Get(server.URL + "/v1/messages/" + responseMessage.Id.Hex())
		if err != nil {
			log.Fatal(err)
		}

		decoder = json.NewDecoder(res.Body)
		decoder.Decode(&responseMessage)

		assertEqual(t ,responseMessage.Body, "testGetById", "Invalid message")
	 })

	 t.Run("Test Delete Message", func (t *testing.T){
		newMessage := models.Message{Body: "testDelete"}
		buffer,_:= json.Marshal(newMessage)

		res, err := http.Post( server.URL + "/v1/messages","application/json", bytes.NewBuffer(buffer))
		if err != nil {
			log.Fatal(err)
		}

		var responseMessage models.Message
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&responseMessage)

		client := &http.Client{}
    	req, err := http.NewRequest("DELETE", server.URL + "/v1/messages/" + responseMessage.Id.Hex(), nil)
   		 if err != nil {
   		   fmt.Println(err)
   		     return
  		  }

	  	res, err = client.Do(req)
    		if err != nil {
        		fmt.Println(err)
        	return
  		  }
		  defer res.Body.Close()
		  
		  assertEqual(t, res.StatusCode, 200, "Unsuccessful delete")
		
		  res, err = http.Get(server.URL + "/v1/messages/" + responseMessage.Id.Hex())
		  if err != nil {
			  log.Fatal(err)
		  }
		  assertEqual(t, res.StatusCode, 404, "failed to  delete by id")
	 })


}

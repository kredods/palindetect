package messages

import (
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/kredods/palindetect/server/db"
	"github.com/kredods/palindetect/server/models"
	"github.com/kredods/palindetect/server/controllers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", getMessage)
	router.Delete("/{id}", deleteMessage)
	router.Post("/", createMessage)
	router.Get("/", getAllMessages)
	router.Put("/{id}", updateMessage)
	return router
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dbId := bson.ObjectIdHex(id)
	
	messageCollection, session, err := getMessageCollection()
	defer session.Close()
	if err != nil {
		http.Error(w, "Unable to Connect to Database", 500 )
		return
	}

	var message models.Message
	err = messageCollection.Find(bson.M{"id":dbId}).One(&message)

	if err != nil {
		http.Error(w, "Cannot find message", 404 )
		return
	}

	render.JSON(w, r, message) // A chi router helper for serializing and returning json
}

func updateMessage(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var message models.Message
    err := decoder.Decode(&message)
    if err != nil {
		http.Error(w, "Unable to Decode Json", 400 )
		return
    }
	id := chi.URLParam(r, "id")
	dbId := bson.ObjectIdHex(id)
	message.IsPalindrome = controllers.IsPalindrome(message.Body)
	messageCollection, session, err := getMessageCollection()
	defer session.Close()
	if err != nil {
		http.Error(w, "Unable to Connect to Database", 500 )
		return
	}

	err = messageCollection.Update(bson.M{"id": dbId}, message)
	if err != nil {
		http.Error(w, "Could not Edit message", 400)
		return
	}

	response := make(map[string]string)
	response["message"] = "Updated Message successfully"
	render.JSON(w, r, response) 
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dbId := bson.ObjectIdHex(id)
	
	messageCollection, session, err := getMessageCollection()
	defer session.Close()
	if err != nil {
		http.Error(w, "Unable to Connect to Database", 500 )
		return
	}

	err = messageCollection.Remove(bson.M{"id": dbId})
	if err != nil {
		http.Error(w, "Unable to delete message", 404 )
		return
	}

	response := make(map[string]string)
	response["message"] = "Deleted Message successfully"
	render.JSON(w, r, response) 
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newMessage models.Message
    err := decoder.Decode(&newMessage)
    if err != nil {
		http.Error(w, "Unable to Decode Json", 400 )
		return
	}
	
	newMessage.IsPalindrome = controllers.IsPalindrome(newMessage.Body)
	messageCollection, session, err := getMessageCollection()
	defer session.Close()
	if err != nil {
		http.Error(w, "Unable to Connect to Database", 500 )
		return
	}
	newMessage.Id = bson.NewObjectId()
	err = messageCollection.Insert(newMessage);
	if err != nil {
		http.Error(w, "Failed to insert message", 500 )
		return
	}
	response := make(map[string]string)
	response["message"] = "Created Message Successfully"
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, newMessage) 
}

func getAllMessages(w http.ResponseWriter, r *http.Request) {

	messageCollection, session, err := getMessageCollection()
	defer session.Close()
	if err != nil {
		return
	}
	var Messages []models.Message
	err = messageCollection.Find(bson.M{}).All(&Messages)
	if err != nil {
		return
	}
	render.JSON(w, r, Messages) 
}

func getMessageCollection() (messageCollection *mgo.Collection, session *mgo.Session, err error) {
	session, err = db.GetSession()
	database := session.DB(db.GetDBName())
	if err != nil {
		return
	}
	messageCollection = database.C("messages")
	return
}
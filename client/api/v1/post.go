package v1

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"main/utils"
	"net/http"
	"net/url"
	"strconv"
)

func AddPoolMember(w http.ResponseWriter, r *http.Request) {
	target_id := r.PathValue("id")
	target_id = url.PathEscape(target_id)

	user_id, err := utils.GetUserId()
	if err != nil {
		log.Println("[X] Error AddPoolMember utils.GetUserId: ", err)
		http.Error(w, "Failed to invite a new member", http.StatusInternalServerError)
		return
	}

	if strconv.Itoa(user_id) == target_id {
		log.Printf("[X] Error usr%d is adding himself 🥶🥶\n", user_id)
		http.Error(w, "Can't add yourself !", http.StatusBadRequest)
		return
	}

	body := `{"user_id":` + strconv.Itoa(user_id) + `}`
	response, err := http.Post(
		"http://localhost:3032/v1/pool/add/"+target_id, //hardcoded for now
		"application/json",
		bytes.NewBufferString(body))
	if err != nil {
		http.Error(w, "Failed to invite a new member", http.StatusInternalServerError)
		log.Println("[X] Error AddPoolMember http.Post: ", err)
		return
	}
	defer response.Body.Close()

	poolResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", poolResponse)
}

func DeletePoolMember(w http.ResponseWriter, r *http.Request) {
	target_id := r.PathValue("id")
	target_id = url.PathEscape(target_id)

	user_id, err := utils.GetUserId()
	if err != nil {
		http.Error(w, "Failed to delete pool member", http.StatusInternalServerError)
		log.Println("[X] Error DeletePoolMember utils.GetUserId: ", err)
		return
	}

	if strconv.Itoa(user_id) == target_id {
		log.Printf("[X] Error usr%d is deleting himself 🗣🗣\n", user_id)
		http.Error(w, "Can't delete yourself (yet)", http.StatusBadRequest)
		return
	}

	body := `{"user_id":` + strconv.Itoa(user_id) + `}`
	response, err := http.Post(
		"http://localhost:3032/v1/pool/delete/"+target_id, //hardcoded for now
		"application/json",
		bytes.NewBufferString(body))
	if err != nil {
		http.Error(w, "Failed to delete pool member", http.StatusInternalServerError)
		log.Println("[X] Error DeletePoolMember http.Post: ", err)
		return
	}
	defer response.Body.Close()

	poolResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", poolResponse)
}
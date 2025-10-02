package v1

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"main/utils"
	"net/http"
	"strconv"
)

func ClearUser(w http.ResponseWriter, r *http.Request) {
	user_id, err := utils.GetUserId()
	if err != nil {
		http.Error(w, "Failed to delete pool member", http.StatusInternalServerError)
		log.Println("[X] Error DeletePoolMember utils.GetUserId: ", err)
		return
	}

	body := `{"user_id":` + strconv.Itoa(user_id) + `}`
	request, err := http.NewRequest(
		http.MethodPut,
		"http://localhost:3032/v1/users/me"+strconv.Itoa(user_id), //hardcoded for now
		bytes.NewBufferString(body))
	if err != nil {
		http.Error(w, "Failed to clear user", http.StatusInternalServerError)
		log.Println("[X] Error ClearUser http.MethodPut: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", response)
}

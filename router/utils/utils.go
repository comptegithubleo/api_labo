package utils

import (
	"errors"
	"io"
	"log"
	"os/exec"
)

// executes an os commands, returns stdout & stderr
// test the function with stdin stdout : "sh", "-c", "echo stdout; echo 1>&2 stderr"
func Exec(name string, args ...string) (string, string) {

	cmd := exec.Command(name, args...)
	stdoutBytes, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrBytes, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		return "", err.Error()
	}

	stdout_, _ := io.ReadAll(stdoutBytes)
	stderr_, _ := io.ReadAll(stderrBytes)

	// we do not error check, we already read stderr and let callee deal with it (throw error or not)
	// still needed because we wait for command to stop and all stdout stderr pipe to be read
	cmd.Wait()

	return string(stdout_), string(stderr_)
}

func AddConnection(user_a, user_b int) error {
	users, err := GetJSONUsers()
	if err != nil {
		return err
	}

	count := 0
	for i := range users {
		if users[i].ID == user_a {
			users[i].Connections = append(users[i].Connections, user_b)
			count++
		}
		if users[i].ID == user_b {
			users[i].Connections = append(users[i].Connections, user_a)
			count++
		}
		if count == 2 {
			break
		}
	}

	if count != 2 {
		return errors.New("Failed to append connections bilaterally. Check or reset user connections")
	}

	err = WriteJSONUsers(users)
	if err != nil {
		return err
	}

	return nil
}

func AreConnected(user_a, user_b int) (bool, error) {
	users, err := GetJSONUsers()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.ID == user_a || user.ID == user_b {
			for _, connection := range user.Connections {
				if connection == user_a || connection == user_b {
					return true, nil
				}
			}
			break
		}
	}

	return false, nil
}

func AddInvite(user_a, user_b int) error {
	// if already connected, return
	connected, err := AreConnected(user_a, user_b)
	if err != nil {
		return err
	}
	if connected {
		return errors.New("Already connected 👑🫑")
	}

	invites, err := GetJSONInvites()
	if err != nil {
		return err
	}
	for i := 0; i < len(invites); i++ {
		if invites[i].From == user_a && invites[i].To == user_b {
			return errors.New("Already invited 🐡💀") // already exists do nothing
		}
		if invites[i].From == user_b && invites[i].To == user_a {
			// accept invite : remove invite && add user connections
			invites[i] = invites[len(invites)-1]
			invites = invites[:len(invites)-1]
			err = WriteJSONInvites(invites)
			if err != nil {
				return err
			}

			AddConnection(user_a, user_b)
			log.Printf("usr%d -> usr%d : Invite accepted, connection created 👺🚿", user_a, user_b)
			return nil
		}
	}

	log.Printf("usr%d -> usr%d : Invite created 🎀✨", user_a, user_b)
	// invite was not found, create one
	invites = append(invites, PendingInvite{
		From: user_a,
		To:   user_b,
	})

	err = WriteJSONInvites(invites)
	if err != nil {
		return err
	}

	return nil
}

func RemoveInvite(invites PendingInvites, user_a, user_b int) {
	for i := 0; i < len(invites); i++ {
		if invites[i].From == user_a && invites[i].To == user_b ||
			invites[i].From == user_b && invites[i].To == user_a {
			//remove from invite, add connection to users.json
			// https://go.dev/wiki/SliceTricks#delete-without-preserving-order because faster
			invites[i] = invites[len(invites)-1]
			invites = invites[:len(invites)-1]
		}
	}
}

func RemoveConnection(user_a, user_b int) error {

	users, err := GetJSONUsers()
	if err != nil {
		return err
	}

	connected, err := AreConnected(user_a, user_b)
	if err != nil {
		return err
	}
	if !connected {
		return errors.New("Can't remove connection between two unconnected users")
	}

	count := 0
	for i := range users {
		if users[i].ID == user_a {
			for j := 0; j < len(users[i].Connections); j++ {
				if users[i].Connections[j] == user_b {
					users[i].Connections[j] = users[i].Connections[len(users[i].Connections)-1]
					users[i].Connections = users[i].Connections[:len(users[i].Connections)-1]
					count++
					break
				}
			}
		}
		if users[i].ID == user_b {
			for j := 0; j < len(users[i].Connections); j++ {
				if users[i].Connections[j] == user_a {
					users[i].Connections[j] = users[i].Connections[len(users[i].Connections)-1]
					users[i].Connections = users[i].Connections[:len(users[i].Connections)-1]
					count++
					break
				}
			}
		}
		if count == 2 {
			log.Printf("usr%d <-> usr%d : Connection removed ⛓️‍💥👁\n", user_a, user_b)
			break
		}
	}

	if count != 2 {
		return errors.New("Failed to remove connections bilaterally. Check or clean user connections")
	}

	err = WriteJSONUsers(users)
	if err != nil {
		return err
	}

	return nil
}

/* 	stdout, stderr := Exec("./scripts/getUsers.sh")
fmt.Println(stderr)

users := strings.Split(strings.TrimSpace(stdout), "\n")
fmt.Println(users) */

package digitalocean

// import (
// 	"fmt"
// 	"os/exec"
// 	"os/user"
// 	"path"
// 	"strings"
// 	"testing"
// )

// func TestKeyFingerprint(t *testing.T) {
// 	t.Logf("Testing getting a key via it's fingerprint\n")
//
// 	usr, err := user.Current()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// Assumes ssh key is id_rsa.pub""
// 	p := path.Join(usr.HomeDir, ".ssh", "id_rsa.pub")
//
// 	out, err := exec.Command("ssh-keygen", "-lf", p).Output()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	fp := strings.Split(string(out), " ")[1]
// 	fmt.Printf("Fingerprint %s\n", fp)
//
// 	token, err := digitalocean.EnvAuth()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cli := NewClient(token)
// 	key, err := cli.Get(fp)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("%+v\n", key)
// }

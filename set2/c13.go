package main

import (
	"cryptopals/lib"
	"cryptopals/lib/key"
	"fmt"

	// "log"
	"strings"
)

var verbose = false

func main() {
	key := key.RandKey()
	oracle := func(email string) (ct []byte) {
		p := profile_for(email)
		if p != nil {
			ct = p.encrypt(key)
		}
		return
	}
	/*
		verify_admin := func(ct []byte) bool {
			p := decrypt(ct, key)
			return (p != nil) && p.role == "admin"
		}*/
	extract := func(ct []byte) (p *profile) { return decrypt(ct, key) }
	attacker(oracle, extract)
}

/*
profile_for("foo@bar.com")
... and it should produce:

{
  email: 'foo@bar.com',
  uid: 10,
  role: 'user'
}
... encoded as:

email=foo@bar.com&uid=10&role=user
*/

type profile struct {
	email string
	uid   string
	role  string
}

type pair struct{ key, value string }

func decode(cookie string) (p *profile) {
	p = new(profile)
	split := strings.Split(cookie, "&")
	if len(split) != 3 {
		if verbose {
			fmt.Println("invalid data split", cookie)
		}
		return nil
	}
	keys := []string{"email", "uid", "role"}
	for i, k := range split {
		splat := strings.Split(k, "=")
		if len(splat) != 2 {
			if verbose {
				fmt.Println("invalid data splat", cookie)
			}
			return nil
		}
		if splat[0] != keys[i] {
			if verbose {
				fmt.Println("invalid data key", cookie)
			}
			return nil
		}
		//switch in a for loop is a weeeeeee bit ugly but lets us avoid some copy-paste... toss-up?
		switch i {
		case 0:
			p.email = splat[1]
		case 1:
			p.uid = splat[1]
		case 2:
			p.role = splat[1]
		default:
			//uhh???
			fmt.Println("invalid data range", cookie)
			return nil
		}
	}
	return
}

func (p *profile) encode() string {
	return fmt.Sprintf("email=%s&uid=%s&role=%s", p.email, p.uid, p.role)
}
func (p *profile) String() string {
	return fmt.Sprintf("email: %s\nuid: %s\nrole: %s\n", p.email, p.uid, p.role)
}
func profile_for(email string) (p *profile) {
	p = new(profile)
	email = strings.Replace(email, "&", "ampersand", -1)
	p.email = strings.Replace(email, "=", "equal", -1)
	p.uid = "totallyrandom"
	p.role = "user"
	return
}

func (p *profile) encrypt(key []byte) []byte {
	pt := p.encode()
	padded := lib.Pkcs7pad([]byte(pt), 16)
	return lib.Encrypt_aes_ecb(padded, key)
}
func decrypt(ct []byte, key []byte) (p *profile) {
	pt := lib.Decrypt_aes_ecb(ct, key)
	return decode(string(pt))
}

//Using only the user input to profile_for() (as an oracle to generate "valid"
//ciphertexts) and the ciphertexts themselves, make a role=admin profile.
func attacker(oracle func(string) []byte, extract func([]byte) (p *profile)) {
	baselineCT := oracle("user@email.com")
	baselineProf := extract(baselineCT)
	//figure out where role field starts
	//how to automate?
	workingCT := make([]byte, len(baselineCT))
	//	i := len(baselineCT) - 1
	var i int
	fmt.Println(i, baselineProf.String())
	var b byte = 0
outer:
	for ; ; b++ {
		i = len(baselineCT) - 1
		//fmt.Println(b, i, "start")
		for ; i > 0; i-- {
			//fiddle with characters until p.role[0] changes
			copy(workingCT, lib.Pkcs7strip(baselineCT, 16))
			workingCT[i] = b //'0' //^(workingCT[i])
			w := extract(workingCT)
			if w == nil {
				if verbose {
					fmt.Println(i, "nil")
				}
				continue
			}
			if w.role == baselineProf.role {
				if verbose {
					fmt.Println(b, i, "role unchanged")
				}
				continue
			}
			if w.role[0] != baselineProf.role[0] {
				fmt.Println(i, "found role start at", i, b)
				break outer
			}
			fmt.Println(b, i, "role:", w.role)
		}
		if b == 255 {
			//overflow otherwise, since the for loop adds then evaluates
			break
		}
	}
	fmt.Println(i, b)
	//i is field start
	//here, the field is last.. TODO allow inserting into cyphertexts
	//strip padding
	//insert random values into cyphertext until p.role==admin
}

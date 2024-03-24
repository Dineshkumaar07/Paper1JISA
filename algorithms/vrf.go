package algorithms

import (
	"crypto/rand"
	"fmt"

	vrf "github.com/YahooArchive/coname/vrf"
)

func Vrf() {
	rnd := rand.Reader

	// Generate Keys
	pk, sk, err := vrf.GenerateKey(rnd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Generated Public Key:", pk)
	fmt.Println("Generated Secret Key:", sk)

	//Compute Method
	message := []byte("VRF Implementation")
	vrfOutput := vrf.Compute(message, sk)
	fmt.Println("Compute Output", vrfOutput)

	//Prove
	vrfOutput, proof := vrf.Prove(message, sk)
	fmt.Println("VRF Output:", vrfOutput)
	fmt.Println("Proof:", proof)

	//Verification
	isValid := vrf.Verify(pk, message, vrfOutput, proof)
	fmt.Println("Is Valid:", isValid)

}

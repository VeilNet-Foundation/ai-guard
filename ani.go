// ai_resistance/ani.go
package ai_resistance

import (
	"veilnet/core/tunnel"
	"veilnet/core/crypto"
	"time"
	"math/rand"
)

func StartAdversarialNoise(nodeIDs []string) {
	for {
		go func() {
			path := randomPath(nodeIDs, 3+rand.Intn(5))
			payload := generateRandomPayload()
			encrypted := crypto.EncryptOnion(payload, extractKeys(path))
			
			pkt := &packet.Packet{
				Content:   encrypted,
				NextHopID: path[0],
				IsLast:    false,
				SessionID: "decoy-" + randString(8),
			}
			
			tunnel.Forward(path[0], pkt)
		}()
		time.Sleep(time.Duration(rand.Intn(200)+50) * time.Millisecond)
	}
}
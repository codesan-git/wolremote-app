package wol

import (
	"fmt"

	"github.com/Ajnasz/wol"
)

func Wake(macAddr string, broadcastIP string) error {
	err := wol.SendPacket(macAddr, broadcastIP)
	if err != nil {
		return fmt.Errorf("failed to send WOL packet to %s: %w", macAddr, err)
	}
	return nil
}

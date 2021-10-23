package mc

import (
	"encoding/json"
	"strings"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	mc "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

type PingList struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
}

func handshake(ip string, conn *mc.Conn) error {
	// HACK: Get the server's IP
	ipNoPort := strings.Split(ip, ":")[0]

	return conn.WritePacket(pk.Marshal(
		0x00,
		pk.VarInt(756),
		pk.String(ipNoPort),
		pk.UnsignedShort(25565),
		pk.Byte(1),
	))
}

func Ping(ip string) (*PingList, error) {
	if !strings.Contains(ip, ":") {
		ip = ip + ":25565"
	}

	conn, err := mc.DialMC(ip)
	if err != nil {
		return nil, err
	}

	err = handshake(ip, conn)
	if err != nil {
		return nil, err
	}

	err = conn.WritePacket(pk.Marshal(
		packetid.PingStart,
	))

	if err != nil {
		return nil, err
	}

	var p pk.Packet
	if err := conn.ReadPacket(&p); err != nil {
		return nil, err
	}

	var str pk.String
	err = p.Scan(&str)

	if err != nil {
		return nil, err
	}

	var ret PingList
	err = json.Unmarshal([]byte(str), &ret)

	return &ret, err
}

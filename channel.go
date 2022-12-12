package main

import (
	"net"
)

type channel struct {
	name    string             //nombre del canal
	members map[net.Addr]*user //almacena la dirección de los miembros del canal
}

//transmision del archivo a los demas usuarios
func (c *channel) fileBroadcast(sender *user, fileName string) {
	if sender.channel != nil {
		for addr, m := range c.members {
			//se enviara el archivo a los miembros, exceptuando al remitente
			if sender.conn.RemoteAddr() != addr {
				m.msg(sender.name + " ha enviado un archivo")
				m.receiveFile(fileName)
			}
		}
	}

}

//transmision de mensaje a otros usuarios
func (c *channel) broadcast(sender *user, msg string) {
	if sender.channel != nil {
		for addr, m := range c.members {
			//se enviara el archivo a los miembros, exceptuando al remitente
			if sender.conn.RemoteAddr() != addr {
				m.msg(msg)
			}
		}
	}

}

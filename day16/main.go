package day16

import (
	"advent2021/helper"
	"math"
)

func asBits(val rune) []int {
	switch val {
	case '0':
		return []int{0, 0, 0, 0}
	case '1':
		return []int{0, 0, 0, 1}
	case '2':
		return []int{0, 0, 1, 0}
	case '3':
		return []int{0, 0, 1, 1}
	case '4':
		return []int{0, 1, 0, 0}
	case '5':
		return []int{0, 1, 0, 1}
	case '6':
		return []int{0, 1, 1, 0}
	case '7':
		return []int{0, 1, 1, 1}
	case '8':
		return []int{1, 0, 0, 0}
	case '9':
		return []int{1, 0, 0, 1}
	case 'A':
		return []int{1, 0, 1, 0}
	case 'B':
		return []int{1, 0, 1, 1}
	case 'C':
		return []int{1, 1, 0, 0}
	case 'D':
		return []int{1, 1, 0, 1}
	case 'E':
		return []int{1, 1, 1, 0}
	case 'F':
		return []int{1, 1, 1, 1}
	}
	return nil
}

func asInt(val []int) int {
	var n int
	for i, b := range val {
		n += b * int(math.Pow(2, float64(len(val)-i-1)))
	}
	return n
}

type Packet struct {
	Version    int
	Type       int
	Value      int
	SubPackets []Packet
}

func parsePacket(data []int, offset int) (Packet, int) {
	packet := Packet{
		Version: asInt(data[offset : offset+3]),
		Type:    asInt(data[offset+3 : offset+6]),
	}
	offset += 6

	if packet.Type == 4 {
		literal := []int{}
		for {
			continueBit := data[offset]
			literal = append(literal, data[offset+1:offset+5]...)
			offset += 5
			if continueBit == 0 {
				break
			}
		}
		packet.Value = asInt(literal)
	} else {
		lengthType := data[offset]
		if lengthType == 0 {
			packet.Value = asInt(data[offset+1 : offset+16])
			offset += 16
			startOffset := offset
			for {
				var subPacket Packet
				subPacket, offset = parsePacket(data, offset)
				packet.SubPackets = append(packet.SubPackets, subPacket)
				if offset-startOffset == int(packet.Value) {
					break
				}
			}
		} else {
			packet.Value = asInt(data[offset+1 : offset+12])
			offset += 12
			for i := 0; i < packet.Value; i++ {
				var subPacket Packet
				subPacket, offset = parsePacket(data, offset)
				packet.SubPackets = append(packet.SubPackets, subPacket)
			}
		}
	}
	return packet, offset
}

func sumVersions(packet Packet) int {
	sum := packet.Version
	for _, p := range packet.SubPackets {
		sum += sumVersions(p)
	}
	return sum
}

func evaluatePacket(packet Packet) int {
	switch packet.Type {
	case 0:
		sum := 0
		for _, p := range packet.SubPackets {
			sum += evaluatePacket(p)
		}
		return sum
	case 1:
		product := 1
		for _, p := range packet.SubPackets {
			product *= evaluatePacket(p)
		}
		return product
	case 2:
		min := -1
		for _, p := range packet.SubPackets {
			val := evaluatePacket(p)
			if min == -1 || val < min {
				min = val
			}
		}
		return min
	case 3:
		max := -1
		for _, p := range packet.SubPackets {
			val := evaluatePacket(p)
			if max == -1 || val > max {
				max = val
			}
		}
		return max
	case 4:
		return packet.Value
	case 5:
		if evaluatePacket(packet.SubPackets[0]) > evaluatePacket(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		if evaluatePacket(packet.SubPackets[0]) < evaluatePacket(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if evaluatePacket(packet.SubPackets[0]) == evaluatePacket(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func Part1() {
	input := helper.ReadInput("day16/input.txt")[0]
	data := []int{}
	for _, r := range input {
		data = append(data, asBits(r)...)
	}
	packet, _ := parsePacket(data, 0)

	println(sumVersions(packet))
}

func Part2() {
	input := helper.ReadInput("day16/input.txt")[0]
	data := []int{}
	for _, r := range input {
		data = append(data, asBits(r)...)
	}
	packet, _ := parsePacket(data, 0)

	println(evaluatePacket(packet))
}

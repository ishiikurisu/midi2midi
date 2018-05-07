package main

import (
    "github.com/tarm/serial"
    "time"
)

// plays a MIDI note. Doesn't check to see that cmd is greater than 127, or that
// data values are less than 127.
func noteOn(s *serial.Port, cmd, pitch, velocity byte) {
    s.Write([]byte {
        cmd,
        pitch,
        velocity,
    })
}

// Check [this](https://www.arduino.cc/en/Tutorial/Midi) for reference.
func main() {
    c := &serial.Config {
        Name: "COM4",
        Baud: 31250,
    }
    s, err := serial.OpenPort(c)
    if err != nil {
        
    }

    var note byte
    for true {
        for note = 0x1E; note < 0x5A; note++ {
            noteOn(s, 0x90, note, 0x45)
            time.Sleep(100 * time.Millisecond)
            noteOn(s, 0x90, note, 0x00)
            time.Sleep(100 * time.Millisecond)
        }
    }
}

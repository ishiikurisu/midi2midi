// plays a MIDI note. Doesn't check to see that cmd is greater than 127, or that
// data values are less than 127.
void noteOn(int cmd, int pitch, int velocity) {
    Serial.write(cmd);
    Serial.write(pitch);
    Serial.write(velocity);
}

// Check [this](https://www.arduino.cc/en/Tutorial/Midi) for reference.

void setup() {
    Serial.begin(115200);
}

void loop() {
    int note;

    for (note = 0x1E; note < 0x5A; note++) {
        noteOn(0x90, note, 0x45);
        delay(100);
        noteOn(0x90, note, 0x00);
        delay(100);
    }
}

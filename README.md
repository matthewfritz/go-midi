# Go MIDI

Go implementation of a library to represent MIDI data.

This began as a passion project to learn and implement a workable amount of the MIDI 1.0 specification from scratch so I could have complete control over certain specific digital instruments and DAW controls.

I had been trying to solve a very specific problem I was having with a hardware MIDI controller and its percussion pads too. The library then began growing into a general representation of the specification(s) from there.

## Table of Contents

   * [Installation](#installation)
   * [Roadmap](#roadmap)
      * [MIDI 1.0 Roadmap](#midi-10-roadmap)
      * [MIDI 2.0 Roadmap](#midi-20-roadmap)
   * [Resources](#resources)
      * [Official Specifications](#official-specifications)
      * [Books](#books)
      * [Assorted MIDI and Digital Music Information](#assorted-midi-and-digital-music-information)
         * [Assorted MIDI 1.0 Information](#assorted-midi-10-information)
         * [Digital Music Information](#digital-music-information)

## Installation

```
go get https://github.com/matthewfritz/go-midi
```

## Roadmap

### MIDI 1.0 Roadmap

#### Channel Voice Messages

   * ✅ [Note-On](https://github.com/matthewfritz/go-midi/issues/3)
   * ✅ [Note-Off](https://github.com/matthewfritz/go-midi/issues/4)
   * Polyphonic Key Pressure
   * Channel Pressure
   * Program Change
   * Pitch Bend Change
   * Control Change

#### Channel Voice Message Modifiers

   * ✅ [Velocity](https://github.com/matthewfritz/go-midi/issues/2)
   * ✅ [Running Status](https://github.com/matthewfritz/go-midi/issues/20)
   * Pitch
   * Modulation

#### System Common Messages

   * MTC Quarter Frame
   * Song Position Pointer
   * Song Select
   * Tune Request
   * End of Exclusive Messages (EOX)

#### System Timing Clock Messages

   * MIDI Start
   * MIDI Stop
   * MIDI Continue
   * Active Sensing
   * System Reset

#### System Exclusive Messages

   * Universal Non-Real-Time
   * Universal Real-Time

### MIDI 2.0 Roadmap

Once the [MIDI 1.0 Roadmap](#midi-10-roadmap) is at least halfway done I will start planning out the roadmap for implementing the current MIDI 2.0 specification.

## Resources

### Official Specifications

[Official MIDI Specifications](https://www.midi.org/specifications) - the official MIDI 1.0 and 2.0 specifications from the MIDI Alliance.

### Books

[_The Midi Manual_](https://www.amazon.com/dp/0367549980?psc=1&ref=ppx_yo2ov_dt_b_product_details) by David Miles Huber. Reading this book was the catalyst for breaking into MIDI programming and gave me enough of a technical foundation to begin this project.

### Assorted MIDI and Digital Music Information

#### Assorted MIDI 1.0 Information

[Summary of MIDI 1.0 Messages](https://www.midi.org/specifications-old/item/table-1-summary-of-midi-message) - ten-thousand foot view of MIDI 1.0 messages from the MIDI Association.

[GM 1 Sound Set](https://www.midi.org/specifications-old/item/gm-level-1-sound-set) - general list of the sounds available in General MIDI Level 1 from the MIDI Association.

[MIDI Tutorial for Programmers (Carnegie-Mellon University)](https://www.cs.cmu.edu/~music/cmsip/readings/MIDI%20tutorial%20for%20programmers.html) | [MIDI Tutorial for Programmers (Music-Software-Development.com)](http://www.music-software-development.com/midi-tutorial.html) - tutorial about MIDI 1.0 from a software development perspective.

[David's MIDI Spec](https://www.cs.cmu.edu/~music/cmsip/readings/davids-midi-spec.htm) - distilled version of the MIDI 1.0 specification from 1995 written by David Van Brink.

[General MIDI Instrument Patch Map](https://www.cs.cmu.edu/~music/cmsip/readings/GMSpecs_Patches.htm) - general MIDI instruments and their family groups.

[General MIDI Percussion Key Map](https://www.cs.cmu.edu/~music/cmsip/readings/GMSpecs_PercMap.htm) - general MIDI note numbers within channel 10 (percussion channel).

#### Digital Music Information

[An Introduction to Music Concepts](https://www.cs.cmu.edu/~music/cmsip/readings/music-theory-java.htm) - an introduction to music concepts, geared toward software developers, with examples in Java and written by Roger B. Dannenberg.

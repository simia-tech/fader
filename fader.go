// Copyright 2014 The fader authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package provides an interface to store and fetch items. The implementation
// is responsible for the removal after a expiry period.
//
// Example for a memory fader, that expires items after 2 seconds
//
//     memoryFader := fader.NewMemory(2*time.Second)
//     memoryFader.Open()
//     defer memoryFader.Close()
//
//     memoryFader.Store(item)
//     memoryFader.Size() // => 1
//
//     time.Sleep(3*time.Second)
//     memoryFader.Size() // => 0
//
// The multicast fader can be used to distribute `Store` operations via a multicast
// group. Other instances that listen to the same group, will perform that operation
// on thier own, so that each instance end up with the same data.
//
//    multicastFaderOne := fader.NewMulticast(memoryFaderOne, "224.0.0.1:1888", fader.DefaultKey)
//    multicastFaderOne.Open()
//    defer multicastFaderOne.Close()
//
//    multicastFaderTwo := fader.NewMulticast(memoryFaderTwo, "224.0.0.1:1888", fader.DefaultKey)
//    multicastFaderTwo.Open()
//    defer multicastFaderTwo.Close()
//
//    multicastFaderOne.Store(item)
//    multicastFaderOne.Size() // => 1
//
//    time.Sleep(10*time.Millisecond)
//
//    multicastFaderTwo.Size() // => 1
package fader

type Fader interface {
	Open() error
	Close() error

	Store(Item) error
	Earliest() Item
	Select(string) []Item
	Detect(string) Item
	Size() int
}

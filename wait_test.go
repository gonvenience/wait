// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package wait_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gonvenience/bunt"
	. "github.com/gonvenience/wait"
)

var _ = Describe("wait package tests", func() {
	BeforeEach(func() {
		SetColorSettings(OFF, OFF)
	})

	AfterEach(func() {
		SetColorSettings(AUTO, AUTO)
	})

	Context("elapsed time", func() {
		It("should return human readable elapsed time", func() {
			timeText := func(elapsedTime time.Duration) string {
				text, _ := TimeInfoText(elapsedTime)
				return text
			}

			Expect(timeText(500 * time.Millisecond)).To(BeEquivalentTo("less than a second"))
			Expect(timeText(1 * time.Second)).To(BeEquivalentTo("1 sec"))
			Expect(timeText(2 * time.Second)).To(BeEquivalentTo("2 sec"))
			Expect(timeText(60 * time.Second)).To(BeEquivalentTo("1 min"))
			Expect(timeText(90 * time.Second)).To(BeEquivalentTo("1 min 30 sec"))
			Expect(timeText(120 * time.Second)).To(BeEquivalentTo("2 min"))
			Expect(timeText(90 * time.Minute)).To(BeEquivalentTo("1 h 30 min"))
			Expect(timeText(90*time.Minute + 45*time.Second)).To(BeEquivalentTo("1 h 30 min 45 sec"))
		})
	})
})

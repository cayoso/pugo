/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package pure1

import (
	"testing"
)

func TestPure1Metrics(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetMetrics", testPure1GetMetrics(c))
}

func testPure1GetMetrics(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Metrics.GetMetrics(nil)
		if err != nil {
			t.Fatalf("error getting metrics: %s", err)
		}
	}
}
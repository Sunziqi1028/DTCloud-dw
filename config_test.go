/**
 * @Author: Sun
 * @Description:
 * @File:  config_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/15 17:03
 */

package main

import (
	"fmt"
	"testing"
	"wg315/config"
)

func TestNewConfig(t *testing.T) {
	c := config.NewConfig()
	if c == nil {
		t.Error(c)
	}
	fmt.Println(c)
}

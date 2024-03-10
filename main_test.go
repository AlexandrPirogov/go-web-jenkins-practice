package main

import (
   "testing"
)

func TestFoo(t* testing.T) {
   if 1 != 1 {
      t.Fatal("tests not passed")
   }
}

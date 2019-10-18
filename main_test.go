package main

import "testing"

func TestHandleRequest(t *testing.T) {
  var event BasicEvent 
  event.Input = "helloworld"
  msg, err := HandleRequest(nil, event)
  if err != nil {
    t.Errorf("Error was not nil!")
  }
  assertEqual(t, msg, "Hello World! event.Input=helloworld")
}

func assertEqual(t *testing.T, s1 string, s2 string) {
  if s1 != s2 {
    t.Errorf("ERROR: Strings Don't Match: [%s] [%s]", s1, s2)
  }
}

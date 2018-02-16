package main

import (
	"testing"
)

func TestGetParameterReg(t *testing.T) {
	
	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "waiting") != "13" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "writing") != "12" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "reading") != "10" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "requests") != "2588670" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "handled") != "2607076" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "accepts") != "2607076" {
		t.Error("waiting error")
	}

	if GetParameterReg([]byte(`Active connections: 22
server accepts handled requests
 2607076 2607076 2588670
Reading: 10 Writing: 12 Waiting: 13`), "connections") != "22" {
		t.Error("waiting error")
	}
}


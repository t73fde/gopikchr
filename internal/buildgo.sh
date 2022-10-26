#!/usr/bin/env bash
golemon -q -l -T lempar.go.tpl pikchr.y && gofmt -w pikchr.go

package main

import "embed"

//go:embed frontend/dist/*
var content embed.FS

//go:embed index.html
var index embed.FS

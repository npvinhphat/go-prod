package assets

import "embed"

//go:embed stacks/* templates/*
var Assets embed.FS

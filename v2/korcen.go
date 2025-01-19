package main

import (
	"github.com/fluffy-melli/GradFlow"
	"github.com/fluffy-melli/korcen-go/v2/cache"
)

// Copyright© All rights reserved.
//  _____                 _
// |_   _|_ _ _ __   __ _| |_
//   | |/ _` | '_ \ / _` | __|
//   | | (_| | | | | (_| | |_
//   |_|\__,_|_| |_|\__,_|\__|

func main() {
	GD := GradFlow.NewGradientDescent(cache.PreText(), cache.PreWord(), 0.01, 25)
	GD.Predict("야옹")
}

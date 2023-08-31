package main

import (
	"log"
	"os"

	gl "github.com/alden-scientific/golightly"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please include files for training and validation")
	}
	trainFile := os.Args[1]
	validFile := os.Args[2]

	model := gl.CreateSimpleRegressionModel(trainFile, validFile)
	var numModels int32
	cErr := gl.BoosterNumberOfTotalModel(model, &numModels)
	if cErr != 0 {
		log.Fatal("Error getting number of trees in model")
	}
	cErr = gl.BoosterSaveModel(model, 0, numModels, gl.Importance_gain, "./trained_model.txt")
	if cErr != 0 {
		log.Fatal("Error saving model")
	}
}

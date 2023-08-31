package golightly

import (
	"fmt"
	"log"
	"testing"
)

func TestLightgbm(t *testing.T) {
	InitializeLightgbm("./lib_lightgbm.so")
	var cErr int

	trained := CreateSimpleRegressionModel("winequality_red_train.tsv", "winequality_red_valid.tsv")

	var numModels int32
	cErr = BoosterNumberOfTotalModel(trained, &numModels)
	fmt.Println("saving model to trained_model.txt")
	cErr = BoosterSaveModel(trained, 0, numModels, 1, "./trained_model.txt")

	// get number of features in model
	var numFeatures int32
	cErr = BoosterGetNumFeature(trained, &numFeatures)
	fmt.Println("number of features in model:", numFeatures)

	// get importance scores
	importance := make([]float64, numFeatures)
	cErr = BoosterFeatureImportance(trained, numModels, Importance_gain, importance)
	fmt.Printf("%+v\n", importance)

	// output model to string
	var outLen64 int64
	modelString := make([]byte, 200_000)
	cErr = BoosterSaveModelToString(trained, 0, numModels, 1, 200_000, &outLen64, modelString)
	fmt.Println("length of model in bytes:", outLen64)
	fmt.Println("----------------Start of model example ----------------------")
	fmt.Println(string(modelString[0:200]))
	fmt.Println("----------------End of model example ----------------------")

	// load model file
	modelFile := "./trained_model.txt"
	iter := int32(0)
	var booster Booster
	cErr = BoosterCreateFromModelfile(modelFile, &iter, &booster)
	if cErr != 0 {
		log.Fatal(cErr)
	}

	// configure fast prediction
	data := []float32{7.4, 0.7, 0, 1.9, 0.076, 11, 34, 0.9978, 3.51, 0.56, 9.4}
	dataType := int32(0)
	nCols := int32(len(data))
	startIter := int32(0)
	numIter := int32(64)
	parameter := ""
	var fastConfig FastConfig
	cErr = BoosterPredictForMatSingleRowFastInit(trained, PredictNormal, startIter, numIter, dataType, nCols, parameter, &fastConfig)

	// make prediction
	var out float64
	cErr = BoosterPredictForMatSingleRowFast(fastConfig, data, &outLen64, &out)
	fmt.Println("Predicted quality:", out)

	cErr = BoosterFree(trained)
}

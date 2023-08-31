package golightly

import (
	"fmt"
	"log"

	"github.com/ebitengine/purego"
)

type Dataset uintptr
type Booster uintptr
type FastConfig uintptr
type ByteBuffer uintptr

const Dtype_float32 = int32(0)
const Dtype_float64 = int32(1)
const Importance_split = int32(0)
const Importance_gain = int32(1)
const Col_major = int32(0)
const Row_major = int32(1)
const PredictNormal = int32(0)
const PredictRaw = int32(1)
const PredictLeaf = int32(2)
const PredictContrib = int32(3)

var BoosterAddValidData func(booster Booster, valid_data Dataset) int
var BoosterCalcNumPredict func(booster Booster, num_row int32, predict_type int32, start_iteration int32, num_iteration int32, outlen *int64) int
var BoosterCreate func(train_data Dataset, parameters string, out *Booster) int
var BoosterCreateFromModelfile func(filename string, out_num_iterations *int32, out *Booster) int
var BoosterDumpModel func(booster Booster, start_iteration int32, num_iteration int32, feature_importance_type int32, buffer_len int64, outlen *int64, out_str string) int
var BoosterFeatureImportance func(booster Booster, num_iteration int32, importance_type int32, out_results []float64) int
var BoosterFree func(booster Booster) int
var BoosterGetCurrentIteration func(booster Booster, out_iteration *int32) int
var BoosterGetEval func(booster Booster, data_idx int32, out_len *int32, out_results *[2]float64) int
var BoosterGetEvalCounts func(booster Booster, out *int32) int
var BoosterGetLeafValue func(booster Booster, tree_idx int32, leaf_idx int32, out_val *float64) int
var BoosterGetLinear func(booster Booster, out *int32) int
var BoosterGetLoadedParam func(booster Booster, buffer_len int64, outlen *int64, out_str string) int
var BoosterGetLowerBoundValue func(booster Booster, out_results *float64) int
var BoosterGetNumClasses func(booster Booster, out *int32) int
var BoosterGetNumFeature func(booster Booster, out_len *int32) int
var BoosterGetNumPredict func(booster Booster, data_idx int32, out_len *int64) int
var BoosterGetPredict func(booster Booster, data_idx int32, out_len *int64, out_result *float64) int
var BoosterGetUpperBoundValue func(booster Booster, out_results *float64) int
var BoosterLoadModelFromString func(model_str string, out_num_iterations *int32, out *Booster) int
var BoosterMerge func(booster Booster, other_booster Booster) int
var BoosterNumberOfTotalModel func(booster Booster, out *int32) int
var BoosterNumModelPerIteration func(booster Booster, out_tree_per_iteration *int32) int
var BoosterPredictForFile func(booster Booster, data_filename string, data_has_header int32, predict_type int32, start_iteration int32, num_iteration int32, parameter string, result_filename string) int
var BoosterPredictForMat func(booster Booster, data uintptr, data_type int32, nrow int32, ncol int32, is_row_major int32, predict_type int32, start_iteration int32, num_iteration int32, parameter string, outlen *int64, out_result *float64) int
var BoosterPredictForMats func(booster Booster, data *uintptr, data_type int32, nrow int32, ncol int32, predict_type int32, start_iteration int32, num_iteration int32, parameter string, outlen *int64, out_result *float64) int
var BoosterPredictForMatSingleRow func(booster Booster, data uintptr, data_type int32, ncol int32, is_row_major int32, predict_type int32, start_iteration int32, num_iteration int32, parameter string, outlen *int64, out_result *float64) int
var BoosterPredictForMatSingleRowFast func(fastConfig FastConfig, data []float32, outlen *int64, out_result *float64) int
var BoosterPredictForMatSingleRowFastInit func(booster Booster, predict_type int32, start_iteration int32, num_iteration int32, data_type int32, ncol int32, parameter string, out_fastConfig *FastConfig) int
var BoosterRefit func(booster Booster, leaf_preds *int32, nrow int32, ncol int32) int
var BoosterResetParameter func(booster Booster, parameters string) int
var BoosterResetTrainingData func(booster Booster, train_data Dataset) int
var BoosterRollbackOneIter func(booster Booster) int
var BoosterSaveModel func(booster Booster, start_iteration int32, num_iteration int32, feature_importance_type int32, filename string) int
var BoosterSaveModelToString func(booster Booster, start_iteration int32, num_iteration int32, feature_importance_type int32, buffer_len int64, outlen *int64, out_str []byte) int
var BoosterSetLeafValue func(booster Booster, tree_idx int32, leaf_idx int32, val float64) int
var BoosterShuffleModels func(booster Booster, start_iter int32, end_iter int32) int
var BoosterUpdateOneIter func(booster Booster, is_finished *int32) int
var BoosterUpdateOneIterCustom func(booster Booster, grad *float32, hess *float32, is_finished *int32) int
var BoosterValidateFeatureNames func(booster Booster, data_names *string, data_num_features int32) int // note: data_names was originally const char**
var ByteBufferFree func(byteBuffer ByteBuffer) int
var ByteBufferGetAt func(byteBuffer ByteBuffer, index int32, out_val *uint8) int
var DatasetAddFeaturesFrom func(target Dataset, source Dataset) int
var DatasetCreateByReference func(reference Dataset, num_total_row int64, out *Dataset) int
var DatasetCreateFromFile func(filename string, parameters string, reference Dataset, out *Dataset) int
var DatasetCreateFromMat func(data [][]float64, data_type int32, nrow int32, ncol int32, is_row_major int32, parameters string, reference Dataset, out *Dataset) int
var DatasetCreateFromMats func(nmat int32, data *uintptr, data_type int32, nrow *int32, ncol int32, is_row_major int32, parameters string, reference Dataset, out *Dataset) int
var DatasetDumpText func(dataset Dataset, filename string) int
var DatasetFree func(dataset Dataset) int
var DatasetGetFeatureNumBin func(dataset Dataset, feature int32, out *int32) int
var DatasetGetField func(dataset Dataset, field_name string, out_len *int32, out_ptr *uintptr, out_type *int32) int
var DatasetGetNumData func(dataset Dataset, out *int32) int
var DatasetGetNumFeature func(dataset Dataset, out *int32) int
var DatasetGetSubset func(dataset Dataset, used_row_indices *int32, num_used_row_indices int32, parameters string, out *Dataset) int
var DatasetInitStreaming func(dataset Dataset, has_weights int32, has_init_scores int32, has_queries int32, nclasses int32, nthreads int32, omp_max_threads int32) int
var DatasetMarkFinished func(dataset Dataset) int
var DatasetPushRows func(dataset Dataset, data uintptr, data_type int32, nrow int32, ncol int32, start_row int32) int
var DatasetPushRowsWithMetadata func(dataset Dataset, data uintptr, data_type int32, nrow int32, ncol int32, start_row int32, label *float32, weight *float32, init_score *float64, query *int32, tid int32) int
var DatasetSaveBinary func(dataset Dataset, filename string) int
var DatasetSerializeReferenceToBinary func(dataset Dataset, out *ByteBuffer, out_len *int32) int
var DatasetSetFeatureNames func(dataset Dataset, feature_names *string, num_feature_names int32) int // note: feature_names was originally const char**
var DatasetSetField func(dataset Dataset, field_name string, field_data uintptr, num_element int32, type_data int32) int
var DatasetSetWaitForManualFinish func(dataset Dataset, wait int32) int
var DatasetUpdateParamChecking func(old_parameters string, new_parameters string) int
var DumpParamAliases func(buffer_len int64, out_len *int64, out_str string) int
var FastConfigFree func(fastConfig FastConfig) int
var GetSampleCount func(num_total_row int32, parameters string, out *int32) int
var NetworkFree func() int
var NetworkInit func(machines string, local_listen_port int32, listen_time_out int32, num_machines int32) int
var SampleIndices func(num_total_row int32, parameters string, out uintptr, out_len *int32) int

// var BoosterGetEvalNames func(booster Booster,  len int32, out_len *int32, const size_t buffer_len, size_t* out_buffer_len, char** out_strs) int
var BoosterGetFeatureNames func(booster Booster, len int32, out_len *int32, buffer_len int32, out_buffer_len []int32, out_strs [][]byte) int

// var NetworkInitWithFunctions func(num_machines int32, rank int32, void *reduce_scatter_ext_fun, void *allgather_ext_fun) int
// var RegisterLogCallback func(void (*callback)(const char*)) int
// var DatasetGetFeatureNames func(dataset Dataset, len int32, num_feature_names *int32, const size_t buffer_len, size_t* out_buffer_len, feature_names *string) int
// var DatasetCreateFromSampledColumn func(double **sample_data, sample_indices *int32, ncol int32, num_per_col *int32, num_sample_row int32, num_local_row int32, num_dist_row int64, parameters string, out *Dataset) int
// var DatasetCreateFromSerializedReference func(const void* ref_buffer, ref_buffer_size int32, num_row int64, num_classes int32, parameters string, DatasetHandle* out) int

func InitializeLightgbm(libpath string) {
	libLightgbm, err := purego.Dlopen(libpath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// new registrations
	// purego.RegisterLibFunc(&BoosterGetEvalNames, libLightgbm, "LGBM_BoosterGetEvalNames")
	purego.RegisterLibFunc(&BoosterGetFeatureNames, libLightgbm, "LGBM_BoosterGetFeatureNames")
	// purego.RegisterLibFunc(&BoosterPredictForMat, libLightgbm, "LGBM_BoosterPredictForMat")
	// purego.RegisterLibFunc(&BoosterPredictForMats, libLightgbm, "LGBM_BoosterPredictForMats")
	// purego.RegisterLibFunc(&BoosterPredictForMatSingleRow, libLightgbm, "LGBM_BoosterPredictForMatSingleRow")
	// purego.RegisterLibFunc(&DatasetCreateFromSampledColumn, libLightgbm, "LGBM_DatasetCreateFromSampledColumn")
	// purego.RegisterLibFunc(&DatasetCreateFromSerializedReference, libLightgbm, "LGBM_DatasetCreateFromSerializedReference")
	// purego.RegisterLibFunc(&DatasetGetFeatureNames, libLightgbm, "LGBM_DatasetGetFeatureNames")
	// purego.RegisterLibFunc(&DatasetPushRowsWithMetadata, libLightgbm, "LGBM_DatasetPushRowsWithMetadata")
	// purego.RegisterLibFunc(&NetworkInitWithFunctions, libLightgbm, "LGBM_NetworkInitWithFunctions")
	// purego.RegisterLibFunc(&RegisterLogCallback, libLightgbm, "LGBM_RegisterLogCallback")
	purego.RegisterLibFunc(&BoosterAddValidData, libLightgbm, "LGBM_BoosterAddValidData")
	purego.RegisterLibFunc(&BoosterCalcNumPredict, libLightgbm, "LGBM_BoosterCalcNumPredict")
	purego.RegisterLibFunc(&BoosterCreate, libLightgbm, "LGBM_BoosterCreate")
	purego.RegisterLibFunc(&BoosterCreateFromModelfile, libLightgbm, "LGBM_BoosterCreateFromModelfile")
	purego.RegisterLibFunc(&BoosterDumpModel, libLightgbm, "LGBM_BoosterDumpModel")
	purego.RegisterLibFunc(&BoosterFeatureImportance, libLightgbm, "LGBM_BoosterFeatureImportance")
	purego.RegisterLibFunc(&BoosterFree, libLightgbm, "LGBM_BoosterFree")
	purego.RegisterLibFunc(&BoosterGetCurrentIteration, libLightgbm, "LGBM_BoosterGetCurrentIteration")
	purego.RegisterLibFunc(&BoosterGetEval, libLightgbm, "LGBM_BoosterGetEval")
	purego.RegisterLibFunc(&BoosterGetEvalCounts, libLightgbm, "LGBM_BoosterGetEvalCounts")
	purego.RegisterLibFunc(&BoosterGetLeafValue, libLightgbm, "LGBM_BoosterGetLeafValue")
	purego.RegisterLibFunc(&BoosterGetLinear, libLightgbm, "LGBM_BoosterGetLinear")
	purego.RegisterLibFunc(&BoosterGetLoadedParam, libLightgbm, "LGBM_BoosterGetLoadedParam")
	purego.RegisterLibFunc(&BoosterGetLowerBoundValue, libLightgbm, "LGBM_BoosterGetLowerBoundValue")
	purego.RegisterLibFunc(&BoosterGetNumClasses, libLightgbm, "LGBM_BoosterGetNumClasses")
	purego.RegisterLibFunc(&BoosterGetNumFeature, libLightgbm, "LGBM_BoosterGetNumFeature")
	purego.RegisterLibFunc(&BoosterGetNumPredict, libLightgbm, "LGBM_BoosterGetNumPredict")
	purego.RegisterLibFunc(&BoosterGetPredict, libLightgbm, "LGBM_BoosterGetPredict")
	purego.RegisterLibFunc(&BoosterGetUpperBoundValue, libLightgbm, "LGBM_BoosterGetUpperBoundValue")
	purego.RegisterLibFunc(&BoosterLoadModelFromString, libLightgbm, "LGBM_BoosterLoadModelFromString")
	purego.RegisterLibFunc(&BoosterMerge, libLightgbm, "LGBM_BoosterMerge")
	purego.RegisterLibFunc(&BoosterNumberOfTotalModel, libLightgbm, "LGBM_BoosterNumberOfTotalModel")
	purego.RegisterLibFunc(&BoosterNumModelPerIteration, libLightgbm, "LGBM_BoosterNumModelPerIteration")
	purego.RegisterLibFunc(&BoosterPredictForFile, libLightgbm, "LGBM_BoosterPredictForFile")
	purego.RegisterLibFunc(&BoosterPredictForMatSingleRowFast, libLightgbm, "LGBM_BoosterPredictForMatSingleRowFast")
	purego.RegisterLibFunc(&BoosterPredictForMatSingleRowFastInit, libLightgbm, "LGBM_BoosterPredictForMatSingleRowFastInit")
	purego.RegisterLibFunc(&BoosterRefit, libLightgbm, "LGBM_BoosterRefit")
	purego.RegisterLibFunc(&BoosterResetParameter, libLightgbm, "LGBM_BoosterResetParameter")
	purego.RegisterLibFunc(&BoosterResetTrainingData, libLightgbm, "LGBM_BoosterResetTrainingData")
	purego.RegisterLibFunc(&BoosterRollbackOneIter, libLightgbm, "LGBM_BoosterRollbackOneIter")
	purego.RegisterLibFunc(&BoosterSaveModel, libLightgbm, "LGBM_BoosterSaveModel")
	purego.RegisterLibFunc(&BoosterSaveModelToString, libLightgbm, "LGBM_BoosterSaveModelToString")
	purego.RegisterLibFunc(&BoosterSetLeafValue, libLightgbm, "LGBM_BoosterSetLeafValue")
	purego.RegisterLibFunc(&BoosterShuffleModels, libLightgbm, "LGBM_BoosterShuffleModels")
	purego.RegisterLibFunc(&BoosterUpdateOneIter, libLightgbm, "LGBM_BoosterUpdateOneIter")
	purego.RegisterLibFunc(&BoosterUpdateOneIterCustom, libLightgbm, "LGBM_BoosterUpdateOneIterCustom")
	purego.RegisterLibFunc(&BoosterValidateFeatureNames, libLightgbm, "LGBM_BoosterValidateFeatureNames")
	purego.RegisterLibFunc(&ByteBufferFree, libLightgbm, "LGBM_ByteBufferFree")
	purego.RegisterLibFunc(&ByteBufferGetAt, libLightgbm, "LGBM_ByteBufferGetAt")
	purego.RegisterLibFunc(&DatasetAddFeaturesFrom, libLightgbm, "LGBM_DatasetAddFeaturesFrom")
	purego.RegisterLibFunc(&DatasetCreateByReference, libLightgbm, "LGBM_DatasetCreateByReference")
	purego.RegisterLibFunc(&DatasetCreateFromFile, libLightgbm, "LGBM_DatasetCreateFromFile")
	purego.RegisterLibFunc(&DatasetCreateFromMat, libLightgbm, "LGBM_DatasetCreateFromMat")
	purego.RegisterLibFunc(&DatasetCreateFromMats, libLightgbm, "LGBM_DatasetCreateFromMats")
	purego.RegisterLibFunc(&DatasetDumpText, libLightgbm, "LGBM_DatasetDumpText")
	purego.RegisterLibFunc(&DatasetFree, libLightgbm, "LGBM_DatasetFree")
	purego.RegisterLibFunc(&DatasetGetFeatureNumBin, libLightgbm, "LGBM_DatasetGetFeatureNumBin")
	purego.RegisterLibFunc(&DatasetGetField, libLightgbm, "LGBM_DatasetGetField")
	purego.RegisterLibFunc(&DatasetGetNumData, libLightgbm, "LGBM_DatasetGetNumData")
	purego.RegisterLibFunc(&DatasetGetNumFeature, libLightgbm, "LGBM_DatasetGetNumFeature")
	purego.RegisterLibFunc(&DatasetGetSubset, libLightgbm, "LGBM_DatasetGetSubset")
	purego.RegisterLibFunc(&DatasetInitStreaming, libLightgbm, "LGBM_DatasetInitStreaming")
	purego.RegisterLibFunc(&DatasetMarkFinished, libLightgbm, "LGBM_DatasetMarkFinished")
	purego.RegisterLibFunc(&DatasetPushRows, libLightgbm, "LGBM_DatasetPushRows")
	purego.RegisterLibFunc(&DatasetSaveBinary, libLightgbm, "LGBM_DatasetSaveBinary")
	purego.RegisterLibFunc(&DatasetSerializeReferenceToBinary, libLightgbm, "LGBM_DatasetSerializeReferenceToBinary")
	purego.RegisterLibFunc(&DatasetSetFeatureNames, libLightgbm, "LGBM_DatasetSetFeatureNames")
	purego.RegisterLibFunc(&DatasetSetField, libLightgbm, "LGBM_DatasetSetField")
	purego.RegisterLibFunc(&DatasetSetWaitForManualFinish, libLightgbm, "LGBM_DatasetSetWaitForManualFinish")
	purego.RegisterLibFunc(&DatasetUpdateParamChecking, libLightgbm, "LGBM_DatasetUpdateParamChecking")
	purego.RegisterLibFunc(&DumpParamAliases, libLightgbm, "LGBM_DumpParamAliases")
	purego.RegisterLibFunc(&FastConfigFree, libLightgbm, "LGBM_FastConfigFree")
	purego.RegisterLibFunc(&GetSampleCount, libLightgbm, "LGBM_GetSampleCount")
	purego.RegisterLibFunc(&NetworkFree, libLightgbm, "LGBM_NetworkFree")
	purego.RegisterLibFunc(&NetworkInit, libLightgbm, "LGBM_NetworkInit")
	purego.RegisterLibFunc(&SampleIndices, libLightgbm, "LGBM_SampleIndices")
}

func CreateSimpleRegressionModel(trainFile, validFile string) Booster {
	var cErr int
	InitializeLightgbm("./lib_lightgbm.so")
	var nilDS Dataset
	var trainDS Dataset
	cErr = DatasetCreateFromFile(trainFile, "header=true", nilDS, &trainDS)
	if cErr != 0 {
		log.Fatal("failed to load training file", trainFile)
	}
	var validDS Dataset
	cErr = DatasetCreateFromFile(validFile, "header=true", trainDS, &validDS)
	if cErr != 0 {
		log.Fatal("failed to load validation file", trainFile)
	}

	var model Booster
	cErr = BoosterCreate(trainDS, "objective=regression metric=l2,l1 min_data_in_leaf=10", &model)
	cErr = BoosterAddValidData(model, validDS)
	var isFinished int32
	dataIndex := int32(1)
	var outLen int32
	var outIter int32
	// var outResults float64
	testOut := [2]float64{0, 0}
	minL2 := 100.0
	minL1 := 100.0
	bestIter := int32(0)
	notImprovedIters := 0
	for i := 0; i < 100; i++ {
		cErr = BoosterUpdateOneIter(model, &isFinished)
		cErr = BoosterGetCurrentIteration(model, &outIter)
		cErr = BoosterGetEval(model, dataIndex, &outLen, &testOut)
		l2, l1 := testOut[0], testOut[1]
		switch {
		case l1 < minL1:
			minL1 = l1
			bestIter = outIter
			notImprovedIters = 0
		case l2 < minL2:
			minL2 = l2
			bestIter = outIter
			notImprovedIters = 0
		default:
			notImprovedIters++
		}
		fmt.Print("Booster iter:", outIter, " finished:", isFinished, " ")
		fmt.Printf("l2: %0.4g l1: %0.4g\n", testOut[0], testOut[1])
		if notImprovedIters > 4 {
			break
		}
	}
	var numModels int32
	cErr = BoosterNumberOfTotalModel(model, &numModels)
	for numModels > bestIter {
		BoosterRollbackOneIter(model)
		cErr = BoosterNumberOfTotalModel(model, &numModels)
	}
	cErr = DatasetFree(trainDS)
	cErr = DatasetFree(validDS)
	return model
}

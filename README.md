## Golightly for Machine Learning in Go

### Golightly provides the building blocks for an ML API
Go is widely used to create APIs and other networked services, but the Go ecosystem has few native ML libraries. Golightly provides ML models for regression, classification, ranking, and more.

### Golightly brings the LightGBM ML library into the Go ecosystem
[LightGBM](https://github.com/microsoft/LightGBM) is an MIT-licensed, high performance machine learning framework developed by Microsoft. It is  comparable to XGBoost in predictive performance but much faster to train. For high performance, it is written in C++.

Golightly uses [purego](https://github.com/ebitengine/purego) to dynamically link to the C++ LightGBM library, which avoids CGO and the associated complications to building a Go project.

Golightly is a work in progress with some basic functions working well (e.g., enough to create a prediction API for existing models or train new models from files), but with more work needed to fully recapitulate LightGBM functionality. Contributions are welcome, particularly if you know C/C++.

Golightly is currently a very thin wrapper for LightGBM, but a small number of convenience functions are planned. Currently, the library is only tested on Linux, but should work on MacOS and Windows with minimal effort (just swap the dynamic library file, e.g, dll, dylib).

### Links to LightGBM and Example Data
For Linux, download the [pre-compiled static object for LightGBM](https://github.com/microsoft/LightGBM/releases/download/v4.0.0/lib_lightgbm.so).

Wine quality data from:
Modeling wine preferences by data mining from physicochemical properties
By P. Cortez, A. Cerdeira, Fernando Almeida, Telmo Matos, J. Reis. 2009

Original wine data is from the [UCI ML Repository](https://doi.org/10.24432/C56S3T)

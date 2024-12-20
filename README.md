# Big Data Project: MapReduce, Hive, and Apache Airflow

## Overview
This project demonstrates the application of Big Data technologies for processing and analyzing datasets. The workflow includes:
1. A **MapReduce** implementation (Hadoop Streaming) written in Go, which can be tested locally using Linux or Docker.
2. A **Hive query script**, designed for execution on a GCP-based Hive cluster, to perform data integration and analysis.
3. Automation using **Apache Airflow** to orchestrate the pipeline on GCP.

## Project Structure
- **data/** - Directory containing input and output data.
  - `input/` - Input datasets required for processing.
  - `output/` - Stores the results of data processing.
- **src/** - Source code directory.
  - `mapreduce/` - MapReduce implementation in Go.
    - `mapper.go` - Mapper logic for Hadoop Streaming.
    - `combiner.go` - Combiner logic for optimizing intermediate data.
    - `reducer.go` - Reducer logic for final aggregation.
  - `hive/` - Hive script for processing MapReduce outputs and integrating with additional datasets.
    - `query.hql` - Hive query designed for GCP-based clusters.
  - `airflow/` - Apache Airflow DAG to automate workflow execution.
- **docker/** - Docker-related files for local testing.
  - `Dockerfile` - Builds an environment to test MapReduce locally using Hadoop Streaming.
- **README.md** - This file, explaining the project structure and usage.
- **.gitignore** - Defines files and directories to ignore in the Git repository.

## Usage Context
This project is designed for deployment and execution on a GCP cluster. However:
- The **MapReduce** components (`.go` files) can also be tested locally using Linux tools or Docker. Hadoop does not need to be installed locally, as the pipeline can be simulated with shell commands.

- The **Hive query script** is designed for execution on a GCP-based Hive cluster. While it is technically possible to adapt the script to standard SQL (e.g., for MySQL or another RDBMS) with minor modifications, its correct execution in such environments is not guaranteed.

## How to Run

The instructions below assume that the environment is either Linux-based or a Docker container emulating a Linux environment. This ensures compatibility with the required shell commands and tools.

### Local Testing (MapReduce Only)
1. **Prepare input data**:
   - Place your data files in a directory, e.g., `input_folder`.
2. **Run the MapReduce pipeline**:
   - Use the following shell command:
     ```bash
     cat input_folder | path_to_mapper | path_to_combiner | sort | path_to_reducer > output.tsv
     ```
   - If you're using Windows, run this command inside a Docker container with a Linux environment.
  
   This command allows testing of the MapReduce logic in a Unix-like environment, using standard input/output redirection.
3. **Check results**:
   - The output will be written to `output.tsv`.

### GCP Deployment
1. **Prepare Data**:
   - Upload datasets to the GCP storage bucket (`gs://`).
   - Ensure datasets are structured as:
     - `project/input/dataset1` - For MapReduce.
     - `project/input/dataset2` - For Hive.

2. **Run MapReduce**:
   - Submit the MapReduce job to the Hadoop cluster on GCP.

3. **Run Hive Query**:
   - Execute the Hive script on the GCP-based Hive cluster.

4. **Automate with Airflow**:
   - Configure and deploy the Airflow DAG to orchestrate the workflow.

## Notes
- It is recommended to use small datasets locally for testing MapReduce before deploying to GCP for large-scale processing.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author
Developed as part of a Big Data course project.

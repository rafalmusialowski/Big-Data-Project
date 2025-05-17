# Big Data Project: MapReduce, Hive, and Apache Airflow

## Business Context and Objective

This project focuses on analyzing movie industry data extracted from IMDb. It involves two main datasets:

* **Datasource1**: Contains information about people involved in films, with fields such as:

    * `tconst` – movie identifier
    * `ordering` – person's sequence number in a movie
    * `nconst` – person identifier
    * `role` – role (e.g., actor, actress, self)
    * `job` – job title (if applicable)
    * `characters` – character names (if applicable)

* **Datasource4**: Contains detailed person information, including:

    * `nconst` – person identifier
    * `primaryName` – person's full name
    * `birthYear`, `deathYear` – birth and death years
    * `primaryProfession` – main professions (e.g., actor, director)
    * `knownForTitles` – identifiers of movies the person is known for

### Project Goal

The project aims to determine:

* For each person: the number of films they have acted in and directed, based on **Datasource1**.
* Identify:

    * The top three actors (based on the number of movies they acted in).
    * The top three directors (based on the number of movies they directed).

Using **MapReduce**, we first aggregate per-person counts of acting and directing activities.

Using **Hive**, we then enrich this data by joining it with **Datasource4**, and select the individuals with the highest activity in each profession.

The final output contains:

* `primaryName` – name of the actor or director,
* `role` – whether they are an actor or director,
* `movies` – the number of movies associated with them.

This structured output provides insights into the most prolific individuals in the movie industry based on their roles.

## Technical Approach

The objective of this project was to practically apply core Big Data processing platforms to real-world datasets. The work involved:

* **Processing one dataset using MapReduce (Hadoop Streaming)**: Filtering and aggregating data through programs written in Go.
* **Analyzing and integrating data using Hive**: Joining the processed dataset with a second dataset, followed by additional aggregation, sorting, and filtering.
* **Automating the workflow with Apache Airflow**: Building a DAG to orchestrate the entire process on a Google Cloud Platform (GCP) environment.

The final result is a fully processed and integrated dataset exported in JSON format.

## Overview

This project demonstrates the application of Big Data technologies for processing and analyzing datasets. The workflow includes:

1. A **MapReduce** implementation (Hadoop Streaming) written in Go, which can be tested locally using Linux or Docker.
2. A **Hive query script**, designed for execution on a GCP-based Hive cluster, to perform data integration and analysis.
3. Automation using **Apache Airflow** to orchestrate the pipeline on GCP.

## Project Structure

* **data/** - Directory containing input and output data.

    * `input/` - Input datasets required for processing.

        * `datasource1/` - First dataset used by the MapReduce pipeline.
        * `datasource4/` - Second dataset used by the Hive script.
    * `output/` - Stores the results of data processing.
* **src/** - Source code directory.

    * `mapreduce/` - MapReduce implementation in Go.

        * `mapper.go` - Mapper logic for Hadoop Streaming.
        * `combiner.go` - Combiner logic for optimizing intermediate data.
        * `reducer.go` - Reducer logic for final aggregation.
    * `hive/` - Hive script for processing MapReduce outputs and integrating with additional datasets.

        * `query.hql` - Hive query designed for GCP-based clusters.
    * `airflow/` - Apache Airflow DAG to automate workflow execution.
* **docker/** - Docker-related files for local testing.

    * `Dockerfile` - Builds an environment to test MapReduce locally using Hadoop Streaming.
* **README.md** - This file, explaining the project structure and usage.
* **.gitignore** - Defines files and directories to ignore in the Git repository.

## Usage Context

This project is designed for deployment and execution on a GCP cluster. However:

* The **MapReduce** components (`.go` files) can also be tested locally using Linux tools or Docker. Hadoop does not need to be installed locally, as the pipeline can be simulated with shell commands.

* The **Hive query script** is designed for execution on a GCP-based Hive cluster. While it is technically possible to adapt the script to standard SQL (e.g., for MySQL or another RDBMS) with minor modifications, its correct execution in such environments is not guaranteed.

## How to Run

### Local Testing (MapReduce Only)

1. **Prepare input data**:

    * Download the dataset manually from Google Drive:
      [Google Drive - Source Data](https://drive.google.com/drive/folders/1UUu-A6qtEwHEFl7YBrIojvT_vR35fgEi?usp=sharing)

    * After downloading, organize the data locally into the following structure:

   ```
   project-root/
   └── data/
       └── input/
           ├── datasource1/
           │   ├── part-00000.tsv
           │   ├── part-00001.tsv
           │   └── ...
           └── datasource4/
               ├── part-00000.tsv
               ├── part-00001.tsv
               └── ...
   ```

2. **Run the MapReduce pipeline**:

    * First, build the Docker image from the provided Dockerfile:

      ```bash
      docker build -t bigdata-local -f docker/Dockerfile .
      ```

    * Then, run the container with the project directory mounted inside:

      ```bash
      docker run -it --rm -v $(pwd):/app bigdata-local
      ```
     
      *(On Windows, use `%cd%` instead of `$(pwd)` if running in CMD or PowerShell)*

   * Inside the container, navigate to the `/app` directory if not already there, then run:

      ```bash
      cd /app
      ```

    * Reset the output directory:

      ```bash
      rm -rf data/output/
      mkdir -p data/output/
      ```

    * Run the local Hadoop Streaming MapReduce pipeline:

      ```bash
      cat data/input/datasource1/* | mapper | combiner | sort | reducer > data/output/output.tsv
      ```

   This command allows testing of the MapReduce logic in a Unix-like environment, using standard input/output redirection.

3. **Check results**:

    * The output will be written to `data/output/output.tsv`.
    * This file will be used as input for the Hive script in the next stage of the processing pipeline.

### GCP Deployment

1. **Prepare Data**:

    * Upload datasets to the GCP storage bucket (`gs://`).
    * Ensure datasets are structured as:

        * `project/input/dataset1` - For MapReduce.
        * `project/input/dataset4` - For Hive.

2. **Run MapReduce**:

    * Submit the MapReduce job to the Hadoop cluster on GCP.

3. **Run Hive Query**:

    * Execute the Hive script on the GCP-based Hive cluster.

4. **Automate with Airflow**:

    * Configure and deploy the Airflow DAG to orchestrate the workflow.

## Notes

* It is recommended to use small datasets locally for testing MapReduce before deploying to GCP for large-scale processing.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

Developed as part of a Big Data course project.

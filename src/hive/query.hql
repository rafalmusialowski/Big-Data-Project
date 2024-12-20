
SET input_dir3=${hiveconf:input_dir3};
SET input_dir4=${hiveconf:input_dir4};
SET output_dir6=${hiveconf:output_dir6};

CREATE EXTERNAL TABLE IF NOT EXISTS mapred_output (
    nconst STRING,
    playedMovies INT,
    directedMovies INT
)
ROW FORMAT DELIMITED 
FIELDS TERMINATED BY '\t' 
STORED AS TEXTFILE 
LOCATION '${hiveconf:input_dir3}';

CREATE EXTERNAL TABLE IF NOT EXISTS name_basics (
    nconst STRING,
    primaryName STRING,
    birthYear INT,
    deathYear INT,
    primaryProfession STRING,
    knownForTitles STRING
)
ROW FORMAT DELIMITED 
FIELDS TERMINATED BY '\t' 
STORED AS TEXTFILE 
LOCATION '${hiveconf:input_dir4}'
TBLPROPERTIES ("skip.header.line.count"="1");

CREATE TABLE IF NOT EXISTS top_actors_n_directors AS  
SELECT * FROM (
    SELECT 
        n.primaryName, 
        CASE 
            WHEN n.primaryProfession LIKE '%actress%' THEN 'actress'
            ELSE 'actor'
        END AS role,  
        m.playedMovies AS movies
    FROM mapred_output m
    JOIN name_basics n ON m.nconst = n.nconst
    WHERE (n.primaryProfession LIKE '%actor%' OR n.primaryProfession LIKE '%actress%')
      AND m.playedMovies > 0  
    ORDER BY m.playedMovies DESC
    LIMIT 3
) AS top_actors
UNION
SELECT * FROM (
    SELECT 
        n.primaryName, 
        'director' AS role,  
        m.directedMovies AS movies
    FROM mapred_output m
    JOIN name_basics n ON m.nconst = n.nconst
    WHERE n.primaryProfession LIKE '%director%'  
      AND m.directedMovies > 0  
    ORDER BY m.directedMovies DESC
    LIMIT 3
) AS top_directors
ORDER BY movies DESC;

INSERT OVERWRITE DIRECTORY '${hiveconf:output_dir6}'
ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.JsonSerDe'
SELECT * FROM top_actors_n_directors;
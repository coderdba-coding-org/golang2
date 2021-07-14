===========================
QUERY OUTPUT JSON FORMAT
===========================
-----------------------------------------
QUERYING INFLUX DB IN BROWSER DIRECTLY
-----------------------------------------
Typical result of a topsql query:
{"results":[{"statement_id":0,"series":[{"name":"oracle_topsql_elpsd","columns":["time","count"],"values":[["1970-01-01T00:00:00Z",4]]}]}]}

----------------------------------------------------------------------------------
FROM GOLANG PROGRAM Influx-Client 'result' - "[]client.Result" object
----------------------------------------------------------------------------------
As in URL /trialquery1
[
    {
        "statement_id": 0,
        "Series": [
            {
                "name": "oracle_dbstatus",
                "columns": [
                    "time",
                    "value"
                ],
                "values": [
                    [
                        "2021-07-13T07:46:53.608802522Z",
                        28
                    ]
                ]
            }
        ],
        "Messages": null
    }
]

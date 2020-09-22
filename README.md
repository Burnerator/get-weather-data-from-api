# Application to get weather data using AEMET data API

This is an example of how to consume web services using Go language https://golang.org/.
Application recover weather data from an indicated meteorologic station. The service is provided
by Spanish AEMET (Agencia Estatal de Meteorología)
https://opendata.aemet.es/dist/index.html

## Build

1) Install Go. Follow instructions https://golang.org/doc/install
2) Clone this git project.
3) Go to downloaded folder /get-weather-data-from-api/
4) Execute command

        go build

## Run

1) Execute command

        ./get-weather-data-from-api [startDate] [endDate] [station] [resFile] [apiKey]

    **Params**

* **starDate** --> Filter start date
* **endDate** --> Filter end date
* **station** --> Filter station id
* **resFile** --> Path results file. For instance: ./results.csv
* **apiKey** --> Authorized API key. You can get it in https://opendata.aemet.es/centrodedescargas/inicio

    _Example: ./get-weather-data-from-api 2020-01-03 2020-09-30 3196 ./result.csv eyJhbGciO...hL0wrslt_

2) Get results

        Fecha;Nombre;TMin;TMax;PresMin;PresMax
        2020-09-01;MADRID, CUATRO VIENTOS;15,3;28,4;933,4;938,1
        2020-09-02;MADRID, CUATRO VIENTOS;15,7;30,3;935,8;940,3
        2020-09-03;MADRID, CUATRO VIENTOS;16,0;31,8;940,4;943,9
        2020-09-04;MADRID, CUATRO VIENTOS;18,1;32,7;941,7;946,1
        2020-09-05;MADRID, CUATRO VIENTOS;17,4;33,8;937,6;942,8
        2020-09-06;MADRID, CUATRO VIENTOS;18,5;33,3;935,8;940,0
        2020-09-07;MADRID, CUATRO VIENTOS;14,5;28,2;938,4;943,2
        2020-09-08;MADRID, CUATRO VIENTOS;12,2;27,3;941,6;944,9
        2020-09-09;MADRID, CUATRO VIENTOS;14,7;29,1;940,4;944,5
        2020-09-10;MADRID, CUATRO VIENTOS;18,2;31,0;937,3;942,3
        2020-09-11;MADRID, CUATRO VIENTOS;17,1;32,4;935,2;938,5
        2020-09-12;MADRID, CUATRO VIENTOS;19,2;33,1;937,5;941,8
        2020-09-13;MADRID, CUATRO VIENTOS;19,0;32,8;941,8;944,8

* **Fecha** --> Data date
* **Nombre** --> Meteorologic station name
* **TMin** --> Minimum Temperature (Celsius)
* **TMax** --> Maximum Temperature (Celsius)
* **PresMin** --> Minimum Pressure
* **PresMax** --> Maximum Pressure
* **Prec** --> Precipitations

Results will be showed in screen and saved in disc. File will be created in path file (as indicated in param resFile)

# go-morocco-weather
This program is structured on three parts:
  1. Fetching cities (gountries)
  2. Computing their weather with coordinates (openweathermap)
  3. Saving the data into an excel file (excelize)
  
Don't forget to add dependencies:
 ``` 
 go get github.com/pariz/gountries
 go get github.com/briandowns/openweathermap
 go get github.com/360EntSecGroup-Skylar/excelize 
 ```
 
Optional(for wiki_parser.go): 
```
go get github.com/PuerkitoBio/goquery
```

You can run the program by typing this command:
  ```
  go run main.go
  ```

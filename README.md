<h1>Description</h1>

I made a simple API that returns country information in an organized JSON format in order to practice Golang. <br>
This is meant to be used as an example and a starting off point.

<h1>Examples: </h1>

<b>Path:</b> /countries/name/pan <br>
<b>Returns:</b> <br>
```
{
  "status": "success",
  "statusCode": 200,
  "data": [
    {
      "name": "Japan",
      "capital": "Tokyo",
      "phone_code": "81",
      "currency_name": "Japanese yen",
      "region": "Asia"
    },
    {
      "name": "Panama",
      "capital": "Panama City",
      "phone_code": "507",
      "currency_name": "Panamanian balboa",
      "region": "Americas"
    }
  ]
}
```
<b>Path:</b> /countries/phonecode/56 <br>
<b>Returns:</b> <br>
```
{
  "status": "success",
  "statusCode": 200,
  "data": [
    {
      "name": "Chile",
      "capital": "Santiago",
      "phone_code": "56",
      "currency_name": "Chilean peso",
      "region": "Americas"
    }
  ]
}
```


I got the DB from: https://github.com/dr5hn/countries-states-cities-database?tab=readme-ov-file <br>

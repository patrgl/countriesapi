<h1>Description</h1>

I made a simple API that returns country information in an organized JSON format in order to practice Golang. <br>
This is meant to be used as an example and a starting off point.

<h1>Examples: </h1>

<b>Path:</b> /countries/name/pan <br>
<b>Returns:</b> <br>
{<br>
  "status": "success",<br>
  "statusCode": 200,<br>
  "data": [<br>
    {<br>
      "name": "Japan",<br>
      "capital": "Tokyo",<br>
      "phone_code": "81",<br>
      "currency_name": "Japanese yen",<br>
      "region": "Asia"<br>
    },<br>
    {<br>
      "name": "Panama",<br>
      "capital": "Panama City",<br>
      "phone_code": "507",<br>
      "currency_name": "Panamanian balboa",<br>
      "region": "Americas"<br>
    }<br>
  ]<br>
}<br>
<br>
<b>Path:</b> /countries/phonecode/56<br>
<b>Returns:</b> <br>
{<br>
  "status": "success",<br>
  "statusCode": 200,<br>
  "data": [<br>
    {<br>
      "name": "Chile",<br>
      "capital": "Santiago",<br>
      "phone_code": "56",<br>
      "currency_name": "Chilean peso",<br>
      "region": "Americas"<br>
    }<br>
  ]<br>
}<br>
<br>
<br>
<br>
I got the DB from: https://github.com/dr5hn/countries-states-cities-database?tab=readme-ov-file <br>
